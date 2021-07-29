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

// OauthmodelTokenThirdPartyResponse oauthmodel token third party response
//
// swagger:model oauthmodel.TokenThirdPartyResponse
type OauthmodelTokenThirdPartyResponse struct {

	// platform token
	// Required: true
	PlatformToken *string `json:"platform_token"`

	// sand box id
	SandBoxID string `json:"sand_box_id,omitempty"`
}

// Validate validates this oauthmodel token third party response
func (m *OauthmodelTokenThirdPartyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlatformToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OauthmodelTokenThirdPartyResponse) validatePlatformToken(formats strfmt.Registry) error {

	if err := validate.Required("platform_token", "body", m.PlatformToken); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OauthmodelTokenThirdPartyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OauthmodelTokenThirdPartyResponse) UnmarshalBinary(b []byte) error {
	var res OauthmodelTokenThirdPartyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}