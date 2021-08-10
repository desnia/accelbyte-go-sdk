// Code generated by go-swagger; DO NOT EDIT.

package party

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

// NewPublicGetPartyDataV1Params creates a new PublicGetPartyDataV1Params object
// with the default values initialized.
func NewPublicGetPartyDataV1Params() *PublicGetPartyDataV1Params {
	var ()
	return &PublicGetPartyDataV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicGetPartyDataV1ParamsWithTimeout creates a new PublicGetPartyDataV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicGetPartyDataV1ParamsWithTimeout(timeout time.Duration) *PublicGetPartyDataV1Params {
	var ()
	return &PublicGetPartyDataV1Params{

		timeout: timeout,
	}
}

// NewPublicGetPartyDataV1ParamsWithContext creates a new PublicGetPartyDataV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewPublicGetPartyDataV1ParamsWithContext(ctx context.Context) *PublicGetPartyDataV1Params {
	var ()
	return &PublicGetPartyDataV1Params{

		Context: ctx,
	}
}

// NewPublicGetPartyDataV1ParamsWithHTTPClient creates a new PublicGetPartyDataV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicGetPartyDataV1ParamsWithHTTPClient(client *http.Client) *PublicGetPartyDataV1Params {
	var ()
	return &PublicGetPartyDataV1Params{
		HTTPClient: client,
	}
}

/*PublicGetPartyDataV1Params contains all the parameters to send to the API endpoint
for the public get party data v1 operation typically these are written to a http.Request
*/
type PublicGetPartyDataV1Params struct {

	/*Namespace
	  namespace

	*/
	Namespace string
	/*PartyID
	  Party ID, should follow UUID version 4 without hyphen

	*/
	PartyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public get party data v1 params
func (o *PublicGetPartyDataV1Params) WithTimeout(timeout time.Duration) *PublicGetPartyDataV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public get party data v1 params
func (o *PublicGetPartyDataV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public get party data v1 params
func (o *PublicGetPartyDataV1Params) WithContext(ctx context.Context) *PublicGetPartyDataV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public get party data v1 params
func (o *PublicGetPartyDataV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public get party data v1 params
func (o *PublicGetPartyDataV1Params) WithHTTPClient(client *http.Client) *PublicGetPartyDataV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public get party data v1 params
func (o *PublicGetPartyDataV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the public get party data v1 params
func (o *PublicGetPartyDataV1Params) WithNamespace(namespace string) *PublicGetPartyDataV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public get party data v1 params
func (o *PublicGetPartyDataV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPartyID adds the partyID to the public get party data v1 params
func (o *PublicGetPartyDataV1Params) WithPartyID(partyID string) *PublicGetPartyDataV1Params {
	o.SetPartyID(partyID)
	return o
}

// SetPartyID adds the partyId to the public get party data v1 params
func (o *PublicGetPartyDataV1Params) SetPartyID(partyID string) {
	o.PartyID = partyID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicGetPartyDataV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param partyId
	if err := r.SetPathParam("partyId", o.PartyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
