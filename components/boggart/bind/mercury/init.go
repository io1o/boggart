package mercury

import (
	"github.com/kihamo/boggart/components/boggart"
)

func init() {
	boggart.RegisterDeviceType("mercury", Type{})
}
