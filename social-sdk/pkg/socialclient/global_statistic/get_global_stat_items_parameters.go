// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package global_statistic

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
	"github.com/go-openapi/swag"
)

// NewGetGlobalStatItemsParams creates a new GetGlobalStatItemsParams object
// with the default values initialized.
func NewGetGlobalStatItemsParams() *GetGlobalStatItemsParams {
	var (
		limitDefault  = int32(20)
		offsetDefault = int32(0)
	)
	return &GetGlobalStatItemsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGlobalStatItemsParamsWithTimeout creates a new GetGlobalStatItemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGlobalStatItemsParamsWithTimeout(timeout time.Duration) *GetGlobalStatItemsParams {
	var (
		limitDefault  = int32(20)
		offsetDefault = int32(0)
	)
	return &GetGlobalStatItemsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewGetGlobalStatItemsParamsWithContext creates a new GetGlobalStatItemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGlobalStatItemsParamsWithContext(ctx context.Context) *GetGlobalStatItemsParams {
	var (
		limitDefault  = int32(20)
		offsetDefault = int32(0)
	)
	return &GetGlobalStatItemsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewGetGlobalStatItemsParamsWithHTTPClient creates a new GetGlobalStatItemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGlobalStatItemsParamsWithHTTPClient(client *http.Client) *GetGlobalStatItemsParams {
	var (
		limitDefault  = int32(20)
		offsetDefault = int32(0)
	)
	return &GetGlobalStatItemsParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetGlobalStatItemsParams contains all the parameters to send to the API endpoint
for the get global stat items operation typically these are written to a http.Request
*/
type GetGlobalStatItemsParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Limit*/
	Limit *int32
	/*Namespace
	  namespace

	*/
	Namespace string
	/*Offset*/
	Offset *int32

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the get global stat items params
func (o *GetGlobalStatItemsParams) WithTimeout(timeout time.Duration) *GetGlobalStatItemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get global stat items params
func (o *GetGlobalStatItemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get global stat items params
func (o *GetGlobalStatItemsParams) WithContext(ctx context.Context) *GetGlobalStatItemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get global stat items params
func (o *GetGlobalStatItemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the get global stat items params
func (o *GetGlobalStatItemsParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the get global stat items params
func (o *GetGlobalStatItemsParams) WithHTTPClient(client *http.Client) *GetGlobalStatItemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get global stat items params
func (o *GetGlobalStatItemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get global stat items params
func (o *GetGlobalStatItemsParams) WithLimit(limit *int32) *GetGlobalStatItemsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get global stat items params
func (o *GetGlobalStatItemsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the get global stat items params
func (o *GetGlobalStatItemsParams) WithNamespace(namespace string) *GetGlobalStatItemsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get global stat items params
func (o *GetGlobalStatItemsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the get global stat items params
func (o *GetGlobalStatItemsParams) WithOffset(offset *int32) *GetGlobalStatItemsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get global stat items params
func (o *GetGlobalStatItemsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetGlobalStatItemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
