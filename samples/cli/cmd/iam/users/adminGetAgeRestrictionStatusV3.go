// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package users

import (
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient/users"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/iam"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// AdminGetAgeRestrictionStatusV3Cmd represents the AdminGetAgeRestrictionStatusV3 command
var AdminGetAgeRestrictionStatusV3Cmd = &cobra.Command{
	Use:   "adminGetAgeRestrictionStatusV3",
	Short: "Admin get age restriction status V3",
	Long:  `Admin get age restriction status V3`,
	RunE: func(cmd *cobra.Command, args []string) error {
		usersService := &iam.UsersService{
			Client:          factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		input := &users.AdminGetAgeRestrictionStatusV3Params{
			Namespace: namespace,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := usersService.AdminGetAgeRestrictionStatusV3(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	AdminGetAgeRestrictionStatusV3Cmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = AdminGetAgeRestrictionStatusV3Cmd.MarkFlagRequired("namespace")
}