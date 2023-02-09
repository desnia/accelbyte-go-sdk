// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package admin_reasons

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

	"github.com/AccelByte/accelbyte-go-sdk/reporting-sdk/pkg/reportingclientmodels"
)

// NewUpdateReasonGroupParams creates a new UpdateReasonGroupParams object
// with the default values initialized.
func NewUpdateReasonGroupParams() *UpdateReasonGroupParams {
	var ()
	return &UpdateReasonGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateReasonGroupParamsWithTimeout creates a new UpdateReasonGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateReasonGroupParamsWithTimeout(timeout time.Duration) *UpdateReasonGroupParams {
	var ()
	return &UpdateReasonGroupParams{

		timeout: timeout,
	}
}

// NewUpdateReasonGroupParamsWithContext creates a new UpdateReasonGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateReasonGroupParamsWithContext(ctx context.Context) *UpdateReasonGroupParams {
	var ()
	return &UpdateReasonGroupParams{

		Context: ctx,
	}
}

// NewUpdateReasonGroupParamsWithHTTPClient creates a new UpdateReasonGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateReasonGroupParamsWithHTTPClient(client *http.Client) *UpdateReasonGroupParams {
	var ()
	return &UpdateReasonGroupParams{
		HTTPClient: client,
	}
}

/*
UpdateReasonGroupParams contains all the parameters to send to the API endpoint
for the update reason group operation typically these are written to a http.Request
*/
type UpdateReasonGroupParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Body*/
	Body *reportingclientmodels.RestapiUpdateReasonGroupRequest
	/*GroupID
	  Reason group ID

	*/
	GroupID string
	/*Namespace*/
	Namespace string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the update reason group params
func (o *UpdateReasonGroupParams) WithTimeout(timeout time.Duration) *UpdateReasonGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update reason group params
func (o *UpdateReasonGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update reason group params
func (o *UpdateReasonGroupParams) WithContext(ctx context.Context) *UpdateReasonGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update reason group params
func (o *UpdateReasonGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the update reason group params
func (o *UpdateReasonGroupParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the update reason group params
func (o *UpdateReasonGroupParams) WithHTTPClient(client *http.Client) *UpdateReasonGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update reason group params
func (o *UpdateReasonGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the update reason group params
func (o *UpdateReasonGroupParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// WithBody adds the body to the update reason group params
func (o *UpdateReasonGroupParams) WithBody(body *reportingclientmodels.RestapiUpdateReasonGroupRequest) *UpdateReasonGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update reason group params
func (o *UpdateReasonGroupParams) SetBody(body *reportingclientmodels.RestapiUpdateReasonGroupRequest) {
	o.Body = body
}

// WithGroupID adds the groupID to the update reason group params
func (o *UpdateReasonGroupParams) WithGroupID(groupID string) *UpdateReasonGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the update reason group params
func (o *UpdateReasonGroupParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithNamespace adds the namespace to the update reason group params
func (o *UpdateReasonGroupParams) WithNamespace(namespace string) *UpdateReasonGroupParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update reason group params
func (o *UpdateReasonGroupParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateReasonGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
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
