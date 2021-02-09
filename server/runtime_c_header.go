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

typedef int (*InitModuleFn)(
	NkContext,
	NkLogger,
	NkDb,
	NkModule,
	NkInitializer);

extern void cLoggerDebug(
	const void *ptr,
	NkString s);

extern void cLoggerError(
	const void *ptr,
	NkString s);

extern void cLoggerInfo(
	const void *ptr,
	NkString s);

extern void cLoggerWarn(
	const void *ptr,
	NkString s);

extern void cContextValue(
	const void *ptr,
	NkString key,
	NkString **outvalue);

extern int cModuleAuthenticateApple(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated);

extern int cModuleAuthenticateCustom(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated);

extern int cModuleAuthenticateDevice(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated);

extern int cModuleAuthenticateEmail(
	const void *ptr,
	const NkContext *ctx,
	NkString email,
	NkString password,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated);

extern int cModuleAuthenticateFacebook(
	const void *ptr,
	const NkContext *ctx,
	NkString token,
	bool importfriends,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated);

extern int cModuleAuthenticateFacebookInstantGame(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated);

extern int cModuleAuthenticateGameCenter(
	const void *ptr,
	const NkContext *ctx,
	NkString playerid,
	NkString bundleid,
	NkI64 timestamp,
	NkString salt,
	NkString signature,
	NkString publickeyurl,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated);

extern int cModuleAuthenticateGoogle(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated);

extern int cModuleAuthenticateSteam(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated);

extern int cModuleAuthenticateTokenGenerate(
	const void *ptr,
	NkString userid,
	NkString username,
	NkI64 expiry,
	NkMapString vars,
	NkString **outtoken,
	NkI64 **outexpiry,
	char **outerror);

extern int cModuleAccountGetId(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkAccount **outaccount,
	char **outerror);

extern int cModuleAccountsGetId(
	const void *ptr,
	const NkContext *ctx,
	NkString *userids,
	NkU32 numuserids,
	NkAccount **outaccounts,
	NkU32 **outnumaccounts,
	char **outerror);

extern int cModuleAccountUpdateId(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	NkMapAny metadata,
	NkString displayname,
	NkString timezone,
	NkString location,
	NkString langtag,
	NkString avatarurl,
	char **outerror);

extern int cModuleAccountDeleteId(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	bool recorded,
	char **outerror);

extern int cModuleAccountExportId(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString **outaccount,
	char **outerror);

extern int cModuleUsersGetId(
	const void *ptr,
	const NkContext *ctx,
	const NkString *keys,
	NkU32 numkeys,
	NkUser **outusers,
	NkU32 **outnumusers,
	char **outerror);

extern int cModuleUsersGetUsername(
	const void *ptr,
	const NkContext *ctx,
	const NkString *keys,
	NkU32 numkeys,
	NkUser **outusers,
	NkU32 **outnumusers,
	char **outerror);

extern int cModuleUsersBanId(
	const void *ptr,
	const NkContext *ctx,
	const NkString *userids,
	NkU32 numids,
	char **outerror);

extern int cModuleUsersUnbanId(
	const void *ptr,
	const NkContext *ctx,
	const NkString *userids,
	NkU32 numids,
	char **outerror);

extern int cModuleLinkApple(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror);

extern int cModuleLinkCustom(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror);

extern int cModuleLinkDevice(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror);

extern int cModuleLinkEmail(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString email,
	NkString password,
	char **outerror);

extern int cModuleLinkFacebook(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	NkString token,
	bool importfriends,
	char **outerror);

extern int cModuleLinkFacebookInstantGame(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror);

extern int cModuleLinkGameCenter(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString playerid,
	NkString bundleid,
	NkI64 timestamp,
	NkString salt,
	NkString signature,
	NkString publickeyurl,
	char **outerror);

extern int cModuleLinkGoogle(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror);

extern int cModuleLinkSteam(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror);

extern int cModuleUnlinkApple(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror);

extern int cModuleUnlinkCustom(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror);

extern int cModuleUnlinkDevice(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror);

extern int cModuleUnlinkEmail(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror);

extern int cModuleUnlinkFacebook(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror);

extern int cModuleUnlinkFacebookInstantGame(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror);

extern int cModuleUnlinkGameCenter(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString playerid,
	NkString bundleid,
	NkI64 timestamp,
	NkString salt,
	NkString signature,
	NkString publickeyurl,
	char **outerror);

extern int cModuleUnlinkGoogle(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror);

extern int cModuleUnlinkSteam(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror);

extern int cModuleStreamUserList(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	bool includehidden,
	bool includenothidden,
	NkPresence **outpresences,
	NkU32 **outnumpresences,
	char **outerror);

extern int cModuleStreamUserGet(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkString userid,
	NkString sessionid,
	NkPresenceMeta **outmeta,
	char **outerror);

extern int cModuleStreamUserJoin(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkString userid,
	NkString sessionid,
	bool hidden,
	bool persistence,
	NkString status,
	bool **outjoined,
	char **outerror);

extern int cModuleStreamUserUpdate(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkString userid,
	NkString sessionid,
	bool hidden,
	bool persistence,
	NkString status,
	char **outerror);

extern int cModuleStreamUserLeave(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkString userid,
	NkString sessionid,
	char **outerror);

extern int cModuleStreamUserKick(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkPresence presence,
	char **outerror);

extern int cModuleStreamCount(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkU64 **outcount,
	char **outerror);

extern int cModuleStreamClose(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	char **outerror);

extern int cModuleStreamSend(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkString data,
	NkPresence *presences,
	NkU32 numpresences,
	bool reliable,
	char **outerror);

extern int cModuleStreamSendRaw(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkEnvelope msg,
	const NkPresence *presences,
	NkU32 numpresences,
	bool reliable,
	char **outerror);

extern int cModuleSessionDisconnect(
	const void *ptr,
	const NkContext *ctx,
	NkString sessionid,
	char **outerror);

extern int cModuleMatchCreate(
	const void *ptr,
	const NkContext *ctx,
	NkString module,
	NkMapAny params,
	NkString **outmatchid,
	char **outerror);

extern int cModuleMatchGet(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkMatch **outmatch,
	char **outerror);

extern int cModuleMatchList(
	const void *ptr,
	const NkContext *ctx,
	NkU32 limit,
	bool authoritative,
	NkString label,
	const NkU32 *minsize,
	const NkU32 *maxsize,
	NkString query,
	NkMatch **outmatches,
	NkU32 **outnummatches,
	char **outerror);

extern int cModuleNotificationSend(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString subject,
	NkMapAny content,
	NkU64 code,
	NkString sender,
	bool persistent,
	char **outerror);

extern int cModuleNotificationsSend(
	const void *ptr,
	const NkContext *ctx,
	const NkNotificationSend *notifications,
	NkU32 numnotifications,
	char **outerror);

extern int cModuleWalletUpdate(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkMapI64 changeset,
	NkMapAny metadata,
	bool updateledger,
	NkMapI64 **outupdated,
	NkMapI64 **outprevious,
	char **outerror);

extern int cModuleWalletsUpdate(
	const void *ptr,
	const NkContext *ctx,
	const NkWalletUpdate *updates,
	NkU32 numupdates,
	bool updateledger,
	NkWalletUpdateResult **outresults,
	NkU32 **outnumresults,
	char **outerror);

extern int cModuleWalletLedgerUpdate(
	const void *ptr,
	const NkContext *ctx,
	NkString itemid,
	NkMapAny metadata,
	NkWalletLedgerItem **outitem,
	char **outerror);

extern int cModuleWalletLedgerList(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkU32 limit,
	NkString cursor,
	NkWalletLedgerItem **outitems,
	NkU32 **outnumitems,
	NkString **outcursor,
	char **outerror);

extern int cModuleStorageList(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString collection,
	NkU32 limit,
	NkString cursor,
	NkStorageObject **outobjs,
	NkU32 **outnumobjs,
	NkString **outcursor,
	char **outerror);

extern int cModuleStorageRead(
	const void *ptr,
	const NkContext *ctx,
	const NkStorageRead *reads,
	NkU32 numreads,
	NkStorageObject **outobjs,
	NkU32 **outnumobjs,
	char **outerror);

extern int cModuleStorageWrite(
	const void *ptr,
	const NkContext *ctx,
	const NkStorageWrite *writes,
	NkU32 numwrites,
	NkStorageObjectAck **outacks,
	NkU32 **outnumacks,
	char **outerror);

extern int cModuleStorageDelete(
	const void *ptr,
	const NkContext *ctx,
	const NkStorageDelete *deletes,
	NkU32 numdeletes,
	char **outerror);

extern int cModuleMultiUpdate(
	const void *ptr,
	const NkContext *ctx,
	const NkAccountUpdate *accountupdates,
	NkU32 numaccountupdates,
	const NkStorageWrite *storagewrites,
	NkU32 numstoragewrites,
	const NkWalletUpdate *walletupdates,
	NkU32 numwalletupdates,
	bool updateledger,
	NkStorageObjectAck **outacks,
	NkU32 **outnumacks,
	NkWalletUpdateResult **outresults,
	NkU32 **outnumresults,
	char **outerror);

extern int cModuleLeaderboardCreate(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	bool authoritative,
	NkString sortorder,
	NkString op,
	NkString resetschedule,
	NkMapAny metadata,
	char **outerror);

extern int cModuleLeaderboardRecordsList(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	const NkString *ownerids,
	NkU32 numownerids,
	NkU32 limit,
	NkString cursor,
	NkI64 expiry,
	NkLeaderboardRecord **outrecords,
	NkU32 **outnumrecords,
	NkLeaderboardRecord **outownerrecords,
	NkU32 **outnumownerrecords,
	NkString **outnextcursor,
	NkString **outprevcursor,
	char **outerror);

extern int cModuleLeaderboardRecordWrite(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkI64 score,
	NkI64 subscore,
	NkMapAny metadata,
	NkLeaderboardRecord **outrecord,
	char **outerror);

extern int cModuleLeaderboardDelete(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	char **outerror);

extern int cModuleLeaderboardRecordDelete(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	char **outerror);

extern int cModuleTournamentDelete(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	char **outerror);

extern int cModuleGroupDelete(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	char **outerror);

extern int cModuleTournamentCreate(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString sortorder,
	NkString resetschedule,
	NkMapAny metadata,
	NkString title,
	NkString description,
	NkU32 category,
	NkU32 starttime,
	NkU32 endtime,
	NkU32 duration,
	NkU32 maxsize,
	NkU32 maxnumscore,
	bool joinrequired,
	char **outerror);

extern int cModuleTournamentAddAttempt(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkU64 count,
	char **outerror);

extern int cModuleTournamentJoin(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkString username,
	char **outerror);

extern int cModuleTournamentsGetId(
	const void *ptr,
	const NkContext *ctx,
	const NkString *tournamentids,
	NkU32 numtournamentids,
	NkTournament **outtournaments,
	NkU32 **outnumtournaments,
	char **outerror);

extern int cModuleTournamentList(
	const void *ptr,
	const NkContext *ctx,
	NkU64 catstart,
	NkU64 catend,
	NkU64 starttime,
	NkU64 endtime,
	NkU32 limit,
	NkString cursor,
	NkString id,
	NkTournamentList **outtournaments,
	NkU32 **outnumtournaments,
	char **outerror);

extern int cModuleTournamentRecordsList(
	const void *ptr,
	const NkContext *ctx,
	NkString tournamentid,
	const NkString *ownerids,
	NkU32 numownerids,
	NkU32 limit,
	NkString cursor,
	NkI64 overrideexpiry,
	NkLeaderboardRecord **outrecords,
	NkU32 **outnumrecords,
	NkLeaderboardRecord **outownerrecords,
	NkU32 **outnumownerrecords,
	NkString **outnextcursor,
	NkString **outprevcursor,
	char **outerror);

extern int cModuleTournamentRecordWrite(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkString username,
	NkI64 score,
	NkI64 subscore,
	NkMapAny metadata,
	NkLeaderboardRecord **outrecord,
	char **outerror);

extern int cModuleTournamentRecordsHaystack(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkU32 limit,
	NkI64 expiry,
	NkLeaderboardRecord **outrecords,
	NkU32 **outnumrecords,
	char **outerror);

extern int cModuleGroupsGetId(
	const void *ptr,
	const NkContext *ctx,
	const NkString *groupids,
	NkU32 numgroupids,
	NkGroup **outgroups,
	NkU32 **outnumgroups,
	char **outerror);

extern int cModuleGroupCreate(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString name,
	NkString creatorid,
	NkString langtag,
	NkString description,
	NkString avatarurl,
	bool open,
	NkMapAny metadata,
	NkU32 maxcount,
	NkGroup **outgroup,
	char **outerror);

extern int cModuleGroupUpdate(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString name,
	NkString creatorid,
	NkString langtag,
	NkString description,
	NkString avatarurl,
	bool open,
	NkMapAny metadata,
	NkU32 maxcount,
	char **outerror);

extern int cModuleGroupUserJoin(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	NkString userid,
	NkString username,
	char **outerror);

extern int cModuleGroupUserLeave(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	NkString userid,
	NkString username,
	char **outerror);

extern int cModuleGroupUsersAdd(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	const NkString *userids,
	NkU32 numuserids,
	char **outerror);

extern int cModuleGroupUsersDemote(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	const NkString *userids,
	NkU32 numuserids,
	char **outerror);

extern int cModuleGroupUsersKick(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	const NkString *userids,
	NkU32 numuserids,
	char **outerror);

extern int cModuleGroupUsersPromote(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	const NkString *userids,
	NkU32 numuserids,
	char **outerror);

extern int cModuleGroupUsersList(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	NkU32 limit,
	NkU32 state,
	NkString cursor,
	NkGroupUserListGroupUser **outusers,
	NkU32 **outnumusers,
	NkString **outcursor,
	char **outerror);

extern int cModuleUserGroupsList(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkU32 limit,
	NkU32 state,
	NkString cursor,
	NkUserGroupListUserGroup **outusers,
	NkU32 **outnumusers,
	NkString **outcursor,
	char **outerror);

extern int cModuleFriendsList(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkU32 limit,
	NkU32 state,
	NkString cursor,
	NkFriend **outfriends,
	NkU32 **outnumfriends,
	NkString **outcursor,
	char **outerror);

extern int cModuleEvent(
	const void *ptr,
	const NkContext *ctx,
	NkEvent evt,
	char **outerror);

extern int cInitializerRegisterRpc(
	const void *ptr,
	NkString id,
	const NkRpcFn fn,
	char **outerror);

extern int cInitializerRegisterBeforeRt(
	const void *ptr,
	NkString id,
	const NkBeforeRtCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterRt(
	const void *ptr,
	NkString id,
	const NkAfterRtCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterMatchmakerMatched(
	const void *ptr,
	const NkMatchmakerMatchedCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterMatch(
	const void *ptr,
	NkString name,
	const NkMatchCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterTournamentEnd(
	const void *ptr,
	const NkTournamentCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterTournamentReset(
	const void *ptr,
	const NkTournamentCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterLeaderBoardEnd(
	const void *ptr,
	const NkLeaderBoardCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterLeaderBoardReset(
	const void *ptr,
	const NkLeaderBoardCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeGetAccount(
	const void *ptr,
	const NkCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterGetAccount(
	const void *ptr,
	const NkAfterGetAccountCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeUpdateAccount(
	const void *ptr,
	const NkBeforeUpdateAccountCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterUpdateAccount(
	const void *ptr,
	const NkAfterUpdateAccountCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeSessionRefresh(
	const void *ptr,
	const NkBeforeSessionRefreshCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterSessionRefresh(
	const void *ptr,
	const NkAfterSessionRefreshCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeAuthenticateApple(
	const void *ptr,
	const NkBeforeAuthenticateAppleCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterAuthenticateApple(
	const void *ptr,
	const NkAfterAuthenticateAppleCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeAuthenticateCustom(
	const void *ptr,
	const NkBeforeAuthenticateCustomCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterAuthenticateCustom(
	const void *ptr,
	const NkAfterAuthenticateCustomCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeAuthenticateDevice(
	const void *ptr,
	const NkBeforeAuthenticateDeviceCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterAuthenticateDevice(
	const void *ptr,
	const NkAfterAuthenticateDeviceCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeAuthenticateEmail(
	const void *ptr,
	const NkBeforeAuthenticateEmailCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterAuthenticateEmail(
	const void *ptr,
	const NkAfterAuthenticateEmailCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeAuthenticateFacebook(
	const void *ptr,
	const NkBeforeAuthenticateFacebookCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterAuthenticateFacebook(
	const void *ptr,
	const NkAfterAuthenticateFacebookCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeAuthenticateFacebookInstantGame(
	const void *ptr,
	const NkBeforeAuthenticateFacebookInstantGameCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterAuthenticateFacebookInstantGame(
	const void *ptr,
	const NkAfterAuthenticateFacebookInstantGameCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeAuthenticateGameCenter(
	const void *ptr,
	const NkBeforeAuthenticateGameCenterCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterAuthenticateGameCenter(
	const void *ptr,
	const NkAfterAuthenticateGameCenterCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeAuthenticateGoogle(
	const void *ptr,
	const NkBeforeAuthenticateGoogleCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterAuthenticateGoogle(
	const void *ptr,
	const NkAfterAuthenticateGoogleCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeAuthenticateSteam(
	const void *ptr,
	const NkBeforeAuthenticateSteamCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterAuthenticateSteam(
	const void *ptr,
	const NkAfterAuthenticateSteamCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeListChannelMessages(
	const void *ptr,
	const NkBeforeListChannelMessagesCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterListChannelMessages(
	const void *ptr,
	const NkAfterListChannelMessagesCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeListFriends(
	const void *ptr,
	const NkBeforeListFriendsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterListFriends(
	const void *ptr,
	const NkAfterListFriendsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeAddFriends(
	const void *ptr,
	const NkBeforeAddFriendsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterAddFriends(
	const void *ptr,
	const NkAfterAddFriendsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeDeleteFriends(
	const void *ptr,
	const NkBeforeDeleteFriendsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterDeleteFriends(
	const void *ptr,
	const NkAfterDeleteFriendsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeBlockFriends(
	const void *ptr,
	const NkBeforeBlockFriendsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterBlockFriends(
	const void *ptr,
	const NkAfterBlockFriendsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeImportFacebookFriends(
	const void *ptr,
	const NkBeforeImportFacebookFriendsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterImportFacebookFriends(
	const void *ptr,
	const NkAfterImportFacebookFriendsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeCreateGroup(
	const void *ptr,
	const NkBeforeCreateGroupCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterCreateGroup(
	const void *ptr,
	const NkAfterCreateGroupCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeUpdateGroup(
	const void *ptr,
	const NkBeforeUpdateGroupCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterUpdateGroup(
	const void *ptr,
	const NkAfterUpdateGroupCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeDeleteGroup(
	const void *ptr,
	const NkBeforeDeleteGroupCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterDeleteGroup(
	const void *ptr,
	const NkAfterDeleteGroupCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeJoinGroup(
	const void *ptr,
	const NkBeforeJoinGroupCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterJoinGroup(
	const void *ptr,
	const NkAfterJoinGroupCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeLeaveGroup(
	const void *ptr,
	const NkBeforeLeaveGroupCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterLeaveGroup(
	const void *ptr,
	const NkAfterLeaveGroupCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeAddGroupUsers(
	const void *ptr,
	const NkBeforeAddGroupUsersCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterAddGroupUsers(
	const void *ptr,
	const NkAfterAddGroupUsersCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeBanGroupUsers(
	const void *ptr,
	const NkBeforeBanGroupUsersCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterBanGroupUsers(
	const void *ptr,
	const NkAfterBanGroupUsersCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeKickGroupUsers(
	const void *ptr,
	const NkBeforeKickGroupUsersCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterKickGroupUsers(
	const void *ptr,
	const NkAfterKickGroupUsersCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforePromoteGroupUsers(
	const void *ptr,
	const NkBeforePromoteGroupUsersCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterPromoteGroupUsers(
	const void *ptr,
	const NkAfterPromoteGroupUsersCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeDemoteGroupUsers(
	const void *ptr,
	const NkBeforeDemoteGroupUsersCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterDemoteGroupUsers(
	const void *ptr,
	const NkAfterDemoteGroupUsersCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeListGroupUsers(
	const void *ptr,
	const NkBeforeListGroupUsersCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterListGroupUsers(
	const void *ptr,
	const NkAfterListGroupUsersCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeListUserGroups(
	const void *ptr,
	const NkBeforeListUserGroupsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterListUserGroups(
	const void *ptr,
	const NkAfterListUserGroupsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeListGroups(
	const void *ptr,
	const NkBeforeListGroupsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterListGroups(
	const void *ptr,
	const NkAfterListGroupsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeDeleteLeaderboardRecord(
	const void *ptr,
	const NkBeforeDeleteLeaderboardRecordCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterDeleteLeaderboardRecord(
	const void *ptr,
	const NkAfterDeleteLeaderboardRecordCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeListLeaderboardRecords(
	const void *ptr,
	const NkBeforeListLeaderboardRecordsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterListLeaderboardRecords(
	const void *ptr,
	const NkAfterListLeaderboardRecordsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeWriteLeaderboardRecord(
	const void *ptr,
	const NkBeforeWriteLeaderboardRecordCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterWriteLeaderboardRecord(
	const void *ptr,
	const NkAfterWriteLeaderboardRecordCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeListLeaderboardRecordsAroundOwner(
	const void *ptr,
	const NkBeforeListLeaderboardRecordsAroundOwnerCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterListLeaderboardRecordsAroundOwner(
	const void *ptr,
	const NkAfterListLeaderboardRecordsAroundOwnerCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeLinkApple(
	const void *ptr,
	const NkBeforeLinkAppleCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterLinkApple(
	const void *ptr,
	const NkAfterLinkAppleCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeLinkCustom(
	const void *ptr,
	const NkBeforeLinkCustomCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterLinkCustom(
	const void *ptr,
	const NkAfterLinkCustomCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeLinkDevice(
	const void *ptr,
	const NkBeforeLinkDeviceCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterLinkDevice(
	const void *ptr,
	const NkAfterLinkDeviceCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeLinkEmail(
	const void *ptr,
	const NkBeforeLinkEmailCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterLinkEmail(
	const void *ptr,
	const NkAfterLinkEmailCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeLinkFacebook(
	const void *ptr,
	const NkBeforeLinkFacebookCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterLinkFacebook(
	const void *ptr,
	const NkAfterLinkFacebookCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeLinkFacebookInstantGame(
	const void *ptr,
	const NkBeforeLinkFacebookInstantGameCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterLinkFacebookInstantGame(
	const void *ptr,
	const NkAfterLinkFacebookInstantGameCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeLinkGameCenter(
	const void *ptr,
	const NkBeforeLinkGameCenterCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterLinkGameCenter(
	const void *ptr,
	const NkAfterLinkGameCenterCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeLinkGoogle(
	const void *ptr,
	const NkBeforeLinkGoogleCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterLinkGoogle(
	const void *ptr,
	const NkAfterLinkGoogleCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeLinkSteam(
	const void *ptr,
	const NkBeforeLinkSteamCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterLinkSteam(
	const void *ptr,
	const NkAfterLinkSteamCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeListMatches(
	const void *ptr,
	const NkBeforeListMatchesCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterListMatches(
	const void *ptr,
	const NkAfterListMatchesCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeListNotifications(
	const void *ptr,
	const NkBeforeListNotificationsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterListNotifications(
	const void *ptr,
	const NkAfterListNotificationsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeDeleteNotifications(
	const void *ptr,
	const NkBeforeDeleteNotificationsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterDeleteNotifications(
	const void *ptr,
	const NkAfterDeleteNotificationsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeListStorageObjects(
	const void *ptr,
	const NkBeforeListStorageObjectsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterListStorageObjects(
	const void *ptr,
	const NkAfterListStorageObjectsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeReadStorageObjects(
	const void *ptr,
	const NkBeforeReadStorageObjectsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterReadStorageObjects(
	const void *ptr,
	const NkAfterReadStorageObjectsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeWriteStorageObjects(
	const void *ptr,
	const NkBeforeWriteStorageObjectsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterWriteStorageObjects(
	const void *ptr,
	const NkAfterWriteStorageObjectsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeDeleteStorageObjects(
	const void *ptr,
	const NkBeforeDeleteStorageObjectsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterDeleteStorageObjects(
	const void *ptr,
	const NkAfterDeleteStorageObjectsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeJoinTournament(
	const void *ptr,
	const NkBeforeJoinTournamentCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterJoinTournament(
	const void *ptr,
	const NkAfterJoinTournamentCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeListTournamentRecords(
	const void *ptr,
	const NkBeforeListTournamentRecordsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterListTournamentRecords(
	const void *ptr,
	const NkAfterListTournamentRecordsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeListTournaments(
	const void *ptr,
	const NkBeforeListTournamentsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterListTournaments(
	const void *ptr,
	const NkAfterListTournamentsCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeWriteTournamentRecord(
	const void *ptr,
	const NkBeforeWriteTournamentRecordCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterWriteTournamentRecord(
	const void *ptr,
	const NkAfterWriteTournamentRecordCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeListTournamentRecordsAroundOwner(
	const void *ptr,
	const NkBeforeListTournamentRecordsAroundOwnerCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterListTournamentRecordsAroundOwner(
	const void *ptr,
	const NkAfterListTournamentRecordsAroundOwnerCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeUnlinkApple(
	const void *ptr,
	const NkBeforeUnlinkAppleCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterUnlinkApple(
	const void *ptr,
	const NkAfterUnlinkAppleCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeUnlinkCustom(
	const void *ptr,
	const NkBeforeUnlinkCustomCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterUnlinkCustom(
	const void *ptr,
	const NkAfterUnlinkCustomCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeUnlinkDevice(
	const void *ptr,
	const NkBeforeUnlinkDeviceCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterUnlinkDevice(
	const void *ptr,
	const NkAfterUnlinkDeviceCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeUnlinkEmail(
	const void *ptr,
	const NkBeforeUnlinkEmailCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterUnlinkEmail(
	const void *ptr,
	const NkAfterUnlinkEmailCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeUnlinkFacebook(
	const void *ptr,
	const NkBeforeUnlinkFacebookCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterUnlinkFacebook(
	const void *ptr,
	const NkAfterUnlinkFacebookCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeUnlinkFacebookInstantGame(
	const void *ptr,
	const NkBeforeUnlinkFacebookInstantGameCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterUnlinkFacebookInstantGame(
	const void *ptr,
	const NkAfterUnlinkFacebookInstantGameCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeUnlinkGameCenter(
	const void *ptr,
	const NkBeforeUnlinkGameCenterCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterUnlinkGameCenter(
	const void *ptr,
	const NkAfterUnlinkGameCenterCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeUnlinkGoogle(
	const void *ptr,
	const NkBeforeUnlinkGoogleCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterUnlinkGoogle(
	const void *ptr,
	const NkAfterUnlinkGoogleCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeUnlinkSteam(
	const void *ptr,
	const NkBeforeUnlinkSteamCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterUnlinkSteam(
	const void *ptr,
	const NkAfterUnlinkSteamCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterBeforeGetUsers(
	const void *ptr,
	const NkBeforeGetUsersCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterAfterGetUsers(
	const void *ptr,
	const NkAfterGetUsersCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterEvent(
	const void *ptr,
	const NkEventCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterEventSessionStart(
	const void *ptr,
	const NkEventCallbackFn cb,
	char **outerror);

extern int cInitializerRegisterEventSessionEnd(
	const void *ptr,
	const NkEventCallbackFn cb,
	char **outerror);

int initmodule(
	const void *ptr,
	NkContext ctx,
	NkLogger logger,
	NkDb db,
	NkModule nk,
	NkInitializer initializer)
{
	InitModuleFn fn = (InitModuleFn)ptr;
	return fn(
		ctx,
		logger,
		db,
		nk,
		initializer);
}

void contextvalue(
	const void *ptr,
	NkString key,
	NkString **outvalue)
{
	return cContextValue(
		ptr,
		key,
		outvalue);
}

void loggerdebug(
	const void *ptr,
	NkString s)
{
	cLoggerDebug(
		ptr,
		s);
}

void loggererror(
	const void *ptr,
	NkString s)
{
	cLoggerError(
		ptr,
		s);
}

void loggerinfo(
	const void *ptr,
	NkString s)
{
	cLoggerInfo(
		ptr,
		s);
}

void loggerwarn(
	const void *ptr,
	NkString s)
{
	cLoggerWarn(
		ptr,
		s);
}

int moduleauthenticateapple(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated)
{
	return cModuleAuthenticateApple(
		ptr,
		ctx->ptr,
		userid,
		username,
		create,
		outuserid,
		outusername,
		outerror,
		outcreated);
}

int moduleauthenticatecustom(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated)
{
	return cModuleAuthenticateCustom(
		ptr,
		ctx->ptr,
		userid,
		username,
		create,
		outuserid,
		outusername,
		outerror,
		outcreated);
}

int moduleauthenticatedevice(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated)
{
	return cModuleAuthenticateDevice(
		ptr,
		ctx->ptr,
		userid,
		username,
		create,
		outuserid,
		outusername,
		outerror,
		outcreated);
}

int moduleauthenticateemail(
	const void *ptr,
	const NkContext *ctx,
	NkString email,
	NkString password,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated)
{
	return cModuleAuthenticateEmail(
		ptr,
		ctx->ptr,
		email,
		password,
		username,
		create,
		outuserid,
		outusername,
		outerror,
		outcreated);
}

int moduleauthenticatefacebook(
	const void *ptr,
	const NkContext *ctx,
	NkString token,
	bool importfriends,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated)
{
	return cModuleAuthenticateFacebook(
		ptr,
		ctx->ptr,
		token,
		importfriends,
		username,
		create,
		outuserid,
		outusername,
		outerror,
		outcreated);
}

int moduleauthenticatefacebookinstantgame(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated)
{
	return cModuleAuthenticateFacebookInstantGame(
		ptr,
		ctx->ptr,
		userid,
		username,
		create,
		outuserid,
		outusername,
		outerror,
		outcreated);
}

int moduleauthenticategamecenter(
	const void *ptr,
	const NkContext *ctx,
	NkString playerid,
	NkString bundleid,
	NkI64 timestamp,
	NkString salt,
	NkString signature,
	NkString publickeyurl,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated)
{
	return cModuleAuthenticateGameCenter(
		ptr,
		ctx->ptr,
		playerid,
		bundleid,
		timestamp,
		salt,
		signature,
		publickeyurl,
		username,
		create,
		outuserid,
		outusername,
		outerror,
		outcreated);
}

int moduleauthenticategoogle(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated)
{
	return cModuleAuthenticateGoogle(
		ptr,
		ctx->ptr,
		userid,
		username,
		create,
		outuserid,
		outusername,
		outerror,
		outcreated);
}

int moduleauthenticatesteam(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	char **outuserid,
	char **outusername,
	char **outerror,
	bool **outcreated)
{
	return cModuleAuthenticateSteam(
		ptr,
		ctx->ptr,
		userid,
		username,
		create,
		outuserid,
		outusername,
		outerror,
		outcreated);
}

int moduleauthenticatetokengenerate(
	const void *ptr,
	NkString userid,
	NkString username,
	NkI64 expiry,
	NkMapString vars,
	NkString **outtoken,
	NkI64 **outexpiry,
	char **outerror)
{
	return cModuleAuthenticateTokenGenerate(
		ptr,
		userid,
		username,
		expiry,
		vars,
		outtoken,
		outexpiry,
		outerror);
}

int moduleaccountgetid(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkAccount **outaccount,
	char **outerror)
{
	return cModuleAccountGetId(
		ptr,
		ctx->ptr,
		userid,
		outaccount,
		outerror);
}

int moduleaccountsgetid(
	const void *ptr,
	const NkContext *ctx,
	NkString *userids,
	NkU32 numuserids,
	NkAccount **outaccounts,
	NkU32 **outnumaccounts,
	char **outerror)
{
	return cModuleAccountsGetId(
		ptr,
		ctx->ptr,
		userids,
		numuserids,
		outaccounts,
		outnumaccounts,
		outerror);
}

int moduleaccountupdateid(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	NkMapAny metadata,
	NkString displayname,
	NkString timezone,
	NkString location,
	NkString langtag,
	NkString avatarurl,
	char **outerror)
{
	return cModuleAccountUpdateId(
		ptr,
		ctx->ptr,
		userid,
		username,
		metadata,
		displayname,
		timezone,
		location,
		langtag,
		avatarurl,
		outerror);
}

int moduleaccountdeleteid(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		bool recorded,
		char **outerror)
{
	return cModuleAccountDeleteId(
		ptr,
		ctx,
		userid,
		recorded,
		outerror);
}

int moduleaccountexportid(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString **outaccount,
	char **outerror)
{
	return cModuleAccountExportId(
		ptr,
		ctx,
		userid,
		outaccount,
		outerror);
}

int moduleusersgetid(
	const void *ptr,
	const NkContext *ctx,
	const NkString *keys,
	NkU32 numkeys,
	NkUser **outusers,
	NkU32 **outnumusers,
	char **outerror)
{
	return cModuleUsersGetId(
		ptr,
		ctx,
		keys,
		numkeys,
		outusers,
		outnumusers,
		outerror);
}

int moduleusersgetusername(
	const void *ptr,
	const NkContext *ctx,
	const NkString *keys,
	NkU32 numkeys,
	NkUser **outusers,
	NkU32 **outnumusers,
	char **outerror)
{
	return cModuleUsersGetUsername(
		ptr,
		ctx,
		keys,
		numkeys,
		outusers,
		outnumusers,
		outerror);
}

int moduleusersbanid(
	const void *ptr,
	const NkContext *ctx,
	const NkString *userids,
	NkU32 numids,
	char **outerror)
{
	return cModuleUsersBanId(
		ptr,
		ctx,
		userids,
		numids,
		outerror);
}

int moduleusersunbanid(
	const void *ptr,
	const NkContext *ctx,
	const NkString *userids,
	NkU32 numids,
	char **outerror)
{
	return cModuleUsersUnbanId(
		ptr,
		ctx,
		userids,
		numids,
		outerror);
}

int modulelinkapple(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror)
{
	return cModuleLinkApple(
		ptr,
		ctx,
		userid,
		linkid,
		outerror);
}

int modulelinkcustom(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror)
{
	return cModuleLinkCustom(
		ptr,
		ctx,
		userid,
		linkid,
		outerror);
}

int modulelinkdevice(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror)
{
	return cModuleLinkDevice(
		ptr,
		ctx,
		userid,
		linkid,
		outerror);
}

int modulelinkfacebookinstantgame(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror)
{
	return cModuleLinkFacebookInstantGame(
		ptr,
		ctx,
		userid,
		linkid,
		outerror);
}

int modulelinkgoogle(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror)
{
	return cModuleLinkGoogle(
		ptr,
		ctx,
		userid,
		linkid,
		outerror);
}

int modulelinksteam(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror)
{
	return cModuleLinkSteam(
		ptr,
		ctx,
		userid,
		linkid,
		outerror);
}

int moduleunlinkapple(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror)
{
	return cModuleUnlinkApple(
		ptr,
		ctx,
		userid,
		linkid,
		outerror);
}

int moduleunlinkcustom(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror)
{
	return cModuleUnlinkCustom(
		ptr,
		ctx,
		userid,
		linkid,
		outerror);
}

int moduleunlinkdevice(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror)
{
	return cModuleUnlinkDevice(
		ptr,
		ctx,
		userid,
		linkid,
		outerror);
}

int moduleunlinkemail(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror)
{
	return cModuleUnlinkEmail(
		ptr,
		ctx,
		userid,
		linkid,
		outerror);
}

int moduleunlinkfacebook(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror)
{
	return cModuleUnlinkFacebook(
		ptr,
		ctx,
		userid,
		linkid,
		outerror);
}

int moduleunlinkfacebookinstantgame(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror)
{
	return cModuleUnlinkFacebookInstantGame(
		ptr,
		ctx,
		userid,
		linkid,
		outerror);
}

int moduleunlinkgamecenter(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString playerid,
	NkString bundleid,
	NkI64 timestamp,
	NkString salt,
	NkString signature,
	NkString publickeyurl,
	char **outerror)
{
	return cModuleUnlinkGameCenter(
		ptr,
		ctx,
		userid,
		playerid,
		bundleid,
		timestamp,
		salt,
		signature,
		publickeyurl,
		outerror);
}

int moduleunlinkgoogle(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror)
{
	return cModuleUnlinkGoogle(
		ptr,
		ctx,
		userid,
		linkid,
		outerror);
}

int moduleunlinksteam(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	char **outerror)
{
	return cModuleUnlinkSteam(
		ptr,
		ctx,
		userid,
		linkid,
		outerror);
}

int modulelinkemail(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString email,
	NkString password,
	char **outerror)
{
	return cModuleLinkEmail(
		ptr,
		ctx,
		userid,
		email,
		password,
		outerror);
}

int modulelinkfacebook(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	NkString token,
	bool importfriends,
	char **outerror)
{
	return cModuleLinkFacebook(
		ptr,
		ctx,
		userid,
		username,
		token,
		importfriends,
		outerror);
}

int modulelinkgamecenter(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString playerid,
	NkString bundleid,
	NkI64 timestamp,
	NkString salt,
	NkString signature,
	NkString publickeyurl,
	char **outerror)
{
	return cModuleLinkGameCenter(
		ptr,
		ctx,
		userid,
		playerid,
		bundleid,
		timestamp,
		salt,
		signature,
		publickeyurl,
		outerror
	);
}

int modulestreamuserlist(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	bool includehidden,
	bool includenothidden,
	NkPresence **outpresences,
	NkU32 **outnumpresences,
	char **outerror)
{
	return cModuleStreamUserList(
		ptr,
		mode,
		subject,
		subcontext,
		label,
		includehidden,
		includenothidden,
		outpresences,
		outnumpresences,
		outerror);
}

int modulestreamuserget(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkString userid,
	NkString sessionid,
	NkPresenceMeta **outmeta,
	char **outerror)
{
	return cModuleStreamUserGet(
		ptr,
		mode,
		subject,
		subcontext,
		label,
		userid,
		sessionid,
		outmeta,
		outerror);
}

int modulestreamuserjoin(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkString userid,
	NkString sessionid,
	bool hidden,
	bool persistence,
	NkString status,
	bool **outjoined,
	char **outerror)
{
	return cModuleStreamUserJoin(
		ptr,
		mode,
		subject,
		subcontext,
		label,
		userid,
		sessionid,
		hidden,
		persistence,
		status,
		outjoined,
		outerror);
}

int modulestreamuserupdate(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkString userid,
	NkString sessionid,
	bool hidden,
	bool persistence,
	NkString status,
	char **outerror)
{
	return cModuleStreamUserUpdate(
		ptr,
		mode,
		subject,
		subcontext,
		label,
		userid,
		sessionid,
		hidden,
		persistence,
		status,
		outerror);
}

int modulestreamuserleave(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkString userid,
	NkString sessionid,
	char **outerror)
{
	return cModuleStreamUserLeave(
		ptr,
		mode,
		subject,
		subcontext,
		label,
		userid,
		sessionid,
		outerror);
}

int modulestreamuserkick(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkPresence presence,
	char **outerror)
{
	return cModuleStreamUserKick(
		ptr,
		mode,
		subject,
		subcontext,
		label,
		presence,
		outerror);
}

int modulestreamcount(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkU64 **outcount,
	char **outerror)
{
	return cModuleStreamCount(
		ptr,
		mode,
		subject,
		subcontext,
		label,
		outcount,
		outerror);
}

int modulestreamclose(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	char **outerror)
{
	return cModuleStreamClose(
		ptr,
		mode,
		subject,
		subcontext,
		label,
		outerror);
}

int modulestreamsend(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkString data,
	NkPresence *presences,
	NkU32 numpresences,
	bool reliable,
	char **outerror)
{
	return cModuleStreamSend(
		ptr,
		mode,
		subject,
		subcontext,
		label,
		data,
		presences,
		numpresences,
		reliable,
		outerror);
}

int modulestreamsendraw(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkEnvelope msg,
	const NkPresence *presences,
	NkU32 numpresences,
	bool reliable,
	char **outerror)
{
	return cModuleStreamSendRaw(
		ptr,
		mode,
		subject,
		subcontext,
		label,
		msg,
		presences,
		numpresences,
		reliable,
		outerror);
}

int modulesessiondisconnect(
	const void *ptr,
	const NkContext *ctx,
	NkString sessionid,
	char **outerror)
{
	return cModuleSessionDisconnect(
		ptr,
		ctx,
		sessionid,
		outerror);
}

int modulematchcreate(
	const void *ptr,
	const NkContext *ctx,
	NkString module,
	NkMapAny params,
	NkString **outmatchid,
	char **outerror)
{
	return cModuleMatchCreate(
		ptr,
		ctx,
		module,
		params,
		outmatchid,
		outerror);
}

int modulematchget(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkMatch **outmatch,
	char **outerror)
{
	return cModuleMatchGet(
		ptr,
		ctx,
		id,
		outmatch,
		outerror);
}

int modulematchlist(
	const void *ptr,
	const NkContext *ctx,
	NkU32 limit,
	bool authoritative,
	NkString label,
	const NkU32 *minsize,
	const NkU32 *maxsize,
	NkString query,
	NkMatch **outmatches,
	NkU32 **outnummatches,
	char **outerror)
{
	return cModuleMatchList(
		ptr,
		ctx,
		limit,
		authoritative,
		label,
		minsize,
		maxsize,
		query,
		outmatches,
		outnummatches,
		outerror);
}

int modulenotificationsend(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString subject,
	NkMapAny content,
	NkU64 code,
	NkString sender,
	bool persistent,
	char **outerror)
{
	return cModuleNotificationSend(
		ptr,
		ctx,
		userid,
		subject,
		content,
		code,
		sender,
		persistent,
		outerror);
}

int modulenotificationssend(
	const void *ptr,
	const NkContext *ctx,
	const NkNotificationSend *notifications,
	NkU32 numnotifications,
	char **outerror)
{
	return cModuleNotificationsSend(
		ptr,
		ctx,
		notifications,
		numnotifications,
		outerror);
}

int modulewalletupdate(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkMapI64 changeset,
	NkMapAny metadata,
	bool updateledger,
	NkMapI64 **outupdated,
	NkMapI64 **outprevious,
	char **outerror)
{
	return cModuleWalletUpdate(
		ptr,
		ctx,
		userid,
		changeset,
		metadata,
		updateledger,
		outupdated,
		outprevious,
		outerror);
}

int modulewalletsupdate(
	const void *ptr,
	const NkContext *ctx,
	const NkWalletUpdate *updates,
	NkU32 numupdates,
	bool updateledger,
	NkWalletUpdateResult **outresults,
	NkU32 **outnumresults,
	char **outerror)
{
	return cModuleWalletsUpdate(
		ptr,
		ctx,
		updates,
		numupdates,
		updateledger,
		outresults,
		outnumresults,
		outerror);
}

int modulewalletledgerupdate(
	const void *ptr,
	const NkContext *ctx,
	NkString itemid,
	NkMapAny metadata,
	NkWalletLedgerItem **outitem,
	char **outerror)
{
	return cModuleWalletLedgerUpdate(
		ptr,
		ctx,
		itemid,
		metadata,
		outitem,
		outerror);
}

int modulewalletledgerlist(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkU32 limit,
	NkString cursor,
	NkWalletLedgerItem **outitems,
	NkU32 **outnumitems,
	NkString **outcursor,
	char **outerror)
{
	return cModuleWalletLedgerList(
		ptr,
		ctx,
		userid,
		limit,
		cursor,
		outitems,
		outnumitems,
		outcursor,
		outerror);
}

int modulestoragelist(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString collection,
	NkU32 limit,
	NkString cursor,
	NkStorageObject **outobjs,
	NkU32 **outnumobjs,
	NkString **outcursor,
	char **outerror)
{
	return cModuleStorageList(
		ptr,
		ctx,
		userid,
		collection,
		limit,
		cursor,
		outobjs,
		outnumobjs,
		outcursor,
		outerror);
}

int modulestorageread(
	const void *ptr,
	const NkContext *ctx,
	const NkStorageRead *reads,
	NkU32 numreads,
	NkStorageObject **outobjs,
	NkU32 **outnumobjs,
	char **outerror)
{
	return cModuleStorageRead(
		ptr,
		ctx,
		reads,
		numreads,
		outobjs,
		outnumobjs,
		outerror);
}

int modulestoragewrite(
	const void *ptr,
	const NkContext *ctx,
	const NkStorageWrite *writes,
	NkU32 numwrites,
	NkStorageObjectAck **outacks,
	NkU32 **outnumacks,
	char **outerror)
{
	return cModuleStorageWrite(
		ptr,
		ctx,
		writes,
		numwrites,
		outacks,
		outnumacks,
		outerror);
}

int modulestoragedelete(
	const void *ptr,
	const NkContext *ctx,
	const NkStorageDelete *deletes,
	NkU32 numdeletes,
	char **outerror)
{
	return cModuleStorageDelete(
		ptr,
		ctx,
		deletes,
		numdeletes,
		outerror);
}

int modulemultiupdate(
	const void *ptr,
	const NkContext *ctx,
	const NkAccountUpdate *accountupdates,
	NkU32 numaccountupdates,
	const NkStorageWrite *storagewrites,
	NkU32 numstoragewrites,
	const NkWalletUpdate *walletupdates,
	NkU32 numwalletupdates,
	bool updateledger,
	NkStorageObjectAck **outacks,
	NkU32 **outnumacks,
	NkWalletUpdateResult **outresults,
	NkU32 **outnumresults,
	char **outerror)
{
	return cModuleMultiUpdate(
		ptr,
		ctx,
		accountupdates,
		numaccountupdates,
		storagewrites,
		numstoragewrites,
		walletupdates,
		numwalletupdates,
		updateledger,
		outacks,
		outnumacks,
		outresults,
		outnumresults,
		outerror);
}

int moduleleaderboardcreate(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	bool authoritative,
	NkString sortorder,
	NkString op,
	NkString resetschedule,
	NkMapAny metadata,
	char **outerror)
{
	return cModuleLeaderboardCreate(
		ptr,
		ctx,
		id,
		authoritative,
		sortorder,
		op,
		resetschedule,
		metadata,
		outerror);
}

int moduleleaderboardrecordslist(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	const NkString *ownerids,
	NkU32 numownerids,
	NkU32 limit,
	NkString cursor,
	NkI64 expiry,
	NkLeaderboardRecord **outrecords,
	NkU32 **outnumrecords,
	NkLeaderboardRecord **outownerrecords,
	NkU32 **outnumownerrecords,
	NkString **outnextcursor,
	NkString **outprevcursor,
	char **outerror)
{
	return cModuleLeaderboardRecordsList(
		ptr,
		ctx,
		id,
		ownerids,
		numownerids,
		limit,
		cursor,
		expiry,
		outrecords,
		outnumrecords,
		outownerrecords,
		outnumownerrecords,
		outnextcursor,
		outprevcursor,
		outerror);
}

int moduleleaderboardrecordwrite(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkI64 score,
	NkI64 subscore,
	NkMapAny metadata,
	NkLeaderboardRecord **outrecord,
	char **outerror)
{
	return cModuleLeaderboardRecordWrite(
		ptr,
		ctx,
		id,
		ownerid,
		score,
		subscore,
		metadata,
		outrecord,
		outerror);
}

int moduleleaderboarddelete(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	char **outerror)
{
	return cModuleLeaderboardDelete(
		ptr,
		ctx,
		id,
		ownerid,
		outerror);
}

int moduleleaderboardrecorddelete(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	char **outerror)
{
	return cModuleLeaderboardRecordDelete(
		ptr,
		ctx,
		id,
		ownerid,
		outerror);
}

int moduletournamentdelete(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	char **outerror)
{
	return cModuleTournamentDelete(
		ptr,
		ctx,
		id,
		ownerid,
		outerror);
}

int modulegroupdelete(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	char **outerror)
{
	return cModuleGroupDelete(
		ptr,
		ctx,
		id,
		ownerid,
		outerror);
}

int moduletournamentcreate(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString sortorder,
	NkString resetschedule,
	NkMapAny metadata,
	NkString title,
	NkString description,
	NkU32 category,
	NkU32 starttime,
	NkU32 endtime,
	NkU32 duration,
	NkU32 maxsize,
	NkU32 maxnumscore,
	bool joinrequired,
	char **outerror)
{
	return cModuleTournamentCreate(
		ptr,
		ctx,
		id,
		sortorder,
		resetschedule,
		metadata,
		title,
		description,
		category,
		starttime,
		endtime,
		duration,
		maxsize,
		maxnumscore,
		joinrequired,
		outerror);
}

int moduletournamentaddattempt(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkU64 count,
	char **outerror)
{
	return cModuleTournamentAddAttempt(
		ptr,
		ctx,
		id,
		ownerid,
		count,
		outerror);
}

int moduletournamentjoin(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkString username,
	char **outerror)
{
	return cModuleTournamentJoin(
		ptr,
		ctx,
		id,
		ownerid,
		username,
		outerror);
}

int moduletournamentsgetid(
	const void *ptr,
	const NkContext *ctx,
	const NkString *tournamentids,
	NkU32 numtournamentids,
	NkTournament **outtournaments,
	NkU32 **outnumtournaments,
	char **outerror)
{
	return cModuleTournamentsGetId(
		ptr,
		ctx,
		tournamentids,
		numtournamentids,
		outtournaments,
		outnumtournaments,
		outerror);
}

int moduletournamentlist(
	const void *ptr,
	const NkContext *ctx,
	NkU64 catstart,
	NkU64 catend,
	NkU64 starttime,
	NkU64 endtime,
	NkU32 limit,
	NkString cursor,
	NkString id,
	NkTournamentList **outtournaments,
	NkU32 **outnumtournaments,
	char **outerror)
{
	return cModuleTournamentList(
		ptr,
		ctx,
		catstart,
		catend,
		starttime,
		endtime,
		limit,
		cursor,
		id,
		outtournaments,
		outnumtournaments,
		outerror);
}

int moduletournamentrecordslist(
	const void *ptr,
	const NkContext *ctx,
	NkString tournamentid,
	const NkString *ownerids,
	NkU32 numownerids,
	NkU32 limit,
	NkString cursor,
	NkI64 overrideexpiry,
	NkLeaderboardRecord **outrecords,
	NkU32 **outnumrecords,
	NkLeaderboardRecord **outownerrecords,
	NkU32 **outnumownerrecords,
	NkString **outnextcursor,
	NkString **outprevcursor,
	char **outerror)
{
	return cModuleTournamentRecordsList(
		ptr,
		ctx,
		tournamentid,
		ownerids,
		numownerids,
		limit,
		cursor,
		overrideexpiry,
		outrecords,
		outnumrecords,
		outownerrecords,
		outnumownerrecords,
		outnextcursor,
		outprevcursor,
		outerror);
}

int moduletournamentrecordwrite(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkString username,
	NkI64 score,
	NkI64 subscore,
	NkMapAny metadata,
	NkLeaderboardRecord **outrecord,
	char **outerror)
{
	return cModuleTournamentRecordWrite(
		ptr,
		ctx,
		id,
		ownerid,
		username,
		score,
		subscore,
		metadata,
		outrecord,
		outerror);
}

int moduletournamentrecordshaystack(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkU32 limit,
	NkI64 expiry,
	NkLeaderboardRecord **outrecords,
	NkU32 **outnumrecords,
	char **outerror)
{
	return cModuleTournamentRecordsHaystack(
		ptr,
		ctx,
		id,
		ownerid,
		limit,
		expiry,
		outrecords,
		outnumrecords,
		outerror);
}

int modulegroupsgetid(
	const void *ptr,
	const NkContext *ctx,
	const NkString *groupids,
	NkU32 numgroupids,
	NkGroup **outgroups,
	NkU32 **outnumgroups,
	char **outerror)
{
	return cModuleGroupsGetId(
		ptr,
		ctx,
		groupids,
		numgroupids,
		outgroups,
		outnumgroups,
		outerror);
}

int modulegroupcreate(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString name,
	NkString creatorid,
	NkString langtag,
	NkString description,
	NkString avatarurl,
	bool open,
	NkMapAny metadata,
	NkU32 maxcount,
	NkGroup **outgroup,
	char **outerror)
{
	return cModuleGroupCreate(
		ptr,
		ctx,
		userid,
		name,
		creatorid,
		langtag,
		description,
		avatarurl,
		open,
		metadata,
		maxcount,
		outgroup,
		outerror);
}

int modulegroupupdate(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString name,
	NkString creatorid,
	NkString langtag,
	NkString description,
	NkString avatarurl,
	bool open,
	NkMapAny metadata,
	NkU32 maxcount,
	char **outerror)
{
	return cModuleGroupUpdate(
		ptr,
		ctx,
		userid,
		name,
		creatorid,
		langtag,
		description,
		avatarurl,
		open,
		metadata,
		maxcount,
		outerror);
}

int modulegroupuserjoin(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	NkString userid,
	NkString username,
	char **outerror)
{
	return cModuleGroupUserJoin(
		ptr,
		ctx,
		groupid,
		userid,
		username,
		outerror);
}

int modulegroupuserleave(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	NkString userid,
	NkString username,
	char **outerror)
{
	return cModuleGroupUserLeave(
		ptr,
		ctx,
		groupid,
		userid,
		username,
		outerror);
}

int modulegroupusersadd(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	const NkString *userids,
	NkU32 numuserids,
	char **outerror)
{
	return cModuleGroupUsersAdd(
		ptr,
		ctx,
		groupid,
		userids,
		numuserids,
		outerror);
}

int modulegroupusersdemote(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	const NkString *userids,
	NkU32 numuserids,
	char **outerror)
{
	return cModuleGroupUsersDemote(
		ptr,
		ctx,
		groupid,
		userids,
		numuserids,
		outerror);
}

int modulegroupuserskick(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	const NkString *userids,
	NkU32 numuserids,
	char **outerror)
{
	return cModuleGroupUsersKick(
		ptr,
		ctx,
		groupid,
		userids,
		numuserids,
		outerror);
}

int modulegroupuserspromote(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	const NkString *userids,
	NkU32 numuserids,
	char **outerror)
{
	return cModuleGroupUsersPromote(
		ptr,
		ctx,
		groupid,
		userids,
		numuserids,
		outerror);
}

int modulegroupuserslist(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	NkU32 limit,
	NkU32 state,
	NkString cursor,
	NkGroupUserListGroupUser **outusers,
	NkU32 **outnumusers,
	NkString **outcursor,
	char **outerror)
{
	return cModuleGroupUsersList(
		ptr,
		ctx,
		groupid,
		limit,
		state,
		cursor,
		outusers,
		outnumusers,
		outcursor,
		outerror);
}

int moduleusergroupslist(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkU32 limit,
	NkU32 state,
	NkString cursor,
	NkUserGroupListUserGroup **outusers,
	NkU32 **outnumusers,
	NkString **outcursor,
	char **outerror)
{
	return cModuleUserGroupsList(
		ptr,
		ctx,
		userid,
		limit,
		state,
		cursor,
		outusers,
		outnumusers,
		outcursor,
		outerror);
}

int modulefriendslist(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkU32 limit,
	NkU32 state,
	NkString cursor,
	NkFriend **outfriends,
	NkU32 **outnumfriends,
	NkString **outcursor,
	char **outerror)
{
	return cModuleFriendsList(
		ptr,
		ctx,
		userid,
		limit,
		state,
		cursor,
		outfriends,
		outnumfriends,
		outcursor,
		outerror);
}

int moduleevent(
	const void *ptr,
	const NkContext *ctx,
	NkEvent evt,
	char **outerror)
{
	return cModuleEvent(
		ptr,
		ctx,
		evt,
		outerror);
}

int initializerregisterrpc(
	const void *ptr,
	NkString id,
	const NkRpcFn fn,
	char **outerror)
{
	return cInitializerRegisterRpc(
		ptr,
		id,
		fn,
		outerror);
}

int initializerregisterbeforert(
	const void *ptr,
	NkString id,
	const NkBeforeRtCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeRt(
		ptr,
		id,
		cb,
		outerror);
}

int initializerregisterafterrt(
	const void *ptr,
	NkString id,
	const NkAfterRtCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterRt(
		ptr,
		id,
		cb,
		outerror);
}

int initializerregistermatchmakermatched(
	const void *ptr,
	const NkMatchmakerMatchedCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterMatchmakerMatched(
		ptr,
		cb,
		outerror);
}

int initializerregistermatch(
	const void *ptr,
	NkString name,
	const NkMatchCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterMatch(
		ptr,
		name,
		cb,
		outerror);
}

int initializerregistertournamentend(
	const void *ptr,
	const NkTournamentCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterTournamentEnd(
		ptr,
		cb,
		outerror);
}

int initializerregistertournamentreset(
	const void *ptr,
	const NkTournamentCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterTournamentReset(
		ptr,
		cb,
		outerror);
}

int initializerregisterleaderboardend(
	const void *ptr,
	const NkLeaderBoardCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterLeaderBoardEnd(
		ptr,
		cb,
		outerror);
}

int initializerregisterleaderboardreset(
	const void *ptr,
	const NkLeaderBoardCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterLeaderBoardReset(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforegetaccount(
	const void *ptr,
	const NkCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeGetAccount(
		ptr,
		cb,
		outerror);
}

int initializerregisteraftergetaccount(
	const void *ptr,
	const NkAfterGetAccountCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterGetAccount(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeupdateaccount(
	const void *ptr,
	const NkBeforeUpdateAccountCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeUpdateAccount(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterupdateaccount(
	const void *ptr,
	const NkAfterUpdateAccountCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterUpdateAccount(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforesessionrefresh(
	const void *ptr,
	const NkBeforeSessionRefreshCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeSessionRefresh(
		ptr,
		cb,
		outerror);
}

int initializerregisteraftersessionrefresh(
	const void *ptr,
	const NkAfterSessionRefreshCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterSessionRefresh(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticateapple(
	const void *ptr,
	const NkBeforeAuthenticateAppleCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeAuthenticateApple(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticateapple(
	const void *ptr,
	const NkAfterAuthenticateAppleCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterAuthenticateApple(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticatecustom(
	const void *ptr,
	const NkBeforeAuthenticateCustomCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeAuthenticateCustom(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticatecustom(
	const void *ptr,
	const NkAfterAuthenticateCustomCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterAuthenticateCustom(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticatedevice(
	const void *ptr,
	const NkBeforeAuthenticateDeviceCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeAuthenticateDevice(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticatedevice(
	const void *ptr,
	const NkAfterAuthenticateDeviceCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterAuthenticateDevice(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticateemail(
	const void *ptr,
	const NkBeforeAuthenticateEmailCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeAuthenticateEmail(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticateemail(
	const void *ptr,
	const NkAfterAuthenticateEmailCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterAuthenticateEmail(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticatefacebook(
	const void *ptr,
	const NkBeforeAuthenticateFacebookCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeAuthenticateFacebook(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticatefacebook(
	const void *ptr,
	const NkAfterAuthenticateFacebookCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterAuthenticateFacebook(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticatefacebookinstantgame(
	const void *ptr,
	const NkBeforeAuthenticateFacebookInstantGameCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeAuthenticateFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticatefacebookinstantgame(
	const void *ptr,
	const NkAfterAuthenticateFacebookInstantGameCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterAuthenticateFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticategamecenter(
	const void *ptr,
	const NkBeforeAuthenticateGameCenterCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeAuthenticateGameCenter(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticategamecenter(
	const void *ptr,
	const NkAfterAuthenticateGameCenterCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterAuthenticateGameCenter(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticategoogle(
	const void *ptr,
	const NkBeforeAuthenticateGoogleCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeAuthenticateGoogle(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticategoogle(
	const void *ptr,
	const NkAfterAuthenticateGoogleCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterAuthenticateGoogle(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticatesteam(
	const void *ptr,
	const NkBeforeAuthenticateSteamCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeAuthenticateSteam(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticatesteam(
	const void *ptr,
	const NkAfterAuthenticateSteamCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterAuthenticateSteam(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelistchannelmessages(
	const void *ptr,
	const NkBeforeListChannelMessagesCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeListChannelMessages(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistchannelmessages(
	const void *ptr,
	const NkAfterListChannelMessagesCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterListChannelMessages(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelistfriends(
	const void *ptr,
	const NkBeforeListFriendsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeListFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistfriends(
	const void *ptr,
	const NkAfterListFriendsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterListFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeaddfriends(
	const void *ptr,
	const NkBeforeAddFriendsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeAddFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterafteraddfriends(
	const void *ptr,
	const NkAfterAddFriendsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterAddFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforedeletefriends(
	const void *ptr,
	const NkBeforeDeleteFriendsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeDeleteFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterdeletefriends(
	const void *ptr,
	const NkAfterDeleteFriendsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterDeleteFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeblockfriends(
	const void *ptr,
	const NkBeforeBlockFriendsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeBlockFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterblockfriends(
	const void *ptr,
	const NkAfterBlockFriendsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterBlockFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeimportfacebookfriends(
	const void *ptr,
	const NkBeforeImportFacebookFriendsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeImportFacebookFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterimportfacebookfriends(
	const void *ptr,
	const NkAfterImportFacebookFriendsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterImportFacebookFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforecreategroup(
	const void *ptr,
	const NkBeforeCreateGroupCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeCreateGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisteraftercreategroup(
	const void *ptr,
	const NkAfterCreateGroupCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterCreateGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeupdategroup(
	const void *ptr,
	const NkBeforeUpdateGroupCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeUpdateGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterupdategroup(
	const void *ptr,
	const NkAfterUpdateGroupCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterUpdateGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforedeletegroup(
	const void *ptr,
	const NkBeforeDeleteGroupCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeDeleteGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterdeletegroup(
	const void *ptr,
	const NkAfterDeleteGroupCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterDeleteGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforejoingroup(
	const void *ptr,
	const NkBeforeJoinGroupCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeJoinGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterjoingroup(
	const void *ptr,
	const NkAfterJoinGroupCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterJoinGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeleavegroup(
	const void *ptr,
	const NkBeforeLeaveGroupCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeLeaveGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterleavegroup(
	const void *ptr,
	const NkAfterLeaveGroupCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterLeaveGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeaddgroupusers(
	const void *ptr,
	const NkBeforeAddGroupUsersCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeAddGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterafteraddgroupusers(
	const void *ptr,
	const NkAfterAddGroupUsersCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterAddGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforebangroupusers(
	const void *ptr,
	const NkBeforeBanGroupUsersCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeBanGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterbangroupusers(
	const void *ptr,
	const NkAfterBanGroupUsersCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterBanGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforekickgroupusers(
	const void *ptr,
	const NkBeforeKickGroupUsersCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeKickGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterkickgroupusers(
	const void *ptr,
	const NkAfterKickGroupUsersCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterKickGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforepromotegroupusers(
	const void *ptr,
	const NkBeforePromoteGroupUsersCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforePromoteGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterpromotegroupusers(
	const void *ptr,
	const NkAfterPromoteGroupUsersCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterPromoteGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforedemotegroupusers(
	const void *ptr,
	const NkBeforeDemoteGroupUsersCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeDemoteGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterdemotegroupusers(
	const void *ptr,
	const NkAfterDemoteGroupUsersCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterDemoteGroupUsers(
		ptr,
		cb,
		outerror);
}


int initializerregisterbeforelistgroupusers(
	const void *ptr,
	const NkBeforeListGroupUsersCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeListGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistgroupusers(
	const void *ptr,
	const NkAfterListGroupUsersCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterListGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelistusergroups(
	const void *ptr,
	const NkBeforeListUserGroupsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeListUserGroups(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistusergroups(
	const void *ptr,
	const NkAfterListUserGroupsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterListUserGroups(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelistgroups(
	const void *ptr,
	const NkBeforeListGroupsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeListGroups(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistgroups(
	const void *ptr,
	const NkAfterListGroupsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterListGroups(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforedeleteleaderboardrecord(
	const void *ptr,
	const NkBeforeDeleteLeaderboardRecordCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeDeleteLeaderboardRecord(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterdeleteleaderboardrecord(
	const void *ptr,
	const NkAfterDeleteLeaderboardRecordCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterDeleteLeaderboardRecord(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelistleaderboardrecords(
	const void *ptr,
	const NkBeforeListLeaderboardRecordsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeListLeaderboardRecords(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistleaderboardrecords(
	const void *ptr,
	const NkAfterListLeaderboardRecordsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterListLeaderboardRecords(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforewriteleaderboardrecord(
	const void *ptr,
	const NkBeforeWriteLeaderboardRecordCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeWriteLeaderboardRecord(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterwriteleaderboardrecord(
	const void *ptr,
	const NkAfterWriteLeaderboardRecordCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterWriteLeaderboardRecord(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelistleaderboardrecordsaroundowner(
	const void *ptr,
	const NkBeforeListLeaderboardRecordsAroundOwnerCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeListLeaderboardRecordsAroundOwner(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistleaderboardrecordsaroundowner(
	const void *ptr,
	const NkAfterListLeaderboardRecordsAroundOwnerCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterListLeaderboardRecordsAroundOwner(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinkapple(
	const void *ptr,
	const NkBeforeLinkAppleCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeLinkApple(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinkapple(
	const void *ptr,
	const NkAfterLinkAppleCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterLinkApple(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinkcustom(
	const void *ptr,
	const NkBeforeLinkCustomCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeLinkCustom(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinkcustom(
	const void *ptr,
	const NkAfterLinkCustomCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterLinkCustom(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinkdevice(
	const void *ptr,
	const NkBeforeLinkDeviceCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeLinkDevice(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinkdevice(
	const void *ptr,
	const NkAfterLinkDeviceCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterLinkDevice(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinkemail(
	const void *ptr,
	const NkBeforeLinkEmailCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeLinkEmail(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinkemail(
	const void *ptr,
	const NkAfterLinkEmailCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterLinkEmail(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinkfacebook(
	const void *ptr,
	const NkBeforeLinkFacebookCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeLinkFacebook(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinkfacebook(
	const void *ptr,
	const NkAfterLinkFacebookCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterLinkFacebook(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinkfacebookinstantgame(
	const void *ptr,
	const NkBeforeLinkFacebookInstantGameCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeLinkFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinkfacebookinstantgame(
	const void *ptr,
	const NkAfterLinkFacebookInstantGameCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterLinkFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinkgamecenter(
	const void *ptr,
	const NkBeforeLinkGameCenterCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeLinkGameCenter(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinkgamecenter(
	const void *ptr,
	const NkAfterLinkGameCenterCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterLinkGameCenter(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinkgoogle(
	const void *ptr,
	const NkBeforeLinkGoogleCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeLinkGoogle(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinkgoogle(
	const void *ptr,
	const NkAfterLinkGoogleCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterLinkGoogle(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinksteam(
	const void *ptr,
	const NkBeforeLinkSteamCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeLinkSteam(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinksteam(
	const void *ptr,
	const NkAfterLinkSteamCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterLinkSteam(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelistmatches(
	const void *ptr,
	const NkBeforeListMatchesCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeListMatches(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistmatches(
	const void *ptr,
	const NkAfterListMatchesCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterListMatches(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelistnotifications(
	const void *ptr,
	const NkBeforeListNotificationsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeListNotifications(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistnotifications(
	const void *ptr,
	const NkAfterListNotificationsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterListNotifications(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforedeletenotifications(
	const void *ptr,
	const NkBeforeDeleteNotificationsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeDeleteNotifications(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterdeletenotifications(
	const void *ptr,
	const NkAfterDeleteNotificationsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterDeleteNotifications(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeliststorageobjects(
	const void *ptr,
	const NkBeforeListStorageObjectsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeListStorageObjects(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterliststorageobjects(
	const void *ptr,
	const NkAfterListStorageObjectsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterListStorageObjects(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforereadstorageobjects(
	const void *ptr,
	const NkBeforeReadStorageObjectsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeReadStorageObjects(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterreadstorageobjects(
	const void *ptr,
	const NkAfterReadStorageObjectsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterReadStorageObjects(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforewritestorageobjects(
	const void *ptr,
	const NkBeforeWriteStorageObjectsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeWriteStorageObjects(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterwritestorageobjects(
	const void *ptr,
	const NkAfterWriteStorageObjectsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterWriteStorageObjects(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforedeletestorageobjects(
	const void *ptr,
	const NkBeforeDeleteStorageObjectsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeDeleteStorageObjects(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterdeletestorageobjects(
	const void *ptr,
	const NkAfterDeleteStorageObjectsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterDeleteStorageObjects(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforejointournament(
	const void *ptr,
	const NkBeforeJoinTournamentCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeJoinTournament(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterjointournament(
	const void *ptr,
	const NkAfterJoinTournamentCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterJoinTournament(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelisttournamentrecords(
	const void *ptr,
	const NkBeforeListTournamentRecordsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeListTournamentRecords(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlisttournamentrecords(
	const void *ptr,
	const NkAfterListTournamentRecordsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterListTournamentRecords(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelisttournaments(
	const void *ptr,
	const NkBeforeListTournamentsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeListTournaments(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlisttournaments(
	const void *ptr,
	const NkAfterListTournamentsCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterListTournaments(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforewritetournamentrecord(
	const void *ptr,
	const NkBeforeWriteTournamentRecordCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeWriteTournamentRecord(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterwritetournamentrecord(
	const void *ptr,
	const NkAfterWriteTournamentRecordCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterWriteTournamentRecord(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelisttournamentrecordsaroundowner(
	const void *ptr,
	const NkBeforeListTournamentRecordsAroundOwnerCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeListTournamentRecordsAroundOwner(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlisttournamentrecordsaroundowner(
	const void *ptr,
	const NkAfterListTournamentRecordsAroundOwnerCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterListTournamentRecordsAroundOwner(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinkapple(
	const void *ptr,
	const NkBeforeUnlinkAppleCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeUnlinkApple(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinkapple(
	const void *ptr,
	const NkAfterUnlinkAppleCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterUnlinkApple(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinkcustom(
	const void *ptr,
	const NkBeforeUnlinkCustomCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeUnlinkCustom(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinkcustom(
	const void *ptr,
	const NkAfterUnlinkCustomCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterUnlinkCustom(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinkdevice(
	const void *ptr,
	const NkBeforeUnlinkDeviceCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeUnlinkDevice(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinkdevice(
	const void *ptr,
	const NkAfterUnlinkDeviceCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterUnlinkDevice(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinkemail(
	const void *ptr,
	const NkBeforeUnlinkEmailCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeUnlinkEmail(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinkemail(
	const void *ptr,
	const NkAfterUnlinkEmailCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterUnlinkEmail(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinkfacebook(
	const void *ptr,
	const NkBeforeUnlinkFacebookCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeUnlinkFacebook(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinkfacebook(
	const void *ptr,
	const NkAfterUnlinkFacebookCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterUnlinkFacebook(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinkfacebookinstantgame(
	const void *ptr,
	const NkBeforeUnlinkFacebookInstantGameCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeUnlinkFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinkfacebookinstantgame(
	const void *ptr,
	const NkAfterUnlinkFacebookInstantGameCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterUnlinkFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinkgamecenter(
	const void *ptr,
	const NkBeforeUnlinkGameCenterCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeUnlinkGameCenter(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinkgamecenter(
	const void *ptr,
	const NkAfterUnlinkGameCenterCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterUnlinkGameCenter(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinkgoogle(
	const void *ptr,
	const NkBeforeUnlinkGoogleCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeUnlinkGoogle(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinkgoogle(
	const void *ptr,
	const NkAfterUnlinkGoogleCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterUnlinkGoogle(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinksteam(
	const void *ptr,
	const NkBeforeUnlinkSteamCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeUnlinkSteam(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinksteam(
	const void *ptr,
	const NkAfterUnlinkSteamCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterUnlinkSteam(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforegetusers(
	const void *ptr,
	const NkBeforeGetUsersCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterBeforeGetUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisteraftergetusers(
	const void *ptr,
	const NkAfterGetUsersCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterAfterGetUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterevent(
	const void *ptr,
	const NkEventCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterEvent(
		ptr,
		cb,
		outerror);
}

int initializerregistereventsessionstart(
	const void *ptr,
	const NkEventCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterEventSessionStart(
		ptr,
		cb,
		outerror);
}

int initializerregistereventsessionend(
	const void *ptr,
	const NkEventCallbackFn cb,
	char **outerror)
{
	return cInitializerRegisterEventSessionEnd(
		ptr,
		cb,
		outerror);
}

*/
import "C"

func GoStringN(s C.NkString) string {
	return C.GoStringN(s.p, C.int(s.n))
}
