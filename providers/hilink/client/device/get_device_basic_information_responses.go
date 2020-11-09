// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kihamo/boggart/providers/hilink/models"
)

// GetDeviceBasicInformationReader is a Reader for the GetDeviceBasicInformation structure.
type GetDeviceBasicInformationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceBasicInformationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceBasicInformationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDeviceBasicInformationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceBasicInformationOK creates a GetDeviceBasicInformationOK with default headers values
func NewGetDeviceBasicInformationOK() *GetDeviceBasicInformationOK {
	return &GetDeviceBasicInformationOK{}
}

/*GetDeviceBasicInformationOK handles this case with default header values.

Successful operation
*/
type GetDeviceBasicInformationOK struct {
	Payload *models.DeviceBasicInformation
}

func (o *GetDeviceBasicInformationOK) Error() string {
	return fmt.Sprintf("[GET /api/device/basic_information][%d] getDeviceBasicInformationOK  %+v", 200, o.Payload)
}

func (o *GetDeviceBasicInformationOK) GetPayload() *models.DeviceBasicInformation {
	return o.Payload
}

func (o *GetDeviceBasicInformationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBasicInformation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceBasicInformationDefault creates a GetDeviceBasicInformationDefault with default headers values
func NewGetDeviceBasicInformationDefault(code int) *GetDeviceBasicInformationDefault {
	return &GetDeviceBasicInformationDefault{
		_statusCode: code,
	}
}

/*GetDeviceBasicInformationDefault handles this case with default header values.

Unexpected error
*/
type GetDeviceBasicInformationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get device basic information default response
func (o *GetDeviceBasicInformationDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceBasicInformationDefault) Error() string {
	return fmt.Sprintf("[GET /api/device/basic_information][%d] getDeviceBasicInformation default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceBasicInformationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDeviceBasicInformationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
