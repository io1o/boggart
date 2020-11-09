// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceCompressLogFile device compress log file
//
// swagger:model DeviceCompressLogFile
type DeviceCompressLogFile struct {

	// log path
	LogPath string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
}

// Validate validates this device compress log file
func (m *DeviceCompressLogFile) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceCompressLogFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceCompressLogFile) UnmarshalBinary(b []byte) error {
	var res DeviceCompressLogFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
