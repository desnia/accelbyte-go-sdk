// Code generated by go-swagger; DO NOT EDIT.

package sessionbrowserclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsCountActiveSessionResponse models count active session response
//
// swagger:model models.CountActiveSessionResponse
type ModelsCountActiveSessionResponse struct {

	// custom game
	// Required: true
	CustomGame *int64 `json:"custom_game"`

	// matchmaking game
	// Required: true
	MatchmakingGame *int64 `json:"matchmaking_game"`

	// total
	// Required: true
	Total *int64 `json:"total"`
}

// Validate validates this models count active session response
func (m *ModelsCountActiveSessionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomGame(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatchmakingGame(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsCountActiveSessionResponse) validateCustomGame(formats strfmt.Registry) error {

	if err := validate.Required("custom_game", "body", m.CustomGame); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCountActiveSessionResponse) validateMatchmakingGame(formats strfmt.Registry) error {

	if err := validate.Required("matchmaking_game", "body", m.MatchmakingGame); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCountActiveSessionResponse) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsCountActiveSessionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsCountActiveSessionResponse) UnmarshalBinary(b []byte) error {
	var res ModelsCountActiveSessionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}