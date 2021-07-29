// Code generated by go-swagger; DO NOT EDIT.

package public_game_record

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new public game record API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for public game record API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteGameRecordHandlerV1(params *DeleteGameRecordHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteGameRecordHandlerV1OK, *DeleteGameRecordHandlerV1InternalServerError, error)

	GetGameRecordHandlerV1(params *GetGameRecordHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetGameRecordHandlerV1OK, *GetGameRecordHandlerV1InternalServerError, error)

	PostGameRecordHandlerV1(params *PostGameRecordHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) (*PostGameRecordHandlerV1OK, *PostGameRecordHandlerV1InternalServerError, error)

	PutGameRecordHandlerV1(params *PutGameRecordHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) (*PutGameRecordHandlerV1OK, *PutGameRecordHandlerV1InternalServerError, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteGameRecordHandlerV1 deletes game record

  <table>
	<tr>
		<td>Required Permission</td>
		<td><code>NAMESPACE:{namespace}:CLOUDSAVE:RECORD [DELETE]</code></td>
	</tr>
	<tr>
		<td>Required Scope</td>
		<td><code>social</code></td>
	</tr>
</table>
<br/>

Delete records by its key

*/
func (a *Client) DeleteGameRecordHandlerV1(params *DeleteGameRecordHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteGameRecordHandlerV1OK, *DeleteGameRecordHandlerV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGameRecordHandlerV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteGameRecordHandlerV1",
		Method:             "DELETE",
		PathPattern:        "/cloudsave/v1/namespaces/{namespace}/records/{key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteGameRecordHandlerV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteGameRecordHandlerV1OK:
		return v, nil, nil
	case *DeleteGameRecordHandlerV1InternalServerError:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetGameRecordHandlerV1 gets game record

  <table>
	<tr>
		<td>Required Permission</td>
		<td><code>NAMESPACE:{namespace}:CLOUDSAVE:RECORD [READ]</code></td>
	</tr>
	<tr>
		<td>Required Scope</td>
		<td><code>social</code></td>
	</tr>
</table>
<br/>

Get game record by its key.

*/
func (a *Client) GetGameRecordHandlerV1(params *GetGameRecordHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetGameRecordHandlerV1OK, *GetGameRecordHandlerV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGameRecordHandlerV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGameRecordHandlerV1",
		Method:             "GET",
		PathPattern:        "/cloudsave/v1/namespaces/{namespace}/records/{key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGameRecordHandlerV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *GetGameRecordHandlerV1OK:
		return v, nil, nil
	case *GetGameRecordHandlerV1InternalServerError:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PostGameRecordHandlerV1 creates or append game record

  <table>
	<tr>
		<td>Required Permission</td>
		<td><code>NAMESPACE:{namespace}:CLOUDSAVE:RECORD [CREATE]</code></td>
	</tr>
	<tr>
		<td>Required Scope</td>
		<td><code>social</code></td>
	</tr>
</table>
<br/>

If there's already record, the record will be merged with conditions:
- If field name is already exist, the value will be replaced
- If field name is not exists it will append the field and its value

Example:

Replace value in a specific JSON key
<pre>
// existed record
{
	"foo": "bar"
}

// new update (request body)
{
	"foo": "barUpdated"
}

// result
{
	"foo": "barUpdated"
}
</pre>

Append new json item
<pre>
// existed record
{
	"foo": "bar"
}

// new update (request body)
{
	"foo_new": "bar_new"
}

// result
{
	"foo": "bar",
	"foo_new": "bar_new"
}
</pre>

*/
func (a *Client) PostGameRecordHandlerV1(params *PostGameRecordHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) (*PostGameRecordHandlerV1OK, *PostGameRecordHandlerV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostGameRecordHandlerV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postGameRecordHandlerV1",
		Method:             "POST",
		PathPattern:        "/cloudsave/v1/namespaces/{namespace}/records/{key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostGameRecordHandlerV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *PostGameRecordHandlerV1OK:
		return v, nil, nil
	case *PostGameRecordHandlerV1InternalServerError:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PutGameRecordHandlerV1 creates or replace game record

  <table>
	<tr>
		<td>Required Permission</td>
		<td><code>NAMESPACE:{namespace}:CLOUDSAVE:RECORD [UPDATE]</code></td>
	</tr>
	<tr>
		<td>Required Scope</td>
		<td><code>social</code></td>
	</tr>
</table>
<br/>

If record already exists, it will be replaced with the one from request body (all fields will be
deleted). If record is not exists, it will create a new one with value from request body.

Example:

Replace all records
<pre>
	// existed record
	{
		"foo": "bar"
	}

	// new update (request body)
	{
		"foo_new": "bar_new"
	}

	// result
	{
		"foo_new": "bar_new"
	}
</pre>

*/
func (a *Client) PutGameRecordHandlerV1(params *PutGameRecordHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) (*PutGameRecordHandlerV1OK, *PutGameRecordHandlerV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutGameRecordHandlerV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putGameRecordHandlerV1",
		Method:             "PUT",
		PathPattern:        "/cloudsave/v1/namespaces/{namespace}/records/{key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutGameRecordHandlerV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *PutGameRecordHandlerV1OK:
		return v, nil, nil
	case *PutGameRecordHandlerV1InternalServerError:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}