// Code generated by go-swagger; DO NOT EDIT.

package seasonpassclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Localization localization
//
// swagger:model Localization
type Localization struct {

	// description info
	Description string `json:"description,omitempty"`

	// title info
	Title string `json:"title,omitempty"`
}

// Validate validates this localization
func (m *Localization) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Localization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Localization) UnmarshalBinary(b []byte) error {
	var res Localization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}