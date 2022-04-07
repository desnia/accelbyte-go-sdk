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

// ModelBackupCodesResponseV4 model backup codes response v4
//
// swagger:model model.BackupCodesResponseV4
type ModelBackupCodesResponseV4 struct {

	// generated at
	// Required: true
	GeneratedAt *int64 `json:"generatedAt"`

	// invalid codes
	// Required: true
	InvalidCodes []string `json:"invalidCodes"`

	// valid codes
	// Required: true
	ValidCodes []string `json:"validCodes"`
}

// Validate validates this model backup codes response v4
func (m *ModelBackupCodesResponseV4) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeneratedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvalidCodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidCodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelBackupCodesResponseV4) validateGeneratedAt(formats strfmt.Registry) error {

	if err := validate.Required("generatedAt", "body", m.GeneratedAt); err != nil {
		return err
	}

	return nil
}

func (m *ModelBackupCodesResponseV4) validateInvalidCodes(formats strfmt.Registry) error {

	if err := validate.Required("invalidCodes", "body", m.InvalidCodes); err != nil {
		return err
	}

	return nil
}

func (m *ModelBackupCodesResponseV4) validateValidCodes(formats strfmt.Registry) error {

	if err := validate.Required("validCodes", "body", m.ValidCodes); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelBackupCodesResponseV4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelBackupCodesResponseV4) UnmarshalBinary(b []byte) error {
	var res ModelBackupCodesResponseV4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}