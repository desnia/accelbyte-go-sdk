// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package bans

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

// NewAdminGetBansTypeV3Params creates a new AdminGetBansTypeV3Params object
// with the default values initialized.
func NewAdminGetBansTypeV3Params() *AdminGetBansTypeV3Params {

	return &AdminGetBansTypeV3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetBansTypeV3ParamsWithTimeout creates a new AdminGetBansTypeV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminGetBansTypeV3ParamsWithTimeout(timeout time.Duration) *AdminGetBansTypeV3Params {

	return &AdminGetBansTypeV3Params{

		timeout: timeout,
	}
}

// NewAdminGetBansTypeV3ParamsWithContext creates a new AdminGetBansTypeV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminGetBansTypeV3ParamsWithContext(ctx context.Context) *AdminGetBansTypeV3Params {

	return &AdminGetBansTypeV3Params{

		Context: ctx,
	}
}

// NewAdminGetBansTypeV3ParamsWithHTTPClient creates a new AdminGetBansTypeV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminGetBansTypeV3ParamsWithHTTPClient(client *http.Client) *AdminGetBansTypeV3Params {

	return &AdminGetBansTypeV3Params{
		HTTPClient: client,
	}
}

/*AdminGetBansTypeV3Params contains all the parameters to send to the API endpoint
for the admin get bans type v3 operation typically these are written to a http.Request
*/
type AdminGetBansTypeV3Params struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the admin get bans type v3 params
func (o *AdminGetBansTypeV3Params) WithTimeout(timeout time.Duration) *AdminGetBansTypeV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get bans type v3 params
func (o *AdminGetBansTypeV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get bans type v3 params
func (o *AdminGetBansTypeV3Params) WithContext(ctx context.Context) *AdminGetBansTypeV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get bans type v3 params
func (o *AdminGetBansTypeV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the admin get bans type v3 params
func (o *AdminGetBansTypeV3Params) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the admin get bans type v3 params
func (o *AdminGetBansTypeV3Params) WithHTTPClient(client *http.Client) *AdminGetBansTypeV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get bans type v3 params
func (o *AdminGetBansTypeV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetBansTypeV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
