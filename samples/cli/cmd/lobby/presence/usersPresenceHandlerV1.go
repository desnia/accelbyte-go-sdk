// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package presence

import (
	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclient/presence"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/lobby"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// UsersPresenceHandlerV1Cmd represents the UsersPresenceHandlerV1 command
var UsersPresenceHandlerV1Cmd = &cobra.Command{
	Use:   "usersPresenceHandlerV1",
	Short: "Users presence handler V1",
	Long:  `Users presence handler V1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		presenceService := &lobby.PresenceService{
			Client:          factory.NewLobbyClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userIds, _ := cmd.Flags().GetString("userIds")
		countOnly, _ := cmd.Flags().GetBool("countOnly")
		input := &presence.UsersPresenceHandlerV1Params{
			Namespace: namespace,
			CountOnly: &countOnly,
			UserIds:   userIds,
		}
		ok, err := presenceService.UsersPresenceHandlerV1Short(input, nil)
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
	UsersPresenceHandlerV1Cmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = UsersPresenceHandlerV1Cmd.MarkFlagRequired("namespace")
	UsersPresenceHandlerV1Cmd.Flags().BoolP("countOnly", "", false, "Count only")
	UsersPresenceHandlerV1Cmd.Flags().StringP("userIds", "", "", "User ids")
	_ = UsersPresenceHandlerV1Cmd.MarkFlagRequired("userIds")
}
