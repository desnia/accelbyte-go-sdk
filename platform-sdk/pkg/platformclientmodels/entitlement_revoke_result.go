// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EntitlementRevokeResult entitlement revoke result
//
// swagger:model EntitlementRevokeResult
type EntitlementRevokeResult struct {

	// entitlement Id
	EntitlementID string `json:"entitlementId,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`
}

// Validate validates this entitlement revoke result
func (m *EntitlementRevokeResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EntitlementRevokeResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntitlementRevokeResult) UnmarshalBinary(b []byte) error {
	var res EntitlementRevokeResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}