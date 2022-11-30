// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PlayStationDLCSyncMultiServiceLabelsRequest play station d l c sync multi service labels request
//
// swagger:model PlayStationDLCSyncMultiServiceLabelsRequest
type PlayStationDLCSyncMultiServiceLabelsRequest struct {

	// service labels
	// Unique: true
	ServiceLabels []int32 `json:"serviceLabels"`
}

// Validate validates this play station d l c sync multi service labels request
func (m *PlayStationDLCSyncMultiServiceLabelsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayStationDLCSyncMultiServiceLabelsRequest) validateServiceLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceLabels) { // not required
		return nil
	}

	if err := validate.UniqueItems("serviceLabels", "body", m.ServiceLabels); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlayStationDLCSyncMultiServiceLabelsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayStationDLCSyncMultiServiceLabelsRequest) UnmarshalBinary(b []byte) error {
	var res PlayStationDLCSyncMultiServiceLabelsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}