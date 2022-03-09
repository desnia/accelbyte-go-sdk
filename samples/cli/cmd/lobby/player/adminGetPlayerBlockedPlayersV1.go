// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package player

import (
	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclient/player"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/lobby"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// AdminGetPlayerBlockedPlayersV1Cmd represents the AdminGetPlayerBlockedPlayersV1 command
var AdminGetPlayerBlockedPlayersV1Cmd = &cobra.Command{
	Use:   "adminGetPlayerBlockedPlayersV1",
	Short: "Admin get player blocked players V1",
	Long:  `Admin get player blocked players V1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		playerService := &lobby.PlayerService{
			Client:          factory.NewLobbyClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		input := &player.AdminGetPlayerBlockedPlayersV1Params{
			Namespace: namespace,
			UserID:    userId,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := playerService.AdminGetPlayerBlockedPlayersV1(input)
		if err != nil {
			logrus.Error(err)
			return err
		} else {
			logrus.Infof("Response CLI success", ok)
		}
		return nil
	},
}

func init() {
	AdminGetPlayerBlockedPlayersV1Cmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = AdminGetPlayerBlockedPlayersV1Cmd.MarkFlagRequired("namespace")
	AdminGetPlayerBlockedPlayersV1Cmd.Flags().StringP("userId", "", "", "User id")
	_ = AdminGetPlayerBlockedPlayersV1Cmd.MarkFlagRequired("userId")
}
