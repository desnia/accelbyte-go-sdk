// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package thirdParty

import (
	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclient/third_party"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/lobby"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// AdminGetThirdPartyConfigCmd represents the AdminGetThirdPartyConfig command
var AdminGetThirdPartyConfigCmd = &cobra.Command{
	Use:   "adminGetThirdPartyConfig",
	Short: "Admin get third party config",
	Long:  `Admin get third party config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		thirdPartyService := &lobby.ThirdPartyService{
			Client:          factory.NewLobbyClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		input := &third_party.AdminGetThirdPartyConfigParams{
			Namespace: namespace,
		}
		ok, err := thirdPartyService.AdminGetThirdPartyConfigShort(input, nil)
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
	AdminGetThirdPartyConfigCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = AdminGetThirdPartyConfigCmd.MarkFlagRequired("namespace")
}
