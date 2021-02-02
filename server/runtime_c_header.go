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

typedef int (*InitModuleFn)(NkContext, NkLogger, NkDb, NkModule, NkInitializer);

extern void cLoggerDebug(const void *, NkString);
extern void cLoggerError(const void *, NkString);
extern void cLoggerInfo(const void *, NkString);
extern void cLoggerWarn(const void *, NkString);

extern void cContextValue(const void *, NkString, NkString *);

extern int cModuleAuthenticateApple(const void *, const void *, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int cModuleAuthenticateCustom(const void *, const void *, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int cModuleAuthenticateDevice(const void *, const void *, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int cModuleAuthenticateEmail(const void *, const void *, NkString, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int cModuleAuthenticateFacebook(const void *, const void *, NkString, bool, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int cModuleAuthenticateFacebookInstantGame(const void *, const void *, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int cModuleAuthenticateGameCenter(const void *, const void *, NkString, NkString, NkI64, NkString, NkString, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int cModuleAuthenticateGoogle(const void *, const void *, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int cModuleAuthenticateSteam(const void *, const void *, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);

int initmodule(const void *ptr, NkContext ctx, NkLogger logger, NkDb db, NkModule nk, NkInitializer initializer)
{
	InitModuleFn fn = (InitModuleFn)ptr;
	return fn(ctx, logger, db, nk, initializer);
}

void contextvalue(const void *ptr, NkString key, NkString *outvalue)
{
	return cContextValue(ptr, key, outvalue);
}

void loggerdebug(const void *ptr, NkString s)
{
	cLoggerDebug(ptr, s);
}

void loggererror(const void *ptr, NkString s)
{
	cLoggerError(ptr, s);
}

void loggerinfo(const void *ptr, NkString s)
{
	cLoggerInfo(ptr, s);
}

void loggerwarn(const void *ptr, NkString s)
{
	cLoggerWarn(ptr, s);
}

int moduleauthenticateapple(const void *ptr, NkContext ctx, NkString userid, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateApple(ptr, ctx.ptr, userid, username, create, outuserid, outusername, outerr, outcreated);
}

int moduleauthenticatecustom(const void *ptr, NkContext ctx, NkString userid, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateCustom(ptr, ctx.ptr, userid, username, create, outuserid, outusername, outerr, outcreated);
}

int moduleauthenticatedevice(const void *ptr, NkContext ctx, NkString userid, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateDevice(ptr, ctx.ptr, userid, username, create, outuserid, outusername, outerr, outcreated);
}

int moduleauthenticateemail(const void *ptr, NkContext ctx, NkString email, NkString password, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateEmail(ptr, ctx.ptr, email, password, username, create, outuserid, outusername, outerr, outcreated);
}

int moduleauthenticatefacebook(const void *ptr, NkContext ctx, NkString token, bool importfriends, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateFacebook(ptr, ctx.ptr, token, importfriends, username, create, outuserid, outusername, outerr, outcreated);
}

int moduleauthenticatefacebookinstantgame(const void *ptr, NkContext ctx, NkString userid, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateFacebookInstantGame(ptr, ctx.ptr, userid, username, create, outuserid, outusername, outerr, outcreated);
}

int moduleauthenticategamecenter(const void *ptr, NkContext ctx, NkString playerid, NkString bundleid, NkI64 timestamp, NkString salt, NkString signature, NkString publickeyurl, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateGameCenter(ptr, ctx.ptr, playerid, bundleid, timestamp, salt, signature, publickeyurl, username, create, outuserid, outusername, outerr, outcreated);
}

int moduleauthenticategoogle(const void *ptr, NkContext ctx, NkString userid, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateGoogle(ptr, ctx.ptr, userid, username, create, outuserid, outusername, outerr, outcreated);
}

int moduleauthenticatesteam(const void *ptr, NkContext ctx, NkString userid, NkString username, bool create, NkString *outuserid, NkString *outusername, NkString *outerr, bool *outcreated)
{
	return cModuleAuthenticateSteam(ptr, ctx.ptr, userid, username, create, outuserid, outusername, outerr, outcreated);
}

*/
import "C"

func nkStringGo(s C.NkString) string {
	return C.GoStringN(s.p, C.int(s.n))
}
