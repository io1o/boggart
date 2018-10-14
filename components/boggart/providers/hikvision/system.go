package hikvision

import (
	"context"
	"net/http"
	"time"

	"github.com/kihamo/shadow/components/tracing"
	"github.com/opentracing/opentracing-go"
)

type SystemDeviceInfoResponse struct {
	DeviceName           string `xml:"deviceName"`
	DeviceID             string `xml:"deviceID"`
	DeviceDescription    string `xml:"deviceDescription"`
	DeviceLocation       string `xml:"deviceLocation"`
	SystemContact        string `xml:"systemContact"`
	Model                string `xml:"model"`
	SerialNumber         string `xml:"serialNumber"`
	MacAddress           string `xml:"macAddress"`
	FirmwareVersion      string `xml:"firmwareVersion"`
	FirmwareVersionInfo  string `xml:"firmwareVersionInfo"` // from real device (DS-2DE5220IW-AE)
	FirmwareReleasedDate string `xml:"firmwareReleasedDate"`
	BootVersion          string `xml:"bootVersion"`
	BootReleasedDate     string `xml:"bootReleasedDate"`
	HardwareVersion      string `xml:"hardwareVersion"`
	EncoderVersion       string `xml:"encoderVersion"`
	EncoderReleasedDate  string `xml:"encoderReleasedDate"`
	DecoderVersion       string `xml:"decoderVersion"`
	DecoderReleasedDate  string `xml:"decoderReleasedDate"`
	DeviceType           string `xml:"deviceType"`
	TelecontrolID        uint64 `xml:"telecontrolID"`
	SupportBeep          bool   `xml:"supportBeep"`
	SupportVideoLoss     bool   `xml:"supportVideoLoss"` // from real device (DS-2DE5220IW-AE)
}

type SystemStatusResponse struct {
	CurrentDeviceTime time.Time `xml:"currentDeviceTime"`
	DeviceUpTime      uint64    `xml:"deviceUpTime"`
	Memory            []struct {
		MemoryDescription string          `xml:"memoryDescription"`
		MemoryUsage       overrideFloat64 `xml:"memoryUsage"`
		MemoryAvailable   overrideFloat64 `xml:"memoryAvailable"`
	} `xml:"MemoryList>Memory"`
}

func (a *ISAPI) SystemDeviceInfo(ctx context.Context) (result SystemDeviceInfoResponse, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, ComponentName+".system.device_info")
	defer span.Finish()

	err = a.DoXML(ctx, http.MethodGet, a.address+"/System/deviceInfo", nil, &result)
	if err != nil {
		tracing.SpanError(span, err)
	}

	return result, err
}

func (a *ISAPI) SystemStatus(ctx context.Context) (result SystemStatusResponse, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, ComponentName+".system.status")
	defer span.Finish()

	err = a.DoXML(ctx, http.MethodGet, a.address+"/System/status", nil, &result)
	if err != nil {
		tracing.SpanError(span, err)
	}

	return result, err
}
