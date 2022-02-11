// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package group

import (
	"encoding/json"
	group_ "github.com/AccelByte/accelbyte-go-sdk/group-sdk/pkg/groupclient/group"
	"github.com/AccelByte/accelbyte-go-sdk/group-sdk/pkg/groupclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/group"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// UpdateGroupCustomAttributesPublicV1Cmd represents the UpdateGroupCustomAttributesPublicV1 command
var UpdateGroupCustomAttributesPublicV1Cmd = &cobra.Command{
	Use:   "updateGroupCustomAttributesPublicV1",
	Short: "Update group custom attributes public V1",
	Long:  `Update group custom attributes public V1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		groupService := &group.GroupService{
			Client:          factory.NewGroupClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		bodyString := cmd.Flag("body").Value.String()
		var body *groupclientmodels.ModelsUpdateGroupCustomAttributesRequestV1
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		groupId, _ := cmd.Flags().GetString("groupId")
		namespace, _ := cmd.Flags().GetString("namespace")
		input := &group_.UpdateGroupCustomAttributesPublicV1Params{
			Body:      body,
			GroupID:   groupId,
			Namespace: namespace,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := groupService.UpdateGroupCustomAttributesPublicV1(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	UpdateGroupCustomAttributesPublicV1Cmd.Flags().StringP("body", "", "", "Body")
	_ = UpdateGroupCustomAttributesPublicV1Cmd.MarkFlagRequired("body")
	UpdateGroupCustomAttributesPublicV1Cmd.Flags().StringP("groupId", "", "", "Group id")
	_ = UpdateGroupCustomAttributesPublicV1Cmd.MarkFlagRequired("groupId")
	UpdateGroupCustomAttributesPublicV1Cmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = UpdateGroupCustomAttributesPublicV1Cmd.MarkFlagRequired("namespace")
}