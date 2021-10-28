// Code generated by go-swagger; DO NOT EDIT.

package entitlement

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

// NewPublicGetUserEntitlementOwnershipByItemIDParams creates a new PublicGetUserEntitlementOwnershipByItemIDParams object
// with the default values initialized.
func NewPublicGetUserEntitlementOwnershipByItemIDParams() *PublicGetUserEntitlementOwnershipByItemIDParams {
	var ()
	return &PublicGetUserEntitlementOwnershipByItemIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicGetUserEntitlementOwnershipByItemIDParamsWithTimeout creates a new PublicGetUserEntitlementOwnershipByItemIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicGetUserEntitlementOwnershipByItemIDParamsWithTimeout(timeout time.Duration) *PublicGetUserEntitlementOwnershipByItemIDParams {
	var ()
	return &PublicGetUserEntitlementOwnershipByItemIDParams{

		timeout: timeout,
	}
}

// NewPublicGetUserEntitlementOwnershipByItemIDParamsWithContext creates a new PublicGetUserEntitlementOwnershipByItemIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicGetUserEntitlementOwnershipByItemIDParamsWithContext(ctx context.Context) *PublicGetUserEntitlementOwnershipByItemIDParams {
	var ()
	return &PublicGetUserEntitlementOwnershipByItemIDParams{

		Context: ctx,
	}
}

// NewPublicGetUserEntitlementOwnershipByItemIDParamsWithHTTPClient creates a new PublicGetUserEntitlementOwnershipByItemIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicGetUserEntitlementOwnershipByItemIDParamsWithHTTPClient(client *http.Client) *PublicGetUserEntitlementOwnershipByItemIDParams {
	var ()
	return &PublicGetUserEntitlementOwnershipByItemIDParams{
		HTTPClient: client,
	}
}

/*PublicGetUserEntitlementOwnershipByItemIDParams contains all the parameters to send to the API endpoint
for the public get user entitlement ownership by item Id operation typically these are written to a http.Request
*/
type PublicGetUserEntitlementOwnershipByItemIDParams struct {

	/*EntitlementClazz*/
	EntitlementClazz *string
	/*ItemID*/
	ItemID string
	/*Namespace*/
	Namespace string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public get user entitlement ownership by item Id params
func (o *PublicGetUserEntitlementOwnershipByItemIDParams) WithTimeout(timeout time.Duration) *PublicGetUserEntitlementOwnershipByItemIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public get user entitlement ownership by item Id params
func (o *PublicGetUserEntitlementOwnershipByItemIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public get user entitlement ownership by item Id params
func (o *PublicGetUserEntitlementOwnershipByItemIDParams) WithContext(ctx context.Context) *PublicGetUserEntitlementOwnershipByItemIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public get user entitlement ownership by item Id params
func (o *PublicGetUserEntitlementOwnershipByItemIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public get user entitlement ownership by item Id params
func (o *PublicGetUserEntitlementOwnershipByItemIDParams) WithHTTPClient(client *http.Client) *PublicGetUserEntitlementOwnershipByItemIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public get user entitlement ownership by item Id params
func (o *PublicGetUserEntitlementOwnershipByItemIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntitlementClazz adds the entitlementClazz to the public get user entitlement ownership by item Id params
func (o *PublicGetUserEntitlementOwnershipByItemIDParams) WithEntitlementClazz(entitlementClazz *string) *PublicGetUserEntitlementOwnershipByItemIDParams {
	o.SetEntitlementClazz(entitlementClazz)
	return o
}

// SetEntitlementClazz adds the entitlementClazz to the public get user entitlement ownership by item Id params
func (o *PublicGetUserEntitlementOwnershipByItemIDParams) SetEntitlementClazz(entitlementClazz *string) {
	o.EntitlementClazz = entitlementClazz
}

// WithItemID adds the itemID to the public get user entitlement ownership by item Id params
func (o *PublicGetUserEntitlementOwnershipByItemIDParams) WithItemID(itemID string) *PublicGetUserEntitlementOwnershipByItemIDParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the public get user entitlement ownership by item Id params
func (o *PublicGetUserEntitlementOwnershipByItemIDParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WithNamespace adds the namespace to the public get user entitlement ownership by item Id params
func (o *PublicGetUserEntitlementOwnershipByItemIDParams) WithNamespace(namespace string) *PublicGetUserEntitlementOwnershipByItemIDParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public get user entitlement ownership by item Id params
func (o *PublicGetUserEntitlementOwnershipByItemIDParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the public get user entitlement ownership by item Id params
func (o *PublicGetUserEntitlementOwnershipByItemIDParams) WithUserID(userID string) *PublicGetUserEntitlementOwnershipByItemIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public get user entitlement ownership by item Id params
func (o *PublicGetUserEntitlementOwnershipByItemIDParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicGetUserEntitlementOwnershipByItemIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EntitlementClazz != nil {

		// query param entitlementClazz
		var qrEntitlementClazz string
		if o.EntitlementClazz != nil {
			qrEntitlementClazz = *o.EntitlementClazz
		}
		qEntitlementClazz := qrEntitlementClazz
		if qEntitlementClazz != "" {
			if err := r.SetQueryParam("entitlementClazz", qEntitlementClazz); err != nil {
				return err
			}
		}

	}

	// query param itemId
	qrItemID := o.ItemID
	qItemID := qrItemID
	if qItemID != "" {
		if err := r.SetQueryParam("itemId", qItemID); err != nil {
			return err
		}
	}

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
