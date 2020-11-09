// Code generated by go-swagger; DO NOT EDIT.

package monitoring

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new monitoring API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for monitoring API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetMonitoringStatus(params *GetMonitoringStatusParams) (*GetMonitoringStatusOK, error)

	GetMonitoringTrafficStatistics(params *GetMonitoringTrafficStatisticsParams) (*GetMonitoringTrafficStatisticsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetMonitoringStatus get monitoring status API
*/
func (a *Client) GetMonitoringStatus(params *GetMonitoringStatusParams) (*GetMonitoringStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMonitoringStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMonitoringStatus",
		Method:             "GET",
		PathPattern:        "/api/monitoring/status",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMonitoringStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMonitoringStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetMonitoringStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetMonitoringTrafficStatistics get monitoring traffic statistics API
*/
func (a *Client) GetMonitoringTrafficStatistics(params *GetMonitoringTrafficStatisticsParams) (*GetMonitoringTrafficStatisticsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMonitoringTrafficStatisticsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMonitoringTrafficStatistics",
		Method:             "GET",
		PathPattern:        "/api/monitoring/traffic-statistics",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMonitoringTrafficStatisticsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMonitoringTrafficStatisticsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetMonitoringTrafficStatisticsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
