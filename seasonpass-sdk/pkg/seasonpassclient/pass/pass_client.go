// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package pass

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pass API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pass API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePass(params *CreatePassParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePassCreated, *CreatePassBadRequest, *CreatePassNotFound, *CreatePassConflict, *CreatePassUnprocessableEntity, error)
	CreatePassShort(params *CreatePassParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePassCreated, error)
	DeletePass(params *DeletePassParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePassNoContent, *DeletePassBadRequest, *DeletePassNotFound, *DeletePassConflict, error)
	DeletePassShort(params *DeletePassParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePassNoContent, error)
	GetPass(params *GetPassParams, authInfo runtime.ClientAuthInfoWriter) (*GetPassOK, *GetPassBadRequest, *GetPassNotFound, error)
	GetPassShort(params *GetPassParams, authInfo runtime.ClientAuthInfoWriter) (*GetPassOK, error)
	GrantUserPass(params *GrantUserPassParams, authInfo runtime.ClientAuthInfoWriter) (*GrantUserPassOK, *GrantUserPassBadRequest, error)
	GrantUserPassShort(params *GrantUserPassParams, authInfo runtime.ClientAuthInfoWriter) (*GrantUserPassOK, error)
	QueryPasses(params *QueryPassesParams, authInfo runtime.ClientAuthInfoWriter) (*QueryPassesOK, *QueryPassesBadRequest, *QueryPassesNotFound, error)
	QueryPassesShort(params *QueryPassesParams, authInfo runtime.ClientAuthInfoWriter) (*QueryPassesOK, error)
	UpdatePass(params *UpdatePassParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePassOK, *UpdatePassBadRequest, *UpdatePassNotFound, *UpdatePassConflict, *UpdatePassUnprocessableEntity, error)
	UpdatePassShort(params *UpdatePassParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePassOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Deprecated: 2022-08-10 - Use CreatePassShort instead.

  CreatePass creates a pass

  This API is used to create a pass for a draft season.&lt;p&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:SEASONPASS&#34;, action=1 (CREATE)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: created pass&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) CreatePass(params *CreatePassParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePassCreated, *CreatePassBadRequest, *CreatePassNotFound, *CreatePassConflict, *CreatePassUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePassParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPass",
		Method:             "POST",
		PathPattern:        "/seasonpass/admin/namespaces/{namespace}/seasons/{seasonId}/passes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePassReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreatePassCreated:
		return v, nil, nil, nil, nil, nil

	case *CreatePassBadRequest:
		return nil, v, nil, nil, nil, nil

	case *CreatePassNotFound:
		return nil, nil, v, nil, nil, nil

	case *CreatePassConflict:
		return nil, nil, nil, v, nil, nil

	case *CreatePassUnprocessableEntity:
		return nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  CreatePassShort creates a pass

  This API is used to create a pass for a draft season.&lt;p&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:SEASONPASS&#34;, action=1 (CREATE)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: created pass&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) CreatePassShort(params *CreatePassParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePassCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePassParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPass",
		Method:             "POST",
		PathPattern:        "/seasonpass/admin/namespaces/{namespace}/seasons/{seasonId}/passes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePassReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *CreatePassCreated:
		return v, nil
	case *CreatePassBadRequest:
		return nil, v
	case *CreatePassNotFound:
		return nil, v
	case *CreatePassConflict:
		return nil, v
	case *CreatePassUnprocessableEntity:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use DeletePassShort instead.

  DeletePass deletes a pass

  This API is used to delete a pass permanently, only draft season pass can be deleted. &lt;p&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:SEASONPASS&#34;, action=8 (DELETE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) DeletePass(params *DeletePassParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePassNoContent, *DeletePassBadRequest, *DeletePassNotFound, *DeletePassConflict, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePassParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePass",
		Method:             "DELETE",
		PathPattern:        "/seasonpass/admin/namespaces/{namespace}/seasons/{seasonId}/passes/{code}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePassReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DeletePassNoContent:
		return v, nil, nil, nil, nil

	case *DeletePassBadRequest:
		return nil, v, nil, nil, nil

	case *DeletePassNotFound:
		return nil, nil, v, nil, nil

	case *DeletePassConflict:
		return nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeletePassShort deletes a pass

  This API is used to delete a pass permanently, only draft season pass can be deleted. &lt;p&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:SEASONPASS&#34;, action=8 (DELETE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) DeletePassShort(params *DeletePassParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePassNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePassParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePass",
		Method:             "DELETE",
		PathPattern:        "/seasonpass/admin/namespaces/{namespace}/seasons/{seasonId}/passes/{code}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePassReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *DeletePassNoContent:
		return v, nil
	case *DeletePassBadRequest:
		return nil, v
	case *DeletePassNotFound:
		return nil, v
	case *DeletePassConflict:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use GetPassShort instead.

  GetPass gets a pass

  This API is used to get a pass for a season.&lt;p&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:SEASONPASS&#34;, action=2 (READ)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: pass data&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) GetPass(params *GetPassParams, authInfo runtime.ClientAuthInfoWriter) (*GetPassOK, *GetPassBadRequest, *GetPassNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPassParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPass",
		Method:             "GET",
		PathPattern:        "/seasonpass/admin/namespaces/{namespace}/seasons/{seasonId}/passes/{code}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPassReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetPassOK:
		return v, nil, nil, nil

	case *GetPassBadRequest:
		return nil, v, nil, nil

	case *GetPassNotFound:
		return nil, nil, v, nil

	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetPassShort gets a pass

  This API is used to get a pass for a season.&lt;p&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:SEASONPASS&#34;, action=2 (READ)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: pass data&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) GetPassShort(params *GetPassParams, authInfo runtime.ClientAuthInfoWriter) (*GetPassOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPassParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPass",
		Method:             "GET",
		PathPattern:        "/seasonpass/admin/namespaces/{namespace}/seasons/{seasonId}/passes/{code}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPassReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetPassOK:
		return v, nil
	case *GetPassBadRequest:
		return nil, v
	case *GetPassNotFound:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use GrantUserPassShort instead.

  GrantUserPass grants pass to user

  This API is used to grant pass to user, it will auto enroll if there&#39;s no user season but active published season exist, season only located in non-publisher namespace, otherwise ignore.&lt;p&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:USER:{userId}:SEASONPASS&#34;, action=4 (UPDATE)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: user season data&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) GrantUserPass(params *GrantUserPassParams, authInfo runtime.ClientAuthInfoWriter) (*GrantUserPassOK, *GrantUserPassBadRequest, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGrantUserPassParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "grantUserPass",
		Method:             "POST",
		PathPattern:        "/seasonpass/admin/namespaces/{namespace}/users/{userId}/seasons/current/passes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GrantUserPassReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *GrantUserPassOK:
		return v, nil, nil

	case *GrantUserPassBadRequest:
		return nil, v, nil

	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GrantUserPassShort grants pass to user

  This API is used to grant pass to user, it will auto enroll if there&#39;s no user season but active published season exist, season only located in non-publisher namespace, otherwise ignore.&lt;p&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:USER:{userId}:SEASONPASS&#34;, action=4 (UPDATE)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: user season data&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) GrantUserPassShort(params *GrantUserPassParams, authInfo runtime.ClientAuthInfoWriter) (*GrantUserPassOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGrantUserPassParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "grantUserPass",
		Method:             "POST",
		PathPattern:        "/seasonpass/admin/namespaces/{namespace}/users/{userId}/seasons/current/passes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GrantUserPassReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GrantUserPassOK:
		return v, nil
	case *GrantUserPassBadRequest:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use QueryPassesShort instead.

  QueryPasses queries all passes for a season

  This API is used to query all passes for a season.&lt;p&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:SEASONPASS&#34;, action=2 (READ)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: the list of passes&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) QueryPasses(params *QueryPassesParams, authInfo runtime.ClientAuthInfoWriter) (*QueryPassesOK, *QueryPassesBadRequest, *QueryPassesNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryPassesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryPasses",
		Method:             "GET",
		PathPattern:        "/seasonpass/admin/namespaces/{namespace}/seasons/{seasonId}/passes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryPassesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *QueryPassesOK:
		return v, nil, nil, nil

	case *QueryPassesBadRequest:
		return nil, v, nil, nil

	case *QueryPassesNotFound:
		return nil, nil, v, nil

	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  QueryPassesShort queries all passes for a season

  This API is used to query all passes for a season.&lt;p&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:SEASONPASS&#34;, action=2 (READ)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: the list of passes&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) QueryPassesShort(params *QueryPassesParams, authInfo runtime.ClientAuthInfoWriter) (*QueryPassesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryPassesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryPasses",
		Method:             "GET",
		PathPattern:        "/seasonpass/admin/namespaces/{namespace}/seasons/{seasonId}/passes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryPassesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *QueryPassesOK:
		return v, nil
	case *QueryPassesBadRequest:
		return nil, v
	case *QueryPassesNotFound:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use UpdatePassShort instead.

  UpdatePass updates a pass

  This API is used to update a pass. Only draft season pass can be updated.&lt;p&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:SEASONPASS&#34;, action=4 (UPDATE)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: updated pass&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) UpdatePass(params *UpdatePassParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePassOK, *UpdatePassBadRequest, *UpdatePassNotFound, *UpdatePassConflict, *UpdatePassUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePassParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updatePass",
		Method:             "PATCH",
		PathPattern:        "/seasonpass/admin/namespaces/{namespace}/seasons/{seasonId}/passes/{code}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePassReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdatePassOK:
		return v, nil, nil, nil, nil, nil

	case *UpdatePassBadRequest:
		return nil, v, nil, nil, nil, nil

	case *UpdatePassNotFound:
		return nil, nil, v, nil, nil, nil

	case *UpdatePassConflict:
		return nil, nil, nil, v, nil, nil

	case *UpdatePassUnprocessableEntity:
		return nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdatePassShort updates a pass

  This API is used to update a pass. Only draft season pass can be updated.&lt;p&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:SEASONPASS&#34;, action=4 (UPDATE)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: updated pass&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) UpdatePassShort(params *UpdatePassParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePassOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePassParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updatePass",
		Method:             "PATCH",
		PathPattern:        "/seasonpass/admin/namespaces/{namespace}/seasons/{seasonId}/passes/{code}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePassReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *UpdatePassOK:
		return v, nil
	case *UpdatePassBadRequest:
		return nil, v
	case *UpdatePassNotFound:
		return nil, v
	case *UpdatePassConflict:
		return nil, v
	case *UpdatePassUnprocessableEntity:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
