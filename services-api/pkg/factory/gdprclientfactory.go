// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.
package factory

import (
	"strings"

	"github.com/AccelByte/accelbyte-go-sdk/gdpr-sdk/pkg/gdprclient"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils"
)

var gdprClientInstance *gdprclient.JusticeGdprService

func NewGdprClient(configRepository repository.ConfigRepository) *gdprclient.JusticeGdprService {
	if gdprClientInstance == nil {
		baseUrl := configRepository.GetJusticeBaseUrl()
		xAmazonTraceId := utils.AmazonTraceIdGen()
		userAgent := utils.UserAgentGen()
		if len(baseUrl) > 0 {
			baseUrlSplit := strings.Split(baseUrl, "://")
			httpClientConfig := &gdprclient.TransportConfig{
				Host:     baseUrlSplit[1],
				BasePath: "",
				Schemes:  []string{baseUrlSplit[0]},
			}
			gdprClientInstance = gdprclient.NewHTTPClientWithConfig(nil, httpClientConfig, userAgent, xAmazonTraceId)
		} else {
			gdprClientInstance = gdprclient.NewHTTPClient(nil)
		}
	}

	return gdprClientInstance
}
