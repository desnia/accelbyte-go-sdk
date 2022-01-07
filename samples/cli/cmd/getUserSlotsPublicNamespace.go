// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/social"
	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclient/slot"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getUserSlotsPublicNamespacePublicNamespaceCmd represents the getUserSlotsPublicNamespace command
var getUserSlotsPublicNamespacePublicNamespaceCmd = &cobra.Command{
	Use:   "getUserSlotsPublicNamespace",
	Short: "Get user slots public namespace",
	Long:  `Get user slots public namespace`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("getUserSlotsPublicNamespace called")
		socialService := &social.SlotService{
			Client:          factory.NewSocialClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace := cmd.Flag("namespace").Value.String()
		userId := cmd.Flag("userId").Value.String()
		input := &slot.PublicGetUserNamespaceSlotsParams{
			Namespace: namespace,
			UserID:    userId,
		}
		//nolint:staticcheck // SA1019 To be deprecated later
		//lint:ignore SA1019 Ignore the deprecation warnings
		slots, err := socialService.PublicGetUserNamespaceSlots(input)
		if err != nil {
			return err
		}
		response, err := json.MarshalIndent(slots, "", "    ")
		if err != nil {
			return err
		}
		logrus.Infof("Response %v", string(response))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getUserSlotsPublicNamespacePublicNamespaceCmd)
	getUserSlotsPublicNamespacePublicNamespaceCmd.Flags().StringP("namespace", "n", "", "User namespace")
	_ = getUserSlotsPublicNamespacePublicNamespaceCmd.MarkFlagRequired("namespace")
	getUserSlotsPublicNamespacePublicNamespaceCmd.Flags().StringP("userId", "u", "", "User ID")
	_ = getUserSlotsPublicNamespacePublicNamespaceCmd.MarkFlagRequired("userId")
}