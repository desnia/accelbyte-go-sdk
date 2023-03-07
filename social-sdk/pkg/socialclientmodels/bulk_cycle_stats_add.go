// Code generated by go-swagger; DO NOT EDIT.

package socialclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BulkCycleStatsAdd bulk cycle stats add
//
// swagger:model BulkCycleStatsAdd
type BulkCycleStatsAdd struct {

	// stat codes
	// Required: true
	StatCodes []string `json:"statCodes"`
}

// Validate validates this bulk cycle stats add
func (m *BulkCycleStatsAdd) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatCodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkCycleStatsAdd) validateStatCodes(formats strfmt.Registry) error {

	if err := validate.Required("statCodes", "body", m.StatCodes); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkCycleStatsAdd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkCycleStatsAdd) UnmarshalBinary(b []byte) error {
	var res BulkCycleStatsAdd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}