package octoprint

import (
	"context"
	"strconv"

	"github.com/kihamo/boggart/providers/octoprint/client/job"
	"github.com/kihamo/boggart/providers/octoprint/client/plugin_display_layer_progress"
	"github.com/kihamo/boggart/providers/octoprint/client/printer"
	"github.com/kihamo/boggart/providers/octoprint/models"
	"github.com/kihamo/go-workers"
	"go.uber.org/multierr"
)

func (b *Bind) Tasks() []workers.Task {
	taskUpdater := b.WrapTaskIsOnline(b.taskUpdater)
	taskUpdater.SetTimeout(b.config.UpdaterTimeout)
	taskUpdater.SetRepeats(-1)
	taskUpdater.SetRepeatInterval(b.config.UpdaterInterval)
	taskUpdater.SetName("updater")

	tasks := []workers.Task{
		taskUpdater,
	}

	return tasks
}

func (b *Bind) taskUpdater(ctx context.Context) error {
	/*
		loginParams := authorization.NewLoginParamsWithContext(ctx)
		loginParams.Body.Passive = true

		_, err := b.provider.Authorization.Login(loginParams, nil)
		if err != nil {
			return err
		}
	*/

	stateParams := printer.NewGetPrinterStateParamsWithContext(ctx).
		WithHistory(&[]bool{false}[0]).
		WithExclude([]string{"sd"})

	state, err := b.provider.Printer.GetPrinterState(stateParams, nil)
	address := b.config.Address.Host

	if err != nil {
		if _, ok := err.(*printer.GetPrinterStateConflict); ok {
			err = b.MQTTPublishAsync(ctx, b.config.TopicState.Format(address), "Not operational")
		}

		return err
	}

	if e := b.MQTTPublishAsync(ctx, b.config.TopicState.Format(address), state.Payload.State.Text); e != nil {
		err = multierr.Append(err, e)
	}

	var temperature *models.TemperatureData

	// Bed
	temperature = state.Payload.Temperature.Bed

	metricDeviceTemperatureActual.With("address", address).With("device", "bed").Set(temperature.Actual)
	if e := b.MQTTPublishAsync(ctx, b.config.TopicStateBedTemperatureActual.Format(address), temperature.Actual); e != nil {
		err = multierr.Append(err, e)
	}

	if e := b.MQTTPublishAsync(ctx, b.config.TopicStateBedTemperatureOffset.Format(address), temperature.Offset); e != nil {
		err = multierr.Append(err, e)
	}

	metricDeviceTemperatureTarget.With("address", address).With("device", "bed").Set(temperature.Target)
	if e := b.MQTTPublishAsync(ctx, b.config.TopicStateBedTemperatureTarget.Format(address), temperature.Target); e != nil {
		err = multierr.Append(err, e)
	}

	// HotEnd (tool 0)
	temperature = state.Payload.Temperature.Tool0

	metricDeviceTemperatureActual.With("address", address).With("device", "tool0").Set(temperature.Actual)
	if e := b.MQTTPublishAsync(ctx, b.config.TopicStateTool0TemperatureActual.Format(address), temperature.Actual); e != nil {
		err = multierr.Append(err, e)
	}

	if e := b.MQTTPublishAsync(ctx, b.config.TopicStateTool0TemperatureOffset.Format(address), temperature.Offset); e != nil {
		err = multierr.Append(err, e)
	}

	metricDeviceTemperatureTarget.With("address", address).With("device", "tool0").Set(temperature.Target)
	if e := b.MQTTPublishAsync(ctx, b.config.TopicStateTool0TemperatureTarget.Format(address), temperature.Target); e != nil {
		err = multierr.Append(err, e)
	}

	// Job
	if j, e := b.provider.Job.GetJob(job.NewGetJobParamsWithContext(ctx), nil); e == nil {
		if e := b.MQTTPublishAsync(ctx, b.config.TopicStateJobFileName.Format(address), j.Payload.Job.File.Name); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTTPublishAsync(ctx, b.config.TopicStateJobFileSize.Format(address), j.Payload.Job.File.Size); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTTPublishAsync(ctx, b.config.TopicStateJobProgress.Format(address), j.Payload.Progress.Completion); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTTPublishAsync(ctx, b.config.TopicStateJobTime.Format(address), j.Payload.Progress.PrintTime); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTTPublishAsync(ctx, b.config.TopicStateJobTimeLeft.Format(address), j.Payload.Progress.PrintTimeLeft); e != nil {
			err = multierr.Append(err, e)
		}
	}

	// Layer & Height
	var (
		layerTotal, layerCurrent   uint64
		heightTotal, heightCurrent float64
	)

	if state.Payload.State.Flags.Printing {
		if progress, e := b.provider.PluginDisplayLayerProgress.DisplayLayerProgress(plugin_display_layer_progress.NewDisplayLayerProgressParamsWithContext(ctx), nil); e == nil {
			if value, e := strconv.ParseUint(progress.Payload.Layer.Total, 10, 64); e == nil {
				layerTotal = value
			}

			if progress.Payload.Layer.Current != "-" {
				if value, e := strconv.ParseUint(progress.Payload.Layer.Current, 10, 64); e == nil {
					layerCurrent = value
				}
			}

			if value, e := strconv.ParseFloat(progress.Payload.Height.Total, 64); e == nil {
				heightTotal = value
			}

			if value, e := strconv.ParseFloat(progress.Payload.Height.Current, 64); e == nil {
				heightCurrent = value
			}
		}
	}

	if e := b.MQTTPublishAsync(ctx, b.config.TopicLayerTotal.Format(address), layerTotal); e != nil {
		err = multierr.Append(err, e)
	}

	if e := b.MQTTPublishAsync(ctx, b.config.TopicLayerCurrent.Format(address), layerCurrent); e != nil {
		err = multierr.Append(err, e)
	}

	if e := b.MQTTPublishAsync(ctx, b.config.TopicHeightTotal.Format(address), heightTotal); e != nil {
		err = multierr.Append(err, e)
	}

	if e := b.MQTTPublishAsync(ctx, b.config.TopicHeightCurrent.Format(address), heightCurrent); e != nil {
		err = multierr.Append(err, e)
	}

	return err
}
