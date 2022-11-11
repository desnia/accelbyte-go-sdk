// Code generated by go-swagger; DO NOT EDIT.

package sessionclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ApimodelsDSInformationResponse apimodels d s information response
//
// swagger:model apimodels.DSInformationResponse
type ApimodelsDSInformationResponse struct {

	// requested at
	// Required: true
	RequestedAt *string `json:"RequestedAt"`

	// server
	// Required: true
	Server *ModelsGameServer `json:"Server"`

	// status
	// Required: true
	Status *string `json:"Status"`
}

// Validate validates this apimodels d s information response
func (m *ApimodelsDSInformationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequestedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
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

func (m *ApimodelsDSInformationResponse) validateRequestedAt(formats strfmt.Registry) error {

	if err := validate.Required("RequestedAt", "body", m.RequestedAt); err != nil {
		return err
	}

	return nil
}

func (m *ApimodelsDSInformationResponse) validateServer(formats strfmt.Registry) error {

	if err := validate.Required("Server", "body", m.Server); err != nil {
		return err
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Server")
			}
			return err
		}
	}

	return nil
}

func (m *ApimodelsDSInformationResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("Status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApimodelsDSInformationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApimodelsDSInformationResponse) UnmarshalBinary(b []byte) error {
	var res ApimodelsDSInformationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}