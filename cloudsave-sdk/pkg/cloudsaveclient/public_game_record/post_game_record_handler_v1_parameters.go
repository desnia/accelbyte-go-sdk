// Code generated by go-swagger; DO NOT EDIT.

package public_game_record

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

	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclientmodels"
)

// NewPostGameRecordHandlerV1Params creates a new PostGameRecordHandlerV1Params object
// with the default values initialized.
func NewPostGameRecordHandlerV1Params() *PostGameRecordHandlerV1Params {
	var ()
	return &PostGameRecordHandlerV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostGameRecordHandlerV1ParamsWithTimeout creates a new PostGameRecordHandlerV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostGameRecordHandlerV1ParamsWithTimeout(timeout time.Duration) *PostGameRecordHandlerV1Params {
	var ()
	return &PostGameRecordHandlerV1Params{

		timeout: timeout,
	}
}

// NewPostGameRecordHandlerV1ParamsWithContext creates a new PostGameRecordHandlerV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewPostGameRecordHandlerV1ParamsWithContext(ctx context.Context) *PostGameRecordHandlerV1Params {
	var ()
	return &PostGameRecordHandlerV1Params{

		Context: ctx,
	}
}

// NewPostGameRecordHandlerV1ParamsWithHTTPClient creates a new PostGameRecordHandlerV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostGameRecordHandlerV1ParamsWithHTTPClient(client *http.Client) *PostGameRecordHandlerV1Params {
	var ()
	return &PostGameRecordHandlerV1Params{
		HTTPClient: client,
	}
}

/*PostGameRecordHandlerV1Params contains all the parameters to send to the API endpoint
for the post game record handler v1 operation typically these are written to a http.Request
*/
type PostGameRecordHandlerV1Params struct {

	/*Body*/
	Body cloudsaveclientmodels.ModelsGameRecordRequest
	/*Key
	  key of record

	*/
	Key string
	/*Namespace
	  namespace of the game

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post game record handler v1 params
func (o *PostGameRecordHandlerV1Params) WithTimeout(timeout time.Duration) *PostGameRecordHandlerV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post game record handler v1 params
func (o *PostGameRecordHandlerV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post game record handler v1 params
func (o *PostGameRecordHandlerV1Params) WithContext(ctx context.Context) *PostGameRecordHandlerV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post game record handler v1 params
func (o *PostGameRecordHandlerV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post game record handler v1 params
func (o *PostGameRecordHandlerV1Params) WithHTTPClient(client *http.Client) *PostGameRecordHandlerV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post game record handler v1 params
func (o *PostGameRecordHandlerV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post game record handler v1 params
func (o *PostGameRecordHandlerV1Params) WithBody(body cloudsaveclientmodels.ModelsGameRecordRequest) *PostGameRecordHandlerV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post game record handler v1 params
func (o *PostGameRecordHandlerV1Params) SetBody(body cloudsaveclientmodels.ModelsGameRecordRequest) {
	o.Body = body
}

// WithKey adds the key to the post game record handler v1 params
func (o *PostGameRecordHandlerV1Params) WithKey(key string) *PostGameRecordHandlerV1Params {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the post game record handler v1 params
func (o *PostGameRecordHandlerV1Params) SetKey(key string) {
	o.Key = key
}

// WithNamespace adds the namespace to the post game record handler v1 params
func (o *PostGameRecordHandlerV1Params) WithNamespace(namespace string) *PostGameRecordHandlerV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the post game record handler v1 params
func (o *PostGameRecordHandlerV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *PostGameRecordHandlerV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
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