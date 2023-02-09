// Code generated by go-swagger; DO NOT EDIT.

package reportingclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestapiConfigResponse restapi config response
//
// swagger:model restapi.configResponse
type RestapiConfigResponse struct {

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// reporting limit
	// Required: true
	ReportingLimit *RestapiReportingLimit `json:"reportingLimit"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt"`
}

// Validate validates this restapi config response
func (m *RestapiConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportingLimit(formats); err != nil {
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

func (m *RestapiConfigResponse) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *RestapiConfigResponse) validateReportingLimit(formats strfmt.Registry) error {

	if err := validate.Required("reportingLimit", "body", m.ReportingLimit); err != nil {
		return err
	}

	if m.ReportingLimit != nil {
		if err := m.ReportingLimit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reportingLimit")
			}
			return err
		}
	}

	return nil
}

func (m *RestapiConfigResponse) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestapiConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestapiConfigResponse) UnmarshalBinary(b []byte) error {
	var res RestapiConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
