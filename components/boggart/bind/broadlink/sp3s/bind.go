package sp3s

import (
	"context"

	"github.com/kihamo/boggart/atomic"
	"github.com/kihamo/boggart/components/boggart/di"
	"github.com/kihamo/boggart/providers/broadlink"
)

type Bind struct {
	di.MetaBind
	di.MQTTBind
	di.WorkersBind
	di.LoggerBind
	di.ProbesBind
	di.WidgetBind

	config *Config

	state *atomic.BoolNull
	power *atomic.Float32Null

	provider *broadlink.SP3S
}

func (b *Bind) Run() error {
	b.Meta().SetMAC(b.config.MAC.HardwareAddr)

	return nil
}

func (b *Bind) State() (bool, error) {
	return b.provider.State()
}

func (b *Bind) On(ctx context.Context) error {
	err := b.provider.On()
	if err == nil {
		err = b.taskUpdater(ctx)
	}

	return err
}

func (b *Bind) Off(ctx context.Context) error {
	err := b.provider.Off()
	if err == nil {
		err = b.taskUpdater(ctx)
	}

	return err
}

func (b *Bind) Power() (float64, error) {
	return b.provider.Power()
}
