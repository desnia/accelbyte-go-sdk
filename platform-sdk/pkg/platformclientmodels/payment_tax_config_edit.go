// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaymentTaxConfigEdit A DTO object for updating payment tax config
//
// swagger:model PaymentTaxConfigEdit
type PaymentTaxConfigEdit struct {

	// sandbox taxJar API token
	SandboxTaxJarAPIToken string `json:"sandboxTaxJarApiToken,omitempty"`

	// taxJar API token
	TaxJarAPIToken string `json:"taxJarApiToken,omitempty"`

	// if taxJar integration is enabled
	TaxJarEnabled bool `json:"taxJarEnabled,omitempty"`

	// taxJar product codes mappings, allow item types: APP, COINS, INGAMEITEM, BUNDLE, CODE, SUBSCRIPTION
	TaxJarProductCodesMapping map[string]string `json:"taxJarProductCodesMapping,omitempty"`
}

// Validate validates this payment tax config edit
func (m *PaymentTaxConfigEdit) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentTaxConfigEdit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentTaxConfigEdit) UnmarshalBinary(b []byte) error {
	var res PaymentTaxConfigEdit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}