// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package native_session

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

// NewAdminGetListNativeSessionParams creates a new AdminGetListNativeSessionParams object
// with the default values initialized.
func NewAdminGetListNativeSessionParams() *AdminGetListNativeSessionParams {
	var (
		limitDefault  = int64(20)
		offsetDefault = int64(0)
	)
	return &AdminGetListNativeSessionParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetListNativeSessionParamsWithTimeout creates a new AdminGetListNativeSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminGetListNativeSessionParamsWithTimeout(timeout time.Duration) *AdminGetListNativeSessionParams {
	var (
		limitDefault  = int64(20)
		offsetDefault = int64(0)
	)
	return &AdminGetListNativeSessionParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewAdminGetListNativeSessionParamsWithContext creates a new AdminGetListNativeSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminGetListNativeSessionParamsWithContext(ctx context.Context) *AdminGetListNativeSessionParams {
	var (
		limitDefault  = int64(20)
		offsetDefault = int64(0)
	)
	return &AdminGetListNativeSessionParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewAdminGetListNativeSessionParamsWithHTTPClient creates a new AdminGetListNativeSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminGetListNativeSessionParamsWithHTTPClient(client *http.Client) *AdminGetListNativeSessionParams {
	var (
		limitDefault  = int64(20)
		offsetDefault = int64(0)
	)
	return &AdminGetListNativeSessionParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*AdminGetListNativeSessionParams contains all the parameters to send to the API endpoint
for the admin get list native session operation typically these are written to a http.Request
*/
type AdminGetListNativeSessionParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*Limit
	  Pagination limit

	*/
	Limit *int64
	/*Offset
	  Pagination offset

	*/
	Offset *int64
	/*Order
	  Order of the result by createdAt. Supported: desc (default), asc

	*/
	Order *string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the admin get list native session params
func (o *AdminGetListNativeSessionParams) WithTimeout(timeout time.Duration) *AdminGetListNativeSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get list native session params
func (o *AdminGetListNativeSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get list native session params
func (o *AdminGetListNativeSessionParams) WithContext(ctx context.Context) *AdminGetListNativeSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get list native session params
func (o *AdminGetListNativeSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the admin get list native session params
func (o *AdminGetListNativeSessionParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the admin get list native session params
func (o *AdminGetListNativeSessionParams) WithHTTPClient(client *http.Client) *AdminGetListNativeSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get list native session params
func (o *AdminGetListNativeSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the admin get list native session params
func (o *AdminGetListNativeSessionParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// WithNamespace adds the namespace to the admin get list native session params
func (o *AdminGetListNativeSessionParams) WithNamespace(namespace string) *AdminGetListNativeSessionParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin get list native session params
func (o *AdminGetListNativeSessionParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithLimit adds the limit to the admin get list native session params
func (o *AdminGetListNativeSessionParams) WithLimit(limit *int64) *AdminGetListNativeSessionParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the admin get list native session params
func (o *AdminGetListNativeSessionParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the admin get list native session params
func (o *AdminGetListNativeSessionParams) WithOffset(offset *int64) *AdminGetListNativeSessionParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the admin get list native session params
func (o *AdminGetListNativeSessionParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrder adds the order to the admin get list native session params
func (o *AdminGetListNativeSessionParams) WithOrder(order *string) *AdminGetListNativeSessionParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the admin get list native session params
func (o *AdminGetListNativeSessionParams) SetOrder(order *string) {
	o.Order = order
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetListNativeSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Order != nil {

		// query param order
		var qrOrder string
		if o.Order != nil {
			qrOrder = *o.Order
		}
		qOrder := qrOrder
		if qOrder != "" {
			if err := r.SetQueryParam("order", qOrder); err != nil {
				return err
			}
		}

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
