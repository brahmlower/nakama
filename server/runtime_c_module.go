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
func cModuleAuthenticateApple(pNk unsafe.Pointer, pCtx unsafe.Pointer, token, userName C.NkString, create bool, outUserID, outUserName, outErr *C.NkString, outCreated *bool) int {
	call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	ctx := pointer.Restore(pCtx).(context.Context)
	retUserID, retUserName, retCreated, err := call.nk.AuthenticateApple(ctx, nkStringGo(token), nkStringGo(userName), create)
	outCreated = &retCreated
	outUserID = call.goStringNk(retUserID)
	outUserName = call.goStringNk(retUserName)

	if err != nil {
		outErr = call.goStringNk(err.Error())

		return 1
	}

	return 0
}

//export cModuleAuthenticateCustom
func cModuleAuthenticateCustom(pNk unsafe.Pointer, pCtx unsafe.Pointer, userID, userName C.NkString, create bool, outUserID, outUserName, outErr *C.NkString, outCreated *bool) int {
	call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	ctx := pointer.Restore(pCtx).(context.Context)
	retUserID, retUserName, retCreated, err := call.nk.AuthenticateCustom(ctx, nkStringGo(userID), nkStringGo(userName), create)
	outCreated = &retCreated
	outUserID = call.goStringNk(retUserID)
	outUserName = call.goStringNk(retUserName)

	if err != nil {
		outErr = call.goStringNk(err.Error())

		return 1
	}

	return 0
}

//export cModuleAuthenticateDevice
func cModuleAuthenticateDevice(pNk unsafe.Pointer, pCtx unsafe.Pointer, userID, userName C.NkString, create bool, outUserID, outUserName, outErr *C.NkString, outCreated *bool) int {
	call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	ctx := pointer.Restore(pCtx).(context.Context)
	retUserID, retUserName, retCreated, err := call.nk.AuthenticateDevice(ctx, nkStringGo(userID), nkStringGo(userName), create)
	outCreated = &retCreated
	outUserID = call.goStringNk(retUserID)
	outUserName = call.goStringNk(retUserName)

	if err != nil {
		outErr = call.goStringNk(err.Error())

		return 1
	}

	return 0
}

//export cModuleAuthenticateEmail
func cModuleAuthenticateEmail(pNk unsafe.Pointer, pCtx unsafe.Pointer, email, password, userName C.NkString, create bool, outUserID, outUserName, outErr *C.NkString, outCreated *bool) int {
	call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	ctx := pointer.Restore(pCtx).(context.Context)
	retUserID, retUserName, retCreated, err := call.nk.AuthenticateEmail(ctx, nkStringGo(email), nkStringGo(password), nkStringGo(userName), create)
	outCreated = &retCreated
	outUserID = call.goStringNk(retUserID)
	outUserName = call.goStringNk(retUserName)

	if err != nil {
		outErr = call.goStringNk(err.Error())

		return 1
	}

	return 0
}

//export cModuleAuthenticateFacebook
func cModuleAuthenticateFacebook(pNk unsafe.Pointer, pCtx unsafe.Pointer, token C.NkString, importFriends bool, userName C.NkString, create bool, outUserID, outUserName, outErr *C.NkString, outCreated *bool) int {
	call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	ctx := pointer.Restore(pCtx).(context.Context)
	retUserID, retUserName, retCreated, err := call.nk.AuthenticateFacebook(ctx, nkStringGo(token), importFriends, nkStringGo(userName), create)
	outCreated = &retCreated
	outUserID = call.goStringNk(retUserID)
	outUserName = call.goStringNk(retUserName)

	if err != nil {
		outErr = call.goStringNk(err.Error())

		return 1
	}

	return 0
}

//export cModuleAuthenticateFacebookInstantGame
func cModuleAuthenticateFacebookInstantGame(pNk unsafe.Pointer, pCtx unsafe.Pointer, userID, userName C.NkString, create bool, outUserID, outUserName, outErr *C.NkString, outCreated *bool) int {
	call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	ctx := pointer.Restore(pCtx).(context.Context)
	retUserID, retUserName, retCreated, err := call.nk.AuthenticateFacebookInstantGame(ctx, nkStringGo(userID), nkStringGo(userName), create)
	outCreated = &retCreated
	outUserID = call.goStringNk(retUserID)
	outUserName = call.goStringNk(retUserName)

	if err != nil {
		outErr = call.goStringNk(err.Error())

		return 1
	}

	return 0
}

//export cModuleAuthenticateGameCenter
func cModuleAuthenticateGameCenter(pNk unsafe.Pointer, pCtx unsafe.Pointer, userID, bundleID C.NkString, timestamp int64, salt, signature, publicKeyURL, userName C.NkString, create bool, outUserID, outUserName, outErr *C.NkString, outCreated *bool) int {
	call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	ctx := pointer.Restore(pCtx).(context.Context)
	retUserID, retUserName, retCreated, err := call.nk.AuthenticateGameCenter(ctx, nkStringGo(userID), nkStringGo(bundleID), timestamp, nkStringGo(salt), nkStringGo(signature), nkStringGo(publicKeyURL), nkStringGo(userName), create)
	outCreated = &retCreated
	outUserID = call.goStringNk(retUserID)
	outUserName = call.goStringNk(retUserName)

	if err != nil {
		outErr = call.goStringNk(err.Error())

		return 1
	}

	return 0
}

//export cModuleAuthenticateGoogle
func cModuleAuthenticateGoogle(pNk unsafe.Pointer, pCtx unsafe.Pointer, userID, userName C.NkString, create bool, outUserID, outUserName, outErr *C.NkString, outCreated *bool) int {
	call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	ctx := pointer.Restore(pCtx).(context.Context)
	retUserID, retUserName, retCreated, err := call.nk.AuthenticateGoogle(ctx, nkStringGo(userID), nkStringGo(userName), create)
	outCreated = &retCreated
	outUserID = call.goStringNk(retUserID)
	outUserName = call.goStringNk(retUserName)

	if err != nil {
		outErr = call.goStringNk(err.Error())

		return 1
	}

	return 0
}

//export cModuleAuthenticateSteam
func cModuleAuthenticateSteam(pNk unsafe.Pointer, pCtx unsafe.Pointer, userID, userName C.NkString, create bool, outUserID, outUserName, outErr *C.NkString, outCreated *bool) int {
	call := pointer.Restore(pNk).(*RuntimeCNakamaModuleCall)
	ctx := pointer.Restore(pCtx).(context.Context)
	retUserID, retUserName, retCreated, err := call.nk.AuthenticateSteam(ctx, nkStringGo(userID), nkStringGo(userName), create)
	outCreated = &retCreated
	outUserID = call.goStringNk(retUserID)
	outUserName = call.goStringNk(retUserName)

	if err != nil {
		outErr = call.goStringNk(err.Error())

		return 1
	}

	return 0
}
