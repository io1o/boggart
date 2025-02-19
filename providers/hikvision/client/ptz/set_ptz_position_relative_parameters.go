// Code generated by go-swagger; DO NOT EDIT.

package ptz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/kihamo/boggart/providers/hikvision/models"
)

// NewSetPtzPositionRelativeParams creates a new SetPtzPositionRelativeParams object
// with the default values initialized.
func NewSetPtzPositionRelativeParams() *SetPtzPositionRelativeParams {
	var ()
	return &SetPtzPositionRelativeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetPtzPositionRelativeParamsWithTimeout creates a new SetPtzPositionRelativeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetPtzPositionRelativeParamsWithTimeout(timeout time.Duration) *SetPtzPositionRelativeParams {
	var ()
	return &SetPtzPositionRelativeParams{

		timeout: timeout,
	}
}

// NewSetPtzPositionRelativeParamsWithContext creates a new SetPtzPositionRelativeParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetPtzPositionRelativeParamsWithContext(ctx context.Context) *SetPtzPositionRelativeParams {
	var ()
	return &SetPtzPositionRelativeParams{

		Context: ctx,
	}
}

// NewSetPtzPositionRelativeParamsWithHTTPClient creates a new SetPtzPositionRelativeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetPtzPositionRelativeParamsWithHTTPClient(client *http.Client) *SetPtzPositionRelativeParams {
	var ()
	return &SetPtzPositionRelativeParams{
		HTTPClient: client,
	}
}

/*SetPtzPositionRelativeParams contains all the parameters to send to the API endpoint
for the set ptz position relative operation typically these are written to a http.Request
*/
type SetPtzPositionRelativeParams struct {

	/*PTZData*/
	PTZData *models.PTZData
	/*Channel
	  Channel ID

	*/
	Channel uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set ptz position relative params
func (o *SetPtzPositionRelativeParams) WithTimeout(timeout time.Duration) *SetPtzPositionRelativeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set ptz position relative params
func (o *SetPtzPositionRelativeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set ptz position relative params
func (o *SetPtzPositionRelativeParams) WithContext(ctx context.Context) *SetPtzPositionRelativeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set ptz position relative params
func (o *SetPtzPositionRelativeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set ptz position relative params
func (o *SetPtzPositionRelativeParams) WithHTTPClient(client *http.Client) *SetPtzPositionRelativeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set ptz position relative params
func (o *SetPtzPositionRelativeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPTZData adds the pTZData to the set ptz position relative params
func (o *SetPtzPositionRelativeParams) WithPTZData(pTZData *models.PTZData) *SetPtzPositionRelativeParams {
	o.SetPTZData(pTZData)
	return o
}

// SetPTZData adds the pTZData to the set ptz position relative params
func (o *SetPtzPositionRelativeParams) SetPTZData(pTZData *models.PTZData) {
	o.PTZData = pTZData
}

// WithChannel adds the channel to the set ptz position relative params
func (o *SetPtzPositionRelativeParams) WithChannel(channel uint64) *SetPtzPositionRelativeParams {
	o.SetChannel(channel)
	return o
}

// SetChannel adds the channel to the set ptz position relative params
func (o *SetPtzPositionRelativeParams) SetChannel(channel uint64) {
	o.Channel = channel
}

// WriteToRequest writes these params to a swagger request
func (o *SetPtzPositionRelativeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PTZData != nil {
		if err := r.SetBodyParam(o.PTZData); err != nil {
			return err
		}
	}

	// path param channel
	if err := r.SetPathParam("channel", swag.FormatUint64(o.Channel)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
