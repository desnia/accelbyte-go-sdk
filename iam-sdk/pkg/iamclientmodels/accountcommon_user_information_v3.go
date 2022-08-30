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

// AccountcommonUserInformationV3 accountcommon user information v3
//
// swagger:model accountcommon.UserInformationV3
type AccountcommonUserInformationV3 struct {

	// country
	Country string `json:"country,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// email addresses
	// Required: true
	EmailAddresses []string `json:"emailAddresses"`

	// phone number
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// platform users
	// Required: true
	PlatformUsers []*AccountcommonPlatformUserInformationV3 `json:"platformUsers"`

	// username
	Username string `json:"username,omitempty"`

	// xbox user Id
	XboxUserID string `json:"xboxUserId,omitempty"`
}

// Validate validates this accountcommon user information v3
func (m *AccountcommonUserInformationV3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmailAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountcommonUserInformationV3) validateEmailAddresses(formats strfmt.Registry) error {

	if err := validate.Required("emailAddresses", "body", m.EmailAddresses); err != nil {
		return err
	}

	return nil
}

func (m *AccountcommonUserInformationV3) validatePlatformUsers(formats strfmt.Registry) error {

	if err := validate.Required("platformUsers", "body", m.PlatformUsers); err != nil {
		return err
	}

	for i := 0; i < len(m.PlatformUsers); i++ {
		if swag.IsZero(m.PlatformUsers[i]) { // not required
			continue
		}

		if m.PlatformUsers[i] != nil {
			if err := m.PlatformUsers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("platformUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountcommonUserInformationV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountcommonUserInformationV3) UnmarshalBinary(b []byte) error {
	var res AccountcommonUserInformationV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}