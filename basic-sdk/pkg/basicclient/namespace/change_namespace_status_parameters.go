// Code generated by go-swagger; DO NOT EDIT.

package namespace

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

	"github.com/AccelByte/accelbyte-go-sdk/basic-sdk/pkg/basicclientmodels"
)

// NewChangeNamespaceStatusParams creates a new ChangeNamespaceStatusParams object
// with the default values initialized.
func NewChangeNamespaceStatusParams() *ChangeNamespaceStatusParams {
	var ()
	return &ChangeNamespaceStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeNamespaceStatusParamsWithTimeout creates a new ChangeNamespaceStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeNamespaceStatusParamsWithTimeout(timeout time.Duration) *ChangeNamespaceStatusParams {
	var ()
	return &ChangeNamespaceStatusParams{

		timeout: timeout,
	}
}

// NewChangeNamespaceStatusParamsWithContext creates a new ChangeNamespaceStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeNamespaceStatusParamsWithContext(ctx context.Context) *ChangeNamespaceStatusParams {
	var ()
	return &ChangeNamespaceStatusParams{

		Context: ctx,
	}
}

// NewChangeNamespaceStatusParamsWithHTTPClient creates a new ChangeNamespaceStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeNamespaceStatusParamsWithHTTPClient(client *http.Client) *ChangeNamespaceStatusParams {
	var ()
	return &ChangeNamespaceStatusParams{
		HTTPClient: client,
	}
}

/*ChangeNamespaceStatusParams contains all the parameters to send to the API endpoint
for the change namespace status operation typically these are written to a http.Request
*/
type ChangeNamespaceStatusParams struct {

	/*Body*/
	Body *basicclientmodels.NamespaceStatusUpdate
	/*Namespace
	  namespace, only accept alphabet and numeric

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change namespace status params
func (o *ChangeNamespaceStatusParams) WithTimeout(timeout time.Duration) *ChangeNamespaceStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change namespace status params
func (o *ChangeNamespaceStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change namespace status params
func (o *ChangeNamespaceStatusParams) WithContext(ctx context.Context) *ChangeNamespaceStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change namespace status params
func (o *ChangeNamespaceStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change namespace status params
func (o *ChangeNamespaceStatusParams) WithHTTPClient(client *http.Client) *ChangeNamespaceStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change namespace status params
func (o *ChangeNamespaceStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change namespace status params
func (o *ChangeNamespaceStatusParams) WithBody(body *basicclientmodels.NamespaceStatusUpdate) *ChangeNamespaceStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change namespace status params
func (o *ChangeNamespaceStatusParams) SetBody(body *basicclientmodels.NamespaceStatusUpdate) {
	o.Body = body
}

// WithNamespace adds the namespace to the change namespace status params
func (o *ChangeNamespaceStatusParams) WithNamespace(namespace string) *ChangeNamespaceStatusParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the change namespace status params
func (o *ChangeNamespaceStatusParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeNamespaceStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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