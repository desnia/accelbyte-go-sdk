// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package utility

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

// NewCheckReadinessParams creates a new CheckReadinessParams object
// with the default values initialized.
func NewCheckReadinessParams() *CheckReadinessParams {

	return &CheckReadinessParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckReadinessParamsWithTimeout creates a new CheckReadinessParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckReadinessParamsWithTimeout(timeout time.Duration) *CheckReadinessParams {

	return &CheckReadinessParams{

		timeout: timeout,
	}
}

// NewCheckReadinessParamsWithContext creates a new CheckReadinessParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckReadinessParamsWithContext(ctx context.Context) *CheckReadinessParams {

	return &CheckReadinessParams{

		Context: ctx,
	}
}

// NewCheckReadinessParamsWithHTTPClient creates a new CheckReadinessParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckReadinessParamsWithHTTPClient(client *http.Client) *CheckReadinessParams {

	return &CheckReadinessParams{
		HTTPClient: client,
	}
}

/*CheckReadinessParams contains all the parameters to send to the API endpoint
for the check readiness operation typically these are written to a http.Request
*/
type CheckReadinessParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the check readiness params
func (o *CheckReadinessParams) WithTimeout(timeout time.Duration) *CheckReadinessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check readiness params
func (o *CheckReadinessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check readiness params
func (o *CheckReadinessParams) WithContext(ctx context.Context) *CheckReadinessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check readiness params
func (o *CheckReadinessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the check readiness params
func (o *CheckReadinessParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the check readiness params
func (o *CheckReadinessParams) WithHTTPClient(client *http.Client) *CheckReadinessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check readiness params
func (o *CheckReadinessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CheckReadinessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
