// Code generated by go-swagger; DO NOT EDIT.

package key_group

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
)

// NewGetKeyGroupParams creates a new GetKeyGroupParams object
// with the default values initialized.
func NewGetKeyGroupParams() *GetKeyGroupParams {
	var ()
	return &GetKeyGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKeyGroupParamsWithTimeout creates a new GetKeyGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKeyGroupParamsWithTimeout(timeout time.Duration) *GetKeyGroupParams {
	var ()
	return &GetKeyGroupParams{

		timeout: timeout,
	}
}

// NewGetKeyGroupParamsWithContext creates a new GetKeyGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKeyGroupParamsWithContext(ctx context.Context) *GetKeyGroupParams {
	var ()
	return &GetKeyGroupParams{

		Context: ctx,
	}
}

// NewGetKeyGroupParamsWithHTTPClient creates a new GetKeyGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKeyGroupParamsWithHTTPClient(client *http.Client) *GetKeyGroupParams {
	var ()
	return &GetKeyGroupParams{
		HTTPClient: client,
	}
}

/*GetKeyGroupParams contains all the parameters to send to the API endpoint
for the get key group operation typically these are written to a http.Request
*/
type GetKeyGroupParams struct {

	/*KeyGroupID*/
	KeyGroupID string
	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get key group params
func (o *GetKeyGroupParams) WithTimeout(timeout time.Duration) *GetKeyGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get key group params
func (o *GetKeyGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get key group params
func (o *GetKeyGroupParams) WithContext(ctx context.Context) *GetKeyGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get key group params
func (o *GetKeyGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get key group params
func (o *GetKeyGroupParams) WithHTTPClient(client *http.Client) *GetKeyGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get key group params
func (o *GetKeyGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyGroupID adds the keyGroupID to the get key group params
func (o *GetKeyGroupParams) WithKeyGroupID(keyGroupID string) *GetKeyGroupParams {
	o.SetKeyGroupID(keyGroupID)
	return o
}

// SetKeyGroupID adds the keyGroupId to the get key group params
func (o *GetKeyGroupParams) SetKeyGroupID(keyGroupID string) {
	o.KeyGroupID = keyGroupID
}

// WithNamespace adds the namespace to the get key group params
func (o *GetKeyGroupParams) WithNamespace(namespace string) *GetKeyGroupParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get key group params
func (o *GetKeyGroupParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetKeyGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param keyGroupId
	if err := r.SetPathParam("keyGroupId", o.KeyGroupID); err != nil {
		return err
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