// Code generated by go-swagger; DO NOT EDIT.

package net

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

// NewGetCurrentPLMNParams creates a new GetCurrentPLMNParams object
// with the default values initialized.
func NewGetCurrentPLMNParams() *GetCurrentPLMNParams {

	return &GetCurrentPLMNParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCurrentPLMNParamsWithTimeout creates a new GetCurrentPLMNParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCurrentPLMNParamsWithTimeout(timeout time.Duration) *GetCurrentPLMNParams {

	return &GetCurrentPLMNParams{

		timeout: timeout,
	}
}

// NewGetCurrentPLMNParamsWithContext creates a new GetCurrentPLMNParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCurrentPLMNParamsWithContext(ctx context.Context) *GetCurrentPLMNParams {

	return &GetCurrentPLMNParams{

		Context: ctx,
	}
}

// NewGetCurrentPLMNParamsWithHTTPClient creates a new GetCurrentPLMNParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCurrentPLMNParamsWithHTTPClient(client *http.Client) *GetCurrentPLMNParams {

	return &GetCurrentPLMNParams{
		HTTPClient: client,
	}
}

/*GetCurrentPLMNParams contains all the parameters to send to the API endpoint
for the get current p l m n operation typically these are written to a http.Request
*/
type GetCurrentPLMNParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get current p l m n params
func (o *GetCurrentPLMNParams) WithTimeout(timeout time.Duration) *GetCurrentPLMNParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get current p l m n params
func (o *GetCurrentPLMNParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get current p l m n params
func (o *GetCurrentPLMNParams) WithContext(ctx context.Context) *GetCurrentPLMNParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get current p l m n params
func (o *GetCurrentPLMNParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get current p l m n params
func (o *GetCurrentPLMNParams) WithHTTPClient(client *http.Client) *GetCurrentPLMNParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get current p l m n params
func (o *GetCurrentPLMNParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCurrentPLMNParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
