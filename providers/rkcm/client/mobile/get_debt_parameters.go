// Code generated by go-swagger; DO NOT EDIT.

package mobile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDebtParams creates a new GetDebtParams object
// with the default values initialized.
func NewGetDebtParams() *GetDebtParams {
	var ()
	return &GetDebtParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDebtParamsWithTimeout creates a new GetDebtParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDebtParamsWithTimeout(timeout time.Duration) *GetDebtParams {
	var ()
	return &GetDebtParams{

		timeout: timeout,
	}
}

// NewGetDebtParamsWithContext creates a new GetDebtParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDebtParamsWithContext(ctx context.Context) *GetDebtParams {
	var ()
	return &GetDebtParams{

		Context: ctx,
	}
}

// NewGetDebtParamsWithHTTPClient creates a new GetDebtParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDebtParamsWithHTTPClient(client *http.Client) *GetDebtParams {
	var ()
	return &GetDebtParams{
		HTTPClient: client,
	}
}

/*GetDebtParams contains all the parameters to send to the API endpoint
for the get debt operation typically these are written to a http.Request
*/
type GetDebtParams struct {

	/*Phone
	  Phone number

	*/
	Phone string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get debt params
func (o *GetDebtParams) WithTimeout(timeout time.Duration) *GetDebtParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get debt params
func (o *GetDebtParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get debt params
func (o *GetDebtParams) WithContext(ctx context.Context) *GetDebtParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get debt params
func (o *GetDebtParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get debt params
func (o *GetDebtParams) WithHTTPClient(client *http.Client) *GetDebtParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get debt params
func (o *GetDebtParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPhone adds the phone to the get debt params
func (o *GetDebtParams) WithPhone(phone string) *GetDebtParams {
	o.SetPhone(phone)
	return o
}

// SetPhone adds the phone to the get debt params
func (o *GetDebtParams) SetPhone(phone string) {
	o.Phone = phone
}

// WriteToRequest writes these params to a swagger request
func (o *GetDebtParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param phone
	qrPhone := o.Phone
	qPhone := qrPhone
	if qPhone != "" {
		if err := r.SetQueryParam("phone", qPhone); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
