// Code generated by go-swagger; DO NOT EDIT.

package matchmakingclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsChannelRequest models channel request
//
// swagger:model models.ChannelRequest
type ModelsChannelRequest struct {

	// deployment
	// Required: true
	Deployment string `json:"deployment"`

	// description
	// Required: true
	Description string `json:"description"`

	// find match timeout seconds
	// Required: true
	FindMatchTimeoutSeconds int32 `json:"find_match_timeout_seconds"`

	// game mode
	// Required: true
	GameMode string `json:"game_mode"`

	// joinable
	Joinable *bool `json:"joinable,omitempty"`

	// max delay ms
	// Required: true
	MaxDelayMs int32 `json:"max_delay_ms"`

	// rule set
	// Required: true
	RuleSet *ModelsRuleSet `json:"rule_set"`

	// session queue timeout seconds
	// Required: true
	SessionQueueTimeoutSeconds int32 `json:"session_queue_timeout_seconds"`

	// social matchmaking
	SocialMatchmaking *bool `json:"social_matchmaking,omitempty"`

	// use sub gamemode
	UseSubGamemode *bool `json:"use_sub_gamemode,omitempty"`
}

// Validate validates this models channel request
func (m *ModelsChannelRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeployment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFindMatchTimeoutSeconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGameMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxDelayMs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionQueueTimeoutSeconds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsChannelRequest) validateDeployment(formats strfmt.Registry) error {

	if err := validate.RequiredString("deployment", "body", string(m.Deployment)); err != nil {
		return err
	}

	return nil
}

func (m *ModelsChannelRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.RequiredString("description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *ModelsChannelRequest) validateFindMatchTimeoutSeconds(formats strfmt.Registry) error {

	if err := validate.Required("find_match_timeout_seconds", "body", int32(m.FindMatchTimeoutSeconds)); err != nil {
		return err
	}

	return nil
}

func (m *ModelsChannelRequest) validateGameMode(formats strfmt.Registry) error {

	if err := validate.RequiredString("game_mode", "body", string(m.GameMode)); err != nil {
		return err
	}

	return nil
}

func (m *ModelsChannelRequest) validateMaxDelayMs(formats strfmt.Registry) error {

	if err := validate.Required("max_delay_ms", "body", int32(m.MaxDelayMs)); err != nil {
		return err
	}

	return nil
}

func (m *ModelsChannelRequest) validateRuleSet(formats strfmt.Registry) error {

	if err := validate.Required("rule_set", "body", m.RuleSet); err != nil {
		return err
	}

	if m.RuleSet != nil {
		if err := m.RuleSet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rule_set")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsChannelRequest) validateSessionQueueTimeoutSeconds(formats strfmt.Registry) error {

	if err := validate.Required("session_queue_timeout_seconds", "body", int32(m.SessionQueueTimeoutSeconds)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsChannelRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsChannelRequest) UnmarshalBinary(b []byte) error {
	var res ModelsChannelRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}