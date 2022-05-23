// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package deployment_config

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

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// NewUpdateDeploymentOverrideParams creates a new UpdateDeploymentOverrideParams object
// with the default values initialized.
func NewUpdateDeploymentOverrideParams() *UpdateDeploymentOverrideParams {
	var ()
	return &UpdateDeploymentOverrideParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeploymentOverrideParamsWithTimeout creates a new UpdateDeploymentOverrideParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDeploymentOverrideParamsWithTimeout(timeout time.Duration) *UpdateDeploymentOverrideParams {
	var ()
	return &UpdateDeploymentOverrideParams{

		timeout: timeout,
	}
}

// NewUpdateDeploymentOverrideParamsWithContext creates a new UpdateDeploymentOverrideParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDeploymentOverrideParamsWithContext(ctx context.Context) *UpdateDeploymentOverrideParams {
	var ()
	return &UpdateDeploymentOverrideParams{

		Context: ctx,
	}
}

// NewUpdateDeploymentOverrideParamsWithHTTPClient creates a new UpdateDeploymentOverrideParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDeploymentOverrideParamsWithHTTPClient(client *http.Client) *UpdateDeploymentOverrideParams {
	var ()
	return &UpdateDeploymentOverrideParams{
		HTTPClient: client,
	}
}

/*UpdateDeploymentOverrideParams contains all the parameters to send to the API endpoint
for the update deployment override operation typically these are written to a http.Request
*/
type UpdateDeploymentOverrideParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Body*/
	Body *dsmcclientmodels.ModelsUpdateDeploymentOverrideRequest
	/*Deployment
	  deployment name

	*/
	Deployment string
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*Version
	  version

	*/
	Version string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the update deployment override params
func (o *UpdateDeploymentOverrideParams) WithTimeout(timeout time.Duration) *UpdateDeploymentOverrideParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update deployment override params
func (o *UpdateDeploymentOverrideParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update deployment override params
func (o *UpdateDeploymentOverrideParams) WithContext(ctx context.Context) *UpdateDeploymentOverrideParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update deployment override params
func (o *UpdateDeploymentOverrideParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the update deployment override params
func (o *UpdateDeploymentOverrideParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the update deployment override params
func (o *UpdateDeploymentOverrideParams) WithHTTPClient(client *http.Client) *UpdateDeploymentOverrideParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update deployment override params
func (o *UpdateDeploymentOverrideParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update deployment override params
func (o *UpdateDeploymentOverrideParams) WithBody(body *dsmcclientmodels.ModelsUpdateDeploymentOverrideRequest) *UpdateDeploymentOverrideParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update deployment override params
func (o *UpdateDeploymentOverrideParams) SetBody(body *dsmcclientmodels.ModelsUpdateDeploymentOverrideRequest) {
	o.Body = body
}

// WithDeployment adds the deployment to the update deployment override params
func (o *UpdateDeploymentOverrideParams) WithDeployment(deployment string) *UpdateDeploymentOverrideParams {
	o.SetDeployment(deployment)
	return o
}

// SetDeployment adds the deployment to the update deployment override params
func (o *UpdateDeploymentOverrideParams) SetDeployment(deployment string) {
	o.Deployment = deployment
}

// WithNamespace adds the namespace to the update deployment override params
func (o *UpdateDeploymentOverrideParams) WithNamespace(namespace string) *UpdateDeploymentOverrideParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update deployment override params
func (o *UpdateDeploymentOverrideParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithVersion adds the version to the update deployment override params
func (o *UpdateDeploymentOverrideParams) WithVersion(version string) *UpdateDeploymentOverrideParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update deployment override params
func (o *UpdateDeploymentOverrideParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeploymentOverrideParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deployment
	if err := r.SetPathParam("deployment", o.Deployment); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
