// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewPublicPlatformLinkV2Params creates a new PublicPlatformLinkV2Params object
// with the default values initialized.
func NewPublicPlatformLinkV2Params() *PublicPlatformLinkV2Params {
	var ()
	return &PublicPlatformLinkV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicPlatformLinkV2ParamsWithTimeout creates a new PublicPlatformLinkV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicPlatformLinkV2ParamsWithTimeout(timeout time.Duration) *PublicPlatformLinkV2Params {
	var ()
	return &PublicPlatformLinkV2Params{

		timeout: timeout,
	}
}

// NewPublicPlatformLinkV2ParamsWithContext creates a new PublicPlatformLinkV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewPublicPlatformLinkV2ParamsWithContext(ctx context.Context) *PublicPlatformLinkV2Params {
	var ()
	return &PublicPlatformLinkV2Params{

		Context: ctx,
	}
}

// NewPublicPlatformLinkV2ParamsWithHTTPClient creates a new PublicPlatformLinkV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicPlatformLinkV2ParamsWithHTTPClient(client *http.Client) *PublicPlatformLinkV2Params {
	var ()
	return &PublicPlatformLinkV2Params{
		HTTPClient: client,
	}
}

/*PublicPlatformLinkV2Params contains all the parameters to send to the API endpoint
for the public platform link v2 operation typically these are written to a http.Request
*/
type PublicPlatformLinkV2Params struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Namespace
	  Namespace

	*/
	Namespace string
	/*PlatformID
	  Platform ID

	*/
	PlatformID string
	/*Ticket
	  Ticket from platform

	*/
	Ticket string
	/*UserID
	  User ID

	*/
	UserID string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the public platform link v2 params
func (o *PublicPlatformLinkV2Params) WithTimeout(timeout time.Duration) *PublicPlatformLinkV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public platform link v2 params
func (o *PublicPlatformLinkV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public platform link v2 params
func (o *PublicPlatformLinkV2Params) WithContext(ctx context.Context) *PublicPlatformLinkV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public platform link v2 params
func (o *PublicPlatformLinkV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the public platform link v2 params
func (o *PublicPlatformLinkV2Params) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the public platform link v2 params
func (o *PublicPlatformLinkV2Params) WithHTTPClient(client *http.Client) *PublicPlatformLinkV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public platform link v2 params
func (o *PublicPlatformLinkV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the public platform link v2 params
func (o *PublicPlatformLinkV2Params) WithNamespace(namespace string) *PublicPlatformLinkV2Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public platform link v2 params
func (o *PublicPlatformLinkV2Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPlatformID adds the platformID to the public platform link v2 params
func (o *PublicPlatformLinkV2Params) WithPlatformID(platformID string) *PublicPlatformLinkV2Params {
	o.SetPlatformID(platformID)
	return o
}

// SetPlatformID adds the platformId to the public platform link v2 params
func (o *PublicPlatformLinkV2Params) SetPlatformID(platformID string) {
	o.PlatformID = platformID
}

// WithTicket adds the ticket to the public platform link v2 params
func (o *PublicPlatformLinkV2Params) WithTicket(ticket string) *PublicPlatformLinkV2Params {
	o.SetTicket(ticket)
	return o
}

// SetTicket adds the ticket to the public platform link v2 params
func (o *PublicPlatformLinkV2Params) SetTicket(ticket string) {
	o.Ticket = ticket
}

// WithUserID adds the userID to the public platform link v2 params
func (o *PublicPlatformLinkV2Params) WithUserID(userID string) *PublicPlatformLinkV2Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public platform link v2 params
func (o *PublicPlatformLinkV2Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicPlatformLinkV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param platformId
	if err := r.SetPathParam("platformId", o.PlatformID); err != nil {
		return err
	}

	// form param ticket
	frTicket := o.Ticket
	fTicket := frTicket
	if fTicket != "" {
		if err := r.SetFormParam("ticket", fTicket); err != nil {
			return err
		}
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
