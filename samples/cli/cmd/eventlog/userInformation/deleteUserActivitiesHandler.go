// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package userInformation

import (
	"github.com/AccelByte/accelbyte-go-sdk/eventlog-sdk/pkg/eventlogclient/user_information"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/eventlog"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// DeleteUserActivitiesHandlerCmd represents the DeleteUserActivitiesHandler command
var DeleteUserActivitiesHandlerCmd = &cobra.Command{
	Use:   "deleteUserActivitiesHandler",
	Short: "Delete user activities handler",
	Long:  `Delete user activities handler`,
	RunE: func(cmd *cobra.Command, args []string) error {
		userInformationService := &eventlog.UserInformationService{
			Client:          factory.NewEventlogClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		input := &user_information.DeleteUserActivitiesHandlerParams{
			Namespace: namespace,
			UserID:    userId,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		errInput := userInformationService.DeleteUserActivitiesHandler(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	DeleteUserActivitiesHandlerCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = DeleteUserActivitiesHandlerCmd.MarkFlagRequired("namespace")
	DeleteUserActivitiesHandlerCmd.Flags().StringP("userId", "", "", "User id")
	_ = DeleteUserActivitiesHandlerCmd.MarkFlagRequired("userId")
}