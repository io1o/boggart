// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewAddByIdentParams creates a new AddByIdentParams object
// with the default values initialized.
func NewAddByIdentParams() *AddByIdentParams {
	var ()
	return &AddByIdentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddByIdentParamsWithTimeout creates a new AddByIdentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddByIdentParamsWithTimeout(timeout time.Duration) *AddByIdentParams {
	var ()
	return &AddByIdentParams{

		timeout: timeout,
	}
}

// NewAddByIdentParamsWithContext creates a new AddByIdentParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddByIdentParamsWithContext(ctx context.Context) *AddByIdentParams {
	var ()
	return &AddByIdentParams{

		Context: ctx,
	}
}

// NewAddByIdentParamsWithHTTPClient creates a new AddByIdentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddByIdentParamsWithHTTPClient(client *http.Client) *AddByIdentParams {
	var ()
	return &AddByIdentParams{
		HTTPClient: client,
	}
}

/*AddByIdentParams contains all the parameters to send to the API endpoint
for the add by ident operation typically these are written to a http.Request
*/
type AddByIdentParams struct {

	/*Request*/
	Request AddByIdentBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add by ident params
func (o *AddByIdentParams) WithTimeout(timeout time.Duration) *AddByIdentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add by ident params
func (o *AddByIdentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add by ident params
func (o *AddByIdentParams) WithContext(ctx context.Context) *AddByIdentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add by ident params
func (o *AddByIdentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add by ident params
func (o *AddByIdentParams) WithHTTPClient(client *http.Client) *AddByIdentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add by ident params
func (o *AddByIdentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the add by ident params
func (o *AddByIdentParams) WithRequest(request AddByIdentBody) *AddByIdentParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the add by ident params
func (o *AddByIdentParams) SetRequest(request AddByIdentBody) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *AddByIdentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
