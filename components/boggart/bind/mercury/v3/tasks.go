package v3

import (
	"context"
	"fmt"
	"time"

	"github.com/kihamo/boggart/components/boggart/tasks"
	"github.com/kihamo/boggart/providers/mercury/v3"
	"go.uber.org/multierr"
)

func (b *Bind) Tasks() []tasks.Task {
	return []tasks.Task{
		tasks.NewTask().
			WithName("updater").
			WithHandler(
				b.Workers().WrapTaskIsOnline(
					tasks.HandlerFuncFromShortToLong(b.taskUpdaterHandler),
				),
			).
			WithSchedule(tasks.ScheduleWithDuration(tasks.ScheduleNow(), b.config.UpdaterInterval)),
	}
}

func (b *Bind) taskUpdaterHandler(ctx context.Context) error {
	provider, err := b.Provider()
	if err != nil {
		return err
	}

	sn := b.Meta().SerialNumber()
	if sn == "" {
		var (
			makeDate time.Time
			version  string
		)

		sn, makeDate, version, _, err = provider.ForceReadParameters()
		if err != nil {
			return fmt.Errorf("execute method ForceReadParameters failed with error %v", err)
		}

		b.Meta().SetSerialNumber(sn)

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicMakeDate.Format(sn), makeDate); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicFirmwareVersion.Format(sn), version); e != nil {
			err = multierr.Append(err, e)
		}
	}

	if val, _, _, _, e := provider.ReadArray(v3.ArrayReset, nil, v3.Tariff1); e == nil {
		metricTariff.With("serial_number", sn).With("tariff", "1").Set(float64(val))

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicTariff.Format(sn), val); e != nil {
			err = multierr.Append(err, e)
		}
	} else {
		err = multierr.Append(err, fmt.Errorf("execute method ReadArray failed with error %v", e))
	}

	if p1, p2, p3, e := provider.Voltage(); e == nil {
		metricVoltage.With("serial_number", sn).With("phase", "1").Set(p1)
		metricVoltage.With("serial_number", sn).With("phase", "2").Set(p2)
		metricVoltage.With("serial_number", sn).With("phase", "3").Set(p3)

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicVoltagePhase1.Format(sn), p1); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicVoltagePhase2.Format(sn), p2); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicVoltagePhase3.Format(sn), p3); e != nil {
			err = multierr.Append(err, e)
		}
	} else {
		err = multierr.Append(err, fmt.Errorf("execute method Voltage failed with error %v", e))
	}

	if p1, p2, p3, e := provider.Amperage(); e == nil {
		metricAmperage.With("serial_number", sn).With("phase", "1").Set(p1)
		metricAmperage.With("serial_number", sn).With("phase", "2").Set(p2)
		metricAmperage.With("serial_number", sn).With("phase", "3").Set(p3)

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicAmperagePhase1.Format(sn), p1); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicAmperagePhase2.Format(sn), p2); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicAmperagePhase3.Format(sn), p3); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicAmperageTotal.Format(sn), p1+p2+p3); e != nil {
			err = multierr.Append(err, e)
		}
	} else {
		err = multierr.Append(err, fmt.Errorf("execute method Amperage failed with error %v", e))
	}

	if _, p1, p2, p3, e := provider.Power(v3.PowerNumberP); e == nil {
		metricPower.With("serial_number", sn).With("phase", "1").Set(p1)
		metricPower.With("serial_number", sn).With("phase", "2").Set(p2)
		metricPower.With("serial_number", sn).With("phase", "3").Set(p3)

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicPowerPhase1.Format(sn), p1); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicPowerPhase2.Format(sn), p2); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicPowerPhase3.Format(sn), p3); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTT().PublishAsync(ctx, b.config.TopicPowerTotal.Format(sn), p1+p2+p3); e != nil {
			err = multierr.Append(err, e)
		}
	} else {
		err = multierr.Append(err, fmt.Errorf("execute method Power failed with error %v", e))
	}

	return err
}
