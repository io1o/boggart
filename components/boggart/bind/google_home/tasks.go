package google_home

import (
	"context"
	"errors"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/providers/google/home"
	"github.com/kihamo/boggart/providers/google/home/client/info"
	"github.com/kihamo/go-workers"
	"github.com/kihamo/go-workers/task"
)

func (b *Bind) Tasks() []workers.Task {
	taskLiveness := task.NewFunctionTask(b.taskLiveness)
	taskLiveness.SetTimeout(b.config.LivenessInterval)
	taskLiveness.SetRepeats(-1)
	taskLiveness.SetRepeatInterval(b.config.LivenessTimeout)
	taskLiveness.SetName("liveness-" + b.config.Host.String())

	return []workers.Task{
		taskLiveness,
	}
}

func (b *Bind) taskLiveness(ctx context.Context) (interface{}, error) {
	response, err := b.ClientGoogleHome().Info.GetEurekaInfo(info.NewGetEurekaInfoParams().
		WithOptions(home.EurekaInfoOptionDetail.Value()).
		WithParams(home.EurekaInfoParamDeviceInfo.Value()))

	if err != nil {
		b.UpdateStatus(boggart.BindStatusOffline)
		return nil, nil
	}

	if b.SerialNumber() == "" {
		if response.Payload == nil || response.Payload.MacAddress == "" {
			return nil, errors.New("MAC address not found")
		}

		b.SetSerialNumber(response.Payload.MacAddress)
	}

	b.UpdateStatus(boggart.BindStatusOnline)
	return nil, nil
}
