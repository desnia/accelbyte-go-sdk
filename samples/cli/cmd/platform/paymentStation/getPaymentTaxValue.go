// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package paymentStation

import (
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/payment_station"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/platform"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// GetPaymentTaxValueCmd represents the GetPaymentTaxValue command
var GetPaymentTaxValueCmd = &cobra.Command{
	Use:   "getPaymentTaxValue",
	Short: "Get payment tax value",
	Long:  `Get payment tax value`,
	RunE: func(cmd *cobra.Command, args []string) error {
		paymentStationService := &platform.PaymentStationService{
			Client:          factory.NewPlatformClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		paymentOrderNo, _ := cmd.Flags().GetString("paymentOrderNo")
		paymentProvider, _ := cmd.Flags().GetString("paymentProvider")
		zipCode, _ := cmd.Flags().GetString("zipCode")
		input := &payment_station.GetPaymentTaxValueParams{
			Namespace:       namespace,
			ZipCode:         &zipCode,
			PaymentOrderNo:  paymentOrderNo,
			PaymentProvider: paymentProvider,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := paymentStationService.GetPaymentTaxValue(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	GetPaymentTaxValueCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = GetPaymentTaxValueCmd.MarkFlagRequired("namespace")
	GetPaymentTaxValueCmd.Flags().StringP("zipCode", "", "", "Zip code")
	GetPaymentTaxValueCmd.Flags().StringP("paymentOrderNo", "", "", "Payment order no")
	_ = GetPaymentTaxValueCmd.MarkFlagRequired("paymentOrderNo")
	GetPaymentTaxValueCmd.Flags().StringP("paymentProvider", "", "", "Payment provider")
	_ = GetPaymentTaxValueCmd.MarkFlagRequired("paymentProvider")
}