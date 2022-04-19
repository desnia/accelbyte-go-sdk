// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package catalog_changes

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

// NewQueryChangesParams creates a new QueryChangesParams object
// with the default values initialized.
func NewQueryChangesParams() *QueryChangesParams {
	var (
		limitDefault  = int32(20)
		offsetDefault = int32(0)
		sortByDefault = string("updatedAt:desc")
		statusDefault = string("UNPUBLISHED")
	)
	return &QueryChangesParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
		SortBy: &sortByDefault,
		Status: &statusDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryChangesParamsWithTimeout creates a new QueryChangesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryChangesParamsWithTimeout(timeout time.Duration) *QueryChangesParams {
	var (
		limitDefault  = int32(20)
		offsetDefault = int32(0)
		sortByDefault = string("updatedAt:desc")
		statusDefault = string("UNPUBLISHED")
	)
	return &QueryChangesParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
		SortBy: &sortByDefault,
		Status: &statusDefault,

		timeout: timeout,
	}
}

// NewQueryChangesParamsWithContext creates a new QueryChangesParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryChangesParamsWithContext(ctx context.Context) *QueryChangesParams {
	var (
		limitDefault  = int32(20)
		offsetDefault = int32(0)
		sortByDefault = string("updatedAt:desc")
		statusDefault = string("UNPUBLISHED")
	)
	return &QueryChangesParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
		SortBy: &sortByDefault,
		Status: &statusDefault,

		Context: ctx,
	}
}

// NewQueryChangesParamsWithHTTPClient creates a new QueryChangesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryChangesParamsWithHTTPClient(client *http.Client) *QueryChangesParams {
	var (
		limitDefault  = int32(20)
		offsetDefault = int32(0)
		sortByDefault = string("updatedAt:desc")
		statusDefault = string("UNPUBLISHED")
	)
	return &QueryChangesParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		SortBy:     &sortByDefault,
		Status:     &statusDefault,
		HTTPClient: client,
	}
}

/*QueryChangesParams contains all the parameters to send to the API endpoint
for the query changes operation typically these are written to a http.Request
*/
type QueryChangesParams struct {

	/*Action*/
	Action *string
	/*ItemType*/
	ItemType *string
	/*Limit
	  limit

	*/
	Limit *int32
	/*Namespace*/
	Namespace string
	/*Offset
	  offset

	*/
	Offset *int32
	/*SortBy
	  default is updatedAt:desc, allow values: [createdAt, createdAt:asc, createdAt:desc, updatedAt, updatedAt:asc, updatedAt:desc],and support sort group, eg: sortBy=title:asc,createdAt:desc. Make sure to always use more than one sort if the first sort is not an unique valuefor example, if you wish to sort by title, make sure to include other sort such as sku or createdAt after the first sort, eg: title:asc,updatedAt:desc

	*/
	SortBy *string
	/*Status
	  default is UNPUBLISHED

	*/
	Status *string
	/*StoreID*/
	StoreID string
	/*Type*/
	Type *string
	/*UpdatedAtEnd
	  updated at end , using ISO 8601 format e.g. yyyy-MM-dd'T'HH:mm:ssZZ

	*/
	UpdatedAtEnd *string
	/*UpdatedAtStart
	  updated at start , using ISO 8601 format e.g. yyyy-MM-dd'T'HH:mm:ssZZ

	*/
	UpdatedAtStart *string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the query changes params
func (o *QueryChangesParams) WithTimeout(timeout time.Duration) *QueryChangesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query changes params
func (o *QueryChangesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query changes params
func (o *QueryChangesParams) WithContext(ctx context.Context) *QueryChangesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query changes params
func (o *QueryChangesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the query changes params
func (o *QueryChangesParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the query changes params
func (o *QueryChangesParams) WithHTTPClient(client *http.Client) *QueryChangesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query changes params
func (o *QueryChangesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the query changes params
func (o *QueryChangesParams) WithAction(action *string) *QueryChangesParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the query changes params
func (o *QueryChangesParams) SetAction(action *string) {
	o.Action = action
}

// WithItemType adds the itemType to the query changes params
func (o *QueryChangesParams) WithItemType(itemType *string) *QueryChangesParams {
	o.SetItemType(itemType)
	return o
}

// SetItemType adds the itemType to the query changes params
func (o *QueryChangesParams) SetItemType(itemType *string) {
	o.ItemType = itemType
}

// WithLimit adds the limit to the query changes params
func (o *QueryChangesParams) WithLimit(limit *int32) *QueryChangesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query changes params
func (o *QueryChangesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the query changes params
func (o *QueryChangesParams) WithNamespace(namespace string) *QueryChangesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the query changes params
func (o *QueryChangesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the query changes params
func (o *QueryChangesParams) WithOffset(offset *int32) *QueryChangesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query changes params
func (o *QueryChangesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSortBy adds the sortBy to the query changes params
func (o *QueryChangesParams) WithSortBy(sortBy *string) *QueryChangesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the query changes params
func (o *QueryChangesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithStatus adds the status to the query changes params
func (o *QueryChangesParams) WithStatus(status *string) *QueryChangesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the query changes params
func (o *QueryChangesParams) SetStatus(status *string) {
	o.Status = status
}

// WithStoreID adds the storeID to the query changes params
func (o *QueryChangesParams) WithStoreID(storeID string) *QueryChangesParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the query changes params
func (o *QueryChangesParams) SetStoreID(storeID string) {
	o.StoreID = storeID
}

// WithType adds the typeVar to the query changes params
func (o *QueryChangesParams) WithType(typeVar *string) *QueryChangesParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the query changes params
func (o *QueryChangesParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithUpdatedAtEnd adds the updatedAtEnd to the query changes params
func (o *QueryChangesParams) WithUpdatedAtEnd(updatedAtEnd *string) *QueryChangesParams {
	o.SetUpdatedAtEnd(updatedAtEnd)
	return o
}

// SetUpdatedAtEnd adds the updatedAtEnd to the query changes params
func (o *QueryChangesParams) SetUpdatedAtEnd(updatedAtEnd *string) {
	o.UpdatedAtEnd = updatedAtEnd
}

// WithUpdatedAtStart adds the updatedAtStart to the query changes params
func (o *QueryChangesParams) WithUpdatedAtStart(updatedAtStart *string) *QueryChangesParams {
	o.SetUpdatedAtStart(updatedAtStart)
	return o
}

// SetUpdatedAtStart adds the updatedAtStart to the query changes params
func (o *QueryChangesParams) SetUpdatedAtStart(updatedAtStart *string) {
	o.UpdatedAtStart = updatedAtStart
}

// WriteToRequest writes these params to a swagger request
func (o *QueryChangesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Action != nil {

		// query param action
		var qrAction string
		if o.Action != nil {
			qrAction = *o.Action
		}
		qAction := qrAction
		if qAction != "" {
			if err := r.SetQueryParam("action", qAction); err != nil {
				return err
			}
		}

	}

	if o.ItemType != nil {

		// query param itemType
		var qrItemType string
		if o.ItemType != nil {
			qrItemType = *o.ItemType
		}
		qItemType := qrItemType
		if qItemType != "" {
			if err := r.SetQueryParam("itemType", qItemType); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
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
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string
		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {
			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	// path param storeId
	if err := r.SetPathParam("storeId", o.StoreID); err != nil {
		return err
	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if o.UpdatedAtEnd != nil {

		// query param updatedAtEnd
		var qrUpdatedAtEnd string
		if o.UpdatedAtEnd != nil {
			qrUpdatedAtEnd = *o.UpdatedAtEnd
		}
		qUpdatedAtEnd := qrUpdatedAtEnd
		if qUpdatedAtEnd != "" {
			if err := r.SetQueryParam("updatedAtEnd", qUpdatedAtEnd); err != nil {
				return err
			}
		}

	}

	if o.UpdatedAtStart != nil {

		// query param updatedAtStart
		var qrUpdatedAtStart string
		if o.UpdatedAtStart != nil {
			qrUpdatedAtStart = *o.UpdatedAtStart
		}
		qUpdatedAtStart := qrUpdatedAtStart
		if qUpdatedAtStart != "" {
			if err := r.SetQueryParam("updatedAtStart", qUpdatedAtStart); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}