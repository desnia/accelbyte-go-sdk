// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package deploymentConfig

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclient/deployment_config"
	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/dsmc"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// UpdateRootRegionOverrideCmd represents the UpdateRootRegionOverride command
var UpdateRootRegionOverrideCmd = &cobra.Command{
	Use:   "updateRootRegionOverride",
	Short: "Update root region override",
	Long:  `Update root region override`,
	RunE: func(cmd *cobra.Command, args []string) error {
		deploymentConfigService := &dsmc.DeploymentConfigService{
			Client:          factory.NewDsmcClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		bodyString := cmd.Flag("body").Value.String()
		var body *dsmcclientmodels.ModelsUpdateRegionOverrideRequest
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		deployment, _ := cmd.Flags().GetString("deployment")
		namespace, _ := cmd.Flags().GetString("namespace")
		region, _ := cmd.Flags().GetString("region")
		input := &deployment_config.UpdateRootRegionOverrideParams{
			Body:       body,
			Deployment: deployment,
			Namespace:  namespace,
			Region:     region,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := deploymentConfigService.UpdateRootRegionOverride(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	UpdateRootRegionOverrideCmd.Flags().StringP("body", "", "", "Body")
	_ = UpdateRootRegionOverrideCmd.MarkFlagRequired("body")
	UpdateRootRegionOverrideCmd.Flags().StringP("deployment", "", "", "Deployment")
	_ = UpdateRootRegionOverrideCmd.MarkFlagRequired("deployment")
	UpdateRootRegionOverrideCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = UpdateRootRegionOverrideCmd.MarkFlagRequired("namespace")
	UpdateRootRegionOverrideCmd.Flags().StringP("region", "", "", "Region")
	_ = UpdateRootRegionOverrideCmd.MarkFlagRequired("region")
}