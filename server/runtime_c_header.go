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

extern void cContextValue(void *, NkString, NkString *);

extern int cModuleAuthenticateApple(void *, void *, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int cModuleAuthenticateCustom(void *, void *, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int cModuleAuthenticateDevice(void *, void *, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int cModuleAuthenticateEmail(void *, void *, NkString, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int cModuleAuthenticateFacebook(void *, void *, NkString, bool, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int cModuleAuthenticateFacebookInstantGame(void *, void *, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int cModuleAuthenticateGameCenter(void *, void *, NkString, NkString, NkI64, NkString, NkString, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int cModuleAuthenticateGoogle(void *, void *, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int cModuleAuthenticateSteam(void *, void *, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);

void initmodule(void *ptr, NkLogger logger, NkModule nk)
{
	InitModuleFn fn = (InitModuleFn)ptr;
	fn(logger, nk);
}

void contextvalue(void *ptr, NkString key, NkString *outvalue)
{
	return cContextValue(ptr, key, outvalue);
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

int moduleauthenticateapple(void *ptr, NkContext ctx, NkString userid, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateApple(ptr, ctx.ptr, userid, username, create, outuserid, outusername, outerr, outcreated);
}

int moduleauthenticatecustom(void *ptr, NkContext ctx, NkString userid, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateCustom(ptr, ctx.ptr, userid, username, create, outuserid, outusername, outerr, outcreated);
}

int moduleauthenticatedevice(void *ptr, NkContext ctx, NkString userid, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateDevice(ptr, ctx.ptr, userid, username, create, outuserid, outusername, outerr, outcreated);
}

int moduleauthenticateemail(void *ptr, NkContext ctx, NkString email, NkString password, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateEmail(ptr, ctx.ptr, email, password, username, create, outuserid, outusername, outerr, outcreated);
}

int moduleauthenticatefacebook(void *ptr, NkContext ctx, NkString token, bool importfriends, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateFacebook(ptr, ctx.ptr, token, importfriends, username, create, outuserid, outusername, outerr, outcreated);
}

int moduleauthenticatefacebookinstantgame(void *ptr, NkContext ctx, NkString userid, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateFacebookInstantGame(ptr, ctx.ptr, userid, username, create, outuserid, outusername, outerr, outcreated);
}

int moduleauthenticategamecenter(void *ptr, NkContext ctx, NkString playerid, NkString bundleid, NkI64 timestamp, NkString salt, NkString signature, NkString publickeyurl, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateGameCenter(ptr, ctx.ptr, playerid, bundleid, timestamp, salt, signature, publickeyurl, username, create, outuserid, outusername, outerr, outcreated);
}

int moduleauthenticategoogle(void *ptr, NkContext ctx, NkString userid, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateGoogle(ptr, ctx.ptr, userid, username, create, outuserid, outusername, outerr, outcreated);
}

int moduleauthenticatesteam(void *ptr, NkContext ctx, NkString userid, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateSteam(ptr, ctx.ptr, userid, username, create, outuserid, outusername, outerr, outcreated);
}

*/
import "C"

func nkStringGo(s C.NkString) string {
	return C.GoStringN(s.p, C.int(s.n))
}
