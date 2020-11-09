// Code generated by go-swagger; DO NOT EDIT.

package sms

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
)

// NewGetSMSCountParams creates a new GetSMSCountParams object
// with the default values initialized.
func NewGetSMSCountParams() *GetSMSCountParams {

	return &GetSMSCountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSMSCountParamsWithTimeout creates a new GetSMSCountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSMSCountParamsWithTimeout(timeout time.Duration) *GetSMSCountParams {

	return &GetSMSCountParams{

		timeout: timeout,
	}
}

// NewGetSMSCountParamsWithContext creates a new GetSMSCountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSMSCountParamsWithContext(ctx context.Context) *GetSMSCountParams {

	return &GetSMSCountParams{

		Context: ctx,
	}
}

// NewGetSMSCountParamsWithHTTPClient creates a new GetSMSCountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSMSCountParamsWithHTTPClient(client *http.Client) *GetSMSCountParams {

	return &GetSMSCountParams{
		HTTPClient: client,
	}
}

/*GetSMSCountParams contains all the parameters to send to the API endpoint
for the get s m s count operation typically these are written to a http.Request
*/
type GetSMSCountParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get s m s count params
func (o *GetSMSCountParams) WithTimeout(timeout time.Duration) *GetSMSCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get s m s count params
func (o *GetSMSCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get s m s count params
func (o *GetSMSCountParams) WithContext(ctx context.Context) *GetSMSCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get s m s count params
func (o *GetSMSCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get s m s count params
func (o *GetSMSCountParams) WithHTTPClient(client *http.Client) *GetSMSCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get s m s count params
func (o *GetSMSCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSMSCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
