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
)

// NewGetLocalizationTemplateParams creates a new GetLocalizationTemplateParams object
// with the default values initialized.
func NewGetLocalizationTemplateParams() *GetLocalizationTemplateParams {
	var ()
	return &GetLocalizationTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLocalizationTemplateParamsWithTimeout creates a new GetLocalizationTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLocalizationTemplateParamsWithTimeout(timeout time.Duration) *GetLocalizationTemplateParams {
	var ()
	return &GetLocalizationTemplateParams{

		timeout: timeout,
	}
}

// NewGetLocalizationTemplateParamsWithContext creates a new GetLocalizationTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLocalizationTemplateParamsWithContext(ctx context.Context) *GetLocalizationTemplateParams {
	var ()
	return &GetLocalizationTemplateParams{

		Context: ctx,
	}
}

// NewGetLocalizationTemplateParamsWithHTTPClient creates a new GetLocalizationTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLocalizationTemplateParamsWithHTTPClient(client *http.Client) *GetLocalizationTemplateParams {
	var ()
	return &GetLocalizationTemplateParams{
		HTTPClient: client,
	}
}

/*GetLocalizationTemplateParams contains all the parameters to send to the API endpoint
for the get localization template operation typically these are written to a http.Request
*/
type GetLocalizationTemplateParams struct {

	/*Namespace
	  namespace

	*/
	Namespace string
	/*TemplateLanguage
	  template language

	*/
	TemplateLanguage string
	/*TemplateSlug
	  template slug

	*/
	TemplateSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get localization template params
func (o *GetLocalizationTemplateParams) WithTimeout(timeout time.Duration) *GetLocalizationTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get localization template params
func (o *GetLocalizationTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get localization template params
func (o *GetLocalizationTemplateParams) WithContext(ctx context.Context) *GetLocalizationTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get localization template params
func (o *GetLocalizationTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get localization template params
func (o *GetLocalizationTemplateParams) WithHTTPClient(client *http.Client) *GetLocalizationTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get localization template params
func (o *GetLocalizationTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get localization template params
func (o *GetLocalizationTemplateParams) WithNamespace(namespace string) *GetLocalizationTemplateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get localization template params
func (o *GetLocalizationTemplateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithTemplateLanguage adds the templateLanguage to the get localization template params
func (o *GetLocalizationTemplateParams) WithTemplateLanguage(templateLanguage string) *GetLocalizationTemplateParams {
	o.SetTemplateLanguage(templateLanguage)
	return o
}

// SetTemplateLanguage adds the templateLanguage to the get localization template params
func (o *GetLocalizationTemplateParams) SetTemplateLanguage(templateLanguage string) {
	o.TemplateLanguage = templateLanguage
}

// WithTemplateSlug adds the templateSlug to the get localization template params
func (o *GetLocalizationTemplateParams) WithTemplateSlug(templateSlug string) *GetLocalizationTemplateParams {
	o.SetTemplateSlug(templateSlug)
	return o
}

// SetTemplateSlug adds the templateSlug to the get localization template params
func (o *GetLocalizationTemplateParams) SetTemplateSlug(templateSlug string) {
	o.TemplateSlug = templateSlug
}

// WriteToRequest writes these params to a swagger request
func (o *GetLocalizationTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param templateLanguage
	if err := r.SetPathParam("templateLanguage", o.TemplateLanguage); err != nil {
		return err
	}

	// path param templateSlug
	if err := r.SetPathParam("templateSlug", o.TemplateSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
