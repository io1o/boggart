package nut

import (
	"context"
	"strings"

	"github.com/kihamo/go-workers"
)

func (b *Bind) Tasks() []workers.Task {
	taskStateUpdater := b.Workers().WrapTaskIsOnline(b.taskUpdater)
	taskStateUpdater.SetRepeats(-1)
	taskStateUpdater.SetRepeatInterval(b.config.UpdaterInterval)
	taskStateUpdater.SetName("updater")

	return []workers.Task{
		taskStateUpdater,
	}
}

func (b *Bind) taskUpdater(ctx context.Context) error {
	variables, err := b.Variables()
	if err != nil {
		return err
	}

	sn := b.Meta().SerialNumber()

	if sn == "" {
		for _, v := range variables {
			if v.Name == "device.serial" {
				sn = strings.TrimSpace(v.Value.(string))
				b.Meta().SetSerialNumber(sn)
				break
			}
		}
	}

	if sn == "" {
		return nil
	}

	for _, v := range variables {
		switch v.Name {
		case "ups.load":
			metricLoad.With("serial_number", sn).Set(float64(v.Value.(int)))
		case "input.voltage":
			metricInputVoltage.With("serial_number", sn).Set(v.Value.(float64))
		case "battery.charge":
			metricBatteryCharge.With("serial_number", sn).Set(float64(v.Value.(int)))
		case "battery.runtime":
			metricBatteryRuntime.With("serial_number", sn).Set(float64(v.Value.(int)))
		case "battery.voltage":
			metricBatteryVoltage.With("serial_number", sn).Set(v.Value.(float64))
		}

		// TODO:
		_ = b.MQTT().PublishAsync(ctx, b.config.TopicVariable.Format(sn, v.Name), v.Value)
	}

	return nil
}
