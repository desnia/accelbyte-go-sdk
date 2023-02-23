// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RevocationConfigInfo revocation config info
//
// swagger:model RevocationConfigInfo
type RevocationConfigInfo struct {

	// entitlement
	Entitlement *EntitlementRevocationConfig `json:"entitlement,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// wallet
	Wallet *WalletRevocationConfig `json:"wallet,omitempty"`
}

// Validate validates this revocation config info
func (m *RevocationConfigInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntitlement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWallet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RevocationConfigInfo) validateEntitlement(formats strfmt.Registry) error {

	if swag.IsZero(m.Entitlement) { // not required
		return nil
	}

	if m.Entitlement != nil {
		if err := m.Entitlement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entitlement")
			}
			return err
		}
	}

	return nil
}

func (m *RevocationConfigInfo) validateWallet(formats strfmt.Registry) error {

	if swag.IsZero(m.Wallet) { // not required
		return nil
	}

	if m.Wallet != nil {
		if err := m.Wallet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wallet")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RevocationConfigInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RevocationConfigInfo) UnmarshalBinary(b []byte) error {
	var res RevocationConfigInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}