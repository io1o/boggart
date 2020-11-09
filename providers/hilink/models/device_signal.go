// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceSignal device signal
//
// swagger:model DeviceSignal
type DeviceSignal struct {

	// cell ID
	CellID int64 `json:"CellID,omitempty" xml:"cell_id,omitempty"`

	// e c i o
	ECIO string `json:"ECIO,omitempty" xml:"ecio,omitempty"`

	// l t e band info
	LTEBandInfo string `json:"LTEBandInfo,omitempty" xml:"lte_bandinfo,omitempty"`

	// l t e band width
	LTEBandWidth string `json:"LTEBandWidth,omitempty" xml:"lte_bandwidth,omitempty"`

	// mode
	Mode int64 `json:"Mode,omitempty" xml:"mode,omitempty"`

	// p c i
	PCI int64 `json:"PCI,omitempty" xml:"pci,omitempty"`

	// p s a t t
	PSATT int64 `json:"PSATT,omitempty" xml:"psatt,omitempty"`

	// r s c p
	RSCP string `json:"RSCP,omitempty" xml:"rscp,omitempty"`

	// r s r p
	RSRP string `json:"RSRP,omitempty" xml:"rsrp,omitempty"`

	// r s r q
	RSRQ string `json:"RSRQ,omitempty" xml:"rsrq,omitempty"`

	// r s s i
	RSSI string `json:"RSSI,omitempty" xml:"rssi,omitempty"`

	// s c
	SC int64 `json:"SC,omitempty" xml:"sc,omitempty"`

	// s i n r
	SINR string `json:"SINR,omitempty" xml:"sinr,omitempty"`
}

// Validate validates this device signal
func (m *DeviceSignal) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceSignal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceSignal) UnmarshalBinary(b []byte) error {
	var res DeviceSignal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
