// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.
package cmd

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/iam"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// createStoreCmd represents the createStore command
var createStoreCmd = &cobra.Command{
	Use:   "createStore",
	Short: "Create store",
	Long:  `Create store`,
	RunE: func(cmd *cobra.Command, args []string) error {

		namespace := cmd.Flag("namespace").Value.String()
		defaultLang := cmd.Flag("language").Value.String()
		defaultRegion := cmd.Flag("region").Value.String()
		description := cmd.Flag("description").Value.String()
		title := cmd.Flag("title").Value.String()
		storeService := &service.StoreService{
			OauthService: &iam.OAuth20Service{
				Client:           factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
				ConfigRepository: &repository.ConfigRepositoryImpl{},
				TokenRepository:  &repository.TokenRepositoryImpl{},
			},
			PlatformService: factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
		}
		storeCreateParam := &platformclientmodels.StoreCreate{
			DefaultLanguage:    defaultLang,
			DefaultRegion:      defaultRegion,
			Description:        description,
			SupportedLanguages: []string{},
			SupportedRegions:   []string{},
			Title:              &title,
		}
		createStore, err := storeService.CreateStore(namespace, *storeCreateParam)
		if err != nil {
			return err
		}
		response, err := json.Marshal(createStore)
		if err != nil {
			return err
		}
		logrus.Infof("Response: %s", response)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createStoreCmd)
	createStoreCmd.Flags().StringP("namespace", "n", "", "Store namespace")
	createStoreCmd.MarkFlagRequired("namespace")
	createStoreCmd.Flags().StringP("title", "t", "", "Store title")
	createStoreCmd.MarkFlagRequired("title")
	createStoreCmd.Flags().StringP("region", "r", "", "Store region")
	createStoreCmd.MarkFlagRequired("region")
	createStoreCmd.Flags().StringP("language", "l", "", "Store language")
	createStoreCmd.MarkFlagRequired("language")
	createStoreCmd.Flags().StringP("description", "d", "", "Store description")
}
