package google_home_mini

import (
	"github.com/kihamo/boggart/components/boggart"
)

type Type struct{}

func (t Type) CreateBind(c interface{}) (boggart.DeviceBind, error) {
	config := c.(*Config)

	device := &Bind{
		host: config.Host,
	}
	device.Init()

	return device, nil
}
