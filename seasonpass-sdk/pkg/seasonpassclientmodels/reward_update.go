// Code generated by go-swagger; DO NOT EDIT.

package seasonpassclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RewardUpdate A DTO object for updating reward API call.
//
// swagger:model RewardUpdate
type RewardUpdate struct {

	// currency, required when reward type is CURRENCY
	Currency *RewardCurrency `json:"currency,omitempty"`

	// image, thumbnail for reward
	Image *Image `json:"image,omitempty"`

	// itemId, required when reward type is ITEM, the item type should be one of: INGAMEITEM,COINS,BUNDLE
	ItemID string `json:"itemId,omitempty"`

	// nullFields
	// Unique: true
	NullFields []string `json:"nullFields"`

	// Item quantity or Currency amount, default 1
	Quantity int32 `json:"quantity,omitempty"`

	// type, at current only support ITEM
	// Enum: [ITEM CURRENCY]
	Type string `json:"type,omitempty"`
}

// Validate validates this reward update
func (m *RewardUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNullFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RewardUpdate) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	if m.Currency != nil {
		if err := m.Currency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currency")
			}
			return err
		}
	}

	return nil
}

func (m *RewardUpdate) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *RewardUpdate) validateNullFields(formats strfmt.Registry) error {

	if swag.IsZero(m.NullFields) { // not required
		return nil
	}

	if err := validate.UniqueItems("nullFields", "body", m.NullFields); err != nil {
		return err
	}

	return nil
}

var rewardUpdateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ITEM","CURRENCY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rewardUpdateTypeTypePropEnum = append(rewardUpdateTypeTypePropEnum, v)
	}
}

const (

	// RewardUpdateTypeITEM captures enum value "ITEM"
	RewardUpdateTypeITEM string = "ITEM"

	// RewardUpdateTypeCURRENCY captures enum value "CURRENCY"
	RewardUpdateTypeCURRENCY string = "CURRENCY"
)

// prop value enum
func (m *RewardUpdate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rewardUpdateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RewardUpdate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RewardUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RewardUpdate) UnmarshalBinary(b []byte) error {
	var res RewardUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}