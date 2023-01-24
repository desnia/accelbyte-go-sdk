// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package oAuth20Extension

import (
	"net/http"

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient/o_auth2_0_extension"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/iam"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// RequestTokenByOneTimeLinkCodeResponseV3Cmd represents the RequestTokenByOneTimeLinkCodeResponseV3 command
var RequestTokenByOneTimeLinkCodeResponseV3Cmd = &cobra.Command{
	Use:   "requestTokenByOneTimeLinkCodeResponseV3",
	Short: "Request token by one time link code response V3",
	Long:  `Request token by one time link code response V3`,
	RunE: func(cmd *cobra.Command, args []string) error {
		oAuth20ExtensionService := &iam.OAuth20ExtensionService{
			Client:          factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		clientId, _ := cmd.Flags().GetString("clientId")
		oneTimeLinkCode, _ := cmd.Flags().GetString("oneTimeLinkCode")
		httpClient := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		input := &o_auth2_0_extension.RequestTokenByOneTimeLinkCodeResponseV3Params{
			ClientID:        clientId,
			OneTimeLinkCode: oneTimeLinkCode,
			HTTPClient:      httpClient,
		}
		ok, err := oAuth20ExtensionService.RequestTokenByOneTimeLinkCodeResponseV3Short(input)
		if err != nil {
			logrus.Error(err)

			return err
		} else {
			logrus.Infof("Response CLI success: %+v", ok)
		}

		return nil
	},
}

func init() {
	RequestTokenByOneTimeLinkCodeResponseV3Cmd.Flags().String("clientId", "", "Client id")
	_ = RequestTokenByOneTimeLinkCodeResponseV3Cmd.MarkFlagRequired("client_id")
	RequestTokenByOneTimeLinkCodeResponseV3Cmd.Flags().String("oneTimeLinkCode", "", "One time link code")
	_ = RequestTokenByOneTimeLinkCodeResponseV3Cmd.MarkFlagRequired("oneTimeLinkCode")
}