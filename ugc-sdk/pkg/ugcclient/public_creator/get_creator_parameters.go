// Code generated by go-swagger; DO NOT EDIT.

package public_creator

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

// NewGetCreatorParams creates a new GetCreatorParams object
// with the default values initialized.
func NewGetCreatorParams() *GetCreatorParams {
	var ()
	return &GetCreatorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCreatorParamsWithTimeout creates a new GetCreatorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCreatorParamsWithTimeout(timeout time.Duration) *GetCreatorParams {
	var ()
	return &GetCreatorParams{

		timeout: timeout,
	}
}

// NewGetCreatorParamsWithContext creates a new GetCreatorParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCreatorParamsWithContext(ctx context.Context) *GetCreatorParams {
	var ()
	return &GetCreatorParams{

		Context: ctx,
	}
}

// NewGetCreatorParamsWithHTTPClient creates a new GetCreatorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCreatorParamsWithHTTPClient(client *http.Client) *GetCreatorParams {
	var ()
	return &GetCreatorParams{
		HTTPClient: client,
	}
}

/*GetCreatorParams contains all the parameters to send to the API endpoint
for the get creator operation typically these are written to a http.Request
*/
type GetCreatorParams struct {

	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*UserID
	  user ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get creator params
func (o *GetCreatorParams) WithTimeout(timeout time.Duration) *GetCreatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get creator params
func (o *GetCreatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get creator params
func (o *GetCreatorParams) WithContext(ctx context.Context) *GetCreatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get creator params
func (o *GetCreatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get creator params
func (o *GetCreatorParams) WithHTTPClient(client *http.Client) *GetCreatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get creator params
func (o *GetCreatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get creator params
func (o *GetCreatorParams) WithNamespace(namespace string) *GetCreatorParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get creator params
func (o *GetCreatorParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the get creator params
func (o *GetCreatorParams) WithUserID(userID string) *GetCreatorParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get creator params
func (o *GetCreatorParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCreatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
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
