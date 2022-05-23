// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package event_registry

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

// NewGetRegisteredEventsHandlerParams creates a new GetRegisteredEventsHandlerParams object
// with the default values initialized.
func NewGetRegisteredEventsHandlerParams() *GetRegisteredEventsHandlerParams {

	return &GetRegisteredEventsHandlerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegisteredEventsHandlerParamsWithTimeout creates a new GetRegisteredEventsHandlerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRegisteredEventsHandlerParamsWithTimeout(timeout time.Duration) *GetRegisteredEventsHandlerParams {

	return &GetRegisteredEventsHandlerParams{

		timeout: timeout,
	}
}

// NewGetRegisteredEventsHandlerParamsWithContext creates a new GetRegisteredEventsHandlerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRegisteredEventsHandlerParamsWithContext(ctx context.Context) *GetRegisteredEventsHandlerParams {

	return &GetRegisteredEventsHandlerParams{

		Context: ctx,
	}
}

// NewGetRegisteredEventsHandlerParamsWithHTTPClient creates a new GetRegisteredEventsHandlerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRegisteredEventsHandlerParamsWithHTTPClient(client *http.Client) *GetRegisteredEventsHandlerParams {

	return &GetRegisteredEventsHandlerParams{
		HTTPClient: client,
	}
}

/*GetRegisteredEventsHandlerParams contains all the parameters to send to the API endpoint
for the get registered events handler operation typically these are written to a http.Request
*/
type GetRegisteredEventsHandlerParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the get registered events handler params
func (o *GetRegisteredEventsHandlerParams) WithTimeout(timeout time.Duration) *GetRegisteredEventsHandlerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get registered events handler params
func (o *GetRegisteredEventsHandlerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get registered events handler params
func (o *GetRegisteredEventsHandlerParams) WithContext(ctx context.Context) *GetRegisteredEventsHandlerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get registered events handler params
func (o *GetRegisteredEventsHandlerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the get registered events handler params
func (o *GetRegisteredEventsHandlerParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the get registered events handler params
func (o *GetRegisteredEventsHandlerParams) WithHTTPClient(client *http.Client) *GetRegisteredEventsHandlerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get registered events handler params
func (o *GetRegisteredEventsHandlerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegisteredEventsHandlerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
