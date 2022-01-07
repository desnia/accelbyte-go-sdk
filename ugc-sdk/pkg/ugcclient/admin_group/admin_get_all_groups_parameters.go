// Code generated by go-swagger; DO NOT EDIT.

package admin_group

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

// NewAdminGetAllGroupsParams creates a new AdminGetAllGroupsParams object
// with the default values initialized.
func NewAdminGetAllGroupsParams() *AdminGetAllGroupsParams {
	var (
		limitDefault  = string("1000")
		offsetDefault = string("0")
	)
	return &AdminGetAllGroupsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetAllGroupsParamsWithTimeout creates a new AdminGetAllGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminGetAllGroupsParamsWithTimeout(timeout time.Duration) *AdminGetAllGroupsParams {
	var (
		limitDefault  = string("1000")
		offsetDefault = string("0")
	)
	return &AdminGetAllGroupsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewAdminGetAllGroupsParamsWithContext creates a new AdminGetAllGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminGetAllGroupsParamsWithContext(ctx context.Context) *AdminGetAllGroupsParams {
	var (
		limitDefault  = string("1000")
		offsetDefault = string("0")
	)
	return &AdminGetAllGroupsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewAdminGetAllGroupsParamsWithHTTPClient creates a new AdminGetAllGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminGetAllGroupsParamsWithHTTPClient(client *http.Client) *AdminGetAllGroupsParams {
	var (
		limitDefault  = string("1000")
		offsetDefault = string("0")
	)
	return &AdminGetAllGroupsParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*AdminGetAllGroupsParams contains all the parameters to send to the API endpoint
for the admin get all groups operation typically these are written to a http.Request
*/
type AdminGetAllGroupsParams struct {

	/*Limit
	  number of content per page

	*/
	Limit *string
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*Offset
	  the offset number to retrieve

	*/
	Offset *string
	/*UserID
	  user ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin get all groups params
func (o *AdminGetAllGroupsParams) WithTimeout(timeout time.Duration) *AdminGetAllGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get all groups params
func (o *AdminGetAllGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get all groups params
func (o *AdminGetAllGroupsParams) WithContext(ctx context.Context) *AdminGetAllGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get all groups params
func (o *AdminGetAllGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin get all groups params
func (o *AdminGetAllGroupsParams) WithHTTPClient(client *http.Client) *AdminGetAllGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get all groups params
func (o *AdminGetAllGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the admin get all groups params
func (o *AdminGetAllGroupsParams) WithLimit(limit *string) *AdminGetAllGroupsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the admin get all groups params
func (o *AdminGetAllGroupsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the admin get all groups params
func (o *AdminGetAllGroupsParams) WithNamespace(namespace string) *AdminGetAllGroupsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin get all groups params
func (o *AdminGetAllGroupsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the admin get all groups params
func (o *AdminGetAllGroupsParams) WithOffset(offset *string) *AdminGetAllGroupsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the admin get all groups params
func (o *AdminGetAllGroupsParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithUserID adds the userID to the admin get all groups params
func (o *AdminGetAllGroupsParams) WithUserID(userID string) *AdminGetAllGroupsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the admin get all groups params
func (o *AdminGetAllGroupsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetAllGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit string
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
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
		var qrOffset string
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
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