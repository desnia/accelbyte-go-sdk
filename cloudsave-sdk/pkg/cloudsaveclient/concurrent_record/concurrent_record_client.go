// Code generated by go-swagger; DO NOT EDIT.

package concurrent_record

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new concurrent record API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for concurrent record API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	PutGameRecordConcurrentHandlerV1(params *PutGameRecordConcurrentHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) (*PutGameRecordConcurrentHandlerV1NoContent, *PutGameRecordConcurrentHandlerV1BadRequest, *PutGameRecordConcurrentHandlerV1PreconditionFailed, *PutGameRecordConcurrentHandlerV1InternalServerError, error)

	PutPlayerPublicRecordConcurrentHandlerV1(params *PutPlayerPublicRecordConcurrentHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) (*PutPlayerPublicRecordConcurrentHandlerV1NoContent, *PutPlayerPublicRecordConcurrentHandlerV1BadRequest, *PutPlayerPublicRecordConcurrentHandlerV1PreconditionFailed, *PutPlayerPublicRecordConcurrentHandlerV1InternalServerError, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PutGameRecordConcurrentHandlerV1 creates or replace game record

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

<b>Optimistic Concurrency Control</b><br>
This endpoint implement optimistic concurrency control to avoid race condition.
If the record has been updated since the client fetch it, the server will return HTTP status code 412 (precondition failed)
and client need to redo the operation (fetch data and do update).
Otherwise, the server will process the request.

*/
func (a *Client) PutGameRecordConcurrentHandlerV1(params *PutGameRecordConcurrentHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) (*PutGameRecordConcurrentHandlerV1NoContent, *PutGameRecordConcurrentHandlerV1BadRequest, *PutGameRecordConcurrentHandlerV1PreconditionFailed, *PutGameRecordConcurrentHandlerV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutGameRecordConcurrentHandlerV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putGameRecordConcurrentHandlerV1",
		Method:             "PUT",
		PathPattern:        "/cloudsave/v1/namespaces/{namespace}/concurrent/records/{key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutGameRecordConcurrentHandlerV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PutGameRecordConcurrentHandlerV1NoContent:
		return v, nil, nil, nil, nil
	case *PutGameRecordConcurrentHandlerV1BadRequest:
		return nil, v, nil, nil, nil
	case *PutGameRecordConcurrentHandlerV1PreconditionFailed:
		return nil, nil, v, nil, nil
	case *PutGameRecordConcurrentHandlerV1InternalServerError:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PutPlayerPublicRecordConcurrentHandlerV1 creates or replace player record

  <table>
	<tr>
		<td>Required Permission</td>
		<td><code>NAMESPACE:{namespace}:USER:{userId}:PUBLIC:CLOUDSAVE:RECORD [UPDATE]</code></td>
	</tr>
	<tr>
		<td>Required Scope</td>
		<td><code>social</code></td>
	</tr>
</table>
<br/>

If the record is not exist, it will create. If the record already exist, it will replace the record
instead. And this operation can only be applied to record with <code>isPublic=true</code>.

Example

Replace record
<pre>
// existed record
{
	"foo": "bar"
}

// new record (request body)
{
	"foo_new": "bar_new"
}

// result
{
	"foo_new": "bar_new"
}
</pre>

<b>Optimistic Concurrency Control</b><br>
This endpoint implement optimistic concurrency control to avoid race condition.
If the record has been updated since the client fetch it, the server will return HTTP status code 412 (precondition failed)
and client need to redo the operation (fetch data and do update).
Otherwise, the server will process the request.

*/
func (a *Client) PutPlayerPublicRecordConcurrentHandlerV1(params *PutPlayerPublicRecordConcurrentHandlerV1Params, authInfo runtime.ClientAuthInfoWriter) (*PutPlayerPublicRecordConcurrentHandlerV1NoContent, *PutPlayerPublicRecordConcurrentHandlerV1BadRequest, *PutPlayerPublicRecordConcurrentHandlerV1PreconditionFailed, *PutPlayerPublicRecordConcurrentHandlerV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutPlayerPublicRecordConcurrentHandlerV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putPlayerPublicRecordConcurrentHandlerV1",
		Method:             "PUT",
		PathPattern:        "/cloudsave/v1/namespaces/{namespace}/users/{userID}/concurrent/records/{key}/public",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutPlayerPublicRecordConcurrentHandlerV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PutPlayerPublicRecordConcurrentHandlerV1NoContent:
		return v, nil, nil, nil, nil
	case *PutPlayerPublicRecordConcurrentHandlerV1BadRequest:
		return nil, v, nil, nil, nil
	case *PutPlayerPublicRecordConcurrentHandlerV1PreconditionFailed:
		return nil, nil, v, nil, nil
	case *PutPlayerPublicRecordConcurrentHandlerV1InternalServerError:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}