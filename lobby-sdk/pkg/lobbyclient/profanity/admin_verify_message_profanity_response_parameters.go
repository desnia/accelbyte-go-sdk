// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package profanity

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

	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclientmodels"
)

// NewAdminVerifyMessageProfanityResponseParams creates a new AdminVerifyMessageProfanityResponseParams object
// with the default values initialized.
func NewAdminVerifyMessageProfanityResponseParams() *AdminVerifyMessageProfanityResponseParams {
	var ()
	return &AdminVerifyMessageProfanityResponseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminVerifyMessageProfanityResponseParamsWithTimeout creates a new AdminVerifyMessageProfanityResponseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminVerifyMessageProfanityResponseParamsWithTimeout(timeout time.Duration) *AdminVerifyMessageProfanityResponseParams {
	var ()
	return &AdminVerifyMessageProfanityResponseParams{

		timeout: timeout,
	}
}

// NewAdminVerifyMessageProfanityResponseParamsWithContext creates a new AdminVerifyMessageProfanityResponseParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminVerifyMessageProfanityResponseParamsWithContext(ctx context.Context) *AdminVerifyMessageProfanityResponseParams {
	var ()
	return &AdminVerifyMessageProfanityResponseParams{

		Context: ctx,
	}
}

// NewAdminVerifyMessageProfanityResponseParamsWithHTTPClient creates a new AdminVerifyMessageProfanityResponseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminVerifyMessageProfanityResponseParamsWithHTTPClient(client *http.Client) *AdminVerifyMessageProfanityResponseParams {
	var ()
	return &AdminVerifyMessageProfanityResponseParams{
		HTTPClient: client,
	}
}

/*AdminVerifyMessageProfanityResponseParams contains all the parameters to send to the API endpoint
for the admin verify message profanity response operation typically these are written to a http.Request
*/
type AdminVerifyMessageProfanityResponseParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Body
	  request

	*/
	Body *lobbyclientmodels.ModelsAdminVerifyMessageProfanityRequest
	/*Namespace
	  namespace

	*/
	Namespace string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the admin verify message profanity response params
func (o *AdminVerifyMessageProfanityResponseParams) WithTimeout(timeout time.Duration) *AdminVerifyMessageProfanityResponseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin verify message profanity response params
func (o *AdminVerifyMessageProfanityResponseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin verify message profanity response params
func (o *AdminVerifyMessageProfanityResponseParams) WithContext(ctx context.Context) *AdminVerifyMessageProfanityResponseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin verify message profanity response params
func (o *AdminVerifyMessageProfanityResponseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the admin verify message profanity response params
func (o *AdminVerifyMessageProfanityResponseParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the admin verify message profanity response params
func (o *AdminVerifyMessageProfanityResponseParams) WithHTTPClient(client *http.Client) *AdminVerifyMessageProfanityResponseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin verify message profanity response params
func (o *AdminVerifyMessageProfanityResponseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin verify message profanity response params
func (o *AdminVerifyMessageProfanityResponseParams) WithBody(body *lobbyclientmodels.ModelsAdminVerifyMessageProfanityRequest) *AdminVerifyMessageProfanityResponseParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin verify message profanity response params
func (o *AdminVerifyMessageProfanityResponseParams) SetBody(body *lobbyclientmodels.ModelsAdminVerifyMessageProfanityRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the admin verify message profanity response params
func (o *AdminVerifyMessageProfanityResponseParams) WithNamespace(namespace string) *AdminVerifyMessageProfanityResponseParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin verify message profanity response params
func (o *AdminVerifyMessageProfanityResponseParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *AdminVerifyMessageProfanityResponseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
