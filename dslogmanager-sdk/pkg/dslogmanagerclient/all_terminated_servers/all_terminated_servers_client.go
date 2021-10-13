// Code generated by go-swagger; DO NOT EDIT.

package all_terminated_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new all terminated servers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for all terminated servers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListAllTerminatedServers(params *ListAllTerminatedServersParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllTerminatedServersOK, *ListAllTerminatedServersBadRequest, *ListAllTerminatedServersUnauthorized, *ListAllTerminatedServersInternalServerError, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListAllTerminatedServers retrieves all terminated servers

  ```
Required permission: ADMIN:NAMESPACE:{namespace}:DSLM:SERVER [READ]

This endpoint used to retrieve terminated servers in all namespace
```
*/
func (a *Client) ListAllTerminatedServers(params *ListAllTerminatedServersParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllTerminatedServersOK, *ListAllTerminatedServersBadRequest, *ListAllTerminatedServersUnauthorized, *ListAllTerminatedServersInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllTerminatedServersParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAllTerminatedServers",
		Method:             "GET",
		PathPattern:        "/dslogmanager/servers/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllTerminatedServersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *ListAllTerminatedServersOK:
		return v, nil, nil, nil, nil
	case *ListAllTerminatedServersBadRequest:
		return nil, v, nil, nil, nil
	case *ListAllTerminatedServersUnauthorized:
		return nil, nil, v, nil, nil
	case *ListAllTerminatedServersInternalServerError:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}