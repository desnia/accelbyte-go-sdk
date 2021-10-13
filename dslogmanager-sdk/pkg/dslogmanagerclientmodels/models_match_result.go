// Code generated by go-swagger; DO NOT EDIT.

package dslogmanagerclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsMatchResult models match result
//
// swagger:model models.MatchResult
type ModelsMatchResult struct {

	// game mode
	// Required: true
	GameMode *string `json:"game_mode"`

	// matching allies
	// Required: true
	MatchingAllies []*ModelsRequestMatchingAlly `json:"matching_allies"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// session id
	// Required: true
	SessionID *string `json:"session_id"`
}

// Validate validates this models match result
func (m *ModelsMatchResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGameMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatchingAllies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsMatchResult) validateGameMode(formats strfmt.Registry) error {

	if err := validate.Required("game_mode", "body", m.GameMode); err != nil {
		return err
	}

	return nil
}

func (m *ModelsMatchResult) validateMatchingAllies(formats strfmt.Registry) error {

	if err := validate.Required("matching_allies", "body", m.MatchingAllies); err != nil {
		return err
	}

	for i := 0; i < len(m.MatchingAllies); i++ {
		if swag.IsZero(m.MatchingAllies[i]) { // not required
			continue
		}

		if m.MatchingAllies[i] != nil {
			if err := m.MatchingAllies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matching_allies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsMatchResult) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *ModelsMatchResult) validateSessionID(formats strfmt.Registry) error {

	if err := validate.Required("session_id", "body", m.SessionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsMatchResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsMatchResult) UnmarshalBinary(b []byte) error {
	var res ModelsMatchResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}