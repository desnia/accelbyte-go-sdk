// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package file_upload

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGeneratedUserUploadContentURLParams creates a new GeneratedUserUploadContentURLParams object
// with the default values initialized.
func NewGeneratedUserUploadContentURLParams() *GeneratedUserUploadContentURLParams {
	var ()
	return &GeneratedUserUploadContentURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGeneratedUserUploadContentURLParamsWithTimeout creates a new GeneratedUserUploadContentURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGeneratedUserUploadContentURLParamsWithTimeout(timeout time.Duration) *GeneratedUserUploadContentURLParams {
	var ()
	return &GeneratedUserUploadContentURLParams{

		timeout: timeout,
	}
}

// NewGeneratedUserUploadContentURLParamsWithContext creates a new GeneratedUserUploadContentURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewGeneratedUserUploadContentURLParamsWithContext(ctx context.Context) *GeneratedUserUploadContentURLParams {
	var ()
	return &GeneratedUserUploadContentURLParams{

		Context: ctx,
	}
}

// NewGeneratedUserUploadContentURLParamsWithHTTPClient creates a new GeneratedUserUploadContentURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGeneratedUserUploadContentURLParamsWithHTTPClient(client *http.Client) *GeneratedUserUploadContentURLParams {
	var ()
	return &GeneratedUserUploadContentURLParams{
		HTTPClient: client,
	}
}

/*GeneratedUserUploadContentURLParams contains all the parameters to send to the API endpoint
for the generated user upload content Url operation typically these are written to a http.Request
*/
type GeneratedUserUploadContentURLParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*FileType
	  one of the these types: jpeg, jpg, png, bmp, gif, mp3, bin, webp

	*/
	FileType string
	/*Namespace
	  namespace, only accept alphabet and numeric

	*/
	Namespace string
	/*UserID
	  user's id, should follow UUID version 4 without hyphen

	*/
	UserID string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the generated user upload content Url params
func (o *GeneratedUserUploadContentURLParams) WithTimeout(timeout time.Duration) *GeneratedUserUploadContentURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generated user upload content Url params
func (o *GeneratedUserUploadContentURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generated user upload content Url params
func (o *GeneratedUserUploadContentURLParams) WithContext(ctx context.Context) *GeneratedUserUploadContentURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generated user upload content Url params
func (o *GeneratedUserUploadContentURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the generated user upload content Url params
func (o *GeneratedUserUploadContentURLParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the generated user upload content Url params
func (o *GeneratedUserUploadContentURLParams) WithHTTPClient(client *http.Client) *GeneratedUserUploadContentURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generated user upload content Url params
func (o *GeneratedUserUploadContentURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFileType adds the fileType to the generated user upload content Url params
func (o *GeneratedUserUploadContentURLParams) WithFileType(fileType string) *GeneratedUserUploadContentURLParams {
	o.SetFileType(fileType)
	return o
}

// SetFileType adds the fileType to the generated user upload content Url params
func (o *GeneratedUserUploadContentURLParams) SetFileType(fileType string) {
	o.FileType = fileType
}

// WithNamespace adds the namespace to the generated user upload content Url params
func (o *GeneratedUserUploadContentURLParams) WithNamespace(namespace string) *GeneratedUserUploadContentURLParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the generated user upload content Url params
func (o *GeneratedUserUploadContentURLParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the generated user upload content Url params
func (o *GeneratedUserUploadContentURLParams) WithUserID(userID string) *GeneratedUserUploadContentURLParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the generated user upload content Url params
func (o *GeneratedUserUploadContentURLParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GeneratedUserUploadContentURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param fileType
	qrFileType := o.FileType
	qFileType := qrFileType
	if qFileType != "" {
		if err := r.SetQueryParam("fileType", qFileType); err != nil {
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
