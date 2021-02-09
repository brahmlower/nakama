// Copyright 2021 The Nakama Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

// #include <stdlib.h>
// #include "../include/nakama.h"
import "C"

import (
	"unsafe"
)

//export cInitializerRegisterRpc
func cInitializerRegisterRpc(
	pInit unsafe.Pointer,
	id C.NkString,
	pFn unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeRt
func cInitializerRegisterBeforeRt(
	pInit unsafe.Pointer,
	id C.NkString,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterRt
func cInitializerRegisterAfterRt(
	pInit unsafe.Pointer,
	id C.NkString,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterMatchmakerMatched
func cInitializerRegisterMatchmakerMatched(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterMatch
func cInitializerRegisterMatch(
	pInit unsafe.Pointer,
	name C.NkString,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterTournamentEnd
func cInitializerRegisterTournamentEnd(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterTournamentReset
func cInitializerRegisterTournamentReset(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterLeaderBoardEnd
func cInitializerRegisterLeaderBoardEnd(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterLeaderBoardReset
func cInitializerRegisterLeaderBoardReset(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeGetAccount
func cInitializerRegisterBeforeGetAccount(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterGetAccount
func cInitializerRegisterAfterGetAccount(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeUpdateAccount
func cInitializerRegisterBeforeUpdateAccount(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterUpdateAccount
func cInitializerRegisterAfterUpdateAccount(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeSessionRefresh
func cInitializerRegisterBeforeSessionRefresh(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterSessionRefresh
func cInitializerRegisterAfterSessionRefresh(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeAuthenticateApple
func cInitializerRegisterBeforeAuthenticateApple(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterAuthenticateApple
func cInitializerRegisterAfterAuthenticateApple(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeAuthenticateCustom
func cInitializerRegisterBeforeAuthenticateCustom(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterAuthenticateCustom
func cInitializerRegisterAfterAuthenticateCustom(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeAuthenticateDevice
func cInitializerRegisterBeforeAuthenticateDevice(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterAuthenticateDevice
func cInitializerRegisterAfterAuthenticateDevice(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeAuthenticateEmail
func cInitializerRegisterBeforeAuthenticateEmail(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterAuthenticateEmail
func cInitializerRegisterAfterAuthenticateEmail(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeAuthenticateFacebook
func cInitializerRegisterBeforeAuthenticateFacebook(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterAuthenticateFacebook
func cInitializerRegisterAfterAuthenticateFacebook(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeAuthenticateFacebookInstantGame
func cInitializerRegisterBeforeAuthenticateFacebookInstantGame(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterAuthenticateFacebookInstantGame
func cInitializerRegisterAfterAuthenticateFacebookInstantGame(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeAuthenticateGameCenter
func cInitializerRegisterBeforeAuthenticateGameCenter(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterAuthenticateGameCenter
func cInitializerRegisterAfterAuthenticateGameCenter(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeAuthenticateGoogle
func cInitializerRegisterBeforeAuthenticateGoogle(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterAuthenticateGoogle
func cInitializerRegisterAfterAuthenticateGoogle(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeAuthenticateSteam
func cInitializerRegisterBeforeAuthenticateSteam(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterAuthenticateSteam
func cInitializerRegisterAfterAuthenticateSteam(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeListChannelMessages
func cInitializerRegisterBeforeListChannelMessages(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterListChannelMessages
func cInitializerRegisterAfterListChannelMessages(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeListFriends
func cInitializerRegisterBeforeListFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterListFriends
func cInitializerRegisterAfterListFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeAddFriends
func cInitializerRegisterBeforeAddFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterAddFriends
func cInitializerRegisterAfterAddFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeDeleteFriends
func cInitializerRegisterBeforeDeleteFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterDeleteFriends
func cInitializerRegisterAfterDeleteFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeBlockFriends
func cInitializerRegisterBeforeBlockFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterBlockFriends
func cInitializerRegisterAfterBlockFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeImportFacebookFriends
func cInitializerRegisterBeforeImportFacebookFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterImportFacebookFriends
func cInitializerRegisterAfterImportFacebookFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeCreateGroup
func cInitializerRegisterBeforeCreateGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterCreateGroup
func cInitializerRegisterAfterCreateGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeUpdateGroup
func cInitializerRegisterBeforeUpdateGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterUpdateGroup
func cInitializerRegisterAfterUpdateGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeDeleteGroup
func cInitializerRegisterBeforeDeleteGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterDeleteGroup
func cInitializerRegisterAfterDeleteGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeJoinGroup
func cInitializerRegisterBeforeJoinGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterJoinGroup
func cInitializerRegisterAfterJoinGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeLeaveGroup
func cInitializerRegisterBeforeLeaveGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterLeaveGroup
func cInitializerRegisterAfterLeaveGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeAddGroupUsers
func cInitializerRegisterBeforeAddGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterAddGroupUsers
func cInitializerRegisterAfterAddGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeBanGroupUsers
func cInitializerRegisterBeforeBanGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterBanGroupUsers
func cInitializerRegisterAfterBanGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeKickGroupUsers
func cInitializerRegisterBeforeKickGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterKickGroupUsers
func cInitializerRegisterAfterKickGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforePromoteGroupUsers
func cInitializerRegisterBeforePromoteGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterPromoteGroupUsers
func cInitializerRegisterAfterPromoteGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeDemoteGroupUsers
func cInitializerRegisterBeforeDemoteGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterDemoteGroupUsers
func cInitializerRegisterAfterDemoteGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeListGroupUsers
func cInitializerRegisterBeforeListGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterListGroupUsers
func cInitializerRegisterAfterListGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeListUserGroups
func cInitializerRegisterBeforeListUserGroups(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterListUserGroups
func cInitializerRegisterAfterListUserGroups(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeListGroups
func cInitializerRegisterBeforeListGroups(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterListGroups
func cInitializerRegisterAfterListGroups(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeDeleteLeaderboardRecord
func cInitializerRegisterBeforeDeleteLeaderboardRecord(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterDeleteLeaderboardRecord
func cInitializerRegisterAfterDeleteLeaderboardRecord(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeListLeaderboardRecords
func cInitializerRegisterBeforeListLeaderboardRecords(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterListLeaderboardRecords
func cInitializerRegisterAfterListLeaderboardRecords(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeWriteLeaderboardRecord
func cInitializerRegisterBeforeWriteLeaderboardRecord(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterWriteLeaderboardRecord
func cInitializerRegisterAfterWriteLeaderboardRecord(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeListLeaderboardRecordsAroundOwner
func cInitializerRegisterBeforeListLeaderboardRecordsAroundOwner(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterListLeaderboardRecordsAroundOwner
func cInitializerRegisterAfterListLeaderboardRecordsAroundOwner(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeLinkApple
func cInitializerRegisterBeforeLinkApple(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterLinkApple
func cInitializerRegisterAfterLinkApple(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeLinkCustom
func cInitializerRegisterBeforeLinkCustom(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterLinkCustom
func cInitializerRegisterAfterLinkCustom(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeLinkDevice
func cInitializerRegisterBeforeLinkDevice(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterLinkDevice
func cInitializerRegisterAfterLinkDevice(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeLinkEmail
func cInitializerRegisterBeforeLinkEmail(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterLinkEmail
func cInitializerRegisterAfterLinkEmail(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeLinkFacebook
func cInitializerRegisterBeforeLinkFacebook(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterLinkFacebook
func cInitializerRegisterAfterLinkFacebook(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeLinkFacebookInstantGame
func cInitializerRegisterBeforeLinkFacebookInstantGame(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterLinkFacebookInstantGame
func cInitializerRegisterAfterLinkFacebookInstantGame(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeLinkGameCenter
func cInitializerRegisterBeforeLinkGameCenter(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterLinkGameCenter
func cInitializerRegisterAfterLinkGameCenter(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeLinkGoogle
func cInitializerRegisterBeforeLinkGoogle(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterLinkGoogle
func cInitializerRegisterAfterLinkGoogle(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeLinkSteam
func cInitializerRegisterBeforeLinkSteam(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterLinkSteam
func cInitializerRegisterAfterLinkSteam(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeListMatches
func cInitializerRegisterBeforeListMatches(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterListMatches
func cInitializerRegisterAfterListMatches(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeListNotifications
func cInitializerRegisterBeforeListNotifications(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterListNotifications
func cInitializerRegisterAfterListNotifications(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeDeleteNotifications
func cInitializerRegisterBeforeDeleteNotifications(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterDeleteNotifications
func cInitializerRegisterAfterDeleteNotifications(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeListStorageObjects
func cInitializerRegisterBeforeListStorageObjects(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterListStorageObjects
func cInitializerRegisterAfterListStorageObjects(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeReadStorageObjects
func cInitializerRegisterBeforeReadStorageObjects(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterReadStorageObjects
func cInitializerRegisterAfterReadStorageObjects(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeWriteStorageObjects
func cInitializerRegisterBeforeWriteStorageObjects(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterWriteStorageObjects
func cInitializerRegisterAfterWriteStorageObjects(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeDeleteStorageObjects
func cInitializerRegisterBeforeDeleteStorageObjects(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterDeleteStorageObjects
func cInitializerRegisterAfterDeleteStorageObjects(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeJoinTournament
func cInitializerRegisterBeforeJoinTournament(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterJoinTournament
func cInitializerRegisterAfterJoinTournament(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeListTournamentRecords
func cInitializerRegisterBeforeListTournamentRecords(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterListTournamentRecords
func cInitializerRegisterAfterListTournamentRecords(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeListTournaments
func cInitializerRegisterBeforeListTournaments(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterListTournaments
func cInitializerRegisterAfterListTournaments(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeWriteTournamentRecord
func cInitializerRegisterBeforeWriteTournamentRecord(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterWriteTournamentRecord
func cInitializerRegisterAfterWriteTournamentRecord(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeListTournamentRecordsAroundOwner
func cInitializerRegisterBeforeListTournamentRecordsAroundOwner(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterListTournamentRecordsAroundOwner
func cInitializerRegisterAfterListTournamentRecordsAroundOwner(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeUnlinkApple
func cInitializerRegisterBeforeUnlinkApple(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterUnlinkApple
func cInitializerRegisterAfterUnlinkApple(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeUnlinkCustom
func cInitializerRegisterBeforeUnlinkCustom(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterUnlinkCustom
func cInitializerRegisterAfterUnlinkCustom(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeUnlinkDevice
func cInitializerRegisterBeforeUnlinkDevice(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterUnlinkDevice
func cInitializerRegisterAfterUnlinkDevice(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeUnlinkEmail
func cInitializerRegisterBeforeUnlinkEmail(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterUnlinkEmail
func cInitializerRegisterAfterUnlinkEmail(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeUnlinkFacebook
func cInitializerRegisterBeforeUnlinkFacebook(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterUnlinkFacebook
func cInitializerRegisterAfterUnlinkFacebook(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeUnlinkFacebookInstantGame
func cInitializerRegisterBeforeUnlinkFacebookInstantGame(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterUnlinkFacebookInstantGame
func cInitializerRegisterAfterUnlinkFacebookInstantGame(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeUnlinkGameCenter
func cInitializerRegisterBeforeUnlinkGameCenter(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterUnlinkGameCenter
func cInitializerRegisterAfterUnlinkGameCenter(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeUnlinkGoogle
func cInitializerRegisterBeforeUnlinkGoogle(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterUnlinkGoogle
func cInitializerRegisterAfterUnlinkGoogle(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeUnlinkSteam
func cInitializerRegisterBeforeUnlinkSteam(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterUnlinkSteam
func cInitializerRegisterAfterUnlinkSteam(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterBeforeGetUsers
func cInitializerRegisterBeforeGetUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterAfterGetUsers
func cInitializerRegisterAfterGetUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterEvent
func cInitializerRegisterEvent(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterEventSessionStart
func cInitializerRegisterEventSessionStart(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}

//export cInitializerRegisterEventSessionEnd
func cInitializerRegisterEventSessionEnd(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char) int {
	return 0
}
