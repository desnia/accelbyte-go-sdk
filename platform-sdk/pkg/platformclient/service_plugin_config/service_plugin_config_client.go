// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package service_plugin_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service plugin config API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service plugin config API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteServicePluginConfig(params *DeleteServicePluginConfigParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServicePluginConfigNoContent, error)
	DeleteServicePluginConfigShort(params *DeleteServicePluginConfigParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServicePluginConfigNoContent, error)
	GetServicePluginConfig(params *GetServicePluginConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetServicePluginConfigOK, error)
	GetServicePluginConfigShort(params *GetServicePluginConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetServicePluginConfigOK, error)
	UpdateServicePluginConfig(params *UpdateServicePluginConfigParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServicePluginConfigOK, *UpdateServicePluginConfigUnprocessableEntity, error)
	UpdateServicePluginConfigShort(params *UpdateServicePluginConfigParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServicePluginConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Deprecated: 2022-08-10 - Use DeleteServicePluginConfigShort instead.

  DeleteServicePluginConfig deletes service plugin config

  Delete service plugin config.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=ADMIN:NAMESPACE:{namespace}:CONFIG:SERVICEPLUGIN, action=8 (DELETE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) DeleteServicePluginConfig(params *DeleteServicePluginConfigParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServicePluginConfigNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServicePluginConfigParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteServicePluginConfig",
		Method:             "DELETE",
		PathPattern:        "/platform/admin/namespaces/{namespace}/configs/servicePlugin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteServicePluginConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *DeleteServicePluginConfigNoContent:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteServicePluginConfigShort deletes service plugin config

  Delete service plugin config.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=ADMIN:NAMESPACE:{namespace}:CONFIG:SERVICEPLUGIN, action=8 (DELETE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) DeleteServicePluginConfigShort(params *DeleteServicePluginConfigParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServicePluginConfigNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServicePluginConfigParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteServicePluginConfig",
		Method:             "DELETE",
		PathPattern:        "/platform/admin/namespaces/{namespace}/configs/servicePlugin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteServicePluginConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *DeleteServicePluginConfigNoContent:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use GetServicePluginConfigShort instead.

  GetServicePluginConfig gets service plugin config

  Get service plugin config.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&lt;b&gt;ADMIN:NAMESPACE:{namespace}:CONFIG:SERVICEPLUGIN&lt;/b&gt;, action=2 &lt;b&gt;(READ)&lt;/b&gt;&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) GetServicePluginConfig(params *GetServicePluginConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetServicePluginConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServicePluginConfigParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getServicePluginConfig",
		Method:             "GET",
		PathPattern:        "/platform/admin/namespaces/{namespace}/configs/servicePlugin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServicePluginConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetServicePluginConfigOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetServicePluginConfigShort gets service plugin config

  Get service plugin config.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&lt;b&gt;ADMIN:NAMESPACE:{namespace}:CONFIG:SERVICEPLUGIN&lt;/b&gt;, action=2 &lt;b&gt;(READ)&lt;/b&gt;&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) GetServicePluginConfigShort(params *GetServicePluginConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetServicePluginConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServicePluginConfigParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getServicePluginConfig",
		Method:             "GET",
		PathPattern:        "/platform/admin/namespaces/{namespace}/configs/servicePlugin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServicePluginConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetServicePluginConfigOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use UpdateServicePluginConfigShort instead.

  UpdateServicePluginConfig updates service plugin config service

  Update catalog config. Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=ADMIN:NAMESPACE:{namespace}:CONFIG:SERVICEPLUGIN, action=4 (UPDATE)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: updated service plugin config&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) UpdateServicePluginConfig(params *UpdateServicePluginConfigParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServicePluginConfigOK, *UpdateServicePluginConfigUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateServicePluginConfigParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateServicePluginConfig",
		Method:             "PUT",
		PathPattern:        "/platform/admin/namespaces/{namespace}/configs/servicePlugin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateServicePluginConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateServicePluginConfigOK:
		return v, nil, nil

	case *UpdateServicePluginConfigUnprocessableEntity:
		return nil, v, nil

	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateServicePluginConfigShort updates service plugin config service

  Update catalog config. Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=ADMIN:NAMESPACE:{namespace}:CONFIG:SERVICEPLUGIN, action=4 (UPDATE)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: updated service plugin config&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) UpdateServicePluginConfigShort(params *UpdateServicePluginConfigParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServicePluginConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateServicePluginConfigParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateServicePluginConfig",
		Method:             "PUT",
		PathPattern:        "/platform/admin/namespaces/{namespace}/configs/servicePlugin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateServicePluginConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *UpdateServicePluginConfigOK:
		return v, nil
	case *UpdateServicePluginConfigUnprocessableEntity:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
