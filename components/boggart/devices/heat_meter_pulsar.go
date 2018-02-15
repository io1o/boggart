package devices

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/providers/pulsar"
	"github.com/kihamo/go-workers"
	"github.com/kihamo/go-workers/task"
	"github.com/kihamo/snitch"
)

var (
	metricHeatMeterPulsarTemperatureIn    = snitch.NewGauge(boggart.ComponentName+"_device_heat_meter_pulsar_temperature_in_celsius", "Pulsar temperature in")
	metricHeatMeterPulsarTemperatureOut   = snitch.NewGauge(boggart.ComponentName+"_device_heat_meter_pulsar_temperature_out_celsius", "Pulsar temperature out")
	metricHeatMeterPulsarTemperatureDelta = snitch.NewGauge(boggart.ComponentName+"_device_heat_meter_pulsar_temperature_delta_celsius", "Pulsar temperature delta")
	metricHeatMeterPulsarEnergy           = snitch.NewGauge(boggart.ComponentName+"_device_heat_meter_pulsar_energy_gigacolories", "Pulsar energy")
	metricHeatMeterPulsarConsumption      = snitch.NewGauge(boggart.ComponentName+"_device_heat_meter_pulsar_consumption_cubic_metres_per_hour", "Pulsar consumption")
)

type PulsarHeadMeter struct {
	boggart.DeviceBase

	serialNumber string
	provider     *pulsar.HeatMeter
	interval     time.Duration
}

func NewPulsarHeadMeter(provider *pulsar.HeatMeter, interval time.Duration) *PulsarHeadMeter {
	device := &PulsarHeadMeter{
		serialNumber: hex.EncodeToString(provider.Address()),
		provider:     provider,
		interval:     interval,
	}
	device.Init()
	device.SetDescription("Pulsar heat meter with serial number " + device.serialNumber)

	return device
}

func (d *PulsarHeadMeter) Types() []boggart.DeviceType {
	return []boggart.DeviceType{
		boggart.DeviceTypeHeatMeter,
	}
}

func (d *PulsarHeadMeter) TemperatureIn(context.Context) (float64, error) {
	value, err := d.provider.TemperatureIn()
	if err != nil {
		return -1, err
	}

	return float64(value), nil
}

func (d *PulsarHeadMeter) TemperatureOut(context.Context) (float64, error) {
	value, err := d.provider.TemperatureOut()
	if err != nil {
		return -1, err
	}

	return float64(value), nil
}

func (d *PulsarHeadMeter) TemperatureDelta(context.Context) (float64, error) {
	value, err := d.provider.TemperatureDelta()
	if err != nil {
		return -1, err
	}

	return float64(value), nil
}

func (d *PulsarHeadMeter) Energy(context.Context) (float64, error) {
	value, err := d.provider.Energy()
	if err != nil {
		return -1, err
	}

	return float64(value), nil
}

func (d *PulsarHeadMeter) Consumption(context.Context) (float64, error) {
	value, err := d.provider.Consumption()
	if err != nil {
		return -1, err
	}

	return float64(value), nil
}

func (d *PulsarHeadMeter) Describe(ch chan<- *snitch.Description) {
	metricHeatMeterPulsarTemperatureIn.With("serial_number", d.serialNumber).Describe(ch)
	metricHeatMeterPulsarTemperatureOut.With("serial_number", d.serialNumber).Describe(ch)
	metricHeatMeterPulsarTemperatureDelta.With("serial_number", d.serialNumber).Describe(ch)
	metricHeatMeterPulsarEnergy.With("serial_number", d.serialNumber).Describe(ch)
	metricHeatMeterPulsarConsumption.With("serial_number", d.serialNumber).Describe(ch)
}

func (d *PulsarHeadMeter) Collect(ch chan<- snitch.Metric) {
	metricHeatMeterPulsarTemperatureIn.With("serial_number", d.serialNumber).Collect(ch)
	metricHeatMeterPulsarTemperatureOut.With("serial_number", d.serialNumber).Collect(ch)
	metricHeatMeterPulsarTemperatureDelta.With("serial_number", d.serialNumber).Collect(ch)
	metricHeatMeterPulsarEnergy.With("serial_number", d.serialNumber).Collect(ch)
	metricHeatMeterPulsarConsumption.With("serial_number", d.serialNumber).Collect(ch)
}

func (d *PulsarHeadMeter) Ping(_ context.Context) bool {
	_, err := d.provider.Version()
	return err == nil
}

func (d *PulsarHeadMeter) Tasks() []workers.Task {
	taskUpdater := task.NewFunctionTask(d.updater)
	taskUpdater.SetRepeats(-1)
	taskUpdater.SetRepeatInterval(d.interval)
	taskUpdater.SetName("device-heat-meter-pulsar-updater-" + d.serialNumber)

	return []workers.Task{
		taskUpdater,
	}
}

func (d *PulsarHeadMeter) updater(ctx context.Context) (interface{}, error) {
	if !d.IsEnabled() {
		return nil, nil
	}

	if value, err := d.TemperatureIn(ctx); err == nil {
		metricHeatMeterPulsarTemperatureIn.With("serial_number", d.serialNumber).Set(value)
	} else {
		return nil, err
	}

	if value, err := d.TemperatureOut(ctx); err == nil {
		metricHeatMeterPulsarTemperatureOut.With("serial_number", d.serialNumber).Set(value)
	} else {
		return nil, err
	}

	if value, err := d.TemperatureDelta(ctx); err == nil {
		metricHeatMeterPulsarTemperatureDelta.With("serial_number", d.serialNumber).Set(value)
	} else {
		return nil, err
	}

	if value, err := d.Energy(ctx); err == nil {
		metricHeatMeterPulsarEnergy.With("serial_number", d.serialNumber).Set(value)
	} else {
		return nil, err
	}

	if value, err := d.Consumption(ctx); err == nil {
		metricHeatMeterPulsarConsumption.With("serial_number", d.serialNumber).Set(value)
	} else {
		return nil, err
	}

	return nil, nil
}
