// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package cmd

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/entitlement"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// existsAnyUserActiveEntitlementByItemIdsCmd represents the existsAnyUserActiveEntitlementByItemIds command
var existsAnyUserActiveEntitlementByItemIdsCmd = &cobra.Command{
	Use:   "existsAnyUserActiveEntitlementByItemIds",
	Short: "Exists any user active entitlement by item ids",
	Long:  `Exists any user active entitlement by item ids`,
	RunE: func(cmd *cobra.Command, args []string) error {
		entitlementService := &platform.EntitlementService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		itemIdsString := cmd.Flag("itemIds").Value.String()
		var itemIds []string
		errItemIds := json.Unmarshal([]byte(itemIdsString), &itemIds)
		if errItemIds != nil {
			return errItemIds
		}
		input := &entitlement.ExistsAnyUserActiveEntitlementByItemIdsParams{
			Namespace: namespace,
			UserID:    userId,
			ItemIds:   itemIds,
		}
		//nolint:staticcheck // SA1019 To be deprecated later
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := entitlementService.ExistsAnyUserActiveEntitlementByItemIds(input)
		if err != nil {
			logrus.Error(err)
			return err
		}
		okJson, err := json.MarshalIndent(ok, "", "   ")
		if err != nil {
			return err
		}
		logrus.Infof("Response: %s", string(okJson))
		return nil
	},
}

func init() {
	RootCmd.AddCommand(existsAnyUserActiveEntitlementByItemIdsCmd)
	existsAnyUserActiveEntitlementByItemIdsCmd.Flags().StringP("namespace", "n", " ", "Namespace")
	_ = existsAnyUserActiveEntitlementByItemIdsCmd.MarkFlagRequired("namespace")
	existsAnyUserActiveEntitlementByItemIdsCmd.Flags().StringP("userId", "u", " ", "User id")
	_ = existsAnyUserActiveEntitlementByItemIdsCmd.MarkFlagRequired("userId")
	existsAnyUserActiveEntitlementByItemIdsCmd.Flags().StringP("itemIds", "i", " ", "Item ids")
	_ = existsAnyUserActiveEntitlementByItemIdsCmd.MarkFlagRequired("itemIds")
}
