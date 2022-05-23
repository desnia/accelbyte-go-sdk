// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package data_retrieval

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

// NewPublicGeneratePersonalDataURLParams creates a new PublicGeneratePersonalDataURLParams object
// with the default values initialized.
func NewPublicGeneratePersonalDataURLParams() *PublicGeneratePersonalDataURLParams {
	var ()
	return &PublicGeneratePersonalDataURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicGeneratePersonalDataURLParamsWithTimeout creates a new PublicGeneratePersonalDataURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicGeneratePersonalDataURLParamsWithTimeout(timeout time.Duration) *PublicGeneratePersonalDataURLParams {
	var ()
	return &PublicGeneratePersonalDataURLParams{

		timeout: timeout,
	}
}

// NewPublicGeneratePersonalDataURLParamsWithContext creates a new PublicGeneratePersonalDataURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicGeneratePersonalDataURLParamsWithContext(ctx context.Context) *PublicGeneratePersonalDataURLParams {
	var ()
	return &PublicGeneratePersonalDataURLParams{

		Context: ctx,
	}
}

// NewPublicGeneratePersonalDataURLParamsWithHTTPClient creates a new PublicGeneratePersonalDataURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicGeneratePersonalDataURLParamsWithHTTPClient(client *http.Client) *PublicGeneratePersonalDataURLParams {
	var ()
	return &PublicGeneratePersonalDataURLParams{
		HTTPClient: client,
	}
}

/*PublicGeneratePersonalDataURLParams contains all the parameters to send to the API endpoint
for the public generate personal data URL operation typically these are written to a http.Request
*/
type PublicGeneratePersonalDataURLParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Namespace
	  namespace of the user

	*/
	Namespace string
	/*Password
	  IAM password of the user

	*/
	Password string
	/*RequestDate
	  Request date in RFC3339 format

	*/
	RequestDate string
	/*UserID
	  IAM id of the user

	*/
	UserID string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the public generate personal data URL params
func (o *PublicGeneratePersonalDataURLParams) WithTimeout(timeout time.Duration) *PublicGeneratePersonalDataURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public generate personal data URL params
func (o *PublicGeneratePersonalDataURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public generate personal data URL params
func (o *PublicGeneratePersonalDataURLParams) WithContext(ctx context.Context) *PublicGeneratePersonalDataURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public generate personal data URL params
func (o *PublicGeneratePersonalDataURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the public generate personal data URL params
func (o *PublicGeneratePersonalDataURLParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the public generate personal data URL params
func (o *PublicGeneratePersonalDataURLParams) WithHTTPClient(client *http.Client) *PublicGeneratePersonalDataURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public generate personal data URL params
func (o *PublicGeneratePersonalDataURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the public generate personal data URL params
func (o *PublicGeneratePersonalDataURLParams) WithNamespace(namespace string) *PublicGeneratePersonalDataURLParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public generate personal data URL params
func (o *PublicGeneratePersonalDataURLParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPassword adds the password to the public generate personal data URL params
func (o *PublicGeneratePersonalDataURLParams) WithPassword(password string) *PublicGeneratePersonalDataURLParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the public generate personal data URL params
func (o *PublicGeneratePersonalDataURLParams) SetPassword(password string) {
	o.Password = password
}

// WithRequestDate adds the requestDate to the public generate personal data URL params
func (o *PublicGeneratePersonalDataURLParams) WithRequestDate(requestDate string) *PublicGeneratePersonalDataURLParams {
	o.SetRequestDate(requestDate)
	return o
}

// SetRequestDate adds the requestDate to the public generate personal data URL params
func (o *PublicGeneratePersonalDataURLParams) SetRequestDate(requestDate string) {
	o.RequestDate = requestDate
}

// WithUserID adds the userID to the public generate personal data URL params
func (o *PublicGeneratePersonalDataURLParams) WithUserID(userID string) *PublicGeneratePersonalDataURLParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public generate personal data URL params
func (o *PublicGeneratePersonalDataURLParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicGeneratePersonalDataURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// form param password
	frPassword := o.Password
	fPassword := frPassword
	if fPassword != "" {
		if err := r.SetFormParam("password", fPassword); err != nil {
			return err
		}
	}

	// path param requestDate
	if err := r.SetPathParam("requestDate", o.RequestDate); err != nil {
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
