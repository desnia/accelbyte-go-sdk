// Code generated by go-swagger; DO NOT EDIT.

package lobbyclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsGetLobbyCcuResponse models get lobby ccu response
//
// swagger:model models.GetLobbyCcuResponse
type ModelsGetLobbyCcuResponse struct {

	// count current users
	// Required: true
	CountCurrentUsers *int32 `json:"countCurrentUsers"`
}

// Validate validates this models get lobby ccu response
func (m *ModelsGetLobbyCcuResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountCurrentUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsGetLobbyCcuResponse) validateCountCurrentUsers(formats strfmt.Registry) error {

	if err := validate.Required("countCurrentUsers", "body", m.CountCurrentUsers); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsGetLobbyCcuResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsGetLobbyCcuResponse) UnmarshalBinary(b []byte) error {
	var res ModelsGetLobbyCcuResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}