// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ItemPurchaseConditionValidateResult item purchase condition validate result
//
// swagger:model ItemPurchaseConditionValidateResult
type ItemPurchaseConditionValidateResult struct {

	// item id
	ItemID string `json:"itemId,omitempty"`

	// purchasable
	Purchasable bool `json:"purchasable"`

	// item sku
	Sku string `json:"sku,omitempty"`

	// validate details: list of condition group validate result. index is same as purchase condition group.
	ValidateDetails []*ConditionGroupValidateResult `json:"validateDetails"`
}

// Validate validates this item purchase condition validate result
func (m *ItemPurchaseConditionValidateResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValidateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemPurchaseConditionValidateResult) validateValidateDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidateDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.ValidateDetails); i++ {
		if swag.IsZero(m.ValidateDetails[i]) { // not required
			continue
		}

		if m.ValidateDetails[i] != nil {
			if err := m.ValidateDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validateDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemPurchaseConditionValidateResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemPurchaseConditionValidateResult) UnmarshalBinary(b []byte) error {
	var res ItemPurchaseConditionValidateResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}