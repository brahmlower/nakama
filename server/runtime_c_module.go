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
	"unsafe"
	"context"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/mattn/go-pointer"
)

//export cModuleAuthenticateApple
func cModuleAuthenticateApple(p unsafe.Pointer, ctx unsafe.Pointer, token, userName C.NkString, create bool) {
	pointer.Restore(p).(runtime.NakamaModule).AuthenticateApple(pointer.Restore(ctx).(context.Context), cStringGo(token), cStringGo(userName), create)
}

//export cModuleAuthenticateCustom
func cModuleAuthenticateCustom(p unsafe.Pointer, ctx unsafe.Pointer, userID, userName C.NkString, create bool) {
	pointer.Restore(p).(runtime.NakamaModule).AuthenticateCustom(pointer.Restore(ctx).(context.Context), cStringGo(userID), cStringGo(userName), create)
}

//export cModuleAuthenticateDevice
func cModuleAuthenticateDevice(p unsafe.Pointer, ctx unsafe.Pointer, userID, userName C.NkString, create bool) {
	pointer.Restore(p).(runtime.NakamaModule).AuthenticateDevice(pointer.Restore(ctx).(context.Context), cStringGo(userID), cStringGo(userName), create)
}

//export cModuleAuthenticateEmail
func cModuleAuthenticateEmail(p unsafe.Pointer, ctx unsafe.Pointer, email, password, userName C.NkString, create bool) {
	pointer.Restore(p).(runtime.NakamaModule).AuthenticateEmail(pointer.Restore(ctx).(context.Context), cStringGo(email), cStringGo(password), cStringGo(userName), create)
}

//export cModuleAuthenticateFacebook
func cModuleAuthenticateFacebook(p unsafe.Pointer, ctx unsafe.Pointer, token C.NkString, importFriends bool, userName C.NkString, create bool) {
	pointer.Restore(p).(runtime.NakamaModule).AuthenticateFacebook(pointer.Restore(ctx).(context.Context), cStringGo(token), importFriends, cStringGo(userName), create)
}

//export cModuleAuthenticateFacebookInstantGame
func cModuleAuthenticateFacebookInstantGame(p unsafe.Pointer, ctx unsafe.Pointer, userID, userName C.NkString, create bool) {
	pointer.Restore(p).(runtime.NakamaModule).AuthenticateFacebookInstantGame(pointer.Restore(ctx).(context.Context), cStringGo(userID), cStringGo(userName), create)
}

//export cModuleAuthenticateGameCenter
func cModuleAuthenticateGameCenter(p unsafe.Pointer, ctx unsafe.Pointer, userID, bundleID C.NkString, timestamp int64, salt, signature, publicKeyURL, userName C.NkString, create bool) {
	pointer.Restore(p).(runtime.NakamaModule).AuthenticateGameCenter(pointer.Restore(ctx).(context.Context), cStringGo(userID), cStringGo(bundleID), timestamp, cStringGo(salt), cStringGo(signature), cStringGo(publicKeyURL), cStringGo(userName), create)
}

//export cModuleAuthenticateGoogle
func cModuleAuthenticateGoogle(p unsafe.Pointer, ctx unsafe.Pointer, userID, userName C.NkString, create bool) {
	pointer.Restore(p).(runtime.NakamaModule).AuthenticateGoogle(pointer.Restore(ctx).(context.Context), cStringGo(userID), cStringGo(userName), create)
}

//export cModuleAuthenticateSteam
func cModuleAuthenticateSteam(p unsafe.Pointer, ctx unsafe.Pointer, userID, userName C.NkString, create bool) {
	pointer.Restore(p).(runtime.NakamaModule).AuthenticateSteam(pointer.Restore(ctx).(context.Context), cStringGo(userID), cStringGo(userName), create)
}
