// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package users

import (
	"encoding/json"
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient/users"
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/iam"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CreateUserFromInvitationV3Cmd represents the CreateUserFromInvitationV3 command
var CreateUserFromInvitationV3Cmd = &cobra.Command{
	Use:   "createUserFromInvitationV3",
	Short: "Create user from invitation V3",
	Long:  `Create user from invitation V3`,
	RunE: func(cmd *cobra.Command, args []string) error {
		usersService := &iam.UsersService{
			Client:          factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		bodyString := cmd.Flag("body").Value.String()
		var body *iamclientmodels.ModelUserCreateFromInvitationRequestV3
		errBody := json.Unmarshal([]byte(bodyString), &body)
		if errBody != nil {
			return errBody
		}
		invitationId, _ := cmd.Flags().GetString("invitationId")
		namespace, _ := cmd.Flags().GetString("namespace")
		input := &users.CreateUserFromInvitationV3Params{
			Body:         body,
			InvitationID: invitationId,
			Namespace:    namespace,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		ok, err := usersService.CreateUserFromInvitationV3(input)
		logrus.Infof("Response %v", ok)
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	},
}

func init() {
	CreateUserFromInvitationV3Cmd.Flags().StringP("body", "", " ", "Body")
	_ = CreateUserFromInvitationV3Cmd.MarkFlagRequired("body")
	CreateUserFromInvitationV3Cmd.Flags().StringP("invitationId", "", " ", "Invitation id")
	_ = CreateUserFromInvitationV3Cmd.MarkFlagRequired("invitationId")
	CreateUserFromInvitationV3Cmd.Flags().StringP("namespace", "", " ", "Namespace")
	_ = CreateUserFromInvitationV3Cmd.MarkFlagRequired("namespace")
}
