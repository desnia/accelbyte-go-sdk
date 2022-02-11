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

// AdminUpdateAgeRestrictionConfigV3Cmd represents the AdminUpdateAgeRestrictionConfigV3 command
var AdminUpdateAgeRestrictionConfigV3Cmd = &cobra.Command{
	Use:   "adminUpdateAgeRestrictionConfigV3",
	Short: "Admin update age restriction config V3",
	Long:  `Admin update age restriction config V3`,
	RunE: func(cmd *cobra.Command, args []string) error {
		usersService := &iam.UsersService{
			Client:          factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		bodyString := cmd.Flag("body").Value.String()
		var body *iamclientmodels.ModelAgeRestrictionRequestV3
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		input := &users.AdminUpdateAgeRestrictionConfigV3Params{
			Body:      body,
			Namespace: namespace,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := usersService.AdminUpdateAgeRestrictionConfigV3(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	AdminUpdateAgeRestrictionConfigV3Cmd.Flags().StringP("body", "", "", "Body")
	_ = AdminUpdateAgeRestrictionConfigV3Cmd.MarkFlagRequired("body")
	AdminUpdateAgeRestrictionConfigV3Cmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = AdminUpdateAgeRestrictionConfigV3Cmd.MarkFlagRequired("namespace")
}