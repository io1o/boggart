// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewGetGlobalConfigParams creates a new GetGlobalConfigParams object
// with the default values initialized.
func NewGetGlobalConfigParams() *GetGlobalConfigParams {

	return &GetGlobalConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGlobalConfigParamsWithTimeout creates a new GetGlobalConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGlobalConfigParamsWithTimeout(timeout time.Duration) *GetGlobalConfigParams {

	return &GetGlobalConfigParams{

		timeout: timeout,
	}
}

// NewGetGlobalConfigParamsWithContext creates a new GetGlobalConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGlobalConfigParamsWithContext(ctx context.Context) *GetGlobalConfigParams {

	return &GetGlobalConfigParams{

		Context: ctx,
	}
}

// NewGetGlobalConfigParamsWithHTTPClient creates a new GetGlobalConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGlobalConfigParamsWithHTTPClient(client *http.Client) *GetGlobalConfigParams {

	return &GetGlobalConfigParams{
		HTTPClient: client,
	}
}

/*GetGlobalConfigParams contains all the parameters to send to the API endpoint
for the get global config operation typically these are written to a http.Request
*/
type GetGlobalConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get global config params
func (o *GetGlobalConfigParams) WithTimeout(timeout time.Duration) *GetGlobalConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get global config params
func (o *GetGlobalConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get global config params
func (o *GetGlobalConfigParams) WithContext(ctx context.Context) *GetGlobalConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get global config params
func (o *GetGlobalConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get global config params
func (o *GetGlobalConfigParams) WithHTTPClient(client *http.Client) *GetGlobalConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get global config params
func (o *GetGlobalConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetGlobalConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
