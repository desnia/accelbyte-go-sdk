// Code generated by go-swagger; DO NOT EDIT.

package iamclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OauthmodelTokenWithDeviceCookieResponseV3 oauthmodel token with device cookie response v3
//
// swagger:model oauthmodel.TokenWithDeviceCookieResponseV3
type OauthmodelTokenWithDeviceCookieResponseV3 struct {

	// access token
	// Required: true
	AccessToken *string `json:"access_token"`

	// Authentication Trust Id for device cookie validation. Only exist when login using grant_type=password and no existing Auth-Trust-Id given from request header
	AuthTrustID string `json:"auth_trust_id,omitempty"`

	// bans
	// Required: true
	Bans []*AccountcommonJWTBanV3 `json:"bans"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// expires in
	// Required: true
	ExpiresIn *int32 `json:"expires_in"`

	// is comply
	IsComply bool `json:"is_comply"`

	// jflgs
	Jflgs int32 `json:"jflgs,omitempty"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// namespace roles
	// Required: true
	NamespaceRoles []*AccountcommonNamespaceRole `json:"namespace_roles"`

	// permissions
	// Required: true
	Permissions []*AccountcommonPermissionV3 `json:"permissions"`

	// platform id
	PlatformID string `json:"platform_id,omitempty"`

	// platform user id
	PlatformUserID string `json:"platform_user_id,omitempty"`

	// present if it is user token
	RefreshExpiresIn int32 `json:"refresh_expires_in,omitempty"`

	// present if it is user token
	RefreshToken string `json:"refresh_token,omitempty"`

	// roles
	// Required: true
	Roles []string `json:"roles"`

	// scope
	// Required: true
	Scope *string `json:"scope"`

	// token type
	// Required: true
	TokenType *string `json:"token_type"`

	// present if it is user token
	UserID string `json:"user_id,omitempty"`

	// xuid
	Xuid string `json:"xuid,omitempty"`
}

// Validate validates this oauthmodel token with device cookie response v3
func (m *OauthmodelTokenWithDeviceCookieResponseV3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaceRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OauthmodelTokenWithDeviceCookieResponseV3) validateAccessToken(formats strfmt.Registry) error {

	if err := validate.Required("access_token", "body", m.AccessToken); err != nil {
		return err
	}

	return nil
}

func (m *OauthmodelTokenWithDeviceCookieResponseV3) validateBans(formats strfmt.Registry) error {

	if err := validate.Required("bans", "body", m.Bans); err != nil {
		return err
	}

	for i := 0; i < len(m.Bans); i++ {
		if swag.IsZero(m.Bans[i]) { // not required
			continue
		}

		if m.Bans[i] != nil {
			if err := m.Bans[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OauthmodelTokenWithDeviceCookieResponseV3) validateExpiresIn(formats strfmt.Registry) error {

	if err := validate.Required("expires_in", "body", m.ExpiresIn); err != nil {
		return err
	}

	return nil
}

func (m *OauthmodelTokenWithDeviceCookieResponseV3) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *OauthmodelTokenWithDeviceCookieResponseV3) validateNamespaceRoles(formats strfmt.Registry) error {

	if err := validate.Required("namespace_roles", "body", m.NamespaceRoles); err != nil {
		return err
	}

	for i := 0; i < len(m.NamespaceRoles); i++ {
		if swag.IsZero(m.NamespaceRoles[i]) { // not required
			continue
		}

		if m.NamespaceRoles[i] != nil {
			if err := m.NamespaceRoles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("namespace_roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OauthmodelTokenWithDeviceCookieResponseV3) validatePermissions(formats strfmt.Registry) error {

	if err := validate.Required("permissions", "body", m.Permissions); err != nil {
		return err
	}

	for i := 0; i < len(m.Permissions); i++ {
		if swag.IsZero(m.Permissions[i]) { // not required
			continue
		}

		if m.Permissions[i] != nil {
			if err := m.Permissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OauthmodelTokenWithDeviceCookieResponseV3) validateRoles(formats strfmt.Registry) error {

	if err := validate.Required("roles", "body", m.Roles); err != nil {
		return err
	}

	return nil
}

func (m *OauthmodelTokenWithDeviceCookieResponseV3) validateScope(formats strfmt.Registry) error {

	if err := validate.Required("scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *OauthmodelTokenWithDeviceCookieResponseV3) validateTokenType(formats strfmt.Registry) error {

	if err := validate.Required("token_type", "body", m.TokenType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OauthmodelTokenWithDeviceCookieResponseV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OauthmodelTokenWithDeviceCookieResponseV3) UnmarshalBinary(b []byte) error {
	var res OauthmodelTokenWithDeviceCookieResponseV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}