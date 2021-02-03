// Copyright 2021 The Nakama Authors
//
// Licensed under the Apache License, Version 2.0 (the License);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an AS IS BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// NOTE: In order to implement a c-module, you must provide the following function:
//
// int nk_init_module(NkContext, NkLogger, NkDb, NkModule, NkInitializer);

#ifndef NAKAMA_H
#define NAKAMA_H

#include <stdbool.h>
#include <stddef.h>

#ifdef __cplusplus
extern "C"
{
#endif

	typedef unsigned char NkU8;
	typedef unsigned short NkU16;
	typedef unsigned int NkU32;
	typedef unsigned long long NkU64;

	typedef signed char NkI8;
	typedef short NkI16;
	typedef int NkI32;
	typedef long long NkI64;

	typedef float NkF32;
	typedef double NkF64;

	typedef struct
	{
		const char *p;
		ptrdiff_t n;
	} NkString;

	const NkU8 NK_U8_T = 0;
	const NkU8 NK_U16_T = 1;
	const NkU8 NK_U32_T = 2;
	const NkU8 NK_U64_T = 3;
	const NkU8 NK_I8_T = 4;
	const NkU8 NK_I16_T = 5;
	const NkU8 NK_I32_T = 6;
	const NkU8 NK_I64_T = 7;
	const NkU8 NK_F32_T = 8;
	const NkU8 NK_F64_T = 9;
	const NkU8 NK_STRING_T = 10;
	const NkU8 NK_LAST_T = 11;

	typedef struct NkAny
	{
		NkU8 type;
		union NkValue
		{
			NkU8 u8;
			NkU16 u16;
			NkU32 u32;
			NkU64 u64;
			NkI8 i8;
			NkI16 i16;
			NkI32 i32;
			NkI64 i64;
			NkF32 f32;
			NkF64 f64;
			NkString s;
		} value;
	} NkAny;

	typedef struct NkMapAny
	{
		NkU32 count;
		NkString *keys;
		NkAny *values;
	} NkMapAny;

	typedef struct NkMapI64
	{
		NkU32 count;
		NkString *keys;
		NkI64 *values;
	} NkMapI64;

	typedef struct NkMapString
	{
		NkU32 count;
		NkString *keys;
		NkString *values;
	} NkMapString;

	typedef void (*NkContextValueFn)(
		const void *ptr,
		NkString key,
		NkString **outvalue);

	typedef struct NkContext
	{
		const void *ptr;
		NkContextValueFn value;
	} NkContext;

	typedef NkMapAny (*NkLoggerFieldsFn)(
		const void *ptr);

	typedef void (*NkLoggerLevelFn)(
		const void *ptr,
		NkString s);

	typedef NkLogger (*NkLoggerWithFieldFn)(
		const void *ptr,
		NkString key,
		NkString value);

	typedef NkLogger (*NkLoggerWithFieldsFn)(
		const void *ptr,
		NkMapAny fields);

	typedef struct NkLogger
	{
		const void *ptr;
		NkLoggerLevelFn debug;
		NkLoggerLevelFn error;
		NkLoggerFieldsFn fields;
		NkLoggerLevelFn info;
		NkLoggerLevelFn warn;
		NkLoggerWithFieldFn withfield;
		NkLoggerWithFieldsFn withfields;
	} NkLogger;

	typedef struct NkDb
	{
		const void *ptr;
	} NkDb;

	typedef int (*NkModuleAuthenticateEmailFn)(
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

	typedef int (*NkModuleAuthenticateFacebookFn)(
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

	typedef int (*NkModuleAuthenticateFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkString username,
		bool create,
		NkString **outuserid,
		NkString **outusername,
		NkString **outerror,
		bool **outcreated);

	typedef int (*NkModuleAuthenticateGameCenterFn)(
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

	typedef int (*NkModuleAuthenticateGenerateTokenFn)(
		NkString userid,
		NkString username,
		NkI64 expiry,
		NkMapString vars,
		NkString **outtoken,
		NkI64 **outexpiry,
		NkString **outerror);

	typedef struct NkAccount
	{
		const void *ptr;
	} NkAccount;

	typedef int (*NkModuleAccountGetIdFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkAccount **outaccount,
		NkString **outerror);

	typedef int (*NkModuleAccountsGetIdFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString *userids,
		NkU32 numuserids,
		NkAccount **outaccounts,
		NkU32 **outnumaccounts,
		NkString **outerror);

	typedef int (*NkModuleAccountUpdateIdFn)(
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

	typedef int (*NkModuleAccountDeleteIdFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		bool recorded,
		NkString **outerror);

	typedef int (*NkModuleAccountExportIdFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkString **outaccount,
		NkString **outerror);

	typedef struct NkUser
	{
		const void *ptr;
	} NkUser;

	typedef int (*NkModuleUsersGetFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString *keys,
		NkU32 numkeys,
		NkUser **outusers,
		NkU32 **outnumusers,
		NkString **outerror);

	typedef int (*NkModuleUsersBanFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString *userids,
		NkU32 numids,
		NkString **outerror);

	typedef int (*NkModuleLinkFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkString linkid,
		NkString **outerror);

	typedef int (*NkModuleLinkEmailFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkString email,
		NkString password,
		NkString **outerror);

	typedef int (*NkModuleLinkFacebookFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkString username,
		NkString token,
		bool importfriends,
		NkString **outerror);

	typedef int (*NkModuleLinkGameCenterFn)(
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

	typedef struct NkPresence
	{
		const void *ptr;
	} NkPresence;

	typedef int (*NkModuleStreamUserListFn)(
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

	typedef struct NkPresenceMeta
	{
		const void *ptr;
	} NkPresenceMeta;

	typedef int (*NkModuleStreamUserGetFn)(
		const void *ptr,
		NkU8 mode,
		NkString subject,
		NkString subcontext,
		NkString label,
		NkString userid,
		NkString sessionid,
		NkPresenceMeta **outmeta,
		NkString **outerror);

	typedef int (*NkModuleStreamUserJoinFn)(
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

	typedef int (*NkModuleStreamUserUpdateFn)(
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

	typedef int (*NkModuleStreamUserLeaveFn)(
		const void *ptr,
		NkU8 mode,
		NkString subject,
		NkString subcontext,
		NkString label,
		NkString userid,
		NkString sessionid,
		NkString **outerror);

	typedef int (*NkModuleStreamUserKickFn)(
		const void *ptr,
		NkU8 mode,
		NkString subject,
		NkString subcontext,
		NkString label,
		NkPresence presence,
		NkString **outerror);

	typedef int (*NkModuleStreamCountFn)(
		const void *ptr,
		NkU8 mode,
		NkString subject,
		NkString subcontext,
		NkString label,
		NkU64 **outcount,
		NkString **outerror);

	typedef int (*NkModuleStreamCloseFn)(
		const void *ptr,
		NkU8 mode,
		NkString subject,
		NkString subcontext,
		NkString label,
		NkString **outerror);

	typedef int (*NkModuleStreamSendFn)(
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

	typedef struct NkEnvelope
	{
		const void *ptr;
	} NkEnvelope;

	typedef int (*NkModuleStreamSendRawFn)(
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

	typedef int (*NkModuleSessionDisconnectFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString sessionid,
		NkString **outerror);

	typedef int (*NkModuleMatchCreateFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString module,
		NkMapAny params,
		NkString **outmatchid,
		NkString **outerror);

	typedef struct NkMatch
	{
		const void *ptr;
	} NkMatch;

	typedef int (*NkModuleMatchGetFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString id,
		NkMatch **outmatch,
		NkString **outerror);

	typedef int (*NkModuleMatchListFn)(
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

	typedef int (*NkModuleNotificationSendFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkString subject,
		NkMapAny content,
		NkU64 code,
		NkString sender,
		bool persistent,
		NkString **outerror);

	typedef struct NkNotificationSend
	{
		const void *ptr;
	} NkNotificationSend;

	typedef int (*NkModuleNotificationsSendFn)(
		const void *ptr,
		const NkContext *ctx,
		const NkNotificationSend *notifications,
		NkU32 numnotifications,
		NkString **outerror);

	typedef int (*NkModuleWalletUpdateFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkMapI64 changeset,
		NkMapAny metadata,
		bool updateledger,
		NkMapI64 **outupdated,
		NkMapI64 **outprevious,
		NkString **outerror);

	typedef struct NkWalletUpdate
	{
		const void *ptr;
	} NkWalletUpdate;

	typedef struct NkWalletUpdateResult
	{
		const void *ptr;
	} NkWalletUpdateResult;

	typedef int (*NkModuleWalletsUpdateFn)(
		const void *ptr,
		const NkContext *ctx,
		const NkWalletUpdate *updates,
		NkU32 numupdates,
		bool updateledger,
		NkWalletUpdateResult **outresults,
		NkU32 **outnumresults,
		NkString **outerror);

	typedef struct NkWalletLedgerItem
	{
		const void *ptr;
	} NkWalletLedgerItem;

	typedef int (*NkModuleWalletLedgerUpdateFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString itemid,
		NkMapAny metadata,
		NkWalletLedgerItem **outitem,
		NkString **outerror);

	typedef int (*NkModuleWalletLedgerListFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkU32 limit,
		NkString cursor,
		NkWalletLedgerItem **outitems,
		NkU32 **outnumitems,
		NkString **outcursor,
		NkString **outerror);

	typedef struct NkStorageObject
	{
		const void *ptr;
	} NkStorageObject;

	typedef int (*NkModuleStorageListFn)(
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

	typedef struct NkStorageRead
	{
		const void *ptr;
	} NkStorageRead;

	typedef int (*NkModuleStorageReadFn)(
		const void *ptr,
		const NkContext *ctx,
		const NkStorageRead *reads,
		NkU32 numreads,
		NkStorageObject **outobjs,
		NkU32 **outnumobjs,
		NkString **outerror);

	typedef struct NkStorageWrite
	{
		const void *ptr;
	} NkStorageWrite;

	typedef struct NkStorageObjectAck
	{
		const void *ptr;
	} NkStorageObjectAck;

	typedef int (*NkModuleStorageWriteFn)(
		const void *ptr,
		const NkContext *ctx,
		const NkStorageWrite *writes,
		NkU32 numwrites,
		NkStorageObjectAck **outacks,
		NkU32 **outnumacks,
		NkString **outerror);

	typedef struct NkStorageDelete
	{
		const void *ptr;
	} NkStorageDelete;

	typedef int (*NkModuleStorageDeleteFn)(
		const void *ptr,
		const NkContext *ctx,
		const NkStorageDelete *deletes,
		NkU32 numdeletes,
		NkString **outerror);

	typedef struct NkAccountUpdate
	{
		const void *ptr;
	} NkAccountUpdate;

	typedef int (*NkModuleMultiUpdateFn)(
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

	typedef int (*NkModuleLeaderboardCreateFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString id,
		bool authoritative,
		NkString sortorder,
		NkString op,
		NkString resetschedule,
		NkMapAny metadata,
		NkString **outerror);

	typedef int (*NkModuleLeaderboardRecordsListFn)(
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

	typedef int (*NkModuleLeaderboardRecordWriteFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString id,
		NkString ownerid,
		NkI64 score,
		NkI64 subscore,
		NkMapAny metadata,
		NkLeaderboardRecord **outrecord,
		NkString **outerror);

	typedef int (*NkModuleDeleteFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString id,
		NkString ownerid,
		NkString **outerror);

	typedef int (*NkModuleTournamentCreateFn)(
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

	typedef int (*NkModuleTournamentAddAttemptFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString id,
		NkString ownerid,
		NkU64 count,
		NkString **outerror);

	typedef int (*NkModuleTournamentJoinFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString id,
		NkString ownerid,
		NkString username,
		NkString **outerror);

	typedef struct NkTournament
	{
		const void *ptr;
	} NkTournament;

	typedef int (*NkModuleTournamentsGetIdFn)(
		const void *ptr,
		const NkContext *ctx,
		const NkString *tournamentids,
		NkU32 numtournamentids,
		NkTournament **outtournaments,
		NkU32 **outnumtournaments,
		NkString **outerror);

	typedef struct NkTournamentList
	{
		const void *ptr;
	} NkTournamentList;

	typedef int (*NkModuleTournamentListFn)(
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

	typedef struct NkLeaderboardRecord
	{
		const void *ptr;
	} NkLeaderboardRecord;

	typedef int (*NkModuleTournamentRecordsListFn)(
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

	typedef int (*NkModuleTournamentRecordWriteFn)(
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

	typedef int (*NkModuleTournamentRecordsHaystackFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString id,
		NkString ownerid,
		NkU32 limit,
		NkI64 expiry,
		NkLeaderboardRecord **outrecords,
		NkU32 **outnumrecords,
		NkString **outerror);

	typedef struct NkGroup
	{
		const void *ptr;
	} NkGroup;

	typedef int (*NkModuleGroupsGetIdFn)(
		const void *ptr,
		const NkContext *ctx,
		const NkString *groupids,
		NkU32 numgroupids,
		NkGroup **outgroups,
		NkU32 **outnumgroups,
		NkString **outerror);

	typedef int (*NkModuleGroupCreateFn)(
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

	typedef int (*NkModuleGroupUpdateFn)(
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

	typedef int (*NkModuleGroupUserFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString groupid,
		NkString userid,
		NkString username,
		NkString **outerror);

	typedef int (*NkModuleGroupUsersFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString groupid,
		const NkString *userids,
		NkU32 numuserids,
		NkString **outerror);

	typedef struct NkGroupUserListGroupUser
	{
		const void *ptr;
	} NkGroupUserListGroupUser;

	typedef int (*NkModuleGroupUsersListFn)(
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

	typedef struct NkUserGroupListUserGroup
	{
		const void *ptr;
	} NkUserGroupListUserGroup;

	typedef int (*NkModuleUserGroupsListFn)(
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

	typedef struct NkFriend
	{
		const void *ptr;
	} NkFriend;

	typedef int (*NkModuleFriendsListFn)(
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

	typedef struct NkEvent
	{
		const void *ptr;
	} NkEvent;

	typedef int (*NkModuleEventFn)(
		const void *ptr,
		const NkContext *ctx,
		NkEvent evt,
		NkString **outerror);

	typedef struct NkModule
	{
		const void *ptr;
		NkModuleAuthenticateFn authenticateapple;
		NkModuleAuthenticateFn authenticatecustom;
		NkModuleAuthenticateFn authenticatedevice;
		NkModuleAuthenticateEmailFn authenticateemail;
		NkModuleAuthenticateFacebookFn authenticatefacebook;
		NkModuleAuthenticateFn authenticatefacebookinstantgame;
		NkModuleAuthenticateGameCenterFn authenticategamecenter;
		NkModuleAuthenticateFn authenticategoogle;
		NkModuleAuthenticateFn authenticatesteam;
		NkModuleAuthenticateGenerateTokenFn authenticatetokengenerate;
		NkModuleAccountGetIdFn accountgetid;
		NkModuleAccountsGetIdFn accountsgetid;
		NkModuleAccountUpdateIdFn accountupdateid;
		NkModuleAccountDeleteIdFn accountdeleteid;
		NkModuleAccountExportIdFn accountexportid;
		NkModuleUsersGetFn usersgetid;
		NkModuleUsersGetFn usersgetusername;
		NkModuleUsersBanFn usersbanid;
		NkModuleUsersBanFn usersunbanid;
		NkModuleLinkFn linkapple;
		NkModuleLinkFn linkcustom;
		NkModuleLinkFn linkdevice;
		NkModuleLinkEmailFn linkemail;
		NkModuleLinkFacebookFn linkfacebook;
		NkModuleLinkFn linkfacebookinstantgame;
		NkModuleLinkGameCenterFn linkgamecenter;
		NkModuleLinkFn linkgoogle;
		NkModuleLinkFn linksteam;
		NkModuleLinkFn unlinkapple;
		NkModuleLinkFn unlinkcustom;
		NkModuleLinkFn unlinkdevice;
		NkModuleLinkFn unlinkemail;
		NkModuleLinkFn unlinkfacebook;
		NkModuleLinkFn unlinkfacebookinstantgame;
		NkModuleLinkGameCenterFn unlinkgamecenter;
		NkModuleLinkFn unlinkgoogle;
		NkModuleLinkFn unlinksteam;
		NkModuleStreamUserListFn streamuserlist;
		NkModuleStreamUserGetFn streamuserget;
		NkModuleStreamUserJoinFn streamuserjoin;
		NkModuleStreamUserUpdateFn streamuserupdate;
		NkModuleStreamUserLeaveFn streamuserleave;
		NkModuleStreamUserKickFn streamuserkick;
		NkModuleStreamCountFn streamcount;
		NkModuleStreamCloseFn streamclose;
		NkModuleStreamSendFn streamsend;
		NkModuleStreamSendRawFn streamsendraw;
		NkModuleSessionDisconnectFn sessiondisconnect;
		NkModuleMatchCreateFn matchcreate;
		NkModuleMatchGetFn matchget;
		NkModuleMatchListFn matchlist;
		NkModuleNotificationSendFn notificationsend;
		NkModuleNotificationsSendFn notificationssend;
		NkModuleWalletUpdateFn walletupdate;
		NkModuleWalletsUpdateFn walletsupdate;
		NkModuleWalletLedgerUpdateFn walletledgerupdate;
		NkModuleWalletLedgerListFn walletledgerlist;
		NkModuleStorageListFn storagelist;
		NkModuleStorageReadFn storageread;
		NkModuleStorageWriteFn storagewrite;
		NkModuleStorageDeleteFn storagedelete;
		NkModuleMultiUpdateFn multiupdate;
		NkModuleLeaderboardCreateFn leaderboardcreate;
		NkModuleDeleteFn leaderboarddelete;
		NkModuleLeaderboardRecordsListFn leaderboardrecordslist;
		NkModuleLeaderboardRecordWriteFn leaderboardrecordwrite;
		NkModuleDeleteFn leaderboardrecorddelete;
		NkModuleTournamentCreateFn tournamentcreate;
		NkModuleDeleteFn tournamentdelete;
		NkModuleTournamentAddAttemptFn tournamentaddattempt;
		NkModuleTournamentJoinFn tournamentjoin;
		NkModuleTournamentsGetIdFn tournamentsgetid;
		NkModuleTournamentListFn tournamentlist;
		NkModuleTournamentRecordsListFn tournamentrecordslist;
		NkModuleTournamentRecordWriteFn tournamentrecordwrite;
		NkModuleTournamentRecordsHaystackFn tournamentrecordshaystack;
		NkModuleGroupsGetIdFn groupgetid;
		NkModuleGroupCreateFn groupcreate;
		NkModuleGroupUpdateFn groupupdate;
		NkModuleDeleteFn groupdelete;
		NkModuleGroupUserFn groupuserjoin;
		NkModuleGroupUserFn groupuserleave;
		NkModuleGroupUsersFn groupusersadd;
		NkModuleGroupUsersFn groupuserskick;
		NkModuleGroupUsersFn groupuserspromote;
		NkModuleGroupUsersFn groupusersdemote;
		NkModuleGroupUsersListFn groupuserslist;
		NkModuleUserGroupsListFn usergroupslist;
		NkModuleFriendsListFn friendslist;
		NkModuleEventFn event;
	} NkModule;

	typedef int (*NkRpcFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkString payload,
		NkString **outpayload,
		NkString **outerror);

	typedef int (*NkInitializerRpcFn)(
		const void *ptr,
		NkString id,
		const NkRpcFn fn,
		NkString **outerror);

	typedef int (*NkBeforeRtCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkEnvelope envelope,
		NkEnvelope **outenvelope,
		NkString **outerror);

	typedef int (*NkInitializerBeforeRtFn)(
		const void *ptr,
		NkString id,
		const NkBeforeRtCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterRtCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkEnvelope envelope,
		NkEnvelope **outenvelope,
		NkString **outerror);

	typedef int (*NkInitializerAfterRtFn)(
		const void *ptr,
		NkString id,
		const NkAfterRtCallbackFn cb,
		NkString **outerror);

	typedef struct NkMatchmakerEntry
	{
		const void *ptr;
	} NkMatchmakerEntry;

	typedef int (*NkMatchmakerMatchedCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		const NkMatchmakerEntry *entries,
		int numentries,
		NkString **outmatchid,
		NkString **outerror);

	typedef int (*NkInitializerMatchmakerMatchedFn)(
		const void *ptr,
		const NkMatchmakerMatchedCallbackFn cb,
		NkString **outerror);

	typedef int (*NkMatchCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		void **outmatch,
		NkString **outerror); // TODO: outmatch type!

	typedef int (*NkInitializerMatchFn)(
		const void *ptr,
		NkString name,
		const NkMatchCallbackFn cb,
		NkString **outerror);

	typedef int (*NkTournamentCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkTournament tournament,
		NkI64 end,
		NkI64 reset,
		NkString **outerror);

	typedef int (*NkInitializerTournamentFn)(
		const void *ptr,
		const NkTournamentCallbackFn cb,
		NkString **outerror);

	typedef struct NkLeaderBoard
	{
		const void *ptr;
	} NkLeaderBoard;

	typedef int (*NkLeaderBoardCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkLeaderBoard leaderboard,
		NkI64 reset,
		NkString **outerror);

	typedef int (*NkInitializerLeaderBoardFn)(
		const void *ptr,
		const NkLeaderBoardCallbackFn cb,
		NkString **outerror);

	typedef int (*NkCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkString **outerror);

	typedef int (*NkInitializerBeforeGetAccountFn)(
		const void *ptr,
		const NkCallbackFn cb,
		NkString **outerror);

	typedef struct NkAccount
	{
		const void *ptr;
	} NkAccount;

	typedef int (*NkAfterGetAccountCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccount **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerAfterGetAccountFn)(
		const void *ptr,
		const NkAfterGetAccountCallbackFn cb,
		NkString **outerror);

	typedef struct NkUpdateAccountRequest
	{
		const void *ptr;
	} NkUpdateAccountRequest;

	typedef int (*NkBeforeUpdateAccountCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkUpdateAccountRequest req,
		NkUpdateAccountRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeUpdateAccountFn)(
		const void *ptr,
		const NkBeforeUpdateAccountCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterUpdateAccountCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkUpdateAccountRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterUpdateAccountFn)(
		const void *ptr,
		const NkAfterUpdateAccountCallbackFn cb,
		NkString **outerror);

	typedef struct NkSessionRefreshRequest
	{
		const void *ptr;
	} NkSessionRefreshRequest;

	typedef int (*NkBeforeSessionRefreshCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSessionRefreshRequest req,
		NkSessionRefreshRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeSessionRefreshFn)(
		const void *ptr,
		const NkBeforeSessionRefreshCallbackFn cb,
		NkString **outerror);

	typedef struct NkSession
	{
		const void *ptr;
	} NkSession;

	typedef int (*NkAfterSessionRefreshCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkSessionRefreshRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterSessionRefreshFn)(
		const void *ptr,
		const NkAfterSessionRefreshCallbackFn cb,
		NkString **outerror);

	typedef struct NkAuthenticateAppleRequest
	{
		const void *ptr;
	} NkAuthenticateAppleRequest;

	typedef int (*NkBeforeAuthenticateAppleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAuthenticateAppleRequest req,
		NkAuthenticateAppleRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeAuthenticateAppleFn)(
		const void *ptr,
		const NkBeforeAuthenticateAppleCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterAuthenticateAppleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateAppleRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterAuthenticateAppleFn)(
		const void *ptr,
		const NkAfterAuthenticateAppleCallbackFn cb,
		NkString **outerror);

	typedef struct NkAuthenticateCustomRequest
	{
		const void *ptr;
	} NkAuthenticateCustomRequest;

	typedef int (*NkBeforeAuthenticateCustomCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAuthenticateCustomRequest req,
		NkAuthenticateCustomRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeAuthenticateCustomFn)(
		const void *ptr,
		const NkBeforeAuthenticateCustomCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterAuthenticateCustomCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateCustomRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterAuthenticateCustomFn)(
		const void *ptr,
		const NkAfterAuthenticateCustomCallbackFn cb,
		NkString **outerror);

	typedef struct NkAuthenticateDeviceRequest
	{
		const void *ptr;
	} NkAuthenticateDeviceRequest;

	typedef int (*NkBeforeAuthenticateDeviceCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAuthenticateDeviceRequest req,
		NkAuthenticateDeviceRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeAuthenticateDeviceFn)(
		const void *ptr,
		const NkBeforeAuthenticateDeviceCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterAuthenticateDeviceCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateDeviceRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterAuthenticateDeviceFn)(
		const void *ptr,
		const NkAfterAuthenticateDeviceCallbackFn cb,
		NkString **outerror);

	typedef struct NkAuthenticateEmailRequest
	{
		const void *ptr;
	} NkAuthenticateEmailRequest;

	typedef int (*NkBeforeAuthenticateEmailCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAuthenticateEmailRequest req,
		NkAuthenticateEmailRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeAuthenticateEmailFn)(
		const void *ptr,
		const NkBeforeAuthenticateEmailCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterAuthenticateEmailCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateEmailRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterAuthenticateEmailFn)(
		const void *ptr,
		const NkAfterAuthenticateEmailCallbackFn cb,
		NkString **outerror);

	typedef struct NkAuthenticateFacebookRequest
	{
		const void *ptr;
	} NkAuthenticateFacebookRequest;

	typedef int (*NkBeforeAuthenticateFacebookCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAuthenticateFacebookRequest req,
		NkAuthenticateFacebookRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeAuthenticateFacebookFn)(
		const void *ptr,
		const NkBeforeAuthenticateFacebookCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterAuthenticateFacebookCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateFacebookRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterAuthenticateFacebookFn)(
		const void *ptr,
		const NkAfterAuthenticateFacebookCallbackFn cb,
		NkString **outerror);

	typedef struct NkAuthenticateFacebookInstantGameRequest
	{
		const void *ptr;
	} NkAuthenticateFacebookInstantGameRequest;

	typedef int (*NkBeforeAuthenticateFacebookInstantGameCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAuthenticateFacebookInstantGameRequest req,
		NkAuthenticateFacebookInstantGameRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeAuthenticateFacebookInstantGameFn)(
		const void *ptr,
		const NkBeforeAuthenticateFacebookInstantGameCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterAuthenticateFacebookInstantGameCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateFacebookInstantGameRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterAuthenticateFacebookInstantGameFn)(
		const void *ptr,
		const NkAfterAuthenticateFacebookInstantGameCallbackFn cb,
		NkString **outerror);

	typedef struct NkAuthenticateGameCenterRequest
	{
		const void *ptr;
	} NkAuthenticateGameCenterRequest;

	typedef int (*NkBeforeAuthenticateGameCenterCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAuthenticateGameCenterRequest req,
		NkAuthenticateGameCenterRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeAuthenticateGameCenterFn)(
		const void *ptr,
		const NkBeforeAuthenticateGameCenterCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterAuthenticateGameCenterCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateGameCenterRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterAuthenticateGameCenterFn)(
		const void *ptr,
		const NkAfterAuthenticateGameCenterCallbackFn cb,
		NkString **outerror);

	typedef struct NkAuthenticateGoogleRequest
	{
		const void *ptr;
	} NkAuthenticateGoogleRequest;

	typedef int (*NkBeforeAuthenticateGoogleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAuthenticateGoogleRequest req,
		NkAuthenticateGoogleRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeAuthenticateGoogleFn)(
		const void *ptr,
		const NkBeforeAuthenticateGoogleCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterAuthenticateGoogleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateGoogleRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterAuthenticateGoogleFn)(
		const void *ptr,
		const NkAfterAuthenticateGoogleCallbackFn cb,
		NkString **outerror);

	typedef struct NkAuthenticateSteamRequest
	{
		const void *ptr;
	} NkAuthenticateSteamRequest;

	typedef int (*NkBeforeAuthenticateSteamCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAuthenticateSteamRequest req,
		NkAuthenticateSteamRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeAuthenticateSteamFn)(
		const void *ptr,
		const NkBeforeAuthenticateSteamCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterAuthenticateSteamCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateSteamRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterAuthenticateSteamFn)(
		const void *ptr,
		const NkAfterAuthenticateSteamCallbackFn cb,
		NkString **outerror);

	typedef struct NkListChannelMessagesRequest
	{
		const void *ptr;
	} NkListChannelMessagesRequest;

	typedef int (*NkBeforeListChannelMessagesCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkListChannelMessagesRequest req,
		NkListChannelMessagesRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeListChannelMessagesFn)(
		const void *ptr,
		const NkBeforeListChannelMessagesCallbackFn cb,
		NkString **outerror);

	typedef struct NkChannelMessageList
	{
		const void *ptr;
	} NkChannelMessageList;

	typedef int (*NkAfterListChannelMessagesCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkChannelMessageList msgs,
		NkListChannelMessagesRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterListChannelMessagesFn)(
		const void *ptr,
		const NkAfterListChannelMessagesCallbackFn cb,
		NkString **outerror);

	typedef struct NkListFriendsRequest
	{
		const void *ptr;
	} NkListFriendsRequest;

	typedef int (*NkBeforeListFriendsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkListFriendsRequest req,
		NkListFriendsRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeListFriendsFn)(
		const void *ptr,
		const NkBeforeListFriendsCallbackFn cb,
		NkString **outerror);

	typedef struct NkFriendList
	{
		const void *ptr;
	} NkFriendList;

	typedef int (*NkAfterListFriendsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkFriendList friends,
		NkString **outerror);

	typedef int (*NkInitializerAfterListFriendsFn)(
		const void *ptr,
		const NkAfterListFriendsCallbackFn cb,
		NkString **outerror);

	typedef struct NkAddFriendsRequest
	{
		const void *ptr;
	} NkAddFriendsRequest;

	typedef int (*NkBeforeAddFriendsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAddFriendsRequest req,
		NkAddFriendsRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeAddFriendsFn)(
		const void *ptr,
		const NkBeforeAddFriendsCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterAddFriendsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAddFriendsRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterAddFriendsFn)(
		const void *ptr,
		const NkAfterAddFriendsCallbackFn cb,
		NkString **outerror);

	typedef struct NkDeleteFriendsRequest
	{
		const void *ptr;
	} NkDeleteFriendsRequest;

	typedef int (*NkBeforeDeleteFriendsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDeleteFriendsRequest req,
		NkDeleteFriendsRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeDeleteFriendsFn)(
		const void *ptr,
		const NkBeforeDeleteFriendsCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterDeleteFriendsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDeleteFriendsRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterDeleteFriendsFn)(
		const void *ptr,
		const NkAfterDeleteFriendsCallbackFn cb,
		NkString **outerror);

	typedef struct NkBlockFriendsRequest
	{
		const void *ptr;
	} NkBlockFriendsRequest;

	typedef int (*NkBeforeBlockFriendsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkBlockFriendsRequest req,
		NkBlockFriendsRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeBlockFriendsFn)(
		const void *ptr,
		const NkBeforeBlockFriendsCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterBlockFriendsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkBlockFriendsRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterBlockFriendsFn)(
		const void *ptr,
		const NkAfterBlockFriendsCallbackFn cb,
		NkString **outerror);

	typedef struct NkImportFacebookFriendsRequest
	{
		const void *ptr;
	} NkImportFacebookFriendsRequest;

	typedef int (*NkBeforeImportFacebookFriendsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkImportFacebookFriendsRequest req,
		NkImportFacebookFriendsRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeImportFacebookFriendsFn)(
		const void *ptr,
		const NkBeforeImportFacebookFriendsCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterImportFacebookFriendsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkImportFacebookFriendsRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterImportFacebookFriendsFn)(
		const void *ptr,
		const NkAfterImportFacebookFriendsCallbackFn cb,
		NkString **outerror);

	typedef struct NkCreateGroupRequest
	{
		const void *ptr;
	} NkCreateGroupRequest;

	typedef int (*NkBeforeCreateGroupCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkCreateGroupRequest req,
		NkCreateGroupRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeCreateGroupFn)(
		const void *ptr,
		const NkBeforeCreateGroupCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterCreateGroupCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkGroup group,
		NkCreateGroupRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterCreateGroupFn)(
		const void *ptr,
		const NkAfterCreateGroupCallbackFn cb,
		NkString **outerror);

	typedef struct NkUpdateGroupRequest
	{
		const void *ptr;
	} NkUpdateGroupRequest;

	typedef int (*NkBeforeUpdateGroupCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkUpdateGroupRequest req,
		NkUpdateGroupRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeUpdateGroupFn)(
		const void *ptr,
		const NkBeforeUpdateGroupCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterUpdateGroupCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkUpdateGroupRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterUpdateGroupFn)(
		const void *ptr,
		const NkAfterUpdateGroupCallbackFn cb,
		NkString **outerror);

	typedef struct NkDeleteGroupRequest
	{
		const void *ptr;
	} NkDeleteGroupRequest;

	typedef int (*NkBeforeDeleteGroupCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDeleteGroupRequest req,
		NkDeleteGroupRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeDeleteGroupFn)(
		const void *ptr,
		const NkBeforeDeleteGroupCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterDeleteGroupCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDeleteGroupRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterDeleteGroupFn)(
		const void *ptr,
		const NkAfterDeleteGroupCallbackFn cb,
		NkString **outerror);

	typedef struct NkJoinGroupRequest
	{
		const void *ptr;
	} NkJoinGroupRequest;

	typedef int (*NkBeforeJoinGroupCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkJoinGroupRequest req,
		NkJoinGroupRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeJoinGroupFn)(
		const void *ptr,
		const NkBeforeJoinGroupCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterJoinGroupCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkJoinGroupRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterJoinGroupFn)(
		const void *ptr,
		const NkAfterJoinGroupCallbackFn cb,
		NkString **outerror);

	typedef struct NkLeaveGroupRequest
	{
		const void *ptr;
	} NkLeaveGroupRequest;

	typedef int (*NkBeforeLeaveGroupCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkLeaveGroupRequest req,
		NkLeaveGroupRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeLeaveGroupFn)(
		const void *ptr,
		const NkBeforeLeaveGroupCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterLeaveGroupCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkLeaveGroupRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterLeaveGroupFn)(
		const void *ptr,
		const NkAfterLeaveGroupCallbackFn cb,
		NkString **outerror);

	typedef struct NkAddGroupUsersRequest
	{
		const void *ptr;
	} NkAddGroupUsersRequest;

	typedef int (*NkBeforeAddGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAddGroupUsersRequest req,
		NkAddGroupUsersRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeAddGroupUsersFn)(
		const void *ptr,
		const NkBeforeAddGroupUsersCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterAddGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAddGroupUsersRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterAddGroupUsersFn)(
		const void *ptr,
		const NkAfterAddGroupUsersCallbackFn cb,
		NkString **outerror);

	typedef struct NkBanGroupUsersRequest
	{
		const void *ptr;
	} NkBanGroupUsersRequest;

	typedef int (*NkBeforeBanGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkBanGroupUsersRequest req,
		NkBanGroupUsersRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeBanGroupUsersFn)(
		const void *ptr,
		const NkBeforeBanGroupUsersCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterBanGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkBanGroupUsersRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterBanGroupUsersFn)(
		const void *ptr,
		const NkAfterBanGroupUsersCallbackFn cb,
		NkString **outerror);

	typedef struct NkKickGroupUsersRequest
	{
		const void *ptr;
	} NkKickGroupUsersRequest;

	typedef int (*NkBeforeKickGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkKickGroupUsersRequest req,
		NkKickGroupUsersRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeKickGroupUsersFn)(
		const void *ptr,
		const NkBeforeKickGroupUsersCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterKickGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkKickGroupUsersRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterKickGroupUsersFn)(
		const void *ptr,
		const NkAfterKickGroupUsersCallbackFn cb,
		NkString **outerror);

	typedef struct NkPromoteGroupUsersRequest
	{
		const void *ptr;
	} NkPromoteGroupUsersRequest;

	typedef int (*NkBeforePromoteGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkPromoteGroupUsersRequest req,
		NkPromoteGroupUsersRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforePromoteGroupUsersFn)(
		const void *ptr,
		const NkBeforePromoteGroupUsersCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterPromoteGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkPromoteGroupUsersRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterPromoteGroupUsersFn)(
		const void *ptr,
		const NkAfterPromoteGroupUsersCallbackFn cb,
		NkString **outerror);

	typedef struct NkDemoteGroupUsersRequest
	{
		const void *ptr;
	} NkDemoteGroupUsersRequest;

	typedef int (*NkBeforeDemoteGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDemoteGroupUsersRequest req,
		NkDemoteGroupUsersRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeDemoteGroupUsersFn)(
		const void *ptr,
		const NkBeforeDemoteGroupUsersCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterDemoteGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDemoteGroupUsersRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterDemoteGroupUsersFn)(
		const void *ptr,
		const NkAfterDemoteGroupUsersCallbackFn cb,
		NkString **outerror);

	typedef struct NkListGroupUsersRequest
	{
		const void *ptr;
	} NkCreateGroupRequeNkListGroupUsersRequestst;

	typedef int (*NkBeforeListGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkListGroupUsersRequest req,
		NkListGroupUsersRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeListGroupUsersFn)(
		const void *ptr,
		const NkBeforeListGroupUsersCallbackFn cb,
		NkString **outerror);

	typedef struct NkGroupUserList
	{
		const void *ptr;
	} NkGroupUserList;

	typedef int (*NkAfterListGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkGroupUserList users,
		NkListGroupUsersRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterListGroupUsersFn)(
		const void *ptr,
		const NkAfterListGroupUsersCallbackFn cb,
		NkString **outerror);

	typedef struct NkListUserGroupsRequest
	{
		const void *ptr;
	} NkListUserGroupsRequest;

	typedef int (*NkBeforeListUserGroupsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkListUserGroupsRequest req,
		NkListUserGroupsRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeListUserGroupsFn)(
		const void *ptr,
		const NkBeforeListUserGroupsCallbackFn cb,
		NkString **outerror);

	typedef struct NkUserGroupList
	{
		const void *ptr;
	} NkUserGroupList;

	typedef int (*NkAfterListUserGroupsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkUserGroupList users,
		NkListUserGroupsRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterListUserGroupsFn)(
		const void *ptr,
		const NkAfterListUserGroupsCallbackFn cb,
		NkString **outerror);

	typedef struct NkListGroupsRequest
	{
		const void *ptr;
	} NkListGroupsRequest;

	typedef int (*NkBeforeListGroupsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkListGroupsRequest req,
		NkListGroupsRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeListGroupsFn)(
		const void *ptr,
		const NkBeforeListGroupsCallbackFn cb,
		NkString **outerror);

	typedef struct NkGroupList
	{
		const void *ptr;
	} NkGroupList;

	typedef int (*NkAfterListGroupsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkGroupList users,
		NkListGroupsRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterListGroupsFn)(
		const void *ptr,
		const NkAfterListGroupsCallbackFn cb,
		NkString **outerror);

	typedef struct NkDeleteLeaderboardRecordRequest
	{
		const void *ptr;
	} NkDeleteLeaderboardRecordRequest;

	typedef int (*NkBeforeDeleteLeaderboardRecordCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDeleteLeaderboardRecordRequest req,
		NkDeleteLeaderboardRecordRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeDeleteLeaderboardRecordFn)(
		const void *ptr,
		const NkBeforeDeleteLeaderboardRecordCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterDeleteLeaderboardRecordCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDeleteLeaderboardRecordRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterDeleteLeaderboardRecordFn)(
		const void *ptr,
		const NkAfterDeleteLeaderboardRecordCallbackFn cb,
		NkString **outerror);

	typedef struct NkListLeaderboardRecordsRequest
	{
		const void *ptr;
	} NkListLeaderboardRecordsRequest;

	typedef int (*NkBeforeListLeaderboardRecordsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkListLeaderboardRecordsRequest req,
		NkListLeaderboardRecordsRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeListLeaderboardRecordsFn)(
		const void *ptr,
		const NkBeforeListLeaderboardRecordsCallbackFn cb,
		NkString **outerror);

	typedef struct NkLeaderboardRecordList
	{
		const void *ptr;
	} NkLeaderboardRecordList;

	typedef int (*NkAfterListLeaderboardRecordsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkLeaderboardRecordList records,
		NkListLeaderboardRecordsRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterListLeaderboardRecordsFn)(
		const void *ptr,
		const NkAfterListLeaderboardRecordsCallbackFn cb,
		NkString **outerror);

	typedef struct NkWriteLeaderboardRecordRequest
	{
		const void *ptr;
	} NkWriteLeaderboardRecordRequest;

	typedef int (*NkBeforeWriteLeaderboardRecordCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkWriteLeaderboardRecordRequest req,
		NkWriteLeaderboardRecordRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeWriteLeaderboardRecordFn)(
		const void *ptr,
		const NkBeforeWriteLeaderboardRecordCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterWriteLeaderboardRecordCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkLeaderboardRecord record,
		NkWriteLeaderboardRecordRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterWriteLeaderboardRecordFn)(
		const void *ptr,
		const NkAfterWriteLeaderboardRecordCallbackFn cb,
		NkString **outerror);

	typedef struct NkListLeaderboardRecordsAroundOwnerRequest
	{
		const void *ptr;
	} NkListLeaderboardRecordsAroundOwnerRequest;

	typedef int (*NkBeforeListLeaderboardRecordsAroundOwnerCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkListLeaderboardRecordsAroundOwnerRequest req,
		NkListLeaderboardRecordsAroundOwnerRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeListLeaderboardRecordsAroundOwnerFn)(
		const void *ptr,
		const NkBeforeListLeaderboardRecordsAroundOwnerCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterListLeaderboardRecordsAroundOwnerCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkLeaderboardRecordList records,
		NkListLeaderboardRecordsAroundOwnerRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterListLeaderboardRecordsAroundOwnerFn)(
		const void *ptr,
		const NkAfterListLeaderboardRecordsAroundOwnerCallbackFn cb,
		NkString **outerror);

	typedef struct NkAccountApple
	{
		const void *ptr;
	} NkAccountApple;

	typedef int (*NkBeforeLinkAppleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountApple account,
		NkAccountApple **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeLinkAppleFn)(
		const void *ptr,
		const NkBeforeLinkAppleCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterLinkAppleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountApple account,
		NkString **outerror);

	typedef int (*NkInitializerAfterLinkAppleFn)(
		const void *ptr,
		const NkAfterLinkAppleCallbackFn cb,
		NkString **outerror);

	typedef struct NkAccountCustom
	{
		const void *ptr;
	} NkAccountCustom;

	typedef int (*NkBeforeLinkCustomCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountCustom account,
		NkAccountCustom **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeLinkCustomFn)(
		const void *ptr,
		const NkBeforeLinkCustomCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterLinkCustomCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountCustom account,
		NkString **outerror);

	typedef int (*NkInitializerAfterLinkCustomFn)(
		const void *ptr,
		const NkAfterLinkCustomCallbackFn cb,
		NkString **outerror);

	typedef struct NkAccountDevice
	{
		const void *ptr;
	} NkAccountDevice;

	typedef int (*NkBeforeLinkDeviceCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountDevice account,
		NkAccountDevice **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeLinkDeviceFn)(
		const void *ptr,
		const NkBeforeLinkDeviceCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterLinkDeviceCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountDevice account,
		NkString **outerror);

	typedef int (*NkInitializerAfterLinkDeviceFn)(
		const void *ptr,
		const NkAfterLinkDeviceCallbackFn cb,
		NkString **outerror);

	typedef struct NkAccountEmail
	{
		const void *ptr;
	} NkAccountEmail;

	typedef int (*NkBeforeLinkEmailCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountEmail account,
		NkAccountEmail **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeLinkEmailFn)(
		const void *ptr,
		const NkBeforeLinkEmailCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterLinkEmailCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountEmail account,
		NkString **outerror);

	typedef int (*NkInitializerAfterLinkEmailFn)(
		const void *ptr,
		const NkAfterLinkEmailCallbackFn cb,
		NkString **outerror);

	typedef struct NkLinkFacebookRequest
	{
		const void *ptr;
	} NkLinkFacebookRequest;

	typedef int (*NkBeforeLinkFacebookCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkLinkFacebookRequest req,
		NkLinkFacebookRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeLinkFacebookFn)(
		const void *ptr,
		const NkBeforeLinkFacebookCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterLinkFacebookCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkLinkFacebookRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterLinkFacebookFn)(
		const void *ptr,
		const NkAfterLinkFacebookCallbackFn cb,
		NkString **outerror);

	typedef struct NkAccountFacebookInstantGame
	{
		const void *ptr;
	} NkAccountFacebookInstantGame;

	typedef int (*NkBeforeLinkFacebookInstantGameCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountFacebookInstantGame account,
		NkAccountFacebookInstantGame **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeLinkFacebookInstantGameFn)(
		const void *ptr,
		const NkBeforeLinkFacebookInstantGameCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterLinkFacebookInstantGameCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountFacebookInstantGame account,
		NkString **outerror);

	typedef int (*NkInitializerAfterLinkFacebookInstantGameFn)(
		const void *ptr,
		const NkAfterLinkFacebookInstantGameCallbackFn cb,
		NkString **outerror);

	typedef struct NkAccountGameCenter
	{
		const void *ptr;
	} NkAccountGameCenter;

	typedef int (*NkBeforeLinkGameCenterCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountGameCenter account,
		NkAccountGameCenter **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeLinkGameCenterFn)(
		const void *ptr,
		const NkBeforeLinkGameCenterCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterLinkGameCenterCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountGameCenter account,
		NkString **outerror);

	typedef int (*NkInitializerAfterLinkGameCenterFn)(
		const void *ptr,
		const NkAfterLinkGameCenterCallbackFn cb,
		NkString **outerror);

	typedef struct NkAccountGoogle
	{
		const void *ptr;
	} NkAccountGoogle;

	typedef int (*NkBeforeLinkGoogleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountGoogle account,
		NkAccountGoogle **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeLinkGoogleFn)(
		const void *ptr,
		const NkBeforeLinkGoogleCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterLinkGoogleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountGoogle account,
		NkString **outerror);

	typedef int (*NkInitializerAfterLinkGoogleFn)(
		const void *ptr,
		const NkAfterLinkGoogleCallbackFn cb,
		NkString **outerror);

	typedef struct NkAccountSteam
	{
		const void *ptr;
	} NkAccountSteam;

	typedef int (*NkBeforeLinkSteamCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountSteam account,
		NkAccountSteam **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeLinkSteamFn)(
		const void *ptr,
		const NkBeforeLinkSteamCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterLinkSteamCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountSteam account,
		NkString **outerror);

	typedef int (*NkInitializerAfterLinkSteamFn)(
		const void *ptr,
		const NkAfterLinkSteamCallbackFn cb,
		NkString **outerror);

	typedef struct NkListMatchesRequest
	{
		const void *ptr;
	} NkListMatchesRequest;

	typedef int (*NkBeforeListMatchesCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkListMatchesRequest req,
		NkListMatchesRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeListMatchesFn)(
		const void *ptr,
		const NkBeforeListMatchesCallbackFn cb,
		NkString **outerror);

	typedef struct NkMatchList
	{
		const void *ptr;
	} NkMatchList;

	typedef int (*NkAfterListMatchesCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkMatchList list,
		NkListMatchesRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterListMatchesFn)(
		const void *ptr,
		const NkAfterListMatchesCallbackFn cb,
		NkString **outerror);

	typedef struct NkListNotificationsRequest
	{
		const void *ptr;
	} NkListNotificationsRequest;

	typedef int (*NkBeforeListNotificationsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkListNotificationsRequest req,
		NkListNotificationsRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeListNotificationsFn)(
		const void *ptr,
		const NkBeforeListNotificationsCallbackFn cb,
		NkString **outerror);

	typedef struct NkNotificationList
	{
		const void *ptr;
	} NkNotificationList;

	typedef int (*NkAfterListNotificationsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkNotificationList list,
		NkListNotificationsRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterListNotificationsFn)(
		const void *ptr,
		const NkAfterListNotificationsCallbackFn cb,
		NkString **outerror);

	typedef struct NkDeleteNotificationsRequest
	{
		const void *ptr;
	} NkDeleteNotificationsRequest;

	typedef int (*NkBeforeDeleteNotificationsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDeleteNotificationsRequest req,
		NkDeleteNotificationsRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeDeleteNotificationsFn)(
		const void *ptr,
		const NkBeforeDeleteNotificationsCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterDeleteNotificationsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDeleteNotificationsRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterDeleteNotificationsFn)(
		const void *ptr,
		const NkAfterDeleteNotificationsCallbackFn cb,
		NkString **outerror);

	typedef struct NkListStorageObjectsRequest
	{
		const void *ptr;
	} NkListStorageObjectsRequest;

	typedef int (*NkBeforeListStorageObjectsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkListStorageObjectsRequest req,
		NkListStorageObjectsRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeListStorageObjectsFn)(
		const void *ptr,
		const NkBeforeListStorageObjectsCallbackFn cb,
		NkString **outerror);

	typedef struct NkStorageObjectList
	{
		const void *ptr;
	} NkStorageObjectList;

	typedef int (*NkAfterListStorageObjectsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkStorageObjectList list,
		NkListStorageObjectsRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterListStorageObjectsFn)(
		const void *ptr,
		const NkAfterListStorageObjectsCallbackFn cb,
		NkString **outerror);

	typedef struct NkReadStorageObjectsRequest
	{
		const void *ptr;
	} NkReadStorageObjectsRequest;

	typedef int (*NkBeforeReadStorageObjectsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkReadStorageObjectsRequest req,
		NkReadStorageObjectsRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeReadStorageObjectsFn)(
		const void *ptr,
		const NkBeforeReadStorageObjectsCallbackFn cb,
		NkString **outerror);

	typedef struct NkStorageObjects
	{
		const void *ptr;
	} NkStorageObjects;

	typedef int (*NkAfterReadStorageObjectsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkStorageObjects objs,
		NkReadStorageObjectsRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterReadStorageObjectsFn)(
		const void *ptr,
		const NkAfterReadStorageObjectsCallbackFn cb,
		NkString **outerror);

	typedef struct NkWriteStorageObjectsRequest
	{
		const void *ptr;
	} NkWriteStorageObjectsRequest;

	typedef int (*NkBeforeWriteStorageObjectsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkWriteStorageObjectsRequest req,
		NkWriteStorageObjectsRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeWriteStorageObjectsFn)(
		const void *ptr,
		const NkBeforeWriteStorageObjectsCallbackFn cb,
		NkString **outerror);

	typedef struct NkStorageObjectAcks
	{
		const void *ptr;
	} NkStorageObjectAcks;

	typedef int (*NkAfterWriteStorageObjectsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkStorageObjectAcks acks,
		NkWriteStorageObjectsRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterWriteStorageObjectsFn)(
		const void *ptr,
		const NkAfterWriteStorageObjectsCallbackFn cb,
		NkString **outerror);

	typedef struct NkDeleteStorageObjectsRequest
	{
		const void *ptr;
	} NkDeleteStorageObjectsRequest;

	typedef int (*NkBeforeDeleteStorageObjectsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDeleteStorageObjectsRequest req,
		NkDeleteStorageObjectsRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeDeleteStorageObjectsFn)(
		const void *ptr,
		const NkBeforeDeleteStorageObjectsCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterDeleteStorageObjectsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDeleteStorageObjectsRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterDeleteStorageObjectsFn)(
		const void *ptr,
		const NkAfterDeleteStorageObjectsCallbackFn cb,
		NkString **outerror);

	typedef struct NkJoinTournamentRequest
	{
		const void *ptr;
	} NkJoinTournamentRequest;

	typedef int (*NkBeforeJoinTournamentCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkJoinTournamentRequest req,
		NkJoinTournamentRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeJoinTournamentFn)(
		const void *ptr,
		const NkBeforeJoinTournamentCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterJoinTournamentCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkJoinTournamentRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterJoinTournamentFn)(
		const void *ptr,
		const NkAfterJoinTournamentCallbackFn cb,
		NkString **outerror);

	typedef struct NkListTournamentRecordsRequest
	{
		const void *ptr;
	} NkListTournamentRecordsRequest;

	typedef int (*NkBeforeListTournamentRecordsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkListTournamentRecordsRequest req,
		NkListTournamentRecordsRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeListTournamentRecordsFn)(
		const void *ptr,
		const NkBeforeListTournamentRecordsCallbackFn cb,
		NkString **outerror);

	typedef struct NkTournamentRecordList
	{
		const void *ptr;
	} NkTournamentRecordList;

	typedef int (*NkAfterListTournamentRecordsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkTournamentRecordList list,
		NkListTournamentRecordsRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterListTournamentRecordsFn)(
		const void *ptr,
		const NkAfterListTournamentRecordsCallbackFn cb,
		NkString **outerror);

	typedef struct NkListTournamentsRequest
	{
		const void *ptr;
	} NkListTournamentsRequest;

	typedef int (*NkBeforeListTournamentsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkListTournamentsRequest req,
		NkListTournamentsRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeListTournamentsFn)(
		const void *ptr,
		const NkBeforeListTournamentsCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterListTournamentsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkTournamentList list,
		NkListTournamentsRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterListTournamentsFn)(
		const void *ptr,
		const NkAfterListTournamentsCallbackFn cb,
		NkString **outerror);

	typedef struct NkWriteTournamentRecordRequest
	{
		const void *ptr;
	} NkWriteTournamentRecordRequest;

	typedef int (*NkBeforeWriteTournamentRecordCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkWriteTournamentRecordRequest req,
		NkWriteTournamentRecordRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeWriteTournamentRecordFn)(
		const void *ptr,
		const NkBeforeWriteTournamentRecordCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterWriteTournamentRecordCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkLeaderboardRecord record,
		NkWriteTournamentRecordRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterWriteTournamentRecordFn)(
		const void *ptr,
		const NkAfterWriteTournamentRecordCallbackFn cb,
		NkString **outerror);

	typedef struct NkListTournamentRecordsAroundOwnerRequest
	{
		const void *ptr;
	} NkListTournamentRecordsAroundOwnerRequest;

	typedef int (*NkBeforeListTournamentRecordsAroundOwnerCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkListTournamentRecordsAroundOwnerRequest req,
		NkListTournamentRecordsAroundOwnerRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeListTournamentRecordsAroundOwnerFn)(
		const void *ptr,
		const NkBeforeListTournamentRecordsAroundOwnerCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterListTournamentRecordsAroundOwnerCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkTournamentRecordList list,
		NkListTournamentRecordsAroundOwnerRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterListTournamentRecordsAroundOwnerFn)(
		const void *ptr,
		const NkAfterListTournamentRecordsAroundOwnerCallbackFn cb,
		NkString **outerror);

	typedef int (*NkBeforeUnlinkAppleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountApple account,
		NkAccountApple **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeUnlinkAppleFn)(
		const void *ptr,
		const NkBeforeUnlinkAppleCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterUnlinkAppleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountApple account,
		NkString **outerror);

	typedef int (*NkInitializerAfterUnlinkAppleFn)(
		const void *ptr,
		const NkAfterUnlinkAppleCallbackFn cb,
		NkString **outerror);

	typedef int (*NkBeforeUnlinkCustomCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountCustom account,
		NkAccountCustom **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeUnlinkCustomFn)(
		const void *ptr,
		const NkBeforeUnlinkCustomCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterUnlinkCustomCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountCustom account,
		NkString **outerror);

	typedef int (*NkInitializerAfterUnlinkCustomFn)(
		const void *ptr,
		const NkAfterUnlinkCustomCallbackFn cb,
		NkString **outerror);

	typedef int (*NkBeforeUnlinkDeviceCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountDevice account,
		NkAccountDevice **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeUnlinkDeviceFn)(
		const void *ptr,
		const NkBeforeUnlinkDeviceCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterUnlinkDeviceCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountDevice account,
		NkString **outerror);

	typedef int (*NkInitializerAfterUnlinkDeviceFn)(
		const void *ptr,
		const NkAfterUnlinkDeviceCallbackFn cb,
		NkString **outerror);

	typedef int (*NkBeforeUnlinkEmailCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountEmail account,
		NkAccountEmail **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeUnlinkEmailFn)(
		const void *ptr,
		const NkBeforeUnlinkEmailCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterUnlinkEmailCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountEmail account,
		NkString **outerror);

	typedef int (*NkInitializerAfterUnlinkEmailFn)(
		const void *ptr,
		const NkAfterUnlinkEmailCallbackFn cb,
		NkString **outerror);

	typedef struct NkAccountFacebook
	{
		const void *ptr;
	} NkAccountFacebook;

	typedef int (*NkBeforeUnlinkFacebookCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountFacebook account,
		NkAccountFacebook **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeUnlinkFacebookFn)(
		const void *ptr,
		const NkBeforeUnlinkFacebookCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterUnlinkFacebookCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountFacebook account,
		NkString **outerror);

	typedef int (*NkInitializerAfterUnlinkFacebookFn)(
		const void *ptr,
		const NkAfterUnlinkFacebookCallbackFn cb,
		NkString **outerror);

	typedef int (*NkBeforeUnlinkFacebookInstantGameCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountFacebookInstantGame account,
		NkAccountFacebookInstantGame **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeUnlinkFacebookInstantGameFn)(
		const void *ptr,
		const NkBeforeUnlinkFacebookInstantGameCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterUnlinkFacebookInstantGameCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountFacebookInstantGame account,
		NkString **outerror);

	typedef int (*NkInitializerAfterUnlinkFacebookInstantGameFn)(
		const void *ptr,
		const NkAfterUnlinkFacebookInstantGameCallbackFn cb,
		NkString **outerror);

	typedef int (*NkBeforeUnlinkGameCenterCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountGameCenter account,
		NkAccountGameCenter **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeUnlinkGameCenterFn)(
		const void *ptr,
		const NkBeforeUnlinkGameCenterCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterUnlinkGameCenterCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountGameCenter account,
		NkString **outerror);

	typedef int (*NkInitializerAfterUnlinkGameCenterFn)(
		const void *ptr,
		const NkAfterUnlinkGameCenterCallbackFn cb,
		NkString **outerror);

	typedef int (*NkBeforeUnlinkGoogleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountGoogle account,
		NkAccountGoogle **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeUnlinkGoogleFn)(
		const void *ptr,
		const NkBeforeUnlinkGoogleCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterUnlinkGoogleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountGoogle account,
		NkString **outerror);

	typedef int (*NkInitializerAfterUnlinkGoogleFn)(
		const void *ptr,
		const NkAfterUnlinkGoogleCallbackFn cb,
		NkString **outerror);

	typedef int (*NkBeforeUnlinkSteamCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountSteam account,
		NkAccountSteam **outaccount,
		NkString **outerror);

	typedef int (*NkInitializerBeforeUnlinkSteamFn)(
		const void *ptr,
		const NkBeforeUnlinkSteamCallbackFn cb,
		NkString **outerror);

	typedef int (*NkAfterUnlinkSteamCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountSteam account,
		NkString **outerror);

	typedef int (*NkInitializerAfterUnlinkSteamFn)(
		const void *ptr,
		const NkAfterUnlinkSteamCallbackFn cb,
		NkString **outerror);

	typedef struct NkGetUsersRequest
	{
		const void *ptr;
	} NkGetUsersRequest;

	typedef int (*NkBeforeGetUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkGetUsersRequest req,
		NkGetUsersRequest **outreq,
		NkString **outerror);

	typedef int (*NkInitializerBeforeGetUsersFn)(
		const void *ptr,
		const NkBeforeGetUsersCallbackFn cb,
		NkString **outerror);

	typedef struct NkUsers
	{
		const void *ptr;
	} NkUsers;

	typedef int (*NkAfterGetUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkUsers users,
		NkGetUsersRequest req,
		NkString **outerror);

	typedef int (*NkInitializerAfterGetUsersFn)(
		const void *ptr,
		const NkAfterGetUsersCallbackFn cb,
		NkString **outerror);

	typedef int (*NkEventCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkEvent evt,
		NkString **outerror);

	typedef int (*NkInitializerEventFn)(
		const void *ptr,
		const NkEventCallbackFn cb,
		NkString **outerror);

	typedef int (*NkInitializerEventSessionStartFn)(
		const void *ptr,
		const NkEventCallbackFn cb,
		NkString **outerror);

	typedef int (*NkInitializerEventSessionEndFn)(
		const void *ptr,
		const NkEventCallbackFn cb,
		NkString **outerror);

	typedef struct
	{
		const void *ptr;
		NkInitializerRpcFn registerrpc;
		NkInitializerBeforeRtFn registerbeforert;
		NkInitializerAfterRtFn registerafterrt;
		NkInitializerMatchmakerMatchedFn registermatchmakermatched;
		NkInitializerMatchFn registermatch;
		NkInitializerTournamentFn registertournamentend;
		NkInitializerTournamentFn registertournamentreset;
		NkInitializerLeaderBoardFn registerleaderboardreset;
		NkInitializerBeforeGetAccountFn registerbeforegetaccount;
		NkInitializerAfterGetAccountFn registeraftergetaccount;
		NkInitializerBeforeUpdateAccountFn registerbeforeupdateaccount;
		NkInitializerAfterUpdateAccountFn registerafterupdateaccount;
		NkInitializerBeforeSessionRefreshFn registerbeforesessionrefresh;
		NkInitializerAfterSessionRefreshFn registeraftersessionrefresh;
		NkInitializerBeforeAuthenticateAppleFn registerbeforeauthenticateapple;
		NkInitializerAfterAuthenticateAppleFn registerafterauthenticateapple;
		NkInitializerBeforeAuthenticateCustomFn registerbeforeauthenticatecustom;
		NkInitializerAfterAuthenticateCustomFn registerafterauthenticatecustom;
		NkInitializerBeforeAuthenticateDeviceFn registerbeforeauthenticatedevice;
		NkInitializerAfterAuthenticateDeviceFn registerafterauthenticatedevice;
		NkInitializerBeforeAuthenticateEmailFn registerbeforeauthenticateemail;
		NkInitializerAfterAuthenticateEmailFn registerafterauthenticateemail;
		NkInitializerBeforeAuthenticateFacebookFn registerbeforeauthenticatefacebook;
		NkInitializerAfterAuthenticateFacebookFn registerafterauthenticatefacebook;
		NkInitializerBeforeAuthenticateFacebookInstantGameFn
			registerbeforeauthenticatefacebookinstantgame;
		NkInitializerAfterAuthenticateFacebookInstantGameFn
			registerafterauthenticatefacebookinstantgame;
		NkInitializerBeforeAuthenticateGameCenterFn registerbeforeauthenticategamecenter;
		NkInitializerAfterAuthenticateGameCenterFn registerafterauthenticategamecenter;
		NkInitializerBeforeAuthenticateGoogleFn registerbeforeauthenticategoogle;
		NkInitializerAfterAuthenticateGoogleFn registerafterauthenticategoogle;
		NkInitializerBeforeAuthenticateSteamFn registerbeforeauthenticatesteam;
		NkInitializerAfterAuthenticateSteamFn registerafterauthenticatesteam;
		NkInitializerBeforeListChannelMessagesFn registerbeforelistchannelmessages;
		NkInitializerAfterListChannelMessagesFn registerafterlistchannelmessages;
		NkInitializerBeforeListFriendsFn registerbeforelistfriends;
		NkInitializerAfterListFriendsFn registerafterlistfriends;
		NkInitializerBeforeAddFriendsFn registerbeforeaddfriends;
		NkInitializerAfterAddFriendsFn registerafteraddfriends;
		NkInitializerBeforeDeleteFriendsFn registerbeforedeletefriends;
		NkInitializerAfterDeleteFriendsFn registerafterdeletefriends;
		NkInitializerBeforeBlockFriendsFn registerbeforeblockfriends;
		NkInitializerAfterBlockFriendsFn registerafterblockfriends;
		NkInitializerBeforeImportFacebookFriendsFn registerbeforeimportfacebookfriends;
		NkInitializerAfterImportFacebookFriendsFn registerafterimportfacebookfriends;
		NkInitializerBeforeCreateGroupFn registerbeforecreategroup;
		NkInitializerAfterCreateGroupFn registeraftercreategroup;
		NkInitializerBeforeUpdateGroupFn registerbeforeupdategroup;
		NkInitializerAfterUpdateGroupFn registerafterupdategroup;
		NkInitializerBeforeDeleteGroupFn registerbeforedeletegroup;
		NkInitializerAfterDeleteGroupFn registerafterdeletegroup;
		NkInitializerBeforeJoinGroupFn registerbeforejoingroup;
		NkInitializerAfterJoinGroupFn registerafterjoingroup;
		NkInitializerBeforeLeaveGroupFn registerbeforeleavegroup;
		NkInitializerAfterLeaveGroupFn registerafterleavegroup;
		NkInitializerBeforeAddGroupUsersFn registerbeforeaddgroupusers;
		NkInitializerAfterAddGroupUsersFn registerafteraddgroupusers;
		NkInitializerBeforeBanGroupUsersFn registerbeforebangroupusers;
		NkInitializerAfterBanGroupUsersFn registerafterbangroupusers;
		NkInitializerBeforeKickGroupUsersFn registerbeforekickgroupusers;
		NkInitializerAfterKickGroupUsersFn registerafterkickgroupusers;
		NkInitializerBeforePromoteGroupUsersFn registerbeforepromotegroupusers;
		NkInitializerAfterPromoteGroupUsersFn registerafterpromotegroupusers;
		NkInitializerBeforeDemoteGroupUsersFn registerbeforedemotegroupusers;
		NkInitializerAfterDemoteGroupUsersFn registerafterdemotegroupusers;
		NkInitializerBeforeListGroupUsersFn registerbeforelistgroupusers;
		NkInitializerAfterListGroupUsersFn registerafterlistgroupusers;
		NkInitializerBeforeListUserGroupsFn registerbeforelistusergroups;
		NkInitializerAfterListUserGroupsFn registerafterlistusergroups;
		NkInitializerBeforeListGroupsFn registerbeforelistgroups;
		NkInitializerAfterListGroupsFn registerafterlistgroups;
		NkInitializerBeforeDeleteLeaderboardRecordFn registerbeforedeleteleaderboardrecord;
		NkInitializerAfterDeleteLeaderboardRecordFn registerafterdeleteleaderboardrecord;
		NkInitializerBeforeListLeaderboardRecordsFn registerbeforelistleaderboardrecords;
		NkInitializerAfterListLeaderboardRecordsFn registerafterlistleaderboardrecords;
		NkInitializerBeforeWriteLeaderboardRecordFn registerbeforewriteleaderboardrecord;
		NkInitializerAfterWriteLeaderboardRecordFn registerafterwriteleaderboardrecord;
		NkInitializerBeforeListLeaderboardRecordsAroundOwnerFn
			registerbeforelistleaderboardrecordsaroundowner;
		NkInitializerAfterListLeaderboardRecordsAroundOwnerFn
			registerafterlistleaderboardrecordsaroundowner;
		NkInitializerBeforeLinkAppleFn registerbeforelinkapple;
		NkInitializerAfterLinkAppleFn registerafterlinkapple;
		NkInitializerBeforeLinkCustomFn registerbeforelinkcustom;
		NkInitializerAfterLinkCustomFn registerafterlinkcustom;
		NkInitializerBeforeLinkDeviceFn registerbeforelinkdevice;
		NkInitializerAfterLinkDeviceFn registerafterlinkdevice;
		NkInitializerBeforeLinkEmailFn registerbeforelinkemail;
		NkInitializerAfterLinkEmailFn registerafterlinkemail;
		NkInitializerBeforeLinkFacebookFn registerbeforelinkfacebook;
		NkInitializerAfterLinkFacebookFn registerafterlinkfacebook;
		NkInitializerBeforeLinkFacebookInstantGameFn registerbeforelinkfacebookinstantgame;
		NkInitializerAfterLinkFacebookInstantGameFn registerafterlinkfacebookinstantgame;
		NkInitializerBeforeLinkGameCenterFn registerbeforelinkgamecenter;
		NkInitializerAfterLinkGameCenterFn registerafterlinkgamecenter;
		NkInitializerBeforeLinkGoogleFn registerbeforelinkgoogle;
		NkInitializerAfterLinkGoogleFn registerafterlinkgoogle;
		NkInitializerBeforeLinkSteamFn registerbeforelinksteam;
		NkInitializerAfterLinkSteamFn registerafterlinksteam;
		NkInitializerBeforeListMatchesFn registerbeforelistmatches;
		NkInitializerAfterListMatchesFn registerafterlistmatches;
		NkInitializerBeforeListNotificationsFn registerbeforelistnotifications;
		NkInitializerAfterListNotificationsFn registerafterlistnotifications;
		NkInitializerBeforeDeleteNotificationsFn registerbeforedeletenotifications;
		NkInitializerAfterDeleteNotificationsFn registerafterdeletenotifications;
		NkInitializerBeforeListStorageObjectsFn registerbeforeliststorageobjects;
		NkInitializerAfterListStorageObjectsFn registerafterliststorageobjects;
		NkInitializerBeforeReadStorageObjectsFn registerbeforereadstorageobjects;
		NkInitializerAfterReadStorageObjectsFn registerafterreadstorageobjects;
		NkInitializerBeforeWriteStorageObjectsFn registerbeforewritestorageobjects;
		NkInitializerAfterWriteStorageObjectsFn registerafterwritestorageobjects;
		NkInitializerBeforeDeleteStorageObjectsFn registerbeforedeletestorageobjects;
		NkInitializerAfterDeleteStorageObjectsFn registerafterdeletestorageobjects;
		NkInitializerBeforeJoinTournamentFn registerbeforejointournament;
		NkInitializerAfterJoinTournamentFn registerafterjointournament;
		NkInitializerBeforeListTournamentRecordsFn registerbeforelisttournamentrecords;
		NkInitializerAfterListTournamentRecordsFn registerafterlisttournamentrecords;
		NkInitializerBeforeListTournamentsFn registerbeforelisttournaments;
		NkInitializerAfterListTournamentsFn registerafterlisttournaments;
		NkInitializerBeforeWriteTournamentRecordFn registerbeforewritetournamentrecord;
		NkInitializerAfterWriteTournamentRecordFn registerafterwritetournamentrecord;
		NkInitializerBeforeListTournamentRecordsAroundOwnerFn
			registerbeforelisttournamentrecordsaroundowner;
		NkInitializerAfterListTournamentRecordsAroundOwnerFn
			registerafterlisttournamentrecordsaroundowner;
		NkInitializerBeforeUnlinkAppleFn registerbeforeunlinkapple;
		NkInitializerAfterUnlinkAppleFn registerafterunlinkapple;
		NkInitializerBeforeUnlinkCustomFn registerbeforeunlinkcustom;
		NkInitializerAfterUnlinkCustomFn registerafterunlinkcustom;
		NkInitializerBeforeUnlinkDeviceFn registerbeforeunlinkdevice;
		NkInitializerAfterUnlinkDeviceFn registerafterunlinkdevice;
		NkInitializerBeforeUnlinkEmailFn registerbeforeunlinkemail;
		NkInitializerAfterUnlinkEmailFn registerafterunlinkemail;
		NkInitializerBeforeUnlinkFacebookFn registerbeforeunlinkfacebook;
		NkInitializerAfterUnlinkFacebookFn registerafterunlinkfacebook;
		NkInitializerBeforeUnlinkFacebookInstantGameFn registerbeforeunlinkfacebookinstantgame;
		NkInitializerAfterUnlinkFacebookInstantGameFn registerafterunlinkfacebookinstantgame;
		NkInitializerBeforeUnlinkGameCenterFn registerbeforeunlinkgamecenter;
		NkInitializerAfterUnlinkGameCenterFn registerafterunlinkgamecenter;
		NkInitializerBeforeUnlinkGoogleFn registerbeforeunlinkgoogle;
		NkInitializerAfterUnlinkGoogleFn registerafterunlinkgoogle;
		NkInitializerBeforeUnlinkSteamFn registerbeforeunlinksteam;
		NkInitializerAfterUnlinkSteamFn registerafterunlinksteam;
		NkInitializerBeforeGetUsersFn registerbeforegetusers;
		NkInitializerAfterGetUsersFn registeraftergetusers;
		NkInitializerEventFn registerevent;
		NkInitializerEventSessionStartFn registereventsessionstart;
		NkInitializerEventSessionEndFn registereventsessionend;
	} NkInitializer;

#ifdef __cplusplus
}
#endif

#endif
