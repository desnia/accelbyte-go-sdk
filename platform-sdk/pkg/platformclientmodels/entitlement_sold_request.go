// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EntitlementSoldRequest A DTO object for sell entitlement.
//
// swagger:model EntitlementSoldRequest
type EntitlementSoldRequest struct {

	// Request id(Optional), client should provide a unique request id to perform at most once execution, When a request id is resubmitted, it will return original successful response
	RequestID string `json:"requestId,omitempty"`

	// the use count to sell, 1 for default, durable entitlement only allow 1
	UseCount int32 `json:"useCount,omitempty"`
}

// Validate validates this entitlement sold request
func (m *EntitlementSoldRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EntitlementSoldRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntitlementSoldRequest) UnmarshalBinary(b []byte) error {
	var res EntitlementSoldRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}