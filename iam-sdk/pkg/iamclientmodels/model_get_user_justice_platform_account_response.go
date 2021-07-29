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

// ModelGetUserJusticePlatformAccountResponse model get user justice platform account response
//
// swagger:model model.GetUserJusticePlatformAccountResponse
type ModelGetUserJusticePlatformAccountResponse struct {

	// designated namespace
	// Required: true
	DesignatedNamespace *string `json:"DesignatedNamespace"`

	// user ID
	// Required: true
	UserID *string `json:"UserID"`
}

// Validate validates this model get user justice platform account response
func (m *ModelGetUserJusticePlatformAccountResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDesignatedNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelGetUserJusticePlatformAccountResponse) validateDesignatedNamespace(formats strfmt.Registry) error {

	if err := validate.Required("DesignatedNamespace", "body", m.DesignatedNamespace); err != nil {
		return err
	}

	return nil
}

func (m *ModelGetUserJusticePlatformAccountResponse) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("UserID", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelGetUserJusticePlatformAccountResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelGetUserJusticePlatformAccountResponse) UnmarshalBinary(b []byte) error {
	var res ModelGetUserJusticePlatformAccountResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}