// Code generated by go-swagger; DO NOT EDIT.

package eventlogclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsEventResponse models event response
//
// swagger:model models.EventResponse
type ModelsEventResponse struct {

	// data
	// Required: true
	Data []*ModelsEvent `json:"Data"`

	// pagination
	// Required: true
	Pagination *ModelsPagination `json:"Pagination"`
}

// Validate validates this models event response
func (m *ModelsEventResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsEventResponse) validateData(formats strfmt.Registry) error {

	if err := validate.Required("Data", "body", m.Data); err != nil {
		return err
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsEventResponse) validatePagination(formats strfmt.Registry) error {

	if err := validate.Required("Pagination", "body", m.Pagination); err != nil {
		return err
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsEventResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsEventResponse) UnmarshalBinary(b []byte) error {
	var res ModelsEventResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}