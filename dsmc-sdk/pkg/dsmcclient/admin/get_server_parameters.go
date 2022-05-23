// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewGetServerParams creates a new GetServerParams object
// with the default values initialized.
func NewGetServerParams() *GetServerParams {
	var ()
	return &GetServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerParamsWithTimeout creates a new GetServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServerParamsWithTimeout(timeout time.Duration) *GetServerParams {
	var ()
	return &GetServerParams{

		timeout: timeout,
	}
}

// NewGetServerParamsWithContext creates a new GetServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServerParamsWithContext(ctx context.Context) *GetServerParams {
	var ()
	return &GetServerParams{

		Context: ctx,
	}
}

// NewGetServerParamsWithHTTPClient creates a new GetServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServerParamsWithHTTPClient(client *http.Client) *GetServerParams {
	var ()
	return &GetServerParams{
		HTTPClient: client,
	}
}

/*GetServerParams contains all the parameters to send to the API endpoint
for the get server operation typically these are written to a http.Request
*/
type GetServerParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*PodName
	  name of the DS pod

	*/
	PodName string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the get server params
func (o *GetServerParams) WithTimeout(timeout time.Duration) *GetServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get server params
func (o *GetServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get server params
func (o *GetServerParams) WithContext(ctx context.Context) *GetServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get server params
func (o *GetServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the get server params
func (o *GetServerParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the get server params
func (o *GetServerParams) WithHTTPClient(client *http.Client) *GetServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get server params
func (o *GetServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get server params
func (o *GetServerParams) WithNamespace(namespace string) *GetServerParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get server params
func (o *GetServerParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPodName adds the podName to the get server params
func (o *GetServerParams) WithPodName(podName string) *GetServerParams {
	o.SetPodName(podName)
	return o
}

// SetPodName adds the podName to the get server params
func (o *GetServerParams) SetPodName(podName string) {
	o.PodName = podName
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param podName
	if err := r.SetPathParam("podName", o.PodName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
