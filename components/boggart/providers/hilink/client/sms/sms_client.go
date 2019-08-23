// Code generated by go-swagger; DO NOT EDIT.

package sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new sms API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sms API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetSMSCount get s m s count API
*/
func (a *Client) GetSMSCount(params *GetSMSCountParams) (*GetSMSCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSMSCountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSMSCount",
		Method:             "GET",
		PathPattern:        "/sms/sms-count",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSMSCountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSMSCountOK), nil

}

/*
GetSMSList get s m s list API
*/
func (a *Client) GetSMSList(params *GetSMSListParams) (*GetSMSListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSMSListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSMSList",
		Method:             "POST",
		PathPattern:        "/sms/sms-list",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSMSListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSMSListOK), nil

}

/*
ReadSMS read s m s API
*/
func (a *Client) ReadSMS(params *ReadSMSParams) (*ReadSMSOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadSMSParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readSMS",
		Method:             "POST",
		PathPattern:        "/sms/set-read",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadSMSReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReadSMSOK), nil

}

/*
RemoveSMS remove s m s API
*/
func (a *Client) RemoveSMS(params *RemoveSMSParams) (*RemoveSMSOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveSMSParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeSMS",
		Method:             "POST",
		PathPattern:        "/sms/delete-sms",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemoveSMSReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveSMSOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
