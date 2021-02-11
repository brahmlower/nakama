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
// int nkinitmodule(NkContext, NkLogger, NkDb, NkModule, NkInitializer);

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

	#define NK_U8_T 0
	#define NK_U16_T 1
	#define NK_U32_T 2
	#define NK_U64_T 3
	#define NK_I8_T 4
	#define NK_I16_T 5
	#define NK_I32_T 6
	#define NK_I64_T 7
	#define NK_F32_T 8
	#define NK_F64_T 9
	#define NK_STRING_T 10
	#define NK_LAST_T 11

	typedef union NkValue
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
	} NkValue;

	typedef struct NkAny
	{
		NkU8 ty;
		NkValue val;
	} NkAny;

	typedef struct NkMapAny
	{
		NkU32 cnt;
		NkString *keys;
		NkAny *vals;
	} NkMapAny;

	typedef struct NkMapI64
	{
		NkU32 cnt;
		NkString *keys;
		NkI64 *vals;
	} NkMapI64;

	typedef struct NkMapString
	{
		NkU32 cnt;
		NkString *keys;
		NkString *vals;
	} NkMapString;

	typedef void (*NkContextValueFn)(
		const void *ptr,
		NkString key,
		char **outvalue);

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

	typedef struct NkLogger (*NkLoggerWithFieldFn)(
		const void *ptr,
		NkString key,
		NkString value);

	typedef struct NkLogger (*NkLoggerWithFieldsFn)(
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
		char **outuserid,
		char **outusername,
		char **outerror,
		bool **outcreated);

	typedef int (*NkModuleAuthenticateFacebookFn)(
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

	typedef int (*NkModuleAuthenticateFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkString username,
		bool create,
		char **outuserid,
		char **outusername,
		char **outerror,
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
		char **outuserid,
		char **outusername,
		char **outerror,
		bool **outcreated);

	typedef int (*NkModuleAuthenticateTokenGenerateFn)(
		NkString userid,
		NkString username,
		NkI64 expiry,
		NkMapString vars,
		char **outtoken,
		NkI64 **outexpiry,
		char **outerror);

	typedef struct NkAccount
	{
		const void *ptr;
	} NkAccount;

	typedef int (*NkModuleAccountGetIdFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkAccount **outaccount,
		char **outerror);

	typedef int (*NkModuleAccountsGetIdFn)(
		const void *ptr,
		const NkContext *ctx,
		const NkString *userids,
		NkU32 numuserids,
		NkAccount **outaccounts,
		NkU32 **outnumaccounts,
		char **outerror);

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
		char **outerror);

	typedef int (*NkModuleAccountDeleteIdFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		bool recorded,
		char **outerror);

	typedef int (*NkModuleAccountExportIdFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		char **outaccount,
		char **outerror);

	typedef struct NkUser
	{
		const void *ptr;
	} NkUser;

	typedef int (*NkModuleUsersGetFn)(
		const void *ptr,
		const NkContext *ctx,
		const NkString *keys,
		NkU32 numkeys,
		NkUser **outusers,
		NkU32 **outnumusers,
		char **outerror);

	typedef int (*NkModuleUsersBanFn)(
		const void *ptr,
		const NkContext *ctx,
		const NkString *userids,
		NkU32 numids,
		char **outerror);

	typedef int (*NkModuleLinkFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkString linkid,
		char **outerror);

	typedef int (*NkModuleLinkEmailFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkString email,
		NkString password,
		char **outerror);

	typedef int (*NkModuleLinkFacebookFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkString username,
		NkString token,
		bool importfriends,
		char **outerror);

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
		char **outerror);

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
		char **outerror);

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
		char **outerror);

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
		char **outerror);

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
		char **outerror);

	typedef int (*NkModuleStreamUserLeaveFn)(
		const void *ptr,
		NkU8 mode,
		NkString subject,
		NkString subcontext,
		NkString label,
		NkString userid,
		NkString sessionid,
		char **outerror);

	typedef int (*NkModuleStreamUserKickFn)(
		const void *ptr,
		NkU8 mode,
		NkString subject,
		NkString subcontext,
		NkString label,
		NkPresence presence,
		char **outerror);

	typedef int (*NkModuleStreamCountFn)(
		const void *ptr,
		NkU8 mode,
		NkString subject,
		NkString subcontext,
		NkString label,
		NkU64 **outcount,
		char **outerror);

	typedef int (*NkModuleStreamCloseFn)(
		const void *ptr,
		NkU8 mode,
		NkString subject,
		NkString subcontext,
		NkString label,
		char **outerror);

	typedef int (*NkModuleStreamSendFn)(
		const void *ptr,
		NkU8 mode,
		NkString subject,
		NkString subcontext,
		NkString label,
		NkString data,
		const NkPresence *presences,
		NkU32 numpresences,
		bool reliable,
		char **outerror);

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
		const NkPresence *presences,
		NkU32 numpresences,
		bool reliable,
		char **outerror);

	typedef int (*NkModuleSessionDisconnectFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString sessionid,
		char **outerror);

	typedef int (*NkModuleMatchCreateFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString module,
		NkMapAny params,
		char **outmatchid,
		char **outerror);

	typedef struct NkMatch
	{
		const void *ptr;
	} NkMatch;

	typedef int (*NkModuleMatchGetFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString id,
		NkMatch **outmatch,
		char **outerror);

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
		char **outerror);

	typedef int (*NkModuleNotificationSendFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkString subject,
		NkMapAny content,
		NkU64 code,
		NkString sender,
		bool persistent,
		char **outerror);

	typedef struct NkNotificationSend
	{
		const void *ptr;
	} NkNotificationSend;

	typedef int (*NkModuleNotificationsSendFn)(
		const void *ptr,
		const NkContext *ctx,
		const NkNotificationSend *notifications,
		NkU32 numnotifications,
		char **outerror);

	typedef int (*NkModuleWalletUpdateFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkMapI64 changeset,
		NkMapAny metadata,
		bool updateledger,
		NkMapI64 **outupdated,
		NkMapI64 **outprevious,
		char **outerror);

	typedef struct NkWalletUpdate
	{
		NkString userid;
		NkMapI64 changeset;
		NkMapAny metadata;
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
		char **outerror);

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
		char **outerror);

	typedef int (*NkModuleWalletLedgerListFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString userid,
		NkU32 limit,
		NkString cursor,
		NkWalletLedgerItem **outitems,
		NkU32 **outnumitems,
		char **outcursor,
		char **outerror);

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
		char **outcursor,
		char **outerror);

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
		char **outerror);

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
		char **outerror);

	typedef struct NkStorageDelete
	{
		const void *ptr;
	} NkStorageDelete;

	typedef int (*NkModuleStorageDeleteFn)(
		const void *ptr,
		const NkContext *ctx,
		const NkStorageDelete *deletes,
		NkU32 numdeletes,
		char **outerror);

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
		char **outerror);

	typedef int (*NkModuleLeaderboardCreateFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString id,
		bool authoritative,
		NkString sortorder,
		NkString op,
		NkString resetschedule,
		NkMapAny metadata,
		char **outerror);

	typedef struct NkLeaderboardRecord
	{
		const void *ptr;
	} NkLeaderboardRecord;

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
		char **outnextcursor,
		char **outprevcursor,
		char **outerror);

	typedef int (*NkModuleLeaderboardRecordWriteFn)(
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

	typedef int (*NkModuleDeleteFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString id,
		char **outerror);

	typedef int (*NkModuleLeaderboardRecordDeleteFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString id,
		NkString ownerid,
		char **outerror);

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
		char **outerror);

	typedef int (*NkModuleTournamentAddAttemptFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString id,
		NkString ownerid,
		NkU64 count,
		char **outerror);

	typedef int (*NkModuleTournamentJoinFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString id,
		NkString ownerid,
		NkString username,
		char **outerror);

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
		char **outerror);

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
		char **outerror);

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
		char **outnextcursor,
		char **outprevcursor,
		char **outerror);

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
		char **outerror);

	typedef int (*NkModuleTournamentRecordsHaystackFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString id,
		NkString ownerid,
		NkU32 limit,
		NkI64 expiry,
		NkLeaderboardRecord **outrecords,
		NkU32 **outnumrecords,
		char **outerror);

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
		char **outerror);

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
		char **outerror);

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
		char **outerror);

	typedef int (*NkModuleGroupUserFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString groupid,
		NkString userid,
		NkString username,
		char **outerror);

	typedef int (*NkModuleGroupUsersFn)(
		const void *ptr,
		const NkContext *ctx,
		NkString groupid,
		const NkString *userids,
		NkU32 numuserids,
		char **outerror);

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
		char **outcursor,
		char **outerror);

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
		char **outcursor,
		char **outerror);

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
		char **outcursor,
		char **outerror);

	typedef struct NkEvent
	{
		const void *ptr;
	} NkEvent;

	typedef int (*NkModuleEventFn)(
		const void *ptr,
		const NkContext *ctx,
		NkEvent evt,
		char **outerror);

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
		NkModuleAuthenticateTokenGenerateFn authenticatetokengenerate;
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
		NkModuleLeaderboardRecordDeleteFn leaderboardrecorddelete;
		NkModuleTournamentCreateFn tournamentcreate;
		NkModuleDeleteFn tournamentdelete;
		NkModuleTournamentAddAttemptFn tournamentaddattempt;
		NkModuleTournamentJoinFn tournamentjoin;
		NkModuleTournamentsGetIdFn tournamentsgetid;
		NkModuleTournamentListFn tournamentlist;
		NkModuleTournamentRecordsListFn tournamentrecordslist;
		NkModuleTournamentRecordWriteFn tournamentrecordwrite;
		NkModuleTournamentRecordsHaystackFn tournamentrecordshaystack;
		NkModuleGroupsGetIdFn groupsgetid;
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
		char **outpayload,
		char **outerror);

	typedef int (*NkInitializerRpcFn)(
		const void *ptr,
		NkString id,
		const NkRpcFn fn,
		char **outerror);

	typedef int (*NkBeforeRtCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkEnvelope envelope,
		NkEnvelope **outenvelope,
		char **outerror);

	typedef int (*NkInitializerBeforeRtFn)(
		const void *ptr,
		NkString id,
		const NkBeforeRtCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterRtCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkEnvelope envelope,
		NkEnvelope **outenvelope,
		char **outerror);

	typedef int (*NkInitializerAfterRtFn)(
		const void *ptr,
		NkString id,
		const NkAfterRtCallbackFn cb,
		char **outerror);

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
		char **outmatchid,
		char **outerror);

	typedef int (*NkInitializerMatchmakerMatchedFn)(
		const void *ptr,
		const NkMatchmakerMatchedCallbackFn cb,
		char **outerror);

	typedef int (*NkMatchCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		void **outmatch,
		char **outerror); // TODO: outmatch type!

	typedef int (*NkInitializerMatchFn)(
		const void *ptr,
		NkString name,
		const NkMatchCallbackFn cb,
		char **outerror);

	typedef int (*NkTournamentCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkTournament tournament,
		NkI64 end,
		NkI64 reset,
		char **outerror);

	typedef int (*NkInitializerTournamentFn)(
		const void *ptr,
		const NkTournamentCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerLeaderBoardFn)(
		const void *ptr,
		const NkLeaderBoardCallbackFn cb,
		char **outerror);

	typedef int (*NkCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		char **outerror);

	typedef int (*NkInitializerBeforeGetAccountFn)(
		const void *ptr,
		const NkCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterGetAccountCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccount **outaccount,
		char **outerror);

	typedef int (*NkInitializerAfterGetAccountFn)(
		const void *ptr,
		const NkAfterGetAccountCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeUpdateAccountFn)(
		const void *ptr,
		const NkBeforeUpdateAccountCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterUpdateAccountCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkUpdateAccountRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterUpdateAccountFn)(
		const void *ptr,
		const NkAfterUpdateAccountCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeSessionRefreshFn)(
		const void *ptr,
		const NkBeforeSessionRefreshCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerAfterSessionRefreshFn)(
		const void *ptr,
		const NkAfterSessionRefreshCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeAuthenticateAppleFn)(
		const void *ptr,
		const NkBeforeAuthenticateAppleCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterAuthenticateAppleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateAppleRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterAuthenticateAppleFn)(
		const void *ptr,
		const NkAfterAuthenticateAppleCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeAuthenticateCustomFn)(
		const void *ptr,
		const NkBeforeAuthenticateCustomCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterAuthenticateCustomCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateCustomRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterAuthenticateCustomFn)(
		const void *ptr,
		const NkAfterAuthenticateCustomCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeAuthenticateDeviceFn)(
		const void *ptr,
		const NkBeforeAuthenticateDeviceCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterAuthenticateDeviceCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateDeviceRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterAuthenticateDeviceFn)(
		const void *ptr,
		const NkAfterAuthenticateDeviceCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeAuthenticateEmailFn)(
		const void *ptr,
		const NkBeforeAuthenticateEmailCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterAuthenticateEmailCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateEmailRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterAuthenticateEmailFn)(
		const void *ptr,
		const NkAfterAuthenticateEmailCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeAuthenticateFacebookFn)(
		const void *ptr,
		const NkBeforeAuthenticateFacebookCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterAuthenticateFacebookCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateFacebookRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterAuthenticateFacebookFn)(
		const void *ptr,
		const NkAfterAuthenticateFacebookCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeAuthenticateFacebookInstantGameFn)(
		const void *ptr,
		const NkBeforeAuthenticateFacebookInstantGameCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterAuthenticateFacebookInstantGameCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateFacebookInstantGameRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterAuthenticateFacebookInstantGameFn)(
		const void *ptr,
		const NkAfterAuthenticateFacebookInstantGameCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeAuthenticateGameCenterFn)(
		const void *ptr,
		const NkBeforeAuthenticateGameCenterCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterAuthenticateGameCenterCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateGameCenterRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterAuthenticateGameCenterFn)(
		const void *ptr,
		const NkAfterAuthenticateGameCenterCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeAuthenticateGoogleFn)(
		const void *ptr,
		const NkBeforeAuthenticateGoogleCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterAuthenticateGoogleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateGoogleRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterAuthenticateGoogleFn)(
		const void *ptr,
		const NkAfterAuthenticateGoogleCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeAuthenticateSteamFn)(
		const void *ptr,
		const NkBeforeAuthenticateSteamCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterAuthenticateSteamCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkSession session,
		NkAuthenticateSteamRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterAuthenticateSteamFn)(
		const void *ptr,
		const NkAfterAuthenticateSteamCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeListChannelMessagesFn)(
		const void *ptr,
		const NkBeforeListChannelMessagesCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerAfterListChannelMessagesFn)(
		const void *ptr,
		const NkAfterListChannelMessagesCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeListFriendsFn)(
		const void *ptr,
		const NkBeforeListFriendsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerAfterListFriendsFn)(
		const void *ptr,
		const NkAfterListFriendsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeAddFriendsFn)(
		const void *ptr,
		const NkBeforeAddFriendsCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterAddFriendsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAddFriendsRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterAddFriendsFn)(
		const void *ptr,
		const NkAfterAddFriendsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeDeleteFriendsFn)(
		const void *ptr,
		const NkBeforeDeleteFriendsCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterDeleteFriendsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDeleteFriendsRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterDeleteFriendsFn)(
		const void *ptr,
		const NkAfterDeleteFriendsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeBlockFriendsFn)(
		const void *ptr,
		const NkBeforeBlockFriendsCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterBlockFriendsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkBlockFriendsRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterBlockFriendsFn)(
		const void *ptr,
		const NkAfterBlockFriendsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeImportFacebookFriendsFn)(
		const void *ptr,
		const NkBeforeImportFacebookFriendsCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterImportFacebookFriendsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkImportFacebookFriendsRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterImportFacebookFriendsFn)(
		const void *ptr,
		const NkAfterImportFacebookFriendsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeCreateGroupFn)(
		const void *ptr,
		const NkBeforeCreateGroupCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterCreateGroupCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkGroup group,
		NkCreateGroupRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterCreateGroupFn)(
		const void *ptr,
		const NkAfterCreateGroupCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeUpdateGroupFn)(
		const void *ptr,
		const NkBeforeUpdateGroupCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterUpdateGroupCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkUpdateGroupRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterUpdateGroupFn)(
		const void *ptr,
		const NkAfterUpdateGroupCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeDeleteGroupFn)(
		const void *ptr,
		const NkBeforeDeleteGroupCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterDeleteGroupCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDeleteGroupRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterDeleteGroupFn)(
		const void *ptr,
		const NkAfterDeleteGroupCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeJoinGroupFn)(
		const void *ptr,
		const NkBeforeJoinGroupCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterJoinGroupCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkJoinGroupRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterJoinGroupFn)(
		const void *ptr,
		const NkAfterJoinGroupCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeLeaveGroupFn)(
		const void *ptr,
		const NkBeforeLeaveGroupCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterLeaveGroupCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkLeaveGroupRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterLeaveGroupFn)(
		const void *ptr,
		const NkAfterLeaveGroupCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeAddGroupUsersFn)(
		const void *ptr,
		const NkBeforeAddGroupUsersCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterAddGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAddGroupUsersRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterAddGroupUsersFn)(
		const void *ptr,
		const NkAfterAddGroupUsersCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeBanGroupUsersFn)(
		const void *ptr,
		const NkBeforeBanGroupUsersCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterBanGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkBanGroupUsersRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterBanGroupUsersFn)(
		const void *ptr,
		const NkAfterBanGroupUsersCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeKickGroupUsersFn)(
		const void *ptr,
		const NkBeforeKickGroupUsersCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterKickGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkKickGroupUsersRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterKickGroupUsersFn)(
		const void *ptr,
		const NkAfterKickGroupUsersCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforePromoteGroupUsersFn)(
		const void *ptr,
		const NkBeforePromoteGroupUsersCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterPromoteGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkPromoteGroupUsersRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterPromoteGroupUsersFn)(
		const void *ptr,
		const NkAfterPromoteGroupUsersCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeDemoteGroupUsersFn)(
		const void *ptr,
		const NkBeforeDemoteGroupUsersCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterDemoteGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDemoteGroupUsersRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterDemoteGroupUsersFn)(
		const void *ptr,
		const NkAfterDemoteGroupUsersCallbackFn cb,
		char **outerror);

	typedef struct NkListGroupUsersRequest
	{
		const void *ptr;
	} NkListGroupUsersRequest;

	typedef int (*NkBeforeListGroupUsersCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkListGroupUsersRequest req,
		NkListGroupUsersRequest **outreq,
		char **outerror);

	typedef int (*NkInitializerBeforeListGroupUsersFn)(
		const void *ptr,
		const NkBeforeListGroupUsersCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerAfterListGroupUsersFn)(
		const void *ptr,
		const NkAfterListGroupUsersCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeListUserGroupsFn)(
		const void *ptr,
		const NkBeforeListUserGroupsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerAfterListUserGroupsFn)(
		const void *ptr,
		const NkAfterListUserGroupsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeListGroupsFn)(
		const void *ptr,
		const NkBeforeListGroupsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerAfterListGroupsFn)(
		const void *ptr,
		const NkAfterListGroupsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeDeleteLeaderboardRecordFn)(
		const void *ptr,
		const NkBeforeDeleteLeaderboardRecordCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterDeleteLeaderboardRecordCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDeleteLeaderboardRecordRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterDeleteLeaderboardRecordFn)(
		const void *ptr,
		const NkAfterDeleteLeaderboardRecordCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeListLeaderboardRecordsFn)(
		const void *ptr,
		const NkBeforeListLeaderboardRecordsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerAfterListLeaderboardRecordsFn)(
		const void *ptr,
		const NkAfterListLeaderboardRecordsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeWriteLeaderboardRecordFn)(
		const void *ptr,
		const NkBeforeWriteLeaderboardRecordCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterWriteLeaderboardRecordCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkLeaderboardRecord record,
		NkWriteLeaderboardRecordRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterWriteLeaderboardRecordFn)(
		const void *ptr,
		const NkAfterWriteLeaderboardRecordCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeListLeaderboardRecordsAroundOwnerFn)(
		const void *ptr,
		const NkBeforeListLeaderboardRecordsAroundOwnerCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterListLeaderboardRecordsAroundOwnerCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkLeaderboardRecordList records,
		NkListLeaderboardRecordsAroundOwnerRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterListLeaderboardRecordsAroundOwnerFn)(
		const void *ptr,
		const NkAfterListLeaderboardRecordsAroundOwnerCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeLinkAppleFn)(
		const void *ptr,
		const NkBeforeLinkAppleCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterLinkAppleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountApple account,
		char **outerror);

	typedef int (*NkInitializerAfterLinkAppleFn)(
		const void *ptr,
		const NkAfterLinkAppleCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeLinkCustomFn)(
		const void *ptr,
		const NkBeforeLinkCustomCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterLinkCustomCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountCustom account,
		char **outerror);

	typedef int (*NkInitializerAfterLinkCustomFn)(
		const void *ptr,
		const NkAfterLinkCustomCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeLinkDeviceFn)(
		const void *ptr,
		const NkBeforeLinkDeviceCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterLinkDeviceCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountDevice account,
		char **outerror);

	typedef int (*NkInitializerAfterLinkDeviceFn)(
		const void *ptr,
		const NkAfterLinkDeviceCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeLinkEmailFn)(
		const void *ptr,
		const NkBeforeLinkEmailCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterLinkEmailCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountEmail account,
		char **outerror);

	typedef int (*NkInitializerAfterLinkEmailFn)(
		const void *ptr,
		const NkAfterLinkEmailCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeLinkFacebookFn)(
		const void *ptr,
		const NkBeforeLinkFacebookCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterLinkFacebookCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkLinkFacebookRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterLinkFacebookFn)(
		const void *ptr,
		const NkAfterLinkFacebookCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeLinkFacebookInstantGameFn)(
		const void *ptr,
		const NkBeforeLinkFacebookInstantGameCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterLinkFacebookInstantGameCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountFacebookInstantGame account,
		char **outerror);

	typedef int (*NkInitializerAfterLinkFacebookInstantGameFn)(
		const void *ptr,
		const NkAfterLinkFacebookInstantGameCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeLinkGameCenterFn)(
		const void *ptr,
		const NkBeforeLinkGameCenterCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterLinkGameCenterCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountGameCenter account,
		char **outerror);

	typedef int (*NkInitializerAfterLinkGameCenterFn)(
		const void *ptr,
		const NkAfterLinkGameCenterCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeLinkGoogleFn)(
		const void *ptr,
		const NkBeforeLinkGoogleCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterLinkGoogleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountGoogle account,
		char **outerror);

	typedef int (*NkInitializerAfterLinkGoogleFn)(
		const void *ptr,
		const NkAfterLinkGoogleCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeLinkSteamFn)(
		const void *ptr,
		const NkBeforeLinkSteamCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterLinkSteamCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountSteam account,
		char **outerror);

	typedef int (*NkInitializerAfterLinkSteamFn)(
		const void *ptr,
		const NkAfterLinkSteamCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeListMatchesFn)(
		const void *ptr,
		const NkBeforeListMatchesCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerAfterListMatchesFn)(
		const void *ptr,
		const NkAfterListMatchesCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeListNotificationsFn)(
		const void *ptr,
		const NkBeforeListNotificationsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerAfterListNotificationsFn)(
		const void *ptr,
		const NkAfterListNotificationsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeDeleteNotificationsFn)(
		const void *ptr,
		const NkBeforeDeleteNotificationsCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterDeleteNotificationsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDeleteNotificationsRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterDeleteNotificationsFn)(
		const void *ptr,
		const NkAfterDeleteNotificationsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeListStorageObjectsFn)(
		const void *ptr,
		const NkBeforeListStorageObjectsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerAfterListStorageObjectsFn)(
		const void *ptr,
		const NkAfterListStorageObjectsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeReadStorageObjectsFn)(
		const void *ptr,
		const NkBeforeReadStorageObjectsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerAfterReadStorageObjectsFn)(
		const void *ptr,
		const NkAfterReadStorageObjectsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeWriteStorageObjectsFn)(
		const void *ptr,
		const NkBeforeWriteStorageObjectsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerAfterWriteStorageObjectsFn)(
		const void *ptr,
		const NkAfterWriteStorageObjectsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeDeleteStorageObjectsFn)(
		const void *ptr,
		const NkBeforeDeleteStorageObjectsCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterDeleteStorageObjectsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkDeleteStorageObjectsRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterDeleteStorageObjectsFn)(
		const void *ptr,
		const NkAfterDeleteStorageObjectsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeJoinTournamentFn)(
		const void *ptr,
		const NkBeforeJoinTournamentCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterJoinTournamentCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkJoinTournamentRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterJoinTournamentFn)(
		const void *ptr,
		const NkAfterJoinTournamentCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeListTournamentRecordsFn)(
		const void *ptr,
		const NkBeforeListTournamentRecordsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerAfterListTournamentRecordsFn)(
		const void *ptr,
		const NkAfterListTournamentRecordsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeListTournamentsFn)(
		const void *ptr,
		const NkBeforeListTournamentsCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterListTournamentsCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkTournamentList list,
		NkListTournamentsRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterListTournamentsFn)(
		const void *ptr,
		const NkAfterListTournamentsCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeWriteTournamentRecordFn)(
		const void *ptr,
		const NkBeforeWriteTournamentRecordCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterWriteTournamentRecordCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkLeaderboardRecord record,
		NkWriteTournamentRecordRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterWriteTournamentRecordFn)(
		const void *ptr,
		const NkAfterWriteTournamentRecordCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeListTournamentRecordsAroundOwnerFn)(
		const void *ptr,
		const NkBeforeListTournamentRecordsAroundOwnerCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterListTournamentRecordsAroundOwnerCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkTournamentRecordList list,
		NkListTournamentRecordsAroundOwnerRequest req,
		char **outerror);

	typedef int (*NkInitializerAfterListTournamentRecordsAroundOwnerFn)(
		const void *ptr,
		const NkAfterListTournamentRecordsAroundOwnerCallbackFn cb,
		char **outerror);

	typedef int (*NkBeforeUnlinkAppleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountApple account,
		NkAccountApple **outaccount,
		char **outerror);

	typedef int (*NkInitializerBeforeUnlinkAppleFn)(
		const void *ptr,
		const NkBeforeUnlinkAppleCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterUnlinkAppleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountApple account,
		char **outerror);

	typedef int (*NkInitializerAfterUnlinkAppleFn)(
		const void *ptr,
		const NkAfterUnlinkAppleCallbackFn cb,
		char **outerror);

	typedef int (*NkBeforeUnlinkCustomCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountCustom account,
		NkAccountCustom **outaccount,
		char **outerror);

	typedef int (*NkInitializerBeforeUnlinkCustomFn)(
		const void *ptr,
		const NkBeforeUnlinkCustomCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterUnlinkCustomCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountCustom account,
		char **outerror);

	typedef int (*NkInitializerAfterUnlinkCustomFn)(
		const void *ptr,
		const NkAfterUnlinkCustomCallbackFn cb,
		char **outerror);

	typedef int (*NkBeforeUnlinkDeviceCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountDevice account,
		NkAccountDevice **outaccount,
		char **outerror);

	typedef int (*NkInitializerBeforeUnlinkDeviceFn)(
		const void *ptr,
		const NkBeforeUnlinkDeviceCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterUnlinkDeviceCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountDevice account,
		char **outerror);

	typedef int (*NkInitializerAfterUnlinkDeviceFn)(
		const void *ptr,
		const NkAfterUnlinkDeviceCallbackFn cb,
		char **outerror);

	typedef int (*NkBeforeUnlinkEmailCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountEmail account,
		NkAccountEmail **outaccount,
		char **outerror);

	typedef int (*NkInitializerBeforeUnlinkEmailFn)(
		const void *ptr,
		const NkBeforeUnlinkEmailCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterUnlinkEmailCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountEmail account,
		char **outerror);

	typedef int (*NkInitializerAfterUnlinkEmailFn)(
		const void *ptr,
		const NkAfterUnlinkEmailCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeUnlinkFacebookFn)(
		const void *ptr,
		const NkBeforeUnlinkFacebookCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterUnlinkFacebookCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountFacebook account,
		char **outerror);

	typedef int (*NkInitializerAfterUnlinkFacebookFn)(
		const void *ptr,
		const NkAfterUnlinkFacebookCallbackFn cb,
		char **outerror);

	typedef int (*NkBeforeUnlinkFacebookInstantGameCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountFacebookInstantGame account,
		NkAccountFacebookInstantGame **outaccount,
		char **outerror);

	typedef int (*NkInitializerBeforeUnlinkFacebookInstantGameFn)(
		const void *ptr,
		const NkBeforeUnlinkFacebookInstantGameCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterUnlinkFacebookInstantGameCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountFacebookInstantGame account,
		char **outerror);

	typedef int (*NkInitializerAfterUnlinkFacebookInstantGameFn)(
		const void *ptr,
		const NkAfterUnlinkFacebookInstantGameCallbackFn cb,
		char **outerror);

	typedef int (*NkBeforeUnlinkGameCenterCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountGameCenter account,
		NkAccountGameCenter **outaccount,
		char **outerror);

	typedef int (*NkInitializerBeforeUnlinkGameCenterFn)(
		const void *ptr,
		const NkBeforeUnlinkGameCenterCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterUnlinkGameCenterCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountGameCenter account,
		char **outerror);

	typedef int (*NkInitializerAfterUnlinkGameCenterFn)(
		const void *ptr,
		const NkAfterUnlinkGameCenterCallbackFn cb,
		char **outerror);

	typedef int (*NkBeforeUnlinkGoogleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountGoogle account,
		NkAccountGoogle **outaccount,
		char **outerror);

	typedef int (*NkInitializerBeforeUnlinkGoogleFn)(
		const void *ptr,
		const NkBeforeUnlinkGoogleCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterUnlinkGoogleCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountGoogle account,
		char **outerror);

	typedef int (*NkInitializerAfterUnlinkGoogleFn)(
		const void *ptr,
		const NkAfterUnlinkGoogleCallbackFn cb,
		char **outerror);

	typedef int (*NkBeforeUnlinkSteamCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountSteam account,
		NkAccountSteam **outaccount,
		char **outerror);

	typedef int (*NkInitializerBeforeUnlinkSteamFn)(
		const void *ptr,
		const NkBeforeUnlinkSteamCallbackFn cb,
		char **outerror);

	typedef int (*NkAfterUnlinkSteamCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkDb db,
		NkModule nk,
		NkAccountSteam account,
		char **outerror);

	typedef int (*NkInitializerAfterUnlinkSteamFn)(
		const void *ptr,
		const NkAfterUnlinkSteamCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerBeforeGetUsersFn)(
		const void *ptr,
		const NkBeforeGetUsersCallbackFn cb,
		char **outerror);

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
		char **outerror);

	typedef int (*NkInitializerAfterGetUsersFn)(
		const void *ptr,
		const NkAfterGetUsersCallbackFn cb,
		char **outerror);

	typedef int (*NkEventCallbackFn)(
		NkContext ctx,
		NkLogger logger,
		NkEvent evt,
		char **outerror);

	typedef int (*NkInitializerEventFn)(
		const void *ptr,
		const NkEventCallbackFn cb,
		char **outerror);

	typedef int (*NkInitializerEventSessionStartFn)(
		const void *ptr,
		const NkEventCallbackFn cb,
		char **outerror);

	typedef int (*NkInitializerEventSessionEndFn)(
		const void *ptr,
		const NkEventCallbackFn cb,
		char **outerror);

	typedef struct NkInitializer
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
