// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	static "github.com/kihamo/boggart/providers/openweathermap/static/models"
)

// OneCall one call
//
// swagger:model OneCall
type OneCall struct {

	// current
	Current *OneCallCurrent `json:"current,omitempty"`

	// daily
	Daily []*OneCallDailyItems0 `json:"daily"`

	// hourly
	Hourly []*OneCallHourlyItems0 `json:"hourly"`

	// lat
	Lat float64 `json:"lat,omitempty"`

	// lon
	Lon float64 `json:"lon,omitempty"`

	// minutely
	Minutely []*OneCallMinutelyItems0 `json:"minutely"`

	// timezone
	Timezone string `json:"timezone,omitempty"`

	// timezone offset
	TimezoneOffset uint64 `json:"timezone_offset,omitempty"`
}

// Validate validates this one call
func (m *OneCall) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDaily(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHourly(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinutely(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OneCall) validateCurrent(formats strfmt.Registry) error {

	if swag.IsZero(m.Current) { // not required
		return nil
	}

	if m.Current != nil {
		if err := m.Current.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current")
			}
			return err
		}
	}

	return nil
}

func (m *OneCall) validateDaily(formats strfmt.Registry) error {

	if swag.IsZero(m.Daily) { // not required
		return nil
	}

	for i := 0; i < len(m.Daily); i++ {
		if swag.IsZero(m.Daily[i]) { // not required
			continue
		}

		if m.Daily[i] != nil {
			if err := m.Daily[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("daily" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OneCall) validateHourly(formats strfmt.Registry) error {

	if swag.IsZero(m.Hourly) { // not required
		return nil
	}

	for i := 0; i < len(m.Hourly); i++ {
		if swag.IsZero(m.Hourly[i]) { // not required
			continue
		}

		if m.Hourly[i] != nil {
			if err := m.Hourly[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hourly" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OneCall) validateMinutely(formats strfmt.Registry) error {

	if swag.IsZero(m.Minutely) { // not required
		return nil
	}

	for i := 0; i < len(m.Minutely); i++ {
		if swag.IsZero(m.Minutely[i]) { // not required
			continue
		}

		if m.Minutely[i] != nil {
			if err := m.Minutely[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("minutely" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OneCall) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OneCall) UnmarshalBinary(b []byte) error {
	var res OneCall
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OneCallCurrent one call current
//
// swagger:model OneCallCurrent
type OneCallCurrent struct {

	// clouds
	Clouds uint64 `json:"clouds,omitempty"`

	// dew point
	DewPoint float64 `json:"dew_point,omitempty"`

	// dt
	// Format: date-time
	Dt static.DateTime `json:"dt,omitempty"`

	// feels like
	FeelsLike float64 `json:"feels_like,omitempty"`

	// humidity
	Humidity uint64 `json:"humidity,omitempty"`

	// pop
	Pop float64 `json:"pop,omitempty"`

	// pressure
	Pressure float64 `json:"pressure,omitempty"`

	// rain
	Rain *Rain `json:"rain,omitempty"`

	// snow
	Snow *Snow `json:"snow,omitempty"`

	// sunrise
	// Format: date-time
	Sunrise static.DateTime `json:"sunrise,omitempty"`

	// sunset
	// Format: date-time
	Sunset static.DateTime `json:"sunset,omitempty"`

	// temp
	Temp float64 `json:"temp,omitempty"`

	// uvi
	Uvi float64 `json:"uvi,omitempty"`

	// visibility
	Visibility uint64 `json:"visibility,omitempty"`

	// weather
	Weather []*Weather `json:"weather"`

	// wind gust
	WindGust float64 `json:"wind_gust,omitempty"`

	// wind speed
	WindSpeed float64 `json:"wind_speed,omitempty"`
}

// Validate validates this one call current
func (m *OneCallCurrent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSunrise(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSunset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeather(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OneCallCurrent) validateDt(formats strfmt.Registry) error {

	if swag.IsZero(m.Dt) { // not required
		return nil
	}

	if err := m.Dt.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("current" + "." + "dt")
		}
		return err
	}

	return nil
}

func (m *OneCallCurrent) validateRain(formats strfmt.Registry) error {

	if swag.IsZero(m.Rain) { // not required
		return nil
	}

	if m.Rain != nil {
		if err := m.Rain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current" + "." + "rain")
			}
			return err
		}
	}

	return nil
}

func (m *OneCallCurrent) validateSnow(formats strfmt.Registry) error {

	if swag.IsZero(m.Snow) { // not required
		return nil
	}

	if m.Snow != nil {
		if err := m.Snow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current" + "." + "snow")
			}
			return err
		}
	}

	return nil
}

func (m *OneCallCurrent) validateSunrise(formats strfmt.Registry) error {

	if swag.IsZero(m.Sunrise) { // not required
		return nil
	}

	if err := m.Sunrise.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("current" + "." + "sunrise")
		}
		return err
	}

	return nil
}

func (m *OneCallCurrent) validateSunset(formats strfmt.Registry) error {

	if swag.IsZero(m.Sunset) { // not required
		return nil
	}

	if err := m.Sunset.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("current" + "." + "sunset")
		}
		return err
	}

	return nil
}

func (m *OneCallCurrent) validateWeather(formats strfmt.Registry) error {

	if swag.IsZero(m.Weather) { // not required
		return nil
	}

	for i := 0; i < len(m.Weather); i++ {
		if swag.IsZero(m.Weather[i]) { // not required
			continue
		}

		if m.Weather[i] != nil {
			if err := m.Weather[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("current" + "." + "weather" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OneCallCurrent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OneCallCurrent) UnmarshalBinary(b []byte) error {
	var res OneCallCurrent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OneCallDailyItems0 one call daily items0
//
// swagger:model OneCallDailyItems0
type OneCallDailyItems0 struct {

	// clouds
	Clouds uint64 `json:"clouds,omitempty"`

	// dew point
	DewPoint float64 `json:"dew_point,omitempty"`

	// dt
	// Format: date-time
	Dt static.DateTime `json:"dt,omitempty"`

	// feels like
	FeelsLike *OneCallDailyItems0FeelsLike `json:"feels_like,omitempty"`

	// humidity
	Humidity uint64 `json:"humidity,omitempty"`

	// pop
	Pop float64 `json:"pop,omitempty"`

	// pressure
	Pressure float64 `json:"pressure,omitempty"`

	// rain
	Rain float64 `json:"rain,omitempty"`

	// snow
	Snow float64 `json:"snow,omitempty"`

	// sunrise
	// Format: date-time
	Sunrise static.DateTime `json:"sunrise,omitempty"`

	// sunset
	// Format: date-time
	Sunset static.DateTime `json:"sunset,omitempty"`

	// temp
	Temp *OneCallDailyItems0Temp `json:"temp,omitempty"`

	// uvi
	Uvi float64 `json:"uvi,omitempty"`

	// visibility
	Visibility uint64 `json:"visibility,omitempty"`

	// weather
	Weather []*Weather `json:"weather"`

	// wind deg
	WindDeg uint64 `json:"wind_deg,omitempty"`

	// wind gust
	WindGust float64 `json:"wind_gust,omitempty"`

	// wind speed
	WindSpeed float64 `json:"wind_speed,omitempty"`
}

// Validate validates this one call daily items0
func (m *OneCallDailyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeelsLike(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSunrise(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSunset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeather(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OneCallDailyItems0) validateDt(formats strfmt.Registry) error {

	if swag.IsZero(m.Dt) { // not required
		return nil
	}

	if err := m.Dt.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dt")
		}
		return err
	}

	return nil
}

func (m *OneCallDailyItems0) validateFeelsLike(formats strfmt.Registry) error {

	if swag.IsZero(m.FeelsLike) { // not required
		return nil
	}

	if m.FeelsLike != nil {
		if err := m.FeelsLike.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("feels_like")
			}
			return err
		}
	}

	return nil
}

func (m *OneCallDailyItems0) validateSunrise(formats strfmt.Registry) error {

	if swag.IsZero(m.Sunrise) { // not required
		return nil
	}

	if err := m.Sunrise.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sunrise")
		}
		return err
	}

	return nil
}

func (m *OneCallDailyItems0) validateSunset(formats strfmt.Registry) error {

	if swag.IsZero(m.Sunset) { // not required
		return nil
	}

	if err := m.Sunset.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sunset")
		}
		return err
	}

	return nil
}

func (m *OneCallDailyItems0) validateTemp(formats strfmt.Registry) error {

	if swag.IsZero(m.Temp) { // not required
		return nil
	}

	if m.Temp != nil {
		if err := m.Temp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("temp")
			}
			return err
		}
	}

	return nil
}

func (m *OneCallDailyItems0) validateWeather(formats strfmt.Registry) error {

	if swag.IsZero(m.Weather) { // not required
		return nil
	}

	for i := 0; i < len(m.Weather); i++ {
		if swag.IsZero(m.Weather[i]) { // not required
			continue
		}

		if m.Weather[i] != nil {
			if err := m.Weather[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("weather" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OneCallDailyItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OneCallDailyItems0) UnmarshalBinary(b []byte) error {
	var res OneCallDailyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OneCallDailyItems0FeelsLike one call daily items0 feels like
//
// swagger:model OneCallDailyItems0FeelsLike
type OneCallDailyItems0FeelsLike struct {

	// day
	Day float64 `json:"day,omitempty"`

	// eve
	Eve float64 `json:"eve,omitempty"`

	// morn
	Morn float64 `json:"morn,omitempty"`

	// night
	Night float64 `json:"night,omitempty"`
}

// Validate validates this one call daily items0 feels like
func (m *OneCallDailyItems0FeelsLike) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OneCallDailyItems0FeelsLike) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OneCallDailyItems0FeelsLike) UnmarshalBinary(b []byte) error {
	var res OneCallDailyItems0FeelsLike
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OneCallDailyItems0Temp one call daily items0 temp
//
// swagger:model OneCallDailyItems0Temp
type OneCallDailyItems0Temp struct {

	// day
	Day float64 `json:"day,omitempty"`

	// eve
	Eve float64 `json:"eve,omitempty"`

	// max
	Max float64 `json:"max,omitempty"`

	// min
	Min float64 `json:"min,omitempty"`

	// morn
	Morn float64 `json:"morn,omitempty"`

	// night
	Night float64 `json:"night,omitempty"`
}

// Validate validates this one call daily items0 temp
func (m *OneCallDailyItems0Temp) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OneCallDailyItems0Temp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OneCallDailyItems0Temp) UnmarshalBinary(b []byte) error {
	var res OneCallDailyItems0Temp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OneCallHourlyItems0 one call hourly items0
//
// swagger:model OneCallHourlyItems0
type OneCallHourlyItems0 struct {

	// clouds
	Clouds uint64 `json:"clouds,omitempty"`

	// dew point
	DewPoint float64 `json:"dew_point,omitempty"`

	// dt
	// Format: date-time
	Dt static.DateTime `json:"dt,omitempty"`

	// feels like
	FeelsLike float64 `json:"feels_like,omitempty"`

	// humidity
	Humidity uint64 `json:"humidity,omitempty"`

	// pressure
	Pressure float64 `json:"pressure,omitempty"`

	// rain
	Rain *Rain `json:"rain,omitempty"`

	// snow
	Snow *Snow `json:"snow,omitempty"`

	// temp
	Temp float64 `json:"temp,omitempty"`

	// visibility
	Visibility uint64 `json:"visibility,omitempty"`

	// weather
	Weather []*Weather `json:"weather"`

	// wind gust
	WindGust float64 `json:"wind_gust,omitempty"`

	// wind speed
	WindSpeed float64 `json:"wind_speed,omitempty"`
}

// Validate validates this one call hourly items0
func (m *OneCallHourlyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeather(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OneCallHourlyItems0) validateDt(formats strfmt.Registry) error {

	if swag.IsZero(m.Dt) { // not required
		return nil
	}

	if err := m.Dt.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dt")
		}
		return err
	}

	return nil
}

func (m *OneCallHourlyItems0) validateRain(formats strfmt.Registry) error {

	if swag.IsZero(m.Rain) { // not required
		return nil
	}

	if m.Rain != nil {
		if err := m.Rain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rain")
			}
			return err
		}
	}

	return nil
}

func (m *OneCallHourlyItems0) validateSnow(formats strfmt.Registry) error {

	if swag.IsZero(m.Snow) { // not required
		return nil
	}

	if m.Snow != nil {
		if err := m.Snow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snow")
			}
			return err
		}
	}

	return nil
}

func (m *OneCallHourlyItems0) validateWeather(formats strfmt.Registry) error {

	if swag.IsZero(m.Weather) { // not required
		return nil
	}

	for i := 0; i < len(m.Weather); i++ {
		if swag.IsZero(m.Weather[i]) { // not required
			continue
		}

		if m.Weather[i] != nil {
			if err := m.Weather[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("weather" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OneCallHourlyItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OneCallHourlyItems0) UnmarshalBinary(b []byte) error {
	var res OneCallHourlyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OneCallMinutelyItems0 one call minutely items0
//
// swagger:model OneCallMinutelyItems0
type OneCallMinutelyItems0 struct {

	// dt
	// Format: date-time
	Dt static.DateTime `json:"dt,omitempty"`

	// precipitation
	Precipitation float64 `json:"precipitation,omitempty"`
}

// Validate validates this one call minutely items0
func (m *OneCallMinutelyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OneCallMinutelyItems0) validateDt(formats strfmt.Registry) error {

	if swag.IsZero(m.Dt) { // not required
		return nil
	}

	if err := m.Dt.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dt")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OneCallMinutelyItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OneCallMinutelyItems0) UnmarshalBinary(b []byte) error {
	var res OneCallMinutelyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
