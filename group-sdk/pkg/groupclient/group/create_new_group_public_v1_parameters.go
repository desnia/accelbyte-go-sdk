// Code generated by go-swagger; DO NOT EDIT.

package group

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

	"github.com/AccelByte/accelbyte-go-sdk/group-sdk/pkg/groupclientmodels"
)

// NewCreateNewGroupPublicV1Params creates a new CreateNewGroupPublicV1Params object
// with the default values initialized.
func NewCreateNewGroupPublicV1Params() *CreateNewGroupPublicV1Params {
	var ()
	return &CreateNewGroupPublicV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNewGroupPublicV1ParamsWithTimeout creates a new CreateNewGroupPublicV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateNewGroupPublicV1ParamsWithTimeout(timeout time.Duration) *CreateNewGroupPublicV1Params {
	var ()
	return &CreateNewGroupPublicV1Params{

		timeout: timeout,
	}
}

// NewCreateNewGroupPublicV1ParamsWithContext creates a new CreateNewGroupPublicV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateNewGroupPublicV1ParamsWithContext(ctx context.Context) *CreateNewGroupPublicV1Params {
	var ()
	return &CreateNewGroupPublicV1Params{

		Context: ctx,
	}
}

// NewCreateNewGroupPublicV1ParamsWithHTTPClient creates a new CreateNewGroupPublicV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateNewGroupPublicV1ParamsWithHTTPClient(client *http.Client) *CreateNewGroupPublicV1Params {
	var ()
	return &CreateNewGroupPublicV1Params{
		HTTPClient: client,
	}
}

/*CreateNewGroupPublicV1Params contains all the parameters to send to the API endpoint
for the create new group public v1 operation typically these are written to a http.Request
*/
type CreateNewGroupPublicV1Params struct {

	/*Body*/
	Body *groupclientmodels.ModelsPublicCreateNewGroupRequestV1
	/*Namespace
	  namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create new group public v1 params
func (o *CreateNewGroupPublicV1Params) WithTimeout(timeout time.Duration) *CreateNewGroupPublicV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create new group public v1 params
func (o *CreateNewGroupPublicV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create new group public v1 params
func (o *CreateNewGroupPublicV1Params) WithContext(ctx context.Context) *CreateNewGroupPublicV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create new group public v1 params
func (o *CreateNewGroupPublicV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create new group public v1 params
func (o *CreateNewGroupPublicV1Params) WithHTTPClient(client *http.Client) *CreateNewGroupPublicV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create new group public v1 params
func (o *CreateNewGroupPublicV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create new group public v1 params
func (o *CreateNewGroupPublicV1Params) WithBody(body *groupclientmodels.ModelsPublicCreateNewGroupRequestV1) *CreateNewGroupPublicV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create new group public v1 params
func (o *CreateNewGroupPublicV1Params) SetBody(body *groupclientmodels.ModelsPublicCreateNewGroupRequestV1) {
	o.Body = body
}

// WithNamespace adds the namespace to the create new group public v1 params
func (o *CreateNewGroupPublicV1Params) WithNamespace(namespace string) *CreateNewGroupPublicV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the create new group public v1 params
func (o *CreateNewGroupPublicV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNewGroupPublicV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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