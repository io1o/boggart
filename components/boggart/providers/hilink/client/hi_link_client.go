// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kihamo/boggart/components/boggart/providers/hilink/client/device"
	"github.com/kihamo/boggart/components/boggart/providers/hilink/client/sms"
	"github.com/kihamo/boggart/components/boggart/providers/hilink/client/ussd"
	"github.com/kihamo/boggart/components/boggart/providers/hilink/client/web_server"
)

// Default hi link HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new hi link HTTP client.
func NewHTTPClient(formats strfmt.Registry) *HiLink {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new hi link HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *HiLink {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new hi link client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *HiLink {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(HiLink)
	cli.Transport = transport

	cli.Device = device.New(transport, formats)

	cli.Sms = sms.New(transport, formats)

	cli.Ussd = ussd.New(transport, formats)

	cli.WebServer = web_server.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// HiLink is a client for hi link
type HiLink struct {
	Device *device.Client

	Sms *sms.Client

	Ussd *ussd.Client

	WebServer *web_server.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *HiLink) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Device.SetTransport(transport)

	c.Sms.SetTransport(transport)

	c.Ussd.SetTransport(transport)

	c.WebServer.SetTransport(transport)

}
