// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Version version
// swagger:model Version
type Version struct {

	// api
	API string `json:"api,omitempty"`

	// server
	Server string `json:"server,omitempty"`

	// text
	Text string `json:"text,omitempty"`
}

// Validate validates this version
func (m *Version) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Version) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Version) UnmarshalBinary(b []byte) error {
	var res Version
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
