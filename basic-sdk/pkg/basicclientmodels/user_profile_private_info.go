// Code generated by go-swagger; DO NOT EDIT.

package basicclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserProfilePrivateInfo user profile private info
//
// swagger:model UserProfilePrivateInfo
type UserProfilePrivateInfo struct {

	// avatar large Url
	AvatarLargeURL string `json:"avatarLargeUrl,omitempty"`

	// avatar small Url
	AvatarSmallURL string `json:"avatarSmallUrl,omitempty"`

	// avatar Url
	AvatarURL string `json:"avatarUrl,omitempty"`

	// custom attributes
	CustomAttributes map[string]interface{} `json:"customAttributes,omitempty"`

	// date of birth
	// Format: date
	DateOfBirth strfmt.Date `json:"dateOfBirth,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// language
	Language string `json:"language,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// private custom attributes
	PrivateCustomAttributes map[string]interface{} `json:"privateCustomAttributes,omitempty"`

	// status
	// Enum: [ACTIVE INACTIVE]
	Status string `json:"status,omitempty"`

	// time zone
	TimeZone string `json:"timeZone,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`

	// zip code
	ZipCode string `json:"zipCode,omitempty"`
}

// Validate validates this user profile private info
func (m *UserProfilePrivateInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateOfBirth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserProfilePrivateInfo) validateDateOfBirth(formats strfmt.Registry) error {

	if swag.IsZero(m.DateOfBirth) { // not required
		return nil
	}

	if err := validate.FormatOf("dateOfBirth", "body", "date", m.DateOfBirth.String(), formats); err != nil {
		return err
	}

	return nil
}

var userProfilePrivateInfoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","INACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userProfilePrivateInfoTypeStatusPropEnum = append(userProfilePrivateInfoTypeStatusPropEnum, v)
	}
}

const (

	// UserProfilePrivateInfoStatusACTIVE captures enum value "ACTIVE"
	UserProfilePrivateInfoStatusACTIVE string = "ACTIVE"

	// UserProfilePrivateInfoStatusINACTIVE captures enum value "INACTIVE"
	UserProfilePrivateInfoStatusINACTIVE string = "INACTIVE"
)

// prop value enum
func (m *UserProfilePrivateInfo) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userProfilePrivateInfoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserProfilePrivateInfo) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserProfilePrivateInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserProfilePrivateInfo) UnmarshalBinary(b []byte) error {
	var res UserProfilePrivateInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}