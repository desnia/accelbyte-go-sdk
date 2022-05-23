// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package agreement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new agreement API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for agreement API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AcceptVersionedPolicy(params *AcceptVersionedPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*AcceptVersionedPolicyCreated, error)
	AcceptVersionedPolicyShort(params *AcceptVersionedPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*AcceptVersionedPolicyCreated, error)
	BulkAcceptVersionedPolicy(params *BulkAcceptVersionedPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*BulkAcceptVersionedPolicyCreated, error)
	BulkAcceptVersionedPolicyShort(params *BulkAcceptVersionedPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*BulkAcceptVersionedPolicyCreated, error)
	ChangePreferenceConsent(params *ChangePreferenceConsentParams, authInfo runtime.ClientAuthInfoWriter) (*ChangePreferenceConsentOK, error)
	ChangePreferenceConsentShort(params *ChangePreferenceConsentParams, authInfo runtime.ClientAuthInfoWriter) (*ChangePreferenceConsentOK, error)
	ChangePreferenceConsent1(params *ChangePreferenceConsent1Params, authInfo runtime.ClientAuthInfoWriter) (*ChangePreferenceConsent1OK, *ChangePreferenceConsent1BadRequest, error)
	ChangePreferenceConsent1Short(params *ChangePreferenceConsent1Params, authInfo runtime.ClientAuthInfoWriter) (*ChangePreferenceConsent1OK, error)
	IndirectBulkAcceptVersionedPolicyV2(params *IndirectBulkAcceptVersionedPolicyV2Params, authInfo runtime.ClientAuthInfoWriter) (*IndirectBulkAcceptVersionedPolicyV2Created, error)
	IndirectBulkAcceptVersionedPolicyV2Short(params *IndirectBulkAcceptVersionedPolicyV2Params, authInfo runtime.ClientAuthInfoWriter) (*IndirectBulkAcceptVersionedPolicyV2Created, error)
	IndirectBulkAcceptVersionedPolicy1(params *IndirectBulkAcceptVersionedPolicy1Params, authInfo runtime.ClientAuthInfoWriter) (*IndirectBulkAcceptVersionedPolicy1Created, error)
	IndirectBulkAcceptVersionedPolicy1Short(params *IndirectBulkAcceptVersionedPolicy1Params, authInfo runtime.ClientAuthInfoWriter) (*IndirectBulkAcceptVersionedPolicy1Created, error)
	RetrieveAcceptedAgreements(params *RetrieveAcceptedAgreementsParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAcceptedAgreementsOK, error)
	RetrieveAcceptedAgreementsShort(params *RetrieveAcceptedAgreementsParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAcceptedAgreementsOK, error)
	RetrieveAgreementsPublic(params *RetrieveAgreementsPublicParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAgreementsPublicOK, error)
	RetrieveAgreementsPublicShort(params *RetrieveAgreementsPublicParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAgreementsPublicOK, error)
	RetrieveAllUsersByPolicyVersion(params *RetrieveAllUsersByPolicyVersionParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAllUsersByPolicyVersionOK, *RetrieveAllUsersByPolicyVersionNotFound, error)
	RetrieveAllUsersByPolicyVersionShort(params *RetrieveAllUsersByPolicyVersionParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAllUsersByPolicyVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AcceptVersionedPolicy accepts a policy version

  Accepts a legal policy version. Supply with localized version policy id to accept an agreement.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: login user&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) AcceptVersionedPolicy(params *AcceptVersionedPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*AcceptVersionedPolicyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAcceptVersionedPolicyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClient(&http.Client{Transport: params.RetryPolicy})
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "acceptVersionedPolicy",
		Method:             "POST",
		PathPattern:        "/agreement/public/agreements/localized-policy-versions/{localizedPolicyVersionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AcceptVersionedPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *AcceptVersionedPolicyCreated:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) AcceptVersionedPolicyShort(params *AcceptVersionedPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*AcceptVersionedPolicyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAcceptVersionedPolicyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "acceptVersionedPolicy",
		Method:             "POST",
		PathPattern:        "/agreement/public/agreements/localized-policy-versions/{localizedPolicyVersionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AcceptVersionedPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *AcceptVersionedPolicyCreated:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  BulkAcceptVersionedPolicy bulks accept policy versions

  Accepts many legal policy versions all at once. Supply with localized version policy id to accept an agreement.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: login user&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) BulkAcceptVersionedPolicy(params *BulkAcceptVersionedPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*BulkAcceptVersionedPolicyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBulkAcceptVersionedPolicyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClient(&http.Client{Transport: params.RetryPolicy})
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "bulkAcceptVersionedPolicy",
		Method:             "POST",
		PathPattern:        "/agreement/public/agreements/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BulkAcceptVersionedPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *BulkAcceptVersionedPolicyCreated:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) BulkAcceptVersionedPolicyShort(params *BulkAcceptVersionedPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*BulkAcceptVersionedPolicyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBulkAcceptVersionedPolicyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "bulkAcceptVersionedPolicy",
		Method:             "POST",
		PathPattern:        "/agreement/public/agreements/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BulkAcceptVersionedPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *BulkAcceptVersionedPolicyCreated:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  ChangePreferenceConsent change preference consent API
*/
func (a *Client) ChangePreferenceConsent(params *ChangePreferenceConsentParams, authInfo runtime.ClientAuthInfoWriter) (*ChangePreferenceConsentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangePreferenceConsentParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClient(&http.Client{Transport: params.RetryPolicy})
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changePreferenceConsent",
		Method:             "PATCH",
		PathPattern:        "/agreement/admin/agreements/localized-policy-versions/preferences/namespaces/{namespace}/userId/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangePreferenceConsentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *ChangePreferenceConsentOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) ChangePreferenceConsentShort(params *ChangePreferenceConsentParams, authInfo runtime.ClientAuthInfoWriter) (*ChangePreferenceConsentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangePreferenceConsentParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changePreferenceConsent",
		Method:             "PATCH",
		PathPattern:        "/agreement/admin/agreements/localized-policy-versions/preferences/namespaces/{namespace}/userId/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangePreferenceConsentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *ChangePreferenceConsentOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  ChangePreferenceConsent1 accepts revoke marketing preference consent

  Change marketing preference consent.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: login user&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) ChangePreferenceConsent1(params *ChangePreferenceConsent1Params, authInfo runtime.ClientAuthInfoWriter) (*ChangePreferenceConsent1OK, *ChangePreferenceConsent1BadRequest, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangePreferenceConsent1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClient(&http.Client{Transport: params.RetryPolicy})
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changePreferenceConsent_1",
		Method:             "PATCH",
		PathPattern:        "/agreement/public/agreements/localized-policy-versions/preferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangePreferenceConsent1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *ChangePreferenceConsent1OK:
		return v, nil, nil

	case *ChangePreferenceConsent1BadRequest:
		return nil, v, nil

	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) ChangePreferenceConsent1Short(params *ChangePreferenceConsent1Params, authInfo runtime.ClientAuthInfoWriter) (*ChangePreferenceConsent1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangePreferenceConsent1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changePreferenceConsent_1",
		Method:             "PATCH",
		PathPattern:        "/agreement/public/agreements/localized-policy-versions/preferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangePreferenceConsent1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *ChangePreferenceConsent1OK:
		return v, nil
	case *ChangePreferenceConsent1BadRequest:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  IndirectBulkAcceptVersionedPolicyV2 bulks accept policy versions indirect

  &lt;b&gt;IMPORTANT: GOING TO DEPRECATE&lt;/b&gt;&lt;br/&gt;&lt;br/&gt;Accepts many legal policy versions all at once. Supply with localized version policy id, version policy id, policy id, userId, namespace, country code and client id to accept an agreement. This endpoint used by APIGateway during new user registration.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;NAMESPACE:{namespace}:LEGAL&#34;, action=1 (CREATE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) IndirectBulkAcceptVersionedPolicyV2(params *IndirectBulkAcceptVersionedPolicyV2Params, authInfo runtime.ClientAuthInfoWriter) (*IndirectBulkAcceptVersionedPolicyV2Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndirectBulkAcceptVersionedPolicyV2Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClient(&http.Client{Transport: params.RetryPolicy})
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indirectBulkAcceptVersionedPolicyV2",
		Method:             "POST",
		PathPattern:        "/agreement/public/agreements/policies/namespaces/{namespace}/countries/{countryCode}/clients/{clientId}/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IndirectBulkAcceptVersionedPolicyV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *IndirectBulkAcceptVersionedPolicyV2Created:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) IndirectBulkAcceptVersionedPolicyV2Short(params *IndirectBulkAcceptVersionedPolicyV2Params, authInfo runtime.ClientAuthInfoWriter) (*IndirectBulkAcceptVersionedPolicyV2Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndirectBulkAcceptVersionedPolicyV2Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indirectBulkAcceptVersionedPolicyV2",
		Method:             "POST",
		PathPattern:        "/agreement/public/agreements/policies/namespaces/{namespace}/countries/{countryCode}/clients/{clientId}/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IndirectBulkAcceptVersionedPolicyV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *IndirectBulkAcceptVersionedPolicyV2Created:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  IndirectBulkAcceptVersionedPolicy1 bulks accept policy versions indirect

  Accepts many legal policy versions all at once. Supply with localized version policy id and userId to accept an agreement. This endpoint used by Authentication Service during new user registration.&lt;br&gt;&lt;br/&gt;Available Extra Information to return: &lt;br/&gt;&lt;ul&gt;&lt;li&gt;&lt;b&gt;userIds&lt;/b&gt; : List of userId mapping (&lt;b&gt;IMPORTANT: GOING TO DEPRECATE&lt;/b&gt;)&lt;/li&gt;&lt;/ul&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: login user&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) IndirectBulkAcceptVersionedPolicy1(params *IndirectBulkAcceptVersionedPolicy1Params, authInfo runtime.ClientAuthInfoWriter) (*IndirectBulkAcceptVersionedPolicy1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndirectBulkAcceptVersionedPolicy1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClient(&http.Client{Transport: params.RetryPolicy})
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indirectBulkAcceptVersionedPolicy_1",
		Method:             "POST",
		PathPattern:        "/agreement/public/agreements/policies/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IndirectBulkAcceptVersionedPolicy1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *IndirectBulkAcceptVersionedPolicy1Created:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) IndirectBulkAcceptVersionedPolicy1Short(params *IndirectBulkAcceptVersionedPolicy1Params, authInfo runtime.ClientAuthInfoWriter) (*IndirectBulkAcceptVersionedPolicy1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndirectBulkAcceptVersionedPolicy1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indirectBulkAcceptVersionedPolicy_1",
		Method:             "POST",
		PathPattern:        "/agreement/public/agreements/policies/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IndirectBulkAcceptVersionedPolicy1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *IndirectBulkAcceptVersionedPolicy1Created:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrieveAcceptedAgreements retrieves accepted legal agreements

  This API will return all accepted Legal Agreements for specified user. Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=2 (READ)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrieveAcceptedAgreements(params *RetrieveAcceptedAgreementsParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAcceptedAgreementsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveAcceptedAgreementsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClient(&http.Client{Transport: params.RetryPolicy})
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveAcceptedAgreements",
		Method:             "GET",
		PathPattern:        "/agreement/admin/agreements/policies/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveAcceptedAgreementsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveAcceptedAgreementsOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) RetrieveAcceptedAgreementsShort(params *RetrieveAcceptedAgreementsParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAcceptedAgreementsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveAcceptedAgreementsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveAcceptedAgreements",
		Method:             "GET",
		PathPattern:        "/agreement/admin/agreements/policies/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveAcceptedAgreementsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveAcceptedAgreementsOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrieveAgreementsPublic retrieves the accepted legal agreements

  Retrieve accepted Legal Agreements.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: login user&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrieveAgreementsPublic(params *RetrieveAgreementsPublicParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAgreementsPublicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveAgreementsPublicParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClient(&http.Client{Transport: params.RetryPolicy})
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveAgreementsPublic",
		Method:             "GET",
		PathPattern:        "/agreement/public/agreements/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveAgreementsPublicReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveAgreementsPublicOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) RetrieveAgreementsPublicShort(params *RetrieveAgreementsPublicParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAgreementsPublicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveAgreementsPublicParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveAgreementsPublic",
		Method:             "GET",
		PathPattern:        "/agreement/public/agreements/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveAgreementsPublicReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveAgreementsPublicOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrieveAllUsersByPolicyVersion retrieves all users accepting legal agreements

  This API will return all users who has accepted a specific policy version.Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:*:LEGAL&#34;, action=2 (READ)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) RetrieveAllUsersByPolicyVersion(params *RetrieveAllUsersByPolicyVersionParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAllUsersByPolicyVersionOK, *RetrieveAllUsersByPolicyVersionNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveAllUsersByPolicyVersionParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClient(&http.Client{Transport: params.RetryPolicy})
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveAllUsersByPolicyVersion",
		Method:             "GET",
		PathPattern:        "/agreement/admin/agreements/policy-versions/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveAllUsersByPolicyVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *RetrieveAllUsersByPolicyVersionOK:
		return v, nil, nil

	case *RetrieveAllUsersByPolicyVersionNotFound:
		return nil, v, nil

	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) RetrieveAllUsersByPolicyVersionShort(params *RetrieveAllUsersByPolicyVersionParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAllUsersByPolicyVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveAllUsersByPolicyVersionParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveAllUsersByPolicyVersion",
		Method:             "GET",
		PathPattern:        "/agreement/admin/agreements/policy-versions/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveAllUsersByPolicyVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveAllUsersByPolicyVersionOK:
		return v, nil
	case *RetrieveAllUsersByPolicyVersionNotFound:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
