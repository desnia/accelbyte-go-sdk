// Code generated by go-swagger; DO NOT EDIT.

package legalclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RetrievePolicyPublicResponse retrieve policy public response
//
// swagger:model RetrievePolicyPublicResponse
type RetrievePolicyPublicResponse struct {

	// base policy Id
	// Required: true
	BasePolicyID *string `json:"basePolicyId"`

	// base urls
	BaseUrls []string `json:"baseUrls"`

	// country code
	// Required: true
	CountryCode *string `json:"countryCode"`

	// country group code
	CountryGroupCode string `json:"countryGroupCode,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// is default opted
	// Required: true
	IsDefaultOpted *bool `json:"isDefaultOpted"`

	// is default selection
	// Required: true
	IsDefaultSelection *bool `json:"isDefaultSelection"`

	// is mandatory
	// Required: true
	IsMandatory *bool `json:"isMandatory"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// policy name
	// Required: true
	PolicyName *string `json:"policyName"`

	// policy type
	// Required: true
	PolicyType *string `json:"policyType"`

	// policy versions
	PolicyVersions []*PolicyVersionWithLocalizedVersionObject `json:"policyVersions"`

	// readable Id
	ReadableID string `json:"readableId,omitempty"`

	// should notify on update
	// Required: true
	ShouldNotifyOnUpdate *bool `json:"shouldNotifyOnUpdate"`

	// tags
	// Unique: true
	Tags []string `json:"tags"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this retrieve policy public response
func (m *RetrievePolicyPublicResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBasePolicyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountryCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsDefaultOpted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsDefaultSelection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsMandatory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShouldNotifyOnUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
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

func (m *RetrievePolicyPublicResponse) validateBasePolicyID(formats strfmt.Registry) error {

	if err := validate.Required("basePolicyId", "body", m.BasePolicyID); err != nil {
		return err
	}

	return nil
}

func (m *RetrievePolicyPublicResponse) validateCountryCode(formats strfmt.Registry) error {

	if err := validate.Required("countryCode", "body", m.CountryCode); err != nil {
		return err
	}

	return nil
}

func (m *RetrievePolicyPublicResponse) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RetrievePolicyPublicResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RetrievePolicyPublicResponse) validateIsDefaultOpted(formats strfmt.Registry) error {

	if err := validate.Required("isDefaultOpted", "body", m.IsDefaultOpted); err != nil {
		return err
	}

	return nil
}

func (m *RetrievePolicyPublicResponse) validateIsDefaultSelection(formats strfmt.Registry) error {

	if err := validate.Required("isDefaultSelection", "body", m.IsDefaultSelection); err != nil {
		return err
	}

	return nil
}

func (m *RetrievePolicyPublicResponse) validateIsMandatory(formats strfmt.Registry) error {

	if err := validate.Required("isMandatory", "body", m.IsMandatory); err != nil {
		return err
	}

	return nil
}

func (m *RetrievePolicyPublicResponse) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *RetrievePolicyPublicResponse) validatePolicyName(formats strfmt.Registry) error {

	if err := validate.Required("policyName", "body", m.PolicyName); err != nil {
		return err
	}

	return nil
}

func (m *RetrievePolicyPublicResponse) validatePolicyType(formats strfmt.Registry) error {

	if err := validate.Required("policyType", "body", m.PolicyType); err != nil {
		return err
	}

	return nil
}

func (m *RetrievePolicyPublicResponse) validatePolicyVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.PolicyVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.PolicyVersions); i++ {
		if swag.IsZero(m.PolicyVersions[i]) { // not required
			continue
		}

		if m.PolicyVersions[i] != nil {
			if err := m.PolicyVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policyVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RetrievePolicyPublicResponse) validateShouldNotifyOnUpdate(formats strfmt.Registry) error {

	if err := validate.Required("shouldNotifyOnUpdate", "body", m.ShouldNotifyOnUpdate); err != nil {
		return err
	}

	return nil
}

func (m *RetrievePolicyPublicResponse) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if err := validate.UniqueItems("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

func (m *RetrievePolicyPublicResponse) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RetrievePolicyPublicResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetrievePolicyPublicResponse) UnmarshalBinary(b []byte) error {
	var res RetrievePolicyPublicResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}