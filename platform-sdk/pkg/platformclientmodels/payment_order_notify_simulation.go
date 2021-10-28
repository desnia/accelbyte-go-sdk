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

// PaymentOrderNotifySimulation payment order notify simulation
//
// swagger:model PaymentOrderNotifySimulation
type PaymentOrderNotifySimulation struct {

	// amount
	Amount int32 `json:"amount,omitempty"`

	// currency code
	// Required: true
	CurrencyCode *string `json:"currencyCode"`

	// notify type
	// Required: true
	// Enum: [CHARGE REFUND]
	NotifyType *string `json:"notifyType"`

	// payment provider
	// Required: true
	// Enum: [WALLET XSOLLA ADYEN STRIPE CHECKOUT ALIPAY WXPAY PAYPAL]
	PaymentProvider *string `json:"paymentProvider"`

	// salesTax for xsolla
	SalesTax int32 `json:"salesTax,omitempty"`

	// vat for xsolla
	Vat int32 `json:"vat,omitempty"`
}

// Validate validates this payment order notify simulation
func (m *PaymentOrderNotifySimulation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrencyCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentOrderNotifySimulation) validateCurrencyCode(formats strfmt.Registry) error {

	if err := validate.Required("currencyCode", "body", m.CurrencyCode); err != nil {
		return err
	}

	return nil
}

var paymentOrderNotifySimulationTypeNotifyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CHARGE","REFUND"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentOrderNotifySimulationTypeNotifyTypePropEnum = append(paymentOrderNotifySimulationTypeNotifyTypePropEnum, v)
	}
}

const (

	// PaymentOrderNotifySimulationNotifyTypeCHARGE captures enum value "CHARGE"
	PaymentOrderNotifySimulationNotifyTypeCHARGE string = "CHARGE"

	// PaymentOrderNotifySimulationNotifyTypeREFUND captures enum value "REFUND"
	PaymentOrderNotifySimulationNotifyTypeREFUND string = "REFUND"
)

// prop value enum
func (m *PaymentOrderNotifySimulation) validateNotifyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, paymentOrderNotifySimulationTypeNotifyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PaymentOrderNotifySimulation) validateNotifyType(formats strfmt.Registry) error {

	if err := validate.Required("notifyType", "body", m.NotifyType); err != nil {
		return err
	}

	// value enum
	if err := m.validateNotifyTypeEnum("notifyType", "body", *m.NotifyType); err != nil {
		return err
	}

	return nil
}

var paymentOrderNotifySimulationTypePaymentProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WALLET","XSOLLA","ADYEN","STRIPE","CHECKOUT","ALIPAY","WXPAY","PAYPAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentOrderNotifySimulationTypePaymentProviderPropEnum = append(paymentOrderNotifySimulationTypePaymentProviderPropEnum, v)
	}
}

const (

	// PaymentOrderNotifySimulationPaymentProviderWALLET captures enum value "WALLET"
	PaymentOrderNotifySimulationPaymentProviderWALLET string = "WALLET"

	// PaymentOrderNotifySimulationPaymentProviderXSOLLA captures enum value "XSOLLA"
	PaymentOrderNotifySimulationPaymentProviderXSOLLA string = "XSOLLA"

	// PaymentOrderNotifySimulationPaymentProviderADYEN captures enum value "ADYEN"
	PaymentOrderNotifySimulationPaymentProviderADYEN string = "ADYEN"

	// PaymentOrderNotifySimulationPaymentProviderSTRIPE captures enum value "STRIPE"
	PaymentOrderNotifySimulationPaymentProviderSTRIPE string = "STRIPE"

	// PaymentOrderNotifySimulationPaymentProviderCHECKOUT captures enum value "CHECKOUT"
	PaymentOrderNotifySimulationPaymentProviderCHECKOUT string = "CHECKOUT"

	// PaymentOrderNotifySimulationPaymentProviderALIPAY captures enum value "ALIPAY"
	PaymentOrderNotifySimulationPaymentProviderALIPAY string = "ALIPAY"

	// PaymentOrderNotifySimulationPaymentProviderWXPAY captures enum value "WXPAY"
	PaymentOrderNotifySimulationPaymentProviderWXPAY string = "WXPAY"

	// PaymentOrderNotifySimulationPaymentProviderPAYPAL captures enum value "PAYPAL"
	PaymentOrderNotifySimulationPaymentProviderPAYPAL string = "PAYPAL"
)

// prop value enum
func (m *PaymentOrderNotifySimulation) validatePaymentProviderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, paymentOrderNotifySimulationTypePaymentProviderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PaymentOrderNotifySimulation) validatePaymentProvider(formats strfmt.Registry) error {

	if err := validate.Required("paymentProvider", "body", m.PaymentProvider); err != nil {
		return err
	}

	// value enum
	if err := m.validatePaymentProviderEnum("paymentProvider", "body", *m.PaymentProvider); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentOrderNotifySimulation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentOrderNotifySimulation) UnmarshalBinary(b []byte) error {
	var res PaymentOrderNotifySimulation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}