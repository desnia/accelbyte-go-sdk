// Code generated by go-swagger; DO NOT EDIT.

package groupclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsAssignRoleToMemberRequestV1 models assign role to member request v1
//
// swagger:model models.AssignRoleToMemberRequestV1
type ModelsAssignRoleToMemberRequestV1 struct {

	// user Id
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this models assign role to member request v1
func (m *ModelsAssignRoleToMemberRequestV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsAssignRoleToMemberRequestV1) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsAssignRoleToMemberRequestV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsAssignRoleToMemberRequestV1) UnmarshalBinary(b []byte) error {
	var res ModelsAssignRoleToMemberRequestV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}