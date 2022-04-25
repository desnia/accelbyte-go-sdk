// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package publicLike

import (
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/ugc"
	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclient/public_like"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// GetLikedContentCmd represents the GetLikedContent command
var GetLikedContentCmd = &cobra.Command{
	Use:   "getLikedContent",
	Short: "Get liked content",
	Long:  `Get liked content`,
	RunE: func(cmd *cobra.Command, args []string) error {
		publicLikeService := &ugc.PublicLikeService{
			Client:          factory.NewUgcClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		limit, _ := cmd.Flags().GetInt64("limit")
		offset, _ := cmd.Flags().GetInt64("offset")
		input := &public_like.GetLikedContentParams{
			Namespace: namespace,
			Limit:     &limit,
			Offset:    &offset,
		}
		ok, err := publicLikeService.GetLikedContentShort(input, nil)
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
	GetLikedContentCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = GetLikedContentCmd.MarkFlagRequired("namespace")
	GetLikedContentCmd.Flags().Int64P("limit", "", 20, "Limit")
	GetLikedContentCmd.Flags().Int64P("offset", "", 0, "Offset")
}
