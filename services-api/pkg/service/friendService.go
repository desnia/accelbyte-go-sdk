package service

import (
	"encoding/json"
	"errors"

	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclient"
	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclient/friends"
	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/go-openapi/runtime/client"
	"github.com/sirupsen/logrus"
)

type FriendService struct {
	LobbyClient     *lobbyclient.JusticeLobbyService
	TokenRepository repository.TokenRepository
}

func (friendService *FriendService) AddFriends(friendIds []string, namespace, userId string) error {
	accessToken, err := friendService.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	param := &friends.AddFriendsWithoutConfirmationParams{
		Body: &lobbyclientmodels.ModelBulkAddFriendsRequest{
			FriendIds: friendIds,
		},
		Namespace: namespace,
		UserID:    userId,
	}
	_, badRequest, unauthorized, forbidden, serverError, err :=
		friendService.LobbyClient.Friends.AddFriendsWithoutConfirmation(param, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		errorMsg, _ := json.Marshal(*badRequest.GetPayload())
		logrus.Error(string(errorMsg))
		return badRequest
	}
	if unauthorized != nil {
		errorMsg, _ := json.Marshal(*unauthorized.GetPayload())
		logrus.Error(string(errorMsg))
		return unauthorized
	}
	if serverError != nil {
		errorMsg, _ := json.Marshal(*serverError.GetPayload())
		logrus.Error(string(errorMsg))
		return serverError
	}
	if forbidden != nil {
		errorMsg, _ := json.Marshal(*forbidden.GetPayload())
		logrus.Error(string(errorMsg))
		return forbidden
	}
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (friendService *FriendService) GetFriends(namespace, userId, limit, offset string) (*lobbyclientmodels.ModelGetFriendsResponse, error) {
	accessToken, err := friendService.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	param := &friends.GetListOfFriendsParams{
		Limit:     &limit,
		Namespace: namespace,
		Offset:    &offset,
		UserID:    userId,
	}
	listOfFriends, badRequest, unauthorized, forbidden, serverError, err :=
		friendService.LobbyClient.Friends.GetListOfFriends(param, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		errorMsg, _ := json.Marshal(*badRequest.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, badRequest
	}
	if unauthorized != nil {
		errorMsg, _ := json.Marshal(*unauthorized.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, unauthorized
	}
	if serverError != nil {
		errorMsg, _ := json.Marshal(*serverError.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, serverError
	}
	if forbidden != nil {
		errorMsg, _ := json.Marshal(*forbidden.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, forbidden
	}
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if listOfFriends == nil {
		return nil, errors.New("list of friends is empty")
	}
	return listOfFriends.GetPayload(), nil
}

func (friendService *FriendService) AddFriendsWithoutConfirmation(body *lobbyclientmodels.ModelBulkAddFriendsRequest, namespace, userID string) error {
	accessToken, err := friendService.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	param := &friends.AddFriendsWithoutConfirmationParams{
		Body:      body,
		Namespace: namespace,
		UserID:    userID,
	}
	_, badRequest, unauthorized, forbidden, serverError, err :=
		friendService.LobbyClient.Friends.AddFriendsWithoutConfirmation(param, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		errorMsg, _ := json.Marshal(*badRequest.GetPayload())
		logrus.Error(string(errorMsg))
		return badRequest
	}
	if unauthorized != nil {
		errorMsg, _ := json.Marshal(*unauthorized.GetPayload())
		logrus.Error(string(errorMsg))
		return unauthorized
	}
	if serverError != nil {
		errorMsg, _ := json.Marshal(*serverError.GetPayload())
		logrus.Error(string(errorMsg))
		return serverError
	}
	if forbidden != nil {
		errorMsg, _ := json.Marshal(*forbidden.GetPayload())
		logrus.Error(string(errorMsg))
		return forbidden
	}
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (friendService *FriendService) GetListOfFriends(limit *string, namespace string, offset *string, userID string) (*lobbyclientmodels.ModelGetFriendsResponse, error) {
	accessToken, err := friendService.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	param := &friends.GetListOfFriendsParams{
		Limit:     limit,
		Namespace: namespace,
		Offset:    offset,
		UserID:    userID,
	}
	ok, badRequest, unauthorized, forbidden, serverError, err :=
		friendService.LobbyClient.Friends.GetListOfFriends(param, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		errorMsg, _ := json.Marshal(*badRequest.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, badRequest
	}
	if unauthorized != nil {
		errorMsg, _ := json.Marshal(*unauthorized.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, unauthorized
	}
	if serverError != nil {
		errorMsg, _ := json.Marshal(*serverError.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, serverError
	}
	if forbidden != nil {
		errorMsg, _ := json.Marshal(*forbidden.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, forbidden
	}
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (friendService *FriendService) GetUserFriends(namespace string) ([]*lobbyclientmodels.ModelGetUserFriendsResponse, error) {
	accessToken, err := friendService.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	param := &friends.GetUserFriendsParams{
		Namespace: namespace,
	}
	ok, badRequest, unauthorized, forbidden, notFound, internalServerError, err :=
		friendService.LobbyClient.Friends.GetUserFriends(param, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		errorMsg, _ := json.Marshal(*badRequest.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, badRequest
	}
	if unauthorized != nil {
		errorMsg, _ := json.Marshal(*unauthorized.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, unauthorized
	}
	if forbidden != nil {
		errorMsg, _ := json.Marshal(*forbidden.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, forbidden
	}
	if notFound != nil {
		errorMsg, _ := json.Marshal(*notFound.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, notFound
	}
	if internalServerError != nil {
		errorMsg, _ := json.Marshal(*internalServerError.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, internalServerError
	}
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (friendService *FriendService) GetUserIncomingFriends(namespace string) ([]*lobbyclientmodels.ModelGetUserIncomingFriendsResponse, error) {
	accessToken, err := friendService.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	param := &friends.GetUserIncomingFriendsParams{
		Namespace: namespace,
	}
	ok, badRequest, unauthorized, forbidden, notFound, internalServerError, err :=
		friendService.LobbyClient.Friends.GetUserIncomingFriends(param, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		errorMsg, _ := json.Marshal(*badRequest.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, badRequest
	}
	if unauthorized != nil {
		errorMsg, _ := json.Marshal(*unauthorized.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, unauthorized
	}
	if forbidden != nil {
		errorMsg, _ := json.Marshal(*forbidden.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, forbidden
	}
	if notFound != nil {
		errorMsg, _ := json.Marshal(*notFound.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, notFound
	}
	if internalServerError != nil {
		errorMsg, _ := json.Marshal(*internalServerError.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, internalServerError
	}
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (friendService *FriendService) GetUserOutgoingFriends(namespace string) ([]*lobbyclientmodels.ModelGetUserOutgoingFriendsResponse, error) {
	accessToken, err := friendService.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	param := &friends.GetUserOutgoingFriendsParams{
		Namespace: namespace,
	}
	ok, badRequest, unauthorized, forbidden, notFound, internalServerError, err :=
		friendService.LobbyClient.Friends.GetUserOutgoingFriends(param, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		errorMsg, _ := json.Marshal(*badRequest.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, badRequest
	}
	if unauthorized != nil {
		errorMsg, _ := json.Marshal(*unauthorized.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, unauthorized
	}
	if forbidden != nil {
		errorMsg, _ := json.Marshal(*forbidden.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, forbidden
	}
	if notFound != nil {
		errorMsg, _ := json.Marshal(*notFound.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, notFound
	}
	if internalServerError != nil {
		errorMsg, _ := json.Marshal(*internalServerError.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, internalServerError
	}
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (friendService *FriendService) UserAcceptFriendRequest(body *lobbyclientmodels.ModelUserAcceptFriendRequest, namespace string) error {
	accessToken, err := friendService.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	param := &friends.UserAcceptFriendRequestParams{
		Body:      body,
		Namespace: namespace,
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err :=
		friendService.LobbyClient.Friends.UserAcceptFriendRequest(param, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		errorMsg, _ := json.Marshal(*badRequest.GetPayload())
		logrus.Error(string(errorMsg))
		return badRequest
	}
	if unauthorized != nil {
		errorMsg, _ := json.Marshal(*unauthorized.GetPayload())
		logrus.Error(string(errorMsg))
		return unauthorized
	}
	if forbidden != nil {
		errorMsg, _ := json.Marshal(*forbidden.GetPayload())
		logrus.Error(string(errorMsg))
		return forbidden
	}
	if notFound != nil {
		errorMsg, _ := json.Marshal(*notFound.GetPayload())
		logrus.Error(string(errorMsg))
		return notFound
	}
	if internalServerError != nil {
		errorMsg, _ := json.Marshal(*internalServerError.GetPayload())
		logrus.Error(string(errorMsg))
		return internalServerError
	}
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (friendService *FriendService) UserCancelFriendRequest(body *lobbyclientmodels.ModelUserCancelFriendRequest, namespace string) error {
	accessToken, err := friendService.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	param := &friends.UserCancelFriendRequestParams{
		Body:      body,
		Namespace: namespace,
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err :=
		friendService.LobbyClient.Friends.UserCancelFriendRequest(param, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		errorMsg, _ := json.Marshal(*badRequest.GetPayload())
		logrus.Error(string(errorMsg))
		return badRequest
	}
	if unauthorized != nil {
		errorMsg, _ := json.Marshal(*unauthorized.GetPayload())
		logrus.Error(string(errorMsg))
		return unauthorized
	}
	if forbidden != nil {
		errorMsg, _ := json.Marshal(*forbidden.GetPayload())
		logrus.Error(string(errorMsg))
		return forbidden
	}
	if notFound != nil {
		errorMsg, _ := json.Marshal(*notFound.GetPayload())
		logrus.Error(string(errorMsg))
		return notFound
	}
	if internalServerError != nil {
		errorMsg, _ := json.Marshal(*internalServerError.GetPayload())
		logrus.Error(string(errorMsg))
		return internalServerError
	}
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (friendService *FriendService) UserGetFriendshipStatus(friendID, namespace string) (*lobbyclientmodels.ModelUserGetFriendshipStatusResponse, error) {
	accessToken, err := friendService.TokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	param := &friends.UserGetFriendshipStatusParams{
		FriendID:  friendID,
		Namespace: namespace,
	}
	ok, badRequest, unauthorized, forbidden, internalServerError, err :=
		friendService.LobbyClient.Friends.UserGetFriendshipStatus(param, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		errorMsg, _ := json.Marshal(*badRequest.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, badRequest
	}
	if unauthorized != nil {
		errorMsg, _ := json.Marshal(*unauthorized.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, unauthorized
	}
	if forbidden != nil {
		errorMsg, _ := json.Marshal(*forbidden.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, forbidden
	}
	if internalServerError != nil {
		errorMsg, _ := json.Marshal(*internalServerError.GetPayload())
		logrus.Error(string(errorMsg))
		return nil, internalServerError
	}
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return ok.GetPayload(), nil
}

func (friendService *FriendService) UserRejectFriendRequest(body *lobbyclientmodels.ModelUserRejectFriendRequest, namespace string) error {
	accessToken, err := friendService.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	param := &friends.UserRejectFriendRequestParams{
		Body:      body,
		Namespace: namespace,
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err :=
		friendService.LobbyClient.Friends.UserRejectFriendRequest(param, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		errorMsg, _ := json.Marshal(*badRequest.GetPayload())
		logrus.Error(string(errorMsg))
		return badRequest
	}
	if unauthorized != nil {
		errorMsg, _ := json.Marshal(*unauthorized.GetPayload())
		logrus.Error(string(errorMsg))
		return unauthorized
	}
	if forbidden != nil {
		errorMsg, _ := json.Marshal(*forbidden.GetPayload())
		logrus.Error(string(errorMsg))
		return forbidden
	}
	if notFound != nil {
		errorMsg, _ := json.Marshal(*notFound.GetPayload())
		logrus.Error(string(errorMsg))
		return notFound
	}
	if internalServerError != nil {
		errorMsg, _ := json.Marshal(*internalServerError.GetPayload())
		logrus.Error(string(errorMsg))
		return internalServerError
	}
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (friendService *FriendService) UserRequestFriend(body *lobbyclientmodels.ModelRequestFriendsRequest, namespace string) error {
	accessToken, err := friendService.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	param := &friends.UserRequestFriendParams{
		Body:      body,
		Namespace: namespace,
	}
	_, badRequest, unauthorized, forbidden, notFound, unprocessableEntity, internalServerError, err :=
		friendService.LobbyClient.Friends.UserRequestFriend(param, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		errorMsg, _ := json.Marshal(*badRequest.GetPayload())
		logrus.Error(string(errorMsg))
		return badRequest
	}
	if unauthorized != nil {
		errorMsg, _ := json.Marshal(*unauthorized.GetPayload())
		logrus.Error(string(errorMsg))
		return unauthorized
	}
	if forbidden != nil {
		errorMsg, _ := json.Marshal(*forbidden.GetPayload())
		logrus.Error(string(errorMsg))
		return forbidden
	}
	if notFound != nil {
		errorMsg, _ := json.Marshal(*notFound.GetPayload())
		logrus.Error(string(errorMsg))
		return notFound
	}
	if unprocessableEntity != nil {
		errorMsg, _ := json.Marshal(*unprocessableEntity.GetPayload())
		logrus.Error(string(errorMsg))
		return unprocessableEntity
	}
	if internalServerError != nil {
		errorMsg, _ := json.Marshal(*internalServerError.GetPayload())
		logrus.Error(string(errorMsg))
		return internalServerError
	}
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (friendService *FriendService) UserUnfriendRequest(body *lobbyclientmodels.ModelUserUnfriendRequest, namespace string) error {
	accessToken, err := friendService.TokenRepository.GetToken()
	if err != nil {
		return err
	}
	param := &friends.UserUnfriendRequestParams{
		Body:      body,
		Namespace: namespace,
	}
	_, badRequest, unauthorized, forbidden, notFound, internalServerError, err :=
		friendService.LobbyClient.Friends.UserUnfriendRequest(param, client.BearerToken(*accessToken.AccessToken))
	if badRequest != nil {
		errorMsg, _ := json.Marshal(*badRequest.GetPayload())
		logrus.Error(string(errorMsg))
		return badRequest
	}
	if unauthorized != nil {
		errorMsg, _ := json.Marshal(*unauthorized.GetPayload())
		logrus.Error(string(errorMsg))
		return unauthorized
	}
	if forbidden != nil {
		errorMsg, _ := json.Marshal(*forbidden.GetPayload())
		logrus.Error(string(errorMsg))
		return forbidden
	}
	if notFound != nil {
		errorMsg, _ := json.Marshal(*notFound.GetPayload())
		logrus.Error(string(errorMsg))
		return notFound
	}
	if internalServerError != nil {
		errorMsg, _ := json.Marshal(*internalServerError.GetPayload())
		logrus.Error(string(errorMsg))
		return internalServerError
	}
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
