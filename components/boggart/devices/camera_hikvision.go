package devices

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/providers/hikvision"
	"github.com/kihamo/boggart/components/mqtt"
	"github.com/kihamo/go-workers"
	"github.com/kihamo/go-workers/task"
)

const (
	CameraHikVisionIgnoreInterval = time.Second * 5
	CameraMQTTTopicPrefix         = boggart.ComponentName + "/camera/"
)

type cameraHikVisionPTZChannel struct {
	Channel hikvision.PTZChannel
	Status  *hikvision.PTZStatus
}

type CameraHikVision struct {
	boggart.DeviceBase
	boggart.DeviceSerialNumber

	isapi                 *hikvision.ISAPI
	interval              time.Duration
	mutex                 sync.RWMutex
	alertStreamingHistory map[string]time.Time
	mqtt                  mqtt.Component

	ptzChannels map[uint64]cameraHikVisionPTZChannel
}

func NewCameraHikVision(isapi *hikvision.ISAPI, interval time.Duration, m mqtt.Component) *CameraHikVision {
	device := &CameraHikVision{
		isapi:                 isapi,
		interval:              interval,
		alertStreamingHistory: make(map[string]time.Time),
		mqtt: m,
	}
	device.Init()
	device.SetDescription("HikVision camera")

	return device
}

func (d *CameraHikVision) Types() []boggart.DeviceType {
	return []boggart.DeviceType{
		boggart.DeviceTypeCamera,
	}
}

func (d *CameraHikVision) Ping(ctx context.Context) bool {
	_, err := d.isapi.SystemStatus(ctx)
	return err == nil
}

func (d *CameraHikVision) Tasks() []workers.Task {
	taskSerialNumber := task.NewFunctionTillStopTask(d.taskSerialNumber)
	taskSerialNumber.SetTimeout(time.Second * 5)
	taskSerialNumber.SetRepeats(-1)
	taskSerialNumber.SetRepeatInterval(time.Minute)
	taskSerialNumber.SetName("device-camera-hikvision-serial-number")

	taskPTZStatus := task.NewFunctionTillStopTask(d.taskPTZStatus)
	taskPTZStatus.SetTimeout(time.Second * 5)
	taskPTZStatus.SetRepeats(-1)
	taskPTZStatus.SetRepeatInterval(time.Minute)
	taskPTZStatus.SetName("device-camera-hikvision-ptz-status")

	return []workers.Task{
		taskSerialNumber,
		taskPTZStatus,
	}
}

func (d *CameraHikVision) taskSerialNumber(ctx context.Context) (interface{}, error, bool) {
	if !d.IsEnabled() {
		return nil, nil, false
	}

	deviceInfo, err := d.isapi.SystemDeviceInfo(ctx)
	if err != nil {
		return nil, err, false
	}

	if deviceInfo.SerialNumber == "" {
		return nil, errors.New("device returns empty serial number"), false
	}

	d.SetSerialNumber(deviceInfo.SerialNumber)
	d.SetDescription(d.Description() + " with serial number " + deviceInfo.SerialNumber)

	ptzChannels := make(map[uint64]cameraHikVisionPTZChannel, 0)
	if list, err := d.isapi.PTZChannels(ctx); err == nil {
		for _, channel := range list.Channels {
			ptzChannels[channel.ID] = cameraHikVisionPTZChannel{
				Channel: channel,
			}
		}
	}

	d.mutex.Lock()
	d.ptzChannels = ptzChannels
	d.mutex.Unlock()

	if err := d.startAlertStreaming(); err != nil {
		return nil, err, false
	}

	if len(ptzChannels) == 0 {
		return nil, nil, true
	}

	if err := d.mqtt.Subscribe(d.generateMQTTTopic("ptz", "+", "move"), 0, d.callbackMQTTAbsolute); err != nil {
		return nil, err, false
	}

	if err := d.mqtt.Subscribe(d.generateMQTTTopic("ptz", "+", "absolute"), 0, d.callbackMQTTAbsolute); err != nil {
		return nil, err, false
	}

	if err := d.mqtt.Subscribe(d.generateMQTTTopic("ptz", "+", "continuous"), 0, d.callbackMQTTContinuous); err != nil {
		return nil, err, false
	}

	if err := d.mqtt.Subscribe(d.generateMQTTTopic("ptz", "+", "relative"), 0, d.callbackMQTTRelative); err != nil {
		return nil, err, false
	}

	if err := d.mqtt.Subscribe(d.generateMQTTTopic("ptz", "+", "preset"), 0, d.callbackMQTTPreset); err != nil {
		return nil, err, false
	}

	if err := d.mqtt.Subscribe(d.generateMQTTTopic("ptz", "+", "momentary"), 0, d.callbackMQTTMomentary); err != nil {
		return nil, err, false
	}

	return nil, nil, true
}

func (d *CameraHikVision) taskPTZStatus(ctx context.Context) (interface{}, error, bool) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	if !d.IsEnabled() || d.ptzChannels == nil {
		return nil, nil, false
	}

	if len(d.ptzChannels) == 0 {
		return nil, nil, true
	}

	stop := true

	for id, channel := range d.ptzChannels {
		if !channel.Channel.Enabled {
			continue
		}

		if err := d.updateStatusByChannelId(ctx, id); err != nil {
			log.Printf("failed updated status for %s device in channel %d", d.SerialNumber(), id)
			continue
		}

		stop = false
	}

	return nil, nil, stop
}

func (d *CameraHikVision) startAlertStreaming() error {
	ctx := context.Background()

	stream, err := d.isapi.EventNotificationAlertStream(ctx)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case event := <-stream.NextAlert():
				if !d.IsEnabled() || event.EventState != hikvision.EventEventStateActive {
					continue
				}

				id := fmt.Sprintf("%d-%s", event.DynChannelID, event.EventType)

				d.mutex.Lock()
				lastFire, ok := d.alertStreamingHistory[id]
				d.alertStreamingHistory[id] = event.DateTime
				d.mutex.Unlock()

				if !ok || event.DateTime.Sub(lastFire) > CameraHikVisionIgnoreInterval {
					d.TriggerEvent(ctx, boggart.DeviceEventHikvisionEventNotificationAlert, event, d.SerialNumber())
				}

			case _ = <-stream.NextError():
				// TODO: log errors

			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

func (d *CameraHikVision) updateStatusByChannelId(ctx context.Context, channelId uint64) error {
	channel, ok := d.ptzChannels[channelId]
	if !ok {
		return fmt.Errorf("channel %d not found", channelId)
	}

	status, err := d.isapi.PTZStatus(ctx, channelId)
	if err != nil {
		return err
	}

	channelAsString := strconv.FormatUint(channelId, 10)

	if channel.Status == nil || channel.Status.AbsoluteHigh.Elevation != status.AbsoluteHigh.Elevation {
		d.mqtt.Publish(ctx, d.generateMQTTTopic("ptz", channelAsString, "status", "elevation"), 1, false, strconv.FormatInt(status.AbsoluteHigh.Elevation, 10))
	}

	if channel.Status == nil || channel.Status.AbsoluteHigh.Azimuth != status.AbsoluteHigh.Azimuth {
		d.mqtt.Publish(ctx, d.generateMQTTTopic("ptz", channelAsString, "status", "azimuth"), 1, false, strconv.FormatUint(status.AbsoluteHigh.Azimuth, 10))
	}

	if channel.Status == nil || channel.Status.AbsoluteHigh.AbsoluteZoom != status.AbsoluteHigh.AbsoluteZoom {
		d.mqtt.Publish(ctx, d.generateMQTTTopic("ptz", channelAsString, "status", "zoom"), 1, false, strconv.FormatUint(status.AbsoluteHigh.AbsoluteZoom, 10))
	}

	channel.Status = &status
	d.ptzChannels[channelId] = channel

	return nil
}

func (d *CameraHikVision) checkTopic(topic string) (uint64, error) {
	if !d.IsEnabled() {
		return 0, errors.New("device is disabled")
	}

	if d.ptzChannels == nil || len(d.ptzChannels) == 0 {
		return 0, errors.New("channels is empty")
	}

	parts := mqtt.RouteSplit(topic)

	channelId, err := strconv.ParseUint(parts[4], 10, 64)
	if err != nil {
		return 0, err
	}

	_, ok := d.ptzChannels[channelId]
	if !ok {
		return 0, fmt.Errorf("channel %d not found", channelId)
	}

	return channelId, nil
}

func (d *CameraHikVision) generateMQTTTopic(parts ...string) string {
	return CameraMQTTTopicPrefix + strings.Replace(d.SerialNumber(), "/", "_", -1) + "/" + strings.Join(parts, "/")
}

func (d *CameraHikVision) callbackMQTTAbsolute(ctx context.Context, client mqtt.Component, message mqtt.Message) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Callback for device %s for topic %s with body %s failed with error %s",
				d.SerialNumber(), err.Error(), message.Topic(), string(message.Payload()))
		}
	}()

	channelId, err := d.checkTopic(message.Topic())
	if err != nil {
		return
	}

	var request struct {
		Elevation int64  `json:"elevation,omitempty"`
		Azimuth   uint64 `json:"azimuth,omitempty"`
		Zoom      uint64 `json:"zoom,omitempty"`
	}

	err = json.Unmarshal(message.Payload(), &request)
	if err != nil {
		return
	}

	err = d.isapi.PTZAbsolute(ctx, channelId, request.Elevation, request.Azimuth, request.Zoom)
	if err != nil {
		return
	}

	err = d.updateStatusByChannelId(ctx, channelId)
}

func (d *CameraHikVision) callbackMQTTContinuous(ctx context.Context, client mqtt.Component, message mqtt.Message) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Callback for device %s for topic %s with body %s failed with error %s",
				d.SerialNumber(), err.Error(), message.Topic(), string(message.Payload()))
		}
	}()

	channelId, err := d.checkTopic(message.Topic())
	if err != nil {
		return
	}

	var request struct {
		Pan  int64 `json:"pan,omitempty"`
		Tilt int64 `json:"tilt,omitempty"`
		Zoom int64 `json:"zoom,omitempty"`
	}

	err = json.Unmarshal(message.Payload(), &request)
	if err != nil {
		return
	}

	err = d.isapi.PTZContinuous(ctx, channelId, request.Pan, request.Tilt, request.Zoom)
	if err != nil {
		return
	}

	err = d.updateStatusByChannelId(ctx, channelId)
}

func (d *CameraHikVision) callbackMQTTRelative(ctx context.Context, client mqtt.Component, message mqtt.Message) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Callback for device %s for topic %s with body %s failed with error %s",
				d.SerialNumber(), err.Error(), message.Topic(), string(message.Payload()))
		}
	}()

	channelId, err := d.checkTopic(message.Topic())
	if err != nil {
		return
	}

	var request struct {
		X    int64 `xml:"x,omitempty"`
		Y    int64 `xml:"y,omitempty"`
		Zoom int64 `xml:"zoom,omitempty"`
	}

	err = json.Unmarshal(message.Payload(), &request)
	if err != nil {
		return
	}

	err = d.isapi.PTZRelative(ctx, channelId, request.X, request.Y, request.Zoom)
	if err != nil {
		return
	}

	err = d.updateStatusByChannelId(ctx, channelId)
}

func (d *CameraHikVision) callbackMQTTPreset(ctx context.Context, client mqtt.Component, message mqtt.Message) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Callback for device %s for topic %s with body %s failed with error %s",
				d.SerialNumber(), err.Error(), message.Topic(), string(message.Payload()))
		}
	}()

	channelId, err := d.checkTopic(message.Topic())
	if err != nil {
		return
	}

	presetId, err := strconv.ParseUint(string(message.Payload()), 10, 64)
	if err != nil {
		return
	}

	err = d.isapi.PTZPresetGoTo(ctx, channelId, presetId)
	if err != nil {
		return
	}

	err = d.updateStatusByChannelId(ctx, channelId)
}

func (d *CameraHikVision) callbackMQTTMomentary(ctx context.Context, client mqtt.Component, message mqtt.Message) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Callback for device %s for topic %s with body %s failed with error %s",
				d.SerialNumber(), err.Error(), message.Topic(), string(message.Payload()))
		}
	}()

	channelId, err := d.checkTopic(message.Topic())
	if err != nil {
		return
	}

	var request struct {
		Pan      int64         `json:"pan,omitempty"`
		Tilt     int64         `json:"tilt,omitempty"`
		Zoom     int64         `json:"zoom,omitempty"`
		Duration time.Duration `json:"duration,omitempty"`
	}

	err = json.Unmarshal(message.Payload(), &request)
	if err != nil {
		return
	}

	duration := time.Duration(request.Duration) * time.Millisecond
	err = d.isapi.PTZMomentary(ctx, channelId, request.Pan, request.Tilt, request.Zoom, duration)
	if err != nil {
		return
	}

	err = d.updateStatusByChannelId(ctx, channelId)
}

func (d *CameraHikVision) callbackMQTTMove(ctx context.Context, client mqtt.Component, message mqtt.Message) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Callback for device %s for topic %s with body %s failed with error %s",
				d.SerialNumber(), err.Error(), message.Topic(), string(message.Payload()))
		}
	}()

	channelId, err := d.checkTopic(message.Topic())
	if err != nil {
		return
	}

	var pan, tilt, zoom int64

	switch string(message.Payload()) {
	case "right":
		pan = 1
	case "left":
		pan = -1
	case "up":
		tilt = 1
	case "up-right":
		pan = 1
		tilt = 1
	case "up-left":
		pan = -1
		tilt = 1
	case "down":
		tilt = -1
	case "down-right":
		pan = 1
		tilt = -1
	case "down-left":
		pan = -1
		tilt = -1
	case "narrow":
		zoom = 1
	case "wide":
		zoom = -1
	case "stop":
	default:
		err = fmt.Errorf("unknown operation %s", string(message.Payload()))
	}

	err = d.isapi.PTZContinuous(ctx, channelId, pan, tilt, zoom)
	if err != nil {
		return
	}

	err = d.updateStatusByChannelId(ctx, channelId)
}
