// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_script

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new fulfillment script API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fulfillment script API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateFulfillmentScript(params *CreateFulfillmentScriptParams, authInfo runtime.ClientAuthInfoWriter) (*CreateFulfillmentScriptCreated, *CreateFulfillmentScriptConflict, error)
	CreateFulfillmentScriptShort(params *CreateFulfillmentScriptParams, authInfo runtime.ClientAuthInfoWriter) (*CreateFulfillmentScriptCreated, error)
	DeleteFulfillmentScript(params *DeleteFulfillmentScriptParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteFulfillmentScriptNoContent, error)
	DeleteFulfillmentScriptShort(params *DeleteFulfillmentScriptParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteFulfillmentScriptNoContent, error)
	GetFulfillmentScript(params *GetFulfillmentScriptParams, authInfo runtime.ClientAuthInfoWriter) (*GetFulfillmentScriptOK, *GetFulfillmentScriptNotFound, error)
	GetFulfillmentScriptShort(params *GetFulfillmentScriptParams, authInfo runtime.ClientAuthInfoWriter) (*GetFulfillmentScriptOK, error)
	ListFulfillmentScripts(params *ListFulfillmentScriptsParams, authInfo runtime.ClientAuthInfoWriter) (*ListFulfillmentScriptsOK, error)
	ListFulfillmentScriptsShort(params *ListFulfillmentScriptsParams, authInfo runtime.ClientAuthInfoWriter) (*ListFulfillmentScriptsOK, error)
	TestFulfillmentScriptEval(params *TestFulfillmentScriptEvalParams, authInfo runtime.ClientAuthInfoWriter) (*TestFulfillmentScriptEvalOK, error)
	TestFulfillmentScriptEvalShort(params *TestFulfillmentScriptEvalParams, authInfo runtime.ClientAuthInfoWriter) (*TestFulfillmentScriptEvalOK, error)
	UpdateFulfillmentScript(params *UpdateFulfillmentScriptParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateFulfillmentScriptOK, *UpdateFulfillmentScriptBadRequest, error)
	UpdateFulfillmentScriptShort(params *UpdateFulfillmentScriptParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateFulfillmentScriptOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Deprecated: 2022-08-10 - Use CreateFulfillmentScriptShort instead.

  CreateFulfillmentScript creates fulfillment script

  Create fulfillment script.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:FULFILLMENT&#34;, action=1 (CREATE)&lt;/li&gt;&lt;/ul&gt;Fulfillment scripts are used for adding custom fulfillment logic based on &lt;b&gt;ITEM_TYPE&lt;/b&gt;: [MEDIA,INGAMEITEM] for now, and the custom scripts only cover grantDays.&lt;br&gt;Example for grantDays: &lt;br&gt;&lt;code&gt;order &amp;&amp; ((order.currency &amp;&amp; order.currency.currencyCode) == &#39;LP&#39; || order.isFree) ? 30 : -1&lt;/code&gt;&lt;br&gt;
*/
func (a *Client) CreateFulfillmentScript(params *CreateFulfillmentScriptParams, authInfo runtime.ClientAuthInfoWriter) (*CreateFulfillmentScriptCreated, *CreateFulfillmentScriptConflict, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFulfillmentScriptParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createFulfillmentScript",
		Method:             "POST",
		PathPattern:        "/platform/admin/fulfillment/scripts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateFulfillmentScriptReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *CreateFulfillmentScriptCreated:
		return v, nil, nil

	case *CreateFulfillmentScriptConflict:
		return nil, v, nil

	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  CreateFulfillmentScriptShort creates fulfillment script

  Create fulfillment script.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:FULFILLMENT&#34;, action=1 (CREATE)&lt;/li&gt;&lt;/ul&gt;Fulfillment scripts are used for adding custom fulfillment logic based on &lt;b&gt;ITEM_TYPE&lt;/b&gt;: [MEDIA,INGAMEITEM] for now, and the custom scripts only cover grantDays.&lt;br&gt;Example for grantDays: &lt;br&gt;&lt;code&gt;order &amp;&amp; ((order.currency &amp;&amp; order.currency.currencyCode) == &#39;LP&#39; || order.isFree) ? 30 : -1&lt;/code&gt;&lt;br&gt;
*/
func (a *Client) CreateFulfillmentScriptShort(params *CreateFulfillmentScriptParams, authInfo runtime.ClientAuthInfoWriter) (*CreateFulfillmentScriptCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFulfillmentScriptParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createFulfillmentScript",
		Method:             "POST",
		PathPattern:        "/platform/admin/fulfillment/scripts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateFulfillmentScriptReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *CreateFulfillmentScriptCreated:
		return v, nil
	case *CreateFulfillmentScriptConflict:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use DeleteFulfillmentScriptShort instead.

  DeleteFulfillmentScript deletes fulfillment script

  Delete fulfillment script.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:FULFILLMENT&#34;, action=8 (DELETE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) DeleteFulfillmentScript(params *DeleteFulfillmentScriptParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteFulfillmentScriptNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFulfillmentScriptParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteFulfillmentScript",
		Method:             "DELETE",
		PathPattern:        "/platform/admin/fulfillment/scripts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFulfillmentScriptReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *DeleteFulfillmentScriptNoContent:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteFulfillmentScriptShort deletes fulfillment script

  Delete fulfillment script.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:FULFILLMENT&#34;, action=8 (DELETE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) DeleteFulfillmentScriptShort(params *DeleteFulfillmentScriptParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteFulfillmentScriptNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFulfillmentScriptParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteFulfillmentScript",
		Method:             "DELETE",
		PathPattern:        "/platform/admin/fulfillment/scripts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFulfillmentScriptReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *DeleteFulfillmentScriptNoContent:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use GetFulfillmentScriptShort instead.

  GetFulfillmentScript gets fulfillment script by id

  Get fulfillment script by id.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:FULFILLMENT&#34;, action=2 (READ)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: get fulfillment script&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) GetFulfillmentScript(params *GetFulfillmentScriptParams, authInfo runtime.ClientAuthInfoWriter) (*GetFulfillmentScriptOK, *GetFulfillmentScriptNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFulfillmentScriptParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFulfillmentScript",
		Method:             "GET",
		PathPattern:        "/platform/admin/fulfillment/scripts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFulfillmentScriptReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *GetFulfillmentScriptOK:
		return v, nil, nil

	case *GetFulfillmentScriptNotFound:
		return nil, v, nil

	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetFulfillmentScriptShort gets fulfillment script by id

  Get fulfillment script by id.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:FULFILLMENT&#34;, action=2 (READ)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: get fulfillment script&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) GetFulfillmentScriptShort(params *GetFulfillmentScriptParams, authInfo runtime.ClientAuthInfoWriter) (*GetFulfillmentScriptOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFulfillmentScriptParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFulfillmentScript",
		Method:             "GET",
		PathPattern:        "/platform/admin/fulfillment/scripts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFulfillmentScriptReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetFulfillmentScriptOK:
		return v, nil
	case *GetFulfillmentScriptNotFound:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use ListFulfillmentScriptsShort instead.

  ListFulfillmentScripts lists all fulfillment scripts

  List all fulfillment scripts.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:FULFILLMENT&#34;, action=2 (READ)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) ListFulfillmentScripts(params *ListFulfillmentScriptsParams, authInfo runtime.ClientAuthInfoWriter) (*ListFulfillmentScriptsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListFulfillmentScriptsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listFulfillmentScripts",
		Method:             "GET",
		PathPattern:        "/platform/admin/fulfillment/scripts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListFulfillmentScriptsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *ListFulfillmentScriptsOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  ListFulfillmentScriptsShort lists all fulfillment scripts

  List all fulfillment scripts.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:FULFILLMENT&#34;, action=2 (READ)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) ListFulfillmentScriptsShort(params *ListFulfillmentScriptsParams, authInfo runtime.ClientAuthInfoWriter) (*ListFulfillmentScriptsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListFulfillmentScriptsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listFulfillmentScripts",
		Method:             "GET",
		PathPattern:        "/platform/admin/fulfillment/scripts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListFulfillmentScriptsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *ListFulfillmentScriptsOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use TestFulfillmentScriptEvalShort instead.

  TestFulfillmentScriptEval tests eval fulfillment script

  &lt;b&gt;[TEST FACILITY ONLY]&lt;/b&gt;Test eval fulfillment script.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:FULFILLMENT&#34;, action=2 (READ)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) TestFulfillmentScriptEval(params *TestFulfillmentScriptEvalParams, authInfo runtime.ClientAuthInfoWriter) (*TestFulfillmentScriptEvalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestFulfillmentScriptEvalParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "testFulfillmentScriptEval",
		Method:             "POST",
		PathPattern:        "/platform/admin/fulfillment/scripts/tests/eval",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TestFulfillmentScriptEvalReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *TestFulfillmentScriptEvalOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  TestFulfillmentScriptEvalShort tests eval fulfillment script

  &lt;b&gt;[TEST FACILITY ONLY]&lt;/b&gt;Test eval fulfillment script.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:FULFILLMENT&#34;, action=2 (READ)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) TestFulfillmentScriptEvalShort(params *TestFulfillmentScriptEvalParams, authInfo runtime.ClientAuthInfoWriter) (*TestFulfillmentScriptEvalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestFulfillmentScriptEvalParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "testFulfillmentScriptEval",
		Method:             "POST",
		PathPattern:        "/platform/admin/fulfillment/scripts/tests/eval",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TestFulfillmentScriptEvalReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *TestFulfillmentScriptEvalOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use UpdateFulfillmentScriptShort instead.

  UpdateFulfillmentScript updates fulfillment script

  Update fulfillment script.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:FULFILLMENT&#34;, action=4 (UPDATE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) UpdateFulfillmentScript(params *UpdateFulfillmentScriptParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateFulfillmentScriptOK, *UpdateFulfillmentScriptBadRequest, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFulfillmentScriptParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateFulfillmentScript",
		Method:             "PATCH",
		PathPattern:        "/platform/admin/fulfillment/scripts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateFulfillmentScriptReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateFulfillmentScriptOK:
		return v, nil, nil

	case *UpdateFulfillmentScriptBadRequest:
		return nil, v, nil

	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateFulfillmentScriptShort updates fulfillment script

  Update fulfillment script.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:FULFILLMENT&#34;, action=4 (UPDATE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) UpdateFulfillmentScriptShort(params *UpdateFulfillmentScriptParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateFulfillmentScriptOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFulfillmentScriptParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateFulfillmentScript",
		Method:             "PATCH",
		PathPattern:        "/platform/admin/fulfillment/scripts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateFulfillmentScriptReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *UpdateFulfillmentScriptOK:
		return v, nil
	case *UpdateFulfillmentScriptBadRequest:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
