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

extern bool *mallocbool(NkU32);

extern NkAccount *mallocnkaccount(NkU32);

extern NkI64 *mallocnki64(NkU32);

extern NkUser *mallocnkuser(NkU32);

extern void nkaccountset(NkAccount *, NkU32, NkAccount);

extern void nkuserset(NkUser *, NkU32, NkUser);
*/
import "C"

import (
	"context"
	"unsafe"

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
	retAccounts, err := nk.AccountsGetId(ctx, goStringArray(userIDs, numUserIDs))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numAccounts := C.NkU32(len(retAccounts))
	*outAccounts = C.mallocnkaccount(numAccounts)
	**outNumAccounts = numAccounts // TODO: Malloc this?

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
	retUsers, err := nk.UsersGetId(ctx, goStringArray(keys, numKeys))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numUsers := C.NkU32(len(retUsers))
	*outUsers = C.mallocnkuser(numUsers)
	**outNumUsers = numUsers // TODO: Malloc this?

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
	retUsers, err := nk.UsersGetUsername(ctx, goStringArray(keys, numKeys))
	if err != nil {
		*outErr = C.CString(err.Error())

		return 1
	}

	numUsers := C.NkU32(len(retUsers))
	*outUsers = C.mallocnkuser(numUsers)
	**outNumUsers = numUsers // TODO: Malloc this?

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
	err := nk.UsersBanId(ctx, goStringArray(userIDs, numUserIDs))
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
	err := nk.UsersUnbanId(ctx, goStringArray(userIDs, numUserIDs))
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
	return -1
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
	return -1
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
	return -1
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
	return -1
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

	**outCount = C.NkU64(retCount) // TODO: Malloc this?

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
	return -1
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
	return -1
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
	return -1
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
	outmatches **C.NkMatch,
	outNumMatches **C.NkU32,
	outErr **C.char) int {
	return -1
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
	return -1
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
	return -1
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
	return -1
}

//export cModuleWalletLedgerUpdate
func cModuleWalletLedgerUpdate(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	itemID C.NkString,
	metadata C.NkMapAny,
	outItem **C.NkWalletLedgerItem,
	outErr **C.char) int {
	return -1
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
	outErr **C.char) int {
	return -1
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
	outErr **C.char) int {
	return -1
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
	return -1
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
	return -1
}

//export cModuleStorageDelete
func cModuleStorageDelete(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	deletes *C.NkStorageDelete,
	numDeletes C.NkU32,
	outErr **C.char) int {
	return -1
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
	return -1
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
	outNextCursor **C.NkString,
	outPrevCursor **C.NkString,
	outErr **C.char) int {
	return -1
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
	outErr **C.char) int {
	return -1
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
	return -1
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
	return -1
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
	outErr **C.char) int {
	return -1
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
	outErr **C.char) int {
	return -1
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
	return -1
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
	return -1
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
	return -1
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
	return -1
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
	err := nk.GroupUsersAdd(ctx, goString(groupID), goStringArray(userIDs, numUserIDs))
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
	err := nk.GroupUsersDemote(ctx, goString(groupID), goStringArray(userIDs, numUserIDs))
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
	err := nk.GroupUsersKick(ctx, goString(groupID), goStringArray(userIDs, numUserIDs))
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
	err := nk.GroupUsersPromote(ctx, goString(groupID), goStringArray(userIDs, numUserIDs))
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
	state C.NkU32,
	cursor C.NkString,
	outUsers **C.NkGroupUserListGroupUser,
	outNumUsers **C.NkU32,
	outCursor **C.NkString,
	outErr **C.char) int {
	return -1
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
	outErr **C.char) int {
	return -1
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
	outErr **C.char) int {
	return -1
}

//export cModuleEvent
func cModuleEvent(
	pNk unsafe.Pointer,
	pCtx unsafe.Pointer,
	evt C.NkEvent,
	outErr **C.char) int {
	return -1
}
