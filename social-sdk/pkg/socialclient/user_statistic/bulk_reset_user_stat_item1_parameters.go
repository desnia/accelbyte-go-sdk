// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package user_statistic

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

	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
)

// NewBulkResetUserStatItem1Params creates a new BulkResetUserStatItem1Params object
// with the default values initialized.
func NewBulkResetUserStatItem1Params() *BulkResetUserStatItem1Params {
	var ()
	return &BulkResetUserStatItem1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewBulkResetUserStatItem1ParamsWithTimeout creates a new BulkResetUserStatItem1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewBulkResetUserStatItem1ParamsWithTimeout(timeout time.Duration) *BulkResetUserStatItem1Params {
	var ()
	return &BulkResetUserStatItem1Params{

		timeout: timeout,
	}
}

// NewBulkResetUserStatItem1ParamsWithContext creates a new BulkResetUserStatItem1Params object
// with the default values initialized, and the ability to set a context for a request
func NewBulkResetUserStatItem1ParamsWithContext(ctx context.Context) *BulkResetUserStatItem1Params {
	var ()
	return &BulkResetUserStatItem1Params{

		Context: ctx,
	}
}

// NewBulkResetUserStatItem1ParamsWithHTTPClient creates a new BulkResetUserStatItem1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBulkResetUserStatItem1ParamsWithHTTPClient(client *http.Client) *BulkResetUserStatItem1Params {
	var ()
	return &BulkResetUserStatItem1Params{
		HTTPClient: client,
	}
}

/*BulkResetUserStatItem1Params contains all the parameters to send to the API endpoint
for the bulk reset user stat item 1 operation typically these are written to a http.Request
*/
type BulkResetUserStatItem1Params struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Body*/
	Body []*socialclientmodels.BulkStatItemReset
	/*Namespace
	  namespace

	*/
	Namespace string
	/*UserID
	  user ID

	*/
	UserID string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the bulk reset user stat item 1 params
func (o *BulkResetUserStatItem1Params) WithTimeout(timeout time.Duration) *BulkResetUserStatItem1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk reset user stat item 1 params
func (o *BulkResetUserStatItem1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk reset user stat item 1 params
func (o *BulkResetUserStatItem1Params) WithContext(ctx context.Context) *BulkResetUserStatItem1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk reset user stat item 1 params
func (o *BulkResetUserStatItem1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the bulk reset user stat item 1 params
func (o *BulkResetUserStatItem1Params) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the bulk reset user stat item 1 params
func (o *BulkResetUserStatItem1Params) WithHTTPClient(client *http.Client) *BulkResetUserStatItem1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk reset user stat item 1 params
func (o *BulkResetUserStatItem1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the bulk reset user stat item 1 params
func (o *BulkResetUserStatItem1Params) WithBody(body []*socialclientmodels.BulkStatItemReset) *BulkResetUserStatItem1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the bulk reset user stat item 1 params
func (o *BulkResetUserStatItem1Params) SetBody(body []*socialclientmodels.BulkStatItemReset) {
	o.Body = body
}

// WithNamespace adds the namespace to the bulk reset user stat item 1 params
func (o *BulkResetUserStatItem1Params) WithNamespace(namespace string) *BulkResetUserStatItem1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the bulk reset user stat item 1 params
func (o *BulkResetUserStatItem1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the bulk reset user stat item 1 params
func (o *BulkResetUserStatItem1Params) WithUserID(userID string) *BulkResetUserStatItem1Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the bulk reset user stat item 1 params
func (o *BulkResetUserStatItem1Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *BulkResetUserStatItem1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
