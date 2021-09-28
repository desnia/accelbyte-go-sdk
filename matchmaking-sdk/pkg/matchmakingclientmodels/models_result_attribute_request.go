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

// ModelsResultAttributeRequest models result attribute request
//
// swagger:model models.ResultAttributeRequest
type ModelsResultAttributeRequest struct {

	// attribute
	// Required: true
	Attribute string `json:"attribute"`

	// value
	// Required: true
	Value float64 `json:"value"`
}

// Validate validates this models result attribute request
func (m *ModelsResultAttributeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttribute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsResultAttributeRequest) validateAttribute(formats strfmt.Registry) error {

	if err := validate.RequiredString("attribute", "body", string(m.Attribute)); err != nil {
		return err
	}

	return nil
}

func (m *ModelsResultAttributeRequest) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", float64(m.Value)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsResultAttributeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsResultAttributeRequest) UnmarshalBinary(b []byte) error {
	var res ModelsResultAttributeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}