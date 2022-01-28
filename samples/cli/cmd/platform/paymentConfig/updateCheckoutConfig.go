// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package paymentConfig

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/payment_config"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// UpdateCheckoutConfigCmd represents the UpdateCheckoutConfig command
var UpdateCheckoutConfigCmd = &cobra.Command{
	Use:   "updateCheckoutConfig",
	Short: "Update checkout config",
	Long:  `Update checkout config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		paymentConfigService := &platform.PaymentConfigService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		id_, _ := cmd.Flags().GetString("id")
		bodyString := cmd.Flag("body").Value.String()
		var body *platformclientmodels.CheckoutConfig
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		sandbox, _ := cmd.Flags().GetBool("sandbox")
		validate, _ := cmd.Flags().GetBool("validate")
		input := &payment_config.UpdateCheckoutConfigParams{
			Body:     body,
			ID:       id_,
			Sandbox:  &sandbox,
			Validate: &validate,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := paymentConfigService.UpdateCheckoutConfig(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	UpdateCheckoutConfigCmd.Flags().StringP("body", "", " ", "Body")
	UpdateCheckoutConfigCmd.Flags().StringP("id", "", " ", "Id")
	_ = UpdateCheckoutConfigCmd.MarkFlagRequired("id")
	UpdateCheckoutConfigCmd.Flags().BoolP("sandbox", "", false, "Sandbox")
	UpdateCheckoutConfigCmd.Flags().BoolP("validate", "", false, "Validate")
}
