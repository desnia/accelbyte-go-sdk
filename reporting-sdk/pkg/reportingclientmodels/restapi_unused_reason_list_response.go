// Code generated by go-swagger; DO NOT EDIT.

package reportingclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestapiUnusedReasonListResponse restapi unused reason list response
//
// swagger:model restapi.UnusedReasonListResponse
type RestapiUnusedReasonListResponse struct {

	// reasons
	// Required: true
	Reasons []*RestapiPublicReasonResponse `json:"reasons"`
}

// Validate validates this restapi unused reason list response
func (m *RestapiUnusedReasonListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReasons(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestapiUnusedReasonListResponse) validateReasons(formats strfmt.Registry) error {

	if err := validate.Required("reasons", "body", m.Reasons); err != nil {
		return err
	}

	for i := 0; i < len(m.Reasons); i++ {
		if swag.IsZero(m.Reasons[i]) { // not required
			continue
		}

		if m.Reasons[i] != nil {
			if err := m.Reasons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reasons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestapiUnusedReasonListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestapiUnusedReasonListResponse) UnmarshalBinary(b []byte) error {
	var res RestapiUnusedReasonListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
