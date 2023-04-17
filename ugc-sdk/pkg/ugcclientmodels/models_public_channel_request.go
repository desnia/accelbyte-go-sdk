// Code generated by go-swagger; DO NOT EDIT.

package ugcclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsPublicChannelRequest models public channel request
//
// swagger:model models.PublicChannelRequest
type ModelsPublicChannelRequest struct {

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this models public channel request
func (m *ModelsPublicChannelRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsPublicChannelRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsPublicChannelRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsPublicChannelRequest) UnmarshalBinary(b []byte) error {
	var res ModelsPublicChannelRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}