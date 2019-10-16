// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PrinterState printer state
// swagger:model PrinterState
type PrinterState struct {

	// sd
	Sd *SDState `json:"sd,omitempty"`

	// state
	State *PrinterStateState `json:"state,omitempty"`

	// temperature
	Temperature *PrinterStateTemperature `json:"temperature,omitempty"`
}

// Validate validates this printer state
func (m *PrinterState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemperature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrinterState) validateSd(formats strfmt.Registry) error {

	if swag.IsZero(m.Sd) { // not required
		return nil
	}

	if m.Sd != nil {
		if err := m.Sd.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sd")
			}
			return err
		}
	}

	return nil
}

func (m *PrinterState) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *PrinterState) validateTemperature(formats strfmt.Registry) error {

	if swag.IsZero(m.Temperature) { // not required
		return nil
	}

	if m.Temperature != nil {
		if err := m.Temperature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("temperature")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrinterState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrinterState) UnmarshalBinary(b []byte) error {
	var res PrinterState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrinterStateState printer state state
// swagger:model PrinterStateState
type PrinterStateState struct {

	// flags
	Flags *PrinterStateStateFlags `json:"flags,omitempty"`

	// text
	Text string `json:"text,omitempty"`
}

// Validate validates this printer state state
func (m *PrinterStateState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrinterStateState) validateFlags(formats strfmt.Registry) error {

	if swag.IsZero(m.Flags) { // not required
		return nil
	}

	if m.Flags != nil {
		if err := m.Flags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state" + "." + "flags")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrinterStateState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrinterStateState) UnmarshalBinary(b []byte) error {
	var res PrinterStateState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrinterStateStateFlags printer state state flags
// swagger:model PrinterStateStateFlags
type PrinterStateStateFlags struct {

	// cancelling
	Cancelling bool `json:"cancelling,omitempty"`

	// closed or error
	ClosedOrError bool `json:"closedOrError,omitempty"`

	// error
	Error bool `json:"error,omitempty"`

	// operational
	Operational bool `json:"operational,omitempty"`

	// paused
	Paused bool `json:"paused,omitempty"`

	// pausing
	Pausing bool `json:"pausing,omitempty"`

	// printing
	Printing bool `json:"printing,omitempty"`

	// ready
	Ready bool `json:"ready,omitempty"`

	// sd ready
	SdReady bool `json:"sdReady,omitempty"`
}

// Validate validates this printer state state flags
func (m *PrinterStateStateFlags) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrinterStateStateFlags) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrinterStateStateFlags) UnmarshalBinary(b []byte) error {
	var res PrinterStateStateFlags
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrinterStateTemperature printer state temperature
// swagger:model PrinterStateTemperature
type PrinterStateTemperature struct {

	// bed
	Bed *TemperatureData `json:"bed,omitempty"`

	// history
	History []*PrinterStateTemperatureHistoryItems0 `json:"history"`

	// tool0
	Tool0 *TemperatureData `json:"tool0,omitempty"`

	// tool1
	Tool1 *TemperatureData `json:"tool1,omitempty"`
}

// Validate validates this printer state temperature
func (m *PrinterStateTemperature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTool0(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTool1(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrinterStateTemperature) validateBed(formats strfmt.Registry) error {

	if swag.IsZero(m.Bed) { // not required
		return nil
	}

	if m.Bed != nil {
		if err := m.Bed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("temperature" + "." + "bed")
			}
			return err
		}
	}

	return nil
}

func (m *PrinterStateTemperature) validateHistory(formats strfmt.Registry) error {

	if swag.IsZero(m.History) { // not required
		return nil
	}

	for i := 0; i < len(m.History); i++ {
		if swag.IsZero(m.History[i]) { // not required
			continue
		}

		if m.History[i] != nil {
			if err := m.History[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("temperature" + "." + "history" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrinterStateTemperature) validateTool0(formats strfmt.Registry) error {

	if swag.IsZero(m.Tool0) { // not required
		return nil
	}

	if m.Tool0 != nil {
		if err := m.Tool0.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("temperature" + "." + "tool0")
			}
			return err
		}
	}

	return nil
}

func (m *PrinterStateTemperature) validateTool1(formats strfmt.Registry) error {

	if swag.IsZero(m.Tool1) { // not required
		return nil
	}

	if m.Tool1 != nil {
		if err := m.Tool1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("temperature" + "." + "tool1")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrinterStateTemperature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrinterStateTemperature) UnmarshalBinary(b []byte) error {
	var res PrinterStateTemperature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrinterStateTemperatureHistoryItems0 printer state temperature history items0
// swagger:model PrinterStateTemperatureHistoryItems0
type PrinterStateTemperatureHistoryItems0 struct {

	// bed
	Bed *TemperatureData `json:"bed,omitempty"`

	// time
	Time int64 `json:"time,omitempty"`

	// tool0
	Tool0 *TemperatureData `json:"tool0,omitempty"`

	// tool1
	Tool1 *TemperatureData `json:"tool1,omitempty"`
}

// Validate validates this printer state temperature history items0
func (m *PrinterStateTemperatureHistoryItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTool0(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTool1(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrinterStateTemperatureHistoryItems0) validateBed(formats strfmt.Registry) error {

	if swag.IsZero(m.Bed) { // not required
		return nil
	}

	if m.Bed != nil {
		if err := m.Bed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bed")
			}
			return err
		}
	}

	return nil
}

func (m *PrinterStateTemperatureHistoryItems0) validateTool0(formats strfmt.Registry) error {

	if swag.IsZero(m.Tool0) { // not required
		return nil
	}

	if m.Tool0 != nil {
		if err := m.Tool0.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tool0")
			}
			return err
		}
	}

	return nil
}

func (m *PrinterStateTemperatureHistoryItems0) validateTool1(formats strfmt.Registry) error {

	if swag.IsZero(m.Tool1) { // not required
		return nil
	}

	if m.Tool1 != nil {
		if err := m.Tool1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tool1")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrinterStateTemperatureHistoryItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrinterStateTemperatureHistoryItems0) UnmarshalBinary(b []byte) error {
	var res PrinterStateTemperatureHistoryItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
