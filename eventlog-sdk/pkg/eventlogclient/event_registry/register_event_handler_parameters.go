// Code generated by go-swagger; DO NOT EDIT.

package event_registry

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

	"github.com/AccelByte/accelbyte-go-sdk/eventlog-sdk/pkg/eventlogclientmodels"
)

// NewRegisterEventHandlerParams creates a new RegisterEventHandlerParams object
// with the default values initialized.
func NewRegisterEventHandlerParams() *RegisterEventHandlerParams {
	var ()
	return &RegisterEventHandlerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterEventHandlerParamsWithTimeout creates a new RegisterEventHandlerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRegisterEventHandlerParamsWithTimeout(timeout time.Duration) *RegisterEventHandlerParams {
	var ()
	return &RegisterEventHandlerParams{

		timeout: timeout,
	}
}

// NewRegisterEventHandlerParamsWithContext creates a new RegisterEventHandlerParams object
// with the default values initialized, and the ability to set a context for a request
func NewRegisterEventHandlerParamsWithContext(ctx context.Context) *RegisterEventHandlerParams {
	var ()
	return &RegisterEventHandlerParams{

		Context: ctx,
	}
}

// NewRegisterEventHandlerParamsWithHTTPClient creates a new RegisterEventHandlerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRegisterEventHandlerParamsWithHTTPClient(client *http.Client) *RegisterEventHandlerParams {
	var ()
	return &RegisterEventHandlerParams{
		HTTPClient: client,
	}
}

/*RegisterEventHandlerParams contains all the parameters to send to the API endpoint
for the register event handler operation typically these are written to a http.Request
*/
type RegisterEventHandlerParams struct {

	/*Body*/
	Body *eventlogclientmodels.ModelsEventRegistry

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the register event handler params
func (o *RegisterEventHandlerParams) WithTimeout(timeout time.Duration) *RegisterEventHandlerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register event handler params
func (o *RegisterEventHandlerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register event handler params
func (o *RegisterEventHandlerParams) WithContext(ctx context.Context) *RegisterEventHandlerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register event handler params
func (o *RegisterEventHandlerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register event handler params
func (o *RegisterEventHandlerParams) WithHTTPClient(client *http.Client) *RegisterEventHandlerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register event handler params
func (o *RegisterEventHandlerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the register event handler params
func (o *RegisterEventHandlerParams) WithBody(body *eventlogclientmodels.ModelsEventRegistry) *RegisterEventHandlerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the register event handler params
func (o *RegisterEventHandlerParams) SetBody(body *eventlogclientmodels.ModelsEventRegistry) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterEventHandlerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}