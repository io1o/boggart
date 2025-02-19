// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kihamo/boggart/providers/hikvision/models"
)

// SetImageIrCutFilterReader is a Reader for the SetImageIrCutFilter structure.
type SetImageIrCutFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetImageIrCutFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetImageIrCutFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetImageIrCutFilterOK creates a SetImageIrCutFilterOK with default headers values
func NewSetImageIrCutFilterOK() *SetImageIrCutFilterOK {
	return &SetImageIrCutFilterOK{}
}

/*SetImageIrCutFilterOK handles this case with default header values.

Successful operation
*/
type SetImageIrCutFilterOK struct {
	Payload *models.Status
}

func (o *SetImageIrCutFilterOK) Error() string {
	return fmt.Sprintf("[PUT /Image/channels/{channel}/IrcutFilter][%d] setImageIrCutFilterOK  %+v", 200, o.Payload)
}

func (o *SetImageIrCutFilterOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *SetImageIrCutFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
