// Code generated by go-swagger; DO NOT EDIT.

package admin_player_record

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

// NewAdminPostPlayerRecordHandlerV1Params creates a new AdminPostPlayerRecordHandlerV1Params object
// with the default values initialized.
func NewAdminPostPlayerRecordHandlerV1Params() *AdminPostPlayerRecordHandlerV1Params {
	var ()
	return &AdminPostPlayerRecordHandlerV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminPostPlayerRecordHandlerV1ParamsWithTimeout creates a new AdminPostPlayerRecordHandlerV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminPostPlayerRecordHandlerV1ParamsWithTimeout(timeout time.Duration) *AdminPostPlayerRecordHandlerV1Params {
	var ()
	return &AdminPostPlayerRecordHandlerV1Params{

		timeout: timeout,
	}
}

// NewAdminPostPlayerRecordHandlerV1ParamsWithContext creates a new AdminPostPlayerRecordHandlerV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminPostPlayerRecordHandlerV1ParamsWithContext(ctx context.Context) *AdminPostPlayerRecordHandlerV1Params {
	var ()
	return &AdminPostPlayerRecordHandlerV1Params{

		Context: ctx,
	}
}

// NewAdminPostPlayerRecordHandlerV1ParamsWithHTTPClient creates a new AdminPostPlayerRecordHandlerV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminPostPlayerRecordHandlerV1ParamsWithHTTPClient(client *http.Client) *AdminPostPlayerRecordHandlerV1Params {
	var ()
	return &AdminPostPlayerRecordHandlerV1Params{
		HTTPClient: client,
	}
}

/*AdminPostPlayerRecordHandlerV1Params contains all the parameters to send to the API endpoint
for the admin post player record handler v1 operation typically these are written to a http.Request
*/
type AdminPostPlayerRecordHandlerV1Params struct {

	/*Body*/
	Body cloudsaveclientmodels.ModelsPlayerRecordRequest
	/*Key
	  key of record

	*/
	Key string
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*UserID
	  user ID who own the record

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin post player record handler v1 params
func (o *AdminPostPlayerRecordHandlerV1Params) WithTimeout(timeout time.Duration) *AdminPostPlayerRecordHandlerV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin post player record handler v1 params
func (o *AdminPostPlayerRecordHandlerV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin post player record handler v1 params
func (o *AdminPostPlayerRecordHandlerV1Params) WithContext(ctx context.Context) *AdminPostPlayerRecordHandlerV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin post player record handler v1 params
func (o *AdminPostPlayerRecordHandlerV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin post player record handler v1 params
func (o *AdminPostPlayerRecordHandlerV1Params) WithHTTPClient(client *http.Client) *AdminPostPlayerRecordHandlerV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin post player record handler v1 params
func (o *AdminPostPlayerRecordHandlerV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin post player record handler v1 params
func (o *AdminPostPlayerRecordHandlerV1Params) WithBody(body cloudsaveclientmodels.ModelsPlayerRecordRequest) *AdminPostPlayerRecordHandlerV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin post player record handler v1 params
func (o *AdminPostPlayerRecordHandlerV1Params) SetBody(body cloudsaveclientmodels.ModelsPlayerRecordRequest) {
	o.Body = body
}

// WithKey adds the key to the admin post player record handler v1 params
func (o *AdminPostPlayerRecordHandlerV1Params) WithKey(key string) *AdminPostPlayerRecordHandlerV1Params {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the admin post player record handler v1 params
func (o *AdminPostPlayerRecordHandlerV1Params) SetKey(key string) {
	o.Key = key
}

// WithNamespace adds the namespace to the admin post player record handler v1 params
func (o *AdminPostPlayerRecordHandlerV1Params) WithNamespace(namespace string) *AdminPostPlayerRecordHandlerV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin post player record handler v1 params
func (o *AdminPostPlayerRecordHandlerV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the admin post player record handler v1 params
func (o *AdminPostPlayerRecordHandlerV1Params) WithUserID(userID string) *AdminPostPlayerRecordHandlerV1Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the admin post player record handler v1 params
func (o *AdminPostPlayerRecordHandlerV1Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminPostPlayerRecordHandlerV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param userID
	if err := r.SetPathParam("userID", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}