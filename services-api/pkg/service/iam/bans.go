// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package iam

import (
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient"
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient/bans"
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils/auth"
	"github.com/go-openapi/runtime/client"
)

type BansService struct {
	Client          *iamclient.JusticeIamService
	TokenRepository repository.TokenRepository
}

// Deprecated: Use GetBansTypeShort instead
func (b *BansService) GetBansType(input *bans.GetBansTypeParams) (*iamclientmodels.AccountcommonBans, error) {
	token, err := b.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, err := b.Client.Bans.GetBansType(input, client.BearerToken(*token.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use GetListBanReasonShort instead
func (b *BansService) GetListBanReason(input *bans.GetListBanReasonParams) (*iamclientmodels.AccountcommonBanReasons, error) {
	token, err := b.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, err := b.Client.Bans.GetListBanReason(input, client.BearerToken(*token.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use AdminGetBansTypeV3Short instead
func (b *BansService) AdminGetBansTypeV3(input *bans.AdminGetBansTypeV3Params) (*iamclientmodels.AccountcommonBansV3, error) {
	token, err := b.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, err := b.Client.Bans.AdminGetBansTypeV3(input, client.BearerToken(*token.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use AdminGetListBanReasonV3Short instead
func (b *BansService) AdminGetListBanReasonV3(input *bans.AdminGetListBanReasonV3Params) (*iamclientmodels.AccountcommonBanReasonsV3, error) {
	token, err := b.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, err := b.Client.Bans.AdminGetListBanReasonV3(input, client.BearerToken(*token.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use AdminGetBannedUsersV3Short instead
func (b *BansService) AdminGetBannedUsersV3(input *bans.AdminGetBannedUsersV3Params) (*iamclientmodels.ModelGetUserBanV3Response, error) {
	token, err := b.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, err := b.Client.Bans.AdminGetBannedUsersV3(input, client.BearerToken(*token.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

// Deprecated: Use AdminGetBansTypeWithNamespaceV3Short instead
func (b *BansService) AdminGetBansTypeWithNamespaceV3(input *bans.AdminGetBansTypeWithNamespaceV3Params) (*iamclientmodels.AccountcommonBansV3, error) {
	token, err := b.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, unauthorized, forbidden, err := b.Client.Bans.AdminGetBansTypeWithNamespaceV3(input, client.BearerToken(*token.AccessToken))
	if unauthorized != nil {
		return nil, unauthorized
	}
	if forbidden != nil {
		return nil, forbidden
	}
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

func (b *BansService) GetBansTypeShort(input *bans.GetBansTypeParams) (*iamclientmodels.AccountcommonBans, error) {
	token, err := b.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	authWriter := auth.Compose(
		auth.Bearer(*token.AccessToken),
	)
	ok, err := b.Client.Bans.GetBansTypeShort(input, authWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

func (b *BansService) GetListBanReasonShort(input *bans.GetListBanReasonParams) (*iamclientmodels.AccountcommonBanReasons, error) {
	token, err := b.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	authWriter := auth.Compose(
		auth.Bearer(*token.AccessToken),
	)
	ok, err := b.Client.Bans.GetListBanReasonShort(input, authWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

func (b *BansService) AdminGetBansTypeV3Short(input *bans.AdminGetBansTypeV3Params) (*iamclientmodels.AccountcommonBansV3, error) {
	token, err := b.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	authWriter := auth.Compose(
		auth.Bearer(*token.AccessToken),
	)
	ok, err := b.Client.Bans.AdminGetBansTypeV3Short(input, authWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

func (b *BansService) AdminGetListBanReasonV3Short(input *bans.AdminGetListBanReasonV3Params) (*iamclientmodels.AccountcommonBanReasonsV3, error) {
	token, err := b.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	authWriter := auth.Compose(
		auth.Bearer(*token.AccessToken),
	)
	ok, err := b.Client.Bans.AdminGetListBanReasonV3Short(input, authWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

func (b *BansService) AdminGetBannedUsersV3Short(input *bans.AdminGetBannedUsersV3Params) (*iamclientmodels.ModelGetUserBanV3Response, error) {
	token, err := b.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	authWriter := auth.Compose(
		auth.Bearer(*token.AccessToken),
	)
	ok, err := b.Client.Bans.AdminGetBannedUsersV3Short(input, authWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}

func (b *BansService) AdminGetBansTypeWithNamespaceV3Short(input *bans.AdminGetBansTypeWithNamespaceV3Params) (*iamclientmodels.AccountcommonBansV3, error) {
	token, err := b.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	authWriter := auth.Compose(
		auth.Bearer(*token.AccessToken),
	)
	ok, err := b.Client.Bans.AdminGetBansTypeWithNamespaceV3Short(input, authWriter)
	if err != nil {
		return nil, err
	}

	return ok.GetPayload(), nil
}
