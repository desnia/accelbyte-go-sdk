// Code generated by go-swagger; DO NOT EDIT.

package legalclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RetrievePolicyTypeResponse retrieve policy type response
//
// swagger:model RetrievePolicyTypeResponse
type RetrievePolicyTypeResponse struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// is need document
	// Required: true
	IsNeedDocument *bool `json:"isNeedDocument"`

	// policy type name
	// Required: true
	PolicyTypeName *string `json:"policyTypeName"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this retrieve policy type response
func (m *RetrievePolicyTypeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsNeedDocument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyTypeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RetrievePolicyTypeResponse) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RetrievePolicyTypeResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RetrievePolicyTypeResponse) validateIsNeedDocument(formats strfmt.Registry) error {

	if err := validate.Required("isNeedDocument", "body", m.IsNeedDocument); err != nil {
		return err
	}

	return nil
}

func (m *RetrievePolicyTypeResponse) validatePolicyTypeName(formats strfmt.Registry) error {

	if err := validate.Required("policyTypeName", "body", m.PolicyTypeName); err != nil {
		return err
	}

	return nil
}

func (m *RetrievePolicyTypeResponse) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RetrievePolicyTypeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetrievePolicyTypeResponse) UnmarshalBinary(b []byte) error {
	var res RetrievePolicyTypeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}