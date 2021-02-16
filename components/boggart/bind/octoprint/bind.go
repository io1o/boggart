package octoprint

import (
	"github.com/kihamo/boggart/components/boggart/di"
	"github.com/kihamo/boggart/providers/octoprint"
)

type Bind struct {
	di.ConfigBind
	di.LoggerBind
	di.MetricsBind
	di.MQTTBind
	di.ProbesBind
	di.WorkersBind

	config   *Config
	provider *octoprint.Client
}
