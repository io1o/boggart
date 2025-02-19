// Code generated by go-swagger; DO NOT EDIT.

package printer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kihamo/boggart/providers/octoprint/models"
)

// GetBedStateReader is a Reader for the GetBedState structure.
type GetBedStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBedStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBedStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewGetBedStateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBedStateOK creates a GetBedStateOK with default headers values
func NewGetBedStateOK() *GetBedStateOK {
	return &GetBedStateOK{}
}

/*GetBedStateOK handles this case with default header values.

Successful operation
*/
type GetBedStateOK struct {
	Payload *models.BedState
}

func (o *GetBedStateOK) Error() string {
	return fmt.Sprintf("[GET /api/printer/bed][%d] getBedStateOK  %+v", 200, o.Payload)
}

func (o *GetBedStateOK) GetPayload() *models.BedState {
	return o.Payload
}

func (o *GetBedStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BedState)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBedStateConflict creates a GetBedStateConflict with default headers values
func NewGetBedStateConflict() *GetBedStateConflict {
	return &GetBedStateConflict{}
}

/*GetBedStateConflict handles this case with default header values.

If the printer is not operational
*/
type GetBedStateConflict struct {
}

func (o *GetBedStateConflict) Error() string {
	return fmt.Sprintf("[GET /api/printer/bed][%d] getBedStateConflict ", 409)
}

func (o *GetBedStateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
