// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package gametelemetry_operations

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

// NewAdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams creates a new AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams object
// with the default values initialized.
func NewAdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams() *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams {
	var ()
	return &AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParamsWithTimeout creates a new AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParamsWithTimeout(timeout time.Duration) *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams {
	var ()
	return &AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams{

		timeout: timeout,
	}
}

// NewAdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParamsWithContext creates a new AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParamsWithContext(ctx context.Context) *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams {
	var ()
	return &AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams{

		Context: ctx,
	}
}

// NewAdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParamsWithHTTPClient creates a new AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParamsWithHTTPClient(client *http.Client) *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams {
	var ()
	return &AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams{
		HTTPClient: client,
	}
}

/*AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams contains all the parameters to send to the API endpoint
for the admin get namespace game telemetry v1 admin telemetrynamespace get operation typically these are written to a http.Request
*/
type AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Cookie*/
	Cookie *string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the admin get namespace game telemetry v1 admin telemetrynamespace get params
func (o *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams) WithTimeout(timeout time.Duration) *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get namespace game telemetry v1 admin telemetrynamespace get params
func (o *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get namespace game telemetry v1 admin telemetrynamespace get params
func (o *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams) WithContext(ctx context.Context) *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get namespace game telemetry v1 admin telemetrynamespace get params
func (o *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the admin get namespace game telemetry v1 admin telemetrynamespace get params
func (o *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the admin get namespace game telemetry v1 admin telemetrynamespace get params
func (o *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams) WithHTTPClient(client *http.Client) *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get namespace game telemetry v1 admin telemetrynamespace get params
func (o *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCookie adds the cookie to the admin get namespace game telemetry v1 admin telemetrynamespace get params
func (o *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams) WithCookie(cookie *string) *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams {
	o.SetCookie(cookie)
	return o
}

// SetCookie adds the cookie to the admin get namespace game telemetry v1 admin telemetrynamespace get params
func (o *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams) SetCookie(cookie *string) {
	o.Cookie = cookie
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cookie != nil {

		// header param Cookie
		if err := r.SetHeaderParam("Cookie", *o.Cookie); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
