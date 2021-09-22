// Code generated by go-swagger; DO NOT EDIT.

package achievementclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsAchievementOrderUpdateRequest models achievement order update request
//
// swagger:model models.AchievementOrderUpdateRequest
type ModelsAchievementOrderUpdateRequest struct {

	// target order
	// Required: true
	TargetOrder *int32 `json:"targetOrder"`
}

// Validate validates this models achievement order update request
func (m *ModelsAchievementOrderUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargetOrder(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsAchievementOrderUpdateRequest) validateTargetOrder(formats strfmt.Registry) error {

	if err := validate.Required("targetOrder", "body", m.TargetOrder); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsAchievementOrderUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsAchievementOrderUpdateRequest) UnmarshalBinary(b []byte) error {
	var res ModelsAchievementOrderUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
