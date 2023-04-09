// Code generated by go-swagger; DO NOT EDIT.

package sessionclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ApimodelsDeleteBulkGameSessionsAPIResponse apimodels delete bulk game sessions API response
//
// swagger:model apimodels.DeleteBulkGameSessionsAPIResponse
type ApimodelsDeleteBulkGameSessionsAPIResponse struct {

	// failed
	// Required: true
	Failed []*ApimodelsResponseDeleteBulkGameSessions `json:"failed"`

	// success
	// Required: true
	Success []string `json:"success"`
}

// Validate validates this apimodels delete bulk game sessions API response
func (m *ApimodelsDeleteBulkGameSessionsAPIResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFailed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApimodelsDeleteBulkGameSessionsAPIResponse) validateFailed(formats strfmt.Registry) error {

	if err := validate.Required("failed", "body", m.Failed); err != nil {
		return err
	}

	for i := 0; i < len(m.Failed); i++ {
		if swag.IsZero(m.Failed[i]) { // not required
			continue
		}

		if m.Failed[i] != nil {
			if err := m.Failed[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("failed" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApimodelsDeleteBulkGameSessionsAPIResponse) validateSuccess(formats strfmt.Registry) error {

	if err := validate.Required("success", "body", m.Success); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApimodelsDeleteBulkGameSessionsAPIResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApimodelsDeleteBulkGameSessionsAPIResponse) UnmarshalBinary(b []byte) error {
	var res ApimodelsDeleteBulkGameSessionsAPIResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
