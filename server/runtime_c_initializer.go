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


//export cInitializerRegisterRpc
func cInitializerRegisterRpc(
	pInit unsafe.Pointer,
	id C.NkString,
	pFn unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterRpc(
		ptr,
		id,
		fn,
		outerror);
}

//export cInitializerRegisterBeforeRt
func cInitializerRegisterBeforeRt(
	pInit unsafe.Pointer,
	id C.NkString,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeRt(
		ptr,
		id,
		cb,
		outerror);
}

//export cInitializerRegisterAfterRt
func cInitializerRegisterAfterRt(
	pInit unsafe.Pointer,
	id C.NkString,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterRt(
		ptr,
		id,
		cb,
		outerror);
}

//export cInitializerRegisterMatchmakerMatched
func cInitializerRegisterMatchmakerMatched(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterMatchmakerMatched(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterMatch
func cInitializerRegisterMatch(
	pInit unsafe.Pointer,
	NkString name,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterMatch(
		ptr,
		name,
		cb,
		outerror);
}

//export 
func initializerregistertournamentend(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterTournamentEnd(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterTournamentReset
func cInitializerRegisterTournamentReset(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterTournamentReset(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterLeaderBoardEnd
func cInitializerRegisterLeaderBoardEnd(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterLeaderBoardEnd(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterLeaderBoardReset
func cInitializerRegisterLeaderBoardReset(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterLeaderBoardReset(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeGetAccount
func cInitializerRegisterBeforeGetAccount(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeGetAccount(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterGetAccount
func cInitializerRegisterAfterGetAccount(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterGetAccount(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeUpdateAccount
func cInitializerRegisterBeforeUpdateAccount(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeUpdateAccount(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterUpdateAccount
func cInitializerRegisterAfterUpdateAccount(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterUpdateAccount(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeSessionRefresh
func cInitializerRegisterBeforeSessionRefresh(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeSessionRefresh(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterSessionRefresh
func cInitializerRegisterAfterSessionRefresh(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterSessionRefresh(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeAuthenticateApple
func cInitializerRegisterBeforeAuthenticateApple(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeAuthenticateApple(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterAuthenticateApple
func cInitializerRegisterAfterAuthenticateApple(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterAuthenticateApple(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeAuthenticateCustom
func cInitializerRegisterBeforeAuthenticateCustom(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeAuthenticateCustom(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterAuthenticateCustom
func cInitializerRegisterAfterAuthenticateCustom(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterAuthenticateCustom(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeAuthenticateDevice
func cInitializerRegisterBeforeAuthenticateDevice(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeAuthenticateDevice(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterAuthenticateDevice
func cInitializerRegisterAfterAuthenticateDevice(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterAuthenticateDevice(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeAuthenticateEmail
func cInitializerRegisterBeforeAuthenticateEmail(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeAuthenticateEmail(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterAuthenticateEmail
func cInitializerRegisterAfterAuthenticateEmail(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterAuthenticateEmail(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeAuthenticateFacebook
func cInitializerRegisterBeforeAuthenticateFacebook(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeAuthenticateFacebook(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterAuthenticateFacebook
func cInitializerRegisterAfterAuthenticateFacebook(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterAuthenticateFacebook(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeAuthenticateFacebookInstantGame
func cInitializerRegisterBeforeAuthenticateFacebookInstantGame(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeAuthenticateFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterAuthenticateFacebookInstantGame
func cInitializerRegisterAfterAuthenticateFacebookInstantGame(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterAuthenticateFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeAuthenticateGameCenter
func cInitializerRegisterBeforeAuthenticateGameCenter(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeAuthenticateGameCenter(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterAuthenticateGameCenter
func cInitializerRegisterAfterAuthenticateGameCenter(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterAuthenticateGameCenter(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeAuthenticateGoogle
func cInitializerRegisterBeforeAuthenticateGoogle(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeAuthenticateGoogle(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterAuthenticateGoogle
func cInitializerRegisterAfterAuthenticateGoogle(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterAuthenticateGoogle(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeAuthenticateSteam
func cInitializerRegisterBeforeAuthenticateSteam(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeAuthenticateSteam(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterAuthenticateSteam
func cInitializerRegisterAfterAuthenticateSteam(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterAuthenticateSteam(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeListChannelMessages
func cInitializerRegisterBeforeListChannelMessages(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeListChannelMessages(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterListChannelMessages
func cInitializerRegisterAfterListChannelMessages(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterListChannelMessages(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeListFriends
func cInitializerRegisterBeforeListFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeListFriends(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterListFriends
func cInitializerRegisterAfterListFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterListFriends(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeAddFriends
func cInitializerRegisterBeforeAddFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeAddFriends(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterAddFriends
func cInitializerRegisterAfterAddFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterAddFriends(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeDeleteFriends
func cInitializerRegisterBeforeDeleteFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeDeleteFriends(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterDeleteFriends
func cInitializerRegisterAfterDeleteFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterDeleteFriends(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeBlockFriends
func cInitializerRegisterBeforeBlockFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeBlockFriends(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterBlockFriends
func cInitializerRegisterAfterBlockFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterBlockFriends(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeImportFacebookFriends
func cInitializerRegisterBeforeImportFacebookFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeImportFacebookFriends(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterImportFacebookFriends
func cInitializerRegisterAfterImportFacebookFriends(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterImportFacebookFriends(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeCreateGroup
func cInitializerRegisterBeforeCreateGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeCreateGroup(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterCreateGroup
func cInitializerRegisterAfterCreateGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterCreateGroup(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeUpdateGroup
func cInitializerRegisterBeforeUpdateGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeUpdateGroup(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterUpdateGroup
func cInitializerRegisterAfterUpdateGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterUpdateGroup(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeDeleteGroup
func cInitializerRegisterBeforeDeleteGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeDeleteGroup(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterDeleteGroup
func cInitializerRegisterAfterDeleteGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterDeleteGroup(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeJoinGroup
func cInitializerRegisterBeforeJoinGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeJoinGroup(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterJoinGroup
func cInitializerRegisterAfterJoinGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterJoinGroup(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeLeaveGroup
func cInitializerRegisterBeforeLeaveGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeLeaveGroup(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterLeaveGroup
func cInitializerRegisterAfterLeaveGroup(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterLeaveGroup(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeAddGroupUsers
func cInitializerRegisterBeforeAddGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeAddGroupUsers(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterAddGroupUsers
func cInitializerRegisterAfterAddGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterAddGroupUsers(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeBanGroupUsers
func cInitializerRegisterBeforeBanGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeBanGroupUsers(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterBanGroupUsers
func cInitializerRegisterAfterBanGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterBanGroupUsers(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeKickGroupUsers
func cInitializerRegisterBeforeKickGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeKickGroupUsers(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterKickGroupUsers
func cInitializerRegisterAfterKickGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterKickGroupUsers(
		ptr,
		cb,
		outerror);
}

// export cInitializerRegisterBeforePromoteGroupUsers
func cInitializerRegisterBeforePromoteGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforePromoteGroupUsers(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterPromoteGroupUsers
func cInitializerRegisterAfterPromoteGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterPromoteGroupUsers(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeDemoteGroupUsers
func cInitializerRegisterBeforeDemoteGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeDemoteGroupUsers(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterDemoteGroupUsers
func cInitializerRegisterAfterDemoteGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterDemoteGroupUsers(
		ptr,
		cb,
		outerror);
}


//export cInitializerRegisterBeforeListGroupUsers
func cInitializerRegisterBeforeListGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeListGroupUsers(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterListGroupUsers
func cInitializerRegisterAfterListGroupUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterListGroupUsers(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeListUserGroups
func cInitializerRegisterBeforeListUserGroups(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeListUserGroups(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterListUserGroups
func cInitializerRegisterAfterListUserGroups(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterListUserGroups(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeListGroups
func cInitializerRegisterBeforeListGroups(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeListGroups(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterListGroups
func cInitializerRegisterAfterListGroups(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterListGroups(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeDeleteLeaderboardRecord
func cInitializerRegisterBeforeDeleteLeaderboardRecord(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeDeleteLeaderboardRecord(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterDeleteLeaderboardRecord
func cInitializerRegisterAfterDeleteLeaderboardRecord(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterDeleteLeaderboardRecord(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeListLeaderboardRecords
func cInitializerRegisterBeforeListLeaderboardRecords(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeListLeaderboardRecords(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterListLeaderboardRecords
func cInitializerRegisterAfterListLeaderboardRecords(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterListLeaderboardRecords(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeWriteLeaderboardRecord
func cInitializerRegisterBeforeWriteLeaderboardRecord(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeWriteLeaderboardRecord(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterWriteLeaderboardRecord
func cInitializerRegisterAfterWriteLeaderboardRecord(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterWriteLeaderboardRecord(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeListLeaderboardRecordsAroundOwner
func cInitializerRegisterBeforeListLeaderboardRecordsAroundOwner(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeListLeaderboardRecordsAroundOwner(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterListLeaderboardRecordsAroundOwner
func cInitializerRegisterAfterListLeaderboardRecordsAroundOwner(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterListLeaderboardRecordsAroundOwner(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeLinkApple
func cInitializerRegisterBeforeLinkApple(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeLinkApple(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterLinkApple
func cInitializerRegisterAfterLinkApple(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterLinkApple(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeLinkCustom
func cInitializerRegisterBeforeLinkCustom(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeLinkCustom(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterLinkCustom
func cInitializerRegisterAfterLinkCustom(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterLinkCustom(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeLinkDevice
func cInitializerRegisterBeforeLinkDevice(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeLinkDevice(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterLinkDevice
func cInitializerRegisterAfterLinkDevice(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterLinkDevice(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeLinkEmail
func cInitializerRegisterBeforeLinkEmail(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeLinkEmail(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterLinkEmail
func cInitializerRegisterAfterLinkEmail(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterLinkEmail(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeLinkFacebook
func cInitializerRegisterBeforeLinkFacebook(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeLinkFacebook(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterLinkFacebook
func cInitializerRegisterAfterLinkFacebook(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterLinkFacebook(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeLinkFacebookInstantGame
func cInitializerRegisterBeforeLinkFacebookInstantGame(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeLinkFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterLinkFacebookInstantGame
func cInitializerRegisterAfterLinkFacebookInstantGame(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterLinkFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeLinkGameCenter
func cInitializerRegisterBeforeLinkGameCenter(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeLinkGameCenter(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterLinkGameCenter
func cInitializerRegisterAfterLinkGameCenter(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterLinkGameCenter(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeLinkGoogle
func cInitializerRegisterBeforeLinkGoogle(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeLinkGoogle(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterLinkGoogle
func cInitializerRegisterAfterLinkGoogle(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterLinkGoogle(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeLinkSteam
func cInitializerRegisterBeforeLinkSteam(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeLinkSteam(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterLinkSteam
func cInitializerRegisterAfterLinkSteam(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterLinkSteam(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeListMatches
func cInitializerRegisterBeforeListMatches(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeListMatches(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterListMatches
func cInitializerRegisterAfterListMatches(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterListMatches(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeListNotifications
func cInitializerRegisterBeforeListNotifications(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeListNotifications(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterListNotifications
func cInitializerRegisterAfterListNotifications(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterListNotifications(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeDeleteNotifications
func cInitializerRegisterBeforeDeleteNotifications(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeDeleteNotifications(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterDeleteNotifications
func cInitializerRegisterAfterDeleteNotifications(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterDeleteNotifications(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeListStorageObjects
func cInitializerRegisterBeforeListStorageObjects(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeListStorageObjects(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterListStorageObjects
func cInitializerRegisterAfterListStorageObjects(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterListStorageObjects(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeReadStorageObjects
func cInitializerRegisterBeforeReadStorageObjects(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeReadStorageObjects(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterReadStorageObjects
func cInitializerRegisterAfterReadStorageObjects(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterReadStorageObjects(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeWriteStorageObjects
func cInitializerRegisterBeforeWriteStorageObjects(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeWriteStorageObjects(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterWriteStorageObjects
func cInitializerRegisterAfterWriteStorageObjects(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterWriteStorageObjects(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeDeleteStorageObjects
func cInitializerRegisterBeforeDeleteStorageObjects(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeDeleteStorageObjects(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterDeleteStorageObjects
func cInitializerRegisterAfterDeleteStorageObjects(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterDeleteStorageObjects(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeJoinTournament
func cInitializerRegisterBeforeJoinTournament(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeJoinTournament(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterJoinTournament
func cInitializerRegisterAfterJoinTournament(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterJoinTournament(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeListTournamentRecords
func cInitializerRegisterBeforeListTournamentRecords(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeListTournamentRecords(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterListTournamentRecords
func cInitializerRegisterAfterListTournamentRecords(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterListTournamentRecords(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeListTournaments
func cInitializerRegisterBeforeListTournaments(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeListTournaments(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterListTournaments
func cInitializerRegisterAfterListTournaments(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterListTournaments(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeWriteTournamentRecord
func cInitializerRegisterBeforeWriteTournamentRecord(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeWriteTournamentRecord(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterWriteTournamentRecord
func cInitializerRegisterAfterWriteTournamentRecord(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterWriteTournamentRecord(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeListTournamentRecordsAroundOwner
func cInitializerRegisterBeforeListTournamentRecordsAroundOwner(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeListTournamentRecordsAroundOwner(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterListTournamentRecordsAroundOwner
func cInitializerRegisterAfterListTournamentRecordsAroundOwner(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterListTournamentRecordsAroundOwner(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeUnlinkApple
func cInitializerRegisterBeforeUnlinkApple(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeUnlinkApple(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterUnlinkApple
func cInitializerRegisterAfterUnlinkApple(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterUnlinkApple(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeUnlinkCustom
func cInitializerRegisterBeforeUnlinkCustom(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeUnlinkCustom(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterUnlinkCustom
func cInitializerRegisterAfterUnlinkCustom(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterUnlinkCustom(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeUnlinkDevice
func cInitializerRegisterBeforeUnlinkDevice(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeUnlinkDevice(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterUnlinkDevice
func cInitializerRegisterAfterUnlinkDevice(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterUnlinkDevice(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeUnlinkEmail
func cInitializerRegisterBeforeUnlinkEmail(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeUnlinkEmail(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterUnlinkEmail
func cInitializerRegisterAfterUnlinkEmail(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterUnlinkEmail(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeUnlinkFacebook
func cInitializerRegisterBeforeUnlinkFacebook(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeUnlinkFacebook(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterUnlinkFacebook
func cInitializerRegisterAfterUnlinkFacebook(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterUnlinkFacebook(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeUnlinkFacebookInstantGame
func cInitializerRegisterBeforeUnlinkFacebookInstantGame(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeUnlinkFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterUnlinkFacebookInstantGame
func cInitializerRegisterAfterUnlinkFacebookInstantGame(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterUnlinkFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeUnlinkGameCenter
func cInitializerRegisterBeforeUnlinkGameCenter(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeUnlinkGameCenter(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterUnlinkGameCenter
func cInitializerRegisterAfterUnlinkGameCenter(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterUnlinkGameCenter(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeUnlinkGoogle
func cInitializerRegisterBeforeUnlinkGoogle(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeUnlinkGoogle(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterUnlinkGoogle
func cInitializerRegisterAfterUnlinkGoogle(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterUnlinkGoogle(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeUnlinkSteam
func cInitializerRegisterBeforeUnlinkSteam(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeUnlinkSteam(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterUnlinkSteam
func cInitializerRegisterAfterUnlinkSteam(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterUnlinkSteam(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterBeforeGetUsers
func cInitializerRegisterBeforeGetUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterBeforeGetUsers(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterAfterGetUsers
func cInitializerRegisterAfterGetUsers(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterAfterGetUsers(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterEvent
func cInitializerRegisterEvent(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterEvent(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterEventSessionStart
func cInitializerRegisterEventSessionStart(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterEventSessionStart(
		ptr,
		cb,
		outerror);
}

//export cInitializerRegisterEventSessionEnd
func cInitializerRegisterEventSessionEnd(
	pInit unsafe.Pointer,
	pCB unsafe.Pointer,
	outError **C.char)
{
	return cInitializerRegisterEventSessionEnd(
		ptr,
		cb,
		outerror);
}