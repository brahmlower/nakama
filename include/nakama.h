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
// int nk_init_module(NkContext, NkLogger, NkDb, NkModule, NkInitializer);
//
// Match initializer:
// int nk_init_match(NkContext, NkLogger, NkDb, NkModule);

#ifndef NAKAMA_H
#define NAKAMA_H

#include <stdbool.h>
#include <stddef.h>
//#include "hashmap.h"

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

	typedef struct NkHashMap
	{
		const void *ptr;
	} NkHashMap;

	//--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--//

	typedef void (*NkContextValueFn)(const void *ptr, NkString key, NkString *outvalue);

	typedef struct NkContext
	{
		const void *ptr;
		NkContextValueFn value;
	} NkContext;

	//--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--//

	typedef struct NkHashMap (*NkLoggerFieldsFn)(const void *ptr);

	typedef void (*NkLoggerLevelFn)(const void *ptr, NkString s);

	typedef struct NkLogger (*NkLoggerWithFieldFn)(const void *ptr, NkString key, NkString value);

	typedef struct NkLogger (*NkLoggerWithFieldsFn)(const void *ptr, struct NkHashMap fields);

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

	//--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--//

	typedef struct NkDb
	{
		const void *ptr;
	} NkDb;

	//--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--=--//

	typedef int (*NkModuleAuthenticateEmailFn)(const void *ptr, const NkContext *ctx,
											   NkString email, NkString password, NkString username,
											   bool create, NkString *outuserid,
											   NkString *outusername, NkString *outerror,
											   bool *outcreated);

	typedef int (*NkModuleAuthenticateFacebookFn)(const void *ptr, const NkContext *ctx,
												  NkString token, bool importfriends,
												  NkString username, bool create,
												  NkString *outuserid, NkString *outusername,
												  NkString *outerror, bool *outcreated);

	typedef int (*NkModuleAuthenticateFn)(const void *ptr, const NkContext *ctx, NkString userid,
										  NkString username, bool create, NkString *outuserid,
										  NkString *outusername, NkString *outerror,
										  bool *outcreated);

	typedef int (*NkModuleAuthenticateGameCenterFn)(const void *ptr, const NkContext *ctx,
													NkString playerid, NkString bundleid,
													NkI64 timestamp, NkString salt,
													NkString signature, NkString publickeyurl,
													NkString username, bool create,
													NkString *outuserid, NkString *outusername,
													NkString *outerror, bool *outcreated);

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

	typedef int (*NkRpcFn)(NkContext ctx, NkLogger logger, NkDb db, NkModule nk, NkString payload,
						   NkString *outpayload, NkString *outerror);

	typedef int (*NkInitializerRpcFn)(const void *ptr, NkString id, const NkRpcFn fn,
									  NkString *outerror);

	typedef struct NkEnvelope
	{
		const void *ptr;
	} NkEnvelope;

	typedef int (*NkBeforeRtCallbackFn)(NkContext ctx, NkLogger logger, NkDb db, NkModule nk,
										NkEnvelope envelope, NkEnvelope *outenvelope,
										NkString *outerror);

	typedef int (*NkInitializerBeforeRtFn)(const void *ptr, NkString id,
										   const NkBeforeRtCallbackFn cb,
										   NkString *outerror);

	typedef int (*NkAfterRtCallbackFn)(NkContext ctx, NkLogger logger, NkDb db, NkModule nk,
									   NkEnvelope envelope, NkEnvelope *outenvelope,
									   NkString *outerror);

	typedef int (*NkInitializerAfterRtFn)(const void *ptr, NkString id,
										  const NkAfterRtCallbackFn cb, NkString *outerror);

	typedef struct NkMatchmakerEntry
	{
		const void *ptr;
	} NkMatchmakerEntry;

	typedef int (*NkMatchmakerMatchedCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												 NkModule nk, const NkMatchmakerEntry *entries,
												 int numentries, NkString *outmatchid,
												 NkString *outerror);

	typedef int (*NkInitializerMatchmakerMatchedFn)(const void *ptr,
													const NkMatchmakerMatchedCallbackFn cb,
													NkString *outerror);

	typedef int (*NkMatchCallbackFn)(NkContext ctx, NkLogger logger, NkDb db, NkModule nk,
									 void *outmatch, NkString *outerror); // TODO: outmatch type!

	typedef int (*NkInitializerMatchFn)(const void *ptr, NkString name, const NkMatchCallbackFn cb,
										NkString *outerror);

	typedef struct NkTournament
	{
		const void *ptr;
	} NkTournament;

	typedef int (*NkTournamentCallbackFn)(NkContext ctx, NkLogger logger, NkDb db, NkModule nk,
										  NkTournament tournament, NkI64 end, NkI64 reset,
										  NkString *outerror);

	typedef int (*NkInitializerTournamentFn)(const void *ptr, const NkTournamentCallbackFn cb,
											 NkString *outerror);

	typedef struct NkLeaderBoard
	{
		const void *ptr;
	} NkLeaderBoard;

	typedef int (*NkLeaderBoardCallbackFn)(NkContext ctx, NkLogger logger, NkDb db, NkModule nk,
										   NkLeaderBoard leaderboard, NkI64 reset,
										   NkString *outerror);

	typedef int (*NkInitializerLeaderBoardFn)(const void *ptr, const NkLeaderBoardCallbackFn cb,
											  NkString *outerror);

	typedef int (*NkCallbackFn)(NkContext ctx, NkLogger logger, NkDb db, NkModule nk,
								NkString *outerror);

	typedef int (*NkInitializerBeforeGetAccountFn)(const void *ptr, const NkCallbackFn cb,
												   NkString *outerror);

	typedef struct NkAccount
	{
		const void *ptr;
	} NkAccount;

	typedef int (*NkAfterGetAccountCallbackFn)(NkContext ctx, NkLogger logger, NkDb db, NkModule nk,
											   NkAccount *outaccount, NkString *outerror);

	typedef int (*NkInitializerAfterGetAccountFn)(const void *ptr,
												  const NkAfterGetAccountCallbackFn cb,
												  NkString *outerror);

	typedef struct NkUpdateAccountRequest
	{
		const void *ptr;
	} NkUpdateAccountRequest;

	typedef int (*NkBeforeUpdateAccountCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												   NkModule nk, NkUpdateAccountRequest req,
												   NkUpdateAccountRequest *outreq,
												   NkString *outerror);

	typedef int (*NkInitializerBeforeUpdateAccountFn)(const void *ptr,
													  const NkBeforeUpdateAccountCallbackFn cb,
													  NkString *outerror);

	typedef int (*NkAfterUpdateAccountCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												  NkModule nk, NkUpdateAccountRequest req,
												  NkString *outerror);

	typedef int (*NkInitializerAfterUpdateAccountFn)(const void *ptr,
													 const NkAfterUpdateAccountCallbackFn cb,
													 NkString *outerror);

	typedef struct NkSessionRefreshRequest
	{
		const void *ptr;
	} NkSessionRefreshRequest;

	typedef int (*NkBeforeSessionRefreshCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
													NkModule nk, NkSessionRefreshRequest req,
													NkSessionRefreshRequest *outreq,
													NkString *outerror);

	typedef int (*NkInitializerBeforeSessionRefreshFn)(const void *ptr,
													   const NkBeforeSessionRefreshCallbackFn cb,
													   NkString *outerror);

	typedef struct NkSession
	{
		const void *ptr;
	} NkSession;

	typedef int (*NkAfterSessionRefreshCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												   NkModule nk, NkSession session,
												   NkSessionRefreshRequest req, NkString *outerror);

	typedef int (*NkInitializerAfterSessionRefreshFn)(const void *ptr,
													  const NkAfterSessionRefreshCallbackFn cb,
													  NkString *outerror);

	typedef struct NkAuthenticateAppleRequest
	{
		const void *ptr;
	} NkAuthenticateAppleRequest;

	typedef int (*NkBeforeAuthenticateAppleCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
													   NkModule nk, NkAuthenticateAppleRequest req,
													   NkAuthenticateAppleRequest *outreq,
													   NkString *outerror);

	typedef int (*NkInitializerBeforeAuthenticateAppleFn)(const void *ptr,
														  const NkBeforeAuthenticateAppleCallbackFn cb,
														  NkString *outerror);

	typedef int (*NkAfterAuthenticateAppleCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
													  NkModule nk, NkSession session,
													  NkAuthenticateAppleRequest req,
													  NkString *outerror);

	typedef int (*NkInitializerAfterAuthenticateAppleFn)(const void *ptr,
														 const NkAfterAuthenticateAppleCallbackFn cb,
														 NkString *outerror);

	typedef struct NkAuthenticateCustomRequest
	{
		const void *ptr;
	} NkAuthenticateCustomRequest;

	typedef int (*NkBeforeAuthenticateCustomCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
														NkModule nk, NkAuthenticateCustomRequest req,
														NkAuthenticateCustomRequest *outreq,
														NkString *outerror);

	typedef int (*NkInitializerBeforeAuthenticateCustomFn)(const void *ptr,
														   const NkBeforeAuthenticateCustomCallbackFn cb,
														   NkString *outerror);

	typedef int (*NkAfterAuthenticateCustomCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
													   NkModule nk, NkSession session,
													   NkAuthenticateCustomRequest req,
													   NkString *outerror);

	typedef int (*NkInitializerAfterAuthenticateCustomFn)(const void *ptr,
														  const NkAfterAuthenticateCustomCallbackFn cb,
														  NkString *outerror);

	typedef struct NkAuthenticateDeviceRequest
	{
		const void *ptr;
	} NkAuthenticateDeviceRequest;

	typedef int (*NkBeforeAuthenticateDeviceCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
														NkModule nk, NkAuthenticateDeviceRequest req,
														NkAuthenticateDeviceRequest *outreq,
														NkString *outerror);

	typedef int (*NkInitializerBeforeAuthenticateDeviceFn)(const void *ptr,
														   const NkBeforeAuthenticateDeviceCallbackFn cb,
														   NkString *outerror);

	typedef int (*NkAfterAuthenticateDeviceCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
													   NkModule nk, NkSession session,
													   NkAuthenticateDeviceRequest req,
													   NkString *outerror);

	typedef int (*NkInitializerAfterAuthenticateDeviceFn)(const void *ptr,
														  const NkAfterAuthenticateDeviceCallbackFn cb,
														  NkString *outerror);

	typedef struct NkAuthenticateEmailRequest
	{
		const void *ptr;
	} NkAuthenticateEmailRequest;

	typedef int (*NkBeforeAuthenticateEmailCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
													   NkModule nk, NkAuthenticateEmailRequest req,
													   NkAuthenticateEmailRequest *outreq,
													   NkString *outerror);

	typedef int (*NkInitializerBeforeAuthenticateEmailFn)(const void *ptr,
														  const NkBeforeAuthenticateEmailCallbackFn cb,
														  NkString *outerror);

	typedef int (*NkAfterAuthenticateEmailCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
													  NkModule nk, NkSession session,
													  NkAuthenticateEmailRequest req,
													  NkString *outerror);

	typedef int (*NkInitializerAfterAuthenticateEmailFn)(const void *ptr,
														 const NkAfterAuthenticateEmailCallbackFn cb,
														 NkString *outerror);

	typedef struct NkAuthenticateFacebookRequest
	{
		const void *ptr;
	} NkAuthenticateFacebookRequest;

	typedef int (*NkBeforeAuthenticateFacebookCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
														  NkModule nk, NkAuthenticateFacebookRequest req,
														  NkAuthenticateFacebookRequest *outreq,
														  NkString *outerror);

	typedef int (*NkInitializerBeforeAuthenticateFacebookFn)(const void *ptr,
															 const NkBeforeAuthenticateFacebookCallbackFn cb,
															 NkString *outerror);

	typedef int (*NkAfterAuthenticateFacebookCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
														 NkModule nk, NkSession session,
														 NkAuthenticateFacebookRequest req,
														 NkString *outerror);

	typedef int (*NkInitializerAfterAuthenticateFacebookFn)(const void *ptr,
															const NkAfterAuthenticateFacebookCallbackFn cb,
															NkString *outerror);

	typedef struct NkAuthenticateFacebookInstantGameRequest
	{
		const void *ptr;
	} NkAuthenticateFacebookInstantGameRequest;

	typedef int (*NkBeforeAuthenticateFacebookInstantGameCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
																	 NkModule nk, NkAuthenticateFacebookInstantGameRequest req,
																	 NkAuthenticateFacebookInstantGameRequest *outreq,
																	 NkString *outerror);

	typedef int (*NkInitializerBeforeAuthenticateFacebookInstantGameFn)(const void *ptr,
																		const NkBeforeAuthenticateFacebookInstantGameCallbackFn cb,
																		NkString *outerror);

	typedef int (*NkAfterAuthenticateFacebookInstantGameCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
																	NkModule nk, NkSession session,
																	NkAuthenticateFacebookInstantGameRequest req,
																	NkString *outerror);

	typedef int (*NkInitializerAfterAuthenticateFacebookInstantGameFn)(const void *ptr,
																	   const NkAfterAuthenticateFacebookInstantGameCallbackFn cb,
																	   NkString *outerror);

	typedef struct NkAuthenticateGameCenterRequest
	{
		const void *ptr;
	} NkAuthenticateGameCenterRequest;

	typedef int (*NkBeforeAuthenticateGameCenterCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
															NkModule nk, NkAuthenticateGameCenterRequest req,
															NkAuthenticateGameCenterRequest *outreq,
															NkString *outerror);

	typedef int (*NkInitializerBeforeAuthenticateGameCenterFn)(const void *ptr,
															   const NkBeforeAuthenticateGameCenterCallbackFn cb,
															   NkString *outerror);

	typedef int (*NkAfterAuthenticateGameCenterCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
														   NkModule nk, NkSession session,
														   NkAuthenticateGameCenterRequest req,
														   NkString *outerror);

	typedef int (*NkInitializerAfterAuthenticateGameCenterFn)(const void *ptr,
															  const NkAfterAuthenticateGameCenterCallbackFn cb,
															  NkString *outerror);

	typedef struct NkAuthenticateGoogleRequest
	{
		const void *ptr;
	} NkAuthenticateGoogleRequest;

	typedef int (*NkBeforeAuthenticateGoogleCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
														NkModule nk, NkAuthenticateGoogleRequest req,
														NkAuthenticateGoogleRequest *outreq,
														NkString *outerror);

	typedef int (*NkInitializerBeforeAuthenticateGoogleFn)(const void *ptr,
														   const NkBeforeAuthenticateGoogleCallbackFn cb,
														   NkString *outerror);

	typedef int (*NkAfterAuthenticateGoogleCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
													   NkModule nk, NkSession session,
													   NkAuthenticateGoogleRequest req,
													   NkString *outerror);

	typedef int (*NkInitializerAfterAuthenticateGoogleFn)(const void *ptr,
														  const NkAfterAuthenticateGoogleCallbackFn cb,
														  NkString *outerror);

	typedef struct NkAuthenticateSteamRequest
	{
		const void *ptr;
	} NkAuthenticateSteamRequest;

	typedef int (*NkBeforeAuthenticateSteamCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
													   NkModule nk, NkAuthenticateSteamRequest req,
													   NkAuthenticateSteamRequest *outreq,
													   NkString *outerror);

	typedef int (*NkInitializerBeforeAuthenticateSteamFn)(const void *ptr,
														  const NkBeforeAuthenticateSteamCallbackFn cb,
														  NkString *outerror);

	typedef int (*NkAfterAuthenticateSteamCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
													  NkModule nk, NkSession session,
													  NkAuthenticateSteamRequest req,
													  NkString *outerror);

	typedef int (*NkInitializerAfterAuthenticateSteamFn)(const void *ptr,
														 const NkAfterAuthenticateSteamCallbackFn cb,
														 NkString *outerror);

	typedef struct NkListChannelMessagesRequest
	{
		const void *ptr;
	} NkListChannelMessagesRequest;

	typedef int (*NkBeforeListChannelMessagesCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
														 NkModule nk,
														 NkListChannelMessagesRequest req,
														 NkListChannelMessagesRequest *outreq,
														 NkString *outerror);

	typedef int (*NkInitializerBeforeListChannelMessagesFn)(const void *ptr,
															const NkBeforeListChannelMessagesCallbackFn cb,
															NkString *outerror);

	typedef struct NkChannelMessageList
	{
		const void *ptr;
	} NkChannelMessageList;

	typedef int (*NkAfterListChannelMessagesCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
														NkModule nk, NkChannelMessageList msgs,
														NkListChannelMessagesRequest req,
														NkString *outerror);

	typedef int (*NkInitializerAfterListChannelMessagesFn)(const void *ptr,
														   const NkAfterListChannelMessagesCallbackFn cb,
														   NkString *outerror);

	typedef struct NkListFriendsRequest
	{
		const void *ptr;
	} NkListFriendsRequest;

	typedef int (*NkBeforeListFriendsCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												 NkModule nk, NkListFriendsRequest req,
												 NkListFriendsRequest *outreq, NkString *outerror);

	typedef int (*NkInitializerBeforeListFriendsFn)(const void *ptr,
													const NkBeforeListFriendsCallbackFn cb,
													NkString *outerror);

	typedef struct NkFriendList
	{
		const void *ptr;
	} NkFriendList;

	typedef int (*NkAfterListFriendsCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												NkModule nk, NkFriendList friends,
												NkString *outerror);

	typedef int (*NkInitializerAfterListFriendsFn)(const void *ptr,
												   const NkAfterListFriendsCallbackFn cb,
												   NkString *outerror);

	typedef struct NkAddFriendsRequest
	{
		const void *ptr;
	} NkAddFriendsRequest;

	typedef int (*NkBeforeAddFriendsCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												NkModule nk, NkAddFriendsRequest req,
												NkAddFriendsRequest *outreq, NkString *outerror);

	typedef int (*NkInitializerBeforeAddFriendsFn)(const void *ptr,
												   const NkBeforeAddFriendsCallbackFn cb,
												   NkString *outerror);

	typedef int (*NkAfterAddFriendsCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
											   NkModule nk, NkAddFriendsRequest req,
											   NkString *outerror);

	typedef int (*NkInitializerAfterAddFriendsFn)(const void *ptr,
												  const NkAfterAddFriendsCallbackFn cb,
												  NkString *outerror);

	typedef struct NkDeleteFriendsRequest
	{
		const void *ptr;
	} NkDeleteFriendsRequest;

	typedef int (*NkBeforeDeleteFriendsCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												   NkModule nk, NkDeleteFriendsRequest req,
												   NkDeleteFriendsRequest *outreq,
												   NkString *outerror);

	typedef int (*NkInitializerBeforeDeleteFriendsFn)(const void *ptr,
													  const NkBeforeDeleteFriendsCallbackFn cb,
													  NkString *outerror);

	typedef int (*NkAfterDeleteFriendsCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												  NkModule nk, NkDeleteFriendsRequest req,
												  NkString *outerror);

	typedef int (*NkInitializerAfterDeleteFriendsFn)(const void *ptr,
													 const NkAfterDeleteFriendsCallbackFn cb,
													 NkString *outerror);

	typedef struct NkBlockFriendsRequest
	{
		const void *ptr;
	} NkBlockFriendsRequest;

	typedef int (*NkBeforeBlockFriendsCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												  NkModule nk, NkBlockFriendsRequest req,
												  NkBlockFriendsRequest *outreq,
												  NkString *outerror);

	typedef int (*NkInitializerBeforeBlockFriendsFn)(const void *ptr,
													 const NkBeforeBlockFriendsCallbackFn cb,
													 NkString *outerror);

	typedef int (*NkAfterBlockFriendsCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												 NkModule nk, NkBlockFriendsRequest req,
												 NkString *outerror);

	typedef int (*NkInitializerAfterBlockFriendsFn)(const void *ptr,
													const NkAfterBlockFriendsCallbackFn cb,
													NkString *outerror);

	typedef struct NkImportFacebookFriendsRequest
	{
		const void *ptr;
	} NkImportFacebookFriendsRequest;

	typedef int (*NkBeforeImportFacebookFriendsCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
														   NkModule nk,
														   NkImportFacebookFriendsRequest req,
														   NkImportFacebookFriendsRequest *outreq,
														   NkString *outerror);

	typedef int (*NkInitializerBeforeImportFacebookFriendsFn)(const void *ptr,
															  const NkBeforeImportFacebookFriendsCallbackFn cb,
															  NkString *outerror);

	typedef int (*NkAfterImportFacebookFriendsCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
														  NkModule nk,
														  NkImportFacebookFriendsRequest req,
														  NkString *outerror);

	typedef int (*NkInitializerAfterImportFacebookFriendsFn)(const void *ptr,
															 const NkAfterImportFacebookFriendsCallbackFn cb,
															 NkString *outerror);

	typedef struct NkCreateGroupRequest
	{
		const void *ptr;
	} NkCreateGroupRequest;

	typedef int (*NkBeforeCreateGroupCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												 NkModule nk, NkCreateGroupRequest req,
												 NkCreateGroupRequest *outreq, NkString *outerror);

	typedef int (*NkInitializerBeforeCreateGroupFn)(const void *ptr,
													const NkBeforeCreateGroupCallbackFn cb,
													NkString *outerror);

	typedef struct NkGroup
	{
		const void *ptr;
	} NkGroup;

	typedef int (*NkAfterCreateGroupCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												NkModule nk, NkGroup group,
												NkCreateGroupRequest req, NkString *outerror);

	typedef int (*NkInitializerAfterCreateGroupFn)(const void *ptr,
												   const NkAfterCreateGroupCallbackFn cb,
												   NkString *outerror);

	typedef struct NkUpdateGroupRequest
	{
		const void *ptr;
	} NkUpdateGroupRequest;

	typedef int (*NkBeforeUpdateGroupCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												 NkModule nk, NkUpdateGroupRequest req,
												 NkUpdateGroupRequest *outreq, NkString *outerror);

	typedef int (*NkInitializerBeforeUpdateGroupFn)(const void *ptr,
													const NkBeforeUpdateGroupCallbackFn cb,
													NkString *outerror);

	typedef int (*NkAfterUpdateGroupCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												NkModule nk, NkUpdateGroupRequest req,
												NkString *outerror);

	typedef int (*NkInitializerAfterUpdateGroupFn)(const void *ptr,
												   const NkAfterUpdateGroupCallbackFn cb,
												   NkString *outerror);

	typedef struct NkDeleteGroupRequest
	{
		const void *ptr;
	} NkDeleteGroupRequest;

	typedef int (*NkBeforeDeleteGroupCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												 NkModule nk, NkDeleteGroupRequest req,
												 NkDeleteGroupRequest *outreq, NkString *outerror);

	typedef int (*NkInitializerBeforeDeleteGroupFn)(const void *ptr,
													const NkBeforeDeleteGroupCallbackFn cb,
													NkString *outerror);

	typedef int (*NkAfterDeleteGroupCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												NkModule nk, NkDeleteGroupRequest req,
												NkString *outerror);

	typedef int (*NkInitializerAfterDeleteGroupFn)(const void *ptr,
												   const NkAfterDeleteGroupCallbackFn cb,
												   NkString *outerror);

	typedef struct NkJoinGroupRequest
	{
		const void *ptr;
	} NkJoinGroupRequest;

	typedef int (*NkBeforeJoinGroupCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
											   NkModule nk, NkJoinGroupRequest req,
											   NkJoinGroupRequest *outreq, NkString *outerror);

	typedef int (*NkInitializerBeforeJoinGroupFn)(const void *ptr,
												  const NkBeforeJoinGroupCallbackFn cb,
												  NkString *outerror);

	typedef int (*NkAfterJoinGroupCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
											  NkModule nk, NkJoinGroupRequest req,
											  NkString *outerror);

	typedef int (*NkInitializerAfterJoinGroupFn)(const void *ptr,
												 const NkAfterJoinGroupCallbackFn cb,
												 NkString *outerror);

	typedef struct NkLeaveGroupRequest
	{
		const void *ptr;
	} NkLeaveGroupRequest;

	typedef int (*NkBeforeLeaveGroupCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												NkModule nk, NkLeaveGroupRequest req,
												NkLeaveGroupRequest *outreq, NkString *outerror);

	typedef int (*NkInitializerBeforeLeaveGroupFn)(const void *ptr,
												   const NkBeforeLeaveGroupCallbackFn cb,
												   NkString *outerror);

	typedef int (*NkAfterLeaveGroupCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
											   NkModule nk, NkLeaveGroupRequest req,
											   NkString *outerror);

	typedef int (*NkInitializerAfterLeaveGroupFn)(const void *ptr,
												  const NkAfterLeaveGroupCallbackFn cb,
												  NkString *outerror);

	typedef struct NkAddGroupUsersRequest
	{
		const void *ptr;
	} NkAddGroupUsersRequest;

	typedef int (*NkBeforeAddGroupUsersCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												   NkModule nk, NkAddGroupUsersRequest req,
												   NkAddGroupUsersRequest *outreq,
												   NkString *outerror);

	typedef int (*NkInitializerBeforeAddGroupUsersFn)(const void *ptr,
													  const NkBeforeAddGroupUsersCallbackFn cb,
													  NkString *outerror);

	typedef int (*NkAfterAddGroupUsersCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												  NkModule nk, NkAddGroupUsersRequest req,
												  NkString *outerror);

	typedef int (*NkInitializerAfterAddGroupUsersFn)(const void *ptr,
													 const NkAfterAddGroupUsersCallbackFn cb,
													 NkString *outerror);

	typedef struct NkBanGroupUsersRequest
	{
		const void *ptr;
	} NkBanGroupUsersRequest;

	typedef int (*NkBeforeBanGroupUsersCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												   NkModule nk, NkBanGroupUsersRequest req,
												   NkBanGroupUsersRequest *outreq,
												   NkString *outerror);

	typedef int (*NkInitializerBeforeBanGroupUsersFn)(const void *ptr,
													  const NkBeforeBanGroupUsersCallbackFn cb,
													  NkString *outerror);

	typedef int (*NkAfterBanGroupUsersCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												  NkModule nk, NkBanGroupUsersRequest req,
												  NkString *outerror);

	typedef int (*NkInitializerAfterBanGroupUsersFn)(const void *ptr,
													 const NkAfterBanGroupUsersCallbackFn cb,
													 NkString *outerror);

	typedef struct NkKickGroupUsersRequest
	{
		const void *ptr;
	} NkKickGroupUsersRequest;

	typedef int (*NkBeforeKickGroupUsersCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
													NkModule nk, NkKickGroupUsersRequest req,
													NkKickGroupUsersRequest *outreq,
													NkString *outerror);

	typedef int (*NkInitializerBeforeKickGroupUsersFn)(const void *ptr,
													   const NkBeforeKickGroupUsersCallbackFn cb,
													   NkString *outerror);

	typedef int (*NkAfterKickGroupUsersCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
												   NkModule nk, NkKickGroupUsersRequest req,
												   NkString *outerror);

	typedef int (*NkInitializerAfterKickGroupUsersFn)(const void *ptr,
													  const NkAfterKickGroupUsersCallbackFn cb,
													  NkString *outerror);

	typedef struct NkPromoteGroupUsersRequest
	{
		const void *ptr;
	} NkPromoteGroupUsersRequest;

	typedef int (*NkBeforePromoteGroupUsersCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
													   NkModule nk, NkPromoteGroupUsersRequest req,
													   NkPromoteGroupUsersRequest *outreq,
													   NkString *outerror);

	typedef int (*NkInitializerBeforePromoteGroupUsersFn)(const void *ptr,
														  const NkBeforePromoteGroupUsersCallbackFn cb,
														  NkString *outerror);

	typedef int (*NkAfterPromoteGroupUsersCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
													  NkModule nk, NkPromoteGroupUsersRequest req,
													  NkString *outerror);

	typedef int (*NkInitializerAfterPromoteGroupUsersFn)(const void *ptr,
														 const NkAfterPromoteGroupUsersCallbackFn cb,
														 NkString *outerror);

	typedef struct NkDemoteGroupUsersRequest
	{
		const void *ptr;
	} NkDemoteGroupUsersRequest;

	typedef int (*NkBeforeDemoteGroupUsersCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
													  NkModule nk, NkDemoteGroupUsersRequest req,
													  NkDemoteGroupUsersRequest *outreq,
													  NkString *outerror);

	typedef int (*NkInitializerBeforeDemoteGroupUsersFn)(const void *ptr,
														 const NkBeforeDemoteGroupUsersCallbackFn cb,
														 NkString *outerror);

	typedef int (*NkAfterDemoteGroupUsersCallbackFn)(NkContext ctx, NkLogger logger, NkDb db,
													 NkModule nk, NkDemoteGroupUsersRequest req,
													 NkString *outerror);

	typedef int (*NkInitializerAfterDemoteGroupUsersFn)(const void *ptr,
														const NkAfterDemoteGroupUsersCallbackFn cb,
														NkString *outerror);

	// RegisterBeforeListGroupUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.ListGroupUsersRequest) (*api.ListGroupUsersRequest, error)) error
	// RegisterAfterListGroupUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, out *api.GroupUserList, in *api.ListGroupUsersRequest) error) error
	// RegisterBeforeListUserGroups(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.ListUserGroupsRequest) (*api.ListUserGroupsRequest, error)) error
	// RegisterAfterListUserGroups(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, out *api.UserGroupList, in *api.ListUserGroupsRequest) error) error
	// RegisterBeforeListGroups(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.ListGroupsRequest) (*api.ListGroupsRequest, error)) error
	// RegisterAfterListGroups(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, out *api.GroupList, in *api.ListGroupsRequest) error) error
	// RegisterBeforeDeleteLeaderboardRecord(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.DeleteLeaderboardRecordRequest) (*api.DeleteLeaderboardRecordRequest, error)) error
	// RegisterAfterDeleteLeaderboardRecord(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.DeleteLeaderboardRecordRequest) error) error
	// RegisterBeforeListLeaderboardRecords(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.ListLeaderboardRecordsRequest) (*api.ListLeaderboardRecordsRequest, error)) error
	// RegisterAfterListLeaderboardRecords(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, out *api.LeaderboardRecordList, in *api.ListLeaderboardRecordsRequest) error) error
	// RegisterBeforeWriteLeaderboardRecord(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.WriteLeaderboardRecordRequest) (*api.WriteLeaderboardRecordRequest, error)) error
	// RegisterAfterWriteLeaderboardRecord(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, out *api.LeaderboardRecord, in *api.WriteLeaderboardRecordRequest) error) error
	// RegisterBeforeListLeaderboardRecordsAroundOwner(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.ListLeaderboardRecordsAroundOwnerRequest) (*api.ListLeaderboardRecordsAroundOwnerRequest, error)) error
	// RegisterAfterListLeaderboardRecordsAroundOwner(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, out *api.LeaderboardRecordList, in *api.ListLeaderboardRecordsAroundOwnerRequest) error) error
	// RegisterBeforeLinkApple(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountApple) (*api.AccountApple, error)) error
	// RegisterAfterLinkApple(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountApple) error) error
	// RegisterBeforeLinkCustom(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountCustom) (*api.AccountCustom, error)) error
	// RegisterAfterLinkCustom(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountCustom) error) error
	// RegisterBeforeLinkDevice(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountDevice) (*api.AccountDevice, error)) error
	// RegisterAfterLinkDevice(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountDevice) error) error
	// RegisterBeforeLinkEmail(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountEmail) (*api.AccountEmail, error)) error
	// RegisterAfterLinkEmail(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountEmail) error) error
	// RegisterBeforeLinkFacebook(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.LinkFacebookRequest) (*api.LinkFacebookRequest, error)) error
	// RegisterAfterLinkFacebook(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.LinkFacebookRequest) error) error
	// RegisterBeforeLinkFacebookInstantGame(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountFacebookInstantGame) (*api.AccountFacebookInstantGame, error)) error
	// RegisterAfterLinkFacebookInstantGame(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountFacebookInstantGame) error) error
	// RegisterBeforeLinkGameCenter(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountGameCenter) (*api.AccountGameCenter, error)) error
	// RegisterAfterLinkGameCenter(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountGameCenter) error) error
	// RegisterBeforeLinkGoogle(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountGoogle) (*api.AccountGoogle, error)) error
	// RegisterAfterLinkGoogle(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountGoogle) error) error
	// RegisterBeforeLinkSteam(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountSteam) (*api.AccountSteam, error)) error
	// RegisterAfterLinkSteam(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountSteam) error) error
	// RegisterBeforeListMatches(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.ListMatchesRequest) (*api.ListMatchesRequest, error)) error
	// RegisterAfterListMatches(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, out *api.MatchList, in *api.ListMatchesRequest) error) error
	// RegisterBeforeListNotifications(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.ListNotificationsRequest) (*api.ListNotificationsRequest, error)) error
	// RegisterAfterListNotifications(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, out *api.NotificationList, in *api.ListNotificationsRequest) error) error
	// RegisterBeforeDeleteNotification(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.DeleteNotificationsRequest) (*api.DeleteNotificationsRequest, error)) error
	// RegisterAfterDeleteNotification(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.DeleteNotificationsRequest) error) error
	// RegisterBeforeListStorageObjects(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.ListStorageObjectsRequest) (*api.ListStorageObjectsRequest, error)) error
	// RegisterAfterListStorageObjects(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, out *api.StorageObjectList, in *api.ListStorageObjectsRequest) error) error
	// RegisterBeforeReadStorageObjects(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.ReadStorageObjectsRequest) (*api.ReadStorageObjectsRequest, error)) error
	// RegisterAfterReadStorageObjects(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, out *api.StorageObjects, in *api.ReadStorageObjectsRequest) error) error
	// RegisterBeforeWriteStorageObjects(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.WriteStorageObjectsRequest) (*api.WriteStorageObjectsRequest, error)) error
	// RegisterAfterWriteStorageObjects(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, out *api.StorageObjectAcks, in *api.WriteStorageObjectsRequest) error) error
	// RegisterBeforeDeleteStorageObjects(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.DeleteStorageObjectsRequest) (*api.DeleteStorageObjectsRequest, error)) error
	// RegisterAfterDeleteStorageObjects(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.DeleteStorageObjectsRequest) error) error
	// RegisterBeforeJoinTournament(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.JoinTournamentRequest) (*api.JoinTournamentRequest, error)) error
	// RegisterAfterJoinTournament(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.JoinTournamentRequest) error) error
	// RegisterBeforeListTournamentRecords(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.ListTournamentRecordsRequest) (*api.ListTournamentRecordsRequest, error)) error
	// RegisterAfterListTournamentRecords(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, out *api.TournamentRecordList, in *api.ListTournamentRecordsRequest) error) error
	// RegisterBeforeListTournaments(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.ListTournamentsRequest) (*api.ListTournamentsRequest, error)) error
	// RegisterAfterListTournaments(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, out *api.TournamentList, in *api.ListTournamentsRequest) error) error
	// RegisterBeforeWriteTournamentRecord(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.WriteTournamentRecordRequest) (*api.WriteTournamentRecordRequest, error)) error
	// RegisterAfterWriteTournamentRecord(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, out *api.LeaderboardRecord, in *api.WriteTournamentRecordRequest) error) error
	// RegisterBeforeListTournamentRecordsAroundOwner(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.ListTournamentRecordsAroundOwnerRequest) (*api.ListTournamentRecordsAroundOwnerRequest, error)) error
	// RegisterAfterListTournamentRecordsAroundOwner(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, out *api.TournamentRecordList, in *api.ListTournamentRecordsAroundOwnerRequest) error) error
	// RegisterBeforeUnlinkApple(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountApple) (*api.AccountApple, error)) error
	// RegisterAfterUnlinkApple(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountApple) error) error
	// RegisterBeforeUnlinkCustom(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountCustom) (*api.AccountCustom, error)) error
	// RegisterAfterUnlinkCustom(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountCustom) error) error
	// RegisterBeforeUnlinkDevice(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountDevice) (*api.AccountDevice, error)) error
	// RegisterAfterUnlinkDevice(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountDevice) error) error
	// RegisterBeforeUnlinkEmail(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountEmail) (*api.AccountEmail, error)) error
	// RegisterAfterUnlinkEmail(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountEmail) error) error
	// RegisterBeforeUnlinkFacebook(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountFacebook) (*api.AccountFacebook, error)) error
	// RegisterAfterUnlinkFacebook(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountFacebook) error) error
	// RegisterBeforeUnlinkFacebookInstantGame(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountFacebookInstantGame) (*api.AccountFacebookInstantGame, error)) error
	// RegisterAfterUnlinkFacebookInstantGame(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountFacebookInstantGame) error) error
	// RegisterBeforeUnlinkGameCenter(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountGameCenter) (*api.AccountGameCenter, error)) error
	// RegisterAfterUnlinkGameCenter(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountGameCenter) error) error
	// RegisterBeforeUnlinkGoogle(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountGoogle) (*api.AccountGoogle, error)) error
	// RegisterAfterUnlinkGoogle(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountGoogle) error) error
	// RegisterBeforeUnlinkSteam(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountSteam) (*api.AccountSteam, error)) error
	// RegisterAfterUnlinkSteam(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.AccountSteam) error) error
	// RegisterBeforeGetUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, in *api.GetUsersRequest) (*api.GetUsersRequest, error)) error
	// RegisterAfterGetUsers(fn func(ctx context.Context, logger Logger, db *sql.DB, nk NakamaModule, out *api.Users, in *api.GetUsersRequest) error) error
	// RegisterEvent(fn func(ctx context.Context, logger Logger, evt *api.Event)) error
	// RegisterEventSessionStart(fn func(ctx context.Context, logger Logger, evt *api.Event)) error
	// RegisterEventSessionEnd(fn func(ctx context.Context, logger Logger, evt *api.Event)) error

	typedef struct
	{
		const void *ptr;

		// registerrpc registers a function with the given ID. This ID can be used within client
		// code to send an RPC message to execute the function and return the result. Results are
		// always returned as a JSON string (or optionally empty string).
		//
		// If there is an issue with the RPC call, return an empty string and the associated error
		// which will be returned to the client.
		NkInitializerRpcFn registerrpc;

		// registerbeforert registers a function for a message. The registered function will be
		// called after the message has been processed in the pipeline.
		//
		// The custom code will be executed asynchronously after the response message has been sent
		// to a client
		//
		// Message names can be found here:
		// https://heroiclabs.com/docs/runtime-code-basics/#message-names
		NkInitializerBeforeRtFn registerbeforert;

		// registerafterrt registers a function for a message. Any function may be registered to
		// intercept a message received from a client and operate on it (or reject it) based on
		// custom logic.
		//
		// This is useful to enforce specific rules on top of the standard features in the server.
		//
		// You can return `NULL` instead of the `rtapi.Envelope` and this will disable disable that
		// particular server functionality.
		//
		// Message names can be found here:
		// https://heroiclabs.com/docs/runtime-code-basics/#message-names
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

		// RegisterBeforeAuthenticateCustom can be used to perform pre-authentication checks.
		//
		// You can use this to process the input (such as decoding custom tokens) and ensure
		// inter-compatibility between Nakama and your own custom system.
		NkInitializerBeforeAuthenticateCustomFn registerbeforeauthenticatecustom;

		// RegisterAfterAuthenticateCustom can be used to perform after successful authentication
		// checks.
		//
		// For instance, you can run special logic if the account was just created like adding them
		// to newcomers leaderboard.
		NkInitializerAfterAuthenticateCustomFn registerafterauthenticatecustom;

		NkInitializerBeforeAuthenticateDeviceFn registerbeforeauthenticatedevice;
		NkInitializerAfterAuthenticateDeviceFn registerafterauthenticatedevice;
		NkInitializerBeforeAuthenticateEmailFn registerbeforeauthenticateemail;
		NkInitializerAfterAuthenticateEmailFn registerafterauthenticateemail;
		NkInitializerBeforeAuthenticateFacebookFn registerbeforeauthenticatefacebook;
		NkInitializerAfterAuthenticateFacebookFn registerafterauthenticatefacebook;
		NkInitializerBeforeAuthenticateFacebookInstantGameFn registerbeforeauthenticatefacebookinstantgame;
		NkInitializerAfterAuthenticateFacebookInstantGameFn registerafterauthenticatefacebookinstantgame;
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
		NkInitializerBeforeListLeaderboardRecordsAroundOwnerFn registerbeforelistleaderboardrecordsaroundowner;
		NkInitializerAfterListLeaderboardRecordsAroundOwnerFn registerafterlistleaderboardrecordsaroundowner;
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
		NkInitializerBeforeDeleteNotificationFn registerbeforedeletenotification;
		NkInitializerAfterDeleteNotificationFn registerafterdeletenotification;
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
		NkInitializerBeforeListTournamentRecordsAroundOwnerFn registerbeforelisttournamentrecordsaroundowner;
		NkInitializerAfterListTournamentRecordsAroundOwnerFn registerafterlisttournamentrecordsaroundowner;
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
