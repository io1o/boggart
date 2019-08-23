// Code generated by go-swagger; DO NOT EDIT.

package ussd

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kihamo/boggart/components/boggart/providers/hilink/models"
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
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
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
	return fmt.Sprintf("[GET /ussd/status][%d] getUSSDStatusOK  %+v", 200, o.Payload)
}

func (o *GetUSSDStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.USSDStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
