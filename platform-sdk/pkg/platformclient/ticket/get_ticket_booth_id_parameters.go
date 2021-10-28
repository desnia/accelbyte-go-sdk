// Code generated by go-swagger; DO NOT EDIT.

package ticket

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

// NewGetTicketBoothIDParams creates a new GetTicketBoothIDParams object
// with the default values initialized.
func NewGetTicketBoothIDParams() *GetTicketBoothIDParams {
	var ()
	return &GetTicketBoothIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTicketBoothIDParamsWithTimeout creates a new GetTicketBoothIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTicketBoothIDParamsWithTimeout(timeout time.Duration) *GetTicketBoothIDParams {
	var ()
	return &GetTicketBoothIDParams{

		timeout: timeout,
	}
}

// NewGetTicketBoothIDParamsWithContext creates a new GetTicketBoothIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTicketBoothIDParamsWithContext(ctx context.Context) *GetTicketBoothIDParams {
	var ()
	return &GetTicketBoothIDParams{

		Context: ctx,
	}
}

// NewGetTicketBoothIDParamsWithHTTPClient creates a new GetTicketBoothIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTicketBoothIDParamsWithHTTPClient(client *http.Client) *GetTicketBoothIDParams {
	var ()
	return &GetTicketBoothIDParams{
		HTTPClient: client,
	}
}

/*GetTicketBoothIDParams contains all the parameters to send to the API endpoint
for the get ticket booth ID operation typically these are written to a http.Request
*/
type GetTicketBoothIDParams struct {

	/*BoothName*/
	BoothName string
	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ticket booth ID params
func (o *GetTicketBoothIDParams) WithTimeout(timeout time.Duration) *GetTicketBoothIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ticket booth ID params
func (o *GetTicketBoothIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ticket booth ID params
func (o *GetTicketBoothIDParams) WithContext(ctx context.Context) *GetTicketBoothIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ticket booth ID params
func (o *GetTicketBoothIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ticket booth ID params
func (o *GetTicketBoothIDParams) WithHTTPClient(client *http.Client) *GetTicketBoothIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ticket booth ID params
func (o *GetTicketBoothIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBoothName adds the boothName to the get ticket booth ID params
func (o *GetTicketBoothIDParams) WithBoothName(boothName string) *GetTicketBoothIDParams {
	o.SetBoothName(boothName)
	return o
}

// SetBoothName adds the boothName to the get ticket booth ID params
func (o *GetTicketBoothIDParams) SetBoothName(boothName string) {
	o.BoothName = boothName
}

// WithNamespace adds the namespace to the get ticket booth ID params
func (o *GetTicketBoothIDParams) WithNamespace(namespace string) *GetTicketBoothIDParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get ticket booth ID params
func (o *GetTicketBoothIDParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetTicketBoothIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param boothName
	if err := r.SetPathParam("boothName", o.BoothName); err != nil {
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