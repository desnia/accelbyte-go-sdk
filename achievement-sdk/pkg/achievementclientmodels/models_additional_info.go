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

// ModelsAdditionalInfo models additional info
//
// swagger:model models.AdditionalInfo
type ModelsAdditionalInfo struct {

	// number of achievements
	// Required: true
	NumberOfAchievements *int64 `json:"numberOfAchievements"`

	// number of hidden achievements
	// Required: true
	NumberOfHiddenAchievements *int64 `json:"numberOfHiddenAchievements"`

	// number of visible achievements
	// Required: true
	NumberOfVisibleAchievements *int64 `json:"numberOfVisibleAchievements"`
}

// Validate validates this models additional info
func (m *ModelsAdditionalInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNumberOfAchievements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberOfHiddenAchievements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberOfVisibleAchievements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsAdditionalInfo) validateNumberOfAchievements(formats strfmt.Registry) error {

	if err := validate.Required("numberOfAchievements", "body", m.NumberOfAchievements); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAdditionalInfo) validateNumberOfHiddenAchievements(formats strfmt.Registry) error {

	if err := validate.Required("numberOfHiddenAchievements", "body", m.NumberOfHiddenAchievements); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAdditionalInfo) validateNumberOfVisibleAchievements(formats strfmt.Registry) error {

	if err := validate.Required("numberOfVisibleAchievements", "body", m.NumberOfVisibleAchievements); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsAdditionalInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsAdditionalInfo) UnmarshalBinary(b []byte) error {
	var res ModelsAdditionalInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
