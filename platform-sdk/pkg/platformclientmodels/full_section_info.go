// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FullSectionInfo full section info
//
// swagger:model FullSectionInfo
type FullSectionInfo struct {

	// active
	// Required: true
	Active *bool `json:"active"`

	// view created time
	// Required: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt"`

	// display order
	DisplayOrder int32 `json:"displayOrder,omitempty"`

	// end date
	// Required: true
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate"`

	// section extension
	Ext map[string]interface{} `json:"ext,omitempty"`

	// fix period rotation config, don't allow null when rotation type is FIXED_PERIOD
	FixedPeriodRotationConfig *FixedPeriodRotationConfig `json:"fixedPeriodRotationConfig,omitempty"`

	// section item namings
	ItemNamings []*ItemNaming `json:"itemNamings"`

	// section items
	Items []*SectionItem `json:"items"`

	// Localization, key language, value localization content
	// Required: true
	Localizations map[string]Localization `json:"localizations"`

	// name
	// Required: true
	Name *string `json:"name"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// rotation type, default is NONE
	// Enum: [CUSTOM FIXED_PERIOD NONE]
	RotationType string `json:"rotationType,omitempty"`

	// id
	// Required: true
	SectionID *string `json:"sectionId"`

	// start date
	// Required: true
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate"`

	// view updated time
	// Required: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt"`

	// view id
	ViewID string `json:"viewId,omitempty"`

	// view name
	ViewName string `json:"viewName,omitempty"`
}

// Validate validates this full section info
func (m *FullSectionInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFixedPeriodRotationConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemNamings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRotationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSectionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
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

func (m *FullSectionInfo) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", m.Active); err != nil {
		return err
	}

	return nil
}

func (m *FullSectionInfo) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FullSectionInfo) validateEndDate(formats strfmt.Registry) error {

	if err := validate.Required("endDate", "body", strfmt.DateTime(m.EndDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FullSectionInfo) validateFixedPeriodRotationConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.FixedPeriodRotationConfig) { // not required
		return nil
	}

	if m.FixedPeriodRotationConfig != nil {
		if err := m.FixedPeriodRotationConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fixedPeriodRotationConfig")
			}
			return err
		}
	}

	return nil
}

func (m *FullSectionInfo) validateItemNamings(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemNamings) { // not required
		return nil
	}

	for i := 0; i < len(m.ItemNamings); i++ {
		if swag.IsZero(m.ItemNamings[i]) { // not required
			continue
		}

		if m.ItemNamings[i] != nil {
			if err := m.ItemNamings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("itemNamings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FullSectionInfo) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FullSectionInfo) validateLocalizations(formats strfmt.Registry) error {

	for k := range m.Localizations {

		if err := validate.Required("localizations"+"."+k, "body", m.Localizations[k]); err != nil {
			return err
		}
		if val, ok := m.Localizations[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *FullSectionInfo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FullSectionInfo) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

var fullSectionInfoTypeRotationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CUSTOM","FIXED_PERIOD","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fullSectionInfoTypeRotationTypePropEnum = append(fullSectionInfoTypeRotationTypePropEnum, v)
	}
}

const (

	// FullSectionInfoRotationTypeCUSTOM captures enum value "CUSTOM"
	FullSectionInfoRotationTypeCUSTOM string = "CUSTOM"

	// FullSectionInfoRotationTypeFIXEDPERIOD captures enum value "FIXED_PERIOD"
	FullSectionInfoRotationTypeFIXEDPERIOD string = "FIXED_PERIOD"

	// FullSectionInfoRotationTypeNONE captures enum value "NONE"
	FullSectionInfoRotationTypeNONE string = "NONE"
)

// prop value enum
func (m *FullSectionInfo) validateRotationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fullSectionInfoTypeRotationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FullSectionInfo) validateRotationType(formats strfmt.Registry) error {

	if swag.IsZero(m.RotationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRotationTypeEnum("rotationType", "body", m.RotationType); err != nil {
		return err
	}

	return nil
}

func (m *FullSectionInfo) validateSectionID(formats strfmt.Registry) error {

	if err := validate.Required("sectionId", "body", m.SectionID); err != nil {
		return err
	}

	return nil
}

func (m *FullSectionInfo) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("startDate", "body", strfmt.DateTime(m.StartDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FullSectionInfo) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FullSectionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FullSectionInfo) UnmarshalBinary(b []byte) error {
	var res FullSectionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
