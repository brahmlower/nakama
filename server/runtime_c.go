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

import (
	"context"
	"database/sql"
	"strings"

	"github.com/dop251/goja"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/jsonpb"
	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/rtapi"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/heroiclabs/nakama/v3/social"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

func CheckRuntimeProviderC(logger *zap.Logger, config Config) error {
	modCache, err := cacheJavascriptModules(logger, config.GetRuntime().Path, config.GetRuntime().JsEntrypoint)
	if err != nil {
		return err
	}
	rp := &RuntimeProviderJS{
		logger: logger,
		config: config,
	}
	_, _, err = evalRuntimeModules(rp, modCache, nil, nil, func(RuntimeExecutionMode, string) {}, true)
	if err != nil {
		logger.Error("Failed to load C module.", zap.Error(err))
	}
	return err
}

func NewRuntimeProviderC(logger, startupLogger *zap.Logger, db *sql.DB, jsonpbMarshaler *jsonpb.Marshaler, jsonpbUnmarshaler *jsonpb.Unmarshaler, config Config, socialClient *social.Client, leaderboardCache LeaderboardCache, leaderboardRankCache LeaderboardRankCache, leaderboardScheduler LeaderboardScheduler, sessionRegistry SessionRegistry, matchRegistry MatchRegistry, tracker Tracker, metrics *Metrics, streamManager StreamManager, router MessageRouter, eventFn RuntimeEventCustomFunction, path, entrypoint string, matchProvider *MatchProvider) ([]string, map[string]RuntimeRpcFunction, map[string]RuntimeBeforeRtFunction, map[string]RuntimeAfterRtFunction, *RuntimeBeforeReqFunctions, *RuntimeAfterReqFunctions, RuntimeMatchmakerMatchedFunction, RuntimeTournamentEndFunction, RuntimeTournamentResetFunction, RuntimeLeaderboardResetFunction, error) {
	startupLogger.Info("Initialising JavaScript runtime provider", zap.String("path", path), zap.String("entrypoint", entrypoint))

	modCache, err := cacheJavascriptModules(startupLogger, path, entrypoint)
	if err != nil {
		startupLogger.Fatal("Failed to load JavaScript files", zap.Error(err))
	}

	jsJsonpbMarshaler := &jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  jsonpbMarshaler.EnumsAsInts,
		EmitDefaults: jsonpbMarshaler.EmitDefaults,
		Indent:       jsonpbMarshaler.Indent,
	}

	runtimeProviderJS := &RuntimeProviderJS{
		config:               config,
		logger:               logger,
		db:                   db,
		eventFn:              eventFn,
		matchCreateFn:        matchProvider.CreateMatch,
		matchRegistry:        matchRegistry,
		jsonpbMarshaler:      jsJsonpbMarshaler,
		jsonpbUnmarshaler:    jsonpbUnmarshaler,
		socialClient:         socialClient,
		leaderboardCache:     leaderboardCache,
		leaderboardRankCache: leaderboardRankCache,
		sessionRegistry:      sessionRegistry,
		tracker:              tracker,
		streamManager:        streamManager,
		router:               router,
		metrics:              metrics,
		poolCh:               make(chan *RuntimeJS, config.GetRuntime().JsMaxCount),
		maxCount:             uint32(config.GetRuntime().JsMaxCount),
		currentCount:         atomic.NewUint32(uint32(config.GetRuntime().JsMinCount)),
	}

	rpcFunctions := make(map[string]RuntimeRpcFunction, 0)
	beforeRtFunctions := make(map[string]RuntimeBeforeRtFunction, 0)
	afterRtFunctions := make(map[string]RuntimeAfterRtFunction, 0)
	beforeReqFunctions := &RuntimeBeforeReqFunctions{}
	afterReqFunctions := &RuntimeAfterReqFunctions{}
	var matchmakerMatchedFunction RuntimeMatchmakerMatchedFunction
	var tournamentEndFunction RuntimeTournamentEndFunction
	var tournamentResetFunction RuntimeTournamentResetFunction
	var leaderboardResetFunction RuntimeLeaderboardResetFunction

	callbacks, matchCallbacks, err := evalRuntimeModules(runtimeProviderJS, modCache, matchProvider, leaderboardScheduler, func(mode RuntimeExecutionMode, id string) {
		switch mode {
		case RuntimeExecutionModeRPC:
			rpcFunctions[id] = func(ctx context.Context, queryParams map[string][]string, userID, username string, vars map[string]string, expiry int64, sessionID, clientIP, clientPort, payload string) (string, error, codes.Code) {
				return runtimeProviderJS.Rpc(ctx, id, queryParams, userID, username, vars, expiry, sessionID, clientIP, clientPort, payload)
			}
		case RuntimeExecutionModeBefore:
			if strings.HasPrefix(id, strings.ToLower(RTAPI_PREFIX)) {
				beforeRtFunctions[id] = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, sessionID, clientIP, clientPort string, envelope *rtapi.Envelope) (*rtapi.Envelope, error) {
					return runtimeProviderJS.BeforeRt(ctx, id, logger, userID, username, vars, expiry, sessionID, clientIP, clientPort, envelope)
				}
			} else if strings.HasPrefix(id, strings.ToLower(API_PREFIX)) {
				shortID := strings.TrimPrefix(id, strings.ToLower(API_PREFIX))
				switch shortID {
				case "getaccount":
					beforeReqFunctions.beforeGetAccountFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string) (error, codes.Code) {
						_, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil)
						if err != nil {
							return err, code
						}
						return nil, 0
					}
				case "updateaccount":
					beforeReqFunctions.beforeUpdateAccountFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.UpdateAccountRequest) (*api.UpdateAccountRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.UpdateAccountRequest), nil, 0
					}
				case "sessionrefresh":
					beforeReqFunctions.beforeSessionRefreshFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.SessionRefreshRequest) (*api.SessionRefreshRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.SessionRefreshRequest), nil, 0
					}
				case "authenticateapple":
					beforeReqFunctions.beforeAuthenticateAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateAppleRequest) (*api.AuthenticateAppleRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AuthenticateAppleRequest), nil, 0
					}
				case "authenticatecustom":
					beforeReqFunctions.beforeAuthenticateCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateCustomRequest) (*api.AuthenticateCustomRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AuthenticateCustomRequest), nil, 0
					}
				case "authenticatedevice":
					beforeReqFunctions.beforeAuthenticateDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateDeviceRequest) (*api.AuthenticateDeviceRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AuthenticateDeviceRequest), nil, 0
					}
				case "authenticateemail":
					beforeReqFunctions.beforeAuthenticateEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateEmailRequest) (*api.AuthenticateEmailRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AuthenticateEmailRequest), nil, 0
					}
				case "authenticatefacebook":
					beforeReqFunctions.beforeAuthenticateFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateFacebookRequest) (*api.AuthenticateFacebookRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AuthenticateFacebookRequest), nil, 0
					}
				case "authenticatefacebookinstantgame":
					beforeReqFunctions.beforeAuthenticateFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateFacebookInstantGameRequest) (*api.AuthenticateFacebookInstantGameRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AuthenticateFacebookInstantGameRequest), nil, 0
					}
				case "authenticategamecenter":
					beforeReqFunctions.beforeAuthenticateGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateGameCenterRequest) (*api.AuthenticateGameCenterRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AuthenticateGameCenterRequest), nil, 0
					}
				case "authenticategoogle":
					beforeReqFunctions.beforeAuthenticateGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateGoogleRequest) (*api.AuthenticateGoogleRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AuthenticateGoogleRequest), nil, 0
					}
				case "authenticatesteam":
					beforeReqFunctions.beforeAuthenticateSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateSteamRequest) (*api.AuthenticateSteamRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AuthenticateSteamRequest), nil, 0
					}
				case "listchannelmessages":
					beforeReqFunctions.beforeListChannelMessagesFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListChannelMessagesRequest) (*api.ListChannelMessagesRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.ListChannelMessagesRequest), nil, 0
					}
				case "listfriends":
					beforeReqFunctions.beforeListFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListFriendsRequest) (*api.ListFriendsRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.ListFriendsRequest), nil, 0
					}
				case "addfriends":
					beforeReqFunctions.beforeAddFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AddFriendsRequest) (*api.AddFriendsRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AddFriendsRequest), nil, 0
					}
				case "deletefriends":
					beforeReqFunctions.beforeDeleteFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteFriendsRequest) (*api.DeleteFriendsRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.DeleteFriendsRequest), nil, 0
					}
				case "blockfriends":
					beforeReqFunctions.beforeBlockFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.BlockFriendsRequest) (*api.BlockFriendsRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.BlockFriendsRequest), nil, 0
					}
				case "importfacebookfriends":
					beforeReqFunctions.beforeImportFacebookFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ImportFacebookFriendsRequest) (*api.ImportFacebookFriendsRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.ImportFacebookFriendsRequest), nil, 0
					}
				case "creategroup":
					beforeReqFunctions.beforeCreateGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.CreateGroupRequest) (*api.CreateGroupRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.CreateGroupRequest), nil, 0
					}
				case "updategroup":
					beforeReqFunctions.beforeUpdateGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.UpdateGroupRequest) (*api.UpdateGroupRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.UpdateGroupRequest), nil, 0
					}
				case "deletegroup":
					beforeReqFunctions.beforeDeleteGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteGroupRequest) (*api.DeleteGroupRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.DeleteGroupRequest), nil, 0
					}
				case "joingroup":
					beforeReqFunctions.beforeJoinGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.JoinGroupRequest) (*api.JoinGroupRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.JoinGroupRequest), nil, 0
					}
				case "leavegroup":
					beforeReqFunctions.beforeLeaveGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.LeaveGroupRequest) (*api.LeaveGroupRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.LeaveGroupRequest), nil, 0
					}
				case "addgroupusers":
					beforeReqFunctions.beforeAddGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AddGroupUsersRequest) (*api.AddGroupUsersRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AddGroupUsersRequest), nil, 0
					}
				case "bangroupusers":
					beforeReqFunctions.beforeBanGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.BanGroupUsersRequest) (*api.BanGroupUsersRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.BanGroupUsersRequest), nil, 0
					}
				case "kickgroupusers":
					beforeReqFunctions.beforeKickGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.KickGroupUsersRequest) (*api.KickGroupUsersRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.KickGroupUsersRequest), nil, 0
					}
				case "promotegroupusers":
					beforeReqFunctions.beforePromoteGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.PromoteGroupUsersRequest) (*api.PromoteGroupUsersRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.PromoteGroupUsersRequest), nil, 0
					}
				case "listgroupusers":
					beforeReqFunctions.beforeListGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListGroupUsersRequest) (*api.ListGroupUsersRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.ListGroupUsersRequest), nil, 0
					}
				case "listusergroups":
					beforeReqFunctions.beforeListUserGroupsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListUserGroupsRequest) (*api.ListUserGroupsRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.ListUserGroupsRequest), nil, 0
					}
				case "listgroups":
					beforeReqFunctions.beforeListGroupsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListGroupsRequest) (*api.ListGroupsRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.ListGroupsRequest), nil, 0
					}
				case "deleteleaderboardrecord":
					beforeReqFunctions.beforeDeleteLeaderboardRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteLeaderboardRecordRequest) (*api.DeleteLeaderboardRecordRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.DeleteLeaderboardRecordRequest), nil, 0
					}
				case "listleaderboardrecords":
					beforeReqFunctions.beforeListLeaderboardRecordsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListLeaderboardRecordsRequest) (*api.ListLeaderboardRecordsRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.ListLeaderboardRecordsRequest), nil, 0
					}
				case "writeleaderboardrecord":
					beforeReqFunctions.beforeWriteLeaderboardRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.WriteLeaderboardRecordRequest) (*api.WriteLeaderboardRecordRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.WriteLeaderboardRecordRequest), nil, 0
					}
				case "listleaderboardrecordsaroundowner":
					beforeReqFunctions.beforeListLeaderboardRecordsAroundOwnerFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListLeaderboardRecordsAroundOwnerRequest) (*api.ListLeaderboardRecordsAroundOwnerRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.ListLeaderboardRecordsAroundOwnerRequest), nil, 0
					}
				case "linkapple":
					beforeReqFunctions.beforeLinkAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountApple) (*api.AccountApple, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountApple), nil, 0
					}
				case "linkcustom":
					beforeReqFunctions.beforeLinkCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountCustom) (*api.AccountCustom, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountCustom), nil, 0
					}
				case "linkdevice":
					beforeReqFunctions.beforeLinkDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountDevice) (*api.AccountDevice, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountDevice), nil, 0
					}
				case "linkemail":
					beforeReqFunctions.beforeLinkEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountEmail) (*api.AccountEmail, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountEmail), nil, 0
					}
				case "linkfacebook":
					beforeReqFunctions.beforeLinkFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.LinkFacebookRequest) (*api.LinkFacebookRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.LinkFacebookRequest), nil, 0
					}
				case "linkfacebookinstantgame":
					beforeReqFunctions.beforeLinkFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebookInstantGame) (*api.AccountFacebookInstantGame, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountFacebookInstantGame), nil, 0
					}
				case "linkgamecenter":
					beforeReqFunctions.beforeLinkGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGameCenter) (*api.AccountGameCenter, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountGameCenter), nil, 0
					}
				case "linkgoogle":
					beforeReqFunctions.beforeLinkGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGoogle) (*api.AccountGoogle, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountGoogle), nil, 0
					}
				case "linksteam":
					beforeReqFunctions.beforeLinkSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountSteam) (*api.AccountSteam, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountSteam), nil, 0
					}
				case "listmatches":
					beforeReqFunctions.beforeListMatchesFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListMatchesRequest) (*api.ListMatchesRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.ListMatchesRequest), nil, 0
					}
				case "listnotifications":
					beforeReqFunctions.beforeListNotificationsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListNotificationsRequest) (*api.ListNotificationsRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.ListNotificationsRequest), nil, 0
					}
				case "deletenotification":
					beforeReqFunctions.beforeDeleteNotificationFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteNotificationsRequest) (*api.DeleteNotificationsRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.DeleteNotificationsRequest), nil, 0
					}
				case "liststorageobjects":
					beforeReqFunctions.beforeListStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListStorageObjectsRequest) (*api.ListStorageObjectsRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.ListStorageObjectsRequest), nil, 0
					}
				case "readstorageobjects":
					beforeReqFunctions.beforeReadStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ReadStorageObjectsRequest) (*api.ReadStorageObjectsRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.ReadStorageObjectsRequest), nil, 0
					}
				case "writestorageobjects":
					beforeReqFunctions.beforeWriteStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.WriteStorageObjectsRequest) (*api.WriteStorageObjectsRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.WriteStorageObjectsRequest), nil, 0
					}
				case "deletestorageobjects":
					beforeReqFunctions.beforeDeleteStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteStorageObjectsRequest) (*api.DeleteStorageObjectsRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.DeleteStorageObjectsRequest), nil, 0
					}
				case "jointournament":
					beforeReqFunctions.beforeJoinTournamentFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.JoinTournamentRequest) (*api.JoinTournamentRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.JoinTournamentRequest), nil, 0
					}
				case "listtournamentrecords":
					beforeReqFunctions.beforeListTournamentRecordsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListTournamentRecordsRequest) (*api.ListTournamentRecordsRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.ListTournamentRecordsRequest), nil, 0
					}
				case "listtournaments":
					beforeReqFunctions.beforeListTournamentsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListTournamentsRequest) (*api.ListTournamentsRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.ListTournamentsRequest), nil, 0
					}
				case "writetournamentrecord":
					beforeReqFunctions.beforeWriteTournamentRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.WriteTournamentRecordRequest) (*api.WriteTournamentRecordRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.WriteTournamentRecordRequest), nil, 0
					}
				case "listtournamentrecordsaroundowner":
					beforeReqFunctions.beforeListTournamentRecordsAroundOwnerFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListTournamentRecordsAroundOwnerRequest) (*api.ListTournamentRecordsAroundOwnerRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.ListTournamentRecordsAroundOwnerRequest), nil, 0
					}
				case "unlinkapple":
					beforeReqFunctions.beforeUnlinkAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountApple) (*api.AccountApple, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountApple), nil, 0
					}
				case "unlinkcustom":
					beforeReqFunctions.beforeUnlinkCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountCustom) (*api.AccountCustom, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountCustom), nil, 0
					}
				case "unlinkdevice":
					beforeReqFunctions.beforeUnlinkDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountDevice) (*api.AccountDevice, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountDevice), nil, 0
					}
				case "unlinkemail":
					beforeReqFunctions.beforeUnlinkEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountEmail) (*api.AccountEmail, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountEmail), nil, 0
					}
				case "unlinkfacebook":
					beforeReqFunctions.beforeUnlinkFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebook) (*api.AccountFacebook, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountFacebook), nil, 0
					}
				case "unlinkfacebookinstantgame":
					beforeReqFunctions.beforeUnlinkFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebookInstantGame) (*api.AccountFacebookInstantGame, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountFacebookInstantGame), nil, 0
					}
				case "unlinkgamecenter":
					beforeReqFunctions.beforeUnlinkGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGameCenter) (*api.AccountGameCenter, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountGameCenter), nil, 0
					}
				case "unlinkgoogle":
					beforeReqFunctions.beforeUnlinkGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGoogle) (*api.AccountGoogle, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountGoogle), nil, 0
					}
				case "unlinksteam":
					beforeReqFunctions.beforeUnlinkSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountSteam) (*api.AccountSteam, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.AccountSteam), nil, 0
					}
				case "getusers":
					beforeReqFunctions.beforeGetUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.GetUsersRequest) (*api.GetUsersRequest, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.GetUsersRequest), nil, 0
					}
				case "event":
					beforeReqFunctions.beforeEventFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.Event) (*api.Event, error, codes.Code) {
						result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
						if result == nil || err != nil {
							return nil, err, code
						}
						return result.(*api.Event), nil, 0
					}
				}
			}
		case RuntimeExecutionModeAfter:
			if strings.HasPrefix(id, strings.ToLower(RTAPI_PREFIX)) {
				afterRtFunctions[id] = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, sessionID, clientIP, clientPort string, envelope *rtapi.Envelope) error {
					return runtimeProviderJS.AfterRt(ctx, id, logger, userID, username, vars, expiry, sessionID, clientIP, clientPort, envelope)
				}
			} else if strings.HasPrefix(id, strings.ToLower(API_PREFIX)) {
				shortID := strings.TrimPrefix(id, strings.ToLower(API_PREFIX))
				switch shortID {
				case "getaccount":
					afterReqFunctions.afterGetAccountFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Account) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, nil)
					}
				case "updateaccount":
					afterReqFunctions.afterUpdateAccountFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.UpdateAccountRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "sessionrefresh":
					afterReqFunctions.afterSessionRefreshFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.SessionRefreshRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "authenticateapple":
					afterReqFunctions.afterAuthenticateAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateAppleRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "authenticatecustom":
					afterReqFunctions.afterAuthenticateCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateCustomRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "authenticatedevice":
					afterReqFunctions.afterAuthenticateDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateDeviceRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "authenticateemail":
					afterReqFunctions.afterAuthenticateEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateEmailRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "authenticatefacebook":
					afterReqFunctions.afterAuthenticateFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateFacebookRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "authenticatefacebookinstantgame":
					afterReqFunctions.afterAuthenticateFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateFacebookInstantGameRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "authenticategamecenter":
					afterReqFunctions.afterAuthenticateGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateGameCenterRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "authenticategoogle":
					afterReqFunctions.afterAuthenticateGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateGoogleRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "authenticatesteam":
					afterReqFunctions.afterAuthenticateSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateSteamRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "listchannelmessages":
					afterReqFunctions.afterListChannelMessagesFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.ChannelMessageList, in *api.ListChannelMessagesRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "listfriends":
					afterReqFunctions.afterListFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.FriendList) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, nil)
					}
				case "addfriends":
					afterReqFunctions.afterAddFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AddFriendsRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "deletefriends":
					afterReqFunctions.afterDeleteFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteFriendsRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "blockfriends":
					afterReqFunctions.afterBlockFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.BlockFriendsRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "importfacebookfriends":
					afterReqFunctions.afterImportFacebookFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ImportFacebookFriendsRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "creategroup":
					afterReqFunctions.afterCreateGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Group, in *api.CreateGroupRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "updategroup":
					afterReqFunctions.afterUpdateGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.UpdateGroupRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "deletegroup":
					afterReqFunctions.afterDeleteGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteGroupRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "joingroup":
					afterReqFunctions.afterJoinGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.JoinGroupRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "leavegroup":
					afterReqFunctions.afterLeaveGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.LeaveGroupRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "addgroupusers":
					afterReqFunctions.afterAddGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AddGroupUsersRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "bangroupusers":
					afterReqFunctions.afterBanGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.BanGroupUsersRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "kickgroupusers":
					afterReqFunctions.afterKickGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.KickGroupUsersRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "promotegroupusers":
					afterReqFunctions.afterPromoteGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.PromoteGroupUsersRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "listgroupusers":
					afterReqFunctions.afterListGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.GroupUserList, in *api.ListGroupUsersRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "listusergroups":
					afterReqFunctions.afterListUserGroupsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.UserGroupList, in *api.ListUserGroupsRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "listgroups":
					afterReqFunctions.afterListGroupsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.GroupList, in *api.ListGroupsRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "deleteleaderboardrecord":
					afterReqFunctions.afterDeleteLeaderboardRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteLeaderboardRecordRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "listleaderboardrecords":
					afterReqFunctions.afterListLeaderboardRecordsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.LeaderboardRecordList, in *api.ListLeaderboardRecordsRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "writeleaderboardrecord":
					afterReqFunctions.afterWriteLeaderboardRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.LeaderboardRecord, in *api.WriteLeaderboardRecordRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "listleaderboardrecordsaroundowner":
					afterReqFunctions.afterListLeaderboardRecordsAroundOwnerFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.LeaderboardRecordList, in *api.ListLeaderboardRecordsAroundOwnerRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "linkapple":
					afterReqFunctions.afterLinkAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountApple) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "linkcustom":
					afterReqFunctions.afterLinkCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountCustom) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "linkdevice":
					afterReqFunctions.afterLinkDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountDevice) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "linkemail":
					afterReqFunctions.afterLinkEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountEmail) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "linkfacebook":
					afterReqFunctions.afterLinkFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.LinkFacebookRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "linkfacebookinstantgame":
					afterReqFunctions.afterLinkFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebookInstantGame) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "linkgamecenter":
					afterReqFunctions.afterLinkGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGameCenter) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "linkgoogle":
					afterReqFunctions.afterLinkGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGoogle) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "linksteam":
					afterReqFunctions.afterLinkSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountSteam) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "listmatches":
					afterReqFunctions.afterListMatchesFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.MatchList, in *api.ListMatchesRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "listnotifications":
					afterReqFunctions.afterListNotificationsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.NotificationList, in *api.ListNotificationsRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "deletenotification":
					afterReqFunctions.afterDeleteNotificationFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteNotificationsRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "liststorageobjects":
					afterReqFunctions.afterListStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.StorageObjectList, in *api.ListStorageObjectsRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "readstorageobjects":
					afterReqFunctions.afterReadStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.StorageObjects, in *api.ReadStorageObjectsRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "writestorageobjects":
					afterReqFunctions.afterWriteStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.StorageObjectAcks, in *api.WriteStorageObjectsRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "deletestorageobjects":
					afterReqFunctions.afterDeleteStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteStorageObjectsRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "jointournament":
					afterReqFunctions.afterJoinTournamentFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.JoinTournamentRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "listtournamentrecords":
					afterReqFunctions.afterListTournamentRecordsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.TournamentRecordList, in *api.ListTournamentRecordsRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "listtournaments":
					afterReqFunctions.afterListTournamentsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.TournamentList, in *api.ListTournamentsRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "writetournamentrecord":
					afterReqFunctions.afterWriteTournamentRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.LeaderboardRecord, in *api.WriteTournamentRecordRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "listtournamentrecordsaroundowner":
					afterReqFunctions.afterListTournamentRecordsAroundOwnerFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.TournamentRecordList, in *api.ListTournamentRecordsAroundOwnerRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "unlinkapple":
					afterReqFunctions.afterUnlinkAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountApple) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "unlinkcustom":
					afterReqFunctions.afterUnlinkCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountCustom) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "unlinkdevice":
					afterReqFunctions.afterUnlinkDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountDevice) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "unlinkemail":
					afterReqFunctions.afterUnlinkEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountEmail) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "unlinkfacebook":
					afterReqFunctions.afterUnlinkFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebook) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "unlinkfacebookinstantgame":
					afterReqFunctions.afterUnlinkFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebookInstantGame) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "unlinkgamecenter":
					afterReqFunctions.afterUnlinkGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGameCenter) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "unlinkgoogle":
					afterReqFunctions.afterUnlinkGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGoogle) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "unlinksteam":
					afterReqFunctions.afterUnlinkSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountSteam) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				case "getusers":
					afterReqFunctions.afterGetUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Users, in *api.GetUsersRequest) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
					}
				case "event":
					afterReqFunctions.afterEventFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.Event) error {
						return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
					}
				}
			}
		case RuntimeExecutionModeMatchmaker:
			matchmakerMatchedFunction = func(ctx context.Context, entries []*MatchmakerEntry) (string, bool, error) {
				return runtimeProviderJS.MatchmakerMatched(ctx, entries)
			}
		case RuntimeExecutionModeTournamentEnd:
			tournamentEndFunction = func(ctx context.Context, tournament *api.Tournament, end, reset int64) error {
				return runtimeProviderJS.TournamentEnd(ctx, tournament, end, reset)
			}
		case RuntimeExecutionModeTournamentReset:
			tournamentResetFunction = func(ctx context.Context, tournament *api.Tournament, end, reset int64) error {
				return runtimeProviderJS.TournamentReset(ctx, tournament, end, reset)
			}
		case RuntimeExecutionModeLeaderboardReset:
			leaderboardResetFunction = func(ctx context.Context, leaderboard runtime.Leaderboard, reset int64) error {
				return runtimeProviderJS.LeaderboardReset(ctx, leaderboard, reset)
			}
		}
	}, false)
	if err != nil {
		logger.Error("Failed to eval JavaScript modules.", zap.Error(err))
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	matchProvider.RegisterCreateFn("javascript",
		func(ctx context.Context, logger *zap.Logger, id uuid.UUID, node string, stopped *atomic.Bool, name string) (RuntimeMatchCore, error) {
			mc := matchCallbacks.Get(name)
			if mc == nil {
				return nil, nil
			}

			return NewRuntimeJavascriptMatchCore(logger, name, db, jsonpbMarshaler, jsonpbUnmarshaler, config, socialClient, leaderboardCache, leaderboardRankCache, leaderboardScheduler, sessionRegistry, matchRegistry, tracker, streamManager, router, matchProvider.CreateMatch, eventFn, id, node, stopped, mc)
		})

	runtimeProviderJS.newFn = func() *RuntimeJS {
		runtime := goja.New()

		jsLogger := NewJsLogger(logger)
		jsLoggerValue := runtime.ToValue(jsLogger.Constructor(runtime))
		jsLoggerInst, err := runtime.New(jsLoggerValue)
		if err != nil {
			logger.Fatal("Failed to initialize JavaScript runtime", zap.Error(err))
		}

		nakamaModule := NewRuntimeJavascriptNakamaModule(logger, db, jsonpbMarshaler, jsonpbUnmarshaler, config, socialClient, leaderboardCache, leaderboardRankCache, leaderboardScheduler, sessionRegistry, matchRegistry, tracker, streamManager, router, eventFn, matchProvider.CreateMatch)
		nk := runtime.ToValue(nakamaModule.Constructor(runtime))
		nkInst, err := runtime.New(nk)
		if err != nil {
			logger.Fatal("Failed to initialize JavaScript runtime", zap.Error(err))
		}

		return &RuntimeJS{
			logger:       logger,
			jsLoggerInst: jsLoggerInst,
			nkInst:       nkInst,
			node:         config.GetName(),
			vm:           runtime,
			env:          runtime.ToValue(config.GetRuntime().Environment),
			callbacks:    callbacks,
		}
	}

	startupLogger.Info("JavaScript runtime modules loaded")

	// Warm up the pool.
	startupLogger.Info("Allocating minimum JavaScript runtime pool", zap.Int("count", config.GetRuntime().JsMinCount))
	if len(modCache.Names) > 0 {
		// Only if there are runtime modules to load.
		for i := 0; i < config.GetRuntime().JsMinCount; i++ {
			runtimeProviderJS.poolCh <- runtimeProviderJS.newFn()
		}
		runtimeProviderJS.metrics.GaugeJsRuntimes(float64(config.GetRuntime().JsMinCount))
	}
	startupLogger.Info("Allocated minimum JavaScript runtime pool")

	return modCache.Names, rpcFunctions, beforeRtFunctions, afterRtFunctions, beforeReqFunctions, afterReqFunctions, matchmakerMatchedFunction, tournamentEndFunction, tournamentResetFunction, leaderboardResetFunction, nil
}