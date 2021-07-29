// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RewardItem reward item
//
// swagger:model RewardItem
type RewardItem struct {

	// item Id
	ItemID string `json:"itemId,omitempty"`

	// quantity
	Quantity int32 `json:"quantity,omitempty"`
}

// Validate validates this reward item
func (m *RewardItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RewardItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RewardItem) UnmarshalBinary(b []byte) error {
	var res RewardItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}