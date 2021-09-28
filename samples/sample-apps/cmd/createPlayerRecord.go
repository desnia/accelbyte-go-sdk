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

// createPlayerRecordCmd represents the createPlayerRecord command
var createPlayerRecordCmd = &cobra.Command{
	Use:   "createPlayerRecord",
	Short: "Create player records",
	Long:  `Create player records`,
	RunE: func(cmd *cobra.Command, args []string) error {
		userId := cmd.Flag("userId").Value.String()
		namespace := cmd.Flag("namespace").Value.String()
		key := cmd.Flag("key").Value.String()
		var content map[string]string
		contentString := cmd.Flag("content").Value.String()
		err := json.Unmarshal([]byte(contentString), &content)
		if err != nil {
			return err
		}
		cloudSaveService := &service.CloudSaveService{
			CloudSaveServiceClient: factory.NewCloudSaveClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository:        &repository.TokenRepositoryImpl{},
		}
		err = cloudSaveService.PostPlayerPublicRecordHandlerV1(userId, namespace, key, content)
		if err != nil {
			return err
		}
		logrus.Info("Player records successfully created")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createPlayerRecordCmd)
	createPlayerRecordCmd.Flags().StringP("userId", "u", "", "User ID")
	createPlayerRecordCmd.MarkFlagRequired("userId")
	createPlayerRecordCmd.Flags().StringP("namespace", "n", "", "User namespace")
	createPlayerRecordCmd.MarkFlagRequired("namespace")
	createPlayerRecordCmd.Flags().StringP("key", "k", "", "Player records key")
	createPlayerRecordCmd.MarkFlagRequired("key")
	createPlayerRecordCmd.Flags().StringP("content", "c", "", "Player records content. In json format, example: '{\"key1\", \"value1\"}'")
	createPlayerRecordCmd.MarkFlagRequired("content")
}