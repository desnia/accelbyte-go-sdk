// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package o_auth2_0

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
	"github.com/go-openapi/swag"
)

// NewTokenGrantV3Params creates a new TokenGrantV3Params object
// with the default values initialized.
func NewTokenGrantV3Params() *TokenGrantV3Params {
	var (
		grantTypeDefault = string("authorization_code")
	)
	return &TokenGrantV3Params{
		GrantType: grantTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTokenGrantV3ParamsWithTimeout creates a new TokenGrantV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewTokenGrantV3ParamsWithTimeout(timeout time.Duration) *TokenGrantV3Params {
	var (
		grantTypeDefault = string("authorization_code")
	)
	return &TokenGrantV3Params{
		GrantType: grantTypeDefault,

		timeout: timeout,
	}
}

// NewTokenGrantV3ParamsWithContext creates a new TokenGrantV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewTokenGrantV3ParamsWithContext(ctx context.Context) *TokenGrantV3Params {
	var (
		grantTypeDefault = string("authorization_code")
	)
	return &TokenGrantV3Params{
		GrantType: grantTypeDefault,

		Context: ctx,
	}
}

// NewTokenGrantV3ParamsWithHTTPClient creates a new TokenGrantV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTokenGrantV3ParamsWithHTTPClient(client *http.Client) *TokenGrantV3Params {
	var (
		grantTypeDefault = string("authorization_code")
	)
	return &TokenGrantV3Params{
		GrantType:  grantTypeDefault,
		HTTPClient: client,
	}
}

/*TokenGrantV3Params contains all the parameters to send to the API endpoint
for the token grant v3 operation typically these are written to a http.Request
*/
type TokenGrantV3Params struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*ClientID
	  client_id (used with grant type 'authorization_code')

	*/
	ClientID *string
	/*Code
	  The authorization code received from the authorization server (used with grant type 'authorization_code')

	*/
	Code *string
	/*CodeVerifier
	  Code verifier received from the authorization server

	*/
	CodeVerifier *string
	/*DeviceID
	  DeviceID (Used on grant type 'password' to track login history) ex. 90252d14544846d79f367148e3f9a3d9

	*/
	DeviceID *string
	/*ExtendExp
	  Extend expiration date of refresh token. Only available for grant type 'password'

	*/
	ExtendExp *bool
	/*GrantType
	  Grant Type

	*/
	GrantType string
	/*Password
	  Password (used with grant type 'password'

	*/
	Password *string
	/*RedirectURI
	  Redirect URI (used with grant type 'authorization_code')

	*/
	RedirectURI *string
	/*RefreshToken
	  Refresh Token (used with grant type 'refresh_token'). This field is optional if the request header provides the "refresh_token" cookie

	*/
	RefreshToken *string
	/*Username
	  User Name (used with grant type 'password'

	*/
	Username *string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the token grant v3 params
func (o *TokenGrantV3Params) WithTimeout(timeout time.Duration) *TokenGrantV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the token grant v3 params
func (o *TokenGrantV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the token grant v3 params
func (o *TokenGrantV3Params) WithContext(ctx context.Context) *TokenGrantV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the token grant v3 params
func (o *TokenGrantV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the token grant v3 params
func (o *TokenGrantV3Params) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the token grant v3 params
func (o *TokenGrantV3Params) WithHTTPClient(client *http.Client) *TokenGrantV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the token grant v3 params
func (o *TokenGrantV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the token grant v3 params
func (o *TokenGrantV3Params) WithClientID(clientID *string) *TokenGrantV3Params {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the token grant v3 params
func (o *TokenGrantV3Params) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithCode adds the code to the token grant v3 params
func (o *TokenGrantV3Params) WithCode(code *string) *TokenGrantV3Params {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the token grant v3 params
func (o *TokenGrantV3Params) SetCode(code *string) {
	o.Code = code
}

// WithCodeVerifier adds the codeVerifier to the token grant v3 params
func (o *TokenGrantV3Params) WithCodeVerifier(codeVerifier *string) *TokenGrantV3Params {
	o.SetCodeVerifier(codeVerifier)
	return o
}

// SetCodeVerifier adds the codeVerifier to the token grant v3 params
func (o *TokenGrantV3Params) SetCodeVerifier(codeVerifier *string) {
	o.CodeVerifier = codeVerifier
}

// WithDeviceID adds the deviceID to the token grant v3 params
func (o *TokenGrantV3Params) WithDeviceID(deviceID *string) *TokenGrantV3Params {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the token grant v3 params
func (o *TokenGrantV3Params) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithExtendExp adds the extendExp to the token grant v3 params
func (o *TokenGrantV3Params) WithExtendExp(extendExp *bool) *TokenGrantV3Params {
	o.SetExtendExp(extendExp)
	return o
}

// SetExtendExp adds the extendExp to the token grant v3 params
func (o *TokenGrantV3Params) SetExtendExp(extendExp *bool) {
	o.ExtendExp = extendExp
}

// WithGrantType adds the grantType to the token grant v3 params
func (o *TokenGrantV3Params) WithGrantType(grantType string) *TokenGrantV3Params {
	o.SetGrantType(grantType)
	return o
}

// SetGrantType adds the grantType to the token grant v3 params
func (o *TokenGrantV3Params) SetGrantType(grantType string) {
	o.GrantType = grantType
}

// WithPassword adds the password to the token grant v3 params
func (o *TokenGrantV3Params) WithPassword(password *string) *TokenGrantV3Params {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the token grant v3 params
func (o *TokenGrantV3Params) SetPassword(password *string) {
	o.Password = password
}

// WithRedirectURI adds the redirectURI to the token grant v3 params
func (o *TokenGrantV3Params) WithRedirectURI(redirectURI *string) *TokenGrantV3Params {
	o.SetRedirectURI(redirectURI)
	return o
}

// SetRedirectURI adds the redirectUri to the token grant v3 params
func (o *TokenGrantV3Params) SetRedirectURI(redirectURI *string) {
	o.RedirectURI = redirectURI
}

// WithRefreshToken adds the refreshToken to the token grant v3 params
func (o *TokenGrantV3Params) WithRefreshToken(refreshToken *string) *TokenGrantV3Params {
	o.SetRefreshToken(refreshToken)
	return o
}

// SetRefreshToken adds the refreshToken to the token grant v3 params
func (o *TokenGrantV3Params) SetRefreshToken(refreshToken *string) {
	o.RefreshToken = refreshToken
}

// WithUsername adds the username to the token grant v3 params
func (o *TokenGrantV3Params) WithUsername(username *string) *TokenGrantV3Params {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the token grant v3 params
func (o *TokenGrantV3Params) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *TokenGrantV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientID != nil {

		// form param client_id
		var frClientID string
		if o.ClientID != nil {
			frClientID = *o.ClientID
		}
		fClientID := frClientID
		if fClientID != "" {
			if err := r.SetFormParam("client_id", fClientID); err != nil {
				return err
			}
		}

	}

	if o.Code != nil {

		// form param code
		var frCode string
		if o.Code != nil {
			frCode = *o.Code
		}
		fCode := frCode
		if fCode != "" {
			if err := r.SetFormParam("code", fCode); err != nil {
				return err
			}
		}

	}

	if o.CodeVerifier != nil {

		// form param code_verifier
		var frCodeVerifier string
		if o.CodeVerifier != nil {
			frCodeVerifier = *o.CodeVerifier
		}
		fCodeVerifier := frCodeVerifier
		if fCodeVerifier != "" {
			if err := r.SetFormParam("code_verifier", fCodeVerifier); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// header param device_id
		if err := r.SetHeaderParam("device_id", *o.DeviceID); err != nil {
			return err
		}

	}

	if o.ExtendExp != nil {

		// form param extend_exp
		var frExtendExp bool
		if o.ExtendExp != nil {
			frExtendExp = *o.ExtendExp
		}
		fExtendExp := swag.FormatBool(frExtendExp)
		if fExtendExp != "" {
			if err := r.SetFormParam("extend_exp", fExtendExp); err != nil {
				return err
			}
		}

	}

	// form param grant_type
	frGrantType := o.GrantType
	fGrantType := frGrantType
	if fGrantType != "" {
		if err := r.SetFormParam("grant_type", fGrantType); err != nil {
			return err
		}
	}

	if o.Password != nil {

		// form param password
		var frPassword string
		if o.Password != nil {
			frPassword = *o.Password
		}
		fPassword := frPassword
		if fPassword != "" {
			if err := r.SetFormParam("password", fPassword); err != nil {
				return err
			}
		}

	}

	if o.RedirectURI != nil {

		// form param redirect_uri
		var frRedirectURI string
		if o.RedirectURI != nil {
			frRedirectURI = *o.RedirectURI
		}
		fRedirectURI := frRedirectURI
		if fRedirectURI != "" {
			if err := r.SetFormParam("redirect_uri", fRedirectURI); err != nil {
				return err
			}
		}

	}

	if o.RefreshToken != nil {

		// form param refresh_token
		var frRefreshToken string
		if o.RefreshToken != nil {
			frRefreshToken = *o.RefreshToken
		}
		fRefreshToken := frRefreshToken
		if fRefreshToken != "" {
			if err := r.SetFormParam("refresh_token", fRefreshToken); err != nil {
				return err
			}
		}

	}

	if o.Username != nil {

		// form param username
		var frUsername string
		if o.Username != nil {
			frUsername = *o.Username
		}
		fUsername := frUsername
		if fUsername != "" {
			if err := r.SetFormParam("username", fUsername); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
