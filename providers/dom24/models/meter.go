// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Meter meter
//
// swagger:model Meter
type Meter struct {

	// address
	Address string `json:"Address,omitempty"`

	// auto value getting only
	AutoValueGettingOnly bool `json:"AutoValueGettingOnly,omitempty"`

	// custom name
	CustomName string `json:"CustomName,omitempty"`

	// factory number
	FactoryNumber string `json:"FactoryNumber,omitempty"`

	// house Id
	HouseID uint64 `json:"HouseId,omitempty"`

	// ID
	ID uint64 `json:"ID,omitempty"`

	// is disabled
	IsDisabled bool `json:"IsDisabled,omitempty"`

	// last checkup date
	LastCheckupDate string `json:"LastCheckupDate,omitempty"`

	// next checkup date
	NextCheckupDate string `json:"NextCheckupDate,omitempty"`

	// number of decimal places
	NumberOfDecimalPlaces uint64 `json:"NumberOfDecimalPlaces,omitempty"`

	// number of integer part
	NumberOfIntegerPart uint64 `json:"NumberOfIntegerPart,omitempty"`

	// period message
	PeriodMessage string `json:"PeriodMessage,omitempty"`

	// recheck interval
	RecheckInterval uint64 `json:"RecheckInterval,omitempty"`

	// resource
	Resource string `json:"Resource,omitempty"`

	// start date
	StartDate string `json:"StartDate,omitempty"`

	// start value
	StartValue float64 `json:"StartValue,omitempty"`

	// start value t2
	StartValueT2 float64 `json:"StartValueT2,omitempty"`

	// start value t3
	StartValueT3 float64 `json:"StartValueT3,omitempty"`

	// tariff1 name
	Tariff1Name string `json:"Tariff1Name,omitempty"`

	// tariff2 name
	Tariff2Name string `json:"Tariff2Name,omitempty"`

	// tariff3 name
	Tariff3Name string `json:"Tariff3Name,omitempty"`

	// tariff number
	TariffNumber string `json:"TariffNumber,omitempty"`

	// tariff number int
	TariffNumberInt uint64 `json:"TariffNumberInt,omitempty"`

	// unique num
	UniqueNum string `json:"UniqueNum,omitempty"`

	// units
	Units string `json:"Units,omitempty"`

	// values
	Values []*MeterValuesItems0 `json:"Values"`

	// values can add
	ValuesCanAdd bool `json:"ValuesCanAdd,omitempty"`

	// values end day
	ValuesEndDay uint64 `json:"ValuesEndDay,omitempty"`

	// values period end is current
	ValuesPeriodEndIsCurrent bool `json:"ValuesPeriodEndIsCurrent,omitempty"`

	// values period start is current
	ValuesPeriodStartIsCurrent bool `json:"ValuesPeriodStartIsCurrent,omitempty"`

	// values start day
	ValuesStartDay uint64 `json:"ValuesStartDay,omitempty"`

	// ident
	Ident string `json:"ident,omitempty"`
}

// Validate validates this meter
func (m *Meter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Meter) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Meter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Meter) UnmarshalBinary(b []byte) error {
	var res Meter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MeterValuesItems0 meter values items0
//
// swagger:model MeterValuesItems0
type MeterValuesItems0 struct {

	// is current period
	IsCurrentPeriod bool `json:"IsCurrentPeriod,omitempty"`

	// period
	Period string `json:"Period,omitempty"`

	// value
	Value float64 `json:"Value,omitempty"`

	// value t2
	ValueT2 float64 `json:"ValueT2,omitempty"`

	// value t3
	ValueT3 float64 `json:"ValueT3,omitempty"`
}

// Validate validates this meter values items0
func (m *MeterValuesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MeterValuesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MeterValuesItems0) UnmarshalBinary(b []byte) error {
	var res MeterValuesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
