// Code generated by go-swagger; DO NOT EDIT.

package order

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

// NewGetOrderParams creates a new GetOrderParams object
// with the default values initialized.
func NewGetOrderParams() *GetOrderParams {
	var ()
	return &GetOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrderParamsWithTimeout creates a new GetOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrderParamsWithTimeout(timeout time.Duration) *GetOrderParams {
	var ()
	return &GetOrderParams{

		timeout: timeout,
	}
}

// NewGetOrderParamsWithContext creates a new GetOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrderParamsWithContext(ctx context.Context) *GetOrderParams {
	var ()
	return &GetOrderParams{

		Context: ctx,
	}
}

// NewGetOrderParamsWithHTTPClient creates a new GetOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrderParamsWithHTTPClient(client *http.Client) *GetOrderParams {
	var ()
	return &GetOrderParams{
		HTTPClient: client,
	}
}

/*GetOrderParams contains all the parameters to send to the API endpoint
for the get order operation typically these are written to a http.Request
*/
type GetOrderParams struct {

	/*Namespace*/
	Namespace string
	/*OrderNo*/
	OrderNo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get order params
func (o *GetOrderParams) WithTimeout(timeout time.Duration) *GetOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get order params
func (o *GetOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get order params
func (o *GetOrderParams) WithContext(ctx context.Context) *GetOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get order params
func (o *GetOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get order params
func (o *GetOrderParams) WithHTTPClient(client *http.Client) *GetOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get order params
func (o *GetOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get order params
func (o *GetOrderParams) WithNamespace(namespace string) *GetOrderParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get order params
func (o *GetOrderParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOrderNo adds the orderNo to the get order params
func (o *GetOrderParams) WithOrderNo(orderNo string) *GetOrderParams {
	o.SetOrderNo(orderNo)
	return o
}

// SetOrderNo adds the orderNo to the get order params
func (o *GetOrderParams) SetOrderNo(orderNo string) {
	o.OrderNo = orderNo
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param orderNo
	if err := r.SetPathParam("orderNo", o.OrderNo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}