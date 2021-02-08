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

// #include "../include/nakama.h"
import "C"

import (
	"context"
	"unsafe"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/mattn/go-pointer"
)

type RuntimeCNakamaModuleCall struct {
	nk     runtime.NakamaModule
	allocs []unsafe.Pointer
}

func (c *RuntimeCNakamaModuleCall) goStringNk(s string) *C.NkString {
	var ret *C.NkString
	ret.p = C.CString(s)
	ret.n = C.long(len(s))

	c.allocs = append(c.allocs, unsafe.Pointer(ret.p))

	return ret
}

//export cModuleAuthenticateApple
func cModuleAuthenticateApple(pNk unsafe.Pointer, pCtx unsafe.Pointer, token, userName C.NkString, create bool, outUserID, outUserName, outErr **C.char, outCreated **bool) int {
	// call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	// ctx := pointer.Restore(pCtx).(context.Context)
	// retUserID, retUserName, retCreated, err := call.nk.AuthenticateApple(ctx, nkStringGo(token), nkStringGo(userName), create)
	// outCreated = &retCreated
	// outUserID = call.goStringNk(retUserID)
	// outUserName = call.goStringNk(retUserName)

	// if err != nil {
	// 	outErr = call.goStringNk(err.Error())

	// 	return 1
	// }

	return 0
}

//export cModuleAuthenticateCustom
func cModuleAuthenticateCustom(pNk unsafe.Pointer, pCtx unsafe.Pointer, userID, userName C.NkString, create bool, outUserID, outUserName, outErr **C.char, outCreated **bool) int {
	// call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	// ctx := pointer.Restore(pCtx).(context.Context)
	// retUserID, retUserName, retCreated, err := call.nk.AuthenticateCustom(ctx, nkStringGo(userID), nkStringGo(userName), create)
	// outCreated = &retCreated
	// outUserID = call.goStringNk(retUserID)
	// outUserName = call.goStringNk(retUserName)

	// if err != nil {
	// 	outErr = call.goStringNk(err.Error())

	// 	return 1
	// }

	return 0
}

//export cModuleAuthenticateDevice
func cModuleAuthenticateDevice(pNk unsafe.Pointer, pCtx unsafe.Pointer, userID, userName C.NkString, create bool, outUserID, outUserName, outErr **C.char, outCreated **bool) int {
	// call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	// ctx := pointer.Restore(pCtx).(context.Context)
	// retUserID, retUserName, retCreated, err := call.nk.AuthenticateDevice(ctx, nkStringGo(userID), nkStringGo(userName), create)
	// outCreated = &retCreated
	// outUserID = call.goStringNk(retUserID)
	// outUserName = call.goStringNk(retUserName)

	// if err != nil {
	// 	outErr = call.goStringNk(err.Error())

	// 	return 1
	// }

	return 0
}

//export cModuleAuthenticateEmail
func cModuleAuthenticateEmail(pNk unsafe.Pointer, pCtx unsafe.Pointer, email, password, userName C.NkString, create bool, outUserID, outUserName, outErr **C.char, outCreated ***bool) int {
	// call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	// ctx := pointer.Restore(pCtx).(context.Context)
	// retUserID, retUserName, retCreated, err := call.nk.AuthenticateEmail(ctx, nkStringGo(email), nkStringGo(password), nkStringGo(userName), create)
	// outCreated = &retCreated
	// outUserID = call.goStringNk(retUserID)
	// outUserName = call.goStringNk(retUserName)

	// if err != nil {
	// 	outErr = call.goStringNk(err.Error())

	// 	return 1
	// }

	return 0
}

//export cModuleAuthenticateFacebook
func cModuleAuthenticateFacebook(pNk unsafe.Pointer, pCtx unsafe.Pointer, token C.NkString, importFriends bool, userName C.NkString, create bool, outUserID, outUserName, outErr **C.char, outCreated **bool) int {
	// call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	// ctx := pointer.Restore(pCtx).(context.Context)
	// retUserID, retUserName, retCreated, err := call.nk.AuthenticateFacebook(ctx, nkStringGo(token), importFriends, nkStringGo(userName), create)
	// outCreated = &retCreated
	// outUserID = call.goStringNk(retUserID)
	// outUserName = call.goStringNk(retUserName)

	// if err != nil {
	// 	outErr = call.goStringNk(err.Error())

	// 	return 1
	// }

	return 0
}

//export cModuleAuthenticateFacebookInstantGame
func cModuleAuthenticateFacebookInstantGame(pNk unsafe.Pointer, pCtx unsafe.Pointer, userID, userName C.NkString, create bool, outUserID, outUserName, outErr **C.char, outCreated **bool) int {
	// call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	// ctx := pointer.Restore(pCtx).(context.Context)
	// retUserID, retUserName, retCreated, err := call.nk.AuthenticateFacebookInstantGame(ctx, nkStringGo(userID), nkStringGo(userName), create)
	// outCreated = &retCreated
	// outUserID = call.goStringNk(retUserID)
	// outUserName = call.goStringNk(retUserName)

	// if err != nil {
	// 	outErr = call.goStringNk(err.Error())

	// 	return 1
	// }

	return 0
}

//export cModuleAuthenticateGameCenter
func cModuleAuthenticateGameCenter(pNk unsafe.Pointer, pCtx unsafe.Pointer, userID, bundleID C.NkString, timestamp int64, salt, signature, publicKeyURL, userName C.NkString, create bool, outUserID, outUserName, outErr **C.char, outCreated **bool) int {
	// call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	// ctx := pointer.Restore(pCtx).(context.Context)
	// retUserID, retUserName, retCreated, err := call.nk.AuthenticateGameCenter(ctx, nkStringGo(userID), nkStringGo(bundleID), timestamp, nkStringGo(salt), nkStringGo(signature), nkStringGo(publicKeyURL), nkStringGo(userName), create)
	// outCreated = &retCreated
	// outUserID = call.goStringNk(retUserID)
	// outUserName = call.goStringNk(retUserName)

	// if err != nil {
	// 	outErr = call.goStringNk(err.Error())

	// 	return 1
	// }

	return 0
}

//export cModuleAuthenticateGoogle
func cModuleAuthenticateGoogle(pNk unsafe.Pointer, pCtx unsafe.Pointer, userID, userName C.NkString, create bool, outUserID, outUserName, outErr **C.char, outCreated **bool) int {
	// call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	// ctx := pointer.Restore(pCtx).(context.Context)
	// retUserID, retUserName, retCreated, err := call.nk.AuthenticateGoogle(ctx, nkStringGo(userID), nkStringGo(userName), create)
	// outCreated = &retCreated
	// outUserID = call.goStringNk(retUserID)
	// outUserName = call.goStringNk(retUserName)

	// if err != nil {
	// 	outErr = call.goStringNk(err.Error())

	// 	return 1
	// }

	return 0
}

//export cModuleAuthenticateSteam
func cModuleAuthenticateSteam(pNk unsafe.Pointer, pCtx unsafe.Pointer, userID, userName C.NkString, create bool, outUserID, outUserName, outErr **C.char, outCreated **bool) int {
	// call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	// ctx := pointer.Restore(pCtx).(context.Context)
	// retUserID, retUserName, retCreated, err := call.nk.AuthenticateSteam(ctx, nkStringGo(userID), nkStringGo(userName), create)
	// outCreated = &retCreated
	// outUserID = call.goStringNk(retUserID)
	// outUserName = call.goStringNk(retUserName)

	// if err != nil {
	// 	outErr = call.goStringNk(err.Error())

	// 	return 1
	// }

	return 0
}

//export cModuleAuthenticateTokenGenerate
func cModuleAuthenticateTokenGenerate(
	pNk unsafe.Pointer,
	userID C.NkString,
	userName C.NkString,
	expiry C.NkI64,
	vars C.NkMapString,
	outToken **C.char,
	outExpiry **C.NkI64,
	outError **C.char) int {
	return 0
}

//export cModuleAccountGetId
func cModuleAccountGetId(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	outAccount **C.NkAccount,
	outError **C.char) int {
	return 0
}

//export cModuleAccountsGetId
func cModuleAccountsGetId(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userIDs *C.NkString,
	numUserIDSs C.NkU32,
	outAccounts **C.NkAccount,
	outNumAccounts **C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleAccountUpdateId
func cModuleAccountUpdateId(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	userName C.NkString,
	metadata C.NkMapAny,
	displayName C.NkString,
	timeZone C.NkString,
	location C.NkString,
	langTag C.NkString,
	avatarURL C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleAccountDeleteId
func cModuleAccountDeleteId(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	recorded bool,
	outError **C.char) int {
	return 0
}

//export cModuleAccountExportId
func cModuleAccountExportId(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	outAccount **C.char,
	outError **C.char) int {
	return 0
}

//export cModuleUsersGetId
func cModuleUsersGetId(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	keys *C.NkString,
	numKeys C.NkU32,
	outUsers **C.NkUser,
	outNumUsers **C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleUsersGetUsername
func cModuleUsersGetUsername(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	keys *C.NkString,
	numKeys C.NkU32,
	outUsers **C.NkUser,
	outNumUsers **C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleUsersBanId
func cModuleUsersBanId(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userIDs *C.NkString,
	numUserIDs C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleUsersUnbanId
func cModuleUsersUnbanId(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userIDs *C.NkString,
	numUserIDs C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleLinkApple
func cModuleLinkApple(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleLinkCustom
func cModuleLinkCustom(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleLinkDevice
func cModuleLinkDevice(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleLinkFacebookInstantGame
func cModuleLinkFacebookInstantGame(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleLinkGoogle
func cModuleLinkGoogle(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleLinkSteam
func cModuleLinkSteam(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleUnlinkApple
func cModuleUnlinkApple(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleUnlinkCustom
func cModuleUnlinkCustom(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleUnlinkDevice
func cModuleUnlinkDevice(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleUnlinkEmail
func cModuleUnlinkEmail(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleUnlinkFacebook
func cModuleUnlinkFacebook(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleUnlinkFacebookInstantGame
func cModuleUnlinkFacebookInstantGame(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleUnlinkGameCenter
func cModuleUnlinkGameCenter(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	playerID C.NkString,
	bundleID C.NkString,
	timestamp C.NkI64,
	salt C.NkString,
	signature C.NkString,
	publicKeyURL C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleUnlinkGoogle
func cModuleUnlinkGoogle(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleUnlinkSteam
func cModuleUnlinkSteam(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outError **C.char) int {
	return 0
}

//export
func cModuleLinkEmail(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	email C.NkString,
	password C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleLinkFacebook
func cModuleLinkFacebook(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	userName C.NkString,
	token C.NkString,
	importFriends bool,
	outError **C.char) int {
	return 0
}

//export cModuleLinkGameCenter
func cModuleLinkGameCenter(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	playerID C.NkString,
	bundleID C.NkString,
	timestamp C.NkI64,
	salt C.NkString,
	signature C.NkString,
	publicKeyURL C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleStreamUserList
func cModuleStreamUserList(
	pNk unsafe.Pointer,
	mode C.NkU8,
	subject C.NkString,
	subContext C.NkString,
	label C.NkString,
	includeHidden bool,
	includeNotHidden bool,
	outPresences **C.NkPresence,
	outNumPresences **C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleStreamUserGet
func cModuleStreamUserGet(
	pNk unsafe.Pointer,
	mode C.NkU8,
	subject C.NkString,
	subContext C.NkString,
	label C.NkString,
	userID C.NkString,
	sessionID C.NkString,
	outMeta **C.NkPresenceMeta,
	outError **C.char) int {
	return 0
}

//export cModuleStreamUserJoin
func cModuleStreamUserJoin(
	pNk unsafe.Pointer,
	mode C.NkU8,
	subject C.NkString,
	subContext C.NkString,
	label C.NkString,
	userID C.NkString,
	sessionID C.NkString,
	hidden bool,
	persistence bool,
	status C.NkString,
	outJoined **bool,
	outError **C.char) int {
	return 0
}

//export cModuleStreamUserUpdate
func cModuleStreamUserUpdate(
	pNk unsafe.Pointer,
	mode C.NkU8,
	subject C.NkString,
	subContext C.NkString,
	label C.NkString,
	userID C.NkString,
	sessionID C.NkString,
	hidden bool,
	persistence bool,
	status C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleStreamUserLeave
func cModuleStreamUserLeave(
	pNk unsafe.Pointer,
	mode C.NkU8,
	subject C.NkString,
	subContext C.NkString,
	label C.NkString,
	userID C.NkString,
	sessionID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleStreamUserKick
func cModuleStreamUserKick(
	pNk unsafe.Pointer,
	mode C.NkU8,
	subject C.NkString,
	subContext C.NkString,
	label C.NkString,
	presence C.NkPresence,
	outError **C.char) int {
	return 0
}

//export cModuleStreamCount
func cModuleStreamCount(
	pNk unsafe.Pointer,
	mode C.NkU8,
	subject C.NkString,
	subContext C.NkString,
	label C.NkString,
	outCount **C.NkU64,
	outError **C.char) int {
	return 0
}

//export
func cModuleStreamClose(
	pNk unsafe.Pointer,
	mode C.NkU8,
	subject C.NkString,
	subContext C.NkString,
	label C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleStreamSend
func cModuleStreamSend(
	pNk unsafe.Pointer,
	mode C.NkU8,
	subject C.NkString,
	subContext C.NkString,
	label C.NkString,
	data C.NkString,
	presences *C.NkPresence,
	numPresences C.NkU32,
	reliable bool,
	outError **C.char) int {
	return 0
}

//export cModuleStreamSendRaw
func cModuleStreamSendRaw(
	pNk unsafe.Pointer,
	mode C.NkU8,
	subject C.NkString,
	subContext C.NkString,
	label C.NkString,
	msg C.NkEnvelope,
	presences *C.NkPresence,
	numPresences C.NkU32,
	reliable bool,
	outError **C.char) int {
	return 0
}

//export cModuleSessionDisconnect
func cModuleSessionDisconnect(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	sessionID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleMatchCreate
func cModuleMatchCreate(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	module C.NkString,
	params C.NkMapAny,
	outMatchId **C.char,
	outError **C.char) int {
	return 0
}

//export cModuleMatchGet
func cModuleMatchGet(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	outMatch **C.NkMatch,
	outError **C.char) int {
	return 0
}

//export cModuleMatchList
func cModuleMatchList(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	limit C.NkU32,
	authoritative bool,
	label C.NkString,
	minSize *C.NkU32,
	maxSize *C.NkU32,
	query C.NkString,
	outmatches **C.NkMatch,
	outNumMatches **C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleNotificationSend
func cModuleNotificationSend(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	subject C.NkString,
	content C.NkMapAny,
	code C.NkU64,
	sender C.NkString,
	persistent bool,
	outError **C.char) int {
	return 0
}

//export cModuleNotificationsSend
func cModuleNotificationsSend(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	notifications *C.NkNotificationSend,
	numNotifications C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleWalletUpdate
func cModuleWalletUpdate(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	changeset C.NkMapI64,
	metadata C.NkMapAny,
	updateLedger bool,
	outUpdated **C.NkMapI64,
	outPrevious **C.NkMapI64,
	outError **C.char) int {
	return 0
}

//export cModuleWalletsUpdate
func cModuleWalletsUpdate(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	updates *C.NkWalletUpdate,
	numUpdates C.NkU32,
	updateLedger bool,
	outResults **C.NkWalletUpdateResult,
	outNumResults **C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleWalletLedgerUpdate
func cModuleWalletLedgerUpdate(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	itemID C.NkString,
	metadata C.NkMapAny,
	outItem **C.NkWalletLedgerItem,
	outError **C.char) int {
	return 0
}

//export cModuleWalletLedgerList
func cModuleWalletLedgerList(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	limit C.NkU32,
	cursor C.NkString,
	outItems **C.NkWalletLedgerItem,
	outNumItems **C.NkU32,
	outCursor **C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleStorageList
func cModuleStorageList(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	collection C.NkString,
	limit C.NkU32,
	cursor C.NkString,
	outobjs **C.NkStorageObject,
	outNumObjs **C.NkU32,
	outCursor **C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleStorageRead
func cModuleStorageRead(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	reads *C.NkStorageRead,
	numReads C.NkU32,
	outObjs **C.NkStorageObject,
	outNumObjs **C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleStorageWrite
func cModuleStorageWrite(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	writes *C.NkStorageWrite,
	numWrites C.NkU32,
	outAcks **C.NkStorageObjectAck,
	outNumAcks **C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleStorageDelete
func cModuleStorageDelete(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	deletes *C.NkStorageDelete,
	numDeletes C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleMultiUpdate
func cModuleMultiUpdate(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	accountUpdates *C.NkAccountUpdate,
	numAccountUpdates C.NkU32,
	storageWrites *C.NkStorageWrite,
	numStorageWrites C.NkU32,
	walletUpdates *C.NkWalletUpdate,
	numWalletUpdates C.NkU32,
	updateLedger bool,
	outAcks **C.NkStorageObjectAck,
	outNumAcks **C.NkU32,
	outResults **C.NkWalletUpdateResult,
	outNumResults **C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleLeaderboardCreate
func cModuleLeaderboardCreate(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	authoritative bool,
	sortOrder C.NkString,
	op C.NkString,
	resetSchedule C.NkString,
	metadata C.NkMapAny,
	outError **C.char) int {
	return 0
}

//export cModuleLeaderboardRecordsList
func cModuleLeaderboardRecordsList(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	ownerIDs *C.NkString,
	numOwnerIDs C.NkU32,
	limit C.NkU32,
	cursor C.NkString,
	expiry C.NkI64,
	outRecords **C.NkLeaderboardRecord,
	outNumRecords **C.NkU32,
	outOwnerRecords **C.NkLeaderboardRecord,
	outNumOwnerRecords **C.NkU32,
	outNextCursor **C.NkString,
	outPrevCursor **C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleLeaderboardRecordWrite
func cModuleLeaderboardRecordWrite(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	ownerID C.NkString,
	score C.NkI64,
	subscore C.NkI64,
	metadata C.NkMapAny,
	outRecord **C.NkLeaderboardRecord,
	outError **C.char) int {
	return 0
}

//export cModuleLeaderboardDelete
func cModuleLeaderboardDelete(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	ownerID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleLeaderboardRecordDelete
func cModuleLeaderboardRecordDelete(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	ownerID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleTournamentDelete
func cModuleTournamentDelete(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	ownerID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleGroupDelete
func cModuleGroupDelete(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	ownerID C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleTournamentCreate
func cModuleTournamentCreate(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	sortOrder C.NkString,
	resetSchedule C.NkString,
	metadata C.NkMapAny,
	title C.NkString,
	description C.NkString,
	category C.NkU32,
	startTime C.NkU32,
	endTime C.NkU32,
	duration C.NkU32,
	maxSize C.NkU32,
	maxNumScore C.NkU32,
	joinRequired bool,
	outError **C.char) int {
	return 0
}

//export cModuleTournamentAddAttempt
func cModuleTournamentAddAttempt(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	ownerID C.NkString,
	count C.NkU64,
	outError **C.char) int {
	return 0
}

//export cModuleTournamentJoin
func cModuleTournamentJoin(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	ownerID C.NkString,
	userName C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleTournamentsGetId
func cModuleTournamentsGetId(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	tournamentIDs *C.NkString,
	numTournamentIDs C.NkU32,
	outTournaments **C.NkTournament,
	outNumTournaments **C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleTournamentList
func cModuleTournamentList(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	catStart C.NkU64,
	catEnd C.NkU64,
	startTime C.NkU64,
	endTime C.NkU64,
	limit C.NkU32,
	cursor C.NkString,
	id C.NkString,
	outTournaments **C.NkTournamentList,
	outNumTournaments **C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleTournamentRecordsList
func cModuleTournamentRecordsList(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	tournamentID C.NkString,
	ownerIDs *C.NkString,
	numOwnerIDs C.NkU32,
	limit C.NkU32,
	cursor C.NkString,
	overrideExpiry C.NkI64,
	outRecords **C.NkLeaderboardRecord,
	outNumRecords **C.NkU32,
	outOwnerRecords **C.NkLeaderboardRecord,
	outNumOwnerRecords **C.NkU32,
	outNextCursor **C.NkString,
	outPrevCursor **C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleTournamentRecordWrite
func cModuleTournamentRecordWrite(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	ownerID C.NkString,
	userName C.NkString,
	score C.NkI64,
	subscore C.NkI64,
	metadata C.NkMapAny,
	outRecord **C.NkLeaderboardRecord,
	outError **C.char) int {
	return 0
}

//export cModuleTournamentRecordsHaystack
func cModuleTournamentRecordsHaystack(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	ownerID C.NkString,
	limit C.NkU32,
	expiry C.NkI64,
	outRecords **C.NkLeaderboardRecord,
	outNumRecords **C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleGroupsGetId
func cModuleGroupsGetId(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	groupIDs *C.NkString,
	numGroupIDs C.NkU32,
	outGroups **C.NkGroup,
	outNumGroups **C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleGroupCreate
func cModuleGroupCreate(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	name C.NkString,
	creatorID C.NkString,
	langTag C.NkString,
	description C.NkString,
	avatarURL C.NkString,
	open bool,
	metadata C.NkMapAny,
	maxCount C.NkU32,
	outGroup **C.NkGroup,
	outError **C.char) int {
	return 0
}

//export cModuleGroupUpdate
func cModuleGroupUpdate(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	name C.NkString,
	creatorID C.NkString,
	langTag C.NkString,
	description C.NkString,
	avatarURL C.NkString,
	open bool,
	metadata C.NkMapAny,
	maxCount C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleGroupUserJoin
func cModuleGroupUserJoin(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	groupID C.NkString,
	userID C.NkString,
	userName C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleGroupUserLeave
func cModuleGroupUserLeave(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	groupID C.NkString,
	userID C.NkString,
	userName C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleGroupUsersAdd
func cModuleGroupUsersAdd(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	groupID C.NkString,
	userIDs *C.NkString,
	numUserIDs C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleGroupUsersDemote
func cModuleGroupUsersDemote(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	groupID C.NkString,
	userIDs *C.NkString,
	numUserIDs C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleGroupUsersKick
func cModuleGroupUsersKick(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	groupID C.NkString,
	userIDs *C.NkString,
	numUserIDs C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleGroupUsersPromote
func cModuleGroupUsersPromote(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	groupID C.NkString,
	userIDs *C.NkString,
	numUserIDs C.NkU32,
	outError **C.char) int {
	return 0
}

//export cModuleGroupUsersList
func cModuleGroupUsersList(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	groupID C.NkString,
	limit C.NkU32,
	state C.NkU32,
	cursor C.NkString,
	outUsers **C.NkGroupUserListGroupUser,
	outNumUsers **C.NkU32,
	outCursor **C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleUserGroupsList
func cModuleUserGroupsList(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	limit C.NkU32,
	state C.NkU32,
	cursor C.NkString,
	outUsers **C.NkUserGroupListUserGroup,
	outNumUsers **C.NkU32,
	outCursor **C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleFriendsList
func cModuleFriendsList(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	limit C.NkU32,
	state C.NkU32,
	cursor C.NkString,
	outFriends **C.NkFriend,
	outNumFriends **C.NkU32,
	outCursor **C.NkString,
	outError **C.char) int {
	return 0
}

//export cModuleEvent
func cModuleEvent(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	evt C.NkEvent,
	outError **C.char) int {
	return 0
}
