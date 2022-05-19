// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package wallet

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

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// NewUpdatePlatformWalletConfigParams creates a new UpdatePlatformWalletConfigParams object
// with the default values initialized.
func NewUpdatePlatformWalletConfigParams() *UpdatePlatformWalletConfigParams {
	var ()
	return &UpdatePlatformWalletConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePlatformWalletConfigParamsWithTimeout creates a new UpdatePlatformWalletConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePlatformWalletConfigParamsWithTimeout(timeout time.Duration) *UpdatePlatformWalletConfigParams {
	var ()
	return &UpdatePlatformWalletConfigParams{

		timeout: timeout,
	}
}

// NewUpdatePlatformWalletConfigParamsWithContext creates a new UpdatePlatformWalletConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePlatformWalletConfigParamsWithContext(ctx context.Context) *UpdatePlatformWalletConfigParams {
	var ()
	return &UpdatePlatformWalletConfigParams{

		Context: ctx,
	}
}

// NewUpdatePlatformWalletConfigParamsWithHTTPClient creates a new UpdatePlatformWalletConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePlatformWalletConfigParamsWithHTTPClient(client *http.Client) *UpdatePlatformWalletConfigParams {
	var ()
	return &UpdatePlatformWalletConfigParams{
		HTTPClient: client,
	}
}

/*UpdatePlatformWalletConfigParams contains all the parameters to send to the API endpoint
for the update platform wallet config operation typically these are written to a http.Request
*/
type UpdatePlatformWalletConfigParams struct {

	/*Body*/
	Body *platformclientmodels.PlatformWalletConfigUpdate
	/*Namespace
	  Namespace

	*/
	Namespace string
	/*Platform*/
	Platform string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the update platform wallet config params
func (o *UpdatePlatformWalletConfigParams) WithTimeout(timeout time.Duration) *UpdatePlatformWalletConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update platform wallet config params
func (o *UpdatePlatformWalletConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update platform wallet config params
func (o *UpdatePlatformWalletConfigParams) WithContext(ctx context.Context) *UpdatePlatformWalletConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update platform wallet config params
func (o *UpdatePlatformWalletConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the update platform wallet config params
func (o *UpdatePlatformWalletConfigParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the update platform wallet config params
func (o *UpdatePlatformWalletConfigParams) WithHTTPClient(client *http.Client) *UpdatePlatformWalletConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update platform wallet config params
func (o *UpdatePlatformWalletConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update platform wallet config params
func (o *UpdatePlatformWalletConfigParams) WithBody(body *platformclientmodels.PlatformWalletConfigUpdate) *UpdatePlatformWalletConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update platform wallet config params
func (o *UpdatePlatformWalletConfigParams) SetBody(body *platformclientmodels.PlatformWalletConfigUpdate) {
	o.Body = body
}

// WithNamespace adds the namespace to the update platform wallet config params
func (o *UpdatePlatformWalletConfigParams) WithNamespace(namespace string) *UpdatePlatformWalletConfigParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update platform wallet config params
func (o *UpdatePlatformWalletConfigParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPlatform adds the platform to the update platform wallet config params
func (o *UpdatePlatformWalletConfigParams) WithPlatform(platform string) *UpdatePlatformWalletConfigParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the update platform wallet config params
func (o *UpdatePlatformWalletConfigParams) SetPlatform(platform string) {
	o.Platform = platform
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePlatformWalletConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param platform
	if err := r.SetPathParam("platform", o.Platform); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}