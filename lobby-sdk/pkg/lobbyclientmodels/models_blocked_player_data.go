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

// ModelsBlockedPlayerData models blocked player data
//
// swagger:model models.BlockedPlayerData
type ModelsBlockedPlayerData struct {

	// blocked at
	// Required: true
	// Format: date-time
	BlockedAt *strfmt.DateTime `json:"blockedAt"`

	// blocked user Id
	// Required: true
	BlockedUserID *string `json:"blockedUserId"`
}

// Validate validates this models blocked player data
func (m *ModelsBlockedPlayerData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlockedUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsBlockedPlayerData) validateBlockedAt(formats strfmt.Registry) error {

	if err := validate.Required("blockedAt", "body", m.BlockedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("blockedAt", "body", "date-time", m.BlockedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelsBlockedPlayerData) validateBlockedUserID(formats strfmt.Registry) error {

	if err := validate.Required("blockedUserId", "body", m.BlockedUserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsBlockedPlayerData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsBlockedPlayerData) UnmarshalBinary(b []byte) error {
	var res ModelsBlockedPlayerData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
