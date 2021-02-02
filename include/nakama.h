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

// NOTE: In order to implement a c-module, you must provide the following functions:
//
// Module entrypoint:
// void nk_init_module(NkContext, NkLogger, NkDb, NkModule, NkInitializer);
//
// Match initializer:
// void *nk_init_match(NkContext, NkLogger, NkDb, NkModule);

#ifndef NAKAMA_H
#define NAKAMA_H

#include <stdbool.h>
#include <stddef.h>
#include "hashmap.h"

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

	//--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--//

	typedef void (*NkContextValueFn)(void *ptr, NkString key, NkString *outvalue);

	typedef struct
	{
		void *ptr;
		NkContextValueFn value;
	} NkContext;

	//--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--//

	typedef struct hashmap_s (*NkLoggerFieldsFn)(void *ptr);

	typedef void (*NkLoggerLevelFn)(void *ptr, NkString s);

	typedef struct NkLogger (*NkLoggerWithFieldFn)(void *ptr, NkString key, NkString value);

	typedef struct NkLogger (*NkLoggerWithFieldsFn)(void *ptr, struct hashmap_s fields);

	typedef struct NkLogger
	{
		void *ptr;
		NkLoggerLevelFn debug;
		NkLoggerLevelFn error;
		NkLoggerFieldsFn fields;
		NkLoggerLevelFn info;
		NkLoggerLevelFn warn;
		NkLoggerWithFieldFn withfield;
		NkLoggerWithFieldsFn withfields;
	} NkLogger;

	//--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--//

	typedef struct
	{
		void *ptr;
	} NkDb;

	//--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--//

	typedef int (*NkModuleAuthenticateEmailFn)(void *ptr, NkContext *ctx, NkString email,
											   NkString password, NkString username, bool create,
											   NkString *outuserid, NkString *outusername,
											   NkString *outerror, bool *outcreated);

	typedef int (*NkModuleAuthenticateFacebookFn)(void *ptr, NkContext *ctx, NkString token,
												  bool importfriends, NkString username,
												  bool create, NkString *outuserid,
												  NkString *outusername, NkString *outerror,
												  bool *outcreated);

	typedef int (*NkModuleAuthenticateFn)(void *ptr, NkContext *ctx, NkString userid,
										  NkString username, bool create, NkString *outuserid,
										  NkString *outusername, NkString *outerror,
										  bool *outcreated);

	typedef int (*NkModuleAuthenticateGameCenterFn)(void *ptr, NkContext *ctx, NkString playerid,
													NkString bundleid, NkI64 timestamp,
													NkString salt, NkString signature,
													NkString publickeyurl, NkString username,
													bool create, NkString *outuserid,
													NkString *outusername, NkString *outerror,
													bool *outcreated);

	typedef struct
	{
		void *ptr;
		NkModuleAuthenticateFn authenticateapple;
		NkModuleAuthenticateFn authenticatecustom;
		NkModuleAuthenticateFn authenticatedevice;
		NkModuleAuthenticateEmailFn authenticateemail;
		NkModuleAuthenticateFacebookFn authenticatefacebook;
		NkModuleAuthenticateFn authenticatefacebookinstantgame;
		NkModuleAuthenticateGameCenterFn authenticategamecenter;
		NkModuleAuthenticateFn authenticategoogle;
		NkModuleAuthenticateFn authenticatesteam;

		// NkModuleAuthenticateGenerateTokenFn authenticatetokengenerate;// AuthenticateTokenGenerate(userID, username string, exp int64, vars map[string]string) (string, int64, error)

		// AccountGetId(ctx context.Context, userID string) (*api.Account, error)
		// AccountsGetId(ctx context.Context, userIDs []string) ([]*api.Account, error)
		// AccountUpdateId(ctx context.Context, userID, username string, metadata map[string]interface{}, displayName, timezone, location, langTag, avatarUrl string) error

		// AccountDeleteId(ctx context.Context, userID string, recorded bool) error
		// AccountExportId(ctx context.Context, userID string) (string, error)

		// UsersGetId(ctx context.Context, userIDs []string) ([]*api.User, error)
		// UsersGetUsername(ctx context.Context, usernames []string) ([]*api.User, error)
		// UsersBanId(ctx context.Context, userIDs []string) error
		// UsersUnbanId(ctx context.Context, userIDs []string) error

		// LinkApple(ctx context.Context, userID, token string) error
		// LinkCustom(ctx context.Context, userID, customID string) error
		// LinkDevice(ctx context.Context, userID, deviceID string) error
		// LinkEmail(ctx context.Context, userID, email, password string) error
		// LinkFacebook(ctx context.Context, userID, username, token string, importFriends bool) error
		// LinkFacebookInstantGame(ctx context.Context, userID, signedPlayerInfo string) error
		// LinkGameCenter(ctx context.Context, userID, playerID, bundleID string, timestamp int64, salt, signature, publicKeyUrl string) error
		// LinkGoogle(ctx context.Context, userID, token string) error
		// LinkSteam(ctx context.Context, userID, token string) error

		// ReadFile(path string) (*os.File, error)

		// UnlinkApple(ctx context.Context, userID, token string) error
		// UnlinkCustom(ctx context.Context, userID, customID string) error
		// UnlinkDevice(ctx context.Context, userID, deviceID string) error
		// UnlinkEmail(ctx context.Context, userID, email string) error
		// UnlinkFacebook(ctx context.Context, userID, token string) error
		// UnlinkFacebookInstantGame(ctx context.Context, userID, signedPlayerInfo string) error
		// UnlinkGameCenter(ctx context.Context, userID, playerID, bundleID string, timestamp int64, salt, signature, publicKeyUrl string) error
		// UnlinkGoogle(ctx context.Context, userID, token string) error
		// UnlinkSteam(ctx context.Context, userID, token string) error

		// StreamUserList(mode uint8, subject, subcontext, label string, includeHidden, includeNotHidden bool) ([]Presence, error)
		// StreamUserGet(mode uint8, subject, subcontext, label, userID, sessionID string) (PresenceMeta, error)
		// StreamUserJoin(mode uint8, subject, subcontext, label, userID, sessionID string, hidden, persistence bool, status string) (bool, error)
		// StreamUserUpdate(mode uint8, subject, subcontext, label, userID, sessionID string, hidden, persistence bool, status string) error
		// StreamUserLeave(mode uint8, subject, subcontext, label, userID, sessionID string) error
		// StreamUserKick(mode uint8, subject, subcontext, label string, presence Presence) error
		// StreamCount(mode uint8, subject, subcontext, label string) (int, error)
		// StreamClose(mode uint8, subject, subcontext, label string) error
		// StreamSend(mode uint8, subject, subcontext, label, data string, presences []Presence, reliable bool) error
		// StreamSendRaw(mode uint8, subject, subcontext, label string, msg *rtapi.Envelope, presences []Presence, reliable bool) error

		// SessionDisconnect(ctx context.Context, sessionID string) error

		// MatchCreate(ctx context.Context, module string, params map[string]interface{}) (string, error)
		// MatchGet(ctx context.Context, id string) (*api.Match, error)
		// MatchList(ctx context.Context, limit int, authoritative bool, label string, minSize, maxSize *int, query string) ([]*api.Match, error)

		// NotificationSend(ctx context.Context, userID, subject string, content map[string]interface{}, code int, sender string, persistent bool) error
		// NotificationsSend(ctx context.Context, notifications []*NotificationSend) error

		// WalletUpdate(ctx context.Context, userID string, changeset map[string]int64, metadata map[string]interface{}, updateLedger bool) (map[string]int64, map[string]int64, error)
		// WalletsUpdate(ctx context.Context, updates []*WalletUpdate, updateLedger bool) ([]*WalletUpdateResult, error)
		// WalletLedgerUpdate(ctx context.Context, itemID string, metadata map[string]interface{}) (WalletLedgerItem, error)
		// WalletLedgerList(ctx context.Context, userID string, limit int, cursor string) ([]WalletLedgerItem, string, error)

		// StorageList(ctx context.Context, userID, collection string, limit int, cursor string) ([]*api.StorageObject, string, error)
		// StorageRead(ctx context.Context, reads []*StorageRead) ([]*api.StorageObject, error)
		// StorageWrite(ctx context.Context, writes []*StorageWrite) ([]*api.StorageObjectAck, error)
		// StorageDelete(ctx context.Context, deletes []*StorageDelete) error

		// MultiUpdate(ctx context.Context, accountUpdates []*AccountUpdate, storageWrites []*StorageWrite, walletUpdates []*WalletUpdate, updateLedger bool) ([]*api.StorageObjectAck, []*WalletUpdateResult, error)

		// LeaderboardCreate(ctx context.Context, id string, authoritative bool, sortOrder, operator, resetSchedule string, metadata map[string]interface{}) error
		// LeaderboardDelete(ctx context.Context, id string) error
		// LeaderboardRecordsList(ctx context.Context, id string, ownerIDs []string, limit int, cursor string, expiry int64) ([]*api.LeaderboardRecord, []*api.LeaderboardRecord, string, string, error)
		// LeaderboardRecordWrite(ctx context.Context, id, ownerID, username string, score, subscore int64, metadata map[string]interface{}) (*api.LeaderboardRecord, error)
		// LeaderboardRecordDelete(ctx context.Context, id, ownerID string) error

		// TournamentCreate(ctx context.Context, id string, sortOrder, operator, resetSchedule string, metadata map[string]interface{}, title, description string, category, startTime, endTime, duration, maxSize, maxNumScore int, joinRequired bool) error
		// TournamentDelete(ctx context.Context, id string) error
		// TournamentAddAttempt(ctx context.Context, id, ownerID string, count int) error
		// TournamentJoin(ctx context.Context, id, ownerID, username string) error
		// TournamentsGetId(ctx context.Context, tournamentIDs []string) ([]*api.Tournament, error)
		// TournamentList(ctx context.Context, categoryStart, categoryEnd, startTime, endTime, limit int, cursor string) (*api.TournamentList, error)
		// TournamentRecordsList(ctx context.Context, tournamentId string, ownerIDs []string, limit int, cursor string, overrideExpiry int64) ([]*api.LeaderboardRecord, []*api.LeaderboardRecord, string, string, error)
		// TournamentRecordWrite(ctx context.Context, id, ownerID, username string, score, subscore int64, metadata map[string]interface{}) (*api.LeaderboardRecord, error)
		// TournamentRecordsHaystack(ctx context.Context, id, ownerID string, limit int, expiry int64) ([]*api.LeaderboardRecord, error)

		// GroupsGetId(ctx context.Context, groupIDs []string) ([]*api.Group, error)
		// GroupCreate(ctx context.Context, userID, name, creatorID, langTag, description, avatarUrl string, open bool, metadata map[string]interface{}, maxCount int) (*api.Group, error)
		// GroupUpdate(ctx context.Context, id, name, creatorID, langTag, description, avatarUrl string, open bool, metadata map[string]interface{}, maxCount int) error
		// GroupDelete(ctx context.Context, id string) error
		// GroupUserJoin(ctx context.Context, groupID, userID, username string) error
		// GroupUserLeave(ctx context.Context, groupID, userID, username string) error
		// GroupUsersAdd(ctx context.Context, groupID string, userIDs []string) error
		// GroupUsersKick(ctx context.Context, groupID string, userIDs []string) error
		// GroupUsersPromote(ctx context.Context, groupID string, userIDs []string) error
		// GroupUsersDemote(ctx context.Context, groupID string, userIDs []string) error
		// GroupUsersList(ctx context.Context, id string, limit int, state *int, cursor string) ([]*api.GroupUserList_GroupUser, string, error)
		// UserGroupsList(ctx context.Context, userID string, limit int, state *int, cursor string) ([]*api.UserGroupList_UserGroup, string, error)

		// FriendsList(ctx context.Context, userID string, limit int, state *int, cursor string) ([]*api.Friend, string, error)

		// Event(ctx context.Context, evt *api.Event) error
	} NkModule;

	//--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--//

	typedef struct
	{
		void *ptr;
	} NkInitializer;

#ifdef __cplusplus
}
#endif

#endif
