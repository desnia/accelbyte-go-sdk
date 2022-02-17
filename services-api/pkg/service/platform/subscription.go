// Copyright (c) 2018 - 2021
// AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package platform

import (
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclient/subscription"
	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/go-openapi/runtime/client"
)

type SubscriptionService struct {
	Client          *platformclient.JusticePlatformService
	TokenRepository repository.TokenRepository
}

// Deprecated: Use QuerySubscriptionsShort instead
func (s *SubscriptionService) QuerySubscriptions(input *subscription.QuerySubscriptionsParams) (*platformclientmodels.SubscriptionPagingSlicedResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.QuerySubscriptions(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use RecurringChargeSubscriptionShort instead
func (s *SubscriptionService) RecurringChargeSubscription(input *subscription.RecurringChargeSubscriptionParams) (*platformclientmodels.RecurringChargeResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.RecurringChargeSubscription(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use QueryUserSubscriptionsShort instead
func (s *SubscriptionService) QueryUserSubscriptions(input *subscription.QueryUserSubscriptionsParams) (*platformclientmodels.SubscriptionPagingSlicedResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.QueryUserSubscriptions(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use GetUserSubscriptionActivitiesShort instead
func (s *SubscriptionService) GetUserSubscriptionActivities(input *subscription.GetUserSubscriptionActivitiesParams) (*platformclientmodels.SubscriptionActivityPagingSlicedResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.GetUserSubscriptionActivities(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use PlatformSubscribeSubscriptionShort instead
func (s *SubscriptionService) PlatformSubscribeSubscription(input *subscription.PlatformSubscribeSubscriptionParams) (*platformclientmodels.SubscriptionInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, created, badRequest, notFound, unprocessableEntity, err := s.Client.Subscription.PlatformSubscribeSubscription(input, client.BearerToken(*accessToken.AccessToken))
	if created != nil {
		return nil, created
	}
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
	}
	if unprocessableEntity != nil {
		return nil, unprocessableEntity
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use CheckUserSubscriptionSubscribableByItemIDShort instead
func (s *SubscriptionService) CheckUserSubscriptionSubscribableByItemID(input *subscription.CheckUserSubscriptionSubscribableByItemIDParams) (*platformclientmodels.Subscribable, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.CheckUserSubscriptionSubscribableByItemID(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use GetUserSubscriptionShort instead
func (s *SubscriptionService) GetUserSubscription(input *subscription.GetUserSubscriptionParams) (*platformclientmodels.SubscriptionInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, err := s.Client.Subscription.GetUserSubscription(input, client.BearerToken(*accessToken.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use DeleteUserSubscriptionShort instead
func (s *SubscriptionService) DeleteUserSubscription(input *subscription.DeleteUserSubscriptionParams) error {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, err = s.Client.Subscription.DeleteUserSubscription(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return err
	}
	return nil
}

// Deprecated: Use CancelSubscriptionShort instead
func (s *SubscriptionService) CancelSubscription(input *subscription.CancelSubscriptionParams) (*platformclientmodels.SubscriptionInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, conflict, err := s.Client.Subscription.CancelSubscription(input, client.BearerToken(*accessToken.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if conflict != nil {
		return nil, conflict
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use GrantDaysToSubscriptionShort instead
func (s *SubscriptionService) GrantDaysToSubscription(input *subscription.GrantDaysToSubscriptionParams) (*platformclientmodels.SubscriptionInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, err := s.Client.Subscription.GrantDaysToSubscription(input, client.BearerToken(*accessToken.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use GetUserSubscriptionBillingHistoriesShort instead
func (s *SubscriptionService) GetUserSubscriptionBillingHistories(input *subscription.GetUserSubscriptionBillingHistoriesParams) (*platformclientmodels.BillingHistoryPagingSlicedResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.GetUserSubscriptionBillingHistories(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use ProcessUserSubscriptionNotificationShort instead
func (s *SubscriptionService) ProcessUserSubscriptionNotification(input *subscription.ProcessUserSubscriptionNotificationParams) error {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, err := s.Client.Subscription.ProcessUserSubscriptionNotification(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if err != nil {
		return err
	}
	return nil
}

// Deprecated: Use PublicQueryUserSubscriptionsShort instead
func (s *SubscriptionService) PublicQueryUserSubscriptions(input *subscription.PublicQueryUserSubscriptionsParams) (*platformclientmodels.SubscriptionPagingSlicedResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.PublicQueryUserSubscriptions(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use PublicSubscribeSubscriptionShort instead
func (s *SubscriptionService) PublicSubscribeSubscription(input *subscription.PublicSubscribeSubscriptionParams) error {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, badRequest, forbidden, notFound, conflict, unprocessableEntity, err := s.Client.Subscription.PublicSubscribeSubscription(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return badRequest
	}
	if forbidden != nil {
		return forbidden
	}
	if notFound != nil {
		return notFound
	}
	if conflict != nil {
		return conflict
	}
	if unprocessableEntity != nil {
		return unprocessableEntity
	}
	if err != nil {
		return err
	}
	return nil
}

// Deprecated: Use PublicCheckUserSubscriptionSubscribableByItemIDShort instead
func (s *SubscriptionService) PublicCheckUserSubscriptionSubscribableByItemID(input *subscription.PublicCheckUserSubscriptionSubscribableByItemIDParams) (*platformclientmodels.Subscribable, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.PublicCheckUserSubscriptionSubscribableByItemID(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use PublicGetUserSubscriptionShort instead
func (s *SubscriptionService) PublicGetUserSubscription(input *subscription.PublicGetUserSubscriptionParams) (*platformclientmodels.SubscriptionInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, err := s.Client.Subscription.PublicGetUserSubscription(input, client.BearerToken(*accessToken.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use PublicChangeSubscriptionBillingAccountShort instead
func (s *SubscriptionService) PublicChangeSubscriptionBillingAccount(input *subscription.PublicChangeSubscriptionBillingAccountParams) (*platformclientmodels.SubscriptionInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, badRequest, notFound, conflict, err := s.Client.Subscription.PublicChangeSubscriptionBillingAccount(input, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		return nil, badRequest
	}
	if notFound != nil {
		return nil, notFound
	}
	if conflict != nil {
		return nil, conflict
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use PublicCancelSubscriptionShort instead
func (s *SubscriptionService) PublicCancelSubscription(input *subscription.PublicCancelSubscriptionParams) (*platformclientmodels.SubscriptionInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, notFound, conflict, err := s.Client.Subscription.PublicCancelSubscription(input, client.BearerToken(*accessToken.AccessToken))
	if notFound != nil {
		return nil, notFound
	}
	if conflict != nil {
		return nil, conflict
	}
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

// Deprecated: Use PublicGetUserSubscriptionBillingHistoriesShort instead
func (s *SubscriptionService) PublicGetUserSubscriptionBillingHistories(input *subscription.PublicGetUserSubscriptionBillingHistoriesParams) (*platformclientmodels.BillingHistoryPagingSlicedResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.PublicGetUserSubscriptionBillingHistories(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SubscriptionService) QuerySubscriptionsShort(input *subscription.QuerySubscriptionsParams) (*platformclientmodels.SubscriptionPagingSlicedResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.QuerySubscriptionsShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SubscriptionService) RecurringChargeSubscriptionShort(input *subscription.RecurringChargeSubscriptionParams) (*platformclientmodels.RecurringChargeResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.RecurringChargeSubscriptionShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SubscriptionService) QueryUserSubscriptionsShort(input *subscription.QueryUserSubscriptionsParams) (*platformclientmodels.SubscriptionPagingSlicedResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.QueryUserSubscriptionsShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SubscriptionService) GetUserSubscriptionActivitiesShort(input *subscription.GetUserSubscriptionActivitiesParams) (*platformclientmodels.SubscriptionActivityPagingSlicedResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.GetUserSubscriptionActivitiesShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SubscriptionService) PlatformSubscribeSubscriptionShort(input *subscription.PlatformSubscribeSubscriptionParams) (*platformclientmodels.SubscriptionInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.PlatformSubscribeSubscriptionShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SubscriptionService) CheckUserSubscriptionSubscribableByItemIDShort(input *subscription.CheckUserSubscriptionSubscribableByItemIDParams) (*platformclientmodels.Subscribable, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.CheckUserSubscriptionSubscribableByItemIDShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SubscriptionService) GetUserSubscriptionShort(input *subscription.GetUserSubscriptionParams) (*platformclientmodels.SubscriptionInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.GetUserSubscriptionShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SubscriptionService) DeleteUserSubscriptionShort(input *subscription.DeleteUserSubscriptionParams) error {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, err = s.Client.Subscription.DeleteUserSubscriptionShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return err
	}
	return nil
}

func (s *SubscriptionService) CancelSubscriptionShort(input *subscription.CancelSubscriptionParams) (*platformclientmodels.SubscriptionInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.CancelSubscriptionShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SubscriptionService) GrantDaysToSubscriptionShort(input *subscription.GrantDaysToSubscriptionParams) (*platformclientmodels.SubscriptionInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.GrantDaysToSubscriptionShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SubscriptionService) GetUserSubscriptionBillingHistoriesShort(input *subscription.GetUserSubscriptionBillingHistoriesParams) (*platformclientmodels.BillingHistoryPagingSlicedResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.GetUserSubscriptionBillingHistoriesShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SubscriptionService) ProcessUserSubscriptionNotificationShort(input *subscription.ProcessUserSubscriptionNotificationParams) error {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, err = s.Client.Subscription.ProcessUserSubscriptionNotificationShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return err
	}
	return nil
}

func (s *SubscriptionService) PublicQueryUserSubscriptionsShort(input *subscription.PublicQueryUserSubscriptionsParams) (*platformclientmodels.SubscriptionPagingSlicedResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.PublicQueryUserSubscriptionsShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SubscriptionService) PublicSubscribeSubscriptionShort(input *subscription.PublicSubscribeSubscriptionParams) error {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	_, err = s.Client.Subscription.PublicSubscribeSubscriptionShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return err
	}
	return nil
}

func (s *SubscriptionService) PublicCheckUserSubscriptionSubscribableByItemIDShort(input *subscription.PublicCheckUserSubscriptionSubscribableByItemIDParams) (*platformclientmodels.Subscribable, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.PublicCheckUserSubscriptionSubscribableByItemIDShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SubscriptionService) PublicGetUserSubscriptionShort(input *subscription.PublicGetUserSubscriptionParams) (*platformclientmodels.SubscriptionInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.PublicGetUserSubscriptionShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SubscriptionService) PublicChangeSubscriptionBillingAccountShort(input *subscription.PublicChangeSubscriptionBillingAccountParams) (*platformclientmodels.SubscriptionInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.PublicChangeSubscriptionBillingAccountShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SubscriptionService) PublicCancelSubscriptionShort(input *subscription.PublicCancelSubscriptionParams) (*platformclientmodels.SubscriptionInfo, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.PublicCancelSubscriptionShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (s *SubscriptionService) PublicGetUserSubscriptionBillingHistoriesShort(input *subscription.PublicGetUserSubscriptionBillingHistoriesParams) (*platformclientmodels.BillingHistoryPagingSlicedResult, error) {
	accessToken, err := s.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	ok, err := s.Client.Subscription.PublicGetUserSubscriptionBillingHistoriesShort(input, client.BearerToken(*accessToken.AccessToken))
	if err != nil {
		return nil, err
	}
	return ok.GetPayload(), nil
}
