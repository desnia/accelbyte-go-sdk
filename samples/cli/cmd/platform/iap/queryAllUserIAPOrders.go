// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package iap

import (
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/i_a_p"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// QueryAllUserIAPOrdersCmd represents the QueryAllUserIAPOrders command
var QueryAllUserIAPOrdersCmd = &cobra.Command{
	Use:   "queryAllUserIAPOrders",
	Short: "Query all user IAP orders",
	Long:  `Query all user IAP orders`,
	RunE: func(cmd *cobra.Command, args []string) error {
		iapService := &platform.IAPService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		input := &i_a_p.QueryAllUserIAPOrdersParams{
			Namespace: namespace,
			UserID:    userId,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := iapService.QueryAllUserIAPOrders(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	QueryAllUserIAPOrdersCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = QueryAllUserIAPOrdersCmd.MarkFlagRequired("namespace")
	QueryAllUserIAPOrdersCmd.Flags().StringP("userId", "", "", "User id")
	_ = QueryAllUserIAPOrdersCmd.MarkFlagRequired("userId")
}