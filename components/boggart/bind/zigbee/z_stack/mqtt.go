package z_stack

import (
	"context"

	"github.com/kihamo/boggart/components/mqtt"
)

func (b *Bind) syncPermitJoin() {
	if sn := b.Meta().SerialNumber(); sn != "" {
		b.MQTT().PublishAsync(context.TODO(), b.config.TopicPermitJoinState.Format(sn), b.client.PermitJoinEnabled())
	}
}

func (b *Bind) MQTTSubscribers() []mqtt.Subscriber {
	return []mqtt.Subscriber{
		mqtt.NewSubscriber(b.config.TopicPermitJoin, 0, b.MQTT().WrapSubscribeDeviceIsOnline(
			func(ctx context.Context, _ mqtt.Component, message mqtt.Message) (err error) {
				if !b.MQTT().CheckSerialNumberInTopic(message.Topic(), 2) {
					return nil
				}

				if message.IsTrue() {
					err = b.client.PermitJoin(ctx, b.permitJoinDuration())
				} else {
					err = b.client.PermitJoinDisable(ctx)
				}

				return err
			})),
	}
}
