// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/model"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils/connectionutils"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils/parser"
	"github.com/AccelByte/sample-apps/cmd"
	"github.com/AccelByte/sample-apps/pkg/repository"
	"github.com/AccelByte/sample-apps/pkg/utils"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

var (
	reader        *bufio.Reader
	connMgr       *utils.ConnectionManagerImpl
	partyService  *service.PartyService
	friendService *service.FriendServiceWebsocket
	chatService   *service.ChatService
)

const (
	infoCmd                   = "1"
	createCmd                 = "2"
	leaveCmd                  = "3"
	inviteCmd                 = "4"
	joinCmd                   = "5"
	kickCmd                   = "6"
	rejectCmd                 = "7"
	promoteLeaderCmd          = "8"
	chatCmd                   = "9"
	partyChatCmd              = "10"
	getOnlineFriendsCmd       = "11"
	getOfflineNotificationCmd = "12"
	getFriendsCmd             = "13"
	setUserStatusCmd          = "14"
	requestFriendsCmd         = "15"
	listIncomingFriendsCmd    = "16"
	listOutgoingFriendsCmd    = "17"
	acceptFriendsCmd          = "18"
	rejectFriendsCmd          = "19"
	cancelFriendsRequestCmd   = "20"
	unfriendCmd               = "21"
	listOfFriendsCmd          = "22"
	getFriendshipStatusCmd    = "23"
	personalChatHistoryCmd    = "24"
	joinDefaultChatChannelCmd = "25"
	sendChatChannelCmd        = "26"
	blockCmd                  = "27"
	unblockCmd                = "28"
	quitCmd                   = "99"
)

func main() {
	args := os.Args
	standbyModeFlag := args[1]
	if standbyModeFlag == "--wsMode" {
		reader = bufio.NewReader(os.Stdin)
		logrus.Info("Enter websocket mode")
		connMgr = &utils.ConnectionManagerImpl{}
		configRepo := &repository.ConfigRepositoryImpl{}
		tokenRepo := &repository.TokenRepositoryImpl{}
		connection, err := connectionutils.NewWebsocketConnection(configRepo, tokenRepo, messageHandler)
		if err != nil {
			panic(err)
		}
		connMgr.Save(connection)
		partyService = &service.PartyService{
			ConfigRepository:  configRepo,
			TokenRepository:   tokenRepo,
			ConnectionManager: connMgr,
		}
		friendService = &service.FriendServiceWebsocket{
			ConfigRepository:  configRepo,
			TokenRepository:   tokenRepo,
			ConnectionManager: connMgr,
		}
		chatService = &service.ChatService{
			ConfigRepository:  configRepo,
			TokenRepository:   tokenRepo,
			ConnectionManager: connMgr,
		}

		serve()

		defer connMgr.Close()
		logrus.Info("Done")
	} else {
		cmd.Execute()
	}
}

//messageHandler is callback function how to handle incoming ws message
var messageHandler = func(dataByte []byte) {
	var msgType string
	msg := decodeWSMessage(string(dataByte))

	if v, ok := msg["type"]; ok {
		msgType = v
	}
	switch msgType {
	case model.TypeInfoResponse:
		logrus.Infof("Receive response type %v", model.TypeInfoResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.InfoResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeCreateResponse:
		logrus.Infof("Receive response type %v", model.TypeCreateResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.CreateResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeLeaveResponse:
		logrus.Infof("Receive response type %v", model.TypeLeaveResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.LeaveResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeLeaveNotif:
		logrus.Infof("Receive response type %v", model.TypeLeaveNotif)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.LeaveNotice)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeInviteResponse:
		logrus.Infof("Receive response type %v", model.TypeInviteResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.InviteResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeInviteNotif:
		logrus.Infof("Receive response type %v", model.TypeInviteNotif)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.InviteNotice)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeGetInvitatedNotif:
		logrus.Infof("Receive response type %v", model.TypeGetInvitatedNotif)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.Invitation)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeJoinResponse:
		logrus.Infof("Receive response type %v", model.TypeJoinResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.JoinResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeJoinNotif:
		logrus.Infof("Receive response type %v", model.TypeJoinNotif)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.JoinNotice)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeKickResponse:
		logrus.Infof("Receive response type %v", model.TypeKickResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.KickResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeKickNotice:
		logrus.Infof("Receive response type %v", model.TypeKickNotice)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.KickNotice)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeRejectResponse:
		logrus.Infof("Receive response type %v", model.TypeRejectResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.RejectResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeRejectNotif:
		logrus.Infof("Receive response type %v", model.TypeRejectNotif)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.RejectNotice)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypePromoteLeaderResponse:
		logrus.Infof("Receive response type %v", model.TypePromoteLeaderResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.PromoteLeaderResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypePersonalChatResponse:
		logrus.Infof("Receive response type %v", model.TypePersonalChatResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.PersonalChatResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypePersonalChatNotif:
		logrus.Infof("Receive response type %v", model.TypePersonalChatNotif)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.PersonalChatNotif)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypePartyChatResponse:
		logrus.Infof("Receive response type %v", model.TypePartyChatResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.PartyChatResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypePartyChatNotif:
		logrus.Infof("Receive response type %v", model.TypePartyChatNotif)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.PartyChatNotif)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeOnlineFriends:
		logrus.Infof("Receive response type %v", model.TypeOnlineFriends)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.OnlineFriendsResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeGetOfflineNotificationResponse:
		logrus.Infof("Receive response type %v", model.TypeGetOfflineNotificationResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.GetOfflineNotificationResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeFriendsPresenceResponse:
		logrus.Infof("Receive response type %v", model.TypeFriendsPresenceResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.ListFriendsPresenceResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeSetUserStatusResponse:
		logrus.Infof("Receive response type %v", model.TypeSetUserStatusResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.SetUserStatusResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeUserStatusNotif:
		logrus.Infof("Receive response type %v", model.TypeUserStatusNotif)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.FriendsPresenceNotif)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeRequestFriendsResponse:
		logrus.Infof("Receive response type %v", model.TypeRequestFriendsResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.RequestFriendsResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeRequestFriendsNotif:
		logrus.Infof("Receive response type %v", model.TypeRequestFriendsNotif)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.RequestFriendsNotif)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeListIncomingFriendsResponse:
		logrus.Infof("Receive response type %v", model.TypeListIncomingFriendsResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.ListIncomingFriendsResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeListOutgoingFriendsResponse:
		logrus.Infof("Receive response type %v", model.TypeListOutgoingFriendsResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.ListOutgoingFriendsResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeAcceptFriendsResponse:
		logrus.Infof("Receive response type %v", model.TypeAcceptFriendsResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.AcceptFriendsResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeAcceptFriendsNotif:
		logrus.Infof("Receive response type %v", model.TypeAcceptFriendsNotif)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.AcceptFriendsNotif)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeRejectFriendsResponse:
		logrus.Infof("Receive response type %v", model.TypeRejectFriendsResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.RejectFriendsResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeCancelFriendsResponse:
		logrus.Infof("Receive response type %v", model.TypeCancelFriendsResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.CancelFriendsResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeUnfriendResponse:
		logrus.Infof("Receive response type %v", model.TypeUnfriendResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.UnfriendResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeUnfriendNotif:
		logrus.Infof("Receive response type %v", model.TypeUnfriendNotif)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.UnfriendNotif)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeListOfFriendsResponse:
		logrus.Infof("Receive response type %v", model.TypeListOfFriendsResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.ListOfFriendsResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeGetFriendshipStatusResponse:
		logrus.Infof("Receive response type %v", model.TypeGetFriendshipStatusResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.GetFriendshipStatusResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypePersonalChatHistoryResponse:
		logrus.Infof("Receive response type %v", model.TypePersonalChatHistoryResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.PersonalChatHistoryResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeJoinDefaultChannelResponse:
		logrus.Infof("Receive response type %v", model.TypeJoinDefaultChannelResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.JoinDefaultChannelResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeSendChannelChatResponse:
		logrus.Infof("Receive response type %v", model.TypeSendChannelChatResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.SendChannelChatResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeChannelChatNotif:
		logrus.Infof("Receive response type %v", model.TypeChannelChatNotif)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.ChannelChatNotif)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeBlockPlayerResponse:
		logrus.Infof("Receive response type %v", model.TypeBlockPlayerResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.BlockPlayerResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeBlockPlayerNotif:
		logrus.Infof("Receive response type %v", model.TypeBlockPlayerNotif)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.PlayerBlockedNotification)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeUnblockPlayerResponse:
		logrus.Infof("Receive response type %v", model.TypeUnblockPlayerResponse)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.UnblockPlayerResponse)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	case model.TypeUnblockPlayerNotif:
		logrus.Infof("Receive response type %v", model.TypeUnblockPlayerNotif)
		unmarshal, err := parser.UnmarshalResponse(dataByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		data := unmarshal.(*model.PlayerUnblockedNotification)
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		logrus.Infof("Response content %v", string(marshal))
		break
	}
}

func printHelp() {
	fmt.Printf(`
Commands:
# Party
%s: User party info
%s: Create party
%s: Leave party
%s: Invite friend
%s: Join party
%s: Kick member
%s: Reject party invitation
%s: Promote new leader

# Chat
%s: 1to1 Chat
%s: Party Chat

# Presence
%s: Get online friends
%s: Get offline notifications
%s: Get Friends Status
%s: Set User Status

# Friends
%s: Request Friends
%s: List of Incoming Friends
%s: List of Outgoing Friends
%s: Accept Friends
%s: Reject Friends
%s: Cancel Friends Request
%s: Unfriend
%s: List of Friends
%s: Get Friendship Status
%s: Load Personal Chat History
%s: Join Default Chat Channel
%s: Send Channel Chat
%s: Block Player
%s: Unblock Player

# Lobby
%s: Quit

`,
		infoCmd,
		createCmd,
		leaveCmd,
		inviteCmd,
		joinCmd,
		kickCmd,
		rejectCmd,
		promoteLeaderCmd,
		chatCmd,
		partyChatCmd,
		getOnlineFriendsCmd,
		getOfflineNotificationCmd,
		getFriendsCmd,
		setUserStatusCmd,
		requestFriendsCmd,
		listIncomingFriendsCmd,
		listOutgoingFriendsCmd,
		acceptFriendsCmd,
		rejectFriendsCmd,
		cancelFriendsRequestCmd,
		unfriendCmd,
		listOfFriendsCmd,
		getFriendshipStatusCmd,
		personalChatHistoryCmd,
		joinDefaultChatChannelCmd,
		sendChatChannelCmd,
		blockCmd,
		unblockCmd,
		quitCmd)
}

func serve() {
	for {
		printHelp()
		command := getInput()
		switch command {

		// as long as ws connection in connMgr not closed,
		case infoCmd:
			err := partyService.GetPartyInfo()
			if err != nil {
				logrus.Error(err)
				return
			}
			break
		case createCmd:
			err := partyService.CreateParty()
			if err != nil {
				logrus.Error(err)
				return
			}
			break
		case leaveCmd:
			err := partyService.LeaveParty()
			if err != nil {
				logrus.Error(err)
				return
			}
			break
		case inviteCmd:
			fmt.Println("Friend ID:")
			friendID := getInput()
			logrus.Debug("Sending invite message")
			err := partyService.InviteParty(friendID)
			if err != nil {
				logrus.Error(err)
				return
			}
			break
		case joinCmd:
			fmt.Println("Party ID:")
			partyID := getInput()
			fmt.Println("Invitation token:")
			token := getInput()
			logrus.Debug("Sending join message")
			err := partyService.JoinParty(partyID, token)
			if err != nil {
				logrus.Error(err)
				return
			}
			break
		case kickCmd:
			fmt.Println("Member ID:")
			id := getInput()
			logrus.Debug("Sending kick message")
			err := partyService.KickPartyMember(id)
			if err != nil {
				logrus.Error(err)
				return
			}
			break
		case rejectCmd:
			fmt.Println("Party ID:")
			partyID := getInput()
			fmt.Println("Invitation token:")
			token := getInput()
			logrus.Debug("Sending reject message")
			err := partyService.RejectPartyInvitation(partyID, token)
			if err != nil {
				logrus.Error(err)
				return
			}
			break
		case promoteLeaderCmd:
			fmt.Println("New leader User ID:")
			leaderUserID := getInput()
			logrus.Debug("Sending kick message")
			err := partyService.PromotePartyLeader(leaderUserID)
			if err != nil {
				logrus.Error(err)
				return
			}
			break
		case chatCmd:
			fmt.Println("Friend ID:")
			friendID := getInput()
			fmt.Println("Message:")
			content := getInput()
			err := chatService.SendPersonalChat(friendID, content)
			if err != nil {
				logrus.Error(err)
			}
			break
		case partyChatCmd:
			fmt.Println("Message:")
			content := getInput()
			err := chatService.SendPartyChat(content)
			if err != nil {
				logrus.Error(err)
			}
			break
		case joinDefaultChatChannelCmd:
			err := chatService.JoinDefaultChannel()
			if err != nil {
				logrus.Error(err)
			}
			break
		case getOnlineFriendsCmd:
			err := friendService.GetOnlineFriends()
			if err != nil {
				logrus.Error(err)
			}
			break
		case getOfflineNotificationCmd:
			err := friendService.GetOfflineNotification()
			if err != nil {
				logrus.Error(err)
			}
			break
		case getFriendsCmd:
			err := friendService.GetFriendPresenceStatus()
			if err != nil {
				logrus.Error(err)
			}
			break
		case setUserStatusCmd:
			fmt.Println("Availability:")
			availability := getInput()
			fmt.Println("Activity:")
			activity := getInput()
			i, err := strconv.Atoi(availability)
			if err != nil {
				logrus.Error(err)
				return
			}
			err = friendService.SetUserStatus(i, activity)
			if err != nil {
				logrus.Error(err)
			}
			break
		case requestFriendsCmd:
			fmt.Println("Friend ID:")
			friendID := getInput()
			err := friendService.RequestFriend(friendID)
			if err != nil {
				logrus.Error(err)
			}
			break
		case listIncomingFriendsCmd:
			err := friendService.GetIncomingFriendRequest()
			if err != nil {
				logrus.Error(err)
			}
			break
		case listOutgoingFriendsCmd:
			err := friendService.GetOutgoingFriendRequest()
			if err != nil {
				logrus.Error(err)
			}
			break
		case acceptFriendsCmd:
			fmt.Println("Friend ID:")
			friendID := getInput()
			err := friendService.AcceptFriendRequest(friendID)
			if err != nil {
				logrus.Error(err)
			}
			break
		case rejectFriendsCmd:
			fmt.Println("Friend ID:")
			friendID := getInput()
			err := friendService.RejectFriendRequest(friendID)
			if err != nil {
				logrus.Error(err)
			}
			break
		case cancelFriendsRequestCmd:
			fmt.Println("Friend ID:")
			friendID := getInput()
			err := friendService.CancelFriendRequest(friendID)
			if err != nil {
				logrus.Error(err)
			}
			break
		case unfriendCmd:
			fmt.Println("Friend ID:")
			friendID := getInput()
			err := friendService.Unfriend(friendID)
			if err != nil {
				logrus.Error(err)
			}
			break
		case listOfFriendsCmd:
			err := friendService.GetFriends()
			if err != nil {
				logrus.Error(err)
			}
			break
		case getFriendshipStatusCmd:
			fmt.Println("Friend ID:")
			friendID := getInput()
			err := friendService.GetFriendshipStatus(friendID)
			if err != nil {
				logrus.Error(err)
			}
			break
		case personalChatHistoryCmd:
			fmt.Println("Friend ID:")
			friendID := getInput()
			err := chatService.GetPersonalChatHistory(friendID)
			if err != nil {
				logrus.Error(err)
			}
			break
		case sendChatChannelCmd:
			fmt.Println("Channel Slug:")
			channelSlug := getInput()
			fmt.Println("Content:")
			content := getInput()
			err := chatService.SendChannelChat(channelSlug, content)
			if err != nil {
				logrus.Error(err)
			}
			break
		case blockCmd:
			fmt.Println("Namespace:")
			namespace := getInput()
			fmt.Println("UserID:")
			userID := getInput()
			fmt.Println("Blocked user ID:")
			blockedUserID := getInput()
			err := friendService.Block(namespace, userID, blockedUserID)
			if err != nil {
				logrus.Error(err)
			}
			break
		case unblockCmd:
			fmt.Println("Namespace:")
			namespace := getInput()
			fmt.Println("UserID:")
			userID := getInput()
			fmt.Println("Blocked user ID:")
			blockedUserID := getInput()
			err := friendService.Unblock(namespace, userID, blockedUserID)
			if err != nil {
				logrus.Error(err)
			}
			break
		case quitCmd:
			logrus.Print("disconnect message: ")
			msg := getInput()
			_ = connMgr.Get().Conn.WriteMessage(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, msg))
			return
		}
	}
}

func getInput() string {
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	return text
}

func decodeWSMessage(msg string) map[string]string {
	scanner := bufio.NewScanner(strings.NewReader(msg))
	res := make(map[string]string)

	for scanner.Scan() {
		str := scanner.Text()
		keyValue := strings.Split(str, ": ")
		if len(keyValue) == 2 {
			res[keyValue[0]] = keyValue[1]
		}
	}

	if err := scanner.Err(); err != nil {
		logrus.Errorf("error reading websocket message: %v", err)
	}

	return res
}