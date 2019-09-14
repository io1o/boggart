package led_wifi

import (
	"context"
	"strconv"
	"strings"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/mqtt"
	"github.com/kihamo/boggart/providers/wifiled"
)

func (b *Bind) MQTTPublishes() []mqtt.Topic {
	return []mqtt.Topic{
		b.config.TopicStatePower,
		b.config.TopicStateColor,
		b.config.TopicStateColorHSV,
		b.config.TopicStateMode,
		b.config.TopicStateSpeed,
	}
}

func (b *Bind) MQTTSubscribers() []mqtt.Subscriber {
	return []mqtt.Subscriber{
		mqtt.NewSubscriber(b.config.TopicPower, 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b.Status, func(ctx context.Context, _ mqtt.Component, message mqtt.Message) error {
			if message.IsTrue() {
				return b.On(ctx)
			}

			return b.Off(ctx)
		})),
		mqtt.NewSubscriber(b.config.TopicColor, 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b.Status, func(ctx context.Context, _ mqtt.Component, message mqtt.Message) error {
			color, err := wifiled.ColorFromString(message.String())
			if err != nil {
				return err
			}

			return b.bulb.SetColorPersist(ctx, *color)
		})),
		mqtt.NewSubscriber(b.config.TopicMode, 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b.Status, func(ctx context.Context, _ mqtt.Component, message mqtt.Message) error {
			mode, err := wifiled.ModeFromString(strings.TrimSpace(message.String()))
			if err != nil {
				return err
			}

			state, err := b.bulb.State(ctx)
			if err != nil {
				return err
			}

			return b.bulb.SetMode(ctx, *mode, state.Speed)
		})),
		mqtt.NewSubscriber(b.config.TopicSpeed, 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b.Status, func(ctx context.Context, _ mqtt.Component, message mqtt.Message) error {
			speed, err := strconv.ParseInt(strings.TrimSpace(message.String()), 10, 64)
			if err != nil {
				return err
			}

			state, err := b.bulb.State(ctx)
			if err != nil {
				return err
			}

			return b.bulb.SetMode(ctx, state.Mode, uint8(speed))
		})),
	}
}
