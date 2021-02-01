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

typedef void (*InitModuleFn)(NkLogger logger, NkModule nk);

extern void cLoggerDebug(void *, NkString);
extern void cLoggerError(void *, NkString);
extern void cLoggerInfo(void *, NkString);
extern void cLoggerWarn(void *, NkString);

extern NkModuleAuthenticateResult cModuleAuthenticateApple(void *, void *, NkString, NkString, bool);
extern NkModuleAuthenticateResult cModuleAuthenticateCustom(void *, void *, NkString, NkString, bool);
extern NkModuleAuthenticateResult cModuleAuthenticateDevice(void *, void *, NkString, NkString, bool);
extern NkModuleAuthenticateResult cModuleAuthenticateEmail(void *, void *, NkString, NkString, NkString, bool);
extern NkModuleAuthenticateResult cModuleAuthenticateFacebook(void *, void *, NkString, bool, NkString, bool);
extern NkModuleAuthenticateResult cModuleAuthenticateFacebookInstantGame(void *, void *, NkString, NkString, bool);
extern NkModuleAuthenticateResult cModuleAuthenticateGameCenter(void *, void *, NkString, NkString, NkI64, NkString, NkString, NkString, NkString, bool);
extern NkModuleAuthenticateResult cModuleAuthenticateGoogle(void *, void *, NkString, NkString, bool);
extern NkModuleAuthenticateResult cModuleAuthenticateSteam(void *, void *, NkString, NkString, bool);

void initmodule(void *ptr, NkLogger logger, NkModule nk)
{
	InitModuleFn fn = (InitModuleFn)ptr;
	fn(logger, nk);
}

void loggerdebug(void *ptr, NkString s)
{
	cLoggerDebug(ptr, s);
}

void loggererror(void *ptr, NkString s)
{
	cLoggerError(ptr, s);
}

void loggerinfo(void *ptr, NkString s)
{
	cLoggerInfo(ptr, s);
}

void loggerwarn(void *ptr, NkString s)
{
	cLoggerWarn(ptr, s);
}

NkModuleAuthenticateResult moduleauthenticateapple(void *ptr, NkContext ctx, NkString userid, NkString username, bool create)
{
	cModuleAuthenticateApple(ptr, ctx.ptr, userid, username, create);
}

NkModuleAuthenticateResult moduleauthenticatecustom(void *ptr, NkContext ctx, NkString userid, NkString username, bool create)
{
	cModuleAuthenticateCustom(ptr, ctx.ptr, userid, username, create);
}

NkModuleAuthenticateResult moduleauthenticatedevice(void *ptr, NkContext ctx, NkString userid, NkString username, bool create)
{
	cModuleAuthenticateDevice(ptr, ctx.ptr, userid, username, create);
}

NkModuleAuthenticateResult moduleauthenticateemail(void *ptr, NkContext ctx, NkString email, NkString password, NkString username, bool create)
{
	cModuleAuthenticateEmail(ptr, ctx.ptr, email, password, username, create);
}

NkModuleAuthenticateResult moduleauthenticatefacebook(void *ptr, NkContext ctx, NkString token, bool importfriends, NkString username, bool create)
{
	cModuleAuthenticateFacebook(ptr, ctx.ptr, token, importfriends, username, create);
}

NkModuleAuthenticateResult moduleauthenticatefacebookinstantgame(void *ptr, NkContext ctx, NkString userid, NkString username, bool create)
{
	cModuleAuthenticateFacebookInstantGame(ptr, ctx.ptr, userid, username, create);
}

NkModuleAuthenticateResult moduleauthenticategamecenter(void *ptr, NkContext ctx, NkString playerid, NkString bundleid, NkI64 timestamp, NkString salt, NkString signature, NkString publickeyurl, NkString username, bool create)
{
	cModuleAuthenticateGameCenter(ptr, ctx.ptr, playerid, bundleid, timestamp, salt, signature, publickeyurl, username, create);
}

NkModuleAuthenticateResult moduleauthenticategoogle(void *ptr, NkContext ctx, NkString userid, NkString username, bool create)
{
	cModuleAuthenticateGoogle(ptr, ctx.ptr, userid, username, create);
}

NkModuleAuthenticateResult moduleauthenticatesteam(void *ptr, NkContext ctx, NkString userid, NkString username, bool create)
{
	cModuleAuthenticateSteam(ptr, ctx.ptr, userid, username, create);
}

*/
import "C"

func cStringGo(s C.NkString) string {
	return C.GoStringN(s.p, C.int(s.n))
}
