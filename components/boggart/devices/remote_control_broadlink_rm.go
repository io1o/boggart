package devices

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/providers/broadlink"
	"github.com/kihamo/boggart/components/mqtt"
	"github.com/kihamo/go-workers"
	"github.com/kihamo/go-workers/task"
	"github.com/kihamo/shadow/components/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

const (
	RemoteControlBroadlinkRMCaptureDuration = time.Second * 15

	RemoteControlBroadlinkRMMQTTTopicCommand         mqtt.Topic = boggart.ComponentName + "/remote-control/+/command"
	RemoteControlBroadlinkRMMQTTTopicRawCount        mqtt.Topic = boggart.ComponentName + "/remote-control/+/command/raw/count"
	RemoteControlBroadlinkRMMQTTTopicRaw             mqtt.Topic = boggart.ComponentName + "/remote-control/+/command/raw"
	RemoteControlBroadlinkRMMQTTTopicIRCount         mqtt.Topic = boggart.ComponentName + "/remote-control/+/command/ir/count"
	RemoteControlBroadlinkRMMQTTTopicIR              mqtt.Topic = boggart.ComponentName + "/remote-control/+/command/ir"
	RemoteControlBroadlinkRMMQTTTopicRF315mhz        mqtt.Topic = boggart.ComponentName + "/remote-control/+/command/rf315mhz"
	RemoteControlBroadlinkRMMQTTTopicRF433mhz        mqtt.Topic = boggart.ComponentName + "/remote-control/+/command/rf433mhz"
	RemoteControlBroadlinkRMMQTTTopicCapture         mqtt.Topic = boggart.ComponentName + "/remote-control/+/command/capture"
	RemoteControlBroadlinkRMMQTTTopicCaptureState    mqtt.Topic = boggart.ComponentName + "/remote-control/+/command/capture/state"
	RemoteControlBroadlinkRMMQTTTopicCaptureIR       mqtt.Topic = boggart.ComponentName + "/remote-control/+/command/capture/ir"
	RemoteControlBroadlinkRMMQTTTopicCaptureRF315mhz mqtt.Topic = boggart.ComponentName + "/remote-control/+/command/capture/rf315mhz"
	RemoteControlBroadlinkRMMQTTTopicCaptureRF433mhz mqtt.Topic = boggart.ComponentName + "/remote-control/+/command/capture/rf433mhz"
)

type BroadlinkRMRemoteControl struct {
	boggart.DeviceBase
	boggart.DeviceSerialNumber
	boggart.DeviceMQTT

	provider *broadlink.RMProPlus
}

func NewBroadlinkRMRemoteControl(provider *broadlink.RMProPlus, m mqtt.Component) *BroadlinkRMRemoteControl {
	device := &BroadlinkRMRemoteControl{
		provider: provider,
	}
	device.Init()
	device.SetSerialNumber(provider.MAC().String())
	device.SetDescription("Socket of Broadlink")

	return device
}

func (d *BroadlinkRMRemoteControl) SetMQTTClient(client mqtt.Component) {
	d.DeviceMQTT.SetMQTTClient(client)

	mac := strings.Replace(d.provider.MAC().String(), ":", "-", -1)
	client.Publish(context.Background(), RemoteControlBroadlinkRMMQTTTopicCaptureState.Format(mac), 2, true, "0")
}

func (d *BroadlinkRMRemoteControl) Types() []boggart.DeviceType {
	return []boggart.DeviceType{
		boggart.DeviceTypeRemoteControl,
	}
}

func (d *BroadlinkRMRemoteControl) Tasks() []workers.Task {
	taskLiveness := task.NewFunctionTask(d.taskLiveness)
	taskLiveness.SetTimeout(time.Second * 5)
	taskLiveness.SetRepeats(-1)
	taskLiveness.SetRepeatInterval(time.Second * 30)
	taskLiveness.SetName("remote-control-broadlink-rm-liveness")

	return []workers.Task{
		taskLiveness,
	}
}

func (d *BroadlinkRMRemoteControl) taskLiveness(ctx context.Context) (interface{}, error) {
	// TODO:
	d.UpdateStatus(boggart.DeviceStatusOnline)

	return nil, nil
}

func (d *BroadlinkRMRemoteControl) MQTTTopics() []mqtt.Topic {
	return []mqtt.Topic{
		RemoteControlBroadlinkRMMQTTTopicCommand,
		RemoteControlBroadlinkRMMQTTTopicRawCount,
		RemoteControlBroadlinkRMMQTTTopicRaw,
		RemoteControlBroadlinkRMMQTTTopicIRCount,
		RemoteControlBroadlinkRMMQTTTopicIR,
		RemoteControlBroadlinkRMMQTTTopicRF315mhz,
		RemoteControlBroadlinkRMMQTTTopicRF433mhz,
		RemoteControlBroadlinkRMMQTTTopicCapture,
		RemoteControlBroadlinkRMMQTTTopicCaptureState,
		RemoteControlBroadlinkRMMQTTTopicCaptureIR,
		RemoteControlBroadlinkRMMQTTTopicCaptureRF315mhz,
		RemoteControlBroadlinkRMMQTTTopicCaptureRF433mhz,
	}
}

func (d *BroadlinkRMRemoteControl) MQTTSubscribers() []mqtt.Subscriber {
	type codeRequest struct {
		Code  string `json:"code"`
		Count int    `json:"count"`
	}

	// создаем таймер который сразу завершается, чтобы получить проиницилизированный таймер
	captureFlush := make(chan struct{}, 1)
	captureDone := make(chan struct{}, 1)

	captureTimer := time.NewTimer(0)
	<-captureTimer.C

	mac := strings.Replace(d.provider.MAC().String(), ":", "-", -1)

	return []mqtt.Subscriber{
		mqtt.NewSubscriber(RemoteControlBroadlinkRMMQTTTopicRawCount.Format(mac), 0, d.wrapMQTTSubscriber("command_raw_count",
			func(ctx context.Context, _ mqtt.Component, message mqtt.Message) error {
				var request codeRequest

				err := json.Unmarshal(message.Payload(), &request)
				if err == nil {
					if span := opentracing.SpanFromContext(ctx); span != nil {
						span.LogFields(
							log.String("code", request.Code),
							log.Int("count", request.Count),
						)
					}

					err = d.provider.SendRemoteControlCodeRawAsString(request.Code, request.Count)
				}

				return err
			})),
		mqtt.NewSubscriber(RemoteControlBroadlinkRMMQTTTopicRaw.Format(mac), 0, d.wrapMQTTSubscriber("command_raw",
			func(_ context.Context, _ mqtt.Component, message mqtt.Message) error {
				return d.provider.SendRemoteControlCodeRawAsString(string(message.Payload()), 0)
			})),
		mqtt.NewSubscriber(RemoteControlBroadlinkRMMQTTTopicIRCount.Format(mac), 0, d.wrapMQTTSubscriber("command_ir_count",
			func(ctx context.Context, _ mqtt.Component, message mqtt.Message) error {
				var request codeRequest

				err := json.Unmarshal(message.Payload(), &request)
				if err == nil {
					if span := opentracing.SpanFromContext(ctx); span != nil {
						span.LogFields(
							log.String("code", request.Code),
							log.Int("count", request.Count),
						)
					}

					err = d.provider.SendIRRemoteControlCodeAsString(request.Code, request.Count)
				}

				return err
			})),
		mqtt.NewSubscriber(RemoteControlBroadlinkRMMQTTTopicIR.Format(mac), 0, d.wrapMQTTSubscriber("command_ir",
			func(_ context.Context, _ mqtt.Component, message mqtt.Message) error {
				return d.provider.SendIRRemoteControlCodeAsString(string(message.Payload()), 0)
			})),
		mqtt.NewSubscriber(RemoteControlBroadlinkRMMQTTTopicRF315mhz.Format(mac), 0, d.wrapMQTTSubscriber("command_rf315mhz",
			func(_ context.Context, _ mqtt.Component, message mqtt.Message) error {
				return d.provider.SendRF315MhzRemoteControlCodeAsString(string(message.Payload()), 0)
			})),
		mqtt.NewSubscriber(RemoteControlBroadlinkRMMQTTTopicRF433mhz.Format(mac), 0, d.wrapMQTTSubscriber("command_rf433mhz",
			func(_ context.Context, _ mqtt.Component, message mqtt.Message) error {
				return d.provider.SendRF433MhzRemoteControlCodeAsString(string(message.Payload()), 0)
			})),
		mqtt.NewSubscriber(RemoteControlBroadlinkRMMQTTTopicCapture.Format(mac), 0, d.wrapMQTTSubscriber("command_capture_start",
			func(ctx context.Context, client mqtt.Component, message mqtt.Message) error {
				if string(message.Payload()) != "1" {
					return nil
				}

				if err := d.provider.StartCaptureRemoteControlCode(); err != nil {
					return err
				}

				// завершаем предыдущие запуски
				if captureTimer.Reset(RemoteControlBroadlinkRMCaptureDuration) {
					captureDone <- struct{}{}
				} else {
					for len(captureDone) > 0 {
						<-captureDone
					}
				}

				// вычитываем каналы, если успели забиться
				for len(captureFlush) > 0 {
					<-captureFlush
				}

				for len(captureTimer.C) > 0 {
					<-captureTimer.C
				}

				// стартуем новую запись
				err := d.MQTTPublish(ctx, RemoteControlBroadlinkRMMQTTTopicCaptureState.Format(mac), 2, true, "1")
				if err != nil {
					return err
				}

				select {
				case <-captureFlush:
				case <-captureTimer.C:

				case <-captureDone:
					return nil
				}

				d.MQTTPublishAsync(ctx, RemoteControlBroadlinkRMMQTTTopicCaptureState.Format(mac), 2, true, "0")

				remoteType, code, err := d.provider.ReadCapturedRemoteControlCodeAsString()
				if err != nil {
					if err != broadlink.ErrSignalNotCaptured {
						return err
					}

					return nil
				}

				var topicCaptureCode string

				switch remoteType {
				case broadlink.RemoteIR:
					topicCaptureCode = RemoteControlBroadlinkRMMQTTTopicCaptureIR.Format(mac)
				case broadlink.RemoteRF315Mhz:
					topicCaptureCode = RemoteControlBroadlinkRMMQTTTopicCaptureRF315mhz.Format(mac)
				case broadlink.RemoteRF433Mhz:
					topicCaptureCode = RemoteControlBroadlinkRMMQTTTopicCaptureRF433mhz.Format(mac)
				}

				if topicCaptureCode != "" {
					if err = d.MQTTPublish(ctx, topicCaptureCode, 0, false, code); err != nil {
						return err
					}
				}

				return nil
			})),
		mqtt.NewSubscriber(RemoteControlBroadlinkRMMQTTTopicCapture.Format(mac), 0, d.wrapMQTTSubscriber("command_capture_stop",
			func(ctx context.Context, client mqtt.Component, message mqtt.Message) error {
				if string(message.Payload()) != "0" {
					return nil
				}

				if len(captureFlush) == 0 {
					captureFlush <- struct{}{}
				}

				return nil
			})),
	}
}

func (d *BroadlinkRMRemoteControl) wrapMQTTSubscriber(operationName string, fn func(context.Context, mqtt.Component, mqtt.Message) error) mqtt.MessageHandler {
	return func(ctx context.Context, client mqtt.Component, message mqtt.Message) {
		if d.Status() != boggart.DeviceStatusOnline {
			return
		}

		span, ctx := tracing.StartSpanFromContext(ctx, boggart.DeviceTypeRemoteControl.String(), operationName)
		span.LogFields(
			log.String("mac", d.provider.MAC().String()),
			log.String("ip", d.provider.Addr().String()))
		defer span.Finish()

		if err := fn(ctx, client, message); err != nil {
			tracing.SpanError(span, err)
		}
	}
}
