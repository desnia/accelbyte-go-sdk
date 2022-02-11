// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package item

import (
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/item"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// PublicBulkGetItemsCmd represents the PublicBulkGetItems command
var PublicBulkGetItemsCmd = &cobra.Command{
	Use:   "publicBulkGetItems",
	Short: "Public bulk get items",
	Long:  `Public bulk get items`,
	RunE: func(cmd *cobra.Command, args []string) error {
		itemService := &platform.ItemService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		itemIds, _ := cmd.Flags().GetString("itemIds")
		language, _ := cmd.Flags().GetString("language")
		region, _ := cmd.Flags().GetString("region")
		storeId, _ := cmd.Flags().GetString("storeId")
		input := &item.PublicBulkGetItemsParams{
			Namespace: namespace,
			Language:  &language,
			Region:    &region,
			StoreID:   &storeId,
			ItemIds:   itemIds,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := itemService.PublicBulkGetItems(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	PublicBulkGetItemsCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = PublicBulkGetItemsCmd.MarkFlagRequired("namespace")
	PublicBulkGetItemsCmd.Flags().StringP("language", "", "", "Language")
	PublicBulkGetItemsCmd.Flags().StringP("region", "", "", "Region")
	PublicBulkGetItemsCmd.Flags().StringP("storeId", "", "", "Store id")
	PublicBulkGetItemsCmd.Flags().StringP("itemIds", "", "", "Item ids")
	_ = PublicBulkGetItemsCmd.MarkFlagRequired("itemIds")
}