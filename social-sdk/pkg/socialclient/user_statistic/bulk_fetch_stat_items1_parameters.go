// Code generated by go-swagger; DO NOT EDIT.

package user_statistic

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

// NewBulkFetchStatItems1Params creates a new BulkFetchStatItems1Params object
// with the default values initialized.
func NewBulkFetchStatItems1Params() *BulkFetchStatItems1Params {
	var ()
	return &BulkFetchStatItems1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewBulkFetchStatItems1ParamsWithTimeout creates a new BulkFetchStatItems1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewBulkFetchStatItems1ParamsWithTimeout(timeout time.Duration) *BulkFetchStatItems1Params {
	var ()
	return &BulkFetchStatItems1Params{

		timeout: timeout,
	}
}

// NewBulkFetchStatItems1ParamsWithContext creates a new BulkFetchStatItems1Params object
// with the default values initialized, and the ability to set a context for a request
func NewBulkFetchStatItems1ParamsWithContext(ctx context.Context) *BulkFetchStatItems1Params {
	var ()
	return &BulkFetchStatItems1Params{

		Context: ctx,
	}
}

// NewBulkFetchStatItems1ParamsWithHTTPClient creates a new BulkFetchStatItems1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBulkFetchStatItems1ParamsWithHTTPClient(client *http.Client) *BulkFetchStatItems1Params {
	var ()
	return &BulkFetchStatItems1Params{
		HTTPClient: client,
	}
}

/*BulkFetchStatItems1Params contains all the parameters to send to the API endpoint
for the bulk fetch stat items 1 operation typically these are written to a http.Request
*/
type BulkFetchStatItems1Params struct {

	/*Namespace
	  namespace

	*/
	Namespace string
	/*StatCode
	  stat code

	*/
	StatCode string
	/*UserIds
	  comma separated user Ids

	*/
	UserIds string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the bulk fetch stat items 1 params
func (o *BulkFetchStatItems1Params) WithTimeout(timeout time.Duration) *BulkFetchStatItems1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk fetch stat items 1 params
func (o *BulkFetchStatItems1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk fetch stat items 1 params
func (o *BulkFetchStatItems1Params) WithContext(ctx context.Context) *BulkFetchStatItems1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk fetch stat items 1 params
func (o *BulkFetchStatItems1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk fetch stat items 1 params
func (o *BulkFetchStatItems1Params) WithHTTPClient(client *http.Client) *BulkFetchStatItems1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk fetch stat items 1 params
func (o *BulkFetchStatItems1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the bulk fetch stat items 1 params
func (o *BulkFetchStatItems1Params) WithNamespace(namespace string) *BulkFetchStatItems1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the bulk fetch stat items 1 params
func (o *BulkFetchStatItems1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithStatCode adds the statCode to the bulk fetch stat items 1 params
func (o *BulkFetchStatItems1Params) WithStatCode(statCode string) *BulkFetchStatItems1Params {
	o.SetStatCode(statCode)
	return o
}

// SetStatCode adds the statCode to the bulk fetch stat items 1 params
func (o *BulkFetchStatItems1Params) SetStatCode(statCode string) {
	o.StatCode = statCode
}

// WithUserIds adds the userIds to the bulk fetch stat items 1 params
func (o *BulkFetchStatItems1Params) WithUserIds(userIds string) *BulkFetchStatItems1Params {
	o.SetUserIds(userIds)
	return o
}

// SetUserIds adds the userIds to the bulk fetch stat items 1 params
func (o *BulkFetchStatItems1Params) SetUserIds(userIds string) {
	o.UserIds = userIds
}

// WriteToRequest writes these params to a swagger request
func (o *BulkFetchStatItems1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// query param statCode
	qrStatCode := o.StatCode
	qStatCode := qrStatCode
	if qStatCode != "" {
		if err := r.SetQueryParam("statCode", qStatCode); err != nil {
			return err
		}
	}

	// query param userIds
	qrUserIds := o.UserIds
	qUserIds := qrUserIds
	if qUserIds != "" {
		if err := r.SetQueryParam("userIds", qUserIds); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}