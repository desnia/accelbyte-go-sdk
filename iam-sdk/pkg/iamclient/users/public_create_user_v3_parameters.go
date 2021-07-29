// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
)

// NewPublicCreateUserV3Params creates a new PublicCreateUserV3Params object
// with the default values initialized.
func NewPublicCreateUserV3Params() *PublicCreateUserV3Params {
	var ()
	return &PublicCreateUserV3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicCreateUserV3ParamsWithTimeout creates a new PublicCreateUserV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicCreateUserV3ParamsWithTimeout(timeout time.Duration) *PublicCreateUserV3Params {
	var ()
	return &PublicCreateUserV3Params{

		timeout: timeout,
	}
}

// NewPublicCreateUserV3ParamsWithContext creates a new PublicCreateUserV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewPublicCreateUserV3ParamsWithContext(ctx context.Context) *PublicCreateUserV3Params {
	var ()
	return &PublicCreateUserV3Params{

		Context: ctx,
	}
}

// NewPublicCreateUserV3ParamsWithHTTPClient creates a new PublicCreateUserV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicCreateUserV3ParamsWithHTTPClient(client *http.Client) *PublicCreateUserV3Params {
	var ()
	return &PublicCreateUserV3Params{
		HTTPClient: client,
	}
}

/*PublicCreateUserV3Params contains all the parameters to send to the API endpoint
for the public create user v3 operation typically these are written to a http.Request
*/
type PublicCreateUserV3Params struct {

	/*Body*/
	Body *iamclientmodels.ModelUserCreateRequestV3
	/*Namespace
	  Namespace, only accept alphabet and numeric

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public create user v3 params
func (o *PublicCreateUserV3Params) WithTimeout(timeout time.Duration) *PublicCreateUserV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public create user v3 params
func (o *PublicCreateUserV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public create user v3 params
func (o *PublicCreateUserV3Params) WithContext(ctx context.Context) *PublicCreateUserV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public create user v3 params
func (o *PublicCreateUserV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public create user v3 params
func (o *PublicCreateUserV3Params) WithHTTPClient(client *http.Client) *PublicCreateUserV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public create user v3 params
func (o *PublicCreateUserV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the public create user v3 params
func (o *PublicCreateUserV3Params) WithBody(body *iamclientmodels.ModelUserCreateRequestV3) *PublicCreateUserV3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the public create user v3 params
func (o *PublicCreateUserV3Params) SetBody(body *iamclientmodels.ModelUserCreateRequestV3) {
	o.Body = body
}

// WithNamespace adds the namespace to the public create user v3 params
func (o *PublicCreateUserV3Params) WithNamespace(namespace string) *PublicCreateUserV3Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public create user v3 params
func (o *PublicCreateUserV3Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *PublicCreateUserV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}