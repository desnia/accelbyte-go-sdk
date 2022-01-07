// Copyright (c) 2018 - 2021
// AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package seasonpass

import (
	"github.com/AccelByte/accelbyte-go-sdk/seasonpass-sdk/pkg/seasonpassclient"
	"github.com/AccelByte/accelbyte-go-sdk/seasonpass-sdk/pkg/seasonpassclient/reward"
	"github.com/AccelByte/accelbyte-go-sdk/seasonpass-sdk/pkg/seasonpassclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

type RewardService struct {
	Client          *seasonpassclient.JusticeSeasonpassService
	TokenRepository repository.TokenRepository
}

// Deprecated: Use GetRewardShort instead
func (r *RewardService) GetReward(input *reward.GetRewardParams) (*seasonpassclientmodels.RewardInfo, error) {
	accessToken, err := r.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, badRequest, err := r.Client.Reward.GetReward(input, client.BearerToken(*accessToken.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if badRequest != nil {
		return nil, badRequest
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use DeleteRewardShort instead
func (r *RewardService) DeleteReward(input *reward.DeleteRewardParams) error {
	accessToken, err := r.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, notFound, conflict, badRequest, err := r.Client.Reward.DeleteReward(input, client.BearerToken(*accessToken.AccessToken))
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

// Deprecated: Use UpdateRewardShort instead
func (r *RewardService) UpdateReward(input *reward.UpdateRewardParams) (*seasonpassclientmodels.RewardInfo, error) {
	accessToken, err := r.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, badRequest, conflict, unprocessableEntity, err := r.Client.Reward.UpdateReward(input, client.BearerToken(*accessToken.AccessToken))
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

// Deprecated: Use QueryRewardsShort instead
func (r *RewardService) QueryRewards(input *reward.QueryRewardsParams) ([]*seasonpassclientmodels.RewardInfo, error) {
	accessToken, err := r.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, err := r.Client.Reward.QueryRewards(input, client.BearerToken(*accessToken.AccessToken))
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

// Deprecated: Use CreateRewardShort instead
func (r *RewardService) CreateReward(input *reward.CreateRewardParams) (*seasonpassclientmodels.RewardInfo, error) {
	accessToken, err := r.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, notFound, badRequest, conflict, unprocessableEntity, err := r.Client.Reward.CreateReward(input, client.BearerToken(*accessToken.AccessToken))
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

// Deprecated: Use PublicClaimUserRewardShort instead
func (r *RewardService) PublicClaimUserReward(input *reward.PublicClaimUserRewardParams) (*seasonpassclientmodels.ClaimableRewards, error) {
	accessToken, err := r.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, badRequest, conflict, err := r.Client.Reward.PublicClaimUserReward(input, client.BearerToken(*accessToken.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if badRequest != nil {
		return nil, badRequest
	}
	if conflict != nil {
		return nil, conflict
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use PublicBulkClaimUserRewardsShort instead
func (r *RewardService) PublicBulkClaimUserRewards(input *reward.PublicBulkClaimUserRewardsParams) (*seasonpassclientmodels.ClaimableRewards, error) {
	accessToken, err := r.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, badRequest, err := r.Client.Reward.PublicBulkClaimUserRewards(input, client.BearerToken(*accessToken.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if badRequest != nil {
		return nil, badRequest
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (r *RewardService) GetRewardShort(input *reward.GetRewardParams, authInfo runtime.ClientAuthInfoWriter) (*seasonpassclientmodels.RewardInfo, error) {
	ok, err := r.Client.Reward.GetRewardShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (r *RewardService) DeleteRewardShort(input *reward.DeleteRewardParams, authInfo runtime.ClientAuthInfoWriter) error {
	_, err := r.Client.Reward.DeleteRewardShort(input, authInfo)
	if err != nil {
		return err
	}
	return nil
}

func (r *RewardService) UpdateRewardShort(input *reward.UpdateRewardParams, authInfo runtime.ClientAuthInfoWriter) (*seasonpassclientmodels.RewardInfo, error) {
	ok, err := r.Client.Reward.UpdateRewardShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (r *RewardService) QueryRewardsShort(input *reward.QueryRewardsParams, authInfo runtime.ClientAuthInfoWriter) ([]*seasonpassclientmodels.RewardInfo, error) {
	ok, err := r.Client.Reward.QueryRewardsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (r *RewardService) CreateRewardShort(input *reward.CreateRewardParams, authInfo runtime.ClientAuthInfoWriter) (*seasonpassclientmodels.RewardInfo, error) {
	created, err := r.Client.Reward.CreateRewardShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (r *RewardService) PublicClaimUserRewardShort(input *reward.PublicClaimUserRewardParams, authInfo runtime.ClientAuthInfoWriter) (*seasonpassclientmodels.ClaimableRewards, error) {
	ok, err := r.Client.Reward.PublicClaimUserRewardShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (r *RewardService) PublicBulkClaimUserRewardsShort(input *reward.PublicBulkClaimUserRewardsParams, authInfo runtime.ClientAuthInfoWriter) (*seasonpassclientmodels.ClaimableRewards, error) {
	ok, err := r.Client.Reward.PublicBulkClaimUserRewardsShort(input, authInfo)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}