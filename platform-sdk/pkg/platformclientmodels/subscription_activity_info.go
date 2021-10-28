// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SubscriptionActivityInfo subscription activity info
//
// swagger:model SubscriptionActivityInfo
type SubscriptionActivityInfo struct {

	// subscription action
	// Required: true
	// Enum: [SUBSCRIBE CANCEL IMMEDIATE_CANCEL RESUBSCRIBE GRANT_DAYS CHANGE_BILLING_ACCOUNT]
	Action *string `json:"action"`

	// subscription already charged cycle
	// Required: true
	ChargedCycles *int32 `json:"chargedCycles"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// subscription current cycle number
	// Required: true
	CurrentCycle *int32 `json:"currentCycle"`

	// subscription grant days, negative indicates decrease
	GrantDays int32 `json:"grantDays,omitempty"`

	// Subscription in fixed cycle trial or not
	InFixedCycleTrial bool `json:"inFixedCycleTrial,omitempty"`

	// Subscription in fixed free days or not
	InFixedFreeDays bool `json:"inFixedFreeDays,omitempty"`

	// Subscription namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// operator
	// Required: true
	Operator *string `json:"operator"`

	// subscription action reason
	Reason string `json:"reason,omitempty"`

	// subscription subscribed by
	// Required: true
	// Enum: [USER PLATFORM]
	SubscribedBy *string `json:"subscribedBy"`

	// Subscription id
	// Required: true
	SubscriptionID *string `json:"subscriptionId"`

	// Subscription already trialed cycles if exist
	TrialedCycles int32 `json:"trialedCycles,omitempty"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`

	// Subscription user id
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this subscription activity info
func (m *SubscriptionActivityInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargedCycles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentCycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscribedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var subscriptionActivityInfoTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUBSCRIBE","CANCEL","IMMEDIATE_CANCEL","RESUBSCRIBE","GRANT_DAYS","CHANGE_BILLING_ACCOUNT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionActivityInfoTypeActionPropEnum = append(subscriptionActivityInfoTypeActionPropEnum, v)
	}
}

const (

	// SubscriptionActivityInfoActionSUBSCRIBE captures enum value "SUBSCRIBE"
	SubscriptionActivityInfoActionSUBSCRIBE string = "SUBSCRIBE"

	// SubscriptionActivityInfoActionCANCEL captures enum value "CANCEL"
	SubscriptionActivityInfoActionCANCEL string = "CANCEL"

	// SubscriptionActivityInfoActionIMMEDIATECANCEL captures enum value "IMMEDIATE_CANCEL"
	SubscriptionActivityInfoActionIMMEDIATECANCEL string = "IMMEDIATE_CANCEL"

	// SubscriptionActivityInfoActionRESUBSCRIBE captures enum value "RESUBSCRIBE"
	SubscriptionActivityInfoActionRESUBSCRIBE string = "RESUBSCRIBE"

	// SubscriptionActivityInfoActionGRANTDAYS captures enum value "GRANT_DAYS"
	SubscriptionActivityInfoActionGRANTDAYS string = "GRANT_DAYS"

	// SubscriptionActivityInfoActionCHANGEBILLINGACCOUNT captures enum value "CHANGE_BILLING_ACCOUNT"
	SubscriptionActivityInfoActionCHANGEBILLINGACCOUNT string = "CHANGE_BILLING_ACCOUNT"
)

// prop value enum
func (m *SubscriptionActivityInfo) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, subscriptionActivityInfoTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionActivityInfo) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionActivityInfo) validateChargedCycles(formats strfmt.Registry) error {

	if err := validate.Required("chargedCycles", "body", m.ChargedCycles); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionActivityInfo) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionActivityInfo) validateCurrentCycle(formats strfmt.Registry) error {

	if err := validate.Required("currentCycle", "body", m.CurrentCycle); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionActivityInfo) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionActivityInfo) validateOperator(formats strfmt.Registry) error {

	if err := validate.Required("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

var subscriptionActivityInfoTypeSubscribedByPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["USER","PLATFORM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionActivityInfoTypeSubscribedByPropEnum = append(subscriptionActivityInfoTypeSubscribedByPropEnum, v)
	}
}

const (

	// SubscriptionActivityInfoSubscribedByUSER captures enum value "USER"
	SubscriptionActivityInfoSubscribedByUSER string = "USER"

	// SubscriptionActivityInfoSubscribedByPLATFORM captures enum value "PLATFORM"
	SubscriptionActivityInfoSubscribedByPLATFORM string = "PLATFORM"
)

// prop value enum
func (m *SubscriptionActivityInfo) validateSubscribedByEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, subscriptionActivityInfoTypeSubscribedByPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionActivityInfo) validateSubscribedBy(formats strfmt.Registry) error {

	if err := validate.Required("subscribedBy", "body", m.SubscribedBy); err != nil {
		return err
	}

	// value enum
	if err := m.validateSubscribedByEnum("subscribedBy", "body", *m.SubscribedBy); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionActivityInfo) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionId", "body", m.SubscriptionID); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionActivityInfo) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionActivityInfo) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionActivityInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionActivityInfo) UnmarshalBinary(b []byte) error {
	var res SubscriptionActivityInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}