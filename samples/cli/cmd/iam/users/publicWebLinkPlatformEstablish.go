// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package users

import (
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient/users"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/iam"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
)

// PublicWebLinkPlatformEstablishCmd represents the PublicWebLinkPlatformEstablish command
var PublicWebLinkPlatformEstablishCmd = &cobra.Command{
	Use:   "publicWebLinkPlatformEstablish",
	Short: "Public web link platform establish",
	Long:  `Public web link platform establish`,
	RunE: func(cmd *cobra.Command, args []string) error {
		usersService := &iam.UsersService{
			Client:          factory.NewIamClient(&repository.ConfigRepositoryImpl{}),
			TokenRepository: &repository.TokenRepositoryImpl{},
		}
		namespace, _ := cmd.Flags().GetString("namespace")
		platformId, _ := cmd.Flags().GetString("platformId")
		state, _ := cmd.Flags().GetString("state")
		httpClient := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		input := &users.PublicWebLinkPlatformEstablishParams{
			Namespace:  namespace,
			PlatformID: platformId,
			State:      state,
			HTTPClient: httpClient,
		}
		//lint:ignore SA1019 Ignore the deprecation warnings
		errInput := usersService.PublicWebLinkPlatformEstablish(input)
		if errInput != nil {
			logrus.Error(errInput)
			return errInput
		}
		return nil
	},
}

func init() {
	PublicWebLinkPlatformEstablishCmd.Flags().StringP("namespace", "", "", "Namespace")
	_ = PublicWebLinkPlatformEstablishCmd.MarkFlagRequired("namespace")
	PublicWebLinkPlatformEstablishCmd.Flags().StringP("platformId", "", "", "Platform id")
	_ = PublicWebLinkPlatformEstablishCmd.MarkFlagRequired("platformId")
	PublicWebLinkPlatformEstablishCmd.Flags().StringP("state", "", "", "State")
	_ = PublicWebLinkPlatformEstablishCmd.MarkFlagRequired("state")
}