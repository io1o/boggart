package boggart

import (
	"github.com/kihamo/boggart/components/mqtt"
	"github.com/kihamo/go-workers"
	"github.com/kihamo/go-workers/event"
	"github.com/kihamo/snitch"
)

var (
	DeviceEventSyslogReceive            = event.NewBaseEvent("SyslogReceive")
	DeviceEventDevicesManagerReady      = event.NewBaseEvent("DevicesManagerReady")
	DeviceEventDeviceRegister           = event.NewBaseEvent("DeviceRegister")
	DeviceEventDeviceDisabledAfterCheck = event.NewBaseEvent("DeviceDisabledAfterCheck")
	DeviceEventDeviceEnabledAfterCheck  = event.NewBaseEvent("DeviceEnabledAfterCheck")
	DeviceEventDeviceEnabled            = event.NewBaseEvent("DeviceEnabled")
	DeviceEventDeviceDisabled           = event.NewBaseEvent("DeviceDisabled")
)

type DeviceStatus uint64

const (
	DeviceStatusUnknown DeviceStatus = iota
	DeviceStatusUninitialized
	DeviceStatusInitializing
	DeviceStatusOnline
	DeviceStatusOffline
	DeviceStatusRemoving
	DeviceStatusRemoved
)

type Device interface {
	Bind() DeviceBind
	ID() string
	Type() string
	Description() string
	Tags() []string
	Config() interface{}
	Tasks() []workers.Task
	Listeners() []workers.ListenerWithEvents
	MQTTSubscribers() []mqtt.Subscriber
	MQTTPublishes() []mqtt.Topic
}

type DeviceList []Device

type DeviceBind interface {
	Status() DeviceStatus
	SerialNumber() string
}

type DeviceBindCloser interface {
	Close() error
}

type DeviceBindHasTasks interface {
	Tasks() []workers.Task
}

type DeviceBindHasListeners interface {
	Listeners() []workers.ListenerWithEvents
}

type DeviceBindHasMQTTClient interface {
	SetMQTTClient(mqtt.Component)
}

type DeviceBindHasMQTTSubscribers interface {
	MQTTSubscribers() []mqtt.Subscriber
}

type DeviceBindHasMQTTPublishes interface {
	MQTTPublishes() []mqtt.Topic
}

type DevicesManager interface {
	snitch.Collector

	Register(device DeviceBind, t string, description string, tags []string, config interface{}) (string, error)
	RegisterWithID(id string, bind DeviceBind, t string, description string, tags []string, config interface{}) error
	Unregister(id string) error
	Device(id string) Device
	Devices() DeviceList
	IsReady() bool
}

func (l DeviceList) MarshalYAML() (interface{}, error) {
	return struct {
		Devices []Device
	}{
		Devices: l,
	}, nil
}
