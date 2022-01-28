// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package matchmaking

import (
	matchmaking_ "github.com/AccelByte/accelbyte-go-sdk/matchmaking-sdk/pkg/matchmakingclient/matchmaking"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/matchmaking"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// PublicGetSingleMatchmakingChannelCmd represents the PublicGetSingleMatchmakingChannel command
var PublicGetSingleMatchmakingChannelCmd = &cobra.Command{
	Use:   "publicGetSingleMatchmakingChannel",
	Short: "Public get single matchmaking channel",
	Long:  `Public get single matchmaking channel`,
	RunE: func(cmd *cobra.Command, args []string) error {
		matchmakingService := &matchmaking.MatchmakingService{
			Client:          factory.NewMatchmakingClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		channelName, _ := cmd.Flags().GetString("channelName")
		namespace, _ := cmd.Flags().GetString("namespace")
		input := &matchmaking_.PublicGetSingleMatchmakingChannelParams{
			ChannelName: channelName,
			Namespace:   namespace,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := matchmakingService.PublicGetSingleMatchmakingChannel(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	PublicGetSingleMatchmakingChannelCmd.Flags().StringP("channelName", "", " ", "Channel name")
	_ = PublicGetSingleMatchmakingChannelCmd.MarkFlagRequired("channelName")
	PublicGetSingleMatchmakingChannelCmd.Flags().StringP("namespace", "", " ", "Namespace")
	_ = PublicGetSingleMatchmakingChannelCmd.MarkFlagRequired("namespace")
}
