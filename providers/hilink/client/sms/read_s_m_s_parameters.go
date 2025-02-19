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

// NewReadSMSParams creates a new ReadSMSParams object
// with the default values initialized.
func NewReadSMSParams() *ReadSMSParams {
	var ()
	return &ReadSMSParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadSMSParamsWithTimeout creates a new ReadSMSParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadSMSParamsWithTimeout(timeout time.Duration) *ReadSMSParams {
	var ()
	return &ReadSMSParams{

		timeout: timeout,
	}
}

// NewReadSMSParamsWithContext creates a new ReadSMSParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadSMSParamsWithContext(ctx context.Context) *ReadSMSParams {
	var ()
	return &ReadSMSParams{

		Context: ctx,
	}
}

// NewReadSMSParamsWithHTTPClient creates a new ReadSMSParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadSMSParamsWithHTTPClient(client *http.Client) *ReadSMSParams {
	var ()
	return &ReadSMSParams{
		HTTPClient: client,
	}
}

/*ReadSMSParams contains all the parameters to send to the API endpoint
for the read s m s operation typically these are written to a http.Request
*/
type ReadSMSParams struct {

	/*Request*/
	Request ReadSMSBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read s m s params
func (o *ReadSMSParams) WithTimeout(timeout time.Duration) *ReadSMSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read s m s params
func (o *ReadSMSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read s m s params
func (o *ReadSMSParams) WithContext(ctx context.Context) *ReadSMSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read s m s params
func (o *ReadSMSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read s m s params
func (o *ReadSMSParams) WithHTTPClient(client *http.Client) *ReadSMSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read s m s params
func (o *ReadSMSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the read s m s params
func (o *ReadSMSParams) WithRequest(request ReadSMSBody) *ReadSMSParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the read s m s params
func (o *ReadSMSParams) SetRequest(request ReadSMSBody) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *ReadSMSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
