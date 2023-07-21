// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package platform_credential

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

// NewAdminDeletePlatformCredentialsParams creates a new AdminDeletePlatformCredentialsParams object
// with the default values initialized.
func NewAdminDeletePlatformCredentialsParams() *AdminDeletePlatformCredentialsParams {
	var ()
	return &AdminDeletePlatformCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminDeletePlatformCredentialsParamsWithTimeout creates a new AdminDeletePlatformCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminDeletePlatformCredentialsParamsWithTimeout(timeout time.Duration) *AdminDeletePlatformCredentialsParams {
	var ()
	return &AdminDeletePlatformCredentialsParams{

		timeout: timeout,
	}
}

// NewAdminDeletePlatformCredentialsParamsWithContext creates a new AdminDeletePlatformCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminDeletePlatformCredentialsParamsWithContext(ctx context.Context) *AdminDeletePlatformCredentialsParams {
	var ()
	return &AdminDeletePlatformCredentialsParams{

		Context: ctx,
	}
}

// NewAdminDeletePlatformCredentialsParamsWithHTTPClient creates a new AdminDeletePlatformCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminDeletePlatformCredentialsParamsWithHTTPClient(client *http.Client) *AdminDeletePlatformCredentialsParams {
	var ()
	return &AdminDeletePlatformCredentialsParams{
		HTTPClient: client,
	}
}

/*AdminDeletePlatformCredentialsParams contains all the parameters to send to the API endpoint
for the admin delete platform credentials operation typically these are written to a http.Request
*/
type AdminDeletePlatformCredentialsParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Namespace
	  namespace of the game

	*/
	Namespace string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the admin delete platform credentials params
func (o *AdminDeletePlatformCredentialsParams) WithTimeout(timeout time.Duration) *AdminDeletePlatformCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin delete platform credentials params
func (o *AdminDeletePlatformCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin delete platform credentials params
func (o *AdminDeletePlatformCredentialsParams) WithContext(ctx context.Context) *AdminDeletePlatformCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin delete platform credentials params
func (o *AdminDeletePlatformCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the admin delete platform credentials params
func (o *AdminDeletePlatformCredentialsParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the admin delete platform credentials params
func (o *AdminDeletePlatformCredentialsParams) WithHTTPClient(client *http.Client) *AdminDeletePlatformCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin delete platform credentials params
func (o *AdminDeletePlatformCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the admin delete platform credentials params
func (o *AdminDeletePlatformCredentialsParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// WithNamespace adds the namespace to the admin delete platform credentials params
func (o *AdminDeletePlatformCredentialsParams) WithNamespace(namespace string) *AdminDeletePlatformCredentialsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin delete platform credentials params
func (o *AdminDeletePlatformCredentialsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *AdminDeletePlatformCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
