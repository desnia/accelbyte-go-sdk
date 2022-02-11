// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package users

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient/users"
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/iam"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// AdminSaveUserRoleV3Cmd represents the AdminSaveUserRoleV3 command
var AdminSaveUserRoleV3Cmd = &cobra.Command{
	Use:   "adminSaveUserRoleV3",
	Short: "Admin save user role V3",
	Long:  `Admin save user role V3`,
	RunE: func(cmd *cobra.Command, args []string) error {
		usersService := &iam.UsersService{
			Client:          factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		bodyString := cmd.Flag("body").Value.String()
		var body []*iamclientmodels.ModelNamespaceRoleRequest
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		input := &users.AdminSaveUserRoleV3Params{
			Body:      body,
			Namespace: namespace,
			UserID:    userId,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		errInput := usersService.AdminSaveUserRoleV3(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	AdminSaveUserRoleV3Cmd.Flags().StringP("body", "", "", "Body")
	_ = AdminSaveUserRoleV3Cmd.MarkFlagRequired("body")
	AdminSaveUserRoleV3Cmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = AdminSaveUserRoleV3Cmd.MarkFlagRequired("namespace")
	AdminSaveUserRoleV3Cmd.Flags().StringP("userId", "", "", "User id")
	_ = AdminSaveUserRoleV3Cmd.MarkFlagRequired("userId")
}