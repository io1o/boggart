package grafana

import (
	"time"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/di"
	"github.com/kihamo/boggart/components/mqtt"
)

type Config struct {
	di.ProbesConfig `mapstructure:",squash" yaml:",inline"`
	di.LoggerConfig `mapstructure:",squash" yaml:",inline"`

	Address         boggart.URL `valid:",required"`
	Debug           bool
	Name            string `valid:",required"`
	Dashboards      []int64
	TopicAnnotation mqtt.Topic `mapstructure:"topic_annotation" yaml:"topic_annotation"`
}

func (t Type) Config() interface{} {
	return &Config{
		ProbesConfig: di.ProbesConfig{
			ReadinessPeriod:  time.Minute,
			ReadinessTimeout: time.Second * 5,
		},
		TopicAnnotation: boggart.ComponentName + "/grafana/+/annotation",
	}
}
