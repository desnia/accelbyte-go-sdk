// Code generated by go-swagger; DO NOT EDIT.

package event

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

// NewGetEventByUserIDHandlerParams creates a new GetEventByUserIDHandlerParams object
// with the default values initialized.
func NewGetEventByUserIDHandlerParams() *GetEventByUserIDHandlerParams {
	var ()
	return &GetEventByUserIDHandlerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventByUserIDHandlerParamsWithTimeout creates a new GetEventByUserIDHandlerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventByUserIDHandlerParamsWithTimeout(timeout time.Duration) *GetEventByUserIDHandlerParams {
	var ()
	return &GetEventByUserIDHandlerParams{

		timeout: timeout,
	}
}

// NewGetEventByUserIDHandlerParamsWithContext creates a new GetEventByUserIDHandlerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventByUserIDHandlerParamsWithContext(ctx context.Context) *GetEventByUserIDHandlerParams {
	var ()
	return &GetEventByUserIDHandlerParams{

		Context: ctx,
	}
}

// NewGetEventByUserIDHandlerParamsWithHTTPClient creates a new GetEventByUserIDHandlerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEventByUserIDHandlerParamsWithHTTPClient(client *http.Client) *GetEventByUserIDHandlerParams {
	var ()
	return &GetEventByUserIDHandlerParams{
		HTTPClient: client,
	}
}

/*GetEventByUserIDHandlerParams contains all the parameters to send to the API endpoint
for the get event by user ID handler operation typically these are written to a http.Request
*/
type GetEventByUserIDHandlerParams struct {

	/*EndDate
	  Ending date. e.g. 2015-03-20T12:27:06.351Z

	*/
	EndDate string
	/*Namespace
	  Namespace

	*/
	Namespace string
	/*Offset
	  Offset to query

	*/
	Offset *float64
	/*PageSize
	  Number of result in a page

	*/
	PageSize float64
	/*StartDate
	  Starting date. e.g. 2015-03-20T12:27:06.351Z

	*/
	StartDate string
	/*UserID
	  User ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) WithTimeout(timeout time.Duration) *GetEventByUserIDHandlerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) WithContext(ctx context.Context) *GetEventByUserIDHandlerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) WithHTTPClient(client *http.Client) *GetEventByUserIDHandlerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) WithEndDate(endDate string) *GetEventByUserIDHandlerParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) SetEndDate(endDate string) {
	o.EndDate = endDate
}

// WithNamespace adds the namespace to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) WithNamespace(namespace string) *GetEventByUserIDHandlerParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) WithOffset(offset *float64) *GetEventByUserIDHandlerParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) SetOffset(offset *float64) {
	o.Offset = offset
}

// WithPageSize adds the pageSize to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) WithPageSize(pageSize float64) *GetEventByUserIDHandlerParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) SetPageSize(pageSize float64) {
	o.PageSize = pageSize
}

// WithStartDate adds the startDate to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) WithStartDate(startDate string) *GetEventByUserIDHandlerParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) SetStartDate(startDate string) {
	o.StartDate = startDate
}

// WithUserID adds the userID to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) WithUserID(userID string) *GetEventByUserIDHandlerParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get event by user ID handler params
func (o *GetEventByUserIDHandlerParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventByUserIDHandlerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param endDate
	qrEndDate := o.EndDate
	qEndDate := qrEndDate
	if qEndDate != "" {
		if err := r.SetQueryParam("endDate", qEndDate); err != nil {
			return err
		}
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset float64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatFloat64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	// query param pageSize
	qrPageSize := o.PageSize
	qPageSize := swag.FormatFloat64(qrPageSize)
	if qPageSize != "" {
		if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
			return err
		}
	}

	// query param startDate
	qrStartDate := o.StartDate
	qStartDate := qrStartDate
	if qStartDate != "" {
		if err := r.SetQueryParam("startDate", qStartDate); err != nil {
			return err
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