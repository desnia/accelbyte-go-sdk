// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package session

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

	"github.com/AccelByte/accelbyte-go-sdk/sessionbrowser-sdk/pkg/sessionbrowserclientmodels"
)

// NewJoinSessionParams creates a new JoinSessionParams object
// with the default values initialized.
func NewJoinSessionParams() *JoinSessionParams {
	var ()
	return &JoinSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJoinSessionParamsWithTimeout creates a new JoinSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJoinSessionParamsWithTimeout(timeout time.Duration) *JoinSessionParams {
	var ()
	return &JoinSessionParams{

		timeout: timeout,
	}
}

// NewJoinSessionParamsWithContext creates a new JoinSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewJoinSessionParamsWithContext(ctx context.Context) *JoinSessionParams {
	var ()
	return &JoinSessionParams{

		Context: ctx,
	}
}

// NewJoinSessionParamsWithHTTPClient creates a new JoinSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJoinSessionParamsWithHTTPClient(client *http.Client) *JoinSessionParams {
	var ()
	return &JoinSessionParams{
		HTTPClient: client,
	}
}

/*JoinSessionParams contains all the parameters to send to the API endpoint
for the join session operation typically these are written to a http.Request
*/
type JoinSessionParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Body*/
	Body *sessionbrowserclientmodels.ModelsJoinGameSessionRequest
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*SessionID
	  session ID

	*/
	SessionID string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the join session params
func (o *JoinSessionParams) WithTimeout(timeout time.Duration) *JoinSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the join session params
func (o *JoinSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the join session params
func (o *JoinSessionParams) WithContext(ctx context.Context) *JoinSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the join session params
func (o *JoinSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the join session params
func (o *JoinSessionParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the join session params
func (o *JoinSessionParams) WithHTTPClient(client *http.Client) *JoinSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the join session params
func (o *JoinSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the join session params
func (o *JoinSessionParams) WithBody(body *sessionbrowserclientmodels.ModelsJoinGameSessionRequest) *JoinSessionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the join session params
func (o *JoinSessionParams) SetBody(body *sessionbrowserclientmodels.ModelsJoinGameSessionRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the join session params
func (o *JoinSessionParams) WithNamespace(namespace string) *JoinSessionParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the join session params
func (o *JoinSessionParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithSessionID adds the sessionID to the join session params
func (o *JoinSessionParams) WithSessionID(sessionID string) *JoinSessionParams {
	o.SetSessionID(sessionID)
	return o
}

// SetSessionID adds the sessionId to the join session params
func (o *JoinSessionParams) SetSessionID(sessionID string) {
	o.SessionID = sessionID
}

// WriteToRequest writes these params to a swagger request
func (o *JoinSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param sessionID
	if err := r.SetPathParam("sessionID", o.SessionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
