// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package adminContent

import (
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/ugc"
	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclient/admin_content"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// AdminDeleteContentScreenshotCmd represents the AdminDeleteContentScreenshot command
var AdminDeleteContentScreenshotCmd = &cobra.Command{
	Use:   "adminDeleteContentScreenshot",
	Short: "Admin delete content screenshot",
	Long:  `Admin delete content screenshot`,
	RunE: func(cmd *cobra.Command, args []string) error {
		adminContentService := &ugc.AdminContentService{
			Client:          factory.NewUgcClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		contentId, _ := cmd.Flags().GetString("contentId")
		namespace, _ := cmd.Flags().GetString("namespace")
		screenshotId, _ := cmd.Flags().GetString("screenshotId")
		input := &admin_content.AdminDeleteContentScreenshotParams{
			ContentID:    contentId,
			Namespace:    namespace,
			ScreenshotID: screenshotId,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		errInput := adminContentService.AdminDeleteContentScreenshot(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	AdminDeleteContentScreenshotCmd.Flags().StringP("contentId", "", "", "Content id")
	_ = AdminDeleteContentScreenshotCmd.MarkFlagRequired("contentId")
	AdminDeleteContentScreenshotCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = AdminDeleteContentScreenshotCmd.MarkFlagRequired("namespace")
	AdminDeleteContentScreenshotCmd.Flags().StringP("screenshotId", "", "", "Screenshot id")
	_ = AdminDeleteContentScreenshotCmd.MarkFlagRequired("screenshotId")
}