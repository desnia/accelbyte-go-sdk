// Code generated by go-swagger; DO NOT EDIT.

package iamclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelAccountProgressionInfo model account progression info
//
// swagger:model model.AccountProgressionInfo
type ModelAccountProgressionInfo struct {

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// For headlessAccount: list of all namespaces from this headless account. For currentAccount: list of all namespaces that conflict with headlessAccount
	// Required: true
	LinkedGames []string `json:"linkedGames"`

	// user name
	UserName string `json:"userName,omitempty"`
}

// Validate validates this model account progression info
func (m *ModelAccountProgressionInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinkedGames(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelAccountProgressionInfo) validateLinkedGames(formats strfmt.Registry) error {

	if err := validate.Required("linkedGames", "body", m.LinkedGames); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelAccountProgressionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelAccountProgressionInfo) UnmarshalBinary(b []byte) error {
	var res ModelAccountProgressionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
