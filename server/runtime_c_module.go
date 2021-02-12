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

/*
#include "../include/nakama.h"

// Start of malloc functions

extern bool *mallocbool(NkU32);

extern NkAccount *mallocnkaccount(NkU32);

extern NkFriend *mallocnkfriend(NkU32);

extern NkGroup *mallocnkgroup(NkU32);

extern NkGroupUserListGroupUser *mallocnkgroupuserlistgroupuser(NkU32);

extern NkLeaderboardRecord *mallocnkleaderboardrecord(NkU32);

extern NkMapI64 *mallocnkmapi64(NkU32);

extern NkMatch *mallocnkmatch(NkU32);

extern NkPresence *mallocnkpresence(NkU32);

extern NkPresenceMeta *mallocnkpresencemeta(NkU32);

extern NkI64 *mallocnki64(NkU32);

extern NkU32 *mallocnku32(NkU32);

extern NkU64 *mallocnku64(NkU32);

extern NkStorageObject *mallocnkstorageobject(NkU32);

extern NkStorageObjectAck *mallocnkstorageobjectack(NkU32);

extern NkTournament *mallocnktournament(NkU32);

extern NkTournamentList *mallocnktournamentlist(NkU32);

extern NkUser *mallocnkuser(NkU32);

extern NkUserGroupListUserGroup *mallocnkusergrouplistusergroup(NkU32);

extern NkWalletLedgerItem *mallocnkwalletledgeritem(NkU32);

extern NkWalletUpdateResult *mallocnkwalletupdateresult(NkU32);

extern void nkaccountset(NkAccount *, NkU32, NkAccount);

extern void nkmapi64set(NkMapI64 *, NkU32, NkString, NkI64);

// Start of c-array member-setting functions

extern void nkfriendset(NkFriend *, NkU32, NkFriend);

extern void nkgroupuserlistgroupuserset(NkGroupUserListGroupUser *, NkU32, NkGroupUserListGroupUser);

extern void nkleaderboardrecordset(NkLeaderboardRecord *, NkU32, NkLeaderboardRecord);

extern void nkgroupset(NkGroup *, NkU32, NkGroup);

extern void nkmatchset(NkMatch *, NkU32, NkMatch);

extern void nkpresenceset(NkPresence *, NkU32, NkPresence);

extern void nkstorageobjectset(NkStorageObject *, NkU32, NkStorageObject);

extern void nkstorageobjectackset(NkStorageObjectAck *, NkU32, NkStorageObjectAck);

extern void nktournamentset(NkTournament *, NkU32, NkTournament);

extern void nkusergrouplistusergroupset(NkUserGroupListUserGroup *, NkU32, NkUserGroupListUserGroup);

extern void nkuserset(NkUser *, NkU32, NkUser);

extern void nkwalletledgeritemset(NkWalletLedgerItem *, NkU32, NkWalletLedgerItem);

extern void nkwalletupdateresultset(NkWalletUpdateResult *, NkU32, NkWalletUpdateResult);
*/
import "C"

import (
	"context"
	"unsafe"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/mattn/go-pointer"
)

//export cModuleAuthenticateApple
func cModuleAuthenticateApple(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	token,
	userName C.NkString,
	create C.bool,
	outUserID,
	outUserName,
	outErr **C.char,
	outCreated **C.bool) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retUserID, retUserName, created, err := nk.AuthenticateApple(ctx, goString(token), goString(userName), bool(create))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	*outUserID = C.CString(retUserID)

	*outUserName = C.CString(retUserName)

	*outCreated = C.mallocbool(1)
	**outCreated = C.bool(created)

	return 0
}

//export cModuleAuthenticateCustom
func cModuleAuthenticateCustom(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID,
	userName C.NkString,
	create C.bool,
	outUserID,
	outUserName,
	outErr **C.char,
	outCreated **C.bool) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retUserID, retUserName, created, err := nk.AuthenticateCustom(ctx, goString(userID), goString(userName), bool(create))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	*outUserID = C.CString(retUserID)

	*outUserName = C.CString(retUserName)

	*outCreated = C.mallocbool(1)
	**outCreated = C.bool(created)

	return 0
}

//export cModuleAuthenticateDevice
func cModuleAuthenticateDevice(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID,
	userName C.NkString,
	create C.bool,
	outUserID,
	outUserName,
	outErr **C.char,
	outCreated **C.bool) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retUserID, retUserName, created, err := nk.AuthenticateDevice(ctx, goString(userID), goString(userName), bool(create))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	*outUserID = C.CString(retUserID)

	*outUserName = C.CString(retUserName)

	*outCreated = C.mallocbool(1)
	**outCreated = C.bool(created)

	return 0
}

//export cModuleAuthenticateEmail
func cModuleAuthenticateEmail(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	email,
	password,
	userName C.NkString,
	create C.bool,
	outUserID,
	outUserName,
	outErr **C.char,
	outCreated **C.bool) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retUserID, retUserName, created, err := nk.AuthenticateEmail(ctx, goString(email), goString(password), goString(userName), bool(create))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	*outUserID = C.CString(retUserID)

	*outUserName = C.CString(retUserName)

	*outCreated = C.mallocbool(1)
	**outCreated = C.bool(created)

	return 0
}

//export cModuleAuthenticateFacebook
func cModuleAuthenticateFacebook(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	token C.NkString,
	importFriends C.bool,
	userName C.NkString,
	create C.bool,
	outUserID,
	outUserName,
	outErr **C.char,
	outCreated **C.bool) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retUserID, retUserName, created, err := nk.AuthenticateFacebook(ctx, goString(token), bool(importFriends), goString(userName), bool(create))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	*outUserID = C.CString(retUserID)

	*outUserName = C.CString(retUserName)

	*outCreated = C.mallocbool(1)
	**outCreated = C.bool(created)

	return 0
}

//export cModuleAuthenticateFacebookInstantGame
func cModuleAuthenticateFacebookInstantGame(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID,
	userName C.NkString,
	create C.bool,
	outUserID,
	outUserName,
	outErr **C.char,
	outCreated **C.bool) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retUserID, retUserName, created, err := nk.AuthenticateFacebookInstantGame(ctx, goString(userID), goString(userName), bool(create))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	*outUserID = C.CString(retUserID)

	*outUserName = C.CString(retUserName)

	*outCreated = C.mallocbool(1)
	**outCreated = C.bool(created)

	return 0
}

//export cModuleAuthenticateGameCenter
func cModuleAuthenticateGameCenter(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID,
	bundleID C.NkString,
	timestamp int64,
	salt,
	signature,
	publicKeyURL,
	userName C.NkString,
	create C.bool,
	outUserID,
	outUserName,
	outErr **C.char,
	outCreated **C.bool) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retUserID, retUserName, created, err := nk.AuthenticateGameCenter(ctx, goString(userID), goString(bundleID), timestamp, goString(salt), goString(signature), goString(publicKeyURL), goString(userName), bool(create))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	*outUserID = C.CString(retUserID)

	*outUserName = C.CString(retUserName)

	*outCreated = C.mallocbool(1)
	**outCreated = C.bool(created)

	return 0
}

//export cModuleAuthenticateGoogle
func cModuleAuthenticateGoogle(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID,
	userName C.NkString,
	create C.bool,
	outUserID,
	outUserName,
	outErr **C.char,
	outCreated **C.bool) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retUserID, retUserName, created, err := nk.AuthenticateGoogle(ctx, goString(userID), goString(userName), bool(create))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	*outUserID = C.CString(retUserID)

	*outUserName = C.CString(retUserName)

	*outCreated = C.mallocbool(1)
	**outCreated = C.bool(created)

	return 0
}

//export cModuleAuthenticateSteam
func cModuleAuthenticateSteam(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID,
	userName C.NkString,
	create C.bool,
	outUserID,
	outUserName,
	outErr **C.char,
	outCreated **C.bool) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retUserID, retUserName, created, err := nk.AuthenticateSteam(ctx, goString(userID), goString(userName), bool(create))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	*outUserID = C.CString(retUserID)

	*outUserName = C.CString(retUserName)

	*outCreated = C.mallocbool(1)
	**outCreated = C.bool(created)

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)

	retToken, retExpiry, err := nk.AuthenticateTokenGenerate(goString(userID), goString(userName), int64(expiry), goMapString(vars))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	*outToken = C.CString(retToken)

	*outExpiry = C.mallocnki64(1)
	**outExpiry = C.NkI64(retExpiry)

	return 0
}

//export cModuleAccountGetId
func cModuleAccountGetId(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	outAccount **C.NkAccount,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retAccount, err := nk.AccountGetId(ctx, goString(userID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	ptr := pointer.Save(retAccount)
	call.refs = append(call.refs, ptr)

	*outAccount = C.mallocnkaccount(1)
	(*outAccount).ptr = ptr

	return 0
}

//export cModuleAccountsGetId
func cModuleAccountsGetId(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userIDs *C.NkString,
	numUserIDs C.NkU32,
	outAccounts **C.NkAccount,
	outNumAccounts **C.NkU32,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retAccounts, err := nk.AccountsGetId(ctx, goStrings(userIDs, numUserIDs))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numAccounts := C.NkU32(len(retAccounts))
	*outNumAccounts = C.mallocnku32(1)
	**outNumAccounts = numAccounts

	*outAccounts = C.mallocnkaccount(numAccounts)
	for idx, retAccount := range retAccounts {
		ptr := pointer.Save(retAccount)
		call.refs = append(call.refs, ptr)

		outAccount := C.NkAccount{}
		outAccount.ptr = ptr

		C.nkaccountset(*outAccounts, C.NkU32(idx), outAccount)
	}

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.AccountUpdateId(ctx, goString(userID), goString(userName), goMapAny(metadata), goString(displayName), goString(timeZone), goString(location), goString(langTag), goString(avatarURL))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleAccountDeleteId
func cModuleAccountDeleteId(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	recorded C.bool,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.AccountDeleteId(ctx, goString(userID), bool(recorded))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleAccountExportId
func cModuleAccountExportId(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	outAccount **C.char,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retAccount, err := nk.AccountExportId(ctx, goString(userID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	*outAccount = C.CString(retAccount)

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retUsers, err := nk.UsersGetId(ctx, goStrings(keys, numKeys))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numUsers := C.NkU32(len(retUsers))
	*outNumUsers = C.mallocnku32(1)
	**outNumUsers = numUsers

	*outUsers = C.mallocnkuser(numUsers)
	for idx, retUser := range retUsers {
		ptr := pointer.Save(retUser)
		call.refs = append(call.refs, ptr)

		outUser := C.NkUser{}
		outUser.ptr = ptr

		C.nkuserset(*outUsers, C.NkU32(idx), outUser)
	}

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retUsers, err := nk.UsersGetUsername(ctx, goStrings(keys, numKeys))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numUsers := C.NkU32(len(retUsers))
	*outNumUsers = C.mallocnku32(1)
	**outNumUsers = numUsers

	*outUsers = C.mallocnkuser(numUsers)
	for idx, retUser := range retUsers {
		ptr := pointer.Save(retUser)
		call.refs = append(call.refs, ptr)

		outUser := C.NkUser{}
		outUser.ptr = ptr

		C.nkuserset(*outUsers, C.NkU32(idx), outUser)
	}

	return 0
}

//export cModuleUsersBanId
func cModuleUsersBanId(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userIDs *C.NkString,
	numUserIDs C.NkU32,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.UsersBanId(ctx, goStrings(userIDs, numUserIDs))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleUsersUnbanId
func cModuleUsersUnbanId(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userIDs *C.NkString,
	numUserIDs C.NkU32,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.UsersUnbanId(ctx, goStrings(userIDs, numUserIDs))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleLinkApple
func cModuleLinkApple(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.LinkApple(ctx, goString(userID), goString(linkID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleLinkCustom
func cModuleLinkCustom(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.LinkCustom(ctx, goString(userID), goString(linkID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleLinkDevice
func cModuleLinkDevice(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.LinkDevice(ctx, goString(userID), goString(linkID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleLinkFacebookInstantGame
func cModuleLinkFacebookInstantGame(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.LinkFacebookInstantGame(ctx, goString(userID), goString(linkID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleLinkGoogle
func cModuleLinkGoogle(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.LinkGoogle(ctx, goString(userID), goString(linkID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleLinkSteam
func cModuleLinkSteam(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.LinkSteam(ctx, goString(userID), goString(linkID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleUnlinkApple
func cModuleUnlinkApple(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.UnlinkApple(ctx, goString(userID), goString(linkID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleUnlinkCustom
func cModuleUnlinkCustom(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.UnlinkCustom(ctx, goString(userID), goString(linkID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleUnlinkDevice
func cModuleUnlinkDevice(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.UnlinkDevice(ctx, goString(userID), goString(linkID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleUnlinkEmail
func cModuleUnlinkEmail(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.UnlinkEmail(ctx, goString(userID), goString(linkID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleUnlinkFacebook
func cModuleUnlinkFacebook(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.UnlinkFacebook(ctx, goString(userID), goString(linkID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleUnlinkFacebookInstantGame
func cModuleUnlinkFacebookInstantGame(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.UnlinkFacebookInstantGame(ctx, goString(userID), goString(linkID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.UnlinkGameCenter(ctx, goString(userID), goString(playerID), goString(bundleID), int64(timestamp), goString(salt), goString(signature), goString(publicKeyURL))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleUnlinkGoogle
func cModuleUnlinkGoogle(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.UnlinkGoogle(ctx, goString(userID), goString(linkID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleUnlinkSteam
func cModuleUnlinkSteam(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	linkID C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.UnlinkSteam(ctx, goString(userID), goString(linkID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleLinkEmail
func cModuleLinkEmail(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	email C.NkString,
	password C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.LinkEmail(ctx, goString(userID), goString(email), goString(password))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleLinkFacebook
func cModuleLinkFacebook(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	userName C.NkString,
	token C.NkString,
	importFriends C.bool,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.LinkFacebook(ctx, goString(userID), goString(userName), goString(token), bool(importFriends))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.LinkGameCenter(ctx, goString(userID), goString(playerID), goString(bundleID), int64(timestamp), goString(salt), goString(signature), goString(publicKeyURL))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleStreamUserList
func cModuleStreamUserList(
	pNk unsafe.Pointer,
	mode C.NkU8,
	subject C.NkString,
	subContext C.NkString,
	label C.NkString,
	includeHidden C.bool,
	includeNotHidden C.bool,
	outPresences **C.NkPresence,
	outNumPresences **C.NkU32,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)

	retPresences, err := nk.StreamUserList(uint8(mode), goString(subject), goString(subContext), goString(label), bool(includeHidden), bool(includeNotHidden))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numPresences := C.NkU32(len(retPresences))
	*outNumPresences = C.mallocnku32(1)
	**outNumPresences = numPresences

	*outPresences = C.mallocnkpresence(numPresences)
	for idx, presence := range retPresences {
		ptr := pointer.Save(presence)
		call.refs = append(call.refs, ptr)

		outPresence := C.NkPresence{}
		outPresence.ptr = ptr

		C.nkpresenceset(*outPresences, C.NkU32(idx), outPresence)
	}

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)

	retMeta, err := nk.StreamUserGet(uint8(mode), goString(subject), goString(subContext), goString(label), goString(userID), goString(sessionID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	ptr := pointer.Save(retMeta)
	call.refs = append(call.refs, ptr)

	*outMeta = C.mallocnkpresencemeta(1)
	(*outMeta).ptr = ptr

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
	hidden C.bool,
	persistence C.bool,
	status C.NkString,
	outJoined **C.bool,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)

	retJoined, err := nk.StreamUserJoin(uint8(mode), goString(subject), goString(subContext), goString(label), goString(userID), goString(sessionID), bool(hidden), bool(persistence), goString(status))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	*outJoined = C.mallocbool(1)
	**outJoined = C.bool(retJoined)

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
	hidden C.bool,
	persistence C.bool,
	status C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)

	err := nk.StreamUserUpdate(uint8(mode), goString(subject), goString(subContext), goString(label), goString(userID), goString(sessionID), bool(hidden), bool(persistence), goString(status))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)

	err := nk.StreamUserLeave(uint8(mode), goString(subject), goString(subContext), goString(label), goString(userID), goString(sessionID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	user := pointer.Restore(presence.ptr).(runtime.Presence)

	err := nk.StreamUserKick(uint8(mode), goString(subject), goString(subContext), goString(label), user)
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)

	retCount, err := nk.StreamCount(uint8(mode), goString(subject), goString(subContext), goString(label))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	*outCount = C.mallocnku64(1)
	**outCount = C.NkU64(retCount)

	return 0
}

//export cModuleStreamClose
func cModuleStreamClose(
	pNk unsafe.Pointer,
	mode C.NkU8,
	subject C.NkString,
	subContext C.NkString,
	label C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)

	err := nk.StreamClose(uint8(mode), goString(subject), goString(subContext), goString(label))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

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
	reliable C.bool,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)

	err := nk.StreamSend(uint8(mode), goString(subject), goString(subContext), goString(label), goString(data), goPresences(presences, numPresences), bool(reliable))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

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
	reliable C.bool,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)

	err := nk.StreamSendRaw(uint8(mode), goString(subject), goString(subContext), goString(label), goEnvelope(msg), goPresences(presences, numPresences), bool(reliable))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleSessionDisconnect
func cModuleSessionDisconnect(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	sessionID C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.SessionDisconnect(ctx, goString(sessionID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleMatchCreate
func cModuleMatchCreate(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	module C.NkString,
	params C.NkMapAny,
	outMatchID **C.char,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retMatchID, err := nk.MatchCreate(ctx, goString(module), goMapAny(params))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	*outMatchID = C.CString(retMatchID)

	return 0
}

//export cModuleMatchGet
func cModuleMatchGet(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	outMatch **C.NkMatch,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retMatch, err := nk.MatchGet(ctx, goString(id))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	ptr := pointer.Save(retMatch)
	call.refs = append(call.refs, ptr)

	*outMatch = C.mallocnkmatch(1)
	(*outMatch).ptr = ptr

	return 0
}

//export cModuleMatchList
func cModuleMatchList(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	limit C.NkU32,
	authoritative C.bool,
	label C.NkString,
	minSize *C.NkU32,
	maxSize *C.NkU32,
	query C.NkString,
	outMatches **C.NkMatch,
	outNumMatches **C.NkU32,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	var goMinSize *int
	if minSize != nil {
		*goMinSize = int(*minSize)
	}

	var goMaxSize *int
	if maxSize != nil {
		*goMaxSize = int(*maxSize)
	}

	retMatches, err := nk.MatchList(ctx, int(limit), bool(authoritative), goString(label), goMinSize, goMaxSize, goString(query))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numMatches := C.NkU32(len(retMatches))
	*outMatches = C.mallocnkmatch(numMatches)
	*outNumMatches = C.mallocnku32(1)
	**outNumMatches = numMatches

	for idx, match := range retMatches {
		ptr := pointer.Save(match)
		call.refs = append(call.refs, ptr)

		outMatch := C.NkMatch{}
		outMatch.ptr = ptr

		C.nkmatchset(*outMatches, C.NkU32(idx), outMatch)
	}

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
	persistent C.bool,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.NotificationSend(ctx, goString(userID), goString(subject), goMapAny(content), int(code), goString(sender), bool(persistent))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleNotificationsSend
func cModuleNotificationsSend(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	notifications *C.NkNotificationSend,
	numNotifications C.NkU32,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.NotificationsSend(ctx, goNotificationSends(notifications, numNotifications))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleWalletUpdate
func cModuleWalletUpdate(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	changeset C.NkMapI64,
	metadata C.NkMapAny,
	updateLedger C.bool,
	outUpdated **C.NkMapI64,
	outPrevious **C.NkMapI64,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retUpdated, retPrevious, err := nk.WalletUpdate(ctx, goString(userID), goMapI64(changeset), goMapAny(metadata), bool(updateLedger))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	*outUpdated = cNkMapI64(retUpdated)

	*outPrevious = cNkMapI64(retPrevious)

	return 0
}

//export cModuleWalletsUpdate
func cModuleWalletsUpdate(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	updates *C.NkWalletUpdate,
	numUpdates C.NkU32,
	updateLedger C.bool,
	outResults **C.NkWalletUpdateResult,
	outNumResults **C.NkU32,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retResults, err := nk.WalletsUpdate(ctx, goWalletUpdates(updates, numUpdates), bool(updateLedger))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numResults := C.NkU32(len(retResults))
	*outNumResults = C.mallocnku32(1)
	**outNumResults = numResults

	*outResults = C.mallocnkwalletupdateresult(numResults)
	for idx, result := range retResults {
		ptr := pointer.Save(result)
		call.refs = append(call.refs, ptr)

		outResult := C.NkWalletUpdateResult{}
		outResult.ptr = ptr

		C.nkwalletupdateresultset(*outResults, C.NkU32(idx), outResult)
	}

	return 0
}

//export cModuleWalletLedgerUpdate
func cModuleWalletLedgerUpdate(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	itemID C.NkString,
	metadata C.NkMapAny,
	outItem **C.NkWalletLedgerItem,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retItem, err := nk.WalletLedgerUpdate(ctx, goString(itemID), goMapAny(metadata))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	ptr := pointer.Save(retItem)
	call.refs = append(call.refs, ptr)

	*outItem = C.mallocnkwalletledgeritem(1)
	**outItem = C.NkWalletLedgerItem{}
	(**outItem).ptr = ptr

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
	outCursor **C.char,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retItems, retCursor, err := nk.WalletLedgerList(ctx, goString(userID), int(limit), goString(cursor))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numItems := C.NkU32(len(retItems))
	*outNumItems = C.mallocnku32(1)
	**outNumItems = numItems

	*outItems = C.mallocnkwalletledgeritem(numItems)
	for idx, item := range retItems {
		ptr := pointer.Save(item)
		call.refs = append(call.refs, ptr)

		outItem := C.NkWalletLedgerItem{}
		outItem.ptr = ptr

		C.nkwalletledgeritemset(*outItems, C.NkU32(idx), outItem)
	}

	*outCursor = C.CString(retCursor)

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
	outObjs **C.NkStorageObject,
	outNumObjs **C.NkU32,
	outCursor **C.char,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retObjs, retCursor, err := nk.StorageList(ctx, goString(userID), goString(collection), int(limit), goString(cursor))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numObjs := C.NkU32(len(retObjs))
	*outNumObjs = C.mallocnku32(1)
	**outNumObjs = numObjs

	*outObjs = C.mallocnkstorageobject(numObjs)
	for idx, _ := range retObjs {
		outObj := C.NkStorageObject{}
		// TODO

		C.nkstorageobjectset(*outObjs, C.NkU32(idx), outObj)
	}

	*outCursor = C.CString(retCursor)

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retObjs, err := nk.StorageRead(ctx, goStorageReads(reads, numReads))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numObjs := C.NkU32(len(retObjs))
	*outNumObjs = C.mallocnku32(1)
	**outNumObjs = numObjs

	*outObjs = C.mallocnkstorageobject(numObjs)
	for idx, _ := range retObjs {
		outObj := C.NkStorageObject{}
		// TODO

		C.nkstorageobjectset(*outObjs, C.NkU32(idx), outObj)
	}

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retAcks, err := nk.StorageWrite(ctx, goStorageWrites(writes, numWrites))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numAcks := C.NkU32(len(retAcks))
	*outNumAcks = C.mallocnku32(1)
	**outNumAcks = numAcks

	*outAcks = C.mallocnkstorageobjectack(numAcks)
	for idx, ack := range retAcks {
		outAck := C.NkStorageObjectAck{}
		outAck.collection = cNkString(ack.Collection)
		outAck.key = cNkString(ack.Key)
		outAck.userid = cNkString(ack.UserId)
		outAck.version = cNkString(ack.Version)

		C.nkstorageobjectackset(*outAcks, C.NkU32(idx), outAck)
	}

	return 0
}

//export cModuleStorageDelete
func cModuleStorageDelete(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	deletes *C.NkStorageDelete,
	numDeletes C.NkU32,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.StorageDelete(ctx, goStorageDeletes(deletes, numDeletes))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

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
	updateLedger C.bool,
	outAcks **C.NkStorageObjectAck,
	outNumAcks **C.NkU32,
	outResults **C.NkWalletUpdateResult,
	outNumResults **C.NkU32,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retAcks, retResults, err := nk.MultiUpdate(ctx, goAccountUpdates(accountUpdates, numAccountUpdates), goStorageWrites(storageWrites, numStorageWrites), goWalletUpdates(walletUpdates, numWalletUpdates), bool(updateLedger))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numAcks := C.NkU32(len(retAcks))
	*outNumAcks = C.mallocnku32(1)
	**outNumAcks = numAcks

	*outAcks = C.mallocnkstorageobjectack(numAcks)
	for idx, ack := range retAcks {
		outAck := C.NkStorageObjectAck{}
		outAck.collection = cNkString(ack.Collection)
		outAck.key = cNkString(ack.Key)
		outAck.userid = cNkString(ack.UserId)
		outAck.version = cNkString(ack.Version)

		C.nkstorageobjectackset(*outAcks, C.NkU32(idx), outAck)
	}

	numResults := C.NkU32(len(retResults))
	*outNumResults = C.mallocnku32(1)
	**outNumResults = numResults

	*outResults = C.mallocnkwalletupdateresult(numResults)
	for idx, result := range retResults {
		ptr := pointer.Save(result)
		call.refs = append(call.refs, ptr)

		outResult := C.NkWalletUpdateResult{}
		outResult.ptr = ptr

		C.nkwalletupdateresultset(*outResults, C.NkU32(idx), outResult)
	}

	return 0
}

//export cModuleLeaderboardCreate
func cModuleLeaderboardCreate(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	authoritative C.bool,
	sortOrder C.NkString,
	op C.NkString,
	resetSchedule C.NkString,
	metadata C.NkMapAny,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.LeaderboardCreate(ctx, goString(id), bool(authoritative), goString(sortOrder), goString(op), goString(resetSchedule), goMapAny(metadata))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

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
	outNextCursor **C.char,
	outPrevCursor **C.char,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retRecords, retOwnerRecords, retNextCursor, retPrevCursor, err := nk.LeaderboardRecordsList(ctx, goString(id), goStrings(ownerIDs, numOwnerIDs), int(limit), goString(cursor), int64(expiry))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numRecords := C.NkU32(len(retRecords))
	*outNumRecords = C.mallocnku32(1)
	**outNumRecords = numRecords

	*outRecords = C.mallocnkleaderboardrecord(numRecords)
	for idx, _ := range retRecords {
		outRecord := C.NkLeaderboardRecord{}
		// TODO!

		C.nkleaderboardrecordset(*outRecords, C.NkU32(idx), outRecord)
	}

	numOwnerRecords := C.NkU32(len(retOwnerRecords))
	*outNumOwnerRecords = C.mallocnku32(1)
	**outNumOwnerRecords = numOwnerRecords

	*outOwnerRecords = C.mallocnkleaderboardrecord(numOwnerRecords)
	for idx, _ := range retOwnerRecords {
		outRecord := C.NkLeaderboardRecord{}
		// TODO!

		C.nkleaderboardrecordset(*outOwnerRecords, C.NkU32(idx), outRecord)
	}

	*outNextCursor = C.CString(retNextCursor)

	*outPrevCursor = C.CString(retPrevCursor)

	return 0
}

//export cModuleLeaderboardRecordWrite
func cModuleLeaderboardRecordWrite(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	ownerID C.NkString,
	userName C.NkString,
	score C.NkI64,
	subscore C.NkI64,
	metadata C.NkMapAny,
	outRecord **C.NkLeaderboardRecord,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retRecord, err := nk.LeaderboardRecordWrite(ctx, goString(id), goString(ownerID), goString(userName), int64(score), int64(subscore), goMapAny(metadata))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	*outRecord = C.mallocnkleaderboardrecord(1)
	**outRecord = C.NkLeaderboardRecord{}

	// TODO!
	_ = retRecord

	return 0
}

//export cModuleLeaderboardDelete
func cModuleLeaderboardDelete(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.LeaderboardDelete(ctx, goString(id))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleLeaderboardRecordDelete
func cModuleLeaderboardRecordDelete(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	ownerID C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.LeaderboardRecordDelete(ctx, goString(id), goString(ownerID))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleTournamentDelete
func cModuleTournamentDelete(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.TournamentDelete(ctx, goString(id))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleGroupDelete
func cModuleGroupDelete(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.GroupDelete(ctx, goString(id))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleTournamentCreate
func cModuleTournamentCreate(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	sortOrder C.NkString,
	operator C.NkString,
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
	joinRequired C.bool,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.TournamentCreate(ctx, goString(id), goString(sortOrder), goString(operator), goString(resetSchedule), goMapAny(metadata), goString(title), goString(description), int(category), int(startTime), int(endTime), int(duration), int(maxSize), int(maxNumScore), bool(joinRequired))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleTournamentAddAttempt
func cModuleTournamentAddAttempt(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	ownerID C.NkString,
	count C.NkU64,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.TournamentAddAttempt(ctx, goString(id), goString(ownerID), int(count))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleTournamentJoin
func cModuleTournamentJoin(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	id C.NkString,
	ownerID C.NkString,
	userName C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.TournamentJoin(ctx, goString(id), goString(ownerID), goString(userName))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retTournaments, err := nk.TournamentsGetId(ctx, goStrings(tournamentIDs, numTournamentIDs))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numTournaments := C.NkU32(len(retTournaments))
	*outNumTournaments = C.mallocnku32(1)
	**outNumTournaments = numTournaments

	*outTournaments = C.mallocnktournament(numTournaments)
	for idx, _ := range retTournaments {
		outTournament := C.NkTournament{}
		// TODO

		C.nktournamentset(*outTournaments, C.NkU32(idx), outTournament)
	}

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retTournaments, err := nk.TournamentList(ctx, int(catStart), int(catEnd), int(startTime), int(endTime), int(limit), goString(cursor))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	if retTournaments != nil {
		tournaments := C.mallocnktournament(1)
		for idx, _ := range retTournaments.Tournaments {
			tournament := C.NkTournament{}
			// TODO

			C.nktournamentset(tournaments, C.NkU32(idx), tournament)
		}

		*outTournaments = C.mallocnktournamentlist(1)
		**outTournaments = C.NkTournamentList{}
		(*outTournaments).tournaments = tournaments
		(*outTournaments).numtournaments = C.NkU32(len(retTournaments.Tournaments))
		(*outTournaments).cursor = cNkString(retTournaments.Cursor)
	}

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
	outNextCursor **C.char,
	outPrevCursor **C.char,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retRecords, retOwnerRecords, retNextCursor, retPrevCursor, err := nk.TournamentRecordsList(ctx, goString(tournamentID), goStrings(ownerIDs, numOwnerIDs), int(limit), goString(cursor), int64(overrideExpiry))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numRecords := C.NkU32(len(retRecords))
	*outNumRecords = C.mallocnku32(1)
	**outNumRecords = numRecords

	*outRecords = C.mallocnkleaderboardrecord(numRecords)
	for idx, _ := range retRecords {
		outRecord := C.NkLeaderboardRecord{}
		// TODO!

		C.nkleaderboardrecordset(*outRecords, C.NkU32(idx), outRecord)
	}

	numOwnerRecords := C.NkU32(len(retOwnerRecords))
	*outNumOwnerRecords = C.mallocnku32(1)
	**outNumOwnerRecords = numOwnerRecords

	*outOwnerRecords = C.mallocnkleaderboardrecord(numOwnerRecords)
	for idx, _ := range retOwnerRecords {
		outRecord := C.NkLeaderboardRecord{}
		// TODO!

		C.nkleaderboardrecordset(*outOwnerRecords, C.NkU32(idx), outRecord)
	}

	*outNextCursor = C.CString(retNextCursor)

	*outPrevCursor = C.CString(retPrevCursor)

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retRecord, err := nk.TournamentRecordWrite(ctx, goString(id), goString(ownerID), goString(userName), int64(score), int64(subscore), goMapAny(metadata))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	if retRecord != nil {
		*outRecord = C.mallocnkleaderboardrecord(1)
		**outRecord = C.NkLeaderboardRecord{}
		// TODO!
	}

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retRecords, err := nk.TournamentRecordsHaystack(ctx, goString(id), goString(ownerID), int(limit), int64(expiry))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numRecords := C.NkU32(len(retRecords))
	*outNumRecords = C.mallocnku32(1)
	**outNumRecords = numRecords

	*outRecords = C.mallocnkleaderboardrecord(numRecords)
	for idx, _ := range retRecords {
		outRecord := C.NkLeaderboardRecord{}
		// TODO!

		C.nkleaderboardrecordset(*outRecords, C.NkU32(idx), outRecord)
	}

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
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retGroups, err := nk.GroupsGetId(ctx, goStrings(groupIDs, numGroupIDs))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numGroups := C.NkU32(len(retGroups))
	*outNumGroups = C.mallocnku32(1)
	**outNumGroups = numGroups

	*outGroups = C.mallocnkgroup(numGroups)
	for idx, _ := range retGroups {
		group := C.NkGroup{}
		// TODO!

		C.nkgroupset(*outGroups, C.NkU32(idx), group)
	}

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
	open C.bool,
	metadata C.NkMapAny,
	maxCount C.NkU32,
	outGroup **C.NkGroup,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	retGroup, err := nk.GroupCreate(ctx, goString(userID), goString(name), goString(creatorID), goString(langTag), goString(description), goString(avatarURL), bool(open), goMapAny(metadata), int(maxCount))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	if retGroup != nil {
		*outGroup = C.mallocnkgroup(1)
		**outGroup = C.NkGroup{}
		// TODO!
	}

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
	open C.bool,
	metadata C.NkMapAny,
	maxCount C.NkU32,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.GroupUpdate(ctx, goString(userID), goString(name), goString(creatorID), goString(langTag), goString(description), goString(avatarURL), bool(open), goMapAny(metadata), int(maxCount))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleGroupUserJoin
func cModuleGroupUserJoin(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	groupID C.NkString,
	userID C.NkString,
	userName C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.GroupUserJoin(ctx, goString(groupID), goString(userID), goString(userName))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleGroupUserLeave
func cModuleGroupUserLeave(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	groupID C.NkString,
	userID C.NkString,
	userName C.NkString,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.GroupUserLeave(ctx, goString(groupID), goString(userID), goString(userName))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleGroupUsersAdd
func cModuleGroupUsersAdd(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	groupID C.NkString,
	userIDs *C.NkString,
	numUserIDs C.NkU32,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.GroupUsersAdd(ctx, goString(groupID), goStrings(userIDs, numUserIDs))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleGroupUsersDemote
func cModuleGroupUsersDemote(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	groupID C.NkString,
	userIDs *C.NkString,
	numUserIDs C.NkU32,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)
	err := nk.GroupUsersDemote(ctx, goString(groupID), goStrings(userIDs, numUserIDs))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleGroupUsersKick
func cModuleGroupUsersKick(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	groupID C.NkString,
	userIDs *C.NkString,
	numUserIDs C.NkU32,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.GroupUsersKick(ctx, goString(groupID), goStrings(userIDs, numUserIDs))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleGroupUsersPromote
func cModuleGroupUsersPromote(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	groupID C.NkString,
	userIDs *C.NkString,
	numUserIDs C.NkU32,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	err := nk.GroupUsersPromote(ctx, goString(groupID), goStrings(userIDs, numUserIDs))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}

//export cModuleGroupUsersList
func cModuleGroupUsersList(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	groupID C.NkString,
	limit C.NkU32,
	state *C.NkU32,
	cursor C.NkString,
	outUsers **C.NkGroupUserListGroupUser,
	outNumUsers **C.NkU32,
	outCursor **C.char,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	var goState *int
	if state != nil {
		*goState = int(*state)
	}

	retUsers, retCursor, err := nk.GroupUsersList(ctx, goString(groupID), int(limit), goState, goString(cursor))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numUsers := C.NkU32(len(retUsers))
	*outNumUsers = C.mallocnku32(1)
	**outNumUsers = numUsers

	*outUsers = C.mallocnkgroupuserlistgroupuser(numUsers)
	for idx, _ := range retUsers {
		outUser := C.NkGroupUserListGroupUser{}
		// TODO

		C.nkgroupuserlistgroupuserset(*outUsers, C.NkU32(idx), outUser)
	}

	*outCursor = C.CString(retCursor)

	return 0
}

//export cModuleUserGroupsList
func cModuleUserGroupsList(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	limit C.NkU32,
	state *C.NkU32,
	cursor C.NkString,
	outGroups **C.NkUserGroupListUserGroup,
	outNumGroups **C.NkU32,
	outCursor **C.char,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	var goState *int
	if state != nil {
		*goState = int(*state)
	}

	retGroups, retCursor, err := nk.UserGroupsList(ctx, goString(userID), int(limit), goState, goString(cursor))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numGroups := C.NkU32(len(retGroups))
	*outNumGroups = C.mallocnku32(1)
	**outNumGroups = numGroups

	*outGroups = C.mallocnkusergrouplistusergroup(numGroups)
	for idx, _ := range retGroups {
		outGroup := C.NkUserGroupListUserGroup{}
		// TODO

		C.nkusergrouplistusergroupset(*outGroups, C.NkU32(idx), outGroup)
	}

	*outCursor = C.CString(retCursor)

	return 0
}

//export cModuleFriendsList
func cModuleFriendsList(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	userID C.NkString,
	limit C.NkU32,
	state *C.NkU32,
	cursor C.NkString,
	outFriends **C.NkFriend,
	outNumFriends **C.NkU32,
	outCursor **C.char,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)

	var goState *int
	if state != nil {
		*goState = int(*state)
	}

	retFriends, retCursor, err := nk.FriendsList(ctx, goString(userID), int(limit), goState, goString(cursor))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numFriends := C.NkU32(len(retFriends))
	*outNumFriends = C.mallocnku32(1)
	**outNumFriends = numFriends

	*outFriends = C.mallocnkfriend(numFriends)
	for idx, _ := range retFriends {
		outFriend := C.NkFriend{}
		// TODO

		C.nkfriendset(*outFriends, C.NkU32(idx), outFriend)
	}

	*outCursor = C.CString(retCursor)

	return 0
}

//export cModuleEvent
func cModuleEvent(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	evt C.NkEvent,
	outErr **C.char) int {
	call := pointer.Restore(pNk).(*RuntimeCCall)
	nk := pointer.Restore(call.ptr).(runtime.NakamaModule)
	ctx := pointer.Restore(pCtx).(context.Context)
	goEvt := pointer.Restore(evt.ptr).(*api.Event)

	err := nk.Event(ctx, goEvt)
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	return 0
}
