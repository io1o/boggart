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

// GetImageChannelsReader is a Reader for the GetImageChannels structure.
type GetImageChannelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImageChannelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetImageChannelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetImageChannelsOK creates a GetImageChannelsOK with default headers values
func NewGetImageChannelsOK() *GetImageChannelsOK {
	return &GetImageChannelsOK{}
}

/*GetImageChannelsOK handles this case with default header values.

Successful operation
*/
type GetImageChannelsOK struct {
	Payload models.ImageChannels
}

func (o *GetImageChannelsOK) Error() string {
	return fmt.Sprintf("[GET /Image/channels][%d] getImageChannelsOK  %+v", 200, o.Payload)
}

func (o *GetImageChannelsOK) GetPayload() models.ImageChannels {
	return o.Payload
}

func (o *GetImageChannelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
