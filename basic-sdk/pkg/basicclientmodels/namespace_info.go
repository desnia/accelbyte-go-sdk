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

// NamespaceInfo namespace info
//
// swagger:model NamespaceInfo
type NamespaceInfo struct {

	// clientId is only present on multi tenant mode with namespace is not publisher namespace
	ClientID string `json:"clientId,omitempty"`

	// created at
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt,omitempty"`

	// namespace display name
	// Required: true
	DisplayName *string `json:"displayName"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// parentNamespace is only present in multi tenant mode
	ParentNamespace string `json:"parentNamespace,omitempty"`

	// status
	// Enum: [ACTIVE DELETED INACTIVE]
	Status string `json:"status,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this namespace info
func (m *NamespaceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *NamespaceInfo) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NamespaceInfo) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *NamespaceInfo) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

var namespaceInfoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","DELETED","INACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		namespaceInfoTypeStatusPropEnum = append(namespaceInfoTypeStatusPropEnum, v)
	}
}

const (

	// NamespaceInfoStatusACTIVE captures enum value "ACTIVE"
	NamespaceInfoStatusACTIVE string = "ACTIVE"

	// NamespaceInfoStatusDELETED captures enum value "DELETED"
	NamespaceInfoStatusDELETED string = "DELETED"

	// NamespaceInfoStatusINACTIVE captures enum value "INACTIVE"
	NamespaceInfoStatusINACTIVE string = "INACTIVE"
)

// prop value enum
func (m *NamespaceInfo) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, namespaceInfoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NamespaceInfo) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *NamespaceInfo) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NamespaceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamespaceInfo) UnmarshalBinary(b []byte) error {
	var res NamespaceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
