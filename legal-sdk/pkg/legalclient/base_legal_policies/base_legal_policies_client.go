// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package base_legal_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new base legal policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for base legal policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePolicy(params *CreatePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePolicyCreated, *CreatePolicyBadRequest, *CreatePolicyNotFound, *CreatePolicyUnprocessableEntity, error)
	CreatePolicyShort(params *CreatePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePolicyCreated, error)
	PartialUpdatePolicy(params *PartialUpdatePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*PartialUpdatePolicyOK, *PartialUpdatePolicyBadRequest, *PartialUpdatePolicyNotFound, error)
	PartialUpdatePolicyShort(params *PartialUpdatePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*PartialUpdatePolicyOK, error)
	RetrieveAllLegalPolicies(params *RetrieveAllLegalPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAllLegalPoliciesOK, error)
	RetrieveAllLegalPoliciesShort(params *RetrieveAllLegalPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAllLegalPoliciesOK, error)
	RetrieveAllPolicyTypes(params *RetrieveAllPolicyTypesParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAllPolicyTypesOK, error)
	RetrieveAllPolicyTypesShort(params *RetrieveAllPolicyTypesParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAllPolicyTypesOK, error)
	RetrievePolicyCountry(params *RetrievePolicyCountryParams, authInfo runtime.ClientAuthInfoWriter) (*RetrievePolicyCountryOK, *RetrievePolicyCountryNotFound, error)
	RetrievePolicyCountryShort(params *RetrievePolicyCountryParams, authInfo runtime.ClientAuthInfoWriter) (*RetrievePolicyCountryOK, error)
	RetrieveSinglePolicy(params *RetrieveSinglePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveSinglePolicyOK, *RetrieveSinglePolicyNotFound, error)
	RetrieveSinglePolicyShort(params *RetrieveSinglePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveSinglePolicyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Deprecated: Use CreatePolicyShort instead.

  CreatePolicy creates a base legal policy

  Create a legal policy.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=1 (CREATE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) CreatePolicy(params *CreatePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePolicyCreated, *CreatePolicyBadRequest, *CreatePolicyNotFound, *CreatePolicyUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePolicyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPolicy",
		Method:             "POST",
		PathPattern:        "/agreement/admin/base-policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreatePolicyCreated:
		return v, nil, nil, nil, nil

	case *CreatePolicyBadRequest:
		return nil, v, nil, nil, nil

	case *CreatePolicyNotFound:
		return nil, nil, v, nil, nil

	case *CreatePolicyUnprocessableEntity:
		return nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  CreatePolicyShort creates a base legal policy

  Create a legal policy.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=1 (CREATE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) CreatePolicyShort(params *CreatePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePolicyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePolicyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPolicy",
		Method:             "POST",
		PathPattern:        "/agreement/admin/base-policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *CreatePolicyCreated:
		return v, nil
	case *CreatePolicyBadRequest:
		return nil, v
	case *CreatePolicyNotFound:
		return nil, v
	case *CreatePolicyUnprocessableEntity:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: Use PartialUpdatePolicyShort instead.

  PartialUpdatePolicy updates base legal policy

  Update an existing base policy.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=4 (UPDATE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) PartialUpdatePolicy(params *PartialUpdatePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*PartialUpdatePolicyOK, *PartialUpdatePolicyBadRequest, *PartialUpdatePolicyNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartialUpdatePolicyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "partialUpdatePolicy",
		Method:             "PATCH",
		PathPattern:        "/agreement/admin/base-policies/{basePolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PartialUpdatePolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PartialUpdatePolicyOK:
		return v, nil, nil, nil

	case *PartialUpdatePolicyBadRequest:
		return nil, v, nil, nil

	case *PartialUpdatePolicyNotFound:
		return nil, nil, v, nil

	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PartialUpdatePolicyShort updates base legal policy

  Update an existing base policy.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=4 (UPDATE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) PartialUpdatePolicyShort(params *PartialUpdatePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*PartialUpdatePolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartialUpdatePolicyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "partialUpdatePolicy",
		Method:             "PATCH",
		PathPattern:        "/agreement/admin/base-policies/{basePolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PartialUpdatePolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *PartialUpdatePolicyOK:
		return v, nil
	case *PartialUpdatePolicyBadRequest:
		return nil, v
	case *PartialUpdatePolicyNotFound:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: Use RetrieveAllLegalPoliciesShort instead.

  RetrieveAllLegalPolicies retrieves all base legal policy

  Retrieve all base policies.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=2 (READ)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrieveAllLegalPolicies(params *RetrieveAllLegalPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAllLegalPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveAllLegalPoliciesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveAllLegalPolicies",
		Method:             "GET",
		PathPattern:        "/agreement/admin/base-policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveAllLegalPoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveAllLegalPoliciesOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrieveAllLegalPoliciesShort retrieves all base legal policy

  Retrieve all base policies.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=2 (READ)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrieveAllLegalPoliciesShort(params *RetrieveAllLegalPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAllLegalPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveAllLegalPoliciesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveAllLegalPolicies",
		Method:             "GET",
		PathPattern:        "/agreement/admin/base-policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveAllLegalPoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveAllLegalPoliciesOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: Use RetrieveAllPolicyTypesShort instead.

  RetrieveAllPolicyTypes retrieves all policy type

  Retrieve all supported policy types.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=2 (READ)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrieveAllPolicyTypes(params *RetrieveAllPolicyTypesParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAllPolicyTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveAllPolicyTypesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveAllPolicyTypes",
		Method:             "GET",
		PathPattern:        "/agreement/admin/policy-types",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveAllPolicyTypesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveAllPolicyTypesOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrieveAllPolicyTypesShort retrieves all policy type

  Retrieve all supported policy types.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=2 (READ)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrieveAllPolicyTypesShort(params *RetrieveAllPolicyTypesParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAllPolicyTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveAllPolicyTypesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveAllPolicyTypes",
		Method:             "GET",
		PathPattern:        "/agreement/admin/policy-types",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveAllPolicyTypesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveAllPolicyTypesOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: Use RetrievePolicyCountryShort instead.

  RetrievePolicyCountry retrieves a base legal policy based on a particular country

  Retrieve a Base Legal Policy based on a Particular Country.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=2 (READ)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrievePolicyCountry(params *RetrievePolicyCountryParams, authInfo runtime.ClientAuthInfoWriter) (*RetrievePolicyCountryOK, *RetrievePolicyCountryNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrievePolicyCountryParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrievePolicyCountry",
		Method:             "GET",
		PathPattern:        "/agreement/admin/base-policies/{basePolicyId}/countries/{countryCode}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrievePolicyCountryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *RetrievePolicyCountryOK:
		return v, nil, nil

	case *RetrievePolicyCountryNotFound:
		return nil, v, nil

	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrievePolicyCountryShort retrieves a base legal policy based on a particular country

  Retrieve a Base Legal Policy based on a Particular Country.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=2 (READ)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrievePolicyCountryShort(params *RetrievePolicyCountryParams, authInfo runtime.ClientAuthInfoWriter) (*RetrievePolicyCountryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrievePolicyCountryParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrievePolicyCountry",
		Method:             "GET",
		PathPattern:        "/agreement/admin/base-policies/{basePolicyId}/countries/{countryCode}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrievePolicyCountryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrievePolicyCountryOK:
		return v, nil
	case *RetrievePolicyCountryNotFound:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: Use RetrieveSinglePolicyShort instead.

  RetrieveSinglePolicy retrieves a base legal policy

  Retrieve a base policy.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=2 (READ)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrieveSinglePolicy(params *RetrieveSinglePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveSinglePolicyOK, *RetrieveSinglePolicyNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveSinglePolicyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveSinglePolicy",
		Method:             "GET",
		PathPattern:        "/agreement/admin/base-policies/{basePolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveSinglePolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *RetrieveSinglePolicyOK:
		return v, nil, nil

	case *RetrieveSinglePolicyNotFound:
		return nil, v, nil

	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrieveSinglePolicyShort retrieves a base legal policy

  Retrieve a base policy.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=2 (READ)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrieveSinglePolicyShort(params *RetrieveSinglePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveSinglePolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveSinglePolicyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveSinglePolicy",
		Method:             "GET",
		PathPattern:        "/agreement/admin/base-policies/{basePolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveSinglePolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveSinglePolicyOK:
		return v, nil
	case *RetrieveSinglePolicyNotFound:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
