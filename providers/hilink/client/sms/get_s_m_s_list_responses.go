// Code generated by go-swagger; DO NOT EDIT.

package sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kihamo/boggart/providers/hilink/models"
)

// GetSMSListReader is a Reader for the GetSMSList structure.
type GetSMSListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSMSListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSMSListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSMSListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSMSListOK creates a GetSMSListOK with default headers values
func NewGetSMSListOK() *GetSMSListOK {
	return &GetSMSListOK{}
}

/*GetSMSListOK handles this case with default header values.

Successful operation
*/
type GetSMSListOK struct {
	Payload *models.SMSList
}

func (o *GetSMSListOK) Error() string {
	return fmt.Sprintf("[POST /api/sms/sms-list][%d] getSMSListOK  %+v", 200, o.Payload)
}

func (o *GetSMSListOK) GetPayload() *models.SMSList {
	return o.Payload
}

func (o *GetSMSListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SMSList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSMSListDefault creates a GetSMSListDefault with default headers values
func NewGetSMSListDefault(code int) *GetSMSListDefault {
	return &GetSMSListDefault{
		_statusCode: code,
	}
}

/*GetSMSListDefault handles this case with default header values.

Unexpected error
*/
type GetSMSListDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get s m s list default response
func (o *GetSMSListDefault) Code() int {
	return o._statusCode
}

func (o *GetSMSListDefault) Error() string {
	return fmt.Sprintf("[POST /api/sms/sms-list][%d] getSMSList default  %+v", o._statusCode, o.Payload)
}

func (o *GetSMSListDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSMSListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
