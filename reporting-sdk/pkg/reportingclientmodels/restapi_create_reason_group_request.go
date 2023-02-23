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

// RestapiCreateReasonGroupRequest restapi create reason group request
//
// swagger:model restapi.createReasonGroupRequest
type RestapiCreateReasonGroupRequest struct {

	// reason ids
	ReasonIds []string `json:"reasonIds"`

	// max 256 chars
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this restapi create reason group request
func (m *RestapiCreateReasonGroupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestapiCreateReasonGroupRequest) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestapiCreateReasonGroupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestapiCreateReasonGroupRequest) UnmarshalBinary(b []byte) error {
	var res RestapiCreateReasonGroupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}