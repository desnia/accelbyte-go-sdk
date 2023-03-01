// Code generated by go-swagger; DO NOT EDIT.

package socialclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StatCycleUpdate stat cycle update
//
// swagger:model StatCycleUpdate
type StatCycleUpdate struct {

	// Cycle Type
	// Required: true
	// Enum: [ANNUALLY DAILY MONTHLY SEASONAL WEEKLY]
	CycleType *string `json:"cycleType"`

	// description
	Description string `json:"description,omitempty"`

	// end date time
	// Format: date-time
	End *strfmt.DateTime `json:"end,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// reset date
	ResetDate int32 `json:"resetDate,omitempty"`

	// reset day of week
	ResetDay int32 `json:"resetDay,omitempty"`

	// reset month
	ResetMonth int32 `json:"resetMonth,omitempty"`

	// reset time
	// Required: true
	ResetTime *string `json:"resetTime"`

	// reset month
	SeasonPeriod int32 `json:"seasonPeriod,omitempty"`

	// start date time
	// Required: true
	// Format: date-time
	Start strfmt.DateTime `json:"start"`
}

// Validate validates this stat cycle update
func (m *StatCycleUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCycleType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResetTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var statCycleUpdateTypeCycleTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ANNUALLY","DAILY","MONTHLY","SEASONAL","WEEKLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statCycleUpdateTypeCycleTypePropEnum = append(statCycleUpdateTypeCycleTypePropEnum, v)
	}
}

const (

	// StatCycleUpdateCycleTypeANNUALLY captures enum value "ANNUALLY"
	StatCycleUpdateCycleTypeANNUALLY string = "ANNUALLY"

	// StatCycleUpdateCycleTypeDAILY captures enum value "DAILY"
	StatCycleUpdateCycleTypeDAILY string = "DAILY"

	// StatCycleUpdateCycleTypeMONTHLY captures enum value "MONTHLY"
	StatCycleUpdateCycleTypeMONTHLY string = "MONTHLY"

	// StatCycleUpdateCycleTypeSEASONAL captures enum value "SEASONAL"
	StatCycleUpdateCycleTypeSEASONAL string = "SEASONAL"

	// StatCycleUpdateCycleTypeWEEKLY captures enum value "WEEKLY"
	StatCycleUpdateCycleTypeWEEKLY string = "WEEKLY"
)

// prop value enum
func (m *StatCycleUpdate) validateCycleTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statCycleUpdateTypeCycleTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StatCycleUpdate) validateCycleType(formats strfmt.Registry) error {

	if err := validate.Required("cycleType", "body", m.CycleType); err != nil {
		return err
	}

	// value enum
	if err := m.validateCycleTypeEnum("cycleType", "body", *m.CycleType); err != nil {
		return err
	}

	return nil
}

func (m *StatCycleUpdate) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StatCycleUpdate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StatCycleUpdate) validateResetTime(formats strfmt.Registry) error {

	if err := validate.Required("resetTime", "body", m.ResetTime); err != nil {
		return err
	}

	return nil
}

func (m *StatCycleUpdate) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", strfmt.DateTime(m.Start)); err != nil {
		return err
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatCycleUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatCycleUpdate) UnmarshalBinary(b []byte) error {
	var res StatCycleUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
