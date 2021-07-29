// Code generated by go-swagger; DO NOT EDIT.

package category

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

// NewPublicGetDescendantCategoriesParams creates a new PublicGetDescendantCategoriesParams object
// with the default values initialized.
func NewPublicGetDescendantCategoriesParams() *PublicGetDescendantCategoriesParams {
	var ()
	return &PublicGetDescendantCategoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicGetDescendantCategoriesParamsWithTimeout creates a new PublicGetDescendantCategoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicGetDescendantCategoriesParamsWithTimeout(timeout time.Duration) *PublicGetDescendantCategoriesParams {
	var ()
	return &PublicGetDescendantCategoriesParams{

		timeout: timeout,
	}
}

// NewPublicGetDescendantCategoriesParamsWithContext creates a new PublicGetDescendantCategoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicGetDescendantCategoriesParamsWithContext(ctx context.Context) *PublicGetDescendantCategoriesParams {
	var ()
	return &PublicGetDescendantCategoriesParams{

		Context: ctx,
	}
}

// NewPublicGetDescendantCategoriesParamsWithHTTPClient creates a new PublicGetDescendantCategoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicGetDescendantCategoriesParamsWithHTTPClient(client *http.Client) *PublicGetDescendantCategoriesParams {
	var ()
	return &PublicGetDescendantCategoriesParams{
		HTTPClient: client,
	}
}

/*PublicGetDescendantCategoriesParams contains all the parameters to send to the API endpoint
for the public get descendant categories operation typically these are written to a http.Request
*/
type PublicGetDescendantCategoriesParams struct {

	/*CategoryPath*/
	CategoryPath string
	/*Language*/
	Language *string
	/*Namespace
	  Namespace

	*/
	Namespace string
	/*StoreID
	  default is published store id

	*/
	StoreID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public get descendant categories params
func (o *PublicGetDescendantCategoriesParams) WithTimeout(timeout time.Duration) *PublicGetDescendantCategoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public get descendant categories params
func (o *PublicGetDescendantCategoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public get descendant categories params
func (o *PublicGetDescendantCategoriesParams) WithContext(ctx context.Context) *PublicGetDescendantCategoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public get descendant categories params
func (o *PublicGetDescendantCategoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public get descendant categories params
func (o *PublicGetDescendantCategoriesParams) WithHTTPClient(client *http.Client) *PublicGetDescendantCategoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public get descendant categories params
func (o *PublicGetDescendantCategoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategoryPath adds the categoryPath to the public get descendant categories params
func (o *PublicGetDescendantCategoriesParams) WithCategoryPath(categoryPath string) *PublicGetDescendantCategoriesParams {
	o.SetCategoryPath(categoryPath)
	return o
}

// SetCategoryPath adds the categoryPath to the public get descendant categories params
func (o *PublicGetDescendantCategoriesParams) SetCategoryPath(categoryPath string) {
	o.CategoryPath = categoryPath
}

// WithLanguage adds the language to the public get descendant categories params
func (o *PublicGetDescendantCategoriesParams) WithLanguage(language *string) *PublicGetDescendantCategoriesParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the public get descendant categories params
func (o *PublicGetDescendantCategoriesParams) SetLanguage(language *string) {
	o.Language = language
}

// WithNamespace adds the namespace to the public get descendant categories params
func (o *PublicGetDescendantCategoriesParams) WithNamespace(namespace string) *PublicGetDescendantCategoriesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public get descendant categories params
func (o *PublicGetDescendantCategoriesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithStoreID adds the storeID to the public get descendant categories params
func (o *PublicGetDescendantCategoriesParams) WithStoreID(storeID *string) *PublicGetDescendantCategoriesParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the public get descendant categories params
func (o *PublicGetDescendantCategoriesParams) SetStoreID(storeID *string) {
	o.StoreID = storeID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicGetDescendantCategoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param categoryPath
	if err := r.SetPathParam("categoryPath", o.CategoryPath); err != nil {
		return err
	}

	if o.Language != nil {

		// query param language
		var qrLanguage string
		if o.Language != nil {
			qrLanguage = *o.Language
		}
		qLanguage := qrLanguage
		if qLanguage != "" {
			if err := r.SetQueryParam("language", qLanguage); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.StoreID != nil {

		// query param storeId
		var qrStoreID string
		if o.StoreID != nil {
			qrStoreID = *o.StoreID
		}
		qStoreID := qrStoreID
		if qStoreID != "" {
			if err := r.SetQueryParam("storeId", qStoreID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}