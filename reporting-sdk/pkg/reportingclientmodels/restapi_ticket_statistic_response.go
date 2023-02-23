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

// RestapiTicketStatisticResponse restapi ticket statistic response
//
// swagger:model restapi.ticketStatisticResponse
type RestapiTicketStatisticResponse struct {

	// moderated count
	// Required: true
	ModeratedCount *int64 `json:"moderatedCount"`

	// open count
	// Required: true
	OpenCount *int64 `json:"openCount"`

	// total count
	// Required: true
	TotalCount *int64 `json:"totalCount"`
}

// Validate validates this restapi ticket statistic response
func (m *RestapiTicketStatisticResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModeratedCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestapiTicketStatisticResponse) validateModeratedCount(formats strfmt.Registry) error {

	if err := validate.Required("moderatedCount", "body", m.ModeratedCount); err != nil {
		return err
	}

	return nil
}

func (m *RestapiTicketStatisticResponse) validateOpenCount(formats strfmt.Registry) error {

	if err := validate.Required("openCount", "body", m.OpenCount); err != nil {
		return err
	}

	return nil
}

func (m *RestapiTicketStatisticResponse) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("totalCount", "body", m.TotalCount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestapiTicketStatisticResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestapiTicketStatisticResponse) UnmarshalBinary(b []byte) error {
	var res RestapiTicketStatisticResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}