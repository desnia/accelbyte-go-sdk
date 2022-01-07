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
	"github.com/go-openapi/swag"
)

// NewAdminRetrievePlayerRecordsParams creates a new AdminRetrievePlayerRecordsParams object
// with the default values initialized.
func NewAdminRetrievePlayerRecordsParams() *AdminRetrievePlayerRecordsParams {
	var ()
	return &AdminRetrievePlayerRecordsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminRetrievePlayerRecordsParamsWithTimeout creates a new AdminRetrievePlayerRecordsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminRetrievePlayerRecordsParamsWithTimeout(timeout time.Duration) *AdminRetrievePlayerRecordsParams {
	var ()
	return &AdminRetrievePlayerRecordsParams{

		timeout: timeout,
	}
}

// NewAdminRetrievePlayerRecordsParamsWithContext creates a new AdminRetrievePlayerRecordsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminRetrievePlayerRecordsParamsWithContext(ctx context.Context) *AdminRetrievePlayerRecordsParams {
	var ()
	return &AdminRetrievePlayerRecordsParams{

		Context: ctx,
	}
}

// NewAdminRetrievePlayerRecordsParamsWithHTTPClient creates a new AdminRetrievePlayerRecordsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminRetrievePlayerRecordsParamsWithHTTPClient(client *http.Client) *AdminRetrievePlayerRecordsParams {
	var ()
	return &AdminRetrievePlayerRecordsParams{
		HTTPClient: client,
	}
}

/*AdminRetrievePlayerRecordsParams contains all the parameters to send to the API endpoint
for the admin retrieve player records operation typically these are written to a http.Request
*/
type AdminRetrievePlayerRecordsParams struct {

	/*Limit
	  limit

	*/
	Limit *int64
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*Offset
	  offset

	*/
	Offset *int64
	/*UserID
	  user ID who own the record

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin retrieve player records params
func (o *AdminRetrievePlayerRecordsParams) WithTimeout(timeout time.Duration) *AdminRetrievePlayerRecordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin retrieve player records params
func (o *AdminRetrievePlayerRecordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin retrieve player records params
func (o *AdminRetrievePlayerRecordsParams) WithContext(ctx context.Context) *AdminRetrievePlayerRecordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin retrieve player records params
func (o *AdminRetrievePlayerRecordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin retrieve player records params
func (o *AdminRetrievePlayerRecordsParams) WithHTTPClient(client *http.Client) *AdminRetrievePlayerRecordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin retrieve player records params
func (o *AdminRetrievePlayerRecordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the admin retrieve player records params
func (o *AdminRetrievePlayerRecordsParams) WithLimit(limit *int64) *AdminRetrievePlayerRecordsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the admin retrieve player records params
func (o *AdminRetrievePlayerRecordsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the admin retrieve player records params
func (o *AdminRetrievePlayerRecordsParams) WithNamespace(namespace string) *AdminRetrievePlayerRecordsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin retrieve player records params
func (o *AdminRetrievePlayerRecordsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the admin retrieve player records params
func (o *AdminRetrievePlayerRecordsParams) WithOffset(offset *int64) *AdminRetrievePlayerRecordsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the admin retrieve player records params
func (o *AdminRetrievePlayerRecordsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithUserID adds the userID to the admin retrieve player records params
func (o *AdminRetrievePlayerRecordsParams) WithUserID(userID string) *AdminRetrievePlayerRecordsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the admin retrieve player records params
func (o *AdminRetrievePlayerRecordsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminRetrievePlayerRecordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

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