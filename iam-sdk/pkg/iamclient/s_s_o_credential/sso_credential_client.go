// Code generated by go-swagger; DO NOT EDIT.

package s_s_o_credential

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new s s o credential API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for s s o credential API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddSSOLoginPlatformCredential(params *AddSSOLoginPlatformCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*AddSSOLoginPlatformCredentialCreated, *AddSSOLoginPlatformCredentialBadRequest, *AddSSOLoginPlatformCredentialUnauthorized, *AddSSOLoginPlatformCredentialForbidden, *AddSSOLoginPlatformCredentialInternalServerError, error)

	DeleteSSOLoginPlatformCredentialV3(params *DeleteSSOLoginPlatformCredentialV3Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteSSOLoginPlatformCredentialV3NoContent, *DeleteSSOLoginPlatformCredentialV3Unauthorized, *DeleteSSOLoginPlatformCredentialV3Forbidden, *DeleteSSOLoginPlatformCredentialV3NotFound, *DeleteSSOLoginPlatformCredentialV3InternalServerError, error)

	RetrieveAllSSOLoginPlatformCredentialV3(params *RetrieveAllSSOLoginPlatformCredentialV3Params, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAllSSOLoginPlatformCredentialV3OK, *RetrieveAllSSOLoginPlatformCredentialV3Unauthorized, *RetrieveAllSSOLoginPlatformCredentialV3Forbidden, *RetrieveAllSSOLoginPlatformCredentialV3NotFound, *RetrieveAllSSOLoginPlatformCredentialV3InternalServerError, error)

	RetrieveSSOLoginPlatformCredential(params *RetrieveSSOLoginPlatformCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveSSOLoginPlatformCredentialOK, *RetrieveSSOLoginPlatformCredentialUnauthorized, *RetrieveSSOLoginPlatformCredentialForbidden, *RetrieveSSOLoginPlatformCredentialNotFound, *RetrieveSSOLoginPlatformCredentialInternalServerError, error)

	UpdateSSOPlatformCredential(params *UpdateSSOPlatformCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSSOPlatformCredentialOK, *UpdateSSOPlatformCredentialBadRequest, *UpdateSSOPlatformCredentialUnauthorized, *UpdateSSOPlatformCredentialForbidden, *UpdateSSOPlatformCredentialNotFound, *UpdateSSOPlatformCredentialInternalServerError, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddSSOLoginPlatformCredential adds s s o platform credential

  This is the API to Add SSO Platform Credential. It needs ADMIN:NAMESPACE:{namespace}:PLATFORM:{platformId}:SSO [CREATE] resource.<h2>Supported platforms:</h2><ul>
			<li><strong>discourse</strong></li>the ssoUrl of the discourse is the discourse forum url. example: https://forum.example.com
			<li><strong>azure with SAML</strong></li><b>appId</b> is an application identifier in IdP, in azure it's called EntityID
			<b>acsUrl</b> is an endpoint on the service provider where the identity provider will redirect to with its authentication response. example: /iam/v3/sso/saml/azuresaml/authenticate
			<b>federationMetadataUrl</b> is an endpoint on the Identity Provider(IdP) to get IdP federation metadata for service provider to build trust relationship
			</ul>
*/
func (a *Client) AddSSOLoginPlatformCredential(params *AddSSOLoginPlatformCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*AddSSOLoginPlatformCredentialCreated, *AddSSOLoginPlatformCredentialBadRequest, *AddSSOLoginPlatformCredentialUnauthorized, *AddSSOLoginPlatformCredentialForbidden, *AddSSOLoginPlatformCredentialInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddSSOLoginPlatformCredentialParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddSSOLoginPlatformCredential",
		Method:             "POST",
		PathPattern:        "/iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddSSOLoginPlatformCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AddSSOLoginPlatformCredentialCreated:
		return v, nil, nil, nil, nil, nil
	case *AddSSOLoginPlatformCredentialBadRequest:
		return nil, v, nil, nil, nil, nil
	case *AddSSOLoginPlatformCredentialUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *AddSSOLoginPlatformCredentialForbidden:
		return nil, nil, nil, v, nil, nil
	case *AddSSOLoginPlatformCredentialInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteSSOLoginPlatformCredentialV3 deletes s s o platform credential

  This is the API to Delete SSO Platform Credential. It needs ADMIN:NAMESPACE:{namespace}:PLATFORM:{platformId}:SSO [DELETE] resource
*/
func (a *Client) DeleteSSOLoginPlatformCredentialV3(params *DeleteSSOLoginPlatformCredentialV3Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteSSOLoginPlatformCredentialV3NoContent, *DeleteSSOLoginPlatformCredentialV3Unauthorized, *DeleteSSOLoginPlatformCredentialV3Forbidden, *DeleteSSOLoginPlatformCredentialV3NotFound, *DeleteSSOLoginPlatformCredentialV3InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSSOLoginPlatformCredentialV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteSSOLoginPlatformCredentialV3",
		Method:             "DELETE",
		PathPattern:        "/iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSSOLoginPlatformCredentialV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteSSOLoginPlatformCredentialV3NoContent:
		return v, nil, nil, nil, nil, nil
	case *DeleteSSOLoginPlatformCredentialV3Unauthorized:
		return nil, v, nil, nil, nil, nil
	case *DeleteSSOLoginPlatformCredentialV3Forbidden:
		return nil, nil, v, nil, nil, nil
	case *DeleteSSOLoginPlatformCredentialV3NotFound:
		return nil, nil, nil, v, nil, nil
	case *DeleteSSOLoginPlatformCredentialV3InternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrieveAllSSOLoginPlatformCredentialV3 gets all s s o platform credential

  This is the API to Get All Active SSO Platform Credential. It needs ADMIN:NAMESPACE:{namespace}:PLATFORM:*:SSO [READ] resource
*/
func (a *Client) RetrieveAllSSOLoginPlatformCredentialV3(params *RetrieveAllSSOLoginPlatformCredentialV3Params, authInfo runtime.ClientAuthInfoWriter) (*RetrieveAllSSOLoginPlatformCredentialV3OK, *RetrieveAllSSOLoginPlatformCredentialV3Unauthorized, *RetrieveAllSSOLoginPlatformCredentialV3Forbidden, *RetrieveAllSSOLoginPlatformCredentialV3NotFound, *RetrieveAllSSOLoginPlatformCredentialV3InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveAllSSOLoginPlatformCredentialV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RetrieveAllSSOLoginPlatformCredentialV3",
		Method:             "GET",
		PathPattern:        "/iam/v3/admin/namespaces/{namespace}/platforms/sso",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RetrieveAllSSOLoginPlatformCredentialV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *RetrieveAllSSOLoginPlatformCredentialV3OK:
		return v, nil, nil, nil, nil, nil
	case *RetrieveAllSSOLoginPlatformCredentialV3Unauthorized:
		return nil, v, nil, nil, nil, nil
	case *RetrieveAllSSOLoginPlatformCredentialV3Forbidden:
		return nil, nil, v, nil, nil, nil
	case *RetrieveAllSSOLoginPlatformCredentialV3NotFound:
		return nil, nil, nil, v, nil, nil
	case *RetrieveAllSSOLoginPlatformCredentialV3InternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrieveSSOLoginPlatformCredential retrieves s s o platform credential

  This is the API to Get SSO Platform Credential. It needs ADMIN:NAMESPACE:{namespace}:PLATFORM:{platformId}:SSO [READ] resource
*/
func (a *Client) RetrieveSSOLoginPlatformCredential(params *RetrieveSSOLoginPlatformCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveSSOLoginPlatformCredentialOK, *RetrieveSSOLoginPlatformCredentialUnauthorized, *RetrieveSSOLoginPlatformCredentialForbidden, *RetrieveSSOLoginPlatformCredentialNotFound, *RetrieveSSOLoginPlatformCredentialInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveSSOLoginPlatformCredentialParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RetrieveSSOLoginPlatformCredential",
		Method:             "GET",
		PathPattern:        "/iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RetrieveSSOLoginPlatformCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *RetrieveSSOLoginPlatformCredentialOK:
		return v, nil, nil, nil, nil, nil
	case *RetrieveSSOLoginPlatformCredentialUnauthorized:
		return nil, v, nil, nil, nil, nil
	case *RetrieveSSOLoginPlatformCredentialForbidden:
		return nil, nil, v, nil, nil, nil
	case *RetrieveSSOLoginPlatformCredentialNotFound:
		return nil, nil, nil, v, nil, nil
	case *RetrieveSSOLoginPlatformCredentialInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateSSOPlatformCredential updates s s o platform credential

  This is the API to Delete SSO Platform Credential. It needs ADMIN:NAMESPACE:{namespace}:PLATFORM:{platformId}:SSO [UPDATE] resource
*/
func (a *Client) UpdateSSOPlatformCredential(params *UpdateSSOPlatformCredentialParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSSOPlatformCredentialOK, *UpdateSSOPlatformCredentialBadRequest, *UpdateSSOPlatformCredentialUnauthorized, *UpdateSSOPlatformCredentialForbidden, *UpdateSSOPlatformCredentialNotFound, *UpdateSSOPlatformCredentialInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSSOPlatformCredentialParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateSSOPlatformCredential",
		Method:             "PATCH",
		PathPattern:        "/iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSSOPlatformCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateSSOPlatformCredentialOK:
		return v, nil, nil, nil, nil, nil, nil
	case *UpdateSSOPlatformCredentialBadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *UpdateSSOPlatformCredentialUnauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *UpdateSSOPlatformCredentialForbidden:
		return nil, nil, nil, v, nil, nil, nil
	case *UpdateSSOPlatformCredentialNotFound:
		return nil, nil, nil, nil, v, nil, nil
	case *UpdateSSOPlatformCredentialInternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}