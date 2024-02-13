// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package tags

import (
	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclient/tags"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/cloudsave"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// AdminDeleteTagHandlerV1Cmd represents the AdminDeleteTagHandlerV1 command
var AdminDeleteTagHandlerV1Cmd = &cobra.Command{
	Use:   "adminDeleteTagHandlerV1",
	Short: "Admin delete tag handler V1",
	Long:  `Admin delete tag handler V1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagsService := &cloudsave.TagsService{
			Client:          factory.NewCloudsaveClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		tag, _ := cmd.Flags().GetString("tag")
		input := &tags.AdminDeleteTagHandlerV1Params{
			Namespace: namespace,
			Tag:       tag,
		}
		errCreated := tagsService.AdminDeleteTagHandlerV1Short(input)
		if errCreated != nil {
			logrus.Error(errCreated)

			return errCreated
		}

		logrus.Infof("Response CLI success.")

		return nil
	},
}

func init() {
	AdminDeleteTagHandlerV1Cmd.Flags().String("namespace", "", "Namespace")
	_ = AdminDeleteTagHandlerV1Cmd.MarkFlagRequired("namespace")
	AdminDeleteTagHandlerV1Cmd.Flags().String("tag", "", "Tag")
	_ = AdminDeleteTagHandlerV1Cmd.MarkFlagRequired("tag")
}
