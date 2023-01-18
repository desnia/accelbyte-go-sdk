// Copyright (c) 2022 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package auth_test

import (
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/assert"

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils/auth"
)

const dummyAccessToken = "foo"

func TestAuthInfoWriterRefresh(t *testing.T) {
	r := &runtime.TestClientRequest{}

	// mockup client input value
	refreshToken := dummyAccessToken
	token = &iamclientmodels.OauthmodelTokenResponseV3{RefreshToken: refreshToken}
	err := dummyService.TokenRepository.Store(*token)
	if err != nil {
		assert.FailNow(t, "fail to store the token")
	}
	tokenStored, errGetToken := dummyService.TokenRepository.GetToken()
	if errGetToken != nil {
		assert.FailNow(t, "fail to store the token")
	}

	// call the AuthInfoWriter
	if authInfoWriter == nil {
		security = [][]string{
			{"refresh"},
		}
		authInfoWriter = auth.AuthInfoWriter(auth.Session{
			Token:   dummyService.TokenRepository,
			Config:  dummyService.ConfigRepository,
			Refresh: auth.DefaultRefreshTokenImpl(),
		}, security, "")
	}

	writer := auth.AuthInfoWriter(auth.Session{
		Token:   dummyService.TokenRepository,
		Config:  dummyService.ConfigRepository,
		Refresh: auth.DefaultRefreshTokenImpl(),
	}, security, "")
	err = writer.AuthenticateRequest(r, nil)
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, tokenStored.AccessToken, token.AccessToken)
}

func TestAuthInfoWriter_RefreshToken(t *testing.T) {
	refreshToken := dummyAccessToken
	tests := []struct {
		name string
		args iamclientmodels.OauthmodelTokenResponseV3
	}{
		{
			name: "refresh token",
			args: iamclientmodels.OauthmodelTokenResponseV3{RefreshToken: refreshToken},
		},
	}

	r := &runtime.TestClientRequest{}

	for _, tt := range tests {
		t.Parallel()
		t.Run(tt.name, func(t *testing.T) {
			err := dummyService.TokenRepository.Store(tt.args)
			if err != nil {
				assert.FailNow(t, "fail to store the token")
			}
			getRefreshToken, errGetToken := repository.GetRefreshToken(dummyService.TokenRepository)
			if errGetToken != nil {
				assert.FailNow(t, "fail to get the token")
			}

			if authInfoWriter == nil {
				security = [][]string{
					{"refresh"},
				}
				authInfoWriter = auth.AuthInfoWriter(auth.Session{
					Token:   dummyService.TokenRepository,
					Config:  dummyService.ConfigRepository,
					Refresh: auth.DefaultRefreshTokenImpl(),
				}, security, "")
			}

			writer := auth.AuthInfoWriter(auth.Session{
				Token:   dummyService.TokenRepository,
				Config:  dummyService.ConfigRepository,
				Refresh: auth.DefaultRefreshTokenImpl(),
			}, security, "")
			errAuth := writer.AuthenticateRequest(r, nil)
			if errAuth != nil {
				assert.FailNow(t, "errAuth should be nil")
			}

			assert.Nil(t, err, "err should be nil")
			assert.Equal(t, getRefreshToken, refreshToken)
		})
	}
}

func TestDefaultRefreshTokenImpl(t *testing.T) {
	refreshToken := auth.DefaultRefreshTokenImpl()

	refreshToken.SetRefreshIsRunningInBackground(true)

	assert.True(t, refreshToken.DisableAutoRefresh())
	assert.Equal(t, 1.0, refreshToken.GetRefreshRate())
	assert.True(t, refreshToken.IsRefreshInProgress)
}
