// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package users_v4

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

// NewAdminDownloadMyBackupCodesV4Params creates a new AdminDownloadMyBackupCodesV4Params object
// with the default values initialized.
func NewAdminDownloadMyBackupCodesV4Params() *AdminDownloadMyBackupCodesV4Params {

	return &AdminDownloadMyBackupCodesV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminDownloadMyBackupCodesV4ParamsWithTimeout creates a new AdminDownloadMyBackupCodesV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminDownloadMyBackupCodesV4ParamsWithTimeout(timeout time.Duration) *AdminDownloadMyBackupCodesV4Params {

	return &AdminDownloadMyBackupCodesV4Params{

		timeout: timeout,
	}
}

// NewAdminDownloadMyBackupCodesV4ParamsWithContext creates a new AdminDownloadMyBackupCodesV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminDownloadMyBackupCodesV4ParamsWithContext(ctx context.Context) *AdminDownloadMyBackupCodesV4Params {

	return &AdminDownloadMyBackupCodesV4Params{

		Context: ctx,
	}
}

// NewAdminDownloadMyBackupCodesV4ParamsWithHTTPClient creates a new AdminDownloadMyBackupCodesV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminDownloadMyBackupCodesV4ParamsWithHTTPClient(client *http.Client) *AdminDownloadMyBackupCodesV4Params {

	return &AdminDownloadMyBackupCodesV4Params{
		HTTPClient: client,
	}
}

/*AdminDownloadMyBackupCodesV4Params contains all the parameters to send to the API endpoint
for the admin download my backup codes v4 operation typically these are written to a http.Request
*/
type AdminDownloadMyBackupCodesV4Params struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the admin download my backup codes v4 params
func (o *AdminDownloadMyBackupCodesV4Params) WithTimeout(timeout time.Duration) *AdminDownloadMyBackupCodesV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin download my backup codes v4 params
func (o *AdminDownloadMyBackupCodesV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin download my backup codes v4 params
func (o *AdminDownloadMyBackupCodesV4Params) WithContext(ctx context.Context) *AdminDownloadMyBackupCodesV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin download my backup codes v4 params
func (o *AdminDownloadMyBackupCodesV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the admin download my backup codes v4 params
func (o *AdminDownloadMyBackupCodesV4Params) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the admin download my backup codes v4 params
func (o *AdminDownloadMyBackupCodesV4Params) WithHTTPClient(client *http.Client) *AdminDownloadMyBackupCodesV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin download my backup codes v4 params
func (o *AdminDownloadMyBackupCodesV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AdminDownloadMyBackupCodesV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
