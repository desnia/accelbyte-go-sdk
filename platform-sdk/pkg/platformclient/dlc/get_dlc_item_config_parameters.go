// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package dlc

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

// NewGetDLCItemConfigParams creates a new GetDLCItemConfigParams object
// with the default values initialized.
func NewGetDLCItemConfigParams() *GetDLCItemConfigParams {
	var ()
	return &GetDLCItemConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDLCItemConfigParamsWithTimeout creates a new GetDLCItemConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDLCItemConfigParamsWithTimeout(timeout time.Duration) *GetDLCItemConfigParams {
	var ()
	return &GetDLCItemConfigParams{

		timeout: timeout,
	}
}

// NewGetDLCItemConfigParamsWithContext creates a new GetDLCItemConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDLCItemConfigParamsWithContext(ctx context.Context) *GetDLCItemConfigParams {
	var ()
	return &GetDLCItemConfigParams{

		Context: ctx,
	}
}

// NewGetDLCItemConfigParamsWithHTTPClient creates a new GetDLCItemConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDLCItemConfigParamsWithHTTPClient(client *http.Client) *GetDLCItemConfigParams {
	var ()
	return &GetDLCItemConfigParams{
		HTTPClient: client,
	}
}

/*
GetDLCItemConfigParams contains all the parameters to send to the API endpoint
for the get DLC item config operation typically these are written to a http.Request
*/
type GetDLCItemConfigParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Namespace*/
	Namespace string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the get DLC item config params
func (o *GetDLCItemConfigParams) WithTimeout(timeout time.Duration) *GetDLCItemConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get DLC item config params
func (o *GetDLCItemConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get DLC item config params
func (o *GetDLCItemConfigParams) WithContext(ctx context.Context) *GetDLCItemConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get DLC item config params
func (o *GetDLCItemConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the get DLC item config params
func (o *GetDLCItemConfigParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the get DLC item config params
func (o *GetDLCItemConfigParams) WithHTTPClient(client *http.Client) *GetDLCItemConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get DLC item config params
func (o *GetDLCItemConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the get DLC item config params
func (o *GetDLCItemConfigParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// WithNamespace adds the namespace to the get DLC item config params
func (o *GetDLCItemConfigParams) WithNamespace(namespace string) *GetDLCItemConfigParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get DLC item config params
func (o *GetDLCItemConfigParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetDLCItemConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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