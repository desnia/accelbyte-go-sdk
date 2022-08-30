// Code generated by go-swagger; DO NOT EDIT.

package iamclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelSendVerificationLinkRequest model send verification link request
//
// swagger:model model.SendVerificationLinkRequest
type ModelSendVerificationLinkRequest struct {

	// language tag
	LanguageTag string `json:"languageTag,omitempty"`
}

// Validate validates this model send verification link request
func (m *ModelSendVerificationLinkRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelSendVerificationLinkRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelSendVerificationLinkRequest) UnmarshalBinary(b []byte) error {
	var res ModelSendVerificationLinkRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}