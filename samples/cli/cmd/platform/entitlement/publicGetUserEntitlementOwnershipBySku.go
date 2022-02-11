// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package entitlement

import (
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/entitlement"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// PublicGetUserEntitlementOwnershipBySkuCmd represents the PublicGetUserEntitlementOwnershipBySku command
var PublicGetUserEntitlementOwnershipBySkuCmd = &cobra.Command{
	Use:   "publicGetUserEntitlementOwnershipBySku",
	Short: "Public get user entitlement ownership by sku",
	Long:  `Public get user entitlement ownership by sku`,
	RunE: func(cmd *cobra.Command, args []string) error {
		entitlementService := &platform.EntitlementService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		userId, _ := cmd.Flags().GetString("userId")
		sku, _ := cmd.Flags().GetString("sku")
		entitlementClazz, _ := cmd.Flags().GetString("entitlementClazz")
		input := &entitlement.PublicGetUserEntitlementOwnershipBySkuParams{
			Namespace:        namespace,
			UserID:           userId,
			EntitlementClazz: &entitlementClazz,
			Sku:              sku,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := entitlementService.PublicGetUserEntitlementOwnershipBySku(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	PublicGetUserEntitlementOwnershipBySkuCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = PublicGetUserEntitlementOwnershipBySkuCmd.MarkFlagRequired("namespace")
	PublicGetUserEntitlementOwnershipBySkuCmd.Flags().StringP("userId", "", "", "User id")
	_ = PublicGetUserEntitlementOwnershipBySkuCmd.MarkFlagRequired("userId")
	PublicGetUserEntitlementOwnershipBySkuCmd.Flags().StringP("entitlementClazz", "", "", "Entitlement clazz")
	PublicGetUserEntitlementOwnershipBySkuCmd.Flags().StringP("sku", "", "", "Sku")
	_ = PublicGetUserEntitlementOwnershipBySkuCmd.MarkFlagRequired("sku")
}