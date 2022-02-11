// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package userStatistic

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/social"
	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclient/user_statistic"
	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// PublicBulkIncUserStatItem1Cmd represents the PublicBulkIncUserStatItem1 command
var PublicBulkIncUserStatItem1Cmd = &cobra.Command{
	Use:   "publicBulkIncUserStatItem1",
	Short: "Public bulk inc user stat item 1",
	Long:  `Public bulk inc user stat item 1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		userStatisticService := &social.UserStatisticService{
			Client:          factory.NewSocialClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		bodyString := cmd.Flag("body").Value.String()
		var body []*socialclientmodels.BulkStatItemInc
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		input := &user_statistic.PublicBulkIncUserStatItem1Params{
			Body:      body,
			Namespace: namespace,
			UserID:    userId,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := userStatisticService.PublicBulkIncUserStatItem1(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	PublicBulkIncUserStatItem1Cmd.Flags().StringP("body", "", "", "Body")
	PublicBulkIncUserStatItem1Cmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = PublicBulkIncUserStatItem1Cmd.MarkFlagRequired("namespace")
	PublicBulkIncUserStatItem1Cmd.Flags().StringP("userId", "", "", "User id")
	_ = PublicBulkIncUserStatItem1Cmd.MarkFlagRequired("userId")
}