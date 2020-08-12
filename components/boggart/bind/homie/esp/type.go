package esp

import (
	"github.com/kihamo/boggart/atomic"
	"github.com/kihamo/boggart/components/boggart"
)

type Type struct{}

func (t Type) CreateBind(c interface{}) (boggart.Bind, error) {
	config := c.(*Config)

	config.TopicBroadcast = config.TopicBroadcast.Format(config.BaseTopic)
	config.TopicReset = config.TopicReset.Format(config.BaseTopic, config.DeviceID)
	config.TopicRestart = config.TopicRestart.Format(config.BaseTopic, config.DeviceID)
	config.TopicDeviceAttribute = config.TopicDeviceAttribute.Format(config.BaseTopic, config.DeviceID)
	config.TopicDeviceAttributeFirmware = config.TopicDeviceAttributeFirmware.Format(config.BaseTopic, config.DeviceID)
	config.TopicDeviceAttributeImplementation = config.TopicDeviceAttributeImplementation.Format(config.BaseTopic, config.DeviceID)
	config.TopicDeviceAttributeStats = config.TopicDeviceAttributeStats.Format(config.BaseTopic, config.DeviceID)
	config.TopicNodeAttribute = config.TopicNodeAttribute.Format(config.BaseTopic, config.DeviceID)
	config.TopicNodeList = config.TopicNodeList.Format(config.BaseTopic, config.DeviceID)
	config.TopicNodeProperty = config.TopicNodeProperty.Format(config.BaseTopic, config.DeviceID)
	config.TopicSettings = config.TopicSettings.Format(config.BaseTopic, config.DeviceID)
	config.TopicSettingsSet = config.TopicSettingsSet.Format(config.BaseTopic, config.DeviceID)
	config.TopicOTAFirmware = config.TopicOTAFirmware.Format(config.BaseTopic, config.DeviceID)
	config.TopicOTAStatus = config.TopicOTAStatus.Format(config.BaseTopic, config.DeviceID)
	config.TopicOTAEnabled = config.TopicOTAEnabled.Format(config.BaseTopic, config.DeviceID)

	bind := &Bind{
		config:     config,
		otaFlash:   make(chan struct{}, 1),
		lastUpdate: atomic.NewTimeNull(),
	}

	return bind, nil
}
