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

// OauthmodelErrorResponse oauthmodel error response
//
// swagger:model oauthmodel.ErrorResponse
type OauthmodelErrorResponse struct {

	// client Id
	ClientID string `json:"clientId,omitempty"`

	// default factor
	DefaultFactor string `json:"default_factor,omitempty"`

	// error
	// Required: true
	Error *string `json:"error"`

	// error description
	ErrorDescription string `json:"error_description,omitempty"`

	// error uri
	ErrorURI string `json:"error_uri,omitempty"`

	// factors
	Factors []string `json:"factors"`

	// linking token
	LinkingToken string `json:"linkingToken,omitempty"`

	// message variables
	MessageVariables map[string]string `json:"messageVariables,omitempty"`

	// mfa token
	MfaToken string `json:"mfa_token,omitempty"`

	// platform Id
	PlatformID string `json:"platformId,omitempty"`
}

// Validate validates this oauthmodel error response
func (m *OauthmodelErrorResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OauthmodelErrorResponse) validateError(formats strfmt.Registry) error {

	if err := validate.Required("error", "body", m.Error); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OauthmodelErrorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OauthmodelErrorResponse) UnmarshalBinary(b []byte) error {
	var res OauthmodelErrorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
