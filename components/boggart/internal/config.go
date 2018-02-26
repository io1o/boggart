package internal

import (
	"strconv"
	"time"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/protocols/rs485"
	"github.com/kihamo/boggart/components/boggart/providers/pulsar"
	"github.com/kihamo/shadow/components/config"
)

func (c *Component) ConfigVariables() []config.Variable {
	return []config.Variable{
		config.NewVariable(
			boggart.ConfigDevicesManagerCheckInterval,
			config.ValueTypeDuration,
			time.Minute,
			"Health check interval",
			true,
			"Device manager",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigDevicesManagerCheckTimeout,
			config.ValueTypeDuration,
			DefaultTimeoutChecker,
			"Health check timeout",
			true,
			"Device manager",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigListenerTelegramChats,
			config.ValueTypeString,
			nil,
			"Chats for messages",
			false,
			"Listener Telegram",
			[]string{config.ViewTags},
			map[string]interface{}{
				config.ViewOptionTagsDefaultText: "add a chat ID",
			}),
		config.NewVariable(
			boggart.ConfigRS485Address,
			config.ValueTypeString,
			rs485.DefaultSerialAddress,
			"Serial port address",
			false,
			"RS485 protocol",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigRS485Timeout,
			config.ValueTypeDuration,
			rs485.DefaultTimeout,
			"Serial port timeout",
			false,
			"RS485 protocol",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigDoorsEnabled,
			config.ValueTypeBool,
			false,
			"Enabled",
			false,
			"Doors",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigDoorsEntrancePin,
			config.ValueTypeInt,
			nil,
			"Pin for door reed switch",
			false,
			"Doors",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigVideoRecorderHikVisionEnabled,
			config.ValueTypeBool,
			false,
			"Enabled",
			false,
			"HikVision video recorder",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigVideoRecorderHikVisionRepeatInterval,
			config.ValueTypeDuration,
			time.Minute,
			"Repeat interval",
			false,
			"HikVision video recorder",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigVideoRecorderHikVisionHost,
			config.ValueTypeString,
			nil,
			"Host",
			false,
			"HikVision video recorder",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigVideoRecorderHikVisionPort,
			config.ValueTypeInt64,
			nil,
			"Port",
			false,
			"HikVision video recorder",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigVideoRecorderHikVisionUsername,
			config.ValueTypeString,
			"admin",
			"Username",
			false,
			"HikVision video recorder",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigVideoRecorderHikVisionPassword,
			config.ValueTypeString,
			nil,
			"Password",
			false,
			"HikVision video recorder",
			[]string{config.ViewPassword},
			nil),
		config.NewVariable(
			boggart.ConfigCameraHikVisionHallEnabled,
			config.ValueTypeBool,
			false,
			"Enabled",
			false,
			"HikVision on the hall",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigCameraHikVisionHallRepeatInterval,
			config.ValueTypeDuration,
			time.Minute,
			"Repeat interval",
			false,
			"HikVision on the hall",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigCameraHikVisionHallHost,
			config.ValueTypeString,
			nil,
			"Host",
			false,
			"HikVision on the hall",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigCameraHikVisionHallPort,
			config.ValueTypeInt64,
			nil,
			"Port",
			false,
			"HikVision on the hall",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigCameraHikVisionHallUsername,
			config.ValueTypeString,
			"admin",
			"Username",
			false,
			"HikVision on the hall",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigCameraHikVisionHallPassword,
			config.ValueTypeString,
			nil,
			"Password",
			false,
			"HikVision on the hall",
			[]string{config.ViewPassword},
			nil),
		config.NewVariable(
			boggart.ConfigCameraHikVisionHallStreamingChannel,
			config.ValueTypeInt64,
			101,
			"Streaming channel",
			false,
			"HikVision on the hall",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigCameraHikVisionStreetEnabled,
			config.ValueTypeBool,
			false,
			"Enabled",
			false,
			"HikVision on the street",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigCameraHikVisionStreetRepeatInterval,
			config.ValueTypeDuration,
			time.Minute,
			"Repeat interval",
			false,
			"HikVision on the street",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigCameraHikVisionStreetHost,
			config.ValueTypeString,
			nil,
			"Host",
			false,
			"HikVision on the street",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigCameraHikVisionStreetPort,
			config.ValueTypeInt64,
			nil,
			"Port",
			false,
			"HikVision on the street",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigCameraHikVisionStreetUsername,
			config.ValueTypeString,
			"admin",
			"Username",
			false,
			"HikVision on the street",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigCameraHikVisionStreetPassword,
			config.ValueTypeString,
			nil,
			"Password",
			false,
			"HikVision on the street",
			[]string{config.ViewPassword},
			nil),
		config.NewVariable(
			boggart.ConfigCameraHikVisionStreetStreamingChannel,
			config.ValueTypeInt64,
			101,
			"Streaming channel",
			false,
			"HikVision on the street",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigMercuryEnabled,
			config.ValueTypeBool,
			false,
			"Enabled",
			false,
			"Mercury devices",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigMercuryRepeatInterval,
			config.ValueTypeDuration,
			time.Minute*2,
			"Repeat interval",
			false,
			"Mercury devices",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigMercuryDeviceAddress,
			config.ValueTypeString,
			nil,
			"Device address in format XXXXXX (last 6 digits of device serial number)",
			false,
			"Mercury devices",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigMikrotikEnabled,
			config.ValueTypeBool,
			false,
			"Enabled",
			false,
			"Mikrotik devices",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigMikrotikRepeatInterval,
			config.ValueTypeDuration,
			time.Minute*5,
			"Repeat interval",
			false,
			"Mikrotik devices",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigMikrotikAddress,
			config.ValueTypeString,
			"192.168.88.1:8728",
			"API address in format host:port",
			false,
			"Mikrotik devices",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigMikrotikUsername,
			config.ValueTypeString,
			"admin",
			"Username",
			false,
			"Mikrotik devices",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigMikrotikPassword,
			config.ValueTypeString,
			nil,
			"Password",
			false,
			"Mikrotik devices",
			[]string{config.ViewPassword},
			nil),
		config.NewVariable(
			boggart.ConfigMikrotikTimeout,
			config.ValueTypeDuration,
			time.Second*10,
			"Request timeout",
			false,
			"Mikrotik devices",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigMobileEnabled,
			config.ValueTypeBool,
			false,
			"Enabled",
			false,
			"Mobile accounts",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigMobileRepeatInterval,
			config.ValueTypeDuration,
			time.Minute*30,
			"Repeat interval",
			false,
			"Mobile accounts",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigMobileMegafonPhone,
			config.ValueTypeString,
			nil,
			"Phone number",
			false,
			"Mobile Megafon accounts",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigMobileMegafonPassword,
			config.ValueTypeString,
			nil,
			"Password",
			false,
			"Mobile Megafon accounts",
			[]string{config.ViewPassword},
			nil),
		config.NewVariable(
			boggart.ConfigPulsarEnabled,
			config.ValueTypeBool,
			false,
			"Enabled",
			false,
			"Pulsar devices",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigPulsarRepeatInterval,
			config.ValueTypeDuration,
			time.Minute*15,
			"Repeat interval",
			false,
			"Pulsar devices",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigPulsarHeatMeterAddress,
			config.ValueTypeString,
			nil,
			"Device address HEX value (AABBCCDD). If empty system try to find device",
			false,
			"Pulsar devices",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigPulsarColdWaterSerialNumber,
			config.ValueTypeString,
			nil,
			"Device address of cold water meter in format AABBCCDD",
			false,
			"Pulsar devices",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigPulsarColdWaterPulseInput,
			config.ValueTypeUint64,
			pulsar.Input1,
			"Input number of cold water meter in heat meter",
			false,
			"Pulsar devices",
			[]string{config.ViewEnum},
			map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{strconv.FormatUint(pulsar.Input1, 10), "#1"},
					{strconv.FormatUint(pulsar.Input2, 10), "#2"},
				},
			}),
		config.NewVariable(
			boggart.ConfigPulsarColdWaterStartValue,
			config.ValueTypeFloat64,
			0,
			"Start value of cold water meter (in m3)",
			false,
			"Pulsar devices",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigPulsarHotWaterSerialNumber,
			config.ValueTypeString,
			nil,
			"Device address of hot water meter in format AABBCCDD",
			false,
			"Pulsar devices",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigPulsarHotWaterPulseInput,
			config.ValueTypeUint64,
			pulsar.Input2,
			"Input number of hot water meter",
			false,
			"Pulsar devices",
			[]string{config.ViewEnum},
			map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{strconv.FormatUint(pulsar.Input1, 10), "#1"},
					{strconv.FormatUint(pulsar.Input2, 10), "#2"},
				},
			}),
		config.NewVariable(
			boggart.ConfigPulsarHotWaterStartValue,
			config.ValueTypeFloat64,
			0,
			"Start value of hot water (in m3)",
			false,
			"Pulsar devices",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigSoftVideoEnabled,
			config.ValueTypeBool,
			false,
			"Enabled",
			false,
			"SoftVideo provider",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigSoftVideoRepeatInterval,
			config.ValueTypeDuration,
			time.Hour*12,
			"Repeat interval",
			false,
			"SoftVideo provider",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigSoftVideoLogin,
			config.ValueTypeString,
			nil,
			"Login",
			false,
			"SoftVideo provider",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigSoftVideoPassword,
			config.ValueTypeString,
			nil,
			"Password",
			false,
			"SoftVideo provider",
			[]string{config.ViewPassword},
			nil),
		config.NewVariable(
			boggart.ConfigMonitoringExternalURL,
			config.ValueTypeString,
			nil,
			"Monitoring external URL",
			false,
			"Others",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigApcupsdEnabled,
			config.ValueTypeBool,
			false,
			"Enabled",
			false,
			"Apcupsd",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigApcupsdRepeatInterval,
			config.ValueTypeDuration,
			time.Minute,
			"Repeat interval",
			false,
			"Apcupsd",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigApcupsdNISAddress,
			config.ValueTypeString,
			"127.0.0.1:3551",
			"NIS address",
			false,
			"Apcupsd",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigApcupsdFileStatus,
			config.ValueTypeString,
			"/var/log/apcupsd.status",
			"File status",
			false,
			"Apcupsd",
			nil,
			nil),
		config.NewVariable(
			boggart.ConfigApcupsdFileEvents,
			config.ValueTypeString,
			"/var/log/apcupsd.events",
			"File events",
			false,
			"Apcupsd",
			nil,
			nil),
	}
}

func (c *Component) ConfigWatchers() []config.Watcher {
	return []config.Watcher{
		config.NewWatcher(c.Name(), []string{
			boggart.ConfigDevicesManagerCheckInterval,
		}, c.watchDevicesManagerCheckInterval),
		config.NewWatcher(c.Name(), []string{
			boggart.ConfigDevicesManagerCheckTimeout,
		}, c.watchDevicesManagerCheckTimeout),
	}
}

func (c *Component) watchDevicesManagerCheckInterval(_ string, newValue interface{}, _ interface{}) {
	c.devicesManager.SetCheckerTickerDuration(newValue.(time.Duration))
}

func (c *Component) watchDevicesManagerCheckTimeout(_ string, newValue interface{}, _ interface{}) {
	c.devicesManager.SetCheckerTimeout(newValue.(time.Duration))
}
