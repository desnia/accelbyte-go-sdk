// Code generated by go-swagger; DO NOT EDIT.

package player

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

// NewAdminGetPlayerBlockedByPlayersV1Params creates a new AdminGetPlayerBlockedByPlayersV1Params object
// with the default values initialized.
func NewAdminGetPlayerBlockedByPlayersV1Params() *AdminGetPlayerBlockedByPlayersV1Params {
	var ()
	return &AdminGetPlayerBlockedByPlayersV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetPlayerBlockedByPlayersV1ParamsWithTimeout creates a new AdminGetPlayerBlockedByPlayersV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminGetPlayerBlockedByPlayersV1ParamsWithTimeout(timeout time.Duration) *AdminGetPlayerBlockedByPlayersV1Params {
	var ()
	return &AdminGetPlayerBlockedByPlayersV1Params{

		timeout: timeout,
	}
}

// NewAdminGetPlayerBlockedByPlayersV1ParamsWithContext creates a new AdminGetPlayerBlockedByPlayersV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminGetPlayerBlockedByPlayersV1ParamsWithContext(ctx context.Context) *AdminGetPlayerBlockedByPlayersV1Params {
	var ()
	return &AdminGetPlayerBlockedByPlayersV1Params{

		Context: ctx,
	}
}

// NewAdminGetPlayerBlockedByPlayersV1ParamsWithHTTPClient creates a new AdminGetPlayerBlockedByPlayersV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminGetPlayerBlockedByPlayersV1ParamsWithHTTPClient(client *http.Client) *AdminGetPlayerBlockedByPlayersV1Params {
	var ()
	return &AdminGetPlayerBlockedByPlayersV1Params{
		HTTPClient: client,
	}
}

/*AdminGetPlayerBlockedByPlayersV1Params contains all the parameters to send to the API endpoint
for the admin get player blocked by players v1 operation typically these are written to a http.Request
*/
type AdminGetPlayerBlockedByPlayersV1Params struct {

	/*Namespace
	  namespace

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

// WithTimeout adds the timeout to the admin get player blocked by players v1 params
func (o *AdminGetPlayerBlockedByPlayersV1Params) WithTimeout(timeout time.Duration) *AdminGetPlayerBlockedByPlayersV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get player blocked by players v1 params
func (o *AdminGetPlayerBlockedByPlayersV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get player blocked by players v1 params
func (o *AdminGetPlayerBlockedByPlayersV1Params) WithContext(ctx context.Context) *AdminGetPlayerBlockedByPlayersV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get player blocked by players v1 params
func (o *AdminGetPlayerBlockedByPlayersV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin get player blocked by players v1 params
func (o *AdminGetPlayerBlockedByPlayersV1Params) WithHTTPClient(client *http.Client) *AdminGetPlayerBlockedByPlayersV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get player blocked by players v1 params
func (o *AdminGetPlayerBlockedByPlayersV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the admin get player blocked by players v1 params
func (o *AdminGetPlayerBlockedByPlayersV1Params) WithNamespace(namespace string) *AdminGetPlayerBlockedByPlayersV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin get player blocked by players v1 params
func (o *AdminGetPlayerBlockedByPlayersV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the admin get player blocked by players v1 params
func (o *AdminGetPlayerBlockedByPlayersV1Params) WithUserID(userID string) *AdminGetPlayerBlockedByPlayersV1Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the admin get player blocked by players v1 params
func (o *AdminGetPlayerBlockedByPlayersV1Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetPlayerBlockedByPlayersV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
