// Code generated by go-swagger; DO NOT EDIT.

package payment_station

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPublicNormalizePaymentReturnURLParams creates a new PublicNormalizePaymentReturnURLParams object
// with the default values initialized.
func NewPublicNormalizePaymentReturnURLParams() *PublicNormalizePaymentReturnURLParams {
	var ()
	return &PublicNormalizePaymentReturnURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicNormalizePaymentReturnURLParamsWithTimeout creates a new PublicNormalizePaymentReturnURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicNormalizePaymentReturnURLParamsWithTimeout(timeout time.Duration) *PublicNormalizePaymentReturnURLParams {
	var ()
	return &PublicNormalizePaymentReturnURLParams{

		timeout: timeout,
	}
}

// NewPublicNormalizePaymentReturnURLParamsWithContext creates a new PublicNormalizePaymentReturnURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicNormalizePaymentReturnURLParamsWithContext(ctx context.Context) *PublicNormalizePaymentReturnURLParams {
	var ()
	return &PublicNormalizePaymentReturnURLParams{

		Context: ctx,
	}
}

// NewPublicNormalizePaymentReturnURLParamsWithHTTPClient creates a new PublicNormalizePaymentReturnURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicNormalizePaymentReturnURLParamsWithHTTPClient(client *http.Client) *PublicNormalizePaymentReturnURLParams {
	var ()
	return &PublicNormalizePaymentReturnURLParams{
		HTTPClient: client,
	}
}

/*PublicNormalizePaymentReturnURLParams contains all the parameters to send to the API endpoint
for the public normalize payment return Url operation typically these are written to a http.Request
*/
type PublicNormalizePaymentReturnURLParams struct {

	/*PayerID
	  PayPal payer id

	*/
	PayerID *string
	/*Foreinginvoice*/
	Foreinginvoice *string
	/*InvoiceID*/
	InvoiceID *string
	/*Namespace*/
	Namespace string
	/*OrderNo
	  Platform order no

	*/
	OrderNo string
	/*Payload*/
	Payload *string
	/*PaymentOrderNo
	  Platform payment order no

	*/
	PaymentOrderNo string
	/*PaymentProvider
	  Platform payment provider

	*/
	PaymentProvider string
	/*ResultCode*/
	ResultCode *string
	/*ReturnURL*/
	ReturnURL string
	/*Status*/
	Status *string
	/*Token
	  PayPal token

	*/
	Token *string
	/*Type*/
	Type *string
	/*UserID*/
	UserID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithTimeout(timeout time.Duration) *PublicNormalizePaymentReturnURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithContext(ctx context.Context) *PublicNormalizePaymentReturnURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithHTTPClient(client *http.Client) *PublicNormalizePaymentReturnURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayerID adds the payerID to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithPayerID(payerID *string) *PublicNormalizePaymentReturnURLParams {
	o.SetPayerID(payerID)
	return o
}

// SetPayerID adds the payerId to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetPayerID(payerID *string) {
	o.PayerID = payerID
}

// WithForeinginvoice adds the foreinginvoice to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithForeinginvoice(foreinginvoice *string) *PublicNormalizePaymentReturnURLParams {
	o.SetForeinginvoice(foreinginvoice)
	return o
}

// SetForeinginvoice adds the foreinginvoice to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetForeinginvoice(foreinginvoice *string) {
	o.Foreinginvoice = foreinginvoice
}

// WithInvoiceID adds the invoiceID to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithInvoiceID(invoiceID *string) *PublicNormalizePaymentReturnURLParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetInvoiceID(invoiceID *string) {
	o.InvoiceID = invoiceID
}

// WithNamespace adds the namespace to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithNamespace(namespace string) *PublicNormalizePaymentReturnURLParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOrderNo adds the orderNo to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithOrderNo(orderNo string) *PublicNormalizePaymentReturnURLParams {
	o.SetOrderNo(orderNo)
	return o
}

// SetOrderNo adds the orderNo to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetOrderNo(orderNo string) {
	o.OrderNo = orderNo
}

// WithPayload adds the payload to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithPayload(payload *string) *PublicNormalizePaymentReturnURLParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetPayload(payload *string) {
	o.Payload = payload
}

// WithPaymentOrderNo adds the paymentOrderNo to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithPaymentOrderNo(paymentOrderNo string) *PublicNormalizePaymentReturnURLParams {
	o.SetPaymentOrderNo(paymentOrderNo)
	return o
}

// SetPaymentOrderNo adds the paymentOrderNo to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetPaymentOrderNo(paymentOrderNo string) {
	o.PaymentOrderNo = paymentOrderNo
}

// WithPaymentProvider adds the paymentProvider to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithPaymentProvider(paymentProvider string) *PublicNormalizePaymentReturnURLParams {
	o.SetPaymentProvider(paymentProvider)
	return o
}

// SetPaymentProvider adds the paymentProvider to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetPaymentProvider(paymentProvider string) {
	o.PaymentProvider = paymentProvider
}

// WithResultCode adds the resultCode to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithResultCode(resultCode *string) *PublicNormalizePaymentReturnURLParams {
	o.SetResultCode(resultCode)
	return o
}

// SetResultCode adds the resultCode to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetResultCode(resultCode *string) {
	o.ResultCode = resultCode
}

// WithReturnURL adds the returnURL to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithReturnURL(returnURL string) *PublicNormalizePaymentReturnURLParams {
	o.SetReturnURL(returnURL)
	return o
}

// SetReturnURL adds the returnUrl to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetReturnURL(returnURL string) {
	o.ReturnURL = returnURL
}

// WithStatus adds the status to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithStatus(status *string) *PublicNormalizePaymentReturnURLParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetStatus(status *string) {
	o.Status = status
}

// WithToken adds the token to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithToken(token *string) *PublicNormalizePaymentReturnURLParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetToken(token *string) {
	o.Token = token
}

// WithType adds the typeVar to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithType(typeVar *string) *PublicNormalizePaymentReturnURLParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithUserID adds the userID to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) WithUserID(userID *string) *PublicNormalizePaymentReturnURLParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public normalize payment return Url params
func (o *PublicNormalizePaymentReturnURLParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicNormalizePaymentReturnURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PayerID != nil {

		// query param PayerID
		var qrPayerID string
		if o.PayerID != nil {
			qrPayerID = *o.PayerID
		}
		qPayerID := qrPayerID
		if qPayerID != "" {
			if err := r.SetQueryParam("PayerID", qPayerID); err != nil {
				return err
			}
		}

	}

	if o.Foreinginvoice != nil {

		// query param foreinginvoice
		var qrForeinginvoice string
		if o.Foreinginvoice != nil {
			qrForeinginvoice = *o.Foreinginvoice
		}
		qForeinginvoice := qrForeinginvoice
		if qForeinginvoice != "" {
			if err := r.SetQueryParam("foreinginvoice", qForeinginvoice); err != nil {
				return err
			}
		}

	}

	if o.InvoiceID != nil {

		// query param invoice_id
		var qrInvoiceID string
		if o.InvoiceID != nil {
			qrInvoiceID = *o.InvoiceID
		}
		qInvoiceID := qrInvoiceID
		if qInvoiceID != "" {
			if err := r.SetQueryParam("invoice_id", qInvoiceID); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// query param orderNo
	qrOrderNo := o.OrderNo
	qOrderNo := qrOrderNo
	if qOrderNo != "" {
		if err := r.SetQueryParam("orderNo", qOrderNo); err != nil {
			return err
		}
	}

	if o.Payload != nil {

		// query param payload
		var qrPayload string
		if o.Payload != nil {
			qrPayload = *o.Payload
		}
		qPayload := qrPayload
		if qPayload != "" {
			if err := r.SetQueryParam("payload", qPayload); err != nil {
				return err
			}
		}

	}

	// query param paymentOrderNo
	qrPaymentOrderNo := o.PaymentOrderNo
	qPaymentOrderNo := qrPaymentOrderNo
	if qPaymentOrderNo != "" {
		if err := r.SetQueryParam("paymentOrderNo", qPaymentOrderNo); err != nil {
			return err
		}
	}

	// query param paymentProvider
	qrPaymentProvider := o.PaymentProvider
	qPaymentProvider := qrPaymentProvider
	if qPaymentProvider != "" {
		if err := r.SetQueryParam("paymentProvider", qPaymentProvider); err != nil {
			return err
		}
	}

	if o.ResultCode != nil {

		// query param resultCode
		var qrResultCode string
		if o.ResultCode != nil {
			qrResultCode = *o.ResultCode
		}
		qResultCode := qrResultCode
		if qResultCode != "" {
			if err := r.SetQueryParam("resultCode", qResultCode); err != nil {
				return err
			}
		}

	}

	// query param returnUrl
	qrReturnURL := o.ReturnURL
	qReturnURL := qrReturnURL
	if qReturnURL != "" {
		if err := r.SetQueryParam("returnUrl", qReturnURL); err != nil {
			return err
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.Token != nil {

		// query param token
		var qrToken string
		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken
		if qToken != "" {
			if err := r.SetQueryParam("token", qToken); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID string
		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {
			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}