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

// RestapiActionAPIRequest restapi action Api request
//
// swagger:model restapi.actionApiRequest
type RestapiActionAPIRequest struct {

	// Auto mod action ID
	// Required: true
	ActionID *string `json:"actionId"`

	// Auto mod action display name
	// Required: true
	ActionName *string `json:"actionName"`

	// Kafka publish event name
	// Required: true
	EventName *string `json:"eventName"`
}

// Validate validates this restapi action Api request
func (m *RestapiActionAPIRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActionName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestapiActionAPIRequest) validateActionID(formats strfmt.Registry) error {

	if err := validate.Required("actionId", "body", m.ActionID); err != nil {
		return err
	}

	return nil
}

func (m *RestapiActionAPIRequest) validateActionName(formats strfmt.Registry) error {

	if err := validate.Required("actionName", "body", m.ActionName); err != nil {
		return err
	}

	return nil
}

func (m *RestapiActionAPIRequest) validateEventName(formats strfmt.Registry) error {

	if err := validate.Required("eventName", "body", m.EventName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestapiActionAPIRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestapiActionAPIRequest) UnmarshalBinary(b []byte) error {
	var res RestapiActionAPIRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
