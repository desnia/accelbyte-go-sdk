// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package users

import (
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient/users"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/iam"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// AdminGetListCountryAgeRestrictionV3Cmd represents the AdminGetListCountryAgeRestrictionV3 command
var AdminGetListCountryAgeRestrictionV3Cmd = &cobra.Command{
	Use:   "adminGetListCountryAgeRestrictionV3",
	Short: "Admin get list country age restriction V3",
	Long:  `Admin get list country age restriction V3`,
	RunE: func(cmd *cobra.Command, args []string) error {
		usersService := &iam.UsersService{
			Client:          factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		input := &users.AdminGetListCountryAgeRestrictionV3Params{
			Namespace: namespace,
		}
		ok, err := usersService.AdminGetListCountryAgeRestrictionV3Short(input, nil)
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
	AdminGetListCountryAgeRestrictionV3Cmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = AdminGetListCountryAgeRestrictionV3Cmd.MarkFlagRequired("namespace")
}
