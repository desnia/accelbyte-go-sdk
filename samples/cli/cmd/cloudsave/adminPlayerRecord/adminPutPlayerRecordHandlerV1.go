// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package adminPlayerRecord

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclient/admin_player_record"
	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/cloudsave"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// AdminPutPlayerRecordHandlerV1Cmd represents the AdminPutPlayerRecordHandlerV1 command
var AdminPutPlayerRecordHandlerV1Cmd = &cobra.Command{
	Use:   "adminPutPlayerRecordHandlerV1",
	Short: "Admin put player record handler V1",
	Long:  `Admin put player record handler V1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		adminPlayerRecordService := &cloudsave.AdminPlayerRecordService{
			Client:          factory.NewCloudsaveClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		bodyString := cmd.Flag("body").Value.String()
		var body *cloudsaveclientmodels.ModelsPlayerRecordRequest
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		key, _ := cmd.Flags().GetString("key")
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		input := &admin_player_record.AdminPutPlayerRecordHandlerV1Params{
			Body:      body,
			Key:       key,
			Namespace: namespace,
			UserID:    userId,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		errInput := adminPlayerRecordService.AdminPutPlayerRecordHandlerV1(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	AdminPutPlayerRecordHandlerV1Cmd.Flags().StringP("body", "", "", "Body")
	_ = AdminPutPlayerRecordHandlerV1Cmd.MarkFlagRequired("body")
	AdminPutPlayerRecordHandlerV1Cmd.Flags().StringP("key", "", "", "Key")
	_ = AdminPutPlayerRecordHandlerV1Cmd.MarkFlagRequired("key")
	AdminPutPlayerRecordHandlerV1Cmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = AdminPutPlayerRecordHandlerV1Cmd.MarkFlagRequired("namespace")
	AdminPutPlayerRecordHandlerV1Cmd.Flags().StringP("userId", "", "", "User id")
	_ = AdminPutPlayerRecordHandlerV1Cmd.MarkFlagRequired("userId")
}