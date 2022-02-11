// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package session

import (
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/sessionbrowser"
	"github.com/AccelByte/accelbyte-go-sdk/sessionbrowser-sdk/pkg/sessionbrowserclient/session"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// GetRecentPlayerCmd represents the GetRecentPlayer command
var GetRecentPlayerCmd = &cobra.Command{
	Use:   "getRecentPlayer",
	Short: "Get recent player",
	Long:  `Get recent player`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sessionService := &sessionbrowser.SessionService{
			Client:          factory.NewSessionbrowserClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userID, _ := cmd.Flags().GetString("userID")
		input := &session.GetRecentPlayerParams{
			Namespace: namespace,
			UserID:    userID,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := sessionService.GetRecentPlayer(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	GetRecentPlayerCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = GetRecentPlayerCmd.MarkFlagRequired("namespace")
	GetRecentPlayerCmd.Flags().StringP("userID", "", "", "User ID")
	_ = GetRecentPlayerCmd.MarkFlagRequired("userID")
}