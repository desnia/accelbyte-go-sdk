// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package matchmaking_operations

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
)

// NewGetHealthcheckInfoParams creates a new GetHealthcheckInfoParams object
// with the default values initialized.
func NewGetHealthcheckInfoParams() *GetHealthcheckInfoParams {

	return &GetHealthcheckInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHealthcheckInfoParamsWithTimeout creates a new GetHealthcheckInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHealthcheckInfoParamsWithTimeout(timeout time.Duration) *GetHealthcheckInfoParams {

	return &GetHealthcheckInfoParams{

		timeout: timeout,
	}
}

// NewGetHealthcheckInfoParamsWithContext creates a new GetHealthcheckInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHealthcheckInfoParamsWithContext(ctx context.Context) *GetHealthcheckInfoParams {

	return &GetHealthcheckInfoParams{

		Context: ctx,
	}
}

// NewGetHealthcheckInfoParamsWithHTTPClient creates a new GetHealthcheckInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHealthcheckInfoParamsWithHTTPClient(client *http.Client) *GetHealthcheckInfoParams {

	return &GetHealthcheckInfoParams{
		HTTPClient: client,
	}
}

/*GetHealthcheckInfoParams contains all the parameters to send to the API endpoint
for the get healthcheck info operation typically these are written to a http.Request
*/
type GetHealthcheckInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get healthcheck info params
func (o *GetHealthcheckInfoParams) WithTimeout(timeout time.Duration) *GetHealthcheckInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get healthcheck info params
func (o *GetHealthcheckInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get healthcheck info params
func (o *GetHealthcheckInfoParams) WithContext(ctx context.Context) *GetHealthcheckInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get healthcheck info params
func (o *GetHealthcheckInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get healthcheck info params
func (o *GetHealthcheckInfoParams) WithHTTPClient(client *http.Client) *GetHealthcheckInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get healthcheck info params
func (o *GetHealthcheckInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetHealthcheckInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}