// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceBasicInformation device basic information
//
// swagger:model DeviceBasicInformation
type DeviceBasicInformation struct {

	// autoupdate guide status
	AutoupdateGuideStatus bool `json:"AutoupdateGuideStatus,omitempty" xml:"autoupdate_guide_status,omitempty"`

	// classify
	Classify string `json:"Classify,omitempty" xml:"classify,omitempty"`

	// device name
	DeviceName string `json:"DeviceName,omitempty" xml:"devicename,omitempty"`

	// multi mode
	MultiMode bool `json:"MultiMode,omitempty" xml:"multimode,omitempty"`

	// product family
	ProductFamily string `json:"ProductFamily,omitempty" xml:"productfamily,omitempty"`

	// restore default status
	RestoreDefaultStatus bool `json:"RestoreDefaultStatus,omitempty" xml:"restore_default_status,omitempty"`

	// sim save pin enable
	SimSavePinEnable bool `json:"SimSavePinEnable,omitempty" xml:"sim_save_pin_enable,omitempty"`

	// software version
	SoftwareVersion string `json:"SoftwareVersion,omitempty" xml:"SoftwareVersion,omitempty"`

	// web UI version
	WebUIVersion string `json:"WebUIVersion,omitempty" xml:"WebUIVersion,omitempty"`
}

// Validate validates this device basic information
func (m *DeviceBasicInformation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceBasicInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceBasicInformation) UnmarshalBinary(b []byte) error {
	var res DeviceBasicInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
