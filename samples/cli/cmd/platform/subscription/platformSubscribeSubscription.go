// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package subscription

import (
	"encoding/json"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/subscription"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// PlatformSubscribeSubscriptionCmd represents the PlatformSubscribeSubscription command
var PlatformSubscribeSubscriptionCmd = &cobra.Command{
	Use:   "platformSubscribeSubscription",
	Short: "Platform subscribe subscription",
	Long:  `Platform subscribe subscription`,
	RunE: func(cmd *cobra.Command, args []string) error {
		subscriptionService := &platform.SubscriptionService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		bodyString := cmd.Flag("body").Value.String()
		var body *platformclientmodels.PlatformSubscribeRequest
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		input := &subscription.PlatformSubscribeSubscriptionParams{
			Body:      body,
			Namespace: namespace,
			UserID:    userId,
		}
		ok, err := subscriptionService.PlatformSubscribeSubscriptionShort(input, nil)
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
	PlatformSubscribeSubscriptionCmd.Flags().StringP("body", "", "", "Body")
	PlatformSubscribeSubscriptionCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = PlatformSubscribeSubscriptionCmd.MarkFlagRequired("namespace")
	PlatformSubscribeSubscriptionCmd.Flags().StringP("userId", "", "", "User id")
	_ = PlatformSubscribeSubscriptionCmd.MarkFlagRequired("userId")
}
