// Code generated by go-swagger; DO NOT EDIT.

package global

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

// NewGetGlobalModuleSwitchParams creates a new GetGlobalModuleSwitchParams object
// with the default values initialized.
func NewGetGlobalModuleSwitchParams() *GetGlobalModuleSwitchParams {

	return &GetGlobalModuleSwitchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGlobalModuleSwitchParamsWithTimeout creates a new GetGlobalModuleSwitchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGlobalModuleSwitchParamsWithTimeout(timeout time.Duration) *GetGlobalModuleSwitchParams {

	return &GetGlobalModuleSwitchParams{

		timeout: timeout,
	}
}

// NewGetGlobalModuleSwitchParamsWithContext creates a new GetGlobalModuleSwitchParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGlobalModuleSwitchParamsWithContext(ctx context.Context) *GetGlobalModuleSwitchParams {

	return &GetGlobalModuleSwitchParams{

		Context: ctx,
	}
}

// NewGetGlobalModuleSwitchParamsWithHTTPClient creates a new GetGlobalModuleSwitchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGlobalModuleSwitchParamsWithHTTPClient(client *http.Client) *GetGlobalModuleSwitchParams {

	return &GetGlobalModuleSwitchParams{
		HTTPClient: client,
	}
}

/*GetGlobalModuleSwitchParams contains all the parameters to send to the API endpoint
for the get global module switch operation typically these are written to a http.Request
*/
type GetGlobalModuleSwitchParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get global module switch params
func (o *GetGlobalModuleSwitchParams) WithTimeout(timeout time.Duration) *GetGlobalModuleSwitchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get global module switch params
func (o *GetGlobalModuleSwitchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get global module switch params
func (o *GetGlobalModuleSwitchParams) WithContext(ctx context.Context) *GetGlobalModuleSwitchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get global module switch params
func (o *GetGlobalModuleSwitchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get global module switch params
func (o *GetGlobalModuleSwitchParams) WithHTTPClient(client *http.Client) *GetGlobalModuleSwitchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get global module switch params
func (o *GetGlobalModuleSwitchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetGlobalModuleSwitchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
