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

// NewProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams creates a new ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams object
// with the default values initialized.
func NewProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams() *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams {
	var ()
	return &ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParamsWithTimeout creates a new ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParamsWithTimeout(timeout time.Duration) *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams {
	var ()
	return &ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams{

		timeout: timeout,
	}
}

// NewProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParamsWithContext creates a new ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParamsWithContext(ctx context.Context) *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams {
	var ()
	return &ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams{

		Context: ctx,
	}
}

// NewProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParamsWithHTTPClient creates a new ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParamsWithHTTPClient(client *http.Client) *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams {
	var ()
	return &ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams{
		HTTPClient: client,
	}
}

/*ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams contains all the parameters to send to the API endpoint
for the protected get playtime game telemetry v1 protected steam ids steam Id playtime get operation typically these are written to a http.Request
*/
type ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams struct {

	/*Cookie*/
	Cookie *string
	/*Retry*/
	Retry       bool
	RetryPolicy *utils.Retry
	/*SteamID*/
	SteamID string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the protected get playtime game telemetry v1 protected steam ids steam Id playtime get params
func (o *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams) WithTimeout(timeout time.Duration) *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the protected get playtime game telemetry v1 protected steam ids steam Id playtime get params
func (o *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the protected get playtime game telemetry v1 protected steam ids steam Id playtime get params
func (o *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams) WithContext(ctx context.Context) *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the protected get playtime game telemetry v1 protected steam ids steam Id playtime get params
func (o *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the protected get playtime game telemetry v1 protected steam ids steam Id playtime get params
func (o *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the protected get playtime game telemetry v1 protected steam ids steam Id playtime get params
func (o *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams) WithHTTPClient(client *http.Client) *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the protected get playtime game telemetry v1 protected steam ids steam Id playtime get params
func (o *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCookie adds the cookie to the protected get playtime game telemetry v1 protected steam ids steam Id playtime get params
func (o *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams) WithCookie(cookie *string) *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams {
	o.SetCookie(cookie)
	return o
}

// SetCookie adds the cookie to the protected get playtime game telemetry v1 protected steam ids steam Id playtime get params
func (o *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams) SetCookie(cookie *string) {
	o.Cookie = cookie
}

// WithRetry
func (o *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams) WithRetry(retry bool) *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams {
	o.SetRetry(retry)
	return o
}

// SetRetry
func (o *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams) SetRetry(retry bool) {
	o.Retry = retry
}

// WithSteamID adds the steamID to the protected get playtime game telemetry v1 protected steam ids steam Id playtime get params
func (o *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams) WithSteamID(steamID string) *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams {
	o.SetSteamID(steamID)
	return o
}

// SetSteamID adds the steamId to the protected get playtime game telemetry v1 protected steam ids steam Id playtime get params
func (o *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams) SetSteamID(steamID string) {
	o.SteamID = steamID
}

// WriteToRequest writes these params to a swagger request
func (o *ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Retry != true {
		// logic retry
		utils.Retry{
			Enabled:    false,
			RetryCodes: nil,
			MaxTries:   0,
			Backoff:    nil,
			Transport:  nil,
			Verbose:    false,
			Sleeper:    nil,
		}
	}

	// path param steamId
	if err := r.SetPathParam("steamId", o.SteamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
