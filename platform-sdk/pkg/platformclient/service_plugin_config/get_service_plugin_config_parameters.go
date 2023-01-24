// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package service_plugin_config

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

// NewGetServicePluginConfigParams creates a new GetServicePluginConfigParams object
// with the default values initialized.
func NewGetServicePluginConfigParams() *GetServicePluginConfigParams {
	var ()
	return &GetServicePluginConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServicePluginConfigParamsWithTimeout creates a new GetServicePluginConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServicePluginConfigParamsWithTimeout(timeout time.Duration) *GetServicePluginConfigParams {
	var ()
	return &GetServicePluginConfigParams{

		timeout: timeout,
	}
}

// NewGetServicePluginConfigParamsWithContext creates a new GetServicePluginConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServicePluginConfigParamsWithContext(ctx context.Context) *GetServicePluginConfigParams {
	var ()
	return &GetServicePluginConfigParams{

		Context: ctx,
	}
}

// NewGetServicePluginConfigParamsWithHTTPClient creates a new GetServicePluginConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServicePluginConfigParamsWithHTTPClient(client *http.Client) *GetServicePluginConfigParams {
	var ()
	return &GetServicePluginConfigParams{
		HTTPClient: client,
	}
}

/*GetServicePluginConfigParams contains all the parameters to send to the API endpoint
for the get service plugin config operation typically these are written to a http.Request
*/
type GetServicePluginConfigParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Namespace*/
	Namespace string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the get service plugin config params
func (o *GetServicePluginConfigParams) WithTimeout(timeout time.Duration) *GetServicePluginConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service plugin config params
func (o *GetServicePluginConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service plugin config params
func (o *GetServicePluginConfigParams) WithContext(ctx context.Context) *GetServicePluginConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service plugin config params
func (o *GetServicePluginConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the get service plugin config params
func (o *GetServicePluginConfigParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the get service plugin config params
func (o *GetServicePluginConfigParams) WithHTTPClient(client *http.Client) *GetServicePluginConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service plugin config params
func (o *GetServicePluginConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the get service plugin config params
func (o *GetServicePluginConfigParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// WithNamespace adds the namespace to the get service plugin config params
func (o *GetServicePluginConfigParams) WithNamespace(namespace string) *GetServicePluginConfigParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get service plugin config params
func (o *GetServicePluginConfigParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetServicePluginConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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