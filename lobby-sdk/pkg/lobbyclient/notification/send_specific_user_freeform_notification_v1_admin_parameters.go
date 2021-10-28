// Code generated by go-swagger; DO NOT EDIT.

package notification

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

	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclientmodels"
)

// NewSendSpecificUserFreeformNotificationV1AdminParams creates a new SendSpecificUserFreeformNotificationV1AdminParams object
// with the default values initialized.
func NewSendSpecificUserFreeformNotificationV1AdminParams() *SendSpecificUserFreeformNotificationV1AdminParams {
	var ()
	return &SendSpecificUserFreeformNotificationV1AdminParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendSpecificUserFreeformNotificationV1AdminParamsWithTimeout creates a new SendSpecificUserFreeformNotificationV1AdminParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendSpecificUserFreeformNotificationV1AdminParamsWithTimeout(timeout time.Duration) *SendSpecificUserFreeformNotificationV1AdminParams {
	var ()
	return &SendSpecificUserFreeformNotificationV1AdminParams{

		timeout: timeout,
	}
}

// NewSendSpecificUserFreeformNotificationV1AdminParamsWithContext creates a new SendSpecificUserFreeformNotificationV1AdminParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendSpecificUserFreeformNotificationV1AdminParamsWithContext(ctx context.Context) *SendSpecificUserFreeformNotificationV1AdminParams {
	var ()
	return &SendSpecificUserFreeformNotificationV1AdminParams{

		Context: ctx,
	}
}

// NewSendSpecificUserFreeformNotificationV1AdminParamsWithHTTPClient creates a new SendSpecificUserFreeformNotificationV1AdminParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendSpecificUserFreeformNotificationV1AdminParamsWithHTTPClient(client *http.Client) *SendSpecificUserFreeformNotificationV1AdminParams {
	var ()
	return &SendSpecificUserFreeformNotificationV1AdminParams{
		HTTPClient: client,
	}
}

/*SendSpecificUserFreeformNotificationV1AdminParams contains all the parameters to send to the API endpoint
for the send specific user freeform notification v1 admin operation typically these are written to a http.Request
*/
type SendSpecificUserFreeformNotificationV1AdminParams struct {

	/*Body
	  notification content

	*/
	Body *lobbyclientmodels.ModelFreeFormNotificationRequestV1
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

// WithTimeout adds the timeout to the send specific user freeform notification v1 admin params
func (o *SendSpecificUserFreeformNotificationV1AdminParams) WithTimeout(timeout time.Duration) *SendSpecificUserFreeformNotificationV1AdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send specific user freeform notification v1 admin params
func (o *SendSpecificUserFreeformNotificationV1AdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send specific user freeform notification v1 admin params
func (o *SendSpecificUserFreeformNotificationV1AdminParams) WithContext(ctx context.Context) *SendSpecificUserFreeformNotificationV1AdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send specific user freeform notification v1 admin params
func (o *SendSpecificUserFreeformNotificationV1AdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send specific user freeform notification v1 admin params
func (o *SendSpecificUserFreeformNotificationV1AdminParams) WithHTTPClient(client *http.Client) *SendSpecificUserFreeformNotificationV1AdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send specific user freeform notification v1 admin params
func (o *SendSpecificUserFreeformNotificationV1AdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the send specific user freeform notification v1 admin params
func (o *SendSpecificUserFreeformNotificationV1AdminParams) WithBody(body *lobbyclientmodels.ModelFreeFormNotificationRequestV1) *SendSpecificUserFreeformNotificationV1AdminParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send specific user freeform notification v1 admin params
func (o *SendSpecificUserFreeformNotificationV1AdminParams) SetBody(body *lobbyclientmodels.ModelFreeFormNotificationRequestV1) {
	o.Body = body
}

// WithNamespace adds the namespace to the send specific user freeform notification v1 admin params
func (o *SendSpecificUserFreeformNotificationV1AdminParams) WithNamespace(namespace string) *SendSpecificUserFreeformNotificationV1AdminParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the send specific user freeform notification v1 admin params
func (o *SendSpecificUserFreeformNotificationV1AdminParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the send specific user freeform notification v1 admin params
func (o *SendSpecificUserFreeformNotificationV1AdminParams) WithUserID(userID string) *SendSpecificUserFreeformNotificationV1AdminParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the send specific user freeform notification v1 admin params
func (o *SendSpecificUserFreeformNotificationV1AdminParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *SendSpecificUserFreeformNotificationV1AdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
