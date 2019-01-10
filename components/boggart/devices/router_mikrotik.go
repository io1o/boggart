package devices

import (
	"context"
	"errors"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/providers/mikrotik"
	"github.com/kihamo/boggart/components/mqtt"
	"github.com/kihamo/go-workers"
	"github.com/kihamo/go-workers/listener"
	"github.com/kihamo/go-workers/task"
	"github.com/kihamo/snitch"
)

const (
	MikrotikRouterMQTTTopicWiFiMACState         mqtt.Topic = boggart.ComponentName + "/router/+/wifi/clients/+/state"
	MikrotikRouterMQTTTopicWiFiConnectedMAC     mqtt.Topic = boggart.ComponentName + "/router/+/wifi/clients/last/on/mac"
	MikrotikRouterMQTTTopicWiFiDisconnectedMAC  mqtt.Topic = boggart.ComponentName + "/router/+/wifi/clients/last/on/mac"
	MikrotikRouterMQTTTopicVPNLoginState        mqtt.Topic = boggart.ComponentName + "/router/+/vpn/clients/+/state"
	MikrotikRouterMQTTTopicVPNConnectedLogin    mqtt.Topic = boggart.ComponentName + "/router/+/vpn/clients/last/on/login"
	MikrotikRouterMQTTTopicVPNDisconnectedLogin mqtt.Topic = boggart.ComponentName + "/router/+/vpn/clients/last/off/login"
)

var (
	metricRouterMikrotikTrafficReceivedBytes = snitch.NewGauge(boggart.ComponentName+"_device_router_mikrotik_traffic_received_bytes", "Mikrotik traffic received bytes")
	metricRouterMikrotikTrafficSentBytes     = snitch.NewGauge(boggart.ComponentName+"_device_router_mikrotik_traffic_sent_bytes", "Mikrotik traffic sent bytes")
	metricRouterMikrotikWifiClients          = snitch.NewGauge(boggart.ComponentName+"_device_router_mikrotik_wifi_clients_total", "Mikrotik wifi clients online")
	metricRouterMikrotikCPULoad              = snitch.NewGauge(boggart.ComponentName+"_device_router_mikrotik_cpu_load_percent", "CPU load in percents")
	metricRouterMikrotikMemoryUsage          = snitch.NewGauge(boggart.ComponentName+"_device_router_mikrotik_memory_usage_bytes", "Memory usage in Mikrotik router")
	metricRouterMikrotikMemoryAvailable      = snitch.NewGauge(boggart.ComponentName+"_device_router_mikrotik_memory_available_bytes", "Memory available in Mikrotik router")
	metricRouterMikrotikStorageUsage         = snitch.NewGauge(boggart.ComponentName+"_device_router_mikrotik_storage_usage_bytes", "Storage usage in Mikrotik router")
	metricRouterMikrotikStorageAvailable     = snitch.NewGauge(boggart.ComponentName+"_device_router_mikrotik_storage_available_bytes", "Storage available in Mikrotik router")
	metricRouterMikrotikDiskUsage            = snitch.NewGauge(boggart.ComponentName+"_device_router_mikrotik_disk_usage_bytes", "Disk usage in Mikrotik router")
	metricRouterMikrotikDiskAvailable        = snitch.NewGauge(boggart.ComponentName+"_device_router_mikrotik_disk_available_bytes", "Disk available in Mikrotik router")
	metricRouterMikrotikVoltage              = snitch.NewGauge(boggart.ComponentName+"_device_router_mikrotik_voltage_volt", "Voltage")
	metricRouterMikrotikTemperature          = snitch.NewGauge(boggart.ComponentName+"_device_router_mikrotik_temperature_celsius", "Temperature")

	wifiClientRegexp = regexp.MustCompile(`^([^@]+)@([^:\s]+):\s+([^\s,]+)`)
	vpnClientRegexp  = regexp.MustCompile(`^(\S+) logged (in|out), (.+?)$`)
)

type MikrotikRouter struct {
	boggart.DeviceBindBase
	boggart.DeviceSerialNumber
	boggart.DeviceMQTT

	provider     *mikrotik.Client
	host         string
	syslogClient string
	interval     time.Duration
}

type MikrotikRouterListener struct {
	listener.BaseListener

	router *MikrotikRouter
}

type MikrotikRouterMac struct {
	Address string
	ARP     struct {
		IP      string
		Comment string
	}
	DHCP struct {
		Hostname string
	}
}

func (d MikrotikRouter) CreateBind(config map[string]interface{}) (boggart.DeviceBind, error) {
	address, ok := config["address"]
	if !ok {
		return nil, errors.New("config option address isn't set")
	}

	if address == "" {
		return nil, errors.New("config option address is empty")
	}

	u, err := url.Parse(address.(string))
	if err != nil {
		return nil, errors.New("Bad Mikrotik address " + address.(string))
	}

	username := u.User.Username()
	password, _ := u.User.Password()

	device := &MikrotikRouter{
		provider:     mikrotik.NewClient(u.Host, username, password, time.Second*10),
		host:         u.Host + "-" + u.Port(),
		syslogClient: u.Hostname() + ":514",
	}
	device.Init()

	return device, nil
}

func NewMikrotikRouterListener(router *MikrotikRouter) *MikrotikRouterListener {
	l := &MikrotikRouterListener{
		router: router,
	}
	l.Init()

	return l
}

func (d *MikrotikRouter) Describe(ch chan<- *snitch.Description) {
	serialNumber := d.SerialNumber()
	if serialNumber == "" {
		return
	}

	metricRouterMikrotikTrafficReceivedBytes.With("serial_number", serialNumber).Describe(ch)
	metricRouterMikrotikTrafficSentBytes.With("serial_number", serialNumber).Describe(ch)
	metricRouterMikrotikWifiClients.With("serial_number", serialNumber).Describe(ch)
	metricRouterMikrotikCPULoad.With("serial_number", serialNumber).Describe(ch)
	metricRouterMikrotikMemoryUsage.With("serial_number", serialNumber).Describe(ch)
	metricRouterMikrotikMemoryAvailable.With("serial_number", serialNumber).Describe(ch)
	metricRouterMikrotikStorageUsage.With("serial_number", serialNumber).Describe(ch)
	metricRouterMikrotikStorageAvailable.With("serial_number", serialNumber).Describe(ch)
	metricRouterMikrotikDiskUsage.With("serial_number", serialNumber).Describe(ch)
	metricRouterMikrotikDiskAvailable.With("serial_number", serialNumber).Describe(ch)
	metricRouterMikrotikVoltage.With("serial_number", serialNumber).Describe(ch)
	metricRouterMikrotikTemperature.With("serial_number", serialNumber).Describe(ch)
}

func (d *MikrotikRouter) Collect(ch chan<- snitch.Metric) {
	serialNumber := d.SerialNumber()
	if serialNumber == "" {
		return
	}

	metricRouterMikrotikTrafficReceivedBytes.With("serial_number", serialNumber).Collect(ch)
	metricRouterMikrotikTrafficSentBytes.With("serial_number", serialNumber).Collect(ch)
	metricRouterMikrotikWifiClients.With("serial_number", serialNumber).Collect(ch)
	metricRouterMikrotikCPULoad.With("serial_number", serialNumber).Collect(ch)
	metricRouterMikrotikMemoryUsage.With("serial_number", serialNumber).Collect(ch)
	metricRouterMikrotikMemoryAvailable.With("serial_number", serialNumber).Collect(ch)
	metricRouterMikrotikStorageUsage.With("serial_number", serialNumber).Collect(ch)
	metricRouterMikrotikStorageAvailable.With("serial_number", serialNumber).Collect(ch)
	metricRouterMikrotikDiskUsage.With("serial_number", serialNumber).Collect(ch)
	metricRouterMikrotikDiskAvailable.With("serial_number", serialNumber).Collect(ch)
	metricRouterMikrotikVoltage.With("serial_number", serialNumber).Collect(ch)
	metricRouterMikrotikTemperature.With("serial_number", serialNumber).Collect(ch)
}

func (d *MikrotikRouter) Tasks() []workers.Task {
	taskLiveness := task.NewFunctionTask(d.taskLiveness)
	taskLiveness.SetTimeout(time.Second * 5)
	taskLiveness.SetRepeats(-1)
	taskLiveness.SetRepeatInterval(time.Minute)
	taskLiveness.SetName("device-router-mikrotik-liveness")

	taskStateUpdater := task.NewFunctionTask(d.taskStateUpdater)
	taskStateUpdater.SetRepeats(-1)
	taskStateUpdater.SetRepeatInterval(time.Minute * 5)
	taskStateUpdater.SetName("device-router-mikrotik-updater-" + d.host)

	return []workers.Task{
		taskLiveness,
		taskStateUpdater,
	}
}

func (d *MikrotikRouter) Listeners() []workers.ListenerWithEvents {
	return []workers.ListenerWithEvents{
		NewMikrotikRouterListener(d),
	}
}

func (d *MikrotikRouter) Mac(ctx context.Context, mac string) (*MikrotikRouterMac, error) {
	if d.SerialNumber() == "" {
		return nil, errors.New("serial number is empty")
	}

	info := &MikrotikRouterMac{
		Address: mac,
	}

	if table, err := d.provider.IPARP(ctx); err == nil {
		for _, row := range table {
			if row.MacAddress == mac {
				info.ARP.IP = row.Address
				info.ARP.Comment = row.Comment
				break
			}
		}
	} else {
		return nil, err
	}

	if leases, err := d.provider.IPDHCPServerLease(ctx); err == nil {
		for _, lease := range leases {
			if lease.MacAddress == mac {
				info.DHCP.Hostname = lease.MacAddress
				break
			}
		}
	} else {
		return nil, err
	}

	return info, nil
}

func (d *MikrotikRouter) taskLiveness(ctx context.Context) (interface{}, error) {
	system, err := d.provider.SystemRouterboard(ctx)
	if err != nil {
		d.UpdateStatus(boggart.DeviceStatusOffline)
		return nil, err
	}

	if system.SerialNumber == "" {
		d.UpdateStatus(boggart.DeviceStatusOffline)
		return nil, errors.New("serial number is empty")
	}

	d.UpdateStatus(boggart.DeviceStatusOnline)
	if d.SerialNumber() != "" {
		return nil, nil
	}

	d.SetSerialNumber(system.SerialNumber)
	sn := system.SerialNumber

	// wifi clients
	clients, err := d.provider.InterfaceWirelessRegistrationTable(ctx)
	if err != nil {
		return nil, err
	}

	for _, connection := range clients {
		mac, err := d.Mac(ctx, connection.MacAddress)
		if err != nil {
			return nil, err
		}

		login := mqtt.NameReplace(mac.Address)

		d.MQTTPublishAsync(ctx, MikrotikRouterMQTTTopicWiFiConnectedMAC.Format(sn), 0, false, login)
		d.MQTTPublishAsync(ctx, MikrotikRouterMQTTTopicWiFiMACState.Format(sn, login), 0, false, true)
	}

	// vpn clients
	connections, err := d.provider.PPPActive(ctx)
	if err != nil {
		return nil, err
	}

	for _, connection := range connections {
		login := mqtt.NameReplace(connection.Name)

		d.MQTTPublishAsync(ctx, MikrotikRouterMQTTTopicVPNConnectedLogin.Format(sn), 0, false, login)
		d.MQTTPublishAsync(ctx, MikrotikRouterMQTTTopicVPNLoginState.Format(sn, login), 0, false, true)
	}

	return nil, nil
}

func (d *MikrotikRouter) taskStateUpdater(ctx context.Context) (interface{}, error) {
	if d.Status() != boggart.DeviceStatusOnline {
		return nil, nil
	}

	serialNumber := d.SerialNumber()
	if serialNumber == "" {
		return nil, nil
	}

	arp, err := d.provider.IPARP(ctx)
	if err != nil && !mikrotik.IsEmptyResponse(err) {
		return nil, err
	}

	dns, err := d.provider.IPDNSStatic(ctx)
	if err != nil && !mikrotik.IsEmptyResponse(err) {
		return nil, err
	}

	leases, err := d.provider.IPDHCPServerLease(ctx)
	if err != nil && !mikrotik.IsEmptyResponse(err) {
		return nil, err
	}

	// Wifi clients
	clients, err := d.provider.InterfaceWirelessRegistrationTable(ctx)
	if err == nil {
		metricRouterMikrotikWifiClients.With("serial_number", serialNumber).Set(float64(len(clients)))

		for _, client := range clients {
			bytes := strings.Split(client.Bytes, ",")
			if len(bytes) != 2 {
				return nil, err
			}

			name := mikrotik.GetNameByMac(client.MacAddress, arp, dns, leases)

			sent, err := strconv.ParseFloat(bytes[0], 64)
			if err != nil {
				return nil, err
			}

			received, err := strconv.ParseFloat(bytes[1], 64)
			if err == nil {
				metricRouterMikrotikTrafficReceivedBytes.With("serial_number", serialNumber).With(
					"interface", client.Interface,
					"mac", client.MacAddress,
					"name", name).Set(received)
				metricRouterMikrotikTrafficSentBytes.With("serial_number", serialNumber).With(
					"interface", client.Interface,
					"mac", client.MacAddress,
					"name", name).Set(sent)
			} else if !mikrotik.IsEmptyResponse(err) {
				return nil, err
			}
		}
	} else if !mikrotik.IsEmptyResponse(err) {
		return nil, err
	}

	// Ports on mikrotik
	stats, err := d.provider.InterfaceStats(ctx)
	if err == nil {
		for _, stat := range stats {
			metricRouterMikrotikTrafficReceivedBytes.With("serial_number", serialNumber).With(
				"interface", stat.Name,
				"mac", stat.MacAddress).Set(float64(stat.RXByte))
			metricRouterMikrotikTrafficSentBytes.With("serial_number", serialNumber).With(
				"interface", stat.Name,
				"mac", stat.MacAddress).Set(float64(stat.TXByte))
		}
	} else if !mikrotik.IsEmptyResponse(err) {
		return nil, err
	}

	resource, err := d.provider.SystemResource(ctx)
	if err == nil {
		metricRouterMikrotikCPULoad.With("serial_number", serialNumber).Set(float64(resource.CPULoad))
		metricRouterMikrotikMemoryAvailable.With("serial_number", serialNumber).Set(float64(resource.FreeMemory))
		metricRouterMikrotikMemoryUsage.With("serial_number", serialNumber).Set(float64(resource.TotalMemory - resource.FreeMemory))
		metricRouterMikrotikStorageAvailable.With("serial_number", serialNumber).Set(float64(resource.FreeHDDSpace))
		metricRouterMikrotikStorageUsage.With("serial_number", serialNumber).Set(float64(resource.TotalHDDSpace - resource.FreeHDDSpace))
	} else if !mikrotik.IsEmptyResponse(err) {
		return nil, err
	}

	disks, err := d.provider.SystemDisk(ctx)
	if err == nil {
		for _, disk := range disks {
			metricRouterMikrotikDiskUsage.With("serial_number", serialNumber).With(
				"name", disk.Name,
				"label", disk.Label,
			).Set(float64(disk.Size - disk.Free))
			metricRouterMikrotikDiskAvailable.With("serial_number", serialNumber).With(
				"name", disk.Name,
				"label", disk.Label,
			).Set(float64(disk.Free))
		}
	} else if !mikrotik.IsEmptyResponse(err) {
		return nil, err
	}

	health, err := d.provider.SystemHealth(ctx)
	if err == nil {
		metricRouterMikrotikVoltage.With("serial_number", serialNumber).Set(health.Voltage)
		metricRouterMikrotikTemperature.With("serial_number", serialNumber).Set(float64(health.Temperature))
	} else if !mikrotik.IsEmptyResponse(err) {
		return nil, err
	}

	return nil, nil
}

func (d *MikrotikRouter) MQTTTopics() []mqtt.Topic {
	return []mqtt.Topic{
		MikrotikRouterMQTTTopicWiFiMACState,
		MikrotikRouterMQTTTopicWiFiConnectedMAC,
		MikrotikRouterMQTTTopicWiFiDisconnectedMAC,
		MikrotikRouterMQTTTopicVPNLoginState,
		MikrotikRouterMQTTTopicVPNConnectedLogin,
		MikrotikRouterMQTTTopicVPNDisconnectedLogin,
	}
}

func (l *MikrotikRouterListener) Events() []workers.Event {
	return []workers.Event{
		boggart.DeviceEventSyslogReceive,
	}
}

func (l *MikrotikRouterListener) Name() string {
	return boggart.ComponentName + ".device.router.mikrotik." + l.router.SerialNumber()
}

func (l *MikrotikRouterListener) Run(ctx context.Context, event workers.Event, t time.Time, args ...interface{}) {
	switch event {
	case boggart.DeviceEventSyslogReceive:
		message := args[0].(map[string]interface{})

		client, ok := message["client"]
		if !ok || client != l.router.syslogClient {
			return
		}

		tag, ok := message["tag"]
		if !ok {
			return
		}

		content, ok := message["content"]
		if !ok {
			return
		}

		switch tag {
		case "wifi":
			check := wifiClientRegexp.FindStringSubmatch(content.(string))
			if len(check) < 4 {
				return
			}

			if _, err := net.ParseMAC(check[1]); err != nil {
				return
			}

			mac, err := l.router.Mac(ctx, check[1])
			if err != nil {
				return
			}

			sn := l.router.SerialNumber()
			login := mqtt.NameReplace(mac.Address)

			switch check[3] {
			case "connected":
				l.router.MQTTPublishAsync(ctx, MikrotikRouterMQTTTopicWiFiConnectedMAC.Format(sn), 0, false, login)
				l.router.MQTTPublishAsync(ctx, MikrotikRouterMQTTTopicWiFiMACState.Format(sn, login), 0, false, []byte(`1`))
			case "disconnected":
				l.router.MQTTPublishAsync(ctx, MikrotikRouterMQTTTopicWiFiDisconnectedMAC.Format(sn), 0, false, login)
				l.router.MQTTPublishAsync(ctx, MikrotikRouterMQTTTopicWiFiMACState.Format(sn, login), 0, false, []byte(`0`))
			}

		case "vpn":
			check := vpnClientRegexp.FindStringSubmatch(content.(string))
			if len(check) < 2 {
				return
			}

			sn := l.router.SerialNumber()
			login := mqtt.NameReplace(check[1])

			switch check[2] {
			case "in":
				l.router.MQTTPublishAsync(ctx, MikrotikRouterMQTTTopicVPNConnectedLogin.Format(sn), 0, false, login)
				l.router.MQTTPublishAsync(ctx, MikrotikRouterMQTTTopicVPNLoginState.Format(sn, login), 0, false, []byte(`1`))
			case "out":
				l.router.MQTTPublishAsync(ctx, MikrotikRouterMQTTTopicVPNDisconnectedLogin.Format(sn), 0, false, login)
				l.router.MQTTPublishAsync(ctx, MikrotikRouterMQTTTopicVPNLoginState.Format(sn, login), 0, false, []byte(`0`))
			}
		}
	}
}
