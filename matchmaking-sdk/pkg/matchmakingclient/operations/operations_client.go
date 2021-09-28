// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	Healthz(params *HealthzParams, authInfo runtime.ClientAuthInfoWriter) (*HealthzOK, error)

	MatchmakingHealthz(params *MatchmakingHealthzParams, authInfo runtime.ClientAuthInfoWriter) (*MatchmakingHealthzOK, error)

	PublicGetMessages(params *PublicGetMessagesParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetMessagesOK, *PublicGetMessagesInternalServerError, error)

	VersionCheckHandler(params *VersionCheckHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*VersionCheckHandlerOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Healthz healthz API
*/
func (a *Client) Healthz(params *HealthzParams, authInfo runtime.ClientAuthInfoWriter) (*HealthzOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHealthzParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "healthz",
		Method:             "GET",
		PathPattern:        "/healthz",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HealthzReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *HealthzOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  MatchmakingHealthz matchmaking healthz API
*/
func (a *Client) MatchmakingHealthz(params *MatchmakingHealthzParams, authInfo runtime.ClientAuthInfoWriter) (*MatchmakingHealthzOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMatchmakingHealthzParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "matchmakingHealthz",
		Method:             "GET",
		PathPattern:        "/matchmaking/healthz",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MatchmakingHealthzReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *MatchmakingHealthzOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicGetMessages gets service messages

  get the list of messages.
*/
func (a *Client) PublicGetMessages(params *PublicGetMessagesParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetMessagesOK, *PublicGetMessagesInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicGetMessagesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicGetMessages",
		Method:             "GET",
		PathPattern:        "/matchmaking/v1/messages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublicGetMessagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *PublicGetMessagesOK:
		return v, nil, nil
	case *PublicGetMessagesInternalServerError:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  VersionCheckHandler version check handler API
*/
func (a *Client) VersionCheckHandler(params *VersionCheckHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*VersionCheckHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVersionCheckHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "versionCheckHandler",
		Method:             "GET",
		PathPattern:        "/matchmaking/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VersionCheckHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *VersionCheckHandlerOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}