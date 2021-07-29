// Code generated by go-swagger; DO NOT EDIT.

package currency

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new currency API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for currency API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCurrency(params *CreateCurrencyParams, authInfo runtime.ClientAuthInfoWriter) (*CreateCurrencyOK, *CreateCurrencyConflict, *CreateCurrencyUnprocessableEntity, error)

	DeleteCurrency(params *DeleteCurrencyParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCurrencyOK, *DeleteCurrencyNotFound, error)

	GetCurrencyConfig(params *GetCurrencyConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetCurrencyConfigOK, *GetCurrencyConfigNotFound, error)

	GetCurrencySummary(params *GetCurrencySummaryParams, authInfo runtime.ClientAuthInfoWriter) (*GetCurrencySummaryOK, *GetCurrencySummaryNotFound, error)

	ListCurrencies(params *ListCurrenciesParams, authInfo runtime.ClientAuthInfoWriter) (*ListCurrenciesOK, error)

	PublicListCurrencies(params *PublicListCurrenciesParams) (*PublicListCurrenciesOK, error)

	UpdateCurrency(params *UpdateCurrencyParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCurrencyOK, *UpdateCurrencyNotFound, *UpdateCurrencyUnprocessableEntity, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCurrency creates a currency

  Create a currency.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:CURRENCY", action=1 (CREATE)</li><li><i>Returns</i>: created currency</li></ul>
*/
func (a *Client) CreateCurrency(params *CreateCurrencyParams, authInfo runtime.ClientAuthInfoWriter) (*CreateCurrencyOK, *CreateCurrencyConflict, *CreateCurrencyUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCurrencyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCurrency",
		Method:             "POST",
		PathPattern:        "/admin/namespaces/{namespace}/currencies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCurrencyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreateCurrencyOK:
		return v, nil, nil, nil
	case *CreateCurrencyConflict:
		return nil, v, nil, nil
	case *CreateCurrencyUnprocessableEntity:
		return nil, nil, v, nil
	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteCurrency deletes a currency

  Delete a currency by currency code.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:CURRENCY", action=8 (DELETE)</li><li><i>Returns</i>: </li></ul>
*/
func (a *Client) DeleteCurrency(params *DeleteCurrencyParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCurrencyOK, *DeleteCurrencyNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCurrencyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCurrency",
		Method:             "DELETE",
		PathPattern:        "/admin/namespaces/{namespace}/currencies/{currencyCode}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCurrencyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteCurrencyOK:
		return v, nil, nil
	case *DeleteCurrencyNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetCurrencyConfig gets currency config

  <b>[SERVICE COMMUNICATION ONLY]</b> Get currency config by code.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:CURRENCY", action=2 (READ)</li><li><i>Returns</i>: simplified Currency</li></ul>
*/
func (a *Client) GetCurrencyConfig(params *GetCurrencyConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetCurrencyConfigOK, *GetCurrencyConfigNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrencyConfigParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCurrencyConfig",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/currencies/{currencyCode}/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCurrencyConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *GetCurrencyConfigOK:
		return v, nil, nil
	case *GetCurrencyConfigNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetCurrencySummary gets currency summary

  Get currency summary by code.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:CURRENCY", action=2 (READ)</li><li><i>Returns</i>: simplified Currency</li></ul>
*/
func (a *Client) GetCurrencySummary(params *GetCurrencySummaryParams, authInfo runtime.ClientAuthInfoWriter) (*GetCurrencySummaryOK, *GetCurrencySummaryNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrencySummaryParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCurrencySummary",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/currencies/{currencyCode}/summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCurrencySummaryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *GetCurrencySummaryOK:
		return v, nil, nil
	case *GetCurrencySummaryNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  ListCurrencies lists currencies

  List currencies of a namespace.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:CURRENCY", action=2 (READ)</li><li><i>Returns</i>: Currency List</li></ul>
*/
func (a *Client) ListCurrencies(params *ListCurrenciesParams, authInfo runtime.ClientAuthInfoWriter) (*ListCurrenciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCurrenciesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listCurrencies",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/currencies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListCurrenciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *ListCurrenciesOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicListCurrencies lists currencies

  List currencies of a namespace.<br>Other detail info: <ul><li><i>Returns</i>: Currency List</li></ul>
*/
func (a *Client) PublicListCurrencies(params *PublicListCurrenciesParams) (*PublicListCurrenciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicListCurrenciesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicListCurrencies",
		Method:             "GET",
		PathPattern:        "/public/namespaces/{namespace}/currencies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicListCurrenciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *PublicListCurrenciesOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateCurrency updates a currency

  Update a currency by currency code.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:CURRENCY", action=4 (UPDATE)</li><li><i>Returns</i>: updated currency</li></ul>
*/
func (a *Client) UpdateCurrency(params *UpdateCurrencyParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCurrencyOK, *UpdateCurrencyNotFound, *UpdateCurrencyUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCurrencyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCurrency",
		Method:             "PUT",
		PathPattern:        "/admin/namespaces/{namespace}/currencies/{currencyCode}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCurrencyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateCurrencyOK:
		return v, nil, nil, nil
	case *UpdateCurrencyNotFound:
		return nil, v, nil, nil
	case *UpdateCurrencyUnprocessableEntity:
		return nil, nil, v, nil
	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}