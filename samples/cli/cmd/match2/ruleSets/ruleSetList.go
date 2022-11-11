// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated. DO NOT EDIT.

package ruleSets

import (
	"github.com/AccelByte/accelbyte-go-sdk/match2-sdk/pkg/match2client/rule_sets"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/match2"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// RuleSetListCmd represents the RuleSetList command
var RuleSetListCmd = &cobra.Command{
	Use:   "ruleSetList",
	Short: "Rule set list",
	Long:  `Rule set list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ruleSetsService := &match2.RuleSetsService{
			Client:          factory.NewMatch2Client(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		limit, _ := cmd.Flags().GetInt64("limit")
		offset, _ := cmd.Flags().GetInt64("offset")
		input := &rule_sets.RuleSetListParams{
			Namespace: namespace,
			Limit:     &limit,
			Offset:    &offset,
		}
		ok, err := ruleSetsService.RuleSetListShort(input)
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
	RuleSetListCmd.Flags().String("namespace", "", "Namespace")
	_ = RuleSetListCmd.MarkFlagRequired("namespace")
	RuleSetListCmd.Flags().Int64("limit", 20, "Limit")
	RuleSetListCmd.Flags().Int64("offset", 0, "Offset")
}