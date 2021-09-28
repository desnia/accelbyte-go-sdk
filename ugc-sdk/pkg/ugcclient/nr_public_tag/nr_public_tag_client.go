// Code generated by go-swagger; DO NOT EDIT.

package nr_public_tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new nr public tag API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for nr public tag API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetTag(params *GetTagParams, authInfo runtime.ClientAuthInfoWriter) (*GetTagOK, *GetTagUnauthorized, *GetTagNotFound, *GetTagInternalServerError, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetTag gets tags

  Requires valid user token
*/
func (a *Client) GetTag(params *GetTagParams, authInfo runtime.ClientAuthInfoWriter) (*GetTagOK, *GetTagUnauthorized, *GetTagNotFound, *GetTagInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTag",
		Method:             "GET",
		PathPattern:        "/ugc/v1/public/namespaces/{namespace}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTagReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetTagOK:
		return v, nil, nil, nil, nil
	case *GetTagUnauthorized:
		return nil, v, nil, nil, nil
	case *GetTagNotFound:
		return nil, nil, v, nil, nil
	case *GetTagInternalServerError:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}