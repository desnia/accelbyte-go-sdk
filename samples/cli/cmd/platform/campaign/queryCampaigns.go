// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package campaign

import (
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/campaign"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// QueryCampaignsCmd represents the QueryCampaigns command
var QueryCampaignsCmd = &cobra.Command{
	Use:   "queryCampaigns",
	Short: "Query campaigns",
	Long:  `Query campaigns`,
	RunE: func(cmd *cobra.Command, args []string) error {
		campaignService := &platform.CampaignService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		limit, _ := cmd.Flags().GetInt32("limit")
		name, _ := cmd.Flags().GetString("name")
		offset, _ := cmd.Flags().GetInt32("offset")
		tag, _ := cmd.Flags().GetString("tag")
		input := &campaign.QueryCampaignsParams{
			Namespace: namespace,
			Limit:     &limit,
			Name:      &name,
			Offset:    &offset,
			Tag:       &tag,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := campaignService.QueryCampaigns(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	QueryCampaignsCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = QueryCampaignsCmd.MarkFlagRequired("namespace")
	QueryCampaignsCmd.Flags().Int32P("limit", "", 20, "Limit")
	QueryCampaignsCmd.Flags().StringP("name", "", "", "Name")
	QueryCampaignsCmd.Flags().Int32P("offset", "", 0, "Offset")
	QueryCampaignsCmd.Flags().StringP("tag", "", "", "Tag")
}