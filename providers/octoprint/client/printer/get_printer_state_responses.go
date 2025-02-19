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

// GetPrinterStateReader is a Reader for the GetPrinterState structure.
type GetPrinterStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrinterStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrinterStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewGetPrinterStateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrinterStateOK creates a GetPrinterStateOK with default headers values
func NewGetPrinterStateOK() *GetPrinterStateOK {
	return &GetPrinterStateOK{}
}

/*GetPrinterStateOK handles this case with default header values.

Successful operation
*/
type GetPrinterStateOK struct {
	Payload *models.PrinterState
}

func (o *GetPrinterStateOK) Error() string {
	return fmt.Sprintf("[GET /api/printer][%d] getPrinterStateOK  %+v", 200, o.Payload)
}

func (o *GetPrinterStateOK) GetPayload() *models.PrinterState {
	return o.Payload
}

func (o *GetPrinterStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrinterState)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrinterStateConflict creates a GetPrinterStateConflict with default headers values
func NewGetPrinterStateConflict() *GetPrinterStateConflict {
	return &GetPrinterStateConflict{}
}

/*GetPrinterStateConflict handles this case with default header values.

If the printer is not operational
*/
type GetPrinterStateConflict struct {
}

func (o *GetPrinterStateConflict) Error() string {
	return fmt.Sprintf("[GET /api/printer][%d] getPrinterStateConflict ", 409)
}

func (o *GetPrinterStateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
