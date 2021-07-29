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

// LegalAcceptedPoliciesRequest legal accepted policies request
//
// swagger:model legal.AcceptedPoliciesRequest
type LegalAcceptedPoliciesRequest struct {

	// is accepted
	// Required: true
	IsAccepted *bool `json:"isAccepted"`

	// localized policy version Id
	// Required: true
	LocalizedPolicyVersionID *string `json:"localizedPolicyVersionId"`

	// policy Id
	// Required: true
	PolicyID *string `json:"policyId"`

	// policy version Id
	// Required: true
	PolicyVersionID *string `json:"policyVersionId"`
}

// Validate validates this legal accepted policies request
func (m *LegalAcceptedPoliciesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsAccepted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalizedPolicyVersionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LegalAcceptedPoliciesRequest) validateIsAccepted(formats strfmt.Registry) error {

	if err := validate.Required("isAccepted", "body", m.IsAccepted); err != nil {
		return err
	}

	return nil
}

func (m *LegalAcceptedPoliciesRequest) validateLocalizedPolicyVersionID(formats strfmt.Registry) error {

	if err := validate.Required("localizedPolicyVersionId", "body", m.LocalizedPolicyVersionID); err != nil {
		return err
	}

	return nil
}

func (m *LegalAcceptedPoliciesRequest) validatePolicyID(formats strfmt.Registry) error {

	if err := validate.Required("policyId", "body", m.PolicyID); err != nil {
		return err
	}

	return nil
}

func (m *LegalAcceptedPoliciesRequest) validatePolicyVersionID(formats strfmt.Registry) error {

	if err := validate.Required("policyVersionId", "body", m.PolicyVersionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LegalAcceptedPoliciesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LegalAcceptedPoliciesRequest) UnmarshalBinary(b []byte) error {
	var res LegalAcceptedPoliciesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}