// Code generated by go-swagger; DO NOT EDIT.

package user_profile

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

	"github.com/AccelByte/accelbyte-go-sdk/basic-sdk/pkg/basicclientmodels"
)

// NewPublicEnableUserProfileStatusParams creates a new PublicEnableUserProfileStatusParams object
// with the default values initialized.
func NewPublicEnableUserProfileStatusParams() *PublicEnableUserProfileStatusParams {
	var ()
	return &PublicEnableUserProfileStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicEnableUserProfileStatusParamsWithTimeout creates a new PublicEnableUserProfileStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicEnableUserProfileStatusParamsWithTimeout(timeout time.Duration) *PublicEnableUserProfileStatusParams {
	var ()
	return &PublicEnableUserProfileStatusParams{

		timeout: timeout,
	}
}

// NewPublicEnableUserProfileStatusParamsWithContext creates a new PublicEnableUserProfileStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicEnableUserProfileStatusParamsWithContext(ctx context.Context) *PublicEnableUserProfileStatusParams {
	var ()
	return &PublicEnableUserProfileStatusParams{

		Context: ctx,
	}
}

// NewPublicEnableUserProfileStatusParamsWithHTTPClient creates a new PublicEnableUserProfileStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicEnableUserProfileStatusParamsWithHTTPClient(client *http.Client) *PublicEnableUserProfileStatusParams {
	var ()
	return &PublicEnableUserProfileStatusParams{
		HTTPClient: client,
	}
}

/*PublicEnableUserProfileStatusParams contains all the parameters to send to the API endpoint
for the public enable user profile status operation typically these are written to a http.Request
*/
type PublicEnableUserProfileStatusParams struct {

	/*Body*/
	Body *basicclientmodels.UserProfileStatusUpdate
	/*Namespace
	  namespace, only accept alphabet and numeric

	*/
	Namespace string
	/*UserID
	  user's id, should follow UUID version 4 without hyphen

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public enable user profile status params
func (o *PublicEnableUserProfileStatusParams) WithTimeout(timeout time.Duration) *PublicEnableUserProfileStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public enable user profile status params
func (o *PublicEnableUserProfileStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public enable user profile status params
func (o *PublicEnableUserProfileStatusParams) WithContext(ctx context.Context) *PublicEnableUserProfileStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public enable user profile status params
func (o *PublicEnableUserProfileStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public enable user profile status params
func (o *PublicEnableUserProfileStatusParams) WithHTTPClient(client *http.Client) *PublicEnableUserProfileStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public enable user profile status params
func (o *PublicEnableUserProfileStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the public enable user profile status params
func (o *PublicEnableUserProfileStatusParams) WithBody(body *basicclientmodels.UserProfileStatusUpdate) *PublicEnableUserProfileStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the public enable user profile status params
func (o *PublicEnableUserProfileStatusParams) SetBody(body *basicclientmodels.UserProfileStatusUpdate) {
	o.Body = body
}

// WithNamespace adds the namespace to the public enable user profile status params
func (o *PublicEnableUserProfileStatusParams) WithNamespace(namespace string) *PublicEnableUserProfileStatusParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public enable user profile status params
func (o *PublicEnableUserProfileStatusParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the public enable user profile status params
func (o *PublicEnableUserProfileStatusParams) WithUserID(userID string) *PublicEnableUserProfileStatusParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public enable user profile status params
func (o *PublicEnableUserProfileStatusParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicEnableUserProfileStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
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