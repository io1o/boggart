package internal

import (
	"container/list"
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	m "github.com/eclipse/paho.mqtt.golang"
	"github.com/kihamo/boggart/components/mqtt"
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/logging"
	"github.com/kihamo/shadow/components/tracing"
	"github.com/opentracing/opentracing-go/log"
)

type Component struct {
	lostConnections uint64

	application shadow.Application
	components  []shadow.Component
	config      config.Component
	logger      logging.Logger

	mutex         sync.RWMutex
	routes        []dashboard.Route
	client        m.Client
	subscriptions *list.List
}

func (c *Component) Name() string {
	return mqtt.ComponentName
}

func (c *Component) Version() string {
	return mqtt.ComponentVersion
}

func (c *Component) Dependencies() []shadow.Dependency {
	return []shadow.Dependency{
		{
			Name:     config.ComponentName,
			Required: true,
		},
		{
			Name: logging.ComponentName,
		},
		{
			Name: tracing.ComponentName,
		},
	}
}

func (c *Component) Init(a shadow.Application) error {
	c.application = a
	return nil
}

func (c *Component) Run(a shadow.Application, ready chan<- struct{}) (err error) {
	if c.components, err = a.GetComponents(); err != nil {
		return err
	}

	c.subscriptions = list.New()
	c.logger = logging.DefaultLazyLogger(c.Name())

	clientLogger := logging.NewLazyLogger(c.logger, c.Name()+".client")
	m.ERROR = NewMQTTLogger(clientLogger.Error, clientLogger.Errorf)
	m.CRITICAL = NewMQTTLogger(clientLogger.Error, clientLogger.Errorf)
	m.WARN = NewMQTTLogger(clientLogger.Warn, clientLogger.Warnf)
	m.DEBUG = NewMQTTLogger(clientLogger.Debug, clientLogger.Debugf)

	<-a.ReadyComponent(config.ComponentName)
	c.config = a.GetComponent(config.ComponentName).(config.Component)

	if err := c.initClient(); err != nil {
		return err
	}

	ready <- struct{}{}

	return c.initSubscribers()
}

func (c *Component) initClient() error {
	attempts := c.config.Int64(mqtt.ConfigConnectionAttempts)

	// auto reconnect
	//duration := c.config.Duration(mqtt.ConfigConnectionTimeout) + time.Second*30
	duration := time.Second * 5
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	opts := m.NewClientOptions()
	opts.Store = NewStore(c.logger)

	opts.ClientID = c.config.String(mqtt.ConfigClientID)
	opts.Username = c.config.String(mqtt.ConfigUsername)
	opts.Password = c.config.String(mqtt.ConfigPassword)
	opts.ConnectTimeout = c.config.Duration(mqtt.ConfigConnectionTimeout)
	opts.CleanSession = c.config.Bool(mqtt.ConfigClearSession)
	opts.ResumeSubs = c.config.Bool(mqtt.ConfigResumeSubs)
	opts.WriteTimeout = c.config.Duration(mqtt.ConfigWriteTimeout)

	opts.WillEnabled = c.config.Bool(mqtt.ConfigLWTEnabled)
	opts.WillTopic = c.config.String(mqtt.ConfigLWTTopic)
	opts.WillPayload = []byte(c.config.String(mqtt.ConfigLWTPayload))
	opts.WillQos = byte(c.config.Int(mqtt.ConfigLWTQOS))
	opts.WillRetained = c.config.Bool(mqtt.ConfigLWTRetained)

	opts.OnConnect = func(client m.Client) {
		cfg := client.OptionsReader()
		var mqttVersion string

		switch cfg.ProtocolVersion() {
		case 3:
			mqttVersion = "3.1"
		case 4:
			mqttVersion = "3.1.1"
		}

		c.logger.Debug("Connect to MQTT broker", "clientId", cfg.ClientID(), "protocol.version", mqttVersion)
		metricConnect.Inc()

		if atomic.LoadUint64(&c.lostConnections) == 0 {
			return
		}

		for _, sub := range c.Subscriptions() {
			topic := sub.Topic()
			qos := sub.QOS()

			if err := c.clientSubscribe(topic, qos, sub); err != nil {
				c.logger.Error("Resubscribe failed", "topic", topic, "qos", qos, "error", err.Error())
			} else {
				c.logger.Debug("Resubscribe success", "topic", topic, "qos", qos)
			}
		}
	}
	opts.OnConnectionLost = func(client m.Client, reason error) {
		atomic.AddUint64(&c.lostConnections, 1)
		c.logger.Error("Connection lost", "error", reason.Error(), "count", atomic.LoadUint64(&c.lostConnections))
		metricConnectionLost.Inc()
	}

	opts.DefaultPublishHandler = func(_ m.Client, message m.Message) {
		c.logger.Warn("Received that does not match any known subscriptions",
			"topic", message.Topic(),
			"qos", message.Qos(),
			"retained", message.Retained(),
		)
	}

	opts.Servers = make([]*url.URL, 0)
	for _, u := range strings.Split(c.config.String(mqtt.ConfigServers), ";") {
		if p, err := url.Parse(u); err == nil {
			opts.Servers = append(opts.Servers, p)
		}
	}

	for ; true; <-ticker.C {
		client := m.NewClient(opts)
		token := client.Connect()

		token.Wait()
		attempts--

		if err := token.Error(); err != nil {
			c.logger.Error("Init MQTT client failed with error " + err.Error())

			if attempts == 0 {
				c.logger.Error("Init MQTT client failed because exhausted number of connection attempts")
				return errors.New("exhausted number of connection attempts")
			}
		} else {
			c.mutex.Lock()
			c.client = client
			c.mutex.Unlock()

			break
		}
	}

	return nil
}

func (c *Component) initSubscribers() error {
	for _, component := range c.components {
		if subscribers, ok := component.(mqtt.HasSubscribers); ok {
			for _, subscriber := range subscribers.MQTTSubscribers() {
				if err := c.SubscribeSubscriber(subscriber); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (c *Component) Shutdown() (err error) {
	c.mutex.RLock()
	client := c.client
	c.mutex.RUnlock()

	if client != nil {
		if c.config.Bool(mqtt.ConfigLWTEnabled) {
			token := client.Publish(
				c.config.String(mqtt.ConfigLWTTopic),
				byte(c.config.Int(mqtt.ConfigLWTQOS)),
				c.config.Bool(mqtt.ConfigLWTRetained),
				[]byte(c.config.String(mqtt.ConfigLWTPayload)))
			token.Wait()

			err = token.Error()
		}

		client.Disconnect(250)
	}

	return err
}

func (c *Component) Client() m.Client {
	c.mutex.RLock()
	client := c.client
	c.mutex.RUnlock()

	if client != nil {
		return client
	}

	<-c.application.ReadyComponent(c.Name())

	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.client
}

func (c *Component) clientSubscribe(topic string, qos byte, subscription *mqtt.Subscription) error {
	client := c.Client()
	if client == nil {
		return errors.New("can't initialize client of MQTT")
	}

	// wrap tracing
	callback := func(client m.Client, message m.Message) {
		span, ctx := tracing.StartSpanFromContext(context.Background(), c.Name(), "subscribe_callback")
		span = span.SetTag("topic", message.Topic())
		defer span.Finish()

		span.LogFields(
			log.Int("qos", int(message.Qos())),
			log.String("payload", string(message.Payload())),
			log.Bool("retained", message.Retained()),
			log.String("topic.subscribe", topic),
		)

		msg := newMessage(message)

		// в отдельной рутине, так как если зависнет хендлер клиент MQTT не сделает ack на сообщение
		go func() {
			if err := subscription.Callback(ctx, c, msg); err != nil {
				metricSubscriberCalls.With("status", "failure", "topic", topic).Inc()

				tracing.SpanError(span, err)

				r := "0"
				if message.Retained() {
					r = "1"
				}

				c.logger.Error(
					"Call MQTT subscriber failed",
					"error", err.Error(),
					"topic.subscribe", topic,
					"topic.call", message.Topic(),
					"qos", strconv.Itoa(int(qos)),
					"retained", r,
					"payload", msg.String(),
				)
			} else {
				metricSubscriberCalls.With("status", "success", "topic", topic).Inc()
			}
		}()
	}

	token := client.Subscribe(topic, qos, callback)
	token.Wait()

	err := token.Error()
	if err == nil {
		metricSubscribe.With("status", "success", "topic", topic).Inc()
	} else {
		metricSubscribe.With("status", "failure", "topic", topic).Inc()
	}

	return err
}

func (c *Component) Publish(ctx context.Context, topic string, qos byte, retained bool, payload interface{}) (err error) {
	// TODO: cache topics
	payload = c.convertPayload(payload)

	span, _ := tracing.StartSpanFromContext(ctx, c.Name(), "mqtt_publish")
	span = span.SetTag("topic", topic)
	defer span.Finish()

	span.LogFields(
		log.Int("qos", int(qos)),
		log.String("payload", fmt.Sprintf("%v", payload)),
		log.Bool("retained", retained),
	)

	client := c.Client()
	if client != nil {
		token := client.Publish(topic, qos, retained, payload)
		token.Wait()

		err = token.Error()
	} else {
		err = errors.New("can't initialize client of MQTT")
	}

	if err != nil {
		metricPublish.With("status", "failure").Inc()

		tracing.SpanError(span, err)

		r := "0"
		if retained {
			r = "1"
		}

		c.logger.Error(
			"Publish MQTT topic failed",
			"error", err.Error(),
			"topic", topic,
			"qos", strconv.Itoa(int(qos)),
			"retained", r,
			"payload", fmt.Sprintf("%v", payload),
		)
	} else {
		metricPublish.With("status", "success").Inc()
	}

	return err
}

func (c *Component) PublishAsync(ctx context.Context, topic string, qos byte, retained bool, payload interface{}) {
	go func() {
		_ = c.Publish(ctx, topic, qos, retained, payload)
	}()
}

func (c *Component) Unsubscribe(topic string) error {
	client := c.Client()
	if client == nil {
		return errors.New("can't initialize client of MQTT")
	}

	if token := client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	for element := c.subscriptions.Front(); element != nil; element = element.Next() {
		subscription := element.Value.(*mqtt.Subscription)

		for _, subscriber := range subscription.Subscribers() {
			if subscriber.Topic() == topic {
				subscription.RemoveSubscriber(subscriber)
			}
		}

		if subscription.Len() == 0 {
			topic := subscription.Topic()

			token := client.Unsubscribe(topic)
			token.Wait()
			err := token.Error()

			if err == nil {
				c.subscriptions.Remove(element)
				c.logger.Debug("Unsubscribe", "topic", topic)
			} else {
				c.logger.Error("Unsubscribe failed", "error", err.Error())
				return err
			}
		}
	}

	return nil
}

func (c *Component) UnsubscribeSubscriber(subscriber mqtt.Subscriber) error {
	client := c.Client()
	if client == nil {
		return errors.New("can't initialize client of MQTT")
	}

	var element *list.Element

	c.mutex.RLock()
	defer c.mutex.RUnlock()

	for element = c.subscriptions.Front(); element != nil; element = element.Next() {
		subscription := element.Value.(*mqtt.Subscription)

		if result := subscription.RemoveSubscriber(subscriber); result {
			if subscription.Len() == 0 {
				topic := subscriber.Topic()

				token := client.Unsubscribe(topic)
				token.Wait()
				err := token.Error()

				if err == nil {
					c.subscriptions.Remove(element)
					c.logger.Debug("Unsubscribe", "topic", topic)
				} else {
					c.logger.Error("Unsubscribe failed", "error", err.Error())

					// fallback
					subscription.AddSubscriber(subscriber)

					return err
				}
			}

			break
		}
	}

	return nil
}

func (c *Component) UnsubscribeSubscribers(subscribers []mqtt.Subscriber) error {
	for _, subscriber := range subscribers {
		if err := c.UnsubscribeSubscriber(subscriber); err != nil {
			return err
		}
	}

	return nil
}

func (c *Component) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) (mqtt.Subscriber, error) {
	subscriber := mqtt.NewSubscriber(topic, qos, callback)
	if err := c.SubscribeSubscriber(subscriber); err != nil {
		return nil, err
	}

	return subscriber, nil
}

func (c *Component) SubscribeSubscriber(subscriber mqtt.Subscriber) error {
	var (
		element      *list.Element
		subscription *mqtt.Subscription
	)

	topic := subscriber.Topic()
	qos := subscriber.QOS()

	//  ищем подходящие подписки
	c.mutex.RLock()
	for element = c.subscriptions.Front(); element != nil; element = element.Next() {
		s := element.Value.(*mqtt.Subscription)
		if s.Match(topic) {
			s.AddSubscriber(subscriber)
			subscription = s
			break
		}
	}
	c.mutex.RUnlock()

	// если подписка не найдена, то создаем новую
	if subscription == nil {
		subscription = mqtt.NewSubscription(subscriber)
	}

	if err := c.clientSubscribe(topic, qos, subscription); err != nil {
		return err
	}

	c.mutex.Lock()
	if element != nil {
		element.Value = subscription
	} else {
		c.subscriptions.PushBack(subscription)
	}
	c.mutex.Unlock()

	c.logger.Debug("Subscribe success", "topic", topic, "qos", qos)

	return nil
}

func (c *Component) Subscriptions() []*mqtt.Subscription {
	c.mutex.RLock()
	subscriptions := make([]*mqtt.Subscription, 0, c.subscriptions.Len())
	for e := c.subscriptions.Front(); e != nil; e = e.Next() {
		subscriptions = append(subscriptions, e.Value.(*mqtt.Subscription))
	}
	c.mutex.RUnlock()

	return subscriptions
}

func (c *Component) convertPayload(payload interface{}) interface{} {
	switch value := payload.(type) {
	case string, []byte:
		// skip
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 64)
	case int64:
		return strconv.FormatInt(value, 10)
	case int32:
		return strconv.FormatInt(int64(value), 10)
	case int16:
		return strconv.FormatInt(int64(value), 10)
	case int8:
		return strconv.FormatInt(int64(value), 10)
	case int:
		return strconv.FormatInt(int64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case bool:
		if value {
			return PayloadTrue
		}

		return PayloadFalse
	case time.Time:
		return value.Format(time.RFC3339)
	case *time.Time:
		return value.Format(time.RFC3339)
	case time.Duration:
		return strconv.FormatFloat(value.Seconds(), 'f', -1, 64)
	case *time.Duration:
		return strconv.FormatFloat(value.Seconds(), 'f', -1, 64)
	default:
		return fmt.Sprintf("%s", payload)
	}

	return payload
}
