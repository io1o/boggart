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

// PtzRelative ptz relative
//
// swagger:model PtzRelative
type PtzRelative struct {

	// x
	X int64 `json:"X,omitempty" xml:"Relative>positionX,omitempty"`

	// y
	Y int64 `json:"Y,omitempty" xml:"Relative>positionY,omitempty"`

	// zoom
	// Maximum: 100
	// Minimum: -100
	Zoom *int64 `json:"zoom,omitempty" xml:"Relative>relativeZoom,omitempty"`
}

// Validate validates this ptz relative
func (m *PtzRelative) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateZoom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PtzRelative) validateZoom(formats strfmt.Registry) error {

	if swag.IsZero(m.Zoom) { // not required
		return nil
	}

	if err := validate.MinimumInt("zoom", "body", int64(*m.Zoom), -100, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("zoom", "body", int64(*m.Zoom), 100, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PtzRelative) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PtzRelative) UnmarshalBinary(b []byte) error {
	var res PtzRelative
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
