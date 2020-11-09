// Code generated by go-swagger; DO NOT EDIT.

package ussd

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kihamo/boggart/providers/hilink/models"
)

// GetUSSDStatusReader is a Reader for the GetUSSDStatus structure.
type GetUSSDStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUSSDStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUSSDStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUSSDStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUSSDStatusOK creates a GetUSSDStatusOK with default headers values
func NewGetUSSDStatusOK() *GetUSSDStatusOK {
	return &GetUSSDStatusOK{}
}

/*GetUSSDStatusOK handles this case with default header values.

Successful operation
*/
type GetUSSDStatusOK struct {
	Payload *models.USSDStatus
}

func (o *GetUSSDStatusOK) Error() string {
	return fmt.Sprintf("[GET /api/ussd/status][%d] getUSSDStatusOK  %+v", 200, o.Payload)
}

func (o *GetUSSDStatusOK) GetPayload() *models.USSDStatus {
	return o.Payload
}

func (o *GetUSSDStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.USSDStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUSSDStatusDefault creates a GetUSSDStatusDefault with default headers values
func NewGetUSSDStatusDefault(code int) *GetUSSDStatusDefault {
	return &GetUSSDStatusDefault{
		_statusCode: code,
	}
}

/*GetUSSDStatusDefault handles this case with default header values.

Unexpected error
*/
type GetUSSDStatusDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get u s s d status default response
func (o *GetUSSDStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetUSSDStatusDefault) Error() string {
	return fmt.Sprintf("[GET /api/ussd/status][%d] getUSSDStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetUSSDStatusDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUSSDStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
