// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package users

import (
	"context"
	"net/http"
	"time"

	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
)

// NewPublicGetUsersPlatformInfosV3Params creates a new PublicGetUsersPlatformInfosV3Params object
// with the default values initialized.
func NewPublicGetUsersPlatformInfosV3Params() *PublicGetUsersPlatformInfosV3Params {
	var ()
	return &PublicGetUsersPlatformInfosV3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicGetUsersPlatformInfosV3ParamsWithTimeout creates a new PublicGetUsersPlatformInfosV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicGetUsersPlatformInfosV3ParamsWithTimeout(timeout time.Duration) *PublicGetUsersPlatformInfosV3Params {
	var ()
	return &PublicGetUsersPlatformInfosV3Params{

		timeout: timeout,
	}
}

// NewPublicGetUsersPlatformInfosV3ParamsWithContext creates a new PublicGetUsersPlatformInfosV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewPublicGetUsersPlatformInfosV3ParamsWithContext(ctx context.Context) *PublicGetUsersPlatformInfosV3Params {
	var ()
	return &PublicGetUsersPlatformInfosV3Params{

		Context: ctx,
	}
}

// NewPublicGetUsersPlatformInfosV3ParamsWithHTTPClient creates a new PublicGetUsersPlatformInfosV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicGetUsersPlatformInfosV3ParamsWithHTTPClient(client *http.Client) *PublicGetUsersPlatformInfosV3Params {
	var ()
	return &PublicGetUsersPlatformInfosV3Params{
		HTTPClient: client,
	}
}

/*PublicGetUsersPlatformInfosV3Params contains all the parameters to send to the API endpoint
for the public get users platform infos v3 operation typically these are written to a http.Request
*/
type PublicGetUsersPlatformInfosV3Params struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Body*/
	Body *iamclientmodels.ModelUsersPlatformInfosRequestV3
	/*Namespace
	  Namespace, only accept alphabet and numeric

	*/
	Namespace string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the public get users platform infos v3 params
func (o *PublicGetUsersPlatformInfosV3Params) WithTimeout(timeout time.Duration) *PublicGetUsersPlatformInfosV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public get users platform infos v3 params
func (o *PublicGetUsersPlatformInfosV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public get users platform infos v3 params
func (o *PublicGetUsersPlatformInfosV3Params) WithContext(ctx context.Context) *PublicGetUsersPlatformInfosV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public get users platform infos v3 params
func (o *PublicGetUsersPlatformInfosV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the public get users platform infos v3 params
func (o *PublicGetUsersPlatformInfosV3Params) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the public get users platform infos v3 params
func (o *PublicGetUsersPlatformInfosV3Params) WithHTTPClient(client *http.Client) *PublicGetUsersPlatformInfosV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public get users platform infos v3 params
func (o *PublicGetUsersPlatformInfosV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the public get users platform infos v3 params
func (o *PublicGetUsersPlatformInfosV3Params) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// WithBody adds the body to the public get users platform infos v3 params
func (o *PublicGetUsersPlatformInfosV3Params) WithBody(body *iamclientmodels.ModelUsersPlatformInfosRequestV3) *PublicGetUsersPlatformInfosV3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the public get users platform infos v3 params
func (o *PublicGetUsersPlatformInfosV3Params) SetBody(body *iamclientmodels.ModelUsersPlatformInfosRequestV3) {
	o.Body = body
}

// WithNamespace adds the namespace to the public get users platform infos v3 params
func (o *PublicGetUsersPlatformInfosV3Params) WithNamespace(namespace string) *PublicGetUsersPlatformInfosV3Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public get users platform infos v3 params
func (o *PublicGetUsersPlatformInfosV3Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *PublicGetUsersPlatformInfosV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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