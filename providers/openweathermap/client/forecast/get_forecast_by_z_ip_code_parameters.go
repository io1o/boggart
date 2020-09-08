// Code generated by go-swagger; DO NOT EDIT.

package forecast

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
)

// NewGetForecastByZIPCodeParams creates a new GetForecastByZIPCodeParams object
// with the default values initialized.
func NewGetForecastByZIPCodeParams() *GetForecastByZIPCodeParams {
	var ()
	return &GetForecastByZIPCodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetForecastByZIPCodeParamsWithTimeout creates a new GetForecastByZIPCodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetForecastByZIPCodeParamsWithTimeout(timeout time.Duration) *GetForecastByZIPCodeParams {
	var ()
	return &GetForecastByZIPCodeParams{

		timeout: timeout,
	}
}

// NewGetForecastByZIPCodeParamsWithContext creates a new GetForecastByZIPCodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetForecastByZIPCodeParamsWithContext(ctx context.Context) *GetForecastByZIPCodeParams {
	var ()
	return &GetForecastByZIPCodeParams{

		Context: ctx,
	}
}

// NewGetForecastByZIPCodeParamsWithHTTPClient creates a new GetForecastByZIPCodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetForecastByZIPCodeParamsWithHTTPClient(client *http.Client) *GetForecastByZIPCodeParams {
	var ()
	return &GetForecastByZIPCodeParams{
		HTTPClient: client,
	}
}

/*GetForecastByZIPCodeParams contains all the parameters to send to the API endpoint
for the get forecast by z IP code operation typically these are written to a http.Request
*/
type GetForecastByZIPCodeParams struct {

	/*Cnt
	  To limit number of listed cities please setup 'cnt' parameter that specifies the number of lines returned

	*/
	Cnt *uint64
	/*Lang
	  Multilingual support

	*/
	Lang *string
	/*Units
	  Standard, metric, and imperial units are available

	*/
	Units *string
	/*Zip
	  ZIP code

	*/
	Zip string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get forecast by z IP code params
func (o *GetForecastByZIPCodeParams) WithTimeout(timeout time.Duration) *GetForecastByZIPCodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get forecast by z IP code params
func (o *GetForecastByZIPCodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get forecast by z IP code params
func (o *GetForecastByZIPCodeParams) WithContext(ctx context.Context) *GetForecastByZIPCodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get forecast by z IP code params
func (o *GetForecastByZIPCodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get forecast by z IP code params
func (o *GetForecastByZIPCodeParams) WithHTTPClient(client *http.Client) *GetForecastByZIPCodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get forecast by z IP code params
func (o *GetForecastByZIPCodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCnt adds the cnt to the get forecast by z IP code params
func (o *GetForecastByZIPCodeParams) WithCnt(cnt *uint64) *GetForecastByZIPCodeParams {
	o.SetCnt(cnt)
	return o
}

// SetCnt adds the cnt to the get forecast by z IP code params
func (o *GetForecastByZIPCodeParams) SetCnt(cnt *uint64) {
	o.Cnt = cnt
}

// WithLang adds the lang to the get forecast by z IP code params
func (o *GetForecastByZIPCodeParams) WithLang(lang *string) *GetForecastByZIPCodeParams {
	o.SetLang(lang)
	return o
}

// SetLang adds the lang to the get forecast by z IP code params
func (o *GetForecastByZIPCodeParams) SetLang(lang *string) {
	o.Lang = lang
}

// WithUnits adds the units to the get forecast by z IP code params
func (o *GetForecastByZIPCodeParams) WithUnits(units *string) *GetForecastByZIPCodeParams {
	o.SetUnits(units)
	return o
}

// SetUnits adds the units to the get forecast by z IP code params
func (o *GetForecastByZIPCodeParams) SetUnits(units *string) {
	o.Units = units
}

// WithZip adds the zip to the get forecast by z IP code params
func (o *GetForecastByZIPCodeParams) WithZip(zip string) *GetForecastByZIPCodeParams {
	o.SetZip(zip)
	return o
}

// SetZip adds the zip to the get forecast by z IP code params
func (o *GetForecastByZIPCodeParams) SetZip(zip string) {
	o.Zip = zip
}

// WriteToRequest writes these params to a swagger request
func (o *GetForecastByZIPCodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cnt != nil {

		// query param cnt
		var qrCnt uint64
		if o.Cnt != nil {
			qrCnt = *o.Cnt
		}
		qCnt := swag.FormatUint64(qrCnt)
		if qCnt != "" {
			if err := r.SetQueryParam("cnt", qCnt); err != nil {
				return err
			}
		}

	}

	if o.Lang != nil {

		// query param lang
		var qrLang string
		if o.Lang != nil {
			qrLang = *o.Lang
		}
		qLang := qrLang
		if qLang != "" {
			if err := r.SetQueryParam("lang", qLang); err != nil {
				return err
			}
		}

	}

	if o.Units != nil {

		// query param units
		var qrUnits string
		if o.Units != nil {
			qrUnits = *o.Units
		}
		qUnits := qrUnits
		if qUnits != "" {
			if err := r.SetQueryParam("units", qUnits); err != nil {
				return err
			}
		}

	}

	// query param zip
	qrZip := o.Zip
	qZip := qrZip
	if qZip != "" {
		if err := r.SetQueryParam("zip", qZip); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
