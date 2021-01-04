package hikvision

import (
	"context"
	"errors"

	"github.com/kihamo/boggart/components/boggart/tasks"
	"github.com/kihamo/boggart/providers/hikvision/client/content_manager"
	"github.com/kihamo/boggart/providers/hikvision/client/ptz"
	"github.com/kihamo/boggart/providers/hikvision/client/system"
	"go.uber.org/multierr"
)

func (b *Bind) Tasks() []tasks.Task {
	list := []tasks.Task{
		tasks.NewTask().
			WithName("updater").
			WithHandler(
				b.Workers().WrapTaskIsOnline(
					tasks.HandlerWithTimeout(
						tasks.HandlerFuncFromShortToLong(b.taskUpdater), b.config.UpdaterTimeout,
					),
				),
			).
			WithSchedule(tasks.ScheduleWithDuration(nil, b.config.UpdaterInterval)),
	}

	if b.config.PTZEnabled {
		list = append(list,
			tasks.NewTask().
				WithName("ptz").
				WithHandler(
					tasks.HandlerWithTimeout(
						tasks.HandlerFunc(b.taskPTZ), b.config.PTZTimeout,
					),
				).
				WithSchedule(
					tasks.ScheduleWithControl(
						tasks.ScheduleWithDuration(nil, b.config.PTZInterval),
					),
				),
		)
	}

	return list
}

func (b *Bind) taskPTZ(ctx context.Context, _ tasks.Meta, task tasks.Task) error {
	if !b.Meta().Status().IsStatusOnline() {
		return nil
	}

	b.mutex.RLock()
	channels := b.ptzChannels
	b.mutex.RUnlock()

	if channels == nil {
		return nil
	}

	stop := true

	if len(channels) != 0 {
		for id, channel := range channels {
			if !channel.Channel.Enabled {
				continue
			}

			if err := b.updateStatusByChannelID(ctx, id); err != nil {
				b.Logger().Error("Failed updated status",
					"serial_number", b.Meta().SerialNumber(),
					"channel", id,
					"error", err.Error(),
				)

				continue
			}

			stop = false
		}
	}

	if stop {
		if schedule, ok := task.Schedule().(*tasks.ScheduleControl); ok {
			schedule.Disable()
		}
	}

	return nil
}

func (b *Bind) taskUpdater(ctx context.Context) (err error) {
	sn := b.Meta().SerialNumber()

	// first initialization
	if sn == "" {
		deviceInfo, e := b.client.System.GetSystemDeviceInfo(system.NewGetSystemDeviceInfoParamsWithContext(ctx), nil)
		if e != nil {
			return e
		}

		if deviceInfo.Payload.SerialNumber == "" {
			return errors.New("device returns empty serial number")
		}

		if deviceInfo.Payload.MacAddress == "" {
			return errors.New("device returns empty MAC address")
		}

		b.Meta().SetSerialNumber(deviceInfo.Payload.SerialNumber)
		sn = deviceInfo.Payload.SerialNumber

		err = b.Meta().SetMACAsString(deviceInfo.Payload.MacAddress)
		if err != nil {
			return err
		}

		if b.config.EventsEnabled && b.config.EventsStreamingEnabled {
			b.startAlertStreaming()
		}

		if b.config.PTZEnabled {
			ptzChannels := make(map[uint64]PTZChannel)

			if list, e := b.client.Ptz.GetPtzChannels(ptz.NewGetPtzChannelsParamsWithContext(ctx), nil); e == nil {
				for _, channel := range list.Payload {
					ptzChannels[channel.ID] = PTZChannel{
						Channel: channel,
					}
				}
			} else {
				err = multierr.Append(err, e)
			}

			b.mutex.Lock()
			b.ptzChannels = ptzChannels
			b.mutex.Unlock()
		}

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicStateModel.Format(deviceInfo.Payload.SerialNumber), deviceInfo.Payload.Model); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicStateFirmwareVersion.Format(deviceInfo.Payload.SerialNumber), deviceInfo.Payload.FirmwareVersion); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicStateFirmwareReleasedDate.Format(deviceInfo.Payload.SerialNumber), deviceInfo.Payload.FirmwareReleasedDate); e != nil {
			err = multierr.Append(err, e)
		}
	}

	b.mutex.RLock()
	defer b.mutex.RUnlock()

	if status, e := b.client.System.GetStatus(system.NewGetStatusParamsWithContext(ctx), nil); e == nil {
		memoryUsage := int64(status.Payload.MemoryList[0].MemoryUsage) * MB
		memoryAvailable := int64(status.Payload.MemoryList[0].MemoryAvailable) * MB

		metricUpTime.With("serial_number", sn).Set(float64(status.Payload.DeviceUpTime))
		metricMemoryUsage.With("serial_number", sn).Set(float64(memoryUsage))
		metricMemoryAvailable.With("serial_number", sn).Set(float64(memoryAvailable))

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicStateUpTime.Format(sn), status.Payload.DeviceUpTime); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicStateMemoryUsage.Format(sn), memoryUsage); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicStateMemoryAvailable.Format(sn), memoryAvailable); e != nil {
			err = multierr.Append(err, e)
		}
	} else {
		err = multierr.Append(err, e)
	}

	if storage, e := b.client.ContentManager.GetStorage(content_manager.NewGetStorageParamsWithContext(ctx), nil); e == nil {
		for _, hdd := range storage.Payload.HddList {
			if hdd.Name == "" {
				continue
			}

			metricStorageUsage.With("serial_number", sn).With("name", hdd.Name).Set(float64((hdd.Capacity - hdd.FreeSpace) * MB))
			metricStorageAvailable.With("serial_number", sn).With("name", hdd.Name).Set(float64(hdd.FreeSpace * MB))

			if e := b.MQTT().PublishAsync(ctx, b.config.TopicStateHDDCapacity.Format(sn, hdd.ID), hdd.Capacity*MB); e != nil {
				err = multierr.Append(err, e)
			}

			if e := b.MQTT().PublishAsync(ctx, b.config.TopicStateHDDUsage.Format(sn, hdd.ID), (hdd.Capacity-hdd.FreeSpace)*MB); e != nil {
				err = multierr.Append(err, e)
			}

			if e := b.MQTT().PublishAsync(ctx, b.config.TopicStateHDDFree.Format(sn, hdd.ID), hdd.FreeSpace*MB); e != nil {
				err = multierr.Append(err, e)
			}
		}
	} else {
		err = multierr.Append(err, e)
	}

	return err
}
