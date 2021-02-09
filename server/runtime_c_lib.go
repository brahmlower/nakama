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
#include <stdlib.h>
#include "../include/nakama.h"

extern int initmodule(
	const void *ptr,
	NkContext,
	NkLogger,
	NkDb,
	NkModule,
	NkInitializer);

extern NkString contextvalue(
	const void * ptr,
	NkString key);

extern void loggerdebug(
	const void * ptr,
	NkString s);

extern void loggererror(
	const void * ptr,
	NkString s);

extern void loggerinfo(
	const void * ptr,
	NkString s);

extern void loggerwarn(
	const void * ptr,
	NkString s);

extern int moduleauthenticateapple(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
	bool **outcreated);

extern int moduleauthenticatecustom(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
	bool **outcreated);

extern int moduleauthenticatedevice(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
	bool **outcreated);

extern int moduleauthenticateemail(
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

extern int moduleauthenticatefacebook(
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

extern int moduleauthenticatefacebookinstantgame(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
	bool **outcreated);

extern int moduleauthenticategamecenter(
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

extern int moduleauthenticategoogle(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
	bool **outcreated);

extern int moduleauthenticatesteam(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	bool create,
	NkString **outuserid,
	NkString **outusername,
	NkString **outerror,
	bool **outcreated);

extern int moduleauthenticatetokengenerate(
	const void *ptr,
	NkString userid,
	NkString username,
	NkI64 expiry,
	NkMapString vars,
	NkString **outtoken,
	NkI64 **outexpiry,
	NkString **outerror);

extern int moduleaccountgetid(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkAccount **outaccount,
	NkString **outerror);

extern int moduleaccountsgetid(
	const void *ptr,
	const NkContext *ctx,
	NkString *userids,
	NkU32 numuserids,
	NkAccount **outaccounts,
	NkU32 **outnumaccounts,
	NkString **outerror);

extern int moduleaccountupdateid(
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

extern int moduleaccountdeleteid(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	bool recorded,
	NkString **outerror);

extern int moduleaccountexportid(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString **outaccount,
	NkString **outerror);

extern int moduleusersgetid(
	const void *ptr,
	const NkContext *ctx,
	NkString *keys,
	NkU32 numkeys,
	NkUser **outusers,
	NkU32 **outnumusers,
	NkString **outerror);

extern int moduleusersgetusername(
	const void *ptr,
	const NkContext *ctx,
	NkString *keys,
	NkU32 numkeys,
	NkUser **outusers,
	NkU32 **outnumusers,
	NkString **outerror);

extern int moduleusersbanid(
	const void *ptr,
	const NkContext *ctx,
	NkString *userids,
	NkU32 numids,
	NkString **outerror);

extern int moduleusersunbanid(
	const void *ptr,
	const NkContext *ctx,
	NkString *userids,
	NkU32 numids,
	NkString **outerror);

extern int modulelinkapple(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);

extern int modulelinkcustom(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);

extern int modulelinkdevice(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);

extern int modulelinkfacebookinstantgame(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);

extern int modulelinkgoogle(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);

extern int modulelinksteam(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);

extern int moduleunlinkapple(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);

extern int moduleunlinkcustom(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);

extern int moduleunlinkdevice(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);

extern int moduleunlinkemail(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);

extern int moduleunlinkfacebook(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);

extern int moduleunlinkfacebookinstantgame(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);

extern int moduleunlinkgoogle(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);

extern int moduleunlinksteam(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);


extern int modulelinkgamecenter(
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

extern int moduleunlinkgamecenter(
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

extern int modulelinkgoogle(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);

extern int modulelinksteam(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString linkid,
	NkString **outerror);

extern int modulelinkemail(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString email,
	NkString password,
	NkString **outerror);

extern int modulelinkfacebook(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString username,
	NkString token,
	bool importfriends,
	NkString **outerror);

extern int modulelinkgamecenter(
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

extern int modulestreamuserlist(
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

extern int modulestreamuserget(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkString userid,
	NkString sessionid,
	NkPresenceMeta **outmeta,
	NkString **outerror);

extern int modulestreamuserjoin(
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

extern int modulestreamuserupdate(
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

extern int modulestreamuserleave(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkString userid,
	NkString sessionid,
	NkString **outerror);

extern int modulestreamuserkick(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkPresence presence,
	NkString **outerror);

extern int modulestreamcount(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkU64 **outcount,
	NkString **outerror);

extern int modulestreamclose(
	const void *ptr,
	NkU8 mode,
	NkString subject,
	NkString subcontext,
	NkString label,
	NkString **outerror);

extern int modulestreamsend(
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

extern int modulestreamsendraw(
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

extern int modulesessiondisconnect(
	const void *ptr,
	const NkContext *ctx,
	NkString sessionid,
	NkString **outerror);

extern int modulematchcreate(
	const void *ptr,
	const NkContext *ctx,
	NkString module,
	NkMapAny params,
	NkString **outmatchid,
	NkString **outerror);

extern int modulematchget(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkMatch **outmatch,
	NkString **outerror);

extern int modulematchlist(
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

extern int modulenotificationsend(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkString subject,
	NkMapAny content,
	NkU64 code,
	NkString sender,
	bool persistent,
	NkString **outerror);

extern int modulenotificationssend(
	const void *ptr,
	const NkContext *ctx,
	const NkNotificationSend *notifications,
	NkU32 numnotifications,
	NkString **outerror);

extern int modulewalletupdate(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkMapI64 changeset,
	NkMapAny metadata,
	bool updateledger,
	NkMapI64 **outupdated,
	NkMapI64 **outprevious,
	NkString **outerror);

extern int modulewalletsupdate(
	const void *ptr,
	const NkContext *ctx,
	const NkWalletUpdate *updates,
	NkU32 numupdates,
	bool updateledger,
	NkWalletUpdateResult **outresults,
	NkU32 **outnumresults,
	NkString **outerror);

extern int modulewalletledgerupdate(
	const void *ptr,
	const NkContext *ctx,
	NkString itemid,
	NkMapAny metadata,
	NkWalletLedgerItem **outitem,
	NkString **outerror);

extern int modulewalletledgerlist(
	const void *ptr,
	const NkContext *ctx,
	NkString userid,
	NkU32 limit,
	NkString cursor,
	NkWalletLedgerItem **outitems,
	NkU32 **outnumitems,
	NkString **outcursor,
	NkString **outerror);

extern int modulestoragelist(
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

extern int modulestorageread(
	const void *ptr,
	const NkContext *ctx,
	const NkStorageRead *reads,
	NkU32 numreads,
	NkStorageObject **outobjs,
	NkU32 **outnumobjs,
	NkString **outerror);

extern int modulestoragewrite(
	const void *ptr,
	const NkContext *ctx,
	const NkStorageWrite *writes,
	NkU32 numwrites,
	NkStorageObjectAck **outacks,
	NkU32 **outnumacks,
	NkString **outerror);

extern int modulestoragedelete(
	const void *ptr,
	const NkContext *ctx,
	const NkStorageDelete *deletes,
	NkU32 numdeletes,
	NkString **outerror);

extern int modulemultiupdate(
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

extern int moduleleaderboardcreate(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	bool authoritative,
	NkString sortorder,
	NkString op,
	NkString resetschedule,
	NkMapAny metadata,
	NkString **outerror);

extern int moduleleaderboardrecordslist(
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

extern int moduleleaderboardrecordwrite(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkI64 score,
	NkI64 subscore,
	NkMapAny metadata,
	NkLeaderboardRecord **outrecord,
	NkString **outerror);

extern int moduleleaderboarddelete(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkString **outerror);

extern int moduleleaderboardrecorddelete(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkString **outerror);

extern int moduletournamentdelete(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkString **outerror);

extern int modulegroupdelete(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkString **outerror);

extern int moduletournamentcreate(
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

extern int moduletournamentaddattempt(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkU64 count,
	NkString **outerror);

extern int moduletournamentjoin(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkString username,
	NkString **outerror);

extern int moduletournamentsgetid(
	const void *ptr,
	const NkContext *ctx,
	const NkString *tournamentids,
	NkU32 numtournamentids,
	NkTournament **outtournaments,
	NkU32 **outnumtournaments,
	NkString **outerror);

extern int moduletournamentlist(
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

extern int moduletournamentrecordslist(
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

extern int moduletournamentrecordwrite(
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

extern int moduletournamentrecordshaystack(
	const void *ptr,
	const NkContext *ctx,
	NkString id,
	NkString ownerid,
	NkU32 limit,
	NkI64 expiry,
	NkLeaderboardRecord **outrecords,
	NkU32 **outnumrecords,
	NkString **outerror);

extern int modulegroupsgetid(
	const void *ptr,
	const NkContext *ctx,
	const NkString *groupids,
	NkU32 numgroupids,
	NkGroup **outgroups,
	NkU32 **outnumgroups,
	NkString **outerror);

extern int modulegroupcreate(
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

extern int modulegroupupdate(
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

extern int modulegroupuserjoin(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	NkString userid,
	NkString username,
	NkString **outerror);

extern int modulegroupuserleave(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	NkString userid,
	NkString username,
	NkString **outerror);

extern int modulegroupusersadd(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	const NkString *userids,
	NkU32 numuserids,
	NkString **outerror);

extern int modulegroupusersdemote(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	const NkString *userids,
	NkU32 numuserids,
	NkString **outerror);

extern int modulegroupuserskick(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	const NkString *userids,
	NkU32 numuserids,
	NkString **outerror);

extern int modulegroupuserspromote(
	const void *ptr,
	const NkContext *ctx,
	NkString groupid,
	const NkString *userids,
	NkU32 numuserids,
	NkString **outerror);

extern int modulegroupuserslist(
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

extern int moduleusergroupslist(
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

extern int modulefriendslist(
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

extern int moduleevent(
	const void *ptr,
	const NkContext *ctx,
	NkEvent evt,
	NkString **outerror);

extern int initializerregisterrpc(
	const void *ptr,
	NkString id,
	const NkRpcFn fn,
	NkString **outerror);

extern int initializerregisterbeforert(
	const void *ptr,
	NkString id,
	const NkBeforeRtCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterrt(
	const void *ptr,
	NkString id,
	const NkAfterRtCallbackFn cb,
	NkString **outerror);

extern int initializerregistermatchmakermatched(
	const void *ptr,
	const NkMatchmakerMatchedCallbackFn cb,
	NkString **outerror);

extern int initializerregistermatch(
	const void *ptr,
	NkString name,
	const NkMatchCallbackFn cb,
	NkString **outerror);

extern int initializerregistertournamentend (
	const void *ptr,
	const NkTournamentCallbackFn cb,
	NkString **outerror);

extern int initializerregistertournamentreset(
	const void *ptr,
	const NkTournamentCallbackFn cb,
	NkString **outerror);

extern int initializerregisterleaderboardreset(
	const void *ptr,
	const NkLeaderBoardCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforegetaccount(
	const void *ptr,
	const NkCallbackFn cb,
	NkString **outerror);

extern int initializerregisteraftergetaccount(
	const void *ptr,
	const NkAfterGetAccountCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeupdateaccount(
	const void *ptr,
	const NkBeforeUpdateAccountCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterupdateaccount(
	const void *ptr,
	const NkAfterUpdateAccountCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforesessionrefresh(
	const void *ptr,
	const NkBeforeSessionRefreshCallbackFn cb,
	NkString **outerror);

extern int initializerregisteraftersessionrefresh(
	const void *ptr,
	const NkAfterSessionRefreshCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeauthenticateapple(
	const void *ptr,
	const NkBeforeAuthenticateAppleCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterauthenticateapple(
	const void *ptr,
	const NkAfterAuthenticateAppleCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeauthenticatecustom(
	const void *ptr,
	const NkBeforeAuthenticateCustomCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterauthenticatecustom(
	const void *ptr,
	const NkAfterAuthenticateCustomCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeauthenticatedevice(
	const void *ptr,
	const NkBeforeAuthenticateDeviceCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterauthenticatedevice(
	const void *ptr,
	const NkAfterAuthenticateDeviceCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeauthenticateemail(
	const void *ptr,
	const NkBeforeAuthenticateEmailCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterauthenticateemail(
	const void *ptr,
	const NkAfterAuthenticateEmailCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeauthenticatefacebook(
	const void *ptr,
	const NkBeforeAuthenticateFacebookCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterauthenticatefacebook(
	const void *ptr,
	const NkAfterAuthenticateFacebookCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeauthenticatefacebookinstantgame(
	const void *ptr,
	const NkBeforeAuthenticateFacebookInstantGameCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterauthenticatefacebookinstantgame(
	const void *ptr,
	const NkAfterAuthenticateFacebookInstantGameCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeauthenticategamecenter(
	const void *ptr,
	const NkBeforeAuthenticateGameCenterCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterauthenticategamecenter(
	const void *ptr,
	const NkAfterAuthenticateGameCenterCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeauthenticategoogle(
	const void *ptr,
	const NkBeforeAuthenticateGoogleCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterauthenticategoogle(
	const void *ptr,
	const NkAfterAuthenticateGoogleCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeauthenticatesteam(
	const void *ptr,
	const NkBeforeAuthenticateSteamCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterauthenticatesteam(
	const void *ptr,
	const NkAfterAuthenticateSteamCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelistchannelmessages(
	const void *ptr,
	const NkBeforeListChannelMessagesCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlistchannelmessages(
	const void *ptr,
	const NkAfterListChannelMessagesCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelistfriends(
	const void *ptr,
	const NkBeforeListFriendsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlistfriends(
	const void *ptr,
	const NkAfterListFriendsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeaddfriends(
	const void *ptr,
	const NkBeforeAddFriendsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafteraddfriends(
	const void *ptr,
	const NkAfterAddFriendsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforedeletefriends(
	const void *ptr,
	const NkBeforeDeleteFriendsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterdeletefriends(
	const void *ptr,
	const NkAfterDeleteFriendsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeblockfriends(
	const void *ptr,
	const NkBeforeBlockFriendsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterblockfriends(
	const void *ptr,
	const NkAfterBlockFriendsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeimportfacebookfriends(
	const void *ptr,
	const NkBeforeImportFacebookFriendsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterimportfacebookfriends(
	const void *ptr,
	const NkAfterImportFacebookFriendsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforecreategroup(
	const void *ptr,
	const NkBeforeCreateGroupCallbackFn cb,
	NkString **outerror);

extern int initializerregisteraftercreategroup(
	const void *ptr,
	const NkAfterCreateGroupCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeupdategroup(
	const void *ptr,
	const NkBeforeUpdateGroupCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterupdategroup(
	const void *ptr,
	const NkAfterUpdateGroupCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforedeletegroup(
	const void *ptr,
	const NkBeforeDeleteGroupCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterdeletegroup(
	const void *ptr,
	const NkAfterDeleteGroupCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforejoingroup(
	const void *ptr,
	const NkBeforeJoinGroupCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterjoingroup(
	const void *ptr,
	const NkAfterJoinGroupCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeleavegroup(
	const void *ptr,
	const NkBeforeLeaveGroupCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterleavegroup(
	const void *ptr,
	const NkAfterLeaveGroupCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeaddgroupusers(
	const void *ptr,
	const NkBeforeAddGroupUsersCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafteraddgroupusers(
	const void *ptr,
	const NkAfterAddGroupUsersCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforebangroupusers(
	const void *ptr,
	const NkBeforeBanGroupUsersCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterbangroupusers(
	const void *ptr,
	const NkAfterBanGroupUsersCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforekickgroupusers(
	const void *ptr,
	const NkBeforeKickGroupUsersCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterkickgroupusers(
	const void *ptr,
	const NkAfterKickGroupUsersCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforepromotegroupusers(
	const void *ptr,
	const NkBeforePromoteGroupUsersCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterpromotegroupusers(
	const void *ptr,
	const NkAfterPromoteGroupUsersCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforedemotegroupusers(
	const void *ptr,
	const NkBeforeDemoteGroupUsersCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterdemotegroupusers(
	const void *ptr,
	const NkAfterDemoteGroupUsersCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelistgroupusers(
	const void *ptr,
	const NkBeforeListGroupUsersCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlistgroupusers(
	const void *ptr,
	const NkAfterListGroupUsersCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelistusergroups(
	const void *ptr,
	const NkBeforeListUserGroupsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlistusergroups(
	const void *ptr,
	const NkAfterListUserGroupsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelistgroups(
	const void *ptr,
	const NkBeforeListGroupsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlistgroups(
	const void *ptr,
	const NkAfterListGroupsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforedeleteleaderboardrecord(
	const void *ptr,
	const NkBeforeDeleteLeaderboardRecordCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterdeleteleaderboardrecord(
	const void *ptr,
	const NkAfterDeleteLeaderboardRecordCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelistleaderboardrecords(
	const void *ptr,
	const NkBeforeListLeaderboardRecordsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlistleaderboardrecords(
	const void *ptr,
	const NkAfterListLeaderboardRecordsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforewriteleaderboardrecord(
	const void *ptr,
	const NkBeforeWriteLeaderboardRecordCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterwriteleaderboardrecord(
	const void *ptr,
	const NkAfterWriteLeaderboardRecordCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelistleaderboardrecordsaroundowner(
	const void *ptr,
	const NkBeforeListLeaderboardRecordsAroundOwnerCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlistleaderboardrecordsaroundowner(
	const void *ptr,
	const NkAfterListLeaderboardRecordsAroundOwnerCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelinkapple(
	const void *ptr,
	const NkBeforeLinkAppleCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlinkapple(
	const void *ptr,
	const NkAfterLinkAppleCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelinkcustom(
	const void *ptr,
	const NkBeforeLinkCustomCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlinkcustom(
	const void *ptr,
	const NkAfterLinkCustomCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelinkdevice(
	const void *ptr,
	const NkBeforeLinkDeviceCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlinkdevice(
	const void *ptr,
	const NkAfterLinkDeviceCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelinkemail(
	const void *ptr,
	const NkBeforeLinkEmailCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlinkemail(
	const void *ptr,
	const NkAfterLinkEmailCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelinkfacebook(
	const void *ptr,
	const NkBeforeLinkFacebookCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlinkfacebook(
	const void *ptr,
	const NkAfterLinkFacebookCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelinkfacebookinstantgame(
	const void *ptr,
	const NkBeforeLinkFacebookInstantGameCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlinkfacebookinstantgame(
	const void *ptr,
	const NkAfterLinkFacebookInstantGameCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelinkgamecenter(
	const void *ptr,
	const NkBeforeLinkGameCenterCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlinkgamecenter(
	const void *ptr,
	const NkAfterLinkGameCenterCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelinkgoogle(
	const void *ptr,
	const NkBeforeLinkGoogleCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlinkgoogle(
	const void *ptr,
	const NkAfterLinkGoogleCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelinksteam(
	const void *ptr,
	const NkBeforeLinkSteamCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlinksteam(
	const void *ptr,
	const NkAfterLinkSteamCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelistmatches(
	const void *ptr,
	const NkBeforeListMatchesCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlistmatches(
	const void *ptr,
	const NkAfterListMatchesCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelistnotifications(
	const void *ptr,
	const NkBeforeListNotificationsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlistnotifications(
	const void *ptr,
	const NkAfterListNotificationsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforedeletenotifications(
	const void *ptr,
	const NkBeforeDeleteNotificationsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterdeletenotifications(
	const void *ptr,
	const NkAfterDeleteNotificationsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeliststorageobjects(
	const void *ptr,
	const NkBeforeListStorageObjectsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterliststorageobjects(
	const void *ptr,
	const NkAfterListStorageObjectsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforereadstorageobjects(
	const void *ptr,
	const NkBeforeReadStorageObjectsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterreadstorageobjects(
	const void *ptr,
	const NkAfterReadStorageObjectsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforewritestorageobjects(
	const void *ptr,
	const NkBeforeWriteStorageObjectsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterwritestorageobjects(
	const void *ptr,
	const NkAfterWriteStorageObjectsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforedeletestorageobjects(
	const void *ptr,
	const NkBeforeDeleteStorageObjectsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterdeletestorageobjects(
	const void *ptr,
	const NkAfterDeleteStorageObjectsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforejointournament(
	const void *ptr,
	const NkBeforeJoinTournamentCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterjointournament(
	const void *ptr,
	const NkAfterJoinTournamentCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelisttournamentrecords(
	const void *ptr,
	const NkBeforeListTournamentRecordsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlisttournamentrecords(
	const void *ptr,
	const NkAfterListTournamentRecordsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelisttournaments(
	const void *ptr,
	const NkBeforeListTournamentsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlisttournaments(
	const void *ptr,
	const NkAfterListTournamentsCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforewritetournamentrecord(
	const void *ptr,
	const NkBeforeWriteTournamentRecordCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterwritetournamentrecord(
	const void *ptr,
	const NkAfterWriteTournamentRecordCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforelisttournamentrecordsaroundowner(
	const void *ptr,
	const NkBeforeListTournamentRecordsAroundOwnerCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterlisttournamentrecordsaroundowner(
	const void *ptr,
	const NkAfterListTournamentRecordsAroundOwnerCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeunlinkapple(
	const void *ptr,
	const NkBeforeUnlinkAppleCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterunlinkapple(
	const void *ptr,
	const NkAfterUnlinkAppleCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeunlinkcustom(
	const void *ptr,
	const NkBeforeUnlinkCustomCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterunlinkcustom(
	const void *ptr,
	const NkAfterUnlinkCustomCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeunlinkdevice(
	const void *ptr,
	const NkBeforeUnlinkDeviceCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterunlinkdevice(
	const void *ptr,
	const NkAfterUnlinkDeviceCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeunlinkemail(
	const void *ptr,
	const NkBeforeUnlinkEmailCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterunlinkemail(
	const void *ptr,
	const NkAfterUnlinkEmailCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeunlinkfacebook(
	const void *ptr,
	const NkBeforeUnlinkFacebookCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterunlinkfacebook(
	const void *ptr,
	const NkAfterUnlinkFacebookCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeunlinkfacebookinstantgame(
	const void *ptr,
	const NkBeforeUnlinkFacebookInstantGameCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterunlinkfacebookinstantgame(
	const void *ptr,
	const NkAfterUnlinkFacebookInstantGameCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeunlinkgamecenter(
	const void *ptr,
	const NkBeforeUnlinkGameCenterCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterunlinkgamecenter(
	const void *ptr,
	const NkAfterUnlinkGameCenterCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeunlinkgoogle(
	const void *ptr,
	const NkBeforeUnlinkGoogleCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterunlinkgoogle(
	const void *ptr,
	const NkAfterUnlinkGoogleCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforeunlinksteam(
	const void *ptr,
	const NkBeforeUnlinkSteamCallbackFn cb,
	NkString **outerror);

extern int initializerregisterafterunlinksteam(
	const void *ptr,
	const NkAfterUnlinkSteamCallbackFn cb,
	NkString **outerror);

extern int initializerregisterbeforegetusers(
	const void *ptr,
	const NkBeforeGetUsersCallbackFn cb,
	NkString **outerror);

extern int initializerregisteraftergetusers(
	const void *ptr,
	const NkAfterGetUsersCallbackFn cb,
	NkString **outerror);

extern int initializerregisterevent(
	const void *ptr,
	const NkEventCallbackFn cb,
	NkString **outerror);

extern int initializerregistereventsessionstart(
	const void *ptr,
	const NkEventCallbackFn cb,
	NkString **outerror);

extern int initializerregistereventsessionend(
	const void *ptr,
	const NkEventCallbackFn cb,
	NkString **outerror);

*/
import "C"

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/mattn/go-pointer"
)

func cContext(ctx context.Context) C.NkContext {
	ret := C.NkContext{}
	ret.ptr = pointer.Save(ctx)
	ret.value = C.NkContextValueFn(C.contextvalue)

	return ret
}

func cDb(db *sql.DB) C.NkDb {
	ret := C.NkDb{}
	ret.ptr = pointer.Save(db)

	return ret
}

func cInitializer(initializer runtime.Initializer) C.NkInitializer {
	ret := C.NkInitializer{}
	ret.ptr = pointer.Save(initializer)
	ret.registerrpc = C.NkInitializerRpcFn(C.initializerregisterrpc)
	ret.registerbeforert = C.NkInitializerBeforeRtFn(C.initializerregisterbeforert)
	ret.registerafterrt = C.NkInitializerAfterRtFn(C.initializerregisterafterrt)
	ret.registermatchmakermatched = C.NkInitializerMatchmakerMatchedFn(C.initializerregistermatchmakermatched)
	ret.registermatch = C.NkInitializerMatchFn(C.initializerregistermatch)
	ret.registertournamentend = C.NkInitializerTournamentFn(C.initializerregistertournamentend)
	ret.registertournamentreset = C.NkInitializerTournamentFn(C.initializerregistertournamentreset)
	ret.registerleaderboardreset = C.NkInitializerLeaderBoardFn(C.initializerregisterleaderboardreset)
	ret.registerbeforegetaccount = C.NkInitializerBeforeGetAccountFn(C.initializerregisterbeforegetaccount)
	ret.registeraftergetaccount = C.NkInitializerAfterGetAccountFn(C.initializerregisteraftergetaccount)
	ret.registerbeforeupdateaccount = C.NkInitializerBeforeUpdateAccountFn(C.initializerregisterbeforeupdateaccount)
	ret.registerafterupdateaccount = C.NkInitializerAfterUpdateAccountFn(C.initializerregisterafterupdateaccount)
	ret.registerbeforesessionrefresh = C.NkInitializerBeforeSessionRefreshFn(C.initializerregisterbeforesessionrefresh)
	ret.registeraftersessionrefresh = C.NkInitializerAfterSessionRefreshFn(C.initializerregisteraftersessionrefresh)
	ret.registerbeforeauthenticateapple = C.NkInitializerBeforeAuthenticateAppleFn(C.initializerregisterbeforeauthenticateapple)
	ret.registerafterauthenticateapple = C.NkInitializerAfterAuthenticateAppleFn(C.initializerregisterafterauthenticateapple)
	ret.registerbeforeauthenticatecustom = C.NkInitializerBeforeAuthenticateCustomFn(C.initializerregisterbeforeauthenticatecustom)
	ret.registerafterauthenticatecustom = C.NkInitializerAfterAuthenticateCustomFn(C.initializerregisterafterauthenticatecustom)
	ret.registerbeforeauthenticatedevice = C.NkInitializerBeforeAuthenticateDeviceFn(C.initializerregisterbeforeauthenticatedevice)
	ret.registerafterauthenticatedevice = C.NkInitializerAfterAuthenticateDeviceFn(C.initializerregisterafterauthenticatedevice)
	ret.registerbeforeauthenticateemail = C.NkInitializerBeforeAuthenticateEmailFn(C.initializerregisterbeforeauthenticateemail)
	ret.registerafterauthenticateemail = C.NkInitializerAfterAuthenticateEmailFn(C.initializerregisterafterauthenticateemail)
	ret.registerbeforeauthenticatefacebook = C.NkInitializerBeforeAuthenticateFacebookFn(C.initializerregisterbeforeauthenticatefacebook)
	ret.registerafterauthenticatefacebook = C.NkInitializerAfterAuthenticateFacebookFn(C.initializerregisterafterauthenticatefacebook)
	ret.registerbeforeauthenticatefacebookinstantgame = C.NkInitializerBeforeAuthenticateFacebookInstantGameFn(C.initializerregisterbeforeauthenticatefacebookinstantgame)
	ret.registerafterauthenticatefacebookinstantgame = C.NkInitializerAfterAuthenticateFacebookInstantGameFn(C.initializerregisterafterauthenticatefacebookinstantgame)
	ret.registerbeforeauthenticategamecenter = C.NkInitializerBeforeAuthenticateGameCenterFn(C.initializerregisterbeforeauthenticategamecenter)
	ret.registerafterauthenticategamecenter = C.NkInitializerAfterAuthenticateGameCenterFn(C.initializerregisterafterauthenticategamecenter)
	ret.registerbeforeauthenticategoogle = C.NkInitializerBeforeAuthenticateGoogleFn(C.initializerregisterbeforeauthenticategoogle)
	ret.registerafterauthenticategoogle = C.NkInitializerAfterAuthenticateGoogleFn(C.initializerregisterafterauthenticategoogle)
	ret.registerbeforeauthenticatesteam = C.NkInitializerBeforeAuthenticateSteamFn(C.initializerregisterbeforeauthenticatesteam)
	ret.registerafterauthenticatesteam = C.NkInitializerAfterAuthenticateSteamFn(C.initializerregisterafterauthenticatesteam)
	ret.registerbeforelistchannelmessages = C.NkInitializerBeforeListChannelMessagesFn(C.initializerregisterbeforelistchannelmessages)
	ret.registerafterlistchannelmessages = C.NkInitializerAfterListChannelMessagesFn(C.initializerregisterafterlistchannelmessages)
	ret.registerbeforelistfriends = C.NkInitializerBeforeListFriendsFn(C.initializerregisterbeforelistfriends)
	ret.registerafterlistfriends = C.NkInitializerAfterListFriendsFn(C.initializerregisterafterlistfriends)
	ret.registerbeforeaddfriends = C.NkInitializerBeforeAddFriendsFn(C.initializerregisterbeforeaddfriends)
	ret.registerafteraddfriends = C.NkInitializerAfterAddFriendsFn(C.initializerregisterafteraddfriends)
	ret.registerbeforedeletefriends = C.NkInitializerBeforeDeleteFriendsFn(C.initializerregisterbeforedeletefriends)
	ret.registerafterdeletefriends = C.NkInitializerAfterDeleteFriendsFn(C.initializerregisterafterdeletefriends)
	ret.registerbeforeblockfriends = C.NkInitializerBeforeBlockFriendsFn(C.initializerregisterbeforeblockfriends)
	ret.registerafterblockfriends = C.NkInitializerAfterBlockFriendsFn(C.initializerregisterafterblockfriends)
	ret.registerbeforeimportfacebookfriends = C.NkInitializerBeforeImportFacebookFriendsFn(C.initializerregisterbeforeimportfacebookfriends)
	ret.registerafterimportfacebookfriends = C.NkInitializerAfterImportFacebookFriendsFn(C.initializerregisterafterimportfacebookfriends)
	ret.registerbeforecreategroup = C.NkInitializerBeforeCreateGroupFn(C.initializerregisterbeforecreategroup)
	ret.registeraftercreategroup = C.NkInitializerAfterCreateGroupFn(C.initializerregisteraftercreategroup)
	ret.registerbeforeupdategroup = C.NkInitializerBeforeUpdateGroupFn(C.initializerregisterbeforeupdategroup)
	ret.registerafterupdategroup = C.NkInitializerAfterUpdateGroupFn(C.initializerregisterafterupdategroup)
	ret.registerbeforedeletegroup = C.NkInitializerBeforeDeleteGroupFn(C.initializerregisterbeforedeletegroup)
	ret.registerafterdeletegroup = C.NkInitializerAfterDeleteGroupFn(C.initializerregisterafterdeletegroup)
	ret.registerbeforejoingroup = C.NkInitializerBeforeJoinGroupFn(C.initializerregisterbeforejoingroup)
	ret.registerafterjoingroup = C.NkInitializerAfterJoinGroupFn(C.initializerregisterafterjoingroup)
	ret.registerbeforeleavegroup = C.NkInitializerBeforeLeaveGroupFn(C.initializerregisterbeforeleavegroup)
	ret.registerafterleavegroup = C.NkInitializerAfterLeaveGroupFn(C.initializerregisterafterleavegroup)
	ret.registerbeforeaddgroupusers = C.NkInitializerBeforeAddGroupUsersFn(C.initializerregisterbeforeaddgroupusers)
	ret.registerafteraddgroupusers = C.NkInitializerAfterAddGroupUsersFn(C.initializerregisterafteraddgroupusers)
	ret.registerbeforebangroupusers = C.NkInitializerBeforeBanGroupUsersFn(C.initializerregisterbeforebangroupusers)
	ret.registerafterbangroupusers = C.NkInitializerAfterBanGroupUsersFn(C.initializerregisterafterbangroupusers)
	ret.registerbeforekickgroupusers = C.NkInitializerBeforeKickGroupUsersFn(C.initializerregisterbeforekickgroupusers)
	ret.registerafterkickgroupusers = C.NkInitializerAfterKickGroupUsersFn(C.initializerregisterafterkickgroupusers)
	ret.registerbeforepromotegroupusers = C.NkInitializerBeforePromoteGroupUsersFn(C.initializerregisterbeforepromotegroupusers)
	ret.registerafterpromotegroupusers = C.NkInitializerAfterPromoteGroupUsersFn(C.initializerregisterafterpromotegroupusers)
	ret.registerbeforedemotegroupusers = C.NkInitializerBeforeDemoteGroupUsersFn(C.initializerregisterbeforedemotegroupusers)
	ret.registerafterdemotegroupusers = C.NkInitializerAfterDemoteGroupUsersFn(C.initializerregisterafterdemotegroupusers)
	ret.registerbeforelistgroupusers = C.NkInitializerBeforeListGroupUsersFn(C.initializerregisterbeforelistgroupusers)
	ret.registerafterlistgroupusers = C.NkInitializerAfterListGroupUsersFn(C.initializerregisterafterlistgroupusers)
	ret.registerbeforelistusergroups = C.NkInitializerBeforeListUserGroupsFn(C.initializerregisterbeforelistusergroups)
	ret.registerafterlistusergroups = C.NkInitializerAfterListUserGroupsFn(C.initializerregisterafterlistusergroups)
	ret.registerbeforelistgroups = C.NkInitializerBeforeListGroupsFn(C.initializerregisterbeforelistgroups)
	ret.registerafterlistgroups = C.NkInitializerAfterListGroupsFn(C.initializerregisterafterlistgroups)
	ret.registerbeforedeleteleaderboardrecord = C.NkInitializerBeforeDeleteLeaderboardRecordFn(C.initializerregisterbeforedeleteleaderboardrecord)
	ret.registerafterdeleteleaderboardrecord = C.NkInitializerAfterDeleteLeaderboardRecordFn(C.initializerregisterafterdeleteleaderboardrecord)
	ret.registerbeforelistleaderboardrecords = C.NkInitializerBeforeListLeaderboardRecordsFn(C.initializerregisterbeforelistleaderboardrecords)
	ret.registerafterlistleaderboardrecords = C.NkInitializerAfterListLeaderboardRecordsFn(C.initializerregisterafterlistleaderboardrecords)
	ret.registerbeforewriteleaderboardrecord = C.NkInitializerBeforeWriteLeaderboardRecordFn(C.initializerregisterbeforewriteleaderboardrecord)
	ret.registerafterwriteleaderboardrecord = C.NkInitializerAfterWriteLeaderboardRecordFn(C.initializerregisterafterwriteleaderboardrecord)
	ret.registerbeforelistleaderboardrecordsaroundowner = C.NkInitializerBeforeListLeaderboardRecordsAroundOwnerFn(C.initializerregisterbeforelistleaderboardrecordsaroundowner)
	ret.registerafterlistleaderboardrecordsaroundowner = C.NkInitializerAfterListLeaderboardRecordsAroundOwnerFn(C.initializerregisterafterlistleaderboardrecordsaroundowner)
	ret.registerbeforelinkapple = C.NkInitializerBeforeLinkAppleFn(C.initializerregisterbeforelinkapple)
	ret.registerafterlinkapple = C.NkInitializerAfterLinkAppleFn(C.initializerregisterafterlinkapple)
	ret.registerbeforelinkcustom = C.NkInitializerBeforeLinkCustomFn(C.initializerregisterbeforelinkcustom)
	ret.registerafterlinkcustom = C.NkInitializerAfterLinkCustomFn(C.initializerregisterafterlinkcustom)
	ret.registerbeforelinkdevice = C.NkInitializerBeforeLinkDeviceFn(C.initializerregisterbeforelinkdevice)
	ret.registerafterlinkdevice = C.NkInitializerAfterLinkDeviceFn(C.initializerregisterafterlinkdevice)
	ret.registerbeforelinkemail = C.NkInitializerBeforeLinkEmailFn(C.initializerregisterbeforelinkemail)
	ret.registerafterlinkemail = C.NkInitializerAfterLinkEmailFn(C.initializerregisterafterlinkemail)
	ret.registerbeforelinkfacebook = C.NkInitializerBeforeLinkFacebookFn(C.initializerregisterbeforelinkfacebook)
	ret.registerafterlinkfacebook = C.NkInitializerAfterLinkFacebookFn(C.initializerregisterafterlinkfacebook)
	ret.registerbeforelinkfacebookinstantgame = C.NkInitializerBeforeLinkFacebookInstantGameFn(C.initializerregisterbeforelinkfacebookinstantgame)
	ret.registerafterlinkfacebookinstantgame = C.NkInitializerAfterLinkFacebookInstantGameFn(C.initializerregisterafterlinkfacebookinstantgame)
	ret.registerbeforelinkgamecenter = C.NkInitializerBeforeLinkGameCenterFn(C.initializerregisterbeforelinkgamecenter)
	ret.registerafterlinkgamecenter = C.NkInitializerAfterLinkGameCenterFn(C.initializerregisterafterlinkgamecenter)
	ret.registerbeforelinkgoogle = C.NkInitializerBeforeLinkGoogleFn(C.initializerregisterbeforelinkgoogle)
	ret.registerafterlinkgoogle = C.NkInitializerAfterLinkGoogleFn(C.initializerregisterafterlinkgoogle)
	ret.registerbeforelinksteam = C.NkInitializerBeforeLinkSteamFn(C.initializerregisterbeforelinksteam)
	ret.registerafterlinksteam = C.NkInitializerAfterLinkSteamFn(C.initializerregisterafterlinksteam)
	ret.registerbeforelistmatches = C.NkInitializerBeforeListMatchesFn(C.initializerregisterbeforelistmatches)
	ret.registerafterlistmatches = C.NkInitializerAfterListMatchesFn(C.initializerregisterafterlistmatches)
	ret.registerbeforelistnotifications = C.NkInitializerBeforeListNotificationsFn(C.initializerregisterbeforelistnotifications)
	ret.registerafterlistnotifications = C.NkInitializerAfterListNotificationsFn(C.initializerregisterafterlistnotifications)
	ret.registerbeforedeletenotifications = C.NkInitializerBeforeDeleteNotificationsFn(C.initializerregisterbeforedeletenotifications)
	ret.registerafterdeletenotifications = C.NkInitializerAfterDeleteNotificationsFn(C.initializerregisterafterdeletenotifications)
	ret.registerbeforeliststorageobjects = C.NkInitializerBeforeListStorageObjectsFn(C.initializerregisterbeforeliststorageobjects)
	ret.registerafterliststorageobjects = C.NkInitializerAfterListStorageObjectsFn(C.initializerregisterafterliststorageobjects)
	ret.registerbeforereadstorageobjects = C.NkInitializerBeforeReadStorageObjectsFn(C.initializerregisterbeforereadstorageobjects)
	ret.registerafterreadstorageobjects = C.NkInitializerAfterReadStorageObjectsFn(C.initializerregisterafterreadstorageobjects)
	ret.registerbeforewritestorageobjects = C.NkInitializerBeforeWriteStorageObjectsFn(C.initializerregisterbeforewritestorageobjects)
	ret.registerafterwritestorageobjects = C.NkInitializerAfterWriteStorageObjectsFn(C.initializerregisterafterwritestorageobjects)
	ret.registerbeforedeletestorageobjects = C.NkInitializerBeforeDeleteStorageObjectsFn(C.initializerregisterbeforedeletestorageobjects)
	ret.registerafterdeletestorageobjects = C.NkInitializerAfterDeleteStorageObjectsFn(C.initializerregisterafterdeletestorageobjects)
	ret.registerbeforejointournament = C.NkInitializerBeforeJoinTournamentFn(C.initializerregisterbeforejointournament)
	ret.registerafterjointournament = C.NkInitializerAfterJoinTournamentFn(C.initializerregisterafterjointournament)
	ret.registerbeforelisttournamentrecords = C.NkInitializerBeforeListTournamentRecordsFn(C.initializerregisterbeforelisttournamentrecords)
	ret.registerafterlisttournamentrecords = C.NkInitializerAfterListTournamentRecordsFn(C.initializerregisterafterlisttournamentrecords)
	ret.registerbeforelisttournaments = C.NkInitializerBeforeListTournamentsFn(C.initializerregisterbeforelisttournaments)
	ret.registerafterlisttournaments = C.NkInitializerAfterListTournamentsFn(C.initializerregisterafterlisttournaments)
	ret.registerbeforewritetournamentrecord = C.NkInitializerBeforeWriteTournamentRecordFn(C.initializerregisterbeforewritetournamentrecord)
	ret.registerafterwritetournamentrecord = C.NkInitializerAfterWriteTournamentRecordFn(C.initializerregisterafterwritetournamentrecord)
	ret.registerbeforelisttournamentrecordsaroundowner = C.NkInitializerBeforeListTournamentRecordsAroundOwnerFn(C.initializerregisterbeforelisttournamentrecordsaroundowner)
	ret.registerafterlisttournamentrecordsaroundowner = C.NkInitializerAfterListTournamentRecordsAroundOwnerFn(C.initializerregisterafterlisttournamentrecordsaroundowner)
	ret.registerbeforeunlinkapple = C.NkInitializerBeforeUnlinkAppleFn(C.initializerregisterbeforeunlinkapple)
	ret.registerafterunlinkapple = C.NkInitializerAfterUnlinkAppleFn(C.initializerregisterafterunlinkapple)
	ret.registerbeforeunlinkcustom = C.NkInitializerBeforeUnlinkCustomFn(C.initializerregisterbeforeunlinkcustom)
	ret.registerafterunlinkcustom = C.NkInitializerAfterUnlinkCustomFn(C.initializerregisterafterunlinkcustom)
	ret.registerbeforeunlinkdevice = C.NkInitializerBeforeUnlinkDeviceFn(C.initializerregisterbeforeunlinkdevice)
	ret.registerafterunlinkdevice = C.NkInitializerAfterUnlinkDeviceFn(C.initializerregisterafterunlinkdevice)
	ret.registerbeforeunlinkemail = C.NkInitializerBeforeUnlinkEmailFn(C.initializerregisterbeforeunlinkemail)
	ret.registerafterunlinkemail = C.NkInitializerAfterUnlinkEmailFn(C.initializerregisterafterunlinkemail)
	ret.registerbeforeunlinkfacebook = C.NkInitializerBeforeUnlinkFacebookFn(C.initializerregisterbeforeunlinkfacebook)
	ret.registerafterunlinkfacebook = C.NkInitializerAfterUnlinkFacebookFn(C.initializerregisterafterunlinkfacebook)
	ret.registerbeforeunlinkfacebookinstantgame = C.NkInitializerBeforeUnlinkFacebookInstantGameFn(C.initializerregisterbeforeunlinkfacebookinstantgame)
	ret.registerafterunlinkfacebookinstantgame = C.NkInitializerAfterUnlinkFacebookInstantGameFn(C.initializerregisterafterunlinkfacebookinstantgame)
	ret.registerbeforeunlinkgamecenter = C.NkInitializerBeforeUnlinkGameCenterFn(C.initializerregisterbeforeunlinkgamecenter)
	ret.registerafterunlinkgamecenter = C.NkInitializerAfterUnlinkGameCenterFn(C.initializerregisterafterunlinkgamecenter)
	ret.registerbeforeunlinkgoogle = C.NkInitializerBeforeUnlinkGoogleFn(C.initializerregisterbeforeunlinkgoogle)
	ret.registerafterunlinkgoogle = C.NkInitializerAfterUnlinkGoogleFn(C.initializerregisterafterunlinkgoogle)
	ret.registerbeforeunlinksteam = C.NkInitializerBeforeUnlinkSteamFn(C.initializerregisterbeforeunlinksteam)
	ret.registerafterunlinksteam = C.NkInitializerAfterUnlinkSteamFn(C.initializerregisterafterunlinksteam)
	ret.registerbeforegetusers = C.NkInitializerBeforeGetUsersFn(C.initializerregisterbeforegetusers)
	ret.registeraftergetusers = C.NkInitializerAfterGetUsersFn(C.initializerregisteraftergetusers)
	ret.registerevent = C.NkInitializerEventFn(C.initializerregisterevent)
	ret.registereventsessionstart = C.NkInitializerEventSessionStartFn(C.initializerregistereventsessionstart)
	ret.registereventsessionend = C.NkInitializerEventSessionEndFn(C.initializerregistereventsessionend)

	return ret
}

func cLogger(logger runtime.Logger) C.NkLogger {
	ret := C.NkLogger{}
	ret.ptr = pointer.Save(logger)
	ret.debug = C.NkLoggerLevelFn(C.loggerdebug)
	ret.error = C.NkLoggerLevelFn(C.loggererror)
	ret.info = C.NkLoggerLevelFn(C.loggerinfo)
	ret.warn = C.NkLoggerLevelFn(C.loggerwarn)

	return ret
}

func cNakamaModule(nk runtime.NakamaModule) C.NkModule {
	ret := C.NkModule{}
	ret.ptr = pointer.Save(nk)
	ret.authenticateapple = C.NkModuleAuthenticateFn(C.moduleauthenticateapple)
	ret.authenticatecustom = C.NkModuleAuthenticateFn(C.moduleauthenticatecustom)
	ret.authenticatedevice = C.NkModuleAuthenticateFn(C.moduleauthenticatedevice)
	ret.authenticateemail = C.NkModuleAuthenticateEmailFn(C.moduleauthenticateemail)
	ret.authenticatefacebook = C.NkModuleAuthenticateFacebookFn(C.moduleauthenticatefacebook)
	ret.authenticatefacebookinstantgame = C.NkModuleAuthenticateFn(C.moduleauthenticatefacebookinstantgame)
	ret.authenticategamecenter = C.NkModuleAuthenticateGameCenterFn(C.moduleauthenticategamecenter)
	ret.authenticategoogle = C.NkModuleAuthenticateFn(C.moduleauthenticategoogle)
	ret.authenticatesteam = C.NkModuleAuthenticateFn(C.moduleauthenticatesteam)
	ret.authenticatetokengenerate = C.NkModuleAuthenticateTokenGenerateFn(C.moduleauthenticatetokengenerate)
	ret.accountgetid = C.NkModuleAccountGetIdFn(C.moduleaccountgetid)
	ret.accountsgetid = C.NkModuleAccountsGetIdFn(C.moduleaccountsgetid)
	ret.accountupdateid = C.NkModuleAccountUpdateIdFn(C.moduleaccountupdateid)
	ret.accountdeleteid = C.NkModuleAccountDeleteIdFn(C.moduleaccountdeleteid)
	ret.accountexportid = C.NkModuleAccountExportIdFn(C.moduleaccountexportid)
	ret.usersgetid = C.NkModuleUsersGetFn(C.moduleusersgetid)
	ret.usersgetusername = C.NkModuleUsersGetFn(C.moduleusersgetusername)
	ret.usersbanid = C.NkModuleUsersBanFn(C.moduleusersbanid)
	ret.usersunbanid = C.NkModuleUsersBanFn(C.moduleusersunbanid)
	ret.linkapple = C.NkModuleLinkFn(C.modulelinkapple)
	ret.linkcustom = C.NkModuleLinkFn(C.modulelinkcustom)
	ret.linkdevice = C.NkModuleLinkFn(C.modulelinkdevice)
	ret.linkemail = C.NkModuleLinkEmailFn(C.modulelinkemail)
	ret.linkfacebook = C.NkModuleLinkFacebookFn(C.modulelinkfacebook)
	ret.linkfacebookinstantgame = C.NkModuleLinkFn(C.modulelinkfacebookinstantgame)
	ret.linkgamecenter = C.NkModuleLinkGameCenterFn(C.modulelinkgamecenter)
	ret.linkgoogle = C.NkModuleLinkFn(C.modulelinkgoogle)
	ret.linksteam = C.NkModuleLinkFn(C.modulelinksteam)
	ret.unlinkapple = C.NkModuleLinkFn(C.moduleunlinkapple)
	ret.unlinkcustom = C.NkModuleLinkFn(C.moduleunlinkcustom)
	ret.unlinkdevice = C.NkModuleLinkFn(C.moduleunlinkdevice)
	ret.unlinkemail = C.NkModuleLinkFn(C.moduleunlinkemail)
	ret.unlinkfacebook = C.NkModuleLinkFn(C.moduleunlinkfacebook)
	ret.unlinkfacebookinstantgame = C.NkModuleLinkFn(C.moduleunlinkfacebookinstantgame)
	ret.unlinkgamecenter = C.NkModuleLinkGameCenterFn(C.moduleunlinkgamecenter)
	ret.unlinkgoogle = C.NkModuleLinkFn(C.moduleunlinkgoogle)
	ret.unlinksteam = C.NkModuleLinkFn(C.moduleunlinksteam)
	ret.streamuserlist = C.NkModuleStreamUserListFn(C.modulestreamuserlist)
	ret.streamuserget = C.NkModuleStreamUserGetFn(C.modulestreamuserget)
	ret.streamuserjoin = C.NkModuleStreamUserJoinFn(C.modulestreamuserjoin)
	ret.streamuserupdate = C.NkModuleStreamUserUpdateFn(C.modulestreamuserupdate)
	ret.streamuserleave = C.NkModuleStreamUserLeaveFn(C.modulestreamuserleave)
	ret.streamuserkick = C.NkModuleStreamUserKickFn(C.modulestreamuserkick)
	ret.streamcount = C.NkModuleStreamCountFn(C.modulestreamcount)
	ret.streamclose = C.NkModuleStreamCloseFn(C.modulestreamclose)
	ret.streamsend = C.NkModuleStreamSendFn(C.modulestreamsend)
	ret.streamsendraw = C.NkModuleStreamSendRawFn(C.modulestreamsendraw)
	ret.sessiondisconnect = C.NkModuleSessionDisconnectFn(C.modulesessiondisconnect)
	ret.matchcreate = C.NkModuleMatchCreateFn(C.modulematchcreate)
	ret.matchget = C.NkModuleMatchGetFn(C.modulematchget)
	ret.matchlist = C.NkModuleMatchListFn(C.modulematchlist)
	ret.notificationsend = C.NkModuleNotificationSendFn(C.modulenotificationsend)
	ret.notificationssend = C.NkModuleNotificationsSendFn(C.modulenotificationssend)
	ret.walletupdate = C.NkModuleWalletUpdateFn(C.modulewalletupdate)
	ret.walletsupdate = C.NkModuleWalletsUpdateFn(C.modulewalletsupdate)
	ret.walletledgerupdate = C.NkModuleWalletLedgerUpdateFn(C.modulewalletledgerupdate)
	ret.walletledgerlist = C.NkModuleWalletLedgerListFn(C.modulewalletledgerlist)
	ret.storagelist = C.NkModuleStorageListFn(C.modulestoragelist)
	ret.storageread = C.NkModuleStorageReadFn(C.modulestorageread)
	ret.storagewrite = C.NkModuleStorageWriteFn(C.modulestoragewrite)
	ret.storagedelete = C.NkModuleStorageDeleteFn(C.modulestoragedelete)
	ret.multiupdate = C.NkModuleMultiUpdateFn(C.modulemultiupdate)
	ret.leaderboardcreate = C.NkModuleLeaderboardCreateFn(C.moduleleaderboardcreate)
	ret.leaderboarddelete = C.NkModuleDeleteFn(C.moduleleaderboarddelete)
	ret.leaderboardrecordslist = C.NkModuleLeaderboardRecordsListFn(C.moduleleaderboardrecordslist)
	ret.leaderboardrecordwrite = C.NkModuleLeaderboardRecordWriteFn(C.moduleleaderboardrecordwrite)
	ret.leaderboardrecorddelete = C.NkModuleDeleteFn(C.moduleleaderboardrecorddelete)
	ret.tournamentcreate = C.NkModuleTournamentCreateFn(C.moduletournamentcreate)
	ret.tournamentdelete = C.NkModuleDeleteFn(C.moduletournamentdelete)
	ret.tournamentaddattempt = C.NkModuleTournamentAddAttemptFn(C.moduletournamentaddattempt)
	ret.tournamentjoin = C.NkModuleTournamentJoinFn(C.moduletournamentjoin)
	ret.tournamentsgetid = C.NkModuleTournamentsGetIdFn(C.moduletournamentsgetid)
	ret.tournamentlist = C.NkModuleTournamentListFn(C.moduletournamentlist)
	ret.tournamentrecordslist = C.NkModuleTournamentRecordsListFn(C.moduletournamentrecordslist)
	ret.tournamentrecordwrite = C.NkModuleTournamentRecordWriteFn(C.moduletournamentrecordwrite)
	ret.tournamentrecordshaystack = C.NkModuleTournamentRecordsHaystackFn(C.moduletournamentrecordshaystack)
	ret.groupsgetid = C.NkModuleGroupsGetIdFn(C.modulegroupsgetid)
	ret.groupcreate = C.NkModuleGroupCreateFn(C.modulegroupcreate)
	ret.groupupdate = C.NkModuleGroupUpdateFn(C.modulegroupupdate)
	ret.groupdelete = C.NkModuleDeleteFn(C.modulegroupdelete)
	ret.groupuserjoin = C.NkModuleGroupUserFn(C.modulegroupuserjoin)
	ret.groupuserleave = C.NkModuleGroupUserFn(C.modulegroupuserleave)
	ret.groupusersadd = C.NkModuleGroupUsersFn(C.modulegroupusersadd)
	ret.groupuserskick = C.NkModuleGroupUsersFn(C.modulegroupuserskick)
	ret.groupuserspromote = C.NkModuleGroupUsersFn(C.modulegroupuserspromote)
	ret.groupusersdemote = C.NkModuleGroupUsersFn(C.modulegroupusersdemote)
	ret.groupuserslist = C.NkModuleGroupUsersListFn(C.modulegroupuserslist)
	ret.usergroupslist = C.NkModuleUserGroupsListFn(C.moduleusergroupslist)
	ret.friendslist = C.NkModuleFriendsListFn(C.modulefriendslist)
	ret.event = C.NkModuleEventFn(C.moduleevent)

	return ret
}

// CLib holds a C-module runtime library
type CLib struct {
	err  error // set if library failed to load or missing required symbol(s)
	name string
	syms CSymbols
}

func (c *CLib) initModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	if c.err != nil {
		return c.err
	}

	if c.syms.initModule == nil {
		return errors.New("Missing c-module library initialisation function")
	}

	cDb := cDb(db)
	cCtx := cContext(ctx)
	cNk := cNakamaModule(nk)
	cLogger := cLogger(logger)
	cInitializer := cInitializer(initializer)

	err := C.initmodule(c.syms.initModule, cCtx, cLogger, cDb, cNk, cInitializer)

	pointer.Unref(cDb.ptr)
	pointer.Unref(cCtx.ptr)
	pointer.Unref(cLogger.ptr)
	pointer.Unref(cNk.ptr)
	pointer.Unref(cInitializer.ptr)

	if err != 0 {
		return fmt.Errorf("Could not initialize c-module: %d", err)
	}

	return nil
}
