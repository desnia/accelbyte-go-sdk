// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ADTOObjectForOrderCreationOptions a DTO object for order creation options
//
// swagger:model A DTO object for order creation options
type ADTOObjectForOrderCreationOptions struct {

	// skip price validation
	SkipPriceValidation bool `json:"skipPriceValidation"`
}

// Validate validates this a DTO object for order creation options
func (m *ADTOObjectForOrderCreationOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ADTOObjectForOrderCreationOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ADTOObjectForOrderCreationOptions) UnmarshalBinary(b []byte) error {
	var res ADTOObjectForOrderCreationOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}