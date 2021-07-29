// Code generated by go-swagger; DO NOT EDIT.

package payment_config

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

// NewTestWxPayConfigByIDParams creates a new TestWxPayConfigByIDParams object
// with the default values initialized.
func NewTestWxPayConfigByIDParams() *TestWxPayConfigByIDParams {
	var ()
	return &TestWxPayConfigByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestWxPayConfigByIDParamsWithTimeout creates a new TestWxPayConfigByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestWxPayConfigByIDParamsWithTimeout(timeout time.Duration) *TestWxPayConfigByIDParams {
	var ()
	return &TestWxPayConfigByIDParams{

		timeout: timeout,
	}
}

// NewTestWxPayConfigByIDParamsWithContext creates a new TestWxPayConfigByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestWxPayConfigByIDParamsWithContext(ctx context.Context) *TestWxPayConfigByIDParams {
	var ()
	return &TestWxPayConfigByIDParams{

		Context: ctx,
	}
}

// NewTestWxPayConfigByIDParamsWithHTTPClient creates a new TestWxPayConfigByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestWxPayConfigByIDParamsWithHTTPClient(client *http.Client) *TestWxPayConfigByIDParams {
	var ()
	return &TestWxPayConfigByIDParams{
		HTTPClient: client,
	}
}

/*TestWxPayConfigByIDParams contains all the parameters to send to the API endpoint
for the test wx pay config by Id operation typically these are written to a http.Request
*/
type TestWxPayConfigByIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test wx pay config by Id params
func (o *TestWxPayConfigByIDParams) WithTimeout(timeout time.Duration) *TestWxPayConfigByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test wx pay config by Id params
func (o *TestWxPayConfigByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test wx pay config by Id params
func (o *TestWxPayConfigByIDParams) WithContext(ctx context.Context) *TestWxPayConfigByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test wx pay config by Id params
func (o *TestWxPayConfigByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test wx pay config by Id params
func (o *TestWxPayConfigByIDParams) WithHTTPClient(client *http.Client) *TestWxPayConfigByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test wx pay config by Id params
func (o *TestWxPayConfigByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the test wx pay config by Id params
func (o *TestWxPayConfigByIDParams) WithID(id string) *TestWxPayConfigByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the test wx pay config by Id params
func (o *TestWxPayConfigByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TestWxPayConfigByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}