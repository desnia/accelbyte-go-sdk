// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaymentTaxConfigInfo payment tax config info
//
// swagger:model PaymentTaxConfigInfo
type PaymentTaxConfigInfo struct {

	// sandbox taxJar API token
	SandboxTaxJarAPIToken string `json:"sandboxTaxJarApiToken,omitempty"`

	// taxJar API token
	TaxJarAPIToken string `json:"taxJarApiToken,omitempty"`

	// if taxJar integration is enabled
	TaxJarEnabled bool `json:"taxJarEnabled,omitempty"`

	// taxJar product codes mappings
	TaxJarProductCodesMapping map[string]string `json:"taxJarProductCodesMapping,omitempty"`
}

// Validate validates this payment tax config info
func (m *PaymentTaxConfigInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentTaxConfigInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentTaxConfigInfo) UnmarshalBinary(b []byte) error {
	var res PaymentTaxConfigInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}