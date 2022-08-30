// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package item

import (
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/item"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// QueryItems1Cmd represents the QueryItems1 command
var QueryItems1Cmd = &cobra.Command{
	Use:   "queryItems1",
	Short: "Query items 1",
	Long:  `Query items 1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		itemService := &platform.ItemService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		appType, _ := cmd.Flags().GetString("appType")
		availableDate, _ := cmd.Flags().GetString("availableDate")
		baseAppId, _ := cmd.Flags().GetString("baseAppId")
		categoryPath, _ := cmd.Flags().GetString("categoryPath")
		features, _ := cmd.Flags().GetString("features")
		itemStatus, _ := cmd.Flags().GetString("itemStatus")
		itemType, _ := cmd.Flags().GetString("itemType")
		limit, _ := cmd.Flags().GetInt32("limit")
		offset, _ := cmd.Flags().GetInt32("offset")
		region, _ := cmd.Flags().GetString("region")
		sortBy, _ := cmd.Flags().GetString("sortBy")
		storeId, _ := cmd.Flags().GetString("storeId")
		tags, _ := cmd.Flags().GetString("tags")
		targetNamespace, _ := cmd.Flags().GetString("targetNamespace")
		input := &item.QueryItems1Params{
			Namespace:       namespace,
			AppType:         &appType,
			AvailableDate:   &availableDate,
			BaseAppID:       &baseAppId,
			CategoryPath:    &categoryPath,
			Features:        &features,
			ItemStatus:      &itemStatus,
			ItemType:        &itemType,
			Limit:           &limit,
			Offset:          &offset,
			Region:          &region,
			SortBy:          &sortBy,
			StoreID:         &storeId,
			Tags:            &tags,
			TargetNamespace: &targetNamespace,
		}
		ok, err := itemService.QueryItems1Short(input)
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
	QueryItems1Cmd.Flags().String("namespace", "", "Namespace")
	_ = QueryItems1Cmd.MarkFlagRequired("namespace")
	QueryItems1Cmd.Flags().String("appType", "", "App type")
	QueryItems1Cmd.Flags().String("availableDate", "", "Available date")
	QueryItems1Cmd.Flags().String("baseAppId", "", "Base app id")
	QueryItems1Cmd.Flags().String("categoryPath", "", "Category path")
	QueryItems1Cmd.Flags().String("features", "", "Features")
	QueryItems1Cmd.Flags().String("itemStatus", "", "Item status")
	QueryItems1Cmd.Flags().String("itemType", "", "Item type")
	QueryItems1Cmd.Flags().Int32("limit", 20, "Limit")
	QueryItems1Cmd.Flags().Int32("offset", 0, "Offset")
	QueryItems1Cmd.Flags().String("region", "", "Region")
	QueryItems1Cmd.Flags().String("sortBy", "", "Sort by")
	QueryItems1Cmd.Flags().String("storeId", "", "Store id")
	QueryItems1Cmd.Flags().String("tags", "", "Tags")
	QueryItems1Cmd.Flags().String("targetNamespace", "", "Target namespace")
}