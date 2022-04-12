// Code generated by go-swagger; DO NOT EDIT.

package user_action

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user action API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user action API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	BanUsers(params *BanUsersParams, authInfo runtime.ClientAuthInfoWriter) (*BanUsersNoContent, *BanUsersBadRequest, *BanUsersNotFound, *BanUsersUnprocessableEntity, *BanUsersInternalServerError, error)
	BanUsersShort(params *BanUsersParams, authInfo runtime.ClientAuthInfoWriter) (*BanUsersNoContent, error)
	GetActions(params *GetActionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetActionsOK, *GetActionsBadRequest, *GetActionsNotFound, *GetActionsInternalServerError, error)
	GetActionsShort(params *GetActionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetActionsOK, error)
	GetBannedUsers(params *GetBannedUsersParams, authInfo runtime.ClientAuthInfoWriter) (*GetBannedUsersOK, *GetBannedUsersBadRequest, *GetBannedUsersNotFound, *GetBannedUsersUnprocessableEntity, *GetBannedUsersInternalServerError, error)
	GetBannedUsersShort(params *GetBannedUsersParams, authInfo runtime.ClientAuthInfoWriter) (*GetBannedUsersOK, error)
	GetUserStatus(params *GetUserStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserStatusOK, *GetUserStatusBadRequest, *GetUserStatusNotFound, *GetUserStatusUnprocessableEntity, *GetUserStatusInternalServerError, error)
	GetUserStatusShort(params *GetUserStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserStatusOK, error)
	PublicReportUser(params *PublicReportUserParams, authInfo runtime.ClientAuthInfoWriter) (*PublicReportUserNoContent, *PublicReportUserBadRequest, *PublicReportUserUnprocessableEntity, error)
	PublicReportUserShort(params *PublicReportUserParams, authInfo runtime.ClientAuthInfoWriter) (*PublicReportUserNoContent, error)
	ReportUser(params *ReportUserParams, authInfo runtime.ClientAuthInfoWriter) (*ReportUserNoContent, *ReportUserUnprocessableEntity, error)
	ReportUserShort(params *ReportUserParams, authInfo runtime.ClientAuthInfoWriter) (*ReportUserNoContent, error)
	UnBanUsers(params *UnBanUsersParams, authInfo runtime.ClientAuthInfoWriter) (*UnBanUsersNoContent, *UnBanUsersBadRequest, *UnBanUsersNotFound, *UnBanUsersUnprocessableEntity, *UnBanUsersInternalServerError, error)
	UnBanUsersShort(params *UnBanUsersParams, authInfo runtime.ClientAuthInfoWriter) (*UnBanUsersNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BanUsers bans user temporarily or permanently

  Ban user.&lt;br&gt; actionId: 1 means permanent ban, actionId: 10 means Temporary ban.Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&lt;b&gt;&#34;ADMIN:NAMESPACE:{namespace}:ACTION&#34;&lt;/b&gt;, action=4 &lt;b&gt;(UPDATE)&lt;/b&gt;&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) BanUsers(params *BanUsersParams, authInfo runtime.ClientAuthInfoWriter) (*BanUsersNoContent, *BanUsersBadRequest, *BanUsersNotFound, *BanUsersUnprocessableEntity, *BanUsersInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBanUsersParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "banUsers",
		Method:             "POST",
		PathPattern:        "/basic/v1/admin/namespaces/{namespace}/actions/ban",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BanUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *BanUsersNoContent:
		return v, nil, nil, nil, nil, nil

	case *BanUsersBadRequest:
		return nil, v, nil, nil, nil, nil

	case *BanUsersNotFound:
		return nil, nil, v, nil, nil, nil

	case *BanUsersUnprocessableEntity:
		return nil, nil, nil, v, nil, nil

	case *BanUsersInternalServerError:
		return nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) BanUsersShort(params *BanUsersParams, authInfo runtime.ClientAuthInfoWriter) (*BanUsersNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBanUsersParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "banUsers",
		Method:             "POST",
		PathPattern:        "/basic/v1/admin/namespaces/{namespace}/actions/ban",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BanUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *BanUsersNoContent:
		return v, nil
	case *BanUsersBadRequest:
		return nil, v
	case *BanUsersNotFound:
		return nil, v
	case *BanUsersUnprocessableEntity:
		return nil, v
	case *BanUsersInternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetActions gets configured actions

  Get configured actions.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&lt;b&gt;&#34;ADMIN:NAMESPACE:{namespace}:ACTION&#34;&lt;/b&gt;, action=2 &lt;b&gt;(READ)&lt;/b&gt;&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) GetActions(params *GetActionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetActionsOK, *GetActionsBadRequest, *GetActionsNotFound, *GetActionsInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActionsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getActions",
		Method:             "GET",
		PathPattern:        "/basic/v1/admin/namespaces/{namespace}/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetActionsOK:
		return v, nil, nil, nil, nil

	case *GetActionsBadRequest:
		return nil, v, nil, nil, nil

	case *GetActionsNotFound:
		return nil, nil, v, nil, nil

	case *GetActionsInternalServerError:
		return nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) GetActionsShort(params *GetActionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetActionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActionsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getActions",
		Method:             "GET",
		PathPattern:        "/basic/v1/admin/namespaces/{namespace}/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetActionsOK:
		return v, nil
	case *GetActionsBadRequest:
		return nil, v
	case *GetActionsNotFound:
		return nil, v
	case *GetActionsInternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetBannedUsers gets banned user

  Get banned status.&lt;br&gt;Unbanned users will not return, for example: request has 8 userIds, only 5 of then were banned, then the api will these 5 user status.Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&lt;b&gt;&#34;ADMIN:NAMESPACE:{namespace}:ACTION&#34;&lt;/b&gt;, action=2 &lt;b&gt;(READ)&lt;/b&gt;&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) GetBannedUsers(params *GetBannedUsersParams, authInfo runtime.ClientAuthInfoWriter) (*GetBannedUsersOK, *GetBannedUsersBadRequest, *GetBannedUsersNotFound, *GetBannedUsersUnprocessableEntity, *GetBannedUsersInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBannedUsersParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBannedUsers",
		Method:             "GET",
		PathPattern:        "/basic/v1/admin/namespaces/{namespace}/actions/banned",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBannedUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetBannedUsersOK:
		return v, nil, nil, nil, nil, nil

	case *GetBannedUsersBadRequest:
		return nil, v, nil, nil, nil, nil

	case *GetBannedUsersNotFound:
		return nil, nil, v, nil, nil, nil

	case *GetBannedUsersUnprocessableEntity:
		return nil, nil, nil, v, nil, nil

	case *GetBannedUsersInternalServerError:
		return nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) GetBannedUsersShort(params *GetBannedUsersParams, authInfo runtime.ClientAuthInfoWriter) (*GetBannedUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBannedUsersParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBannedUsers",
		Method:             "GET",
		PathPattern:        "/basic/v1/admin/namespaces/{namespace}/actions/banned",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBannedUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetBannedUsersOK:
		return v, nil
	case *GetBannedUsersBadRequest:
		return nil, v
	case *GetBannedUsersNotFound:
		return nil, v
	case *GetBannedUsersUnprocessableEntity:
		return nil, v
	case *GetBannedUsersInternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetUserStatus gets user status

  Get user status.&lt;br&gt;If actionId does not exist, then the user is not banned.If actionId and expires exist, then the user is temporarily banned, if expires does not exist, then the user is permanently banned.Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&lt;b&gt;&#34;ADMIN:NAMESPACE:{namespace}:ACTION&#34;&lt;/b&gt;, action=2 &lt;b&gt;(READ)&lt;/b&gt;&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) GetUserStatus(params *GetUserStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserStatusOK, *GetUserStatusBadRequest, *GetUserStatusNotFound, *GetUserStatusUnprocessableEntity, *GetUserStatusInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserStatusParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserStatus",
		Method:             "GET",
		PathPattern:        "/basic/v1/admin/namespaces/{namespace}/actions/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetUserStatusOK:
		return v, nil, nil, nil, nil, nil

	case *GetUserStatusBadRequest:
		return nil, v, nil, nil, nil, nil

	case *GetUserStatusNotFound:
		return nil, nil, v, nil, nil, nil

	case *GetUserStatusUnprocessableEntity:
		return nil, nil, nil, v, nil, nil

	case *GetUserStatusInternalServerError:
		return nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) GetUserStatusShort(params *GetUserStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserStatusParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserStatus",
		Method:             "GET",
		PathPattern:        "/basic/v1/admin/namespaces/{namespace}/actions/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetUserStatusOK:
		return v, nil
	case *GetUserStatusBadRequest:
		return nil, v
	case *GetUserStatusNotFound:
		return nil, v
	case *GetUserStatusUnprocessableEntity:
		return nil, v
	case *GetUserStatusInternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicReportUser reports a game user

  This API is used to report a game user.&lt;p&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;NAMESPACE:{namespace}:USER:{userId}:ACTION&#34;, action=1 (CREATE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) PublicReportUser(params *PublicReportUserParams, authInfo runtime.ClientAuthInfoWriter) (*PublicReportUserNoContent, *PublicReportUserBadRequest, *PublicReportUserUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicReportUserParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicReportUser",
		Method:             "POST",
		PathPattern:        "/basic/v1/public/namespaces/{namespace}/users/{userId}/actions/report",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicReportUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PublicReportUserNoContent:
		return v, nil, nil, nil

	case *PublicReportUserBadRequest:
		return nil, v, nil, nil

	case *PublicReportUserUnprocessableEntity:
		return nil, nil, v, nil

	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) PublicReportUserShort(params *PublicReportUserParams, authInfo runtime.ClientAuthInfoWriter) (*PublicReportUserNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicReportUserParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicReportUser",
		Method:             "POST",
		PathPattern:        "/basic/v1/public/namespaces/{namespace}/users/{userId}/actions/report",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicReportUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *PublicReportUserNoContent:
		return v, nil
	case *PublicReportUserBadRequest:
		return nil, v
	case *PublicReportUserUnprocessableEntity:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  ReportUser reports a game player for game service

  This API is for game service to report a game player.&lt;p&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:ACTION&#34;, action=1 (CREATE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) ReportUser(params *ReportUserParams, authInfo runtime.ClientAuthInfoWriter) (*ReportUserNoContent, *ReportUserUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReportUserParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reportUser",
		Method:             "POST",
		PathPattern:        "/basic/v1/admin/namespaces/{namespace}/actions/report",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReportUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *ReportUserNoContent:
		return v, nil, nil

	case *ReportUserUnprocessableEntity:
		return nil, v, nil

	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) ReportUserShort(params *ReportUserParams, authInfo runtime.ClientAuthInfoWriter) (*ReportUserNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReportUserParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reportUser",
		Method:             "POST",
		PathPattern:        "/basic/v1/admin/namespaces/{namespace}/actions/report",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReportUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *ReportUserNoContent:
		return v, nil
	case *ReportUserUnprocessableEntity:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UnBanUsers unbans user

  Unban user.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&lt;b&gt;&#34;ADMIN:NAMESPACE:{namespace}:ACTION&#34;&lt;/b&gt;, action=4 &lt;b&gt;(UPDATE)&lt;/b&gt;&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) UnBanUsers(params *UnBanUsersParams, authInfo runtime.ClientAuthInfoWriter) (*UnBanUsersNoContent, *UnBanUsersBadRequest, *UnBanUsersNotFound, *UnBanUsersUnprocessableEntity, *UnBanUsersInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnBanUsersParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "unBanUsers",
		Method:             "POST",
		PathPattern:        "/basic/v1/admin/namespaces/{namespace}/actions/unban",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnBanUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UnBanUsersNoContent:
		return v, nil, nil, nil, nil, nil

	case *UnBanUsersBadRequest:
		return nil, v, nil, nil, nil, nil

	case *UnBanUsersNotFound:
		return nil, nil, v, nil, nil, nil

	case *UnBanUsersUnprocessableEntity:
		return nil, nil, nil, v, nil, nil

	case *UnBanUsersInternalServerError:
		return nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) UnBanUsersShort(params *UnBanUsersParams, authInfo runtime.ClientAuthInfoWriter) (*UnBanUsersNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnBanUsersParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "unBanUsers",
		Method:             "POST",
		PathPattern:        "/basic/v1/admin/namespaces/{namespace}/actions/unban",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnBanUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *UnBanUsersNoContent:
		return v, nil
	case *UnBanUsersBadRequest:
		return nil, v
	case *UnBanUsersNotFound:
		return nil, v
	case *UnBanUsersUnprocessableEntity:
		return nil, v
	case *UnBanUsersInternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
