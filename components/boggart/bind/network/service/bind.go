package service

import (
	"time"

	"github.com/kihamo/boggart/components/boggart"
)

type Bind struct {
	boggart.BindBase
	boggart.BindMQTT

	address         string
	retry           int
	timeout         time.Duration
	updaterInterval time.Duration
}

func (b *Bind) Run() error {
	b.UpdateStatus(boggart.BindStatusOnline)
	return nil
}
