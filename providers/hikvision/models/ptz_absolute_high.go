// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PtzAbsoluteHigh ptz absolute high
//
// swagger:model PtzAbsoluteHigh
type PtzAbsoluteHigh struct {

	// azimuth
	// Maximum: 3600
	// Minimum: 0
	Azimuth *uint64 `json:"azimuth,omitempty" xml:"azimuth,omitempty"`

	// elevation
	// Maximum: 2700
	// Minimum: -900
	Elevation *int64 `json:"elevation,omitempty" xml:"elevation,omitempty"`

	// zoom
	// Maximum: 1000
	// Minimum: 0
	Zoom *uint64 `json:"zoom,omitempty" xml:"absoluteZoom,omitempty"`
}

// Validate validates this ptz absolute high
func (m *PtzAbsoluteHigh) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzimuth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElevation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PtzAbsoluteHigh) validateAzimuth(formats strfmt.Registry) error {

	if swag.IsZero(m.Azimuth) { // not required
		return nil
	}

	if err := validate.MinimumInt("azimuth", "body", int64(*m.Azimuth), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("azimuth", "body", int64(*m.Azimuth), 3600, false); err != nil {
		return err
	}

	return nil
}

func (m *PtzAbsoluteHigh) validateElevation(formats strfmt.Registry) error {

	if swag.IsZero(m.Elevation) { // not required
		return nil
	}

	if err := validate.MinimumInt("elevation", "body", int64(*m.Elevation), -900, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("elevation", "body", int64(*m.Elevation), 2700, false); err != nil {
		return err
	}

	return nil
}

func (m *PtzAbsoluteHigh) validateZoom(formats strfmt.Registry) error {

	if swag.IsZero(m.Zoom) { // not required
		return nil
	}

	if err := validate.MinimumInt("zoom", "body", int64(*m.Zoom), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("zoom", "body", int64(*m.Zoom), 1000, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PtzAbsoluteHigh) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PtzAbsoluteHigh) UnmarshalBinary(b []byte) error {
	var res PtzAbsoluteHigh
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
