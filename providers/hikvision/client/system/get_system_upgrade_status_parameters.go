// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewGetSystemUpgradeStatusParams creates a new GetSystemUpgradeStatusParams object
// with the default values initialized.
func NewGetSystemUpgradeStatusParams() *GetSystemUpgradeStatusParams {

	return &GetSystemUpgradeStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemUpgradeStatusParamsWithTimeout creates a new GetSystemUpgradeStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSystemUpgradeStatusParamsWithTimeout(timeout time.Duration) *GetSystemUpgradeStatusParams {

	return &GetSystemUpgradeStatusParams{

		timeout: timeout,
	}
}

// NewGetSystemUpgradeStatusParamsWithContext creates a new GetSystemUpgradeStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSystemUpgradeStatusParamsWithContext(ctx context.Context) *GetSystemUpgradeStatusParams {

	return &GetSystemUpgradeStatusParams{

		Context: ctx,
	}
}

// NewGetSystemUpgradeStatusParamsWithHTTPClient creates a new GetSystemUpgradeStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSystemUpgradeStatusParamsWithHTTPClient(client *http.Client) *GetSystemUpgradeStatusParams {

	return &GetSystemUpgradeStatusParams{
		HTTPClient: client,
	}
}

/*GetSystemUpgradeStatusParams contains all the parameters to send to the API endpoint
for the get system upgrade status operation typically these are written to a http.Request
*/
type GetSystemUpgradeStatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get system upgrade status params
func (o *GetSystemUpgradeStatusParams) WithTimeout(timeout time.Duration) *GetSystemUpgradeStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system upgrade status params
func (o *GetSystemUpgradeStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system upgrade status params
func (o *GetSystemUpgradeStatusParams) WithContext(ctx context.Context) *GetSystemUpgradeStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system upgrade status params
func (o *GetSystemUpgradeStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system upgrade status params
func (o *GetSystemUpgradeStatusParams) WithHTTPClient(client *http.Client) *GetSystemUpgradeStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system upgrade status params
func (o *GetSystemUpgradeStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemUpgradeStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
