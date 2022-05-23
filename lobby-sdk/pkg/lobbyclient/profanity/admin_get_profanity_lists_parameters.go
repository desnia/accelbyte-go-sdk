// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package profanity

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

// NewAdminGetProfanityListsParams creates a new AdminGetProfanityListsParams object
// with the default values initialized.
func NewAdminGetProfanityListsParams() *AdminGetProfanityListsParams {
	var ()
	return &AdminGetProfanityListsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetProfanityListsParamsWithTimeout creates a new AdminGetProfanityListsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminGetProfanityListsParamsWithTimeout(timeout time.Duration) *AdminGetProfanityListsParams {
	var ()
	return &AdminGetProfanityListsParams{

		timeout: timeout,
	}
}

// NewAdminGetProfanityListsParamsWithContext creates a new AdminGetProfanityListsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminGetProfanityListsParamsWithContext(ctx context.Context) *AdminGetProfanityListsParams {
	var ()
	return &AdminGetProfanityListsParams{

		Context: ctx,
	}
}

// NewAdminGetProfanityListsParamsWithHTTPClient creates a new AdminGetProfanityListsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminGetProfanityListsParamsWithHTTPClient(client *http.Client) *AdminGetProfanityListsParams {
	var ()
	return &AdminGetProfanityListsParams{
		HTTPClient: client,
	}
}

/*AdminGetProfanityListsParams contains all the parameters to send to the API endpoint
for the admin get profanity lists operation typically these are written to a http.Request
*/
type AdminGetProfanityListsParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Namespace
	  namespace

	*/
	Namespace string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the admin get profanity lists params
func (o *AdminGetProfanityListsParams) WithTimeout(timeout time.Duration) *AdminGetProfanityListsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get profanity lists params
func (o *AdminGetProfanityListsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get profanity lists params
func (o *AdminGetProfanityListsParams) WithContext(ctx context.Context) *AdminGetProfanityListsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get profanity lists params
func (o *AdminGetProfanityListsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the admin get profanity lists params
func (o *AdminGetProfanityListsParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the admin get profanity lists params
func (o *AdminGetProfanityListsParams) WithHTTPClient(client *http.Client) *AdminGetProfanityListsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get profanity lists params
func (o *AdminGetProfanityListsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the admin get profanity lists params
func (o *AdminGetProfanityListsParams) WithNamespace(namespace string) *AdminGetProfanityListsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin get profanity lists params
func (o *AdminGetProfanityListsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetProfanityListsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
