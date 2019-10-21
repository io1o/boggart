// Code generated by go-swagger; DO NOT EDIT.

package sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kihamo/boggart/providers/hilink/models"
)

// ReadSMSReader is a Reader for the ReadSMS structure.
type ReadSMSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadSMSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadSMSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReadSMSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadSMSOK creates a ReadSMSOK with default headers values
func NewReadSMSOK() *ReadSMSOK {
	return &ReadSMSOK{}
}

/*ReadSMSOK handles this case with default header values.

Successful operation
*/
type ReadSMSOK struct {
	Payload string
}

func (o *ReadSMSOK) Error() string {
	return fmt.Sprintf("[POST /api/sms/set-read][%d] readSMSOK  %+v", 200, o.Payload)
}

func (o *ReadSMSOK) GetPayload() string {
	return o.Payload
}

func (o *ReadSMSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadSMSDefault creates a ReadSMSDefault with default headers values
func NewReadSMSDefault(code int) *ReadSMSDefault {
	return &ReadSMSDefault{
		_statusCode: code,
	}
}

/*ReadSMSDefault handles this case with default header values.

Unexpected error
*/
type ReadSMSDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the read s m s default response
func (o *ReadSMSDefault) Code() int {
	return o._statusCode
}

func (o *ReadSMSDefault) Error() string {
	return fmt.Sprintf("[POST /api/sms/set-read][%d] readSMS default  %+v", o._statusCode, o.Payload)
}

func (o *ReadSMSDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReadSMSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ReadSMSBody read s m s body
swagger:model ReadSMSBody
*/
type ReadSMSBody struct {

	// index
	Index int64 `json:"Index,omitempty" xml:"Index"`
}

// Validate validates this read s m s body
func (o *ReadSMSBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ReadSMSBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReadSMSBody) UnmarshalBinary(b []byte) error {
	var res ReadSMSBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
