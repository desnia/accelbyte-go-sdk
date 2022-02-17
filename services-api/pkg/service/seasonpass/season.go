// Copyright (c) 2018 - 2021
// AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package seasonpass

import (
	"github.com/AccelByte/accelbyte-go-sdk/seasonpass-sdk/pkg/seasonpassclient"
	"github.com/AccelByte/accelbyte-go-sdk/seasonpass-sdk/pkg/seasonpassclient/season"
	"github.com/AccelByte/accelbyte-go-sdk/seasonpass-sdk/pkg/seasonpassclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/go-openapi/runtime/client"
)

type SeasonService struct {
	Client          *seasonpassclient.JusticeSeasonpassService
	TokenRepository repository.TokenRepository
}

// Deprecated: Use QuerySeasonsShort instead
func (s *SeasonService) QuerySeasons(input *season.QuerySeasonsParams) (*seasonpassclientmodels.ListSeasonInfoPagingSlicedResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, err := s.Client.Season.QuerySeasons(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use CreateSeasonShort instead
func (s *SeasonService) CreateSeason(input *season.CreateSeasonParams) (*seasonpassclientmodels.SeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, badRequest, notFound, unprocessableEntity, err := s.Client.Season.CreateSeason(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
	}
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

// Deprecated: Use GetCurrentSeasonShort instead
func (s *SeasonService) GetCurrentSeason(input *season.GetCurrentSeasonParams) (*seasonpassclientmodels.SeasonSummary, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, err := s.Client.Season.GetCurrentSeason(input, client.BearerToken(*accessToken.AccessToken))
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

// Deprecated: Use GetSeasonShort instead
func (s *SeasonService) GetSeason(input *season.GetSeasonParams) (*seasonpassclientmodels.SeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, err := s.Client.Season.GetSeason(input, client.BearerToken(*accessToken.AccessToken))
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

// Deprecated: Use DeleteSeasonShort instead
func (s *SeasonService) DeleteSeason(input *season.DeleteSeasonParams) error {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, notFound, conflict, err := s.Client.Season.DeleteSeason(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if notFound != nil {
		return notFound
	}
	if conflict != nil {
		return conflict
	}
	if err != nil {
		return err
	}
	return nil
}

// Deprecated: Use UpdateSeasonShort instead
func (s *SeasonService) UpdateSeason(input *season.UpdateSeasonParams) (*seasonpassclientmodels.SeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, conflict, unprocessableEntity, err := s.Client.Season.UpdateSeason(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
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

// Deprecated: Use CloneSeasonShort instead
func (s *SeasonService) CloneSeason(input *season.CloneSeasonParams) (*seasonpassclientmodels.SeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, unprocessableEntity, err := s.Client.Season.CloneSeason(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
	}
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use PublishSeasonShort instead
func (s *SeasonService) PublishSeason(input *season.PublishSeasonParams) (*seasonpassclientmodels.SeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, conflict, err := s.Client.Season.PublishSeason(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
	}
	if conflict != nil {
		return nil, conflict
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use RetireSeasonShort instead
func (s *SeasonService) RetireSeason(input *season.RetireSeasonParams) (*seasonpassclientmodels.SeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, conflict, err := s.Client.Season.RetireSeason(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
	}
	if conflict != nil {
		return nil, conflict
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use UnpublishSeasonShort instead
func (s *SeasonService) UnpublishSeason(input *season.UnpublishSeasonParams) (*seasonpassclientmodels.SeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, conflict, err := s.Client.Season.UnpublishSeason(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
	}
	if conflict != nil {
		return nil, conflict
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use GetUserParticipatedSeasonsShort instead
func (s *SeasonService) GetUserParticipatedSeasons(input *season.GetUserParticipatedSeasonsParams) (*seasonpassclientmodels.ListUserSeasonInfoPagingSlicedResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, err := s.Client.Season.GetUserParticipatedSeasons(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use ExistsAnyPassByPassCodesShort instead
func (s *SeasonService) ExistsAnyPassByPassCodes(input *season.ExistsAnyPassByPassCodesParams) (*seasonpassclientmodels.Ownership, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, err := s.Client.Season.ExistsAnyPassByPassCodes(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use GetCurrentUserSeasonProgressionShort instead
func (s *SeasonService) GetCurrentUserSeasonProgression(input *season.GetCurrentUserSeasonProgressionParams) (*seasonpassclientmodels.UserSeasonSummary, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, err := s.Client.Season.GetCurrentUserSeasonProgression(input, client.BearerToken(*accessToken.AccessToken))
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

// Deprecated: Use CheckSeasonPurchasableShort instead
func (s *SeasonService) CheckSeasonPurchasable(input *season.CheckSeasonPurchasableParams) error {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, notFound, conflict, err := s.Client.Season.CheckSeasonPurchasable(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if notFound != nil {
		return notFound
	}
	if conflict != nil {
		return conflict
	}
	if err != nil {
		return err
	}
	return nil
}

// Deprecated: Use ResetUserSeasonShort instead
func (s *SeasonService) ResetUserSeason(input *season.ResetUserSeasonParams) error {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, err := s.Client.Season.ResetUserSeason(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if err != nil {
		return err
	}
	return nil
}

// Deprecated: Use GetUserSeasonShort instead
func (s *SeasonService) GetUserSeason(input *season.GetUserSeasonParams) (*seasonpassclientmodels.ClaimableUserSeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, err := s.Client.Season.GetUserSeason(input, client.BearerToken(*accessToken.AccessToken))
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

// Deprecated: Use PublicGetCurrentSeasonShort instead
func (s *SeasonService) PublicGetCurrentSeason(input *season.PublicGetCurrentSeasonParams) (*seasonpassclientmodels.LocalizedSeasonInfo, error) {
	ok, badRequest, notFound, err := s.Client.Season.PublicGetCurrentSeason(input)
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

// Deprecated: Use PublicGetCurrentUserSeasonShort instead
func (s *SeasonService) PublicGetCurrentUserSeason(input *season.PublicGetCurrentUserSeasonParams) (*seasonpassclientmodels.ClaimableUserSeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, err := s.Client.Season.PublicGetCurrentUserSeason(input, client.BearerToken(*accessToken.AccessToken))
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

// Deprecated: Use PublicGetUserSeasonShort instead
func (s *SeasonService) PublicGetUserSeason(input *season.PublicGetUserSeasonParams) (*seasonpassclientmodels.ClaimableUserSeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, err := s.Client.Season.PublicGetUserSeason(input, client.BearerToken(*accessToken.AccessToken))
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

func (s *SeasonService) QuerySeasonsShort(input *season.QuerySeasonsParams) (*seasonpassclientmodels.ListSeasonInfoPagingSlicedResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Season.QuerySeasonsShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SeasonService) CreateSeasonShort(input *season.CreateSeasonParams) (*seasonpassclientmodels.SeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	created, err := s.Client.Season.CreateSeasonShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return created.GetPayload(), nil
}

func (s *SeasonService) GetCurrentSeasonShort(input *season.GetCurrentSeasonParams) (*seasonpassclientmodels.SeasonSummary, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Season.GetCurrentSeasonShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SeasonService) GetSeasonShort(input *season.GetSeasonParams) (*seasonpassclientmodels.SeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Season.GetSeasonShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SeasonService) DeleteSeasonShort(input *season.DeleteSeasonParams) error {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, err = s.Client.Season.DeleteSeasonShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return err
	}
	return nil
}

func (s *SeasonService) UpdateSeasonShort(input *season.UpdateSeasonParams) (*seasonpassclientmodels.SeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Season.UpdateSeasonShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SeasonService) CloneSeasonShort(input *season.CloneSeasonParams) (*seasonpassclientmodels.SeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Season.CloneSeasonShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SeasonService) PublishSeasonShort(input *season.PublishSeasonParams) (*seasonpassclientmodels.SeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Season.PublishSeasonShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SeasonService) RetireSeasonShort(input *season.RetireSeasonParams) (*seasonpassclientmodels.SeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Season.RetireSeasonShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SeasonService) UnpublishSeasonShort(input *season.UnpublishSeasonParams) (*seasonpassclientmodels.SeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Season.UnpublishSeasonShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SeasonService) GetUserParticipatedSeasonsShort(input *season.GetUserParticipatedSeasonsParams) (*seasonpassclientmodels.ListUserSeasonInfoPagingSlicedResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Season.GetUserParticipatedSeasonsShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SeasonService) ExistsAnyPassByPassCodesShort(input *season.ExistsAnyPassByPassCodesParams) (*seasonpassclientmodels.Ownership, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Season.ExistsAnyPassByPassCodesShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SeasonService) GetCurrentUserSeasonProgressionShort(input *season.GetCurrentUserSeasonProgressionParams) (*seasonpassclientmodels.UserSeasonSummary, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Season.GetCurrentUserSeasonProgressionShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SeasonService) CheckSeasonPurchasableShort(input *season.CheckSeasonPurchasableParams) error {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, err = s.Client.Season.CheckSeasonPurchasableShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return err
	}
	return nil
}

func (s *SeasonService) ResetUserSeasonShort(input *season.ResetUserSeasonParams) error {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, err = s.Client.Season.ResetUserSeasonShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return err
	}
	return nil
}

func (s *SeasonService) GetUserSeasonShort(input *season.GetUserSeasonParams) (*seasonpassclientmodels.ClaimableUserSeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Season.GetUserSeasonShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SeasonService) PublicGetCurrentSeasonShort(input *season.PublicGetCurrentSeasonParams) (*seasonpassclientmodels.LocalizedSeasonInfo, error) {
	ok, err := s.Client.Season.PublicGetCurrentSeasonShort(input)
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SeasonService) PublicGetCurrentUserSeasonShort(input *season.PublicGetCurrentUserSeasonParams) (*seasonpassclientmodels.ClaimableUserSeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Season.PublicGetCurrentUserSeasonShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SeasonService) PublicGetUserSeasonShort(input *season.PublicGetUserSeasonParams) (*seasonpassclientmodels.ClaimableUserSeasonInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Season.PublicGetUserSeasonShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}
