// Code generated by go-swagger; DO NOT EDIT.

package matchmaking

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

	"github.com/AccelByte/accelbyte-go-sdk/matchmaking-sdk/pkg/matchmakingclientmodels"
)

// NewCreateChannelHandlerParams creates a new CreateChannelHandlerParams object
// with the default values initialized.
func NewCreateChannelHandlerParams() *CreateChannelHandlerParams {
	var ()
	return &CreateChannelHandlerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateChannelHandlerParamsWithTimeout creates a new CreateChannelHandlerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateChannelHandlerParamsWithTimeout(timeout time.Duration) *CreateChannelHandlerParams {
	var ()
	return &CreateChannelHandlerParams{

		timeout: timeout,
	}
}

// NewCreateChannelHandlerParamsWithContext creates a new CreateChannelHandlerParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateChannelHandlerParamsWithContext(ctx context.Context) *CreateChannelHandlerParams {
	var ()
	return &CreateChannelHandlerParams{

		Context: ctx,
	}
}

// NewCreateChannelHandlerParamsWithHTTPClient creates a new CreateChannelHandlerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateChannelHandlerParamsWithHTTPClient(client *http.Client) *CreateChannelHandlerParams {
	var ()
	return &CreateChannelHandlerParams{
		HTTPClient: client,
	}
}

/*CreateChannelHandlerParams contains all the parameters to send to the API endpoint
for the create channel handler operation typically these are written to a http.Request
*/
type CreateChannelHandlerParams struct {

	/*Body*/
	Body *matchmakingclientmodels.ModelsChannelRequest
	/*Namespace
	  namespace of the game

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create channel handler params
func (o *CreateChannelHandlerParams) WithTimeout(timeout time.Duration) *CreateChannelHandlerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create channel handler params
func (o *CreateChannelHandlerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create channel handler params
func (o *CreateChannelHandlerParams) WithContext(ctx context.Context) *CreateChannelHandlerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create channel handler params
func (o *CreateChannelHandlerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create channel handler params
func (o *CreateChannelHandlerParams) WithHTTPClient(client *http.Client) *CreateChannelHandlerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create channel handler params
func (o *CreateChannelHandlerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create channel handler params
func (o *CreateChannelHandlerParams) WithBody(body *matchmakingclientmodels.ModelsChannelRequest) *CreateChannelHandlerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create channel handler params
func (o *CreateChannelHandlerParams) SetBody(body *matchmakingclientmodels.ModelsChannelRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the create channel handler params
func (o *CreateChannelHandlerParams) WithNamespace(namespace string) *CreateChannelHandlerParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the create channel handler params
func (o *CreateChannelHandlerParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *CreateChannelHandlerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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