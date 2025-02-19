package hilink

import (
	"context"
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/kihamo/boggart/atomic"
	"github.com/kihamo/boggart/components/boggart/di"
	"github.com/kihamo/boggart/protocols/swagger"
	"github.com/kihamo/boggart/providers/hilink"
	"github.com/kihamo/boggart/providers/hilink/client/device"
	"github.com/kihamo/boggart/providers/hilink/client/sms"
	"github.com/kihamo/boggart/providers/hilink/client/ussd"
	"github.com/kihamo/boggart/providers/hilink/models"
)

var (
	cmdRegexp = regexp.MustCompile(`^(exec)?\s*(?P<command>.*)$`)
)

type Bind struct {
	di.ConfigBind
	di.LoggerBind
	di.MetaBind
	di.MetricsBind
	di.MQTTBind
	di.ProbesBind
	di.WorkersBind

	client                    *hilink.Client
	operator                  *atomic.String
	limitInternetTrafficIndex *atomic.Int64
	simStatus                 *atomic.Uint32
}

func (b *Bind) config() *Config {
	return b.Config().Bind().(*Config)
}

func (b *Bind) Run() error {
	cfg := b.config()

	b.operator.Set("")
	b.limitInternetTrafficIndex.Set(0)
	b.simStatus.Set(0)

	b.Meta().SetLink(&cfg.Address.URL)

	b.client = hilink.New(cfg.Address.Host, cfg.Debug, swagger.NewLogger(
		func(message string) {
			b.Logger().Info(message)
		},
		func(message string) {
			b.Logger().Debug(message)
		}))

	return nil
}

func (b *Bind) SMS(ctx context.Context, content string, phones ...string) error {
	if len(content) == 0 {
		return errors.New("sms content is empty")
	}

	if len(phones) == 0 {
		return errors.New("sms phones list is empty")
	}

	params := sms.NewSendSMSParamsWithContext(ctx)
	params.Request.Index = -1
	params.Request.Reserved = 1
	params.Request.Date = time.Now().Format(hilink.TimeFormat)
	params.Request.Phones = phones
	params.Request.Content = content
	params.Request.Length = int64(len(content))

	_, err := b.client.Sms.SendSMS(params)
	return err
}

func (b *Bind) USSD(ctx context.Context, content string) (string, error) {
	if content == "" {
		return "", nil
	}

	params := ussd.NewSendUSSDParamsWithContext(ctx).
		WithRequest(&models.USSD{
			Content: content,
		})

	_, err := b.client.Ussd.SendUSSD(params)
	if err != nil {
		return "", err
	}

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()

		default:
			response, err := b.client.Ussd.GetUSSD(ussd.NewGetUSSDParamsWithContext(ctx))
			if err == nil && response.Payload.Content != "" {
				return response.Payload.Content, nil
			}

			time.Sleep(time.Second)
		}
	}
}

func (b *Bind) Balance(ctx context.Context) (float64, error) {
	op, err := b.operatorSettings()
	if err != nil {
		return -1, err
	}

	content, err := b.USSD(ctx, op.BalanceUSSD)
	if err != nil {
		return -1, err
	}

	match := op.BalanceRegexp.FindStringSubmatch(content)

	if len(match) > 0 {
		for i, name := range op.BalanceRegexp.SubexpNames() {
			if name == "value" {
				return strconv.ParseFloat(match[i], 64)
			}
		}
	}

	return 0, errors.New("balance not found")
}

func (b *Bind) operatorSettings() (*operator, error) {
	label := b.operator.Load()

	switch strings.ToLower(label) {
	case "tele2", "tele2 ru":
		return operatorTele2, nil
	}

	return nil, errors.New("operator " + label + " settings isn't found")
}

func (b *Bind) checkSpecialSMS(ctx context.Context, smsItem *models.SMSListMessagesItems0) bool {
	op, err := b.operatorSettings()
	if err != nil {
		return false
	}

	sn := b.Meta().SerialNumber()
	cfg := b.config()

	// limit traffic
	match := op.SMSLimitTrafficRegexp.FindStringSubmatch(smsItem.Content)
	if len(match) > 0 {
		var (
			value float64
			found bool
		)

		for i, name := range op.SMSLimitTrafficRegexp.SubexpNames() {
			if !strings.HasPrefix(name, "value") {
				continue
			}

			found = true

			if v, err := strconv.ParseFloat(match[i], 64); err == nil {
				value += v
			}
		}

		if found {
			if smsItem.Index > b.limitInternetTrafficIndex.Load() {
				value *= op.SMSLimitTrafficFactor

				metricLimitInternetTraffic.With("serial_number", sn).Set(value)
				b.MQTT().PublishAsync(ctx, cfg.TopicLimitInternetTraffic.Format(sn), uint64(value))

				b.limitInternetTrafficIndex.Set(smsItem.Index)
			}

			return true
		}
	}

	// special commands
	if cfg.SMSCommandsEnabled {
		match = cmdRegexp.FindStringSubmatch(smsItem.Content)
		if len(match) > 0 {
			for i, name := range cmdRegexp.SubexpNames() {
				if name == "command" {
					// TODO: надо игнорировать номера операторов

					cmd := strings.ToLower(strings.TrimSpace(match[i]))

					switch cmd {
					case "reboot", "status":
						// поддерживаемые команды
					default:
						b.Logger().Warnf("Unknown command %s from phone number %s ", cmd, smsItem.Phone)
						return false
					}

					// нет разрешенных телефонов с которых можно принимать команды, но команда найдена
					if len(cfg.SMSCommandsAllowedPhones) == 0 {
						return true
					}

					var allowed bool

					for _, phone := range cfg.SMSCommandsAllowedPhones {
						if strings.Compare(phone, smsItem.Phone) == 0 {
							allowed = true
							break
						}
					}

					// выполнение команд с номера не разрешено
					if !allowed {
						b.Logger().Warnf("Unauthorized execute command %s from phone number %s ", cmd, smsItem.Phone)
						return true
					}

					switch cmd {
					case "reboot":
						params := device.NewDeviceControlParamsWithContext(ctx)
						params.Request.Control = 1

						if _, err := b.client.Device.DeviceControl(params); err != nil {
							b.Logger().Errorf("Reboot failed with error %v", err)
						}
					case "status":
						balance, err := b.Balance(ctx)
						if err == nil {
							err = b.SMS(ctx,
								"A'm OK. My balance is "+strconv.FormatFloat(balance, 'f', -1, 64),
								smsItem.Phone)
						}

						if err != nil {
							b.Logger().Errorf("Send status SMS failed with error %v", err)
						}
					}

					return true
				}
			}
		}
	}

	return false
}

func parseInt(s string) (int64, error) {
	s = strings.TrimFunc(s, func(r rune) bool {
		return !(unicode.Is(unicode.Number, r) || unicode.Is(unicode.Pd, r))
	})

	return strconv.ParseInt(s, 10, 64)
}
