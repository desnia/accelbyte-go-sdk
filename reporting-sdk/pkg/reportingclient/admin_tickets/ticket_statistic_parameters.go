// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package admin_tickets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewTicketStatisticParams creates a new TicketStatisticParams object
// with the default values initialized.
func NewTicketStatisticParams() *TicketStatisticParams {
	var ()
	return &TicketStatisticParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTicketStatisticParamsWithTimeout creates a new TicketStatisticParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTicketStatisticParamsWithTimeout(timeout time.Duration) *TicketStatisticParams {
	var ()
	return &TicketStatisticParams{

		timeout: timeout,
	}
}

// NewTicketStatisticParamsWithContext creates a new TicketStatisticParams object
// with the default values initialized, and the ability to set a context for a request
func NewTicketStatisticParamsWithContext(ctx context.Context) *TicketStatisticParams {
	var ()
	return &TicketStatisticParams{

		Context: ctx,
	}
}

// NewTicketStatisticParamsWithHTTPClient creates a new TicketStatisticParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTicketStatisticParamsWithHTTPClient(client *http.Client) *TicketStatisticParams {
	var ()
	return &TicketStatisticParams{
		HTTPClient: client,
	}
}

/*TicketStatisticParams contains all the parameters to send to the API endpoint
for the ticket statistic operation typically these are written to a http.Request
*/
type TicketStatisticParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Category
	  filter statistic by category

	*/
	Category string
	/*ExtensionCategory
	  filter statistic by extension category

	*/
	ExtensionCategory *string
	/*Namespace*/
	Namespace string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the ticket statistic params
func (o *TicketStatisticParams) WithTimeout(timeout time.Duration) *TicketStatisticParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ticket statistic params
func (o *TicketStatisticParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ticket statistic params
func (o *TicketStatisticParams) WithContext(ctx context.Context) *TicketStatisticParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ticket statistic params
func (o *TicketStatisticParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the ticket statistic params
func (o *TicketStatisticParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the ticket statistic params
func (o *TicketStatisticParams) WithHTTPClient(client *http.Client) *TicketStatisticParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ticket statistic params
func (o *TicketStatisticParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the ticket statistic params
func (o *TicketStatisticParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// WithCategory adds the category to the ticket statistic params
func (o *TicketStatisticParams) WithCategory(category string) *TicketStatisticParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the ticket statistic params
func (o *TicketStatisticParams) SetCategory(category string) {
	o.Category = category
}

// WithExtensionCategory adds the extensionCategory to the ticket statistic params
func (o *TicketStatisticParams) WithExtensionCategory(extensionCategory *string) *TicketStatisticParams {
	o.SetExtensionCategory(extensionCategory)
	return o
}

// SetExtensionCategory adds the extensionCategory to the ticket statistic params
func (o *TicketStatisticParams) SetExtensionCategory(extensionCategory *string) {
	o.ExtensionCategory = extensionCategory
}

// WithNamespace adds the namespace to the ticket statistic params
func (o *TicketStatisticParams) WithNamespace(namespace string) *TicketStatisticParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the ticket statistic params
func (o *TicketStatisticParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *TicketStatisticParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param category
	qrCategory := o.Category
	qCategory := qrCategory
	if qCategory != "" {
		if err := r.SetQueryParam("category", qCategory); err != nil {
			return err
		}
	}

	if o.ExtensionCategory != nil {

		// query param extensionCategory
		var qrExtensionCategory string
		if o.ExtensionCategory != nil {
			qrExtensionCategory = *o.ExtensionCategory
		}
		qExtensionCategory := qrExtensionCategory
		if qExtensionCategory != "" {
			if err := r.SetQueryParam("extensionCategory", qExtensionCategory); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// setting the default header value
	if err := r.SetHeaderParam("User-Agent", utils.UserAgentGen()); err != nil {
		return err
	}

	if err := r.SetHeaderParam("X-Amzn-Trace-Id", utils.AmazonTraceIDGen()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
