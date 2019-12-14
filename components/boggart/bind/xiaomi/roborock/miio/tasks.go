package miio

import (
	"context"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/providers/xiaomi/miio/devices/vacuum"
	"github.com/kihamo/go-workers"
	"github.com/kihamo/go-workers/task"
	"go.uber.org/multierr"
)

func (b *Bind) Tasks() []workers.Task {
	taskLiveness := task.NewFunctionTask(b.taskLiveness)
	taskLiveness.SetTimeout(b.config.LivenessTimeout)
	taskLiveness.SetRepeats(-1)
	taskLiveness.SetRepeatInterval(b.config.LivenessInterval)
	taskLiveness.SetName("liveness-" + b.config.Host)

	taskState := b.WrapTaskIsOnline(b.taskUpdater)
	taskState.SetTimeout(b.config.UpdaterTimeout)
	taskState.SetRepeats(-1)
	taskState.SetRepeatInterval(b.config.UpdaterInterval)
	taskState.SetName("updater-" + b.config.Host)

	return []workers.Task{
		taskLiveness,
		taskState,
	}
}

func (b *Bind) taskLiveness(ctx context.Context) (interface{}, error) {
	sn, err := b.device.SerialNumber(ctx)
	if err != nil {
		b.UpdateStatus(boggart.BindStatusOffline)
		return nil, nil
	}

	if b.IsStatusOnline() {
		return nil, nil
	}

	b.UpdateStatus(boggart.BindStatusOnline)

	if b.SerialNumber() == "" {
		b.SetSerialNumber(sn)

		b.taskUpdater(ctx)
	}

	return nil, nil
}

func (b *Bind) taskUpdater(ctx context.Context) error {
	sn := b.SerialNumber()
	if sn == "" {
		return nil
	}

	var err error

	// only statistics
	status, e := b.device.Status(ctx)
	if e == nil {
		metricBattery.With("serial_number", sn).Set(float64(status.Battery))

		if e := b.MQTTPublishAsync(ctx, b.config.TopicBattery.Format(sn), status.Battery); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTTPublishAsync(ctx, b.config.TopicCleanArea.Format(sn), status.CleanArea); e != nil {
			err = multierr.Append(err, e)
		}

		if status.CleanTime > 0 {
			if e := b.MQTTPublishAsync(ctx, b.config.TopicCleanTime.Format(sn), status.CleanTime); e != nil {
				err = multierr.Append(err, e)
			}
		}
	} else {
		err = multierr.Append(err, e)
	}

	consumables, e := b.device.Consumables(ctx)
	if e == nil {
		if consumable, ok := consumables[vacuum.ConsumableFilter]; ok {
			if e := b.MQTTPublishAsync(ctx, b.config.TopicConsumableFilter.Format(sn), consumable); e != nil {
				err = multierr.Append(err, e)
			}
		}

		if consumable, ok := consumables[vacuum.ConsumableBrushMain]; ok {
			if e := b.MQTTPublishAsync(ctx, b.config.TopicConsumableBrushMain.Format(sn), consumable); e != nil {
				err = multierr.Append(err, e)
			}
		}

		if consumable, ok := consumables[vacuum.ConsumableBrushSide]; ok {
			if e := b.MQTTPublishAsync(ctx, b.config.TopicConsumableBrushSide.Format(sn), consumable); e != nil {
				err = multierr.Append(err, e)
			}
		}

		if consumable, ok := consumables[vacuum.ConsumableSensor]; ok {
			if e := b.MQTTPublishAsync(ctx, b.config.TopicConsumableSensor.Format(sn), consumable); e != nil {
				err = multierr.Append(err, e)
			}
		}
	} else {
		err = multierr.Append(err, e)
	}

	if e := b.updateStatus(ctx); e != nil {
		err = multierr.Append(err, e)
	}

	if e := b.updateFanPower(ctx); e != nil {
		err = multierr.Append(err, e)
	}

	if e := b.updateVolume(ctx); e != nil {
		err = multierr.Append(err, e)
	}

	return err
}
