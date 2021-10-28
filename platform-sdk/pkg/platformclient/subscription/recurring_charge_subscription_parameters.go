// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// NewRecurringChargeSubscriptionParams creates a new RecurringChargeSubscriptionParams object
// with the default values initialized.
func NewRecurringChargeSubscriptionParams() *RecurringChargeSubscriptionParams {
	var ()
	return &RecurringChargeSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRecurringChargeSubscriptionParamsWithTimeout creates a new RecurringChargeSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRecurringChargeSubscriptionParamsWithTimeout(timeout time.Duration) *RecurringChargeSubscriptionParams {
	var ()
	return &RecurringChargeSubscriptionParams{

		timeout: timeout,
	}
}

// NewRecurringChargeSubscriptionParamsWithContext creates a new RecurringChargeSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewRecurringChargeSubscriptionParamsWithContext(ctx context.Context) *RecurringChargeSubscriptionParams {
	var ()
	return &RecurringChargeSubscriptionParams{

		Context: ctx,
	}
}

// NewRecurringChargeSubscriptionParamsWithHTTPClient creates a new RecurringChargeSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRecurringChargeSubscriptionParamsWithHTTPClient(client *http.Client) *RecurringChargeSubscriptionParams {
	var ()
	return &RecurringChargeSubscriptionParams{
		HTTPClient: client,
	}
}

/*RecurringChargeSubscriptionParams contains all the parameters to send to the API endpoint
for the recurring charge subscription operation typically these are written to a http.Request
*/
type RecurringChargeSubscriptionParams struct {

	/*Namespace*/
	Namespace string
	/*SubscriptionID*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the recurring charge subscription params
func (o *RecurringChargeSubscriptionParams) WithTimeout(timeout time.Duration) *RecurringChargeSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the recurring charge subscription params
func (o *RecurringChargeSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the recurring charge subscription params
func (o *RecurringChargeSubscriptionParams) WithContext(ctx context.Context) *RecurringChargeSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the recurring charge subscription params
func (o *RecurringChargeSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the recurring charge subscription params
func (o *RecurringChargeSubscriptionParams) WithHTTPClient(client *http.Client) *RecurringChargeSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the recurring charge subscription params
func (o *RecurringChargeSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the recurring charge subscription params
func (o *RecurringChargeSubscriptionParams) WithNamespace(namespace string) *RecurringChargeSubscriptionParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the recurring charge subscription params
func (o *RecurringChargeSubscriptionParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithSubscriptionID adds the subscriptionID to the recurring charge subscription params
func (o *RecurringChargeSubscriptionParams) WithSubscriptionID(subscriptionID string) *RecurringChargeSubscriptionParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the recurring charge subscription params
func (o *RecurringChargeSubscriptionParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *RecurringChargeSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}