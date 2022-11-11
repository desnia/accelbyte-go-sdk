// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package backfill

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

	"github.com/AccelByte/accelbyte-go-sdk/match2-sdk/pkg/match2clientmodels"
)

// NewAcceptBackfillParams creates a new AcceptBackfillParams object
// with the default values initialized.
func NewAcceptBackfillParams() *AcceptBackfillParams {
	var ()
	return &AcceptBackfillParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAcceptBackfillParamsWithTimeout creates a new AcceptBackfillParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAcceptBackfillParamsWithTimeout(timeout time.Duration) *AcceptBackfillParams {
	var ()
	return &AcceptBackfillParams{

		timeout: timeout,
	}
}

// NewAcceptBackfillParamsWithContext creates a new AcceptBackfillParams object
// with the default values initialized, and the ability to set a context for a request
func NewAcceptBackfillParamsWithContext(ctx context.Context) *AcceptBackfillParams {
	var ()
	return &AcceptBackfillParams{

		Context: ctx,
	}
}

// NewAcceptBackfillParamsWithHTTPClient creates a new AcceptBackfillParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAcceptBackfillParamsWithHTTPClient(client *http.Client) *AcceptBackfillParams {
	var ()
	return &AcceptBackfillParams{
		HTTPClient: client,
	}
}

/*AcceptBackfillParams contains all the parameters to send to the API endpoint
for the accept backfill operation typically these are written to a http.Request
*/
type AcceptBackfillParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*BackfillID
	  backfill id

	*/
	BackfillID string
	/*Body*/
	Body *match2clientmodels.APIBackFillAcceptRequest
	/*Namespace
	  namespace of the game

	*/
	Namespace string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the accept backfill params
func (o *AcceptBackfillParams) WithTimeout(timeout time.Duration) *AcceptBackfillParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accept backfill params
func (o *AcceptBackfillParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accept backfill params
func (o *AcceptBackfillParams) WithContext(ctx context.Context) *AcceptBackfillParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accept backfill params
func (o *AcceptBackfillParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the accept backfill params
func (o *AcceptBackfillParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the accept backfill params
func (o *AcceptBackfillParams) WithHTTPClient(client *http.Client) *AcceptBackfillParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accept backfill params
func (o *AcceptBackfillParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the accept backfill params
func (o *AcceptBackfillParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// WithBackfillID adds the backfillID to the accept backfill params
func (o *AcceptBackfillParams) WithBackfillID(backfillID string) *AcceptBackfillParams {
	o.SetBackfillID(backfillID)
	return o
}

// SetBackfillID adds the backfillId to the accept backfill params
func (o *AcceptBackfillParams) SetBackfillID(backfillID string) {
	o.BackfillID = backfillID
}

// WithBody adds the body to the accept backfill params
func (o *AcceptBackfillParams) WithBody(body *match2clientmodels.APIBackFillAcceptRequest) *AcceptBackfillParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the accept backfill params
func (o *AcceptBackfillParams) SetBody(body *match2clientmodels.APIBackFillAcceptRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the accept backfill params
func (o *AcceptBackfillParams) WithNamespace(namespace string) *AcceptBackfillParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the accept backfill params
func (o *AcceptBackfillParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *AcceptBackfillParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param backfillID
	if err := r.SetPathParam("backfillID", o.BackfillID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
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