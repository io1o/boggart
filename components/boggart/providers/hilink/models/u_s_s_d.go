// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// USSD u s s d
// swagger:model USSD
type USSD struct {

	// code type
	CodeType string `json:"CodeType,omitempty" xml:"codeType"`

	// content
	Content string `json:"Content,omitempty" xml:"content"`

	// timeout
	Timeout string `json:"Timeout,omitempty" xml:"timeout"`
}

// Validate validates this u s s d
func (m *USSD) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *USSD) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *USSD) UnmarshalBinary(b []byte) error {
	var res USSD
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
