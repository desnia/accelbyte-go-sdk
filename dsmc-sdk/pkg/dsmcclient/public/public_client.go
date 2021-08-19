// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new public API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for public API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetDefaultProvider(params *GetDefaultProviderParams, authInfo runtime.ClientAuthInfoWriter) (*GetDefaultProviderOK, error)

	ListProviders(params *ListProvidersParams, authInfo runtime.ClientAuthInfoWriter) (*ListProvidersOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetDefaultProvider gets default provider

  This endpoints returns the default provider.
*/
func (a *Client) GetDefaultProvider(params *GetDefaultProviderParams, authInfo runtime.ClientAuthInfoWriter) (*GetDefaultProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDefaultProviderParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDefaultProvider",
		Method:             "GET",
		PathPattern:        "/dsmcontroller/public/provider/default",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDefaultProviderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetDefaultProviderOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  ListProviders lists all supported providers

  This endpoints returns list of supported providers. Armada is the default provider.
*/
func (a *Client) ListProviders(params *ListProvidersParams, authInfo runtime.ClientAuthInfoWriter) (*ListProvidersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProvidersParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListProviders",
		Method:             "GET",
		PathPattern:        "/dsmcontroller/public/providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListProvidersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *ListProvidersOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}