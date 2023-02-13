// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	RetrieveLatestPolicies(params *RetrieveLatestPoliciesParams) (*RetrieveLatestPoliciesOK, error)
	RetrieveLatestPoliciesShort(params *RetrieveLatestPoliciesParams) (*RetrieveLatestPoliciesOK, error)
	RetrieveLatestPoliciesByNamespaceAndCountryPublic(params *RetrieveLatestPoliciesByNamespaceAndCountryPublicParams) (*RetrieveLatestPoliciesByNamespaceAndCountryPublicOK, error)
	RetrieveLatestPoliciesByNamespaceAndCountryPublicShort(params *RetrieveLatestPoliciesByNamespaceAndCountryPublicParams) (*RetrieveLatestPoliciesByNamespaceAndCountryPublicOK, error)
	RetrieveLatestPoliciesPublic(params *RetrieveLatestPoliciesPublicParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveLatestPoliciesPublicOK, *RetrieveLatestPoliciesPublicNotFound, error)
	RetrieveLatestPoliciesPublicShort(params *RetrieveLatestPoliciesPublicParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveLatestPoliciesPublicOK, error)
	RetrievePolicies(params *RetrievePoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*RetrievePoliciesOK, error)
	RetrievePoliciesShort(params *RetrievePoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*RetrievePoliciesOK, error)
	SetDefaultPolicy2(params *SetDefaultPolicy2Params, authInfo runtime.ClientAuthInfoWriter) (*SetDefaultPolicy2OK, *SetDefaultPolicy2BadRequest, error)
	SetDefaultPolicy2Short(params *SetDefaultPolicy2Params, authInfo runtime.ClientAuthInfoWriter) (*SetDefaultPolicy2OK, error)
	UpdatePolicy(params *UpdatePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePolicyOK, *UpdatePolicyBadRequest, error)
	UpdatePolicyShort(params *UpdatePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePolicyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Deprecated: Use RetrieveLatestPoliciesShort instead.

  RetrieveLatestPolicies retrieves latest policies by country

  Retrieve all active latest policies based on a country.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Leave the policyType empty if you want to be responded with all policy type&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Fill the tags if you want to filter the responded policy by tags&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Fill the defaultOnEmpty with true if you want to be responded with default country-specific policy if your requested country is not exist&lt;/i&gt;&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrieveLatestPolicies(params *RetrieveLatestPoliciesParams) (*RetrieveLatestPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveLatestPoliciesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveLatestPolicies",
		Method:             "GET",
		PathPattern:        "/agreement/public/policies/countries/{countryCode}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveLatestPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveLatestPoliciesOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrieveLatestPoliciesShort retrieves latest policies by country

  Retrieve all active latest policies based on a country.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Leave the policyType empty if you want to be responded with all policy type&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Fill the tags if you want to filter the responded policy by tags&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Fill the defaultOnEmpty with true if you want to be responded with default country-specific policy if your requested country is not exist&lt;/i&gt;&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrieveLatestPoliciesShort(params *RetrieveLatestPoliciesParams) (*RetrieveLatestPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveLatestPoliciesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveLatestPolicies",
		Method:             "GET",
		PathPattern:        "/agreement/public/policies/countries/{countryCode}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveLatestPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveLatestPoliciesOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: Use RetrieveLatestPoliciesByNamespaceAndCountryPublicShort instead.

  RetrieveLatestPoliciesByNamespaceAndCountryPublic retrieves latest policies by namespace and country

  Retrieve all active latest policies based on a namespace and country.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Leave the policyType empty if you want to be responded with all policy type&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Fill the tags if you want to filter the responded policy by tags&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Fill the defaultOnEmpty with true if you want to be responded with default country-specific policy if your requested country is not exist&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Fill the alwaysIncludeDefault with true if you want to be responded with always include default policy. If there are duplicate policies (default policies and country specific policies with same base policy) it&#39;ll include policy with same country code, for example:&lt;ul&gt;&lt;li&gt;Document 1 (default): Region US (default), UA&lt;/li&gt;&lt;li&gt;Document 2 (default): Region US (default)&lt;/li&gt;&lt;li&gt;Document 3 (default): Region US (default)&lt;/li&gt;&lt;li&gt;User: Region UA&lt;/li&gt;&lt;li&gt;Query: alwaysIncludeDefault: true&lt;/li&gt;&lt;li&gt;Response: Document 1 (UA), Document 2 (US), Document 3 (US)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrieveLatestPoliciesByNamespaceAndCountryPublic(params *RetrieveLatestPoliciesByNamespaceAndCountryPublicParams) (*RetrieveLatestPoliciesByNamespaceAndCountryPublicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveLatestPoliciesByNamespaceAndCountryPublicParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveLatestPoliciesByNamespaceAndCountryPublic",
		Method:             "GET",
		PathPattern:        "/agreement/public/policies/namespaces/{namespace}/countries/{countryCode}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveLatestPoliciesByNamespaceAndCountryPublicReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveLatestPoliciesByNamespaceAndCountryPublicOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrieveLatestPoliciesByNamespaceAndCountryPublicShort retrieves latest policies by namespace and country

  Retrieve all active latest policies based on a namespace and country.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Leave the policyType empty if you want to be responded with all policy type&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Fill the tags if you want to filter the responded policy by tags&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Fill the defaultOnEmpty with true if you want to be responded with default country-specific policy if your requested country is not exist&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Fill the alwaysIncludeDefault with true if you want to be responded with always include default policy. If there are duplicate policies (default policies and country specific policies with same base policy) it&#39;ll include policy with same country code, for example:&lt;ul&gt;&lt;li&gt;Document 1 (default): Region US (default), UA&lt;/li&gt;&lt;li&gt;Document 2 (default): Region US (default)&lt;/li&gt;&lt;li&gt;Document 3 (default): Region US (default)&lt;/li&gt;&lt;li&gt;User: Region UA&lt;/li&gt;&lt;li&gt;Query: alwaysIncludeDefault: true&lt;/li&gt;&lt;li&gt;Response: Document 1 (UA), Document 2 (US), Document 3 (US)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrieveLatestPoliciesByNamespaceAndCountryPublicShort(params *RetrieveLatestPoliciesByNamespaceAndCountryPublicParams) (*RetrieveLatestPoliciesByNamespaceAndCountryPublicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveLatestPoliciesByNamespaceAndCountryPublicParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveLatestPoliciesByNamespaceAndCountryPublic",
		Method:             "GET",
		PathPattern:        "/agreement/public/policies/namespaces/{namespace}/countries/{countryCode}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveLatestPoliciesByNamespaceAndCountryPublicReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveLatestPoliciesByNamespaceAndCountryPublicOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: Use RetrieveLatestPoliciesPublicShort instead.

  RetrieveLatestPoliciesPublic retrieves latest policies by namespace and country

  Retrieve all active latest policies based on a namespace and country. The country will be read from user token.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Leave the policyType empty if you want to be responded with all policy type&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Fill the tags if you want to filter the responded policy by tags&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Fill the defaultOnEmpty with true if you want to be responded with default country-specific policy if your requested country is not exist&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Fill the alwaysIncludeDefault with true if you want to be responded with always include default policy. If there are duplicate policies (default policies and country specific policies with same base policy) it&#39;ll include policy with same country code, for example:&lt;ul&gt;&lt;li&gt;Document 1 (default): Region US (default), UA&lt;/li&gt;&lt;li&gt;Document 2 (default): Region US (default)&lt;/li&gt;&lt;li&gt;Document 3 (default): Region US (default)&lt;/li&gt;&lt;li&gt;User: Region UA&lt;/li&gt;&lt;li&gt;Query: alwaysIncludeDefault: true&lt;/li&gt;&lt;li&gt;Response: Document 1 (UA), Document 2 (US), Document 3 (US)&lt;/li&gt;&lt;/ul&gt;&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Required permission: login user &lt;/i&gt;&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrieveLatestPoliciesPublic(params *RetrieveLatestPoliciesPublicParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveLatestPoliciesPublicOK, *RetrieveLatestPoliciesPublicNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveLatestPoliciesPublicParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveLatestPoliciesPublic",
		Method:             "GET",
		PathPattern:        "/agreement/public/policies/namespaces/{namespace}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveLatestPoliciesPublicReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *RetrieveLatestPoliciesPublicOK:
		return v, nil, nil

	case *RetrieveLatestPoliciesPublicNotFound:
		return nil, v, nil

	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrieveLatestPoliciesPublicShort retrieves latest policies by namespace and country

  Retrieve all active latest policies based on a namespace and country. The country will be read from user token.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Leave the policyType empty if you want to be responded with all policy type&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Fill the tags if you want to filter the responded policy by tags&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Fill the defaultOnEmpty with true if you want to be responded with default country-specific policy if your requested country is not exist&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Fill the alwaysIncludeDefault with true if you want to be responded with always include default policy. If there are duplicate policies (default policies and country specific policies with same base policy) it&#39;ll include policy with same country code, for example:&lt;ul&gt;&lt;li&gt;Document 1 (default): Region US (default), UA&lt;/li&gt;&lt;li&gt;Document 2 (default): Region US (default)&lt;/li&gt;&lt;li&gt;Document 3 (default): Region US (default)&lt;/li&gt;&lt;li&gt;User: Region UA&lt;/li&gt;&lt;li&gt;Query: alwaysIncludeDefault: true&lt;/li&gt;&lt;li&gt;Response: Document 1 (UA), Document 2 (US), Document 3 (US)&lt;/li&gt;&lt;/ul&gt;&lt;/i&gt;&lt;/li&gt;&lt;li&gt;&lt;i&gt;Required permission: login user &lt;/i&gt;&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrieveLatestPoliciesPublicShort(params *RetrieveLatestPoliciesPublicParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveLatestPoliciesPublicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveLatestPoliciesPublicParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveLatestPoliciesPublic",
		Method:             "GET",
		PathPattern:        "/agreement/public/policies/namespaces/{namespace}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveLatestPoliciesPublicReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveLatestPoliciesPublicOK:
		return v, nil
	case *RetrieveLatestPoliciesPublicNotFound:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: Use RetrievePoliciesShort instead.

  RetrievePolicies retrieves policies by country

  Retrieve all active policies based on a country.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=2 (READ)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrievePolicies(params *RetrievePoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*RetrievePoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrievePoliciesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrievePolicies",
		Method:             "GET",
		PathPattern:        "/agreement/admin/policies/countries/{countryCode}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrievePoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrievePoliciesOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrievePoliciesShort retrieves policies by country

  Retrieve all active policies based on a country.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=2 (READ)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrievePoliciesShort(params *RetrievePoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*RetrievePoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrievePoliciesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrievePolicies",
		Method:             "GET",
		PathPattern:        "/agreement/admin/policies/countries/{countryCode}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrievePoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrievePoliciesOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: Use SetDefaultPolicy2Short instead.

  SetDefaultPolicy2 sets default policy

  Update a policy to be the default.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=4 (UPDATE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) SetDefaultPolicy2(params *SetDefaultPolicy2Params, authInfo runtime.ClientAuthInfoWriter) (*SetDefaultPolicy2OK, *SetDefaultPolicy2BadRequest, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetDefaultPolicy2Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setDefaultPolicy_2",
		Method:             "PATCH",
		PathPattern:        "/agreement/admin/policies/{policyId}/default",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetDefaultPolicy2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *SetDefaultPolicy2OK:
		return v, nil, nil

	case *SetDefaultPolicy2BadRequest:
		return nil, v, nil

	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  SetDefaultPolicy2Short sets default policy

  Update a policy to be the default.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=4 (UPDATE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) SetDefaultPolicy2Short(params *SetDefaultPolicy2Params, authInfo runtime.ClientAuthInfoWriter) (*SetDefaultPolicy2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetDefaultPolicy2Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setDefaultPolicy_2",
		Method:             "PATCH",
		PathPattern:        "/agreement/admin/policies/{policyId}/default",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetDefaultPolicy2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *SetDefaultPolicy2OK:
		return v, nil
	case *SetDefaultPolicy2BadRequest:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: Use UpdatePolicyShort instead.

  UpdatePolicy updates country specific policy

  Update country-specific policy.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=4 (UPDATE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) UpdatePolicy(params *UpdatePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePolicyOK, *UpdatePolicyBadRequest, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePolicyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updatePolicy",
		Method:             "PATCH",
		PathPattern:        "/agreement/admin/policies/{policyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *UpdatePolicyOK:
		return v, nil, nil

	case *UpdatePolicyBadRequest:
		return nil, v, nil

	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdatePolicyShort updates country specific policy

  Update country-specific policy.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=4 (UPDATE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) UpdatePolicyShort(params *UpdatePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePolicyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updatePolicy",
		Method:             "PATCH",
		PathPattern:        "/agreement/admin/policies/{policyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *UpdatePolicyOK:
		return v, nil
	case *UpdatePolicyBadRequest:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
