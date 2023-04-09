// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsumableEntitlementRevocationConfig consumable entitlement revocation config
//
// swagger:model ConsumableEntitlementRevocationConfig
type ConsumableEntitlementRevocationConfig struct {

	// will do revocation if enabled, otherwise will skip do revocation.
	Enabled bool `json:"enabled"`

	// consumable entitlement revocation strategy
	// Enum: [CUSTOM REVOKE_OR_REPORT]
	Strategy string `json:"strategy,omitempty"`
}

// Validate validates this consumable entitlement revocation config
func (m *ConsumableEntitlementRevocationConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var consumableEntitlementRevocationConfigTypeStrategyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CUSTOM","REVOKE_OR_REPORT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consumableEntitlementRevocationConfigTypeStrategyPropEnum = append(consumableEntitlementRevocationConfigTypeStrategyPropEnum, v)
	}
}

const (

	// ConsumableEntitlementRevocationConfigStrategyCUSTOM captures enum value "CUSTOM"
	ConsumableEntitlementRevocationConfigStrategyCUSTOM string = "CUSTOM"

	// ConsumableEntitlementRevocationConfigStrategyREVOKEORREPORT captures enum value "REVOKE_OR_REPORT"
	ConsumableEntitlementRevocationConfigStrategyREVOKEORREPORT string = "REVOKE_OR_REPORT"
)

// prop value enum
func (m *ConsumableEntitlementRevocationConfig) validateStrategyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consumableEntitlementRevocationConfigTypeStrategyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsumableEntitlementRevocationConfig) validateStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.Strategy) { // not required
		return nil
	}

	// value enum
	if err := m.validateStrategyEnum("strategy", "body", m.Strategy); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsumableEntitlementRevocationConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsumableEntitlementRevocationConfig) UnmarshalBinary(b []byte) error {
	var res ConsumableEntitlementRevocationConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
