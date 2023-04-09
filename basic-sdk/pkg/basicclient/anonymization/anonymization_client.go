// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package anonymization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new anonymization API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for anonymization API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AnonymizeUserProfile(params *AnonymizeUserProfileParams, authInfo runtime.ClientAuthInfoWriter) (*AnonymizeUserProfileNoContent, *AnonymizeUserProfileBadRequest, *AnonymizeUserProfileUnauthorized, *AnonymizeUserProfileForbidden, error)
	AnonymizeUserProfileShort(params *AnonymizeUserProfileParams, authInfo runtime.ClientAuthInfoWriter) (*AnonymizeUserProfileNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Deprecated: 2022-08-10 - Use AnonymizeUserProfileShort instead.

  AnonymizeUserProfile anonymizes user profile

  Anonymize user profile.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&lt;b&gt;&#34;ADMIN:NAMESPACE:{namespace}:USER:{userId}:ANONYMIZATION&#34;&lt;/b&gt;, action=8 &lt;b&gt;(DELETE)&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Action code&lt;/i&gt;: 11501&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) AnonymizeUserProfile(params *AnonymizeUserProfileParams, authInfo runtime.ClientAuthInfoWriter) (*AnonymizeUserProfileNoContent, *AnonymizeUserProfileBadRequest, *AnonymizeUserProfileUnauthorized, *AnonymizeUserProfileForbidden, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAnonymizeUserProfileParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "anonymizeUserProfile",
		Method:             "DELETE",
		PathPattern:        "/basic/v1/admin/namespaces/{namespace}/users/{userId}/anonymization/profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AnonymizeUserProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AnonymizeUserProfileNoContent:
		return v, nil, nil, nil, nil

	case *AnonymizeUserProfileBadRequest:
		return nil, v, nil, nil, nil

	case *AnonymizeUserProfileUnauthorized:
		return nil, nil, v, nil, nil

	case *AnonymizeUserProfileForbidden:
		return nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  AnonymizeUserProfileShort anonymizes user profile

  Anonymize user profile.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&lt;b&gt;&#34;ADMIN:NAMESPACE:{namespace}:USER:{userId}:ANONYMIZATION&#34;&lt;/b&gt;, action=8 &lt;b&gt;(DELETE)&lt;/b&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Action code&lt;/i&gt;: 11501&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) AnonymizeUserProfileShort(params *AnonymizeUserProfileParams, authInfo runtime.ClientAuthInfoWriter) (*AnonymizeUserProfileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAnonymizeUserProfileParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "anonymizeUserProfile",
		Method:             "DELETE",
		PathPattern:        "/basic/v1/admin/namespaces/{namespace}/users/{userId}/anonymization/profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AnonymizeUserProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *AnonymizeUserProfileNoContent:
		return v, nil
	case *AnonymizeUserProfileBadRequest:
		return nil, v
	case *AnonymizeUserProfileUnauthorized:
		return nil, v
	case *AnonymizeUserProfileForbidden:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
