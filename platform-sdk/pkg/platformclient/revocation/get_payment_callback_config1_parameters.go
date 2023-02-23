// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package revocation

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

// NewGetPaymentCallbackConfig1Params creates a new GetPaymentCallbackConfig1Params object
// with the default values initialized.
func NewGetPaymentCallbackConfig1Params() *GetPaymentCallbackConfig1Params {
	var ()
	return &GetPaymentCallbackConfig1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentCallbackConfig1ParamsWithTimeout creates a new GetPaymentCallbackConfig1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentCallbackConfig1ParamsWithTimeout(timeout time.Duration) *GetPaymentCallbackConfig1Params {
	var ()
	return &GetPaymentCallbackConfig1Params{

		timeout: timeout,
	}
}

// NewGetPaymentCallbackConfig1ParamsWithContext creates a new GetPaymentCallbackConfig1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentCallbackConfig1ParamsWithContext(ctx context.Context) *GetPaymentCallbackConfig1Params {
	var ()
	return &GetPaymentCallbackConfig1Params{

		Context: ctx,
	}
}

// NewGetPaymentCallbackConfig1ParamsWithHTTPClient creates a new GetPaymentCallbackConfig1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentCallbackConfig1ParamsWithHTTPClient(client *http.Client) *GetPaymentCallbackConfig1Params {
	var ()
	return &GetPaymentCallbackConfig1Params{
		HTTPClient: client,
	}
}

/*GetPaymentCallbackConfig1Params contains all the parameters to send to the API endpoint
for the get payment callback config 1 operation typically these are written to a http.Request
*/
type GetPaymentCallbackConfig1Params struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Namespace*/
	Namespace string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the get payment callback config 1 params
func (o *GetPaymentCallbackConfig1Params) WithTimeout(timeout time.Duration) *GetPaymentCallbackConfig1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment callback config 1 params
func (o *GetPaymentCallbackConfig1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment callback config 1 params
func (o *GetPaymentCallbackConfig1Params) WithContext(ctx context.Context) *GetPaymentCallbackConfig1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment callback config 1 params
func (o *GetPaymentCallbackConfig1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the get payment callback config 1 params
func (o *GetPaymentCallbackConfig1Params) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the get payment callback config 1 params
func (o *GetPaymentCallbackConfig1Params) WithHTTPClient(client *http.Client) *GetPaymentCallbackConfig1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment callback config 1 params
func (o *GetPaymentCallbackConfig1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the get payment callback config 1 params
func (o *GetPaymentCallbackConfig1Params) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// WithNamespace adds the namespace to the get payment callback config 1 params
func (o *GetPaymentCallbackConfig1Params) WithNamespace(namespace string) *GetPaymentCallbackConfig1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get payment callback config 1 params
func (o *GetPaymentCallbackConfig1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentCallbackConfig1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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