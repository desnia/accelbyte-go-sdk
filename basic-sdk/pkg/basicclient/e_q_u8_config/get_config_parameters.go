// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package e_q_u8_config

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

// NewGetConfigParams creates a new GetConfigParams object
// with the default values initialized.
func NewGetConfigParams() *GetConfigParams {
	var ()
	return &GetConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfigParamsWithTimeout creates a new GetConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConfigParamsWithTimeout(timeout time.Duration) *GetConfigParams {
	var ()
	return &GetConfigParams{

		timeout: timeout,
	}
}

// NewGetConfigParamsWithContext creates a new GetConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConfigParamsWithContext(ctx context.Context) *GetConfigParams {
	var ()
	return &GetConfigParams{

		Context: ctx,
	}
}

// NewGetConfigParamsWithHTTPClient creates a new GetConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConfigParamsWithHTTPClient(client *http.Client) *GetConfigParams {
	var ()
	return &GetConfigParams{
		HTTPClient: client,
	}
}

/*GetConfigParams contains all the parameters to send to the API endpoint
for the get config operation typically these are written to a http.Request
*/
type GetConfigParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Namespace
	  namespace, only accept alphabet and numeric

	*/
	Namespace string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the get config params
func (o *GetConfigParams) WithTimeout(timeout time.Duration) *GetConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get config params
func (o *GetConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get config params
func (o *GetConfigParams) WithContext(ctx context.Context) *GetConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get config params
func (o *GetConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the get config params
func (o *GetConfigParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the get config params
func (o *GetConfigParams) WithHTTPClient(client *http.Client) *GetConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get config params
func (o *GetConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get config params
func (o *GetConfigParams) WithNamespace(namespace string) *GetConfigParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get config params
func (o *GetConfigParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
