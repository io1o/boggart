package lg_webos

import (
	"context"
	"errors"
	"strconv"

	"github.com/ghthor/gowol"
	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/mqtt"
)

func (b *Bind) MQTTPublishes() []mqtt.Topic {
	return []mqtt.Topic{
		b.config.TopicStateMute,
		b.config.TopicStateVolume,
		b.config.TopicStateApplication,
		b.config.TopicStateChannelNumber,
		b.config.TopicStatePower,
	}
}

func (b *Bind) MQTTSubscribers() []mqtt.Subscriber {
	return []mqtt.Subscriber{
		mqtt.NewSubscriber(b.config.TopicApplication, 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b.Status, func(_ context.Context, _ mqtt.Component, message mqtt.Message) error {
			if !boggart.CheckSerialNumberInMQTTTopic(b, message.Topic(), 2) {
				return nil
			}

			client := b.Client()
			if client == nil {
				return errors.New("client isn't init")
			}

			_, err := client.ApplicationManagerLaunch(message.String(), nil)
			return err
		})),
		mqtt.NewSubscriber(b.config.TopicMute, 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b.Status, func(_ context.Context, _ mqtt.Component, message mqtt.Message) error {
			if !boggart.CheckSerialNumberInMQTTTopic(b, message.Topic(), 2) {
				return nil
			}

			client := b.Client()
			if client == nil {
				return errors.New("client isn't init")
			}

			return client.AudioSetMute(message.IsTrue())
		})),
		mqtt.NewSubscriber(b.config.TopicVolume, 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b.Status, func(_ context.Context, _ mqtt.Component, message mqtt.Message) error {
			if !boggart.CheckSerialNumberInMQTTTopic(b, message.Topic(), 2) {
				return nil
			}

			vol, err := strconv.ParseInt(message.String(), 10, 64)
			if err != nil {
				return err
			}

			client := b.Client()
			if client == nil {
				return errors.New("client isn't init")
			}

			return client.AudioSetVolume(int(vol))
		})),
		mqtt.NewSubscriber(b.config.TopicVolumeUp, 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b.Status, func(_ context.Context, _ mqtt.Component, message mqtt.Message) error {
			if !boggart.CheckSerialNumberInMQTTTopic(b, message.Topic(), 3) {
				return nil
			}

			client := b.Client()
			if client == nil {
				return errors.New("client isn't init")
			}

			return client.AudioVolumeUp()
		})),
		mqtt.NewSubscriber(b.config.TopicVolumeDown, 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b.Status, func(_ context.Context, _ mqtt.Component, message mqtt.Message) error {
			if !boggart.CheckSerialNumberInMQTTTopic(b, message.Topic(), 3) {
				return nil
			}

			client := b.Client()
			if client == nil {
				return errors.New("client isn't init")
			}

			return client.AudioVolumeDown()
		})),
		mqtt.NewSubscriber(b.config.TopicToast, 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b.Status, func(_ context.Context, _ mqtt.Component, message mqtt.Message) error {
			if !boggart.CheckSerialNumberInMQTTTopic(b, message.Topic(), 2) {
				return nil
			}

			return b.Toast(message.String())
		})),
		mqtt.NewSubscriber(b.config.TopicPower, 0, func(_ context.Context, _ mqtt.Component, message mqtt.Message) error {
			if !boggart.CheckSerialNumberInMQTTTopic(b, message.Topic(), 2) {
				return nil
			}

			if message.IsTrue() {
				return wol.MagicWake(b.SerialNumber(), "255.255.255.255")
			}

			if !b.IsStatusOnline() {
				return errors.New("bind isn't online")
			}

			client := b.Client()
			if client == nil {
				return errors.New("client isn't init")
			}

			return client.SystemTurnOff()
		}),
	}
}
