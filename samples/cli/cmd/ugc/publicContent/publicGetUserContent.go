// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package publicContent

import (
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/ugc"
	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclient/public_content"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// PublicGetUserContentCmd represents the PublicGetUserContent command
var PublicGetUserContentCmd = &cobra.Command{
	Use:   "publicGetUserContent",
	Short: "Public get user content",
	Long:  `Public get user content`,
	RunE: func(cmd *cobra.Command, args []string) error {
		publicContentService := &ugc.PublicContentService{
			Client:          factory.NewUgcClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		limit, _ := cmd.Flags().GetString("limit")
		offset, _ := cmd.Flags().GetString("offset")
		input := &public_content.PublicGetUserContentParams{
			Namespace: namespace,
			UserID:    userId,
			Limit:     &limit,
			Offset:    &offset,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := publicContentService.PublicGetUserContent(input)
		if err != nil {
			logrus.Error(err)
			return err
		} else {
			logrus.Infof("Response CLI success", ok)
		}
		return nil
	},
}

func init() {
	PublicGetUserContentCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = PublicGetUserContentCmd.MarkFlagRequired("namespace")
	PublicGetUserContentCmd.Flags().StringP("userId", "", "", "User id")
	_ = PublicGetUserContentCmd.MarkFlagRequired("userId")
	PublicGetUserContentCmd.Flags().StringP("limit", "", "20", "Limit")
	PublicGetUserContentCmd.Flags().StringP("offset", "", "0", "Offset")
}
