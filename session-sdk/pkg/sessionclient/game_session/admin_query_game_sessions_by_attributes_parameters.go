// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package game_session

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

// NewAdminQueryGameSessionsByAttributesParams creates a new AdminQueryGameSessionsByAttributesParams object
// with the default values initialized.
func NewAdminQueryGameSessionsByAttributesParams() *AdminQueryGameSessionsByAttributesParams {
	var ()
	return &AdminQueryGameSessionsByAttributesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminQueryGameSessionsByAttributesParamsWithTimeout creates a new AdminQueryGameSessionsByAttributesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminQueryGameSessionsByAttributesParamsWithTimeout(timeout time.Duration) *AdminQueryGameSessionsByAttributesParams {
	var ()
	return &AdminQueryGameSessionsByAttributesParams{

		timeout: timeout,
	}
}

// NewAdminQueryGameSessionsByAttributesParamsWithContext creates a new AdminQueryGameSessionsByAttributesParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminQueryGameSessionsByAttributesParamsWithContext(ctx context.Context) *AdminQueryGameSessionsByAttributesParams {
	var ()
	return &AdminQueryGameSessionsByAttributesParams{

		Context: ctx,
	}
}

// NewAdminQueryGameSessionsByAttributesParamsWithHTTPClient creates a new AdminQueryGameSessionsByAttributesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminQueryGameSessionsByAttributesParamsWithHTTPClient(client *http.Client) *AdminQueryGameSessionsByAttributesParams {
	var ()
	return &AdminQueryGameSessionsByAttributesParams{
		HTTPClient: client,
	}
}

/*AdminQueryGameSessionsByAttributesParams contains all the parameters to send to the API endpoint
for the admin query game sessions by attributes operation typically these are written to a http.Request
*/
type AdminQueryGameSessionsByAttributesParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Body*/
	Body map[string]interface{}
	/*Namespace
	  namespace of the game

	*/
	Namespace string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the admin query game sessions by attributes params
func (o *AdminQueryGameSessionsByAttributesParams) WithTimeout(timeout time.Duration) *AdminQueryGameSessionsByAttributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin query game sessions by attributes params
func (o *AdminQueryGameSessionsByAttributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin query game sessions by attributes params
func (o *AdminQueryGameSessionsByAttributesParams) WithContext(ctx context.Context) *AdminQueryGameSessionsByAttributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin query game sessions by attributes params
func (o *AdminQueryGameSessionsByAttributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the admin query game sessions by attributes params
func (o *AdminQueryGameSessionsByAttributesParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the admin query game sessions by attributes params
func (o *AdminQueryGameSessionsByAttributesParams) WithHTTPClient(client *http.Client) *AdminQueryGameSessionsByAttributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin query game sessions by attributes params
func (o *AdminQueryGameSessionsByAttributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the admin query game sessions by attributes params
func (o *AdminQueryGameSessionsByAttributesParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// WithBody adds the body to the admin query game sessions by attributes params
func (o *AdminQueryGameSessionsByAttributesParams) WithBody(body map[string]interface{}) *AdminQueryGameSessionsByAttributesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin query game sessions by attributes params
func (o *AdminQueryGameSessionsByAttributesParams) SetBody(body map[string]interface{}) {
	o.Body = body
}

// WithNamespace adds the namespace to the admin query game sessions by attributes params
func (o *AdminQueryGameSessionsByAttributesParams) WithNamespace(namespace string) *AdminQueryGameSessionsByAttributesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin query game sessions by attributes params
func (o *AdminQueryGameSessionsByAttributesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *AdminQueryGameSessionsByAttributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
