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
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
	bool **outcreated);

extern int cModuleAuthenticateCustom(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
	bool **outcreated);

extern int cModuleAuthenticateDevice(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
	bool **outcreated);

extern int cModuleAuthenticateEmail(
	const void *ptr,
	const NkContext *ctx,
	NkString email,
	NkString password,
	NkString username,
	bool create,
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
	bool **outcreated);

extern int cModuleAuthenticateFacebook(
	const void *ptr,
	const NkContext *ctx,
	NkString token,
	bool importfriends,
	NkString username,
	bool create,
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
	bool **outcreated);

extern int cModuleAuthenticateFacebookInstantGame(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
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
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
	bool **outcreated);

extern int cModuleAuthenticateGoogle(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
	bool **outcreated);

extern int cModuleAuthenticateSteam(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
	bool **outcreated);

extern int cModuleAuthenticateGenerateToken(
	const void *ptr,
	NkString userid,
	NkString username,
	NkI64 expiry,
	NkMapString vars,
	NkString **outtoken,
	NkI64 **outexpiry,
	NkString **outerror);

extern int cModuleAccountGetId(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkAccount **outaccount,
	NkString **outerror);

extern int cModuleAccountsGetId(
	const void *ptr,
	const NkContext *ctx,
	NkString *userids,
	NkU32 numuserids,
	NkAccount **outaccounts,
	NkU32 **outnumaccounts,
	NkString **outerror);

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
	NkString **outerror);

extern int cModuleAccountDeleteId(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	bool recorded,
	NkString **outerror);

extern int cModuleAccountExportId(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString **outaccount,
	NkString **outerror);

extern int cModuleUsersGet(
	const void *ptr,
	const NkContext *ctx,
	NkString *keys,
	NkU32 numkeys,
	NkUser **outusers,
	NkU32 **outnumusers,
	NkString **outerror);

extern int cModuleUsersBan(
	const void *ptr,
	const NkContext *ctx,
	NkString *userids,
	NkU32 numids,
	NkString **outerror);

extern int cModuleLink(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);

extern int cModuleLinkEmail(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString email,
	NkString password,
	NkString **outerror);

extern int cModuleLinkFacebook(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	NkString token,
	bool importfriends,
	NkString **outerror);

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
	NkString **outerror);

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
	NkString **outerror);

extern int cModuleStreamUserGet(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkString userid,
	NkString sessionid,
	NkPresenceMeta **outmeta,
	NkString **outerror);

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
	NkString **outerror);

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
	NkString **outerror);

extern int cModuleStreamUserLeave(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkString userid,
	NkString sessionid,
	NkString **outerror);

extern int cModuleStreamUserKick(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkPresence presence,
	NkString **outerror);

extern int cModuleStreamCount(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkU64 **outcount,
	NkString **outerror);

extern int cModuleStreamClose(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkString **outerror);

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
	NkString **outerror);

extern int cModuleStreamSendRaw(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkEnvelope msg,
	NkPresence *presences,
	NkU32 numpresences,
	bool reliable,
	NkString **outerror);

extern int cModuleSessionDisconnect(
	const void *ptr,
	const NkContext *ctx,
	NkString sessionid,
	NkString **outerror);

extern int cModuleMatchCreate(
	const void *ptr,
	const NkContext *ctx,
	NkString module,
	NkMapAny params,
	NkString **outmatchid,
	NkString **outerror);

extern int cModuleMatchGet(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkMatch **outmatch,
	NkString **outerror);

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
	NkString **outerror);

extern int cModuleNotificationSend(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString subject,
	NkMapAny content,
	NkU64 code,
	NkString sender,
	bool persistent,
	NkString **outerror);

extern int cModuleNotificationsSend(
	const void *ptr,
	const NkContext *ctx,
	const NkNotificationSend *notifications,
	NkU32 numnotifications,
	NkString **outerror);

extern int cModuleWalletUpdate(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkMapI64 changeset,
	NkMapAny metadata,
	bool updateledger,
	NkMapI64 **outupdated,
	NkMapI64 **outprevious,
	NkString **outerror);

extern int cModuleWalletsUpdate(
	const void *ptr,
	const NkContext *ctx,
	const NkWalletUpdate *updates,
	NkU32 numupdates,
	bool updateledger,
	NkWalletUpdateResult **outresults,
	NkU32 **outnumresults,
	NkString **outerror);

extern int cModuleWalletLedgerUpdate(
	const void *ptr,
	const NkContext *ctx,
	NkString itemid,
	NkMapAny metadata,
	NkWalletLedgerItem **outitem,
	NkString **outerror);

extern int cModuleWalletLedgerList(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkU32 limit,
	NkString cursor,
	NkWalletLedgerItem **outitems,
	NkU32 **outnumitems,
	NkString **outcursor,
	NkString **outerror);

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
	NkString **outerror);

extern int cModuleStorageRead(
	const void *ptr,
	const NkContext *ctx,
	const NkStorageRead *reads,
	NkU32 numreads,
	NkStorageObject **outobjs,
	NkU32 **outnumobjs,
	NkString **outerror);

extern int cModuleStorageWrite(
	const void *ptr,
	const NkContext *ctx,
	const NkStorageWrite *writes,
	NkU32 numwrites,
	NkStorageObjectAck **outacks,
	NkU32 **outnumacks,
	NkString **outerror);

extern int cModuleStorageDelete(
	const void *ptr,
	const NkContext *ctx,
	const NkStorageDelete *deletes,
	NkU32 numdeletes,
	NkString **outerror);

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
	NkString **outerror);

extern int cModuleLeaderboardCreate(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	bool authoritative,
	NkString sortorder,
	NkString op,
	NkString resetschedule,
	NkMapAny metadata,
	NkString **outerror);

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
	NkString **outerror);

extern int cModuleLeaderboardRecordWrite(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkI64 score,
	NkI64 subscore,
	NkMapAny metadata,
	NkLeaderboardRecord **outrecord,
	NkString **outerror);

extern int cModuleDelete(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkString **outerror);

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
	NkString **outerror);

extern int cModuleTournamentAddAttempt(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkU64 count,
	NkString **outerror);

extern int cModuleTournamentJoin(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkString username,
	NkString **outerror);

extern int cModuleTournamentsGetId(
	const void *ptr,
	const NkContext *ctx,
	const NkString *tournamentids,
	NkU32 numtournamentids,
	NkTournament **outtournaments,
	NkU32 **outnumtournaments,
	NkString **outerror);

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
	NkString **outerror);

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
	NkString **outerror);

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
	NkString **outerror);

extern int cModuleTournamentRecordsHaystack(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkU32 limit,
	NkI64 expiry,
	NkLeaderboardRecord **outrecords,
	NkU32 **outnumrecords,
	NkString **outerror);

extern int cModuleGroupsGetId(
	const void *ptr,
	const NkContext *ctx,
	const NkString *groupids,
	NkU32 numgroupids,
	NkGroup **outgroups,
	NkU32 **outnumgroups,
	NkString **outerror);

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
	NkString **outerror);

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
	NkString **outerror);

extern int cModuleGroupuUser(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	NkString userid,
	NkString username,
	NkString **outerror);

extern int cModuleGroupUsers(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	const NkString *userids,
	NkU32 numuserids,
	NkString **outerror);

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
	NkString **outerror);

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
	NkString **outerror);

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
	NkString **outerror);

extern int cModuleEvent(
	const void *ptr,
	const NkContext *ctx,
	NkEvent evt,
	NkString **outerror);

extern int cInitializerRpc(
	const void *ptr,
	NkString id,
	const NkRpcFn fn,
	NkString **outerror);

extern int cInitializerBeforeRt(
	const void *ptr,
	NkString id,
	const NkBeforeRtCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterRt(
	const void *ptr,
	NkString id,
	const NkAfterRtCallbackFn cb,
	NkString **outerror);

extern int cInitializerMatchmakerMatched(
	const void *ptr,
	const NkMatchmakerMatchedCallbackFn cb,
	NkString **outerror);

extern int cInitializerMatch(
	const void *ptr,
	NkString name,
	const NkMatchCallbackFn cb,
	NkString **outerror);

extern int cInitializerTournament(
	const void *ptr,
	const NkTournamentCallbackFn cb,
	NkString **outerror);

extern int cInitializerLeaderBoardEnd(
	const void *ptr,
	const NkLeaderBoardCallbackFn cb,
	NkString **outerror);

extern int cInitializerLeaderBoardReset(
	const void *ptr,
	const NkLeaderBoardCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeGetAccount(
	const void *ptr,
	const NkCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterGetAccount(
	const void *ptr,
	const NkAfterGetAccountCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeUpdateAccount(
	const void *ptr,
	const NkBeforeUpdateAccountCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterUpdateAccount(
	const void *ptr,
	const NkAfterUpdateAccountCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeSessionRefresh(
	const void *ptr,
	const NkBeforeSessionRefreshCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterSessionRefresh(
	const void *ptr,
	const NkAfterSessionRefreshCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeAuthenticateApple(
	const void *ptr,
	const NkBeforeAuthenticateAppleCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterAuthenticateApple(
	const void *ptr,
	const NkAfterAuthenticateAppleCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeAuthenticateCustom(
	const void *ptr,
	const NkBeforeAuthenticateCustomCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterAuthenticateCustom(
	const void *ptr,
	const NkAfterAuthenticateCustomCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeAuthenticateDevice(
	const void *ptr,
	const NkBeforeAuthenticateDeviceCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterAuthenticateDevice(
	const void *ptr,
	const NkAfterAuthenticateDeviceCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeAuthenticateEmail(
	const void *ptr,
	const NkBeforeAuthenticateEmailCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterAuthenticateEmail(
	const void *ptr,
	const NkAfterAuthenticateEmailCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeAuthenticateFacebook(
	const void *ptr,
	const NkBeforeAuthenticateFacebookCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterAuthenticateFacebook(
	const void *ptr,
	const NkAfterAuthenticateFacebookCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeAuthenticateFacebookInstantGame(
	const void *ptr,
	const NkBeforeAuthenticateFacebookInstantGameCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterAuthenticateFacebookInstantGame(
	const void *ptr,
	const NkAfterAuthenticateFacebookInstantGameCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeAuthenticateGameCenter(
	const void *ptr,
	const NkBeforeAuthenticateGameCenterCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterAuthenticateGameCenter(
	const void *ptr,
	const NkAfterAuthenticateGameCenterCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeAuthenticateGoogle(
	const void *ptr,
	const NkBeforeAuthenticateGoogleCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterAuthenticateGoogle(
	const void *ptr,
	const NkAfterAuthenticateGoogleCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeAuthenticateSteam(
	const void *ptr,
	const NkBeforeAuthenticateSteamCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterAuthenticateSteam(
	const void *ptr,
	const NkAfterAuthenticateSteamCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeListChannelMessages(
	const void *ptr,
	const NkBeforeListChannelMessagesCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterListChannelMessages(
	const void *ptr,
	const NkAfterListChannelMessagesCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeListFriends(
	const void *ptr,
	const NkBeforeListFriendsCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterListFriends(
	const void *ptr,
	const NkAfterListFriendsCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeAddFriends(
	const void *ptr,
	const NkBeforeAddFriendsCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterAddFriends(
	const void *ptr,
	const NkAfterAddFriendsCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeDeleteFriends(
	const void *ptr,
	const NkBeforeDeleteFriendsCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterDeleteFriends(
	const void *ptr,
	const NkAfterDeleteFriendsCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeBlockFriends(
	const void *ptr,
	const NkBeforeBlockFriendsCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterBlockFriends(
	const void *ptr,
	const NkAfterBlockFriendsCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeImportFacebookFriends(
	const void *ptr,
	const NkBeforeImportFacebookFriendsCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterImportFacebookFriends(
	const void *ptr,
	const NkAfterImportFacebookFriendsCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeCreateGroup(
	const void *ptr,
	const NkBeforeCreateGroupCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterCreateGroup(
	const void *ptr,
	const NkAfterCreateGroupCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeUpdateGroup(
	const void *ptr,
	const NkBeforeUpdateGroupCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterUpdateGroup(
	const void *ptr,
	const NkAfterUpdateGroupCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeDeleteGroup(
	const void *ptr,
	const NkBeforeDeleteGroupCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterDeleteGroup(
	const void *ptr,
	const NkAfterDeleteGroupCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeJoinGroup(
	const void *ptr,
	const NkBeforeJoinGroupCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterJoinGroup(
	const void *ptr,
	const NkAfterJoinGroupCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeLeaveGroup(
	const void *ptr,
	const NkBeforeLeaveGroupCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterLeaveGroup(
	const void *ptr,
	const NkAfterLeaveGroupCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeAddGroupUsers(
	const void *ptr,
	const NkBeforeAddGroupUsersCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterAddGroupUsers(
	const void *ptr,
	const NkAfterAddGroupUsersCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeBanGroupUsers(
	const void *ptr,
	const NkBeforeBanGroupUsersCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterBanGroupUsers(
	const void *ptr,
	const NkAfterBanGroupUsersCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeKickGroupUsers(
	const void *ptr,
	const NkBeforeKickGroupUsersCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterKickGroupUsers(
	const void *ptr,
	const NkAfterKickGroupUsersCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforePromoteGroupUsers(
	const void *ptr,
	const NkBeforePromoteGroupUsersCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterPromoteGroupUsers(
	const void *ptr,
	const NkAfterPromoteGroupUsersCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeDemoteGroupUsers(
	const void *ptr,
	const NkBeforeDemoteGroupUsersCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterDemoteGroupUsers(
	const void *ptr,
	const NkAfterDemoteGroupUsersCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeListGroupUsers(
	const void *ptr,
	const NkBeforeListGroupUsersCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterListGroupUsers(
	const void *ptr,
	const NkAfterListGroupUsersCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeListUserGroups(
	const void *ptr,
	const NkBeforeListUserGroupsCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterListUserGroups(
	const void *ptr,
	const NkAfterListUserGroupsCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeListGroups(
	const void *ptr,
	const NkBeforeListGroupsCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterListGroups(
	const void *ptr,
	const NkAfterListGroupsCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeDeleteLeaderboardRecord(
	const void *ptr,
	const NkBeforeDeleteLeaderboardRecordCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterDeleteLeaderboardRecord(
	const void *ptr,
	const NkAfterDeleteLeaderboardRecordCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeListLeaderboardRecords(
	const void *ptr,
	const NkBeforeListLeaderboardRecordsCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterListLeaderboardRecords(
	const void *ptr,
	const NkAfterListLeaderboardRecordsCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeWriteLeaderboardRecord(
	const void *ptr,
	const NkBeforeWriteLeaderboardRecordCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterWriteLeaderboardRecord(
	const void *ptr,
	const NkAfterWriteLeaderboardRecordCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeListLeaderboardRecordsAroundOwner(
	const void *ptr,
	const NkBeforeListLeaderboardRecordsAroundOwnerCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterListLeaderboardRecordsAroundOwner(
	const void *ptr,
	const NkAfterListLeaderboardRecordsAroundOwnerCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeLinkApple(
	const void *ptr,
	const NkBeforeLinkAppleCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterLinkApple(
	const void *ptr,
	const NkAfterLinkAppleCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeLinkCustom(
	const void *ptr,
	const NkBeforeLinkCustomCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterLinkCustom(
	const void *ptr,
	const NkAfterLinkCustomCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeLinkDevice(
	const void *ptr,
	const NkBeforeLinkDeviceCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterLinkDevice(
	const void *ptr,
	const NkAfterLinkDeviceCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeLinkEmail(
	const void *ptr,
	const NkBeforeLinkEmailCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterLinkEmail(
	const void *ptr,
	const NkAfterLinkEmailCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeLinkFacebook(
	const void *ptr,
	const NkBeforeLinkFacebookCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterLinkFacebook(
	const void *ptr,
	const NkAfterLinkFacebookCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeLinkFacebookInstantGame(
	const void *ptr,
	const NkBeforeLinkFacebookInstantGameCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterLinkFacebookInstantGame(
	const void *ptr,
	const NkAfterLinkFacebookInstantGameCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeLinkGameCenter(
	const void *ptr,
	const NkBeforeLinkGameCenterCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterLinkGameCenter(
	const void *ptr,
	const NkAfterLinkGameCenterCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeLinkGoogle(
	const void *ptr,
	const NkBeforeLinkGoogleCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterLinkGoogle(
	const void *ptr,
	const NkAfterLinkGoogleCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeLinkSteam(
	const void *ptr,
	const NkBeforeLinkSteamCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterLinkSteam(
	const void *ptr,
	const NkAfterLinkSteamCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeListMatches(
	const void *ptr,
	const NkBeforeListMatchesCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterListMatches(
	const void *ptr,
	const NkAfterListMatchesCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeListNotifications(
	const void *ptr,
	const NkBeforeListNotificationsCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterListNotifications(
	const void *ptr,
	const NkAfterListNotificationsCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeDeleteNotifications(
	const void *ptr,
	const NkBeforeDeleteNotificationsCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterDeleteNotifications(
	const void *ptr,
	const NkAfterDeleteNotificationsCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeListStorageObjects(
	const void *ptr,
	const NkBeforeListStorageObjectsCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterListStorageObjects(
	const void *ptr,
	const NkAfterListStorageObjectsCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeReadStorageObjects(
	const void *ptr,
	const NkBeforeReadStorageObjectsCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterReadStorageObjects(
	const void *ptr,
	const NkAfterReadStorageObjectsCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeWriteStorageObjects(
	const void *ptr,
	const NkBeforeWriteStorageObjectsCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterWriteStorageObjects(
	const void *ptr,
	const NkAfterWriteStorageObjectsCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeDeleteStorageObjects(
	const void *ptr,
	const NkBeforeDeleteStorageObjectsCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterDeleteStorageObjects(
	const void *ptr,
	const NkAfterDeleteStorageObjectsCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeJoinTournament(
	const void *ptr,
	const NkBeforeJoinTournamentCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterJoinTournament(
	const void *ptr,
	const NkAfterJoinTournamentCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeListTournamentRecords(
	const void *ptr,
	const NkBeforeListTournamentRecordsCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterListTournamentRecords(
	const void *ptr,
	const NkAfterListTournamentRecordsCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeListTournaments(
	const void *ptr,
	const NkBeforeListTournamentsCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterListTournaments(
	const void *ptr,
	const NkAfterListTournamentsCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeWriteTournamentRecord(
	const void *ptr,
	const NkBeforeWriteTournamentRecordCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterWriteTournamentRecord(
	const void *ptr,
	const NkAfterWriteTournamentRecordCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeListTournamentRecordsAroundOwner(
	const void *ptr,
	const NkBeforeListTournamentRecordsAroundOwnerCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterListTournamentRecordsAroundOwner(
	const void *ptr,
	const NkAfterListTournamentRecordsAroundOwnerCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeUnlinkApple(
	const void *ptr,
	const NkBeforeUnlinkAppleCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterUnlinkApple(
	const void *ptr,
	const NkAfterUnlinkAppleCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeUnlinkCustom(
	const void *ptr,
	const NkBeforeUnlinkCustomCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterUnlinkCustom(
	const void *ptr,
	const NkAfterUnlinkCustomCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeUnlinkDevice(
	const void *ptr,
	const NkBeforeUnlinkDeviceCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterUnlinkDevice(
	const void *ptr,
	const NkAfterUnlinkDeviceCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeUnlinkEmail(
	const void *ptr,
	const NkBeforeUnlinkEmailCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterUnlinkEmail(
	const void *ptr,
	const NkAfterUnlinkEmailCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeUnlinkFacebook(
	const void *ptr,
	const NkBeforeUnlinkFacebookCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterUnlinkFacebook(
	const void *ptr,
	const NkAfterUnlinkFacebookCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeUnlinkFacebookInstantGame(
	const void *ptr,
	const NkBeforeUnlinkFacebookInstantGameCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterUnlinkFacebookInstantGame(
	const void *ptr,
	const NkAfterUnlinkFacebookInstantGameCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeUnlinkGameCenter(
	const void *ptr,
	const NkBeforeUnlinkGameCenterCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterUnlinkGameCenter(
	const void *ptr,
	const NkAfterUnlinkGameCenterCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeUnlinkGoogle(
	const void *ptr,
	const NkBeforeUnlinkGoogleCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterUnlinkGoogle(
	const void *ptr,
	const NkAfterUnlinkGoogleCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeUnlinkSteam(
	const void *ptr,
	const NkBeforeUnlinkSteamCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterUnlinkSteam(
	const void *ptr,
	const NkAfterUnlinkSteamCallbackFn cb,
	NkString **outerror);

extern int cInitializerBeforeGetUsers(
	const void *ptr,
	const NkBeforeGetUsersCallbackFn cb,
	NkString **outerror);

extern int cInitializerAfterGetUsers(
	const void *ptr,
	const NkAfterGetUsersCallbackFn cb,
	NkString **outerror);

extern int cInitializerEvent(
	const void *ptr,
	const NkEventCallbackFn cb,
	NkString **outerror);

extern int cInitializerEventSessionStart(
	const void *ptr,
	const NkEventCallbackFn cb,
	NkString **outerror);

extern int cInitializerEventSessionEnd(
	const void *ptr,
	const NkEventCallbackFn cb,
	NkString **outerror);

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
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
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
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
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
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
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
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
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
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
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
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
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
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
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
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
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
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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

int ModuleAccountDeleteId(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		bool recorded,
		NkString **outerror)
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
	NkString **outerror)
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
	NkString *keys,
	NkU32 numkeys,
	NkUser **outusers,
	NkU32 **outnumusers,
	NkString **outerror)
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
	NkString *keys,
	NkU32 numkeys,
	NkUser **outusers,
	NkU32 **outnumusers,
	NkString **outerror)
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
	NkString *userids,
	NkU32 numids,
	NkString **outerror)
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
	NkString *userids,
	NkU32 numids,
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
{
	return c(
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkPresence *presences,
	NkU32 numpresences,
	bool reliable,
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
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
	NkString **outerror)
{
	return cInitializerRpc(
		ptr,
		id,
		fn,
		outerror);
}

int initializerregisterbeforert(
	const void *ptr,
	NkString id,
	const NkBeforeRtCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeRt(
		ptr,
		id,
		cb,
		outerror);
}

int initializerregisterafterrt(
	const void *ptr,
	NkString id,
	const NkAfterRtCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterRt(
		ptr,
		id,
		cb,
		outerror);
}

int initializerregistermatchmakermatched(
	const void *ptr,
	const NkMatchmakerMatchedCallbackFn cb,
	NkString **outerror)
{
	return cInitializerMatchmakerMatched(
		ptr,
		cb,
		outerror);
}

int initializerregistermatch(
	const void *ptr,
	NkString name,
	const NkMatchCallbackFn cb,
	NkString **outerror)
{
	return cInitializerMatch(
		ptr,
		name,
		cb,
		outerror);
}

int initializerregistertournamentend(
	const void *ptr,
	const NkTournamentCallbackFn cb,
	NkString **outerror)
{
	return cInitializerTournamentEnd(
		ptr,
		cb,
		outerror);
}

int initializerregistertournamentreset(
	const void *ptr,
	const NkTournamentCallbackFn cb,
	NkString **outerror)
{
	return cInitializerTournamentReset(
		ptr,
		cb,
		outerror);
}

int initializerregisterleaderboardend(
	const void *ptr,
	const NkLeaderBoardCallbackFn cb,
	NkString **outerror)
{
	return cInitializerLeaderBoardEnd(
		ptr,
		cb,
		outerror);
}

int initializerregisterleaderboardreset(
	const void *ptr,
	const NkLeaderBoardCallbackFn cb,
	NkString **outerror)
{
	return cInitializerLeaderBoardReset(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforegetaccount(
	const void *ptr,
	const NkCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeGetAccount(
		ptr,
		cb,
		outerror);
}

int initializerregisteraftergetaccount(
	const void *ptr,
	const NkAfterGetAccountCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterGetAccount(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeupdateaccount(
	const void *ptr,
	const NkBeforeUpdateAccountCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeUpdateAccount(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterupdateaccount(
	const void *ptr,
	const NkAfterUpdateAccountCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterUpdateAccount(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforesessionrefresh(
	const void *ptr,
	const NkBeforeSessionRefreshCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeSessionRefresh(
		ptr,
		cb,
		outerror);
}

int initializerregisteraftersessionrefresh(
	const void *ptr,
	const NkAfterSessionRefreshCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterSessionRefresh(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticateapple(
	const void *ptr,
	const NkBeforeAuthenticateAppleCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeAuthenticateApple(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticateapple(
	const void *ptr,
	const NkAfterAuthenticateAppleCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterAuthenticateApple(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticatecustom(
	const void *ptr,
	const NkBeforeAuthenticateCustomCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeAuthenticateCustom(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticatecustom(
	const void *ptr,
	const NkAfterAuthenticateCustomCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterAuthenticateCustom(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticatedevice(
	const void *ptr,
	const NkBeforeAuthenticateDeviceCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeAuthenticateDevice(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticatedevice(
	const void *ptr,
	const NkAfterAuthenticateDeviceCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterAuthenticateDevice(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticateemail(
	const void *ptr,
	const NkBeforeAuthenticateEmailCallbackFn cb,
	NkString **outerror)
{

}

int initializerregisterafterauthenticateemail(
	const void *ptr,
	const NkAfterAuthenticateEmailCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterAuthenticateEmail(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticatefacebook(
	const void *ptr,
	const NkBeforeAuthenticateFacebookCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeAuthenticateFacebook(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticatefacebook(
	const void *ptr,
	const NkAfterAuthenticateFacebookCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterAuthenticateFacebook(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticatefacebookinstantgame(
	const void *ptr,
	const NkBeforeAuthenticateFacebookInstantGameCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeAuthenticateFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticatefacebookinstantgame(
	const void *ptr,
	const NkAfterAuthenticateFacebookInstantGameCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterAuthenticateFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticategamecenter(
	const void *ptr,
	const NkBeforeAuthenticateGameCenterCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeAuthenticateGameCenter(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticategamecenter(
	const void *ptr,
	const NkAfterAuthenticateGameCenterCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterAuthenticateGameCenter(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticategoogle(
	const void *ptr,
	const NkBeforeAuthenticateGoogleCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeAuthenticateGoogle(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticategoogle(
	const void *ptr,
	const NkAfterAuthenticateGoogleCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterAuthenticateGoogle(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeauthenticatesteam(
	const void *ptr,
	const NkBeforeAuthenticateSteamCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeAuthenticateSteam(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterauthenticatesteam(
	const void *ptr,
	const NkAfterAuthenticateSteamCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterAuthenticateSteam(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelistchannelmessages(
	const void *ptr,
	const NkBeforeListChannelMessagesCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeListChannelMessages(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistchannelmessages(
	const void *ptr,
	const NkAfterListChannelMessagesCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterListChannelMessages(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelistfriends(
	const void *ptr,
	const NkBeforeListFriendsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeListFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistfriends(
	const void *ptr,
	const NkAfterListFriendsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterListFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeaddfriends(
	const void *ptr,
	const NkBeforeAddFriendsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeAddFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterafteraddfriends(
	const void *ptr,
	const NkAfterAddFriendsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterAddFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforedeletefriends(
	const void *ptr,
	const NkBeforeDeleteFriendsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeDeleteFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterdeletefriends(
	const void *ptr,
	const NkAfterDeleteFriendsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterDeleteFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeblockfriends(
	const void *ptr,
	const NkBeforeBlockFriendsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeBlockFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterblockfriends(
	const void *ptr,
	const NkAfterBlockFriendsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterBlockFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeimportfacebookfriends(
	const void *ptr,
	const NkBeforeImportFacebookFriendsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeImportFacebookFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterimportfacebookfriends(
	const void *ptr,
	const NkAfterImportFacebookFriendsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterImportFacebookFriends(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforecreategroup(
	const void *ptr,
	const NkBeforeCreateGroupCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeCreateGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisteraftercreategroup(
	const void *ptr,
	const NkAfterCreateGroupCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterCreateGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeupdategroup(
	const void *ptr,
	const NkBeforeUpdateGroupCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeUpdateGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterupdategroup(
	const void *ptr,
	const NkAfterUpdateGroupCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterUpdateGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforedeletegroup(
	const void *ptr,
	const NkBeforeDeleteGroupCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeDeleteGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterdeletegroup(
	const void *ptr,
	const NkAfterDeleteGroupCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterDeleteGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforejoingroup(
	const void *ptr,
	const NkBeforeJoinGroupCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeJoinGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterjoingroup(
	const void *ptr,
	const NkAfterJoinGroupCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterJoinGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeleavegroup(
	const void *ptr,
	const NkBeforeLeaveGroupCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeLeaveGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterleavegroup(
	const void *ptr,
	const NkAfterLeaveGroupCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterLeaveGroup(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeaddgroupusers(
	const void *ptr,
	const NkBeforeAddGroupUsersCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeAddGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterafteraddgroupusers(
	const void *ptr,
	const NkAfterAddGroupUsersCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterAddGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforebangroupusers(
	const void *ptr,
	const NkBeforeBanGroupUsersCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeBanGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterbangroupusers(
	const void *ptr,
	const NkAfterBanGroupUsersCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterBanGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforekickgroupusers(
	const void *ptr,
	const NkBeforeKickGroupUsersCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeKickGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterkickgroupusers(
	const void *ptr,
	const NkAfterKickGroupUsersCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterKickGroupUsers(
		ptr,
		cb,
		outerror);
}

InitializerBeforePromoteGroupUsers(
	const void *ptr,
	const NkBeforePromoteGroupUsersCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforePromoteGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterpromotegroupusers(
	const void *ptr,
	const NkAfterPromoteGroupUsersCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterPromoteGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforedemotegroupusers(
	const void *ptr,
	const NkBeforeDemoteGroupUsersCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeDemoteGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterdemotegroupusers(
	const void *ptr,
	const NkAfterDemoteGroupUsersCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterDemoteGroupUsers(
		ptr,
		cb,
		outerror);
}


int initializerregisterbeforelistgroupusers(
	const void *ptr,
	const NkBeforeListGroupUsersCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeListGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistgroupusers(
	const void *ptr,
	const NkAfterListGroupUsersCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterListGroupUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelistusergroups(
	const void *ptr,
	const NkBeforeListUserGroupsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeListUserGroups(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistusergroups(
	const void *ptr,
	const NkAfterListUserGroupsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterListUserGroups(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelistgroups(
	const void *ptr,
	const NkBeforeListGroupsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeListGroups(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistgroups(
	const void *ptr,
	const NkAfterListGroupsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterListGroups(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforedeleteleaderboardrecord(
	const void *ptr,
	const NkBeforeDeleteLeaderboardRecordCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeDeleteLeaderboardRecord(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterdeleteleaderboardrecord(
	const void *ptr,
	const NkAfterDeleteLeaderboardRecordCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterDeleteLeaderboardRecord(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelistleaderboardrecords(
	const void *ptr,
	const NkBeforeListLeaderboardRecordsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeListLeaderboardRecords(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistleaderboardrecords(
	const void *ptr,
	const NkAfterListLeaderboardRecordsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterListLeaderboardRecords(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforewriteleaderboardrecord(
	const void *ptr,
	const NkBeforeWriteLeaderboardRecordCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeWriteLeaderboardRecord(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterwriteleaderboardrecord(
	const void *ptr,
	const NkAfterWriteLeaderboardRecordCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterWriteLeaderboardRecord(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelistleaderboardrecordsaroundowner(
	const void *ptr,
	const NkBeforeListLeaderboardRecordsAroundOwnerCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeListLeaderboardRecordsAroundOwner(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistleaderboardrecordsaroundowner(
	const void *ptr,
	const NkAfterListLeaderboardRecordsAroundOwnerCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterListLeaderboardRecordsAroundOwner(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinkapple(
	const void *ptr,
	const NkBeforeLinkAppleCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeLinkApple(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinkapple(
	const void *ptr,
	const NkAfterLinkAppleCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterLinkApple(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinkcustom(
	const void *ptr,
	const NkBeforeLinkCustomCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeLinkCustom(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinkcustom(
	const void *ptr,
	const NkAfterLinkCustomCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterLinkCustom(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinkdevice(
	const void *ptr,
	const NkBeforeLinkDeviceCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeLinkDevice(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinkdevice(
	const void *ptr,
	const NkAfterLinkDeviceCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterLinkDevice(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinkemail(
	const void *ptr,
	const NkBeforeLinkEmailCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeLinkEmail(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinkemail(
	const void *ptr,
	const NkAfterLinkEmailCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterLinkEmail(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinkfacebook(
	const void *ptr,
	const NkBeforeLinkFacebookCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeLinkFacebook(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinkfacebook(
	const void *ptr,
	const NkAfterLinkFacebookCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterLinkFacebook(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinkfacebookinstantgame(
	const void *ptr,
	const NkBeforeLinkFacebookInstantGameCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeLinkFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinkfacebookinstantgame(
	const void *ptr,
	const NkAfterLinkFacebookInstantGameCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterLinkFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinkgamecenter(
	const void *ptr,
	const NkBeforeLinkGameCenterCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeLinkGameCenter(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinkgamecenter(
	const void *ptr,
	const NkAfterLinkGameCenterCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterLinkGameCenter(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinkgoogle(
	const void *ptr,
	const NkBeforeLinkGoogleCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeLinkGoogle(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinkgoogle(
	const void *ptr,
	const NkAfterLinkGoogleCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterLinkGoogle(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelinksteam(
	const void *ptr,
	const NkBeforeLinkSteamCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeLinkSteam(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlinksteam(
	const void *ptr,
	const NkAfterLinkSteamCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterLinkSteam(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelistmatches(
	const void *ptr,
	const NkBeforeListMatchesCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeListMatches(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistmatches(
	const void *ptr,
	const NkAfterListMatchesCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterListMatches(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelistnotifications(
	const void *ptr,
	const NkBeforeListNotificationsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeListNotifications(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlistnotifications(
	const void *ptr,
	const NkAfterListNotificationsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterListNotifications(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforedeletenotifications(
	const void *ptr,
	const NkBeforeDeleteNotificationsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeDeleteNotifications(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterdeletenotifications(
	const void *ptr,
	const NkAfterDeleteNotificationsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterDeleteNotifications(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeliststorageobjects(
	const void *ptr,
	const NkBeforeListStorageObjectsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeListStorageObjects(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterliststorageobjects(
	const void *ptr,
	const NkAfterListStorageObjectsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterListStorageObjects(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforereadstorageobjects(
	const void *ptr,
	const NkBeforeReadStorageObjectsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeReadStorageObjects(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterreadstorageobjects(
	const void *ptr,
	const NkAfterReadStorageObjectsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterReadStorageObjects(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforewritestorageobjects(
	const void *ptr,
	const NkBeforeWriteStorageObjectsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeWriteStorageObjects(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterwritestorageobjects(
	const void *ptr,
	const NkAfterWriteStorageObjectsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterWriteStorageObjects(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforedeletestorageobjects(
	const void *ptr,
	const NkBeforeDeleteStorageObjectsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeDeleteStorageObjects(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterdeletestorageobjects(
	const void *ptr,
	const NkAfterDeleteStorageObjectsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterDeleteStorageObjects(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforejointournament(
	const void *ptr,
	const NkBeforeJoinTournamentCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeJoinTournament(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterjointournament(
	const void *ptr,
	const NkAfterJoinTournamentCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterJoinTournament(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelisttournamentrecords(
	const void *ptr,
	const NkBeforeListTournamentRecordsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeListTournamentRecords(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlisttournamentrecords(
	const void *ptr,
	const NkAfterListTournamentRecordsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterListTournamentRecords(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelisttournaments(
	const void *ptr,
	const NkBeforeListTournamentsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeListTournaments(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlisttournaments(
	const void *ptr,
	const NkAfterListTournamentsCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterListTournaments(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforewritetournamentrecord(
	const void *ptr,
	const NkBeforeWriteTournamentRecordCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeWriteTournamentRecord(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterwritetournamentrecord(
	const void *ptr,
	const NkAfterWriteTournamentRecordCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterWriteTournamentRecord(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforelisttournamentrecordsaroundowner(
	const void *ptr,
	const NkBeforeListTournamentRecordsAroundOwnerCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeListTournamentRecordsAroundOwner(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterlisttournamentrecordsaroundowner(
	const void *ptr,
	const NkAfterListTournamentRecordsAroundOwnerCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterListTournamentRecordsAroundOwner(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinkapple(
	const void *ptr,
	const NkBeforeUnlinkAppleCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeUnlinkApple(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinkapple(
	const void *ptr,
	const NkAfterUnlinkAppleCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterUnlinkApple(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinkcustom(
	const void *ptr,
	const NkBeforeUnlinkCustomCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeUnlinkCustom(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinkcustom(
	const void *ptr,
	const NkAfterUnlinkCustomCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterUnlinkCustom(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinkdevice(
	const void *ptr,
	const NkBeforeUnlinkDeviceCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeUnlinkDevice(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinkdevice(
	const void *ptr,
	const NkAfterUnlinkDeviceCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterUnlinkDevice(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinkemail(
	const void *ptr,
	const NkBeforeUnlinkEmailCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeUnlinkEmail(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinkemail(
	const void *ptr,
	const NkAfterUnlinkEmailCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterUnlinkEmail(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinkfacebook(
	const void *ptr,
	const NkBeforeUnlinkFacebookCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeUnlinkFacebook(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinkfacebook(
	const void *ptr,
	const NkAfterUnlinkFacebookCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterUnlinkFacebook(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinkfacebookinstantgame(
	const void *ptr,
	const NkBeforeUnlinkFacebookInstantGameCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeUnlinkFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinkfacebookinstantgame(
	const void *ptr,
	const NkAfterUnlinkFacebookInstantGameCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterUnlinkFacebookInstantGame(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinkgamecenter(
	const void *ptr,
	const NkBeforeUnlinkGameCenterCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeUnlinkGameCenter(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinkgamecenter(
	const void *ptr,
	const NkAfterUnlinkGameCenterCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterUnlinkGameCenter(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinkgoogle(
	const void *ptr,
	const NkBeforeUnlinkGoogleCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeUnlinkGoogle(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinkgoogle(
	const void *ptr,
	const NkAfterUnlinkGoogleCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterUnlinkGoogle(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforeunlinksteam(
	const void *ptr,
	const NkBeforeUnlinkSteamCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeUnlinkSteam(
		ptr,
		cb,
		outerror);
}

int initializerregisterafterunlinksteam(
	const void *ptr,
	const NkAfterUnlinkSteamCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterUnlinkSteam(
		ptr,
		cb,
		outerror);
}

int initializerregisterbeforegetusers(
	const void *ptr,
	const NkBeforeGetUsersCallbackFn cb,
	NkString **outerror)
{
	return cInitializerBeforeGetUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisteraftergetusers(
	const void *ptr,
	const NkAfterGetUsersCallbackFn cb,
	NkString **outerror)
{
	return cInitializerAfterGetUsers(
		ptr,
		cb,
		outerror);
}

int initializerregisterevent(
	const void *ptr,
	const NkEventCallbackFn cb,
	NkString **outerror)
{
	return cInitializerEvent(
		ptr,
		cb,
		outerror);
}

int initializerregistereventsessionstart(
	const void *ptr,
	const NkEventCallbackFn cb,
	NkString **outerror)
{
	return cInitializerEventSessionStart(
		ptr,
		cb,
		outerror);
}

int initializerregistereventsessionend(
	const void *ptr,
	const NkEventCallbackFn cb,
	NkString **outerror)
{
	return cInitializerEventSessionEnd(
		ptr,
		cb,
		outerror);
}

*/
import "C"

func nkStringGo(s C.NkString) string {
	return C.GoStringN(s.p, C.int(s.n))
}
