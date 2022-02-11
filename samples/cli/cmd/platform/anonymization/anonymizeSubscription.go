// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package anonymization

import (
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/anonymization"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// AnonymizeSubscriptionCmd represents the AnonymizeSubscription command
var AnonymizeSubscriptionCmd = &cobra.Command{
	Use:   "anonymizeSubscription",
	Short: "Anonymize subscription",
	Long:  `Anonymize subscription`,
	RunE: func(cmd *cobra.Command, args []string) error {
		anonymizationService := &platform.AnonymizationService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		input := &anonymization.AnonymizeSubscriptionParams{
			Namespace: namespace,
			UserID:    userId,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		errInput := anonymizationService.AnonymizeSubscription(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	AnonymizeSubscriptionCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = AnonymizeSubscriptionCmd.MarkFlagRequired("namespace")
	AnonymizeSubscriptionCmd.Flags().StringP("userId", "", "", "User id")
	_ = AnonymizeSubscriptionCmd.MarkFlagRequired("userId")
}