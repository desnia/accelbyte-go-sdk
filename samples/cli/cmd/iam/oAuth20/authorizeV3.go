// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package oAuth20

import (
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient/o_auth2_0"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/iam"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// AuthorizeV3Cmd represents the AuthorizeV3 command
var AuthorizeV3Cmd = &cobra.Command{
	Use:   "authorizeV3",
	Short: "Authorize V3",
	Long:  `Authorize V3`,
	RunE: func(cmd *cobra.Command, args []string) error {
		oAuth20Service := &iam.OAuth20Service{
			Client:          factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		clientId, _ := cmd.Flags().GetString("clientId")
		responseType, _ := cmd.Flags().GetString("responseType")
		codeChallenge, _ := cmd.Flags().GetString("codeChallenge")
		codeChallengeMethod, _ := cmd.Flags().GetString("codeChallengeMethod")
		redirectUri, _ := cmd.Flags().GetString("redirectUri")
		scope, _ := cmd.Flags().GetString("scope")
		state, _ := cmd.Flags().GetString("state")
		targetAuthPage, _ := cmd.Flags().GetString("targetAuthPage")
		input := &o_auth2_0.AuthorizeV3Params{
			CodeChallenge:       &codeChallenge,
			CodeChallengeMethod: &codeChallengeMethod,
			RedirectURI:         &redirectUri,
			Scope:               &scope,
			State:               &state,
			TargetAuthPage:      &targetAuthPage,
			ClientID:            clientId,
			ResponseType:        responseType,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		errInput := oAuth20Service.AuthorizeV3(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	AuthorizeV3Cmd.Flags().StringP("code_challenge", "", " ", "Code challenge")
	AuthorizeV3Cmd.Flags().StringP("code_challenge_method", "", " ", "Code challenge method")
	AuthorizeV3Cmd.Flags().StringP("redirect_uri", "", " ", "Redirect uri")
	AuthorizeV3Cmd.Flags().StringP("scope", "", " ", "Scope")
	AuthorizeV3Cmd.Flags().StringP("state", "", " ", "State")
	AuthorizeV3Cmd.Flags().StringP("target_auth_page", "", " ", "Target auth page")
	AuthorizeV3Cmd.Flags().StringP("client_id", "", " ", "Client id")
	_ = AuthorizeV3Cmd.MarkFlagRequired("client_id")
	AuthorizeV3Cmd.Flags().StringP("response_type", "", " ", "Response type")
	_ = AuthorizeV3Cmd.MarkFlagRequired("response_type")
}
