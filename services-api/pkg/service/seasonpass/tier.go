// Copyright (c) 2018 - 2021
// AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package seasonpass

import (
	"github.com/AccelByte/accelbyte-go-sdk/seasonpass-sdk/pkg/seasonpassclient"
	"github.com/AccelByte/accelbyte-go-sdk/seasonpass-sdk/pkg/seasonpassclient/tier"
	"github.com/AccelByte/accelbyte-go-sdk/seasonpass-sdk/pkg/seasonpassclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type TierService struct {
	Client          *seasonpassclient.JusticeSeasonpassService
	TokenRepository repository.TokenRepository
}

// Deprecated: Use UpdateTierShort instead
func (t *TierService) UpdateTier(input *tier.UpdateTierParams) (*seasonpassclientmodels.Tier, error) {
	accessToken, err := t.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, badRequest, conflict, unprocessableEntity, err := t.Client.Tier.UpdateTier(input, client.BearerToken(*accessToken.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if badRequest != nil {
		return nil, badRequest
	}
	if conflict != nil {
		return nil, conflict
	}
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use DeleteTierShort instead
func (t *TierService) DeleteTier(input *tier.DeleteTierParams) error {
	accessToken, err := t.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, notFound, conflict, badRequest, err := t.Client.Tier.DeleteTier(input, client.BearerToken(*accessToken.AccessToken))
	if notFound != nil {
		return notFound
	}
	if conflict != nil {
		return conflict
	}
	if badRequest != nil {
		return badRequest
	}
	if err != nil {
		return err
	}
	return nil
}

// Deprecated: Use QueryTiersShort instead
func (t *TierService) QueryTiers(input *tier.QueryTiersParams) (*seasonpassclientmodels.TierPagingSlicedResult, error) {
	accessToken, err := t.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, err := t.Client.Tier.QueryTiers(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use CreateTierShort instead
func (t *TierService) CreateTier(input *tier.CreateTierParams) ([]*seasonpassclientmodels.Tier, error) {
	accessToken, err := t.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, notFound, badRequest, conflict, unprocessableEntity, err := t.Client.Tier.CreateTier(input, client.BearerToken(*accessToken.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if badRequest != nil {
		return nil, badRequest
	}
	if conflict != nil {
		return nil, conflict
	}
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

// Deprecated: Use GrantUserTierShort instead
func (t *TierService) GrantUserTier(input *tier.GrantUserTierParams) (*seasonpassclientmodels.UserSeasonSummary, error) {
	accessToken, err := t.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, err := t.Client.Tier.GrantUserTier(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use GrantUserExpShort instead
func (t *TierService) GrantUserExp(input *tier.GrantUserExpParams) (*seasonpassclientmodels.UserSeasonSummary, error) {
	accessToken, err := t.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, err := t.Client.Tier.GrantUserExp(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (t *TierService) UpdateTierShort(input *tier.UpdateTierParams, authInfo runtime.ClientAuthInfoWriter) (*seasonpassclientmodels.Tier, error) {
	ok, err := t.Client.Tier.UpdateTierShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (t *TierService) DeleteTierShort(input *tier.DeleteTierParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := t.Client.Tier.DeleteTierShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (t *TierService) QueryTiersShort(input *tier.QueryTiersParams, authInfo runtime.ClientAuthInfoWriter) (*seasonpassclientmodels.TierPagingSlicedResult, error) {
	ok, err := t.Client.Tier.QueryTiersShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (t *TierService) CreateTierShort(input *tier.CreateTierParams, authInfo runtime.ClientAuthInfoWriter) ([]*seasonpassclientmodels.Tier, error) {
	created, err := t.Client.Tier.CreateTierShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (t *TierService) GrantUserTierShort(input *tier.GrantUserTierParams, authInfo runtime.ClientAuthInfoWriter) (*seasonpassclientmodels.UserSeasonSummary, error) {
	ok, err := t.Client.Tier.GrantUserTierShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (t *TierService) GrantUserExpShort(input *tier.GrantUserExpParams, authInfo runtime.ClientAuthInfoWriter) (*seasonpassclientmodels.UserSeasonSummary, error) {
	ok, err := t.Client.Tier.GrantUserExpShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}
