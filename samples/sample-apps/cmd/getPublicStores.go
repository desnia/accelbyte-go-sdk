// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.
package cmd

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getPublicStoresCmd represents the getPublicStores command
var getPublicStoresCmd = &cobra.Command{
	Use:   "getPublicStores",
	Short: "Get public stores",
	Long:  `Get public stores`,
	RunE: func(cmd *cobra.Command, args []string) error {
		namespace := cmd.Flag("namespace").Value.String()
		storeService := &service.StoreService{
			OauthService: &service.OauthService{
				IamService:       factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
				ConfigRepository: &repository.ConfigRepositoryImpl{},
				TokenRepository:  &repository.TokenRepositoryImpl{},
			},
			PlatformService: factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
		}
		listStore, err := storeService.PublicListStore(namespace)
		if err != nil {
			return err
		}
		response, err := json.Marshal(listStore)
		if err != nil {
			return err
		}
		logrus.Infof("Response %s", response)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getPublicStoresCmd)
	getPublicStoresCmd.Flags().StringP("namespace", "n", "", "Store namespace")
	getPublicStoresCmd.MarkFlagRequired("namespace")
}