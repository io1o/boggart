package ds18b20

import (
	"time"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/atomic"
)

type Bind struct {
	boggart.BindBase
	boggart.BindMQTT

	temperature *atomic.Float32Null

	livenessInterval time.Duration
	livenessTimeout  time.Duration
	updaterInterval  time.Duration
}
