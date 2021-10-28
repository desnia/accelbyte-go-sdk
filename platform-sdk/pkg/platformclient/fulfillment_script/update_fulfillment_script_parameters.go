// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_script

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

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// NewUpdateFulfillmentScriptParams creates a new UpdateFulfillmentScriptParams object
// with the default values initialized.
func NewUpdateFulfillmentScriptParams() *UpdateFulfillmentScriptParams {
	var ()
	return &UpdateFulfillmentScriptParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFulfillmentScriptParamsWithTimeout creates a new UpdateFulfillmentScriptParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateFulfillmentScriptParamsWithTimeout(timeout time.Duration) *UpdateFulfillmentScriptParams {
	var ()
	return &UpdateFulfillmentScriptParams{

		timeout: timeout,
	}
}

// NewUpdateFulfillmentScriptParamsWithContext creates a new UpdateFulfillmentScriptParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateFulfillmentScriptParamsWithContext(ctx context.Context) *UpdateFulfillmentScriptParams {
	var ()
	return &UpdateFulfillmentScriptParams{

		Context: ctx,
	}
}

// NewUpdateFulfillmentScriptParamsWithHTTPClient creates a new UpdateFulfillmentScriptParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateFulfillmentScriptParamsWithHTTPClient(client *http.Client) *UpdateFulfillmentScriptParams {
	var ()
	return &UpdateFulfillmentScriptParams{
		HTTPClient: client,
	}
}

/*UpdateFulfillmentScriptParams contains all the parameters to send to the API endpoint
for the update fulfillment script operation typically these are written to a http.Request
*/
type UpdateFulfillmentScriptParams struct {

	/*Body*/
	Body *platformclientmodels.FulfillmentScriptUpdate
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update fulfillment script params
func (o *UpdateFulfillmentScriptParams) WithTimeout(timeout time.Duration) *UpdateFulfillmentScriptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update fulfillment script params
func (o *UpdateFulfillmentScriptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update fulfillment script params
func (o *UpdateFulfillmentScriptParams) WithContext(ctx context.Context) *UpdateFulfillmentScriptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update fulfillment script params
func (o *UpdateFulfillmentScriptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update fulfillment script params
func (o *UpdateFulfillmentScriptParams) WithHTTPClient(client *http.Client) *UpdateFulfillmentScriptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update fulfillment script params
func (o *UpdateFulfillmentScriptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update fulfillment script params
func (o *UpdateFulfillmentScriptParams) WithBody(body *platformclientmodels.FulfillmentScriptUpdate) *UpdateFulfillmentScriptParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update fulfillment script params
func (o *UpdateFulfillmentScriptParams) SetBody(body *platformclientmodels.FulfillmentScriptUpdate) {
	o.Body = body
}

// WithID adds the id to the update fulfillment script params
func (o *UpdateFulfillmentScriptParams) WithID(id string) *UpdateFulfillmentScriptParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update fulfillment script params
func (o *UpdateFulfillmentScriptParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFulfillmentScriptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
