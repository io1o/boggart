// Code generated by go-swagger; DO NOT EDIT.

package connection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// SendConnectionCommandReader is a Reader for the SendConnectionCommand structure.
type SendConnectionCommandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendConnectionCommandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSendConnectionCommandNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSendConnectionCommandBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendConnectionCommandNoContent creates a SendConnectionCommandNoContent with default headers values
func NewSendConnectionCommandNoContent() *SendConnectionCommandNoContent {
	return &SendConnectionCommandNoContent{}
}

/*SendConnectionCommandNoContent handles this case with default header values.

Successful operation
*/
type SendConnectionCommandNoContent struct {
}

func (o *SendConnectionCommandNoContent) Error() string {
	return fmt.Sprintf("[POST /api/connection][%d] sendConnectionCommandNoContent ", 204)
}

func (o *SendConnectionCommandNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendConnectionCommandBadRequest creates a SendConnectionCommandBadRequest with default headers values
func NewSendConnectionCommandBadRequest() *SendConnectionCommandBadRequest {
	return &SendConnectionCommandBadRequest{}
}

/*SendConnectionCommandBadRequest handles this case with default header values.

Bad request
*/
type SendConnectionCommandBadRequest struct {
}

func (o *SendConnectionCommandBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/connection][%d] sendConnectionCommandBadRequest ", 400)
}

func (o *SendConnectionCommandBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*SendConnectionCommandBody send connection command body
swagger:model SendConnectionCommandBody
*/
type SendConnectionCommandBody struct {

	// autoconnect
	Autoconnect bool `json:"autoconnect,omitempty"`

	// baudrate
	Baudrate int64 `json:"baudrate,omitempty"`

	// command
	// Enum: [connect disconnect fake_ack]
	Command string `json:"command,omitempty"`

	// port
	Port string `json:"port,omitempty"`

	// printer profile
	PrinterProfile string `json:"printerProfile,omitempty"`

	// save
	Save bool `json:"save,omitempty"`
}

// Validate validates this send connection command body
func (o *SendConnectionCommandBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommand(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sendConnectionCommandBodyTypeCommandPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["connect","disconnect","fake_ack"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sendConnectionCommandBodyTypeCommandPropEnum = append(sendConnectionCommandBodyTypeCommandPropEnum, v)
	}
}

const (

	// SendConnectionCommandBodyCommandConnect captures enum value "connect"
	SendConnectionCommandBodyCommandConnect string = "connect"

	// SendConnectionCommandBodyCommandDisconnect captures enum value "disconnect"
	SendConnectionCommandBodyCommandDisconnect string = "disconnect"

	// SendConnectionCommandBodyCommandFakeAck captures enum value "fake_ack"
	SendConnectionCommandBodyCommandFakeAck string = "fake_ack"
)

// prop value enum
func (o *SendConnectionCommandBody) validateCommandEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sendConnectionCommandBodyTypeCommandPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *SendConnectionCommandBody) validateCommand(formats strfmt.Registry) error {

	if swag.IsZero(o.Command) { // not required
		return nil
	}

	// value enum
	if err := o.validateCommandEnum("body"+"."+"command", "body", o.Command); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SendConnectionCommandBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SendConnectionCommandBody) UnmarshalBinary(b []byte) error {
	var res SendConnectionCommandBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
