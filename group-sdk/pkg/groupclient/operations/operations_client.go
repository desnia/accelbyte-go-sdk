// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated ; DO NOT EDIT.

package operations

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
	IndexHandler(params *IndexHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*IndexHandlerOK, error)
	IndexHandlerShort(params *IndexHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*IndexHandlerOK, error)
	BlockHandler(params *BlockHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*BlockHandlerOK, error)
	BlockHandlerShort(params *BlockHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*BlockHandlerOK, error)
	CmdlineHandler(params *CmdlineHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*CmdlineHandlerOK, error)
	CmdlineHandlerShort(params *CmdlineHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*CmdlineHandlerOK, error)
	GoroutineHandler(params *GoroutineHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*GoroutineHandlerOK, error)
	GoroutineHandlerShort(params *GoroutineHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*GoroutineHandlerOK, error)
	HeapHandler(params *HeapHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*HeapHandlerOK, error)
	HeapHandlerShort(params *HeapHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*HeapHandlerOK, error)
	Profile(params *ProfileParams, authInfo runtime.ClientAuthInfoWriter) (*ProfileOK, error)
	ProfileShort(params *ProfileParams, authInfo runtime.ClientAuthInfoWriter) (*ProfileOK, error)
	SymbolHandler(params *SymbolHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*SymbolHandlerOK, error)
	SymbolHandlerShort(params *SymbolHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*SymbolHandlerOK, error)
	ThreadcreateHandler(params *ThreadcreateHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*ThreadcreateHandlerOK, error)
	ThreadcreateHandlerShort(params *ThreadcreateHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*ThreadcreateHandlerOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Deprecated: 2022-08-10 - Use IndexHandlerShort instead.

IndexHandler
*/
func (a *Client) IndexHandler(params *IndexHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*IndexHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	if params.XFlightId != nil {
		params.SetFlightId(*params.XFlightId)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexHandler",
		Method:             "GET",
		PathPattern:        "/group/debug/pprof",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IndexHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *IndexHandlerOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
IndexHandlerShort
*/
func (a *Client) IndexHandlerShort(params *IndexHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*IndexHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indexHandler",
		Method:             "GET",
		PathPattern:        "/group/debug/pprof",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IndexHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *IndexHandlerOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use BlockHandlerShort instead.

BlockHandler
*/
func (a *Client) BlockHandler(params *BlockHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*BlockHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBlockHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	if params.XFlightId != nil {
		params.SetFlightId(*params.XFlightId)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "blockHandler",
		Method:             "GET",
		PathPattern:        "/group/debug/pprof/block",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BlockHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *BlockHandlerOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
BlockHandlerShort
*/
func (a *Client) BlockHandlerShort(params *BlockHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*BlockHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBlockHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "blockHandler",
		Method:             "GET",
		PathPattern:        "/group/debug/pprof/block",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BlockHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *BlockHandlerOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use CmdlineHandlerShort instead.

CmdlineHandler
*/
func (a *Client) CmdlineHandler(params *CmdlineHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*CmdlineHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCmdlineHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	if params.XFlightId != nil {
		params.SetFlightId(*params.XFlightId)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cmdlineHandler",
		Method:             "GET",
		PathPattern:        "/group/debug/pprof/cmdline",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CmdlineHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *CmdlineHandlerOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
CmdlineHandlerShort
*/
func (a *Client) CmdlineHandlerShort(params *CmdlineHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*CmdlineHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCmdlineHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cmdlineHandler",
		Method:             "GET",
		PathPattern:        "/group/debug/pprof/cmdline",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CmdlineHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *CmdlineHandlerOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use GoroutineHandlerShort instead.

GoroutineHandler
*/
func (a *Client) GoroutineHandler(params *GoroutineHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*GoroutineHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGoroutineHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	if params.XFlightId != nil {
		params.SetFlightId(*params.XFlightId)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "goroutineHandler",
		Method:             "GET",
		PathPattern:        "/group/debug/pprof/goroutine",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GoroutineHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GoroutineHandlerOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
GoroutineHandlerShort
*/
func (a *Client) GoroutineHandlerShort(params *GoroutineHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*GoroutineHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGoroutineHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "goroutineHandler",
		Method:             "GET",
		PathPattern:        "/group/debug/pprof/goroutine",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GoroutineHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GoroutineHandlerOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use HeapHandlerShort instead.

HeapHandler
*/
func (a *Client) HeapHandler(params *HeapHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*HeapHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeapHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	if params.XFlightId != nil {
		params.SetFlightId(*params.XFlightId)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "heapHandler",
		Method:             "GET",
		PathPattern:        "/group/debug/pprof/heap",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HeapHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *HeapHandlerOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
HeapHandlerShort
*/
func (a *Client) HeapHandlerShort(params *HeapHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*HeapHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeapHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "heapHandler",
		Method:             "GET",
		PathPattern:        "/group/debug/pprof/heap",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HeapHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *HeapHandlerOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use ProfileShort instead.

Profile
*/
func (a *Client) Profile(params *ProfileParams, authInfo runtime.ClientAuthInfoWriter) (*ProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProfileParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	if params.XFlightId != nil {
		params.SetFlightId(*params.XFlightId)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "profile",
		Method:             "GET",
		PathPattern:        "/group/debug/pprof/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *ProfileOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
ProfileShort
*/
func (a *Client) ProfileShort(params *ProfileParams, authInfo runtime.ClientAuthInfoWriter) (*ProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProfileParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "profile",
		Method:             "GET",
		PathPattern:        "/group/debug/pprof/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *ProfileOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use SymbolHandlerShort instead.

SymbolHandler
*/
func (a *Client) SymbolHandler(params *SymbolHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*SymbolHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSymbolHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	if params.XFlightId != nil {
		params.SetFlightId(*params.XFlightId)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "symbolHandler",
		Method:             "GET",
		PathPattern:        "/group/debug/pprof/symbol",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SymbolHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *SymbolHandlerOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
SymbolHandlerShort
*/
func (a *Client) SymbolHandlerShort(params *SymbolHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*SymbolHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSymbolHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "symbolHandler",
		Method:             "GET",
		PathPattern:        "/group/debug/pprof/symbol",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SymbolHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *SymbolHandlerOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use ThreadcreateHandlerShort instead.

ThreadcreateHandler
*/
func (a *Client) ThreadcreateHandler(params *ThreadcreateHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*ThreadcreateHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewThreadcreateHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	if params.XFlightId != nil {
		params.SetFlightId(*params.XFlightId)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "threadcreateHandler",
		Method:             "GET",
		PathPattern:        "/group/debug/pprof/threadcreate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ThreadcreateHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *ThreadcreateHandlerOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
ThreadcreateHandlerShort
*/
func (a *Client) ThreadcreateHandlerShort(params *ThreadcreateHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*ThreadcreateHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewThreadcreateHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "threadcreateHandler",
		Method:             "GET",
		PathPattern:        "/group/debug/pprof/threadcreate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ThreadcreateHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *ThreadcreateHandlerOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
