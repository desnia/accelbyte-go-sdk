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

// ModelAuthenticatorKeyResponseV4 model authenticator key response v4
//
// swagger:model model.AuthenticatorKeyResponseV4
type ModelAuthenticatorKeyResponseV4 struct {

	// secret key
	// Required: true
	SecretKey *string `json:"secretKey"`

	// uri
	// Required: true
	URI *string `json:"uri"`
}

// Validate validates this model authenticator key response v4
func (m *ModelAuthenticatorKeyResponseV4) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelAuthenticatorKeyResponseV4) validateSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("secretKey", "body", m.SecretKey); err != nil {
		return err
	}

	return nil
}

func (m *ModelAuthenticatorKeyResponseV4) validateURI(formats strfmt.Registry) error {

	if err := validate.Required("uri", "body", m.URI); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelAuthenticatorKeyResponseV4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelAuthenticatorKeyResponseV4) UnmarshalBinary(b []byte) error {
	var res ModelAuthenticatorKeyResponseV4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}