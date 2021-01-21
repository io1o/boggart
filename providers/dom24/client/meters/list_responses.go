// Code generated by go-swagger; DO NOT EDIT.

package meters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/kihamo/boggart/providers/dom24/models"
)

// ListReader is a Reader for the List structure.
type ListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOK creates a ListOK with default headers values
func NewListOK() *ListOK {
	return &ListOK{}
}

/*ListOK handles this case with default header values.

Successful operation
*/
type ListOK struct {
	Payload *ListOKBody
}

func (o *ListOK) Error() string {
	return fmt.Sprintf("[GET /Meters/List][%d] listOK  %+v", 200, o.Payload)
}

func (o *ListOK) GetPayload() *ListOKBody {
	return o.Payload
}

func (o *ListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUnauthorized creates a ListUnauthorized with default headers values
func NewListUnauthorized() *ListUnauthorized {
	return &ListUnauthorized{}
}

/*ListUnauthorized handles this case with default header values.

Unauthorized
*/
type ListUnauthorized struct {
	Payload *models.Error
}

func (o *ListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Meters/List][%d] listUnauthorized  %+v", 401, o.Payload)
}

func (o *ListUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDefault creates a ListDefault with default headers values
func NewListDefault(code int) *ListDefault {
	return &ListDefault{
		_statusCode: code,
	}
}

/*ListDefault handles this case with default header values.

Unexpected error
*/
type ListDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list default response
func (o *ListDefault) Code() int {
	return o._statusCode
}

func (o *ListDefault) Error() string {
	return fmt.Sprintf("[GET /Meters/List][%d] list default  %+v", o._statusCode, o.Payload)
}

func (o *ListDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListOKBody list o k body
swagger:model ListOKBody
*/
type ListOKBody struct {

	// data
	Data []*models.Meter `json:"Data"`
}

// Validate validates this list o k body
func (o *ListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listOK" + "." + "Data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListOKBody) UnmarshalBinary(b []byte) error {
	var res ListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
