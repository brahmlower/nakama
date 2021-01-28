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
	"errors"
	"path/filepath"
	"strings"
	"sync"


	"github.com/dop251/goja"
	"github.com/golang/protobuf/jsonpb"
	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/rtapi"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/heroiclabs/nakama/v3/social"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

type RuntimeC struct {
	logger       *zap.Logger
	node         string
	nkInst       goja.Value
	jsLoggerInst goja.Value
	env          goja.Value
	vm           *goja.Runtime
	callbacks    *RuntimeJavascriptCallbacks
}

type RuntimeCInitializer struct {
	logger runtime.Logger
	db     *sql.DB
	node   string
	env    map[string]string
	nk     runtime.NakamaModule

	rpc               map[string]RuntimeRpcFunction
	beforeRt          map[string]RuntimeBeforeRtFunction
	afterRt           map[string]RuntimeAfterRtFunction
	beforeReq         *RuntimeBeforeReqFunctions
	afterReq          *RuntimeAfterReqFunctions
	matchmakerMatched RuntimeMatchmakerMatchedFunction
	tournamentEnd     RuntimeTournamentEndFunction
	tournamentReset   RuntimeTournamentResetFunction
	leaderboardReset  RuntimeLeaderboardResetFunction

	eventFunctions        []RuntimeEventFunction
	sessionStartFunctions []RuntimeEventFunction
	sessionEndFunctions   []RuntimeEventFunction

	match     map[string]func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error)
	matchLock *sync.RWMutex
}

func (ri *RuntimeCInitializer) RegisterEvent(fn func(ctx context.Context, logger runtime.Logger, evt *api.Event)) error {
	ri.eventFunctions = append(ri.eventFunctions, fn)
	return nil
}

func (ri *RuntimeCInitializer) RegisterEventSessionStart(fn func(ctx context.Context, logger runtime.Logger, evt *api.Event)) error {
	ri.sessionStartFunctions = append(ri.sessionStartFunctions, fn)
	return nil
}

func (ri *RuntimeCInitializer) RegisterEventSessionEnd(fn func(ctx context.Context, logger runtime.Logger, evt *api.Event)) error {
	ri.sessionEndFunctions = append(ri.sessionEndFunctions, fn)
	return nil
}

func (ri *RuntimeCInitializer) RegisterRpc(id string, fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error)) error {
	id = strings.ToLower(id)
	ri.rpc[id] = func(ctx context.Context, queryParams map[string][]string, userID, username string, vars map[string]string, expiry int64, sessionID, clientIP, clientPort, payload string) (string, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeRPC, queryParams, expiry, userID, username, vars, sessionID, clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, payload)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeRt(id string, fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, envelope *rtapi.Envelope) (*rtapi.Envelope, error)) error {
	id = strings.ToLower(RTAPI_PREFIX + id)
	ri.beforeRt[id] = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, sessionID, clientIP, clientPort string, envelope *rtapi.Envelope) (*rtapi.Envelope, error) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, sessionID, clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, envelope)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterRt(id string, fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, envelope *rtapi.Envelope) error) error {
	id = strings.ToLower(RTAPI_PREFIX + id)
	ri.afterRt[id] = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, sessionID, clientIP, clientPort string, envelope *rtapi.Envelope) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, sessionID, clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, envelope)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeGetAccount(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) error) error {
	ri.beforeReq.beforeGetAccountFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string) (error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		fnErr := fn(ctx, ri.logger, ri.db, ri.nk)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return runtimeErr, codes.Internal
				}
				return runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return fnErr, codes.Internal
		}
		return nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterGetAccount(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Account) error) error {
	ri.afterReq.afterGetAccountFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Account) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeUpdateAccount(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.UpdateAccountRequest) (*api.UpdateAccountRequest, error)) error {
	ri.beforeReq.beforeUpdateAccountFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.UpdateAccountRequest) (*api.UpdateAccountRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterUpdateAccount(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.UpdateAccountRequest) error) error {
	ri.afterReq.afterUpdateAccountFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.UpdateAccountRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeSessionRefresh(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.SessionRefreshRequest) (*api.SessionRefreshRequest, error)) error {
	ri.beforeReq.beforeSessionRefreshFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.SessionRefreshRequest) (*api.SessionRefreshRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterSessionRefresh(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.SessionRefreshRequest) error) error {
	ri.afterReq.afterSessionRefreshFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.SessionRefreshRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeAuthenticateApple(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateAppleRequest) (*api.AuthenticateAppleRequest, error)) error {
	ri.beforeReq.beforeAuthenticateAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateAppleRequest) (*api.AuthenticateAppleRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterAuthenticateApple(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateAppleRequest) error) error {
	ri.afterReq.afterAuthenticateAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateAppleRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeAuthenticateCustom(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateCustomRequest) (*api.AuthenticateCustomRequest, error)) error {
	ri.beforeReq.beforeAuthenticateCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateCustomRequest) (*api.AuthenticateCustomRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterAuthenticateCustom(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateCustomRequest) error) error {
	ri.afterReq.afterAuthenticateCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateCustomRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeAuthenticateDevice(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateDeviceRequest) (*api.AuthenticateDeviceRequest, error)) error {
	ri.beforeReq.beforeAuthenticateDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateDeviceRequest) (*api.AuthenticateDeviceRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterAuthenticateDevice(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateDeviceRequest) error) error {
	ri.afterReq.afterAuthenticateDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateDeviceRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeAuthenticateEmail(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateEmailRequest) (*api.AuthenticateEmailRequest, error)) error {
	ri.beforeReq.beforeAuthenticateEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateEmailRequest) (*api.AuthenticateEmailRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterAuthenticateEmail(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateEmailRequest) error) error {
	ri.afterReq.afterAuthenticateEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateEmailRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeAuthenticateFacebook(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateFacebookRequest) (*api.AuthenticateFacebookRequest, error)) error {
	ri.beforeReq.beforeAuthenticateFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateFacebookRequest) (*api.AuthenticateFacebookRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterAuthenticateFacebook(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateFacebookRequest) error) error {
	ri.afterReq.afterAuthenticateFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateFacebookRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeAuthenticateFacebookInstantGame(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateFacebookInstantGameRequest) (*api.AuthenticateFacebookInstantGameRequest, error)) error {
	ri.beforeReq.beforeAuthenticateFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateFacebookInstantGameRequest) (*api.AuthenticateFacebookInstantGameRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterAuthenticateFacebookInstantGame(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateFacebookInstantGameRequest) error) error {
	ri.afterReq.afterAuthenticateFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateFacebookInstantGameRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeAuthenticateGameCenter(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateGameCenterRequest) (*api.AuthenticateGameCenterRequest, error)) error {
	ri.beforeReq.beforeAuthenticateGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateGameCenterRequest) (*api.AuthenticateGameCenterRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterAuthenticateGameCenter(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateGameCenterRequest) error) error {
	ri.afterReq.afterAuthenticateGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateGameCenterRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeAuthenticateGoogle(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateGoogleRequest) (*api.AuthenticateGoogleRequest, error)) error {
	ri.beforeReq.beforeAuthenticateGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateGoogleRequest) (*api.AuthenticateGoogleRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterAuthenticateGoogle(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateGoogleRequest) error) error {
	ri.afterReq.afterAuthenticateGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateGoogleRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeAuthenticateSteam(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateSteamRequest) (*api.AuthenticateSteamRequest, error)) error {
	ri.beforeReq.beforeAuthenticateSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateSteamRequest) (*api.AuthenticateSteamRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterAuthenticateSteam(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateSteamRequest) error) error {
	ri.afterReq.afterAuthenticateSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateSteamRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeListChannelMessages(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListChannelMessagesRequest) (*api.ListChannelMessagesRequest, error)) error {
	ri.beforeReq.beforeListChannelMessagesFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListChannelMessagesRequest) (*api.ListChannelMessagesRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterListChannelMessages(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.ChannelMessageList, in *api.ListChannelMessagesRequest) error) error {
	ri.afterReq.afterListChannelMessagesFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.ChannelMessageList, in *api.ListChannelMessagesRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeListFriends(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListFriendsRequest) (*api.ListFriendsRequest, error)) error {
	ri.beforeReq.beforeListFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListFriendsRequest) (*api.ListFriendsRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterListFriends(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.FriendList) error) error {
	ri.afterReq.afterListFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.FriendList) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeAddFriends(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AddFriendsRequest) (*api.AddFriendsRequest, error)) error {
	ri.beforeReq.beforeAddFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AddFriendsRequest) (*api.AddFriendsRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterAddFriends(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AddFriendsRequest) error) error {
	ri.afterReq.afterAddFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AddFriendsRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeDeleteFriends(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteFriendsRequest) (*api.DeleteFriendsRequest, error)) error {
	ri.beforeReq.beforeDeleteFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteFriendsRequest) (*api.DeleteFriendsRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterDeleteFriends(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteFriendsRequest) error) error {
	ri.afterReq.afterDeleteFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteFriendsRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeBlockFriends(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.BlockFriendsRequest) (*api.BlockFriendsRequest, error)) error {
	ri.beforeReq.beforeBlockFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.BlockFriendsRequest) (*api.BlockFriendsRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterBlockFriends(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.BlockFriendsRequest) error) error {
	ri.afterReq.afterBlockFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.BlockFriendsRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeImportFacebookFriends(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ImportFacebookFriendsRequest) (*api.ImportFacebookFriendsRequest, error)) error {
	ri.beforeReq.beforeImportFacebookFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ImportFacebookFriendsRequest) (*api.ImportFacebookFriendsRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterImportFacebookFriends(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ImportFacebookFriendsRequest) error) error {
	ri.afterReq.afterImportFacebookFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ImportFacebookFriendsRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeCreateGroup(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.CreateGroupRequest) (*api.CreateGroupRequest, error)) error {
	ri.beforeReq.beforeCreateGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.CreateGroupRequest) (*api.CreateGroupRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterCreateGroup(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Group, in *api.CreateGroupRequest) error) error {
	ri.afterReq.afterCreateGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Group, in *api.CreateGroupRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeUpdateGroup(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.UpdateGroupRequest) (*api.UpdateGroupRequest, error)) error {
	ri.beforeReq.beforeUpdateGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.UpdateGroupRequest) (*api.UpdateGroupRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterUpdateGroup(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.UpdateGroupRequest) error) error {
	ri.afterReq.afterUpdateGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.UpdateGroupRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeDeleteGroup(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteGroupRequest) (*api.DeleteGroupRequest, error)) error {
	ri.beforeReq.beforeDeleteGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteGroupRequest) (*api.DeleteGroupRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterDeleteGroup(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteGroupRequest) error) error {
	ri.afterReq.afterDeleteGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteGroupRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeJoinGroup(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.JoinGroupRequest) (*api.JoinGroupRequest, error)) error {
	ri.beforeReq.beforeJoinGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.JoinGroupRequest) (*api.JoinGroupRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterJoinGroup(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.JoinGroupRequest) error) error {
	ri.afterReq.afterJoinGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.JoinGroupRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeLeaveGroup(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.LeaveGroupRequest) (*api.LeaveGroupRequest, error)) error {
	ri.beforeReq.beforeLeaveGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.LeaveGroupRequest) (*api.LeaveGroupRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterLeaveGroup(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.LeaveGroupRequest) error) error {
	ri.afterReq.afterLeaveGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.LeaveGroupRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeAddGroupUsers(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AddGroupUsersRequest) (*api.AddGroupUsersRequest, error)) error {
	ri.beforeReq.beforeAddGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AddGroupUsersRequest) (*api.AddGroupUsersRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterAddGroupUsers(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AddGroupUsersRequest) error) error {
	ri.afterReq.afterAddGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AddGroupUsersRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeBanGroupUsers(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.BanGroupUsersRequest) (*api.BanGroupUsersRequest, error)) error {
	ri.beforeReq.beforeBanGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.BanGroupUsersRequest) (*api.BanGroupUsersRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterBanGroupUsers(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.BanGroupUsersRequest) error) error {
	ri.afterReq.afterBanGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.BanGroupUsersRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeKickGroupUsers(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.KickGroupUsersRequest) (*api.KickGroupUsersRequest, error)) error {
	ri.beforeReq.beforeKickGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.KickGroupUsersRequest) (*api.KickGroupUsersRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterKickGroupUsers(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.KickGroupUsersRequest) error) error {
	ri.afterReq.afterKickGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.KickGroupUsersRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforePromoteGroupUsers(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.PromoteGroupUsersRequest) (*api.PromoteGroupUsersRequest, error)) error {
	ri.beforeReq.beforePromoteGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.PromoteGroupUsersRequest) (*api.PromoteGroupUsersRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterPromoteGroupUsers(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.PromoteGroupUsersRequest) error) error {
	ri.afterReq.afterPromoteGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.PromoteGroupUsersRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeDemoteGroupUsers(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DemoteGroupUsersRequest) (*api.DemoteGroupUsersRequest, error)) error {
	ri.beforeReq.beforeDemoteGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DemoteGroupUsersRequest) (*api.DemoteGroupUsersRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterDemoteGroupUsers(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DemoteGroupUsersRequest) error) error {
	ri.afterReq.afterDemoteGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DemoteGroupUsersRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeListGroupUsers(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListGroupUsersRequest) (*api.ListGroupUsersRequest, error)) error {
	ri.beforeReq.beforeListGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListGroupUsersRequest) (*api.ListGroupUsersRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterListGroupUsers(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.GroupUserList, in *api.ListGroupUsersRequest) error) error {
	ri.afterReq.afterListGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.GroupUserList, in *api.ListGroupUsersRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeListUserGroups(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListUserGroupsRequest) (*api.ListUserGroupsRequest, error)) error {
	ri.beforeReq.beforeListUserGroupsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListUserGroupsRequest) (*api.ListUserGroupsRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterListUserGroups(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.UserGroupList, in *api.ListUserGroupsRequest) error) error {
	ri.afterReq.afterListUserGroupsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.UserGroupList, in *api.ListUserGroupsRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeListGroups(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListGroupsRequest) (*api.ListGroupsRequest, error)) error {
	ri.beforeReq.beforeListGroupsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListGroupsRequest) (*api.ListGroupsRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterListGroups(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.GroupList, in *api.ListGroupsRequest) error) error {
	ri.afterReq.afterListGroupsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.GroupList, in *api.ListGroupsRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeDeleteLeaderboardRecord(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteLeaderboardRecordRequest) (*api.DeleteLeaderboardRecordRequest, error)) error {
	ri.beforeReq.beforeDeleteLeaderboardRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteLeaderboardRecordRequest) (*api.DeleteLeaderboardRecordRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterDeleteLeaderboardRecord(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteLeaderboardRecordRequest) error) error {
	ri.afterReq.afterDeleteLeaderboardRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteLeaderboardRecordRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeListLeaderboardRecords(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListLeaderboardRecordsRequest) (*api.ListLeaderboardRecordsRequest, error)) error {
	ri.beforeReq.beforeListLeaderboardRecordsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListLeaderboardRecordsRequest) (*api.ListLeaderboardRecordsRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterListLeaderboardRecords(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.LeaderboardRecordList, in *api.ListLeaderboardRecordsRequest) error) error {
	ri.afterReq.afterListLeaderboardRecordsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.LeaderboardRecordList, in *api.ListLeaderboardRecordsRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeWriteLeaderboardRecord(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.WriteLeaderboardRecordRequest) (*api.WriteLeaderboardRecordRequest, error)) error {
	ri.beforeReq.beforeWriteLeaderboardRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.WriteLeaderboardRecordRequest) (*api.WriteLeaderboardRecordRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterWriteLeaderboardRecord(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.LeaderboardRecord, in *api.WriteLeaderboardRecordRequest) error) error {
	ri.afterReq.afterWriteLeaderboardRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.LeaderboardRecord, in *api.WriteLeaderboardRecordRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeListLeaderboardRecordsAroundOwner(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListLeaderboardRecordsAroundOwnerRequest) (*api.ListLeaderboardRecordsAroundOwnerRequest, error)) error {
	ri.beforeReq.beforeListLeaderboardRecordsAroundOwnerFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListLeaderboardRecordsAroundOwnerRequest) (*api.ListLeaderboardRecordsAroundOwnerRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterListLeaderboardRecordsAroundOwner(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.LeaderboardRecordList, in *api.ListLeaderboardRecordsAroundOwnerRequest) error) error {
	ri.afterReq.afterListLeaderboardRecordsAroundOwnerFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.LeaderboardRecordList, in *api.ListLeaderboardRecordsAroundOwnerRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeLinkApple(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountApple) (*api.AccountApple, error)) error {
	ri.beforeReq.beforeLinkAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountApple) (*api.AccountApple, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterLinkApple(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountApple) error) error {
	ri.afterReq.afterLinkAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountApple) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeLinkCustom(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountCustom) (*api.AccountCustom, error)) error {
	ri.beforeReq.beforeLinkCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountCustom) (*api.AccountCustom, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterLinkCustom(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountCustom) error) error {
	ri.afterReq.afterLinkCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountCustom) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeLinkDevice(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountDevice) (*api.AccountDevice, error)) error {
	ri.beforeReq.beforeLinkDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountDevice) (*api.AccountDevice, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterLinkDevice(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountDevice) error) error {
	ri.afterReq.afterLinkDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountDevice) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeLinkEmail(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountEmail) (*api.AccountEmail, error)) error {
	ri.beforeReq.beforeLinkEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountEmail) (*api.AccountEmail, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterLinkEmail(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountEmail) error) error {
	ri.afterReq.afterLinkEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountEmail) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeLinkFacebook(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.LinkFacebookRequest) (*api.LinkFacebookRequest, error)) error {
	ri.beforeReq.beforeLinkFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.LinkFacebookRequest) (*api.LinkFacebookRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterLinkFacebook(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.LinkFacebookRequest) error) error {
	ri.afterReq.afterLinkFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.LinkFacebookRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeLinkFacebookInstantGame(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountFacebookInstantGame) (*api.AccountFacebookInstantGame, error)) error {
	ri.beforeReq.beforeLinkFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebookInstantGame) (*api.AccountFacebookInstantGame, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterLinkFacebookInstantGame(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountFacebookInstantGame) error) error {
	ri.afterReq.afterLinkFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebookInstantGame) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeLinkGameCenter(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGameCenter) (*api.AccountGameCenter, error)) error {
	ri.beforeReq.beforeLinkGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGameCenter) (*api.AccountGameCenter, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterLinkGameCenter(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGameCenter) error) error {
	ri.afterReq.afterLinkGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGameCenter) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeLinkGoogle(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGoogle) (*api.AccountGoogle, error)) error {
	ri.beforeReq.beforeLinkGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGoogle) (*api.AccountGoogle, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterLinkGoogle(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGoogle) error) error {
	ri.afterReq.afterLinkGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGoogle) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeLinkSteam(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountSteam) (*api.AccountSteam, error)) error {
	ri.beforeReq.beforeLinkSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountSteam) (*api.AccountSteam, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterLinkSteam(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountSteam) error) error {
	ri.afterReq.afterLinkSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountSteam) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeListMatches(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListMatchesRequest) (*api.ListMatchesRequest, error)) error {
	ri.beforeReq.beforeListMatchesFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListMatchesRequest) (*api.ListMatchesRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterListMatches(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.MatchList, in *api.ListMatchesRequest) error) error {
	ri.afterReq.afterListMatchesFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.MatchList, in *api.ListMatchesRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeListNotifications(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListNotificationsRequest) (*api.ListNotificationsRequest, error)) error {
	ri.beforeReq.beforeListNotificationsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListNotificationsRequest) (*api.ListNotificationsRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterListNotifications(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.NotificationList, in *api.ListNotificationsRequest) error) error {
	ri.afterReq.afterListNotificationsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.NotificationList, in *api.ListNotificationsRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeDeleteNotification(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteNotificationsRequest) (*api.DeleteNotificationsRequest, error)) error {
	ri.beforeReq.beforeDeleteNotificationFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteNotificationsRequest) (*api.DeleteNotificationsRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterDeleteNotification(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteNotificationsRequest) error) error {
	ri.afterReq.afterDeleteNotificationFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteNotificationsRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeListStorageObjects(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListStorageObjectsRequest) (*api.ListStorageObjectsRequest, error)) error {
	ri.beforeReq.beforeListStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListStorageObjectsRequest) (*api.ListStorageObjectsRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterListStorageObjects(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.StorageObjectList, in *api.ListStorageObjectsRequest) error) error {
	ri.afterReq.afterListStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.StorageObjectList, in *api.ListStorageObjectsRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeReadStorageObjects(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ReadStorageObjectsRequest) (*api.ReadStorageObjectsRequest, error)) error {
	ri.beforeReq.beforeReadStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ReadStorageObjectsRequest) (*api.ReadStorageObjectsRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterReadStorageObjects(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.StorageObjects, in *api.ReadStorageObjectsRequest) error) error {
	ri.afterReq.afterReadStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.StorageObjects, in *api.ReadStorageObjectsRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeWriteStorageObjects(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.WriteStorageObjectsRequest) (*api.WriteStorageObjectsRequest, error)) error {
	ri.beforeReq.beforeWriteStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.WriteStorageObjectsRequest) (*api.WriteStorageObjectsRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterWriteStorageObjects(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.StorageObjectAcks, in *api.WriteStorageObjectsRequest) error) error {
	ri.afterReq.afterWriteStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.StorageObjectAcks, in *api.WriteStorageObjectsRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeDeleteStorageObjects(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteStorageObjectsRequest) (*api.DeleteStorageObjectsRequest, error)) error {
	ri.beforeReq.beforeDeleteStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteStorageObjectsRequest) (*api.DeleteStorageObjectsRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterDeleteStorageObjects(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteStorageObjectsRequest) error) error {
	ri.afterReq.afterDeleteStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteStorageObjectsRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeJoinTournament(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.JoinTournamentRequest) (*api.JoinTournamentRequest, error)) error {
	ri.beforeReq.beforeJoinTournamentFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.JoinTournamentRequest) (*api.JoinTournamentRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterJoinTournament(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.JoinTournamentRequest) error) error {
	ri.afterReq.afterJoinTournamentFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.JoinTournamentRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeListTournamentRecords(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListTournamentRecordsRequest) (*api.ListTournamentRecordsRequest, error)) error {
	ri.beforeReq.beforeListTournamentRecordsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListTournamentRecordsRequest) (*api.ListTournamentRecordsRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterListTournamentRecords(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.TournamentRecordList, in *api.ListTournamentRecordsRequest) error) error {
	ri.afterReq.afterListTournamentRecordsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.TournamentRecordList, in *api.ListTournamentRecordsRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeListTournaments(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListTournamentsRequest) (*api.ListTournamentsRequest, error)) error {
	ri.beforeReq.beforeListTournamentsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListTournamentsRequest) (*api.ListTournamentsRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterListTournaments(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.TournamentList, in *api.ListTournamentsRequest) error) error {
	ri.afterReq.afterListTournamentsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.TournamentList, in *api.ListTournamentsRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeWriteTournamentRecord(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.WriteTournamentRecordRequest) (*api.WriteTournamentRecordRequest, error)) error {
	ri.beforeReq.beforeWriteTournamentRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.WriteTournamentRecordRequest) (*api.WriteTournamentRecordRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterWriteTournamentRecord(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.LeaderboardRecord, in *api.WriteTournamentRecordRequest) error) error {
	ri.afterReq.afterWriteTournamentRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.LeaderboardRecord, in *api.WriteTournamentRecordRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeListTournamentRecordsAroundOwner(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListTournamentRecordsAroundOwnerRequest) (*api.ListTournamentRecordsAroundOwnerRequest, error)) error {
	ri.beforeReq.beforeListTournamentRecordsAroundOwnerFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListTournamentRecordsAroundOwnerRequest) (*api.ListTournamentRecordsAroundOwnerRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterListTournamentRecordsAroundOwner(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.TournamentRecordList, in *api.ListTournamentRecordsAroundOwnerRequest) error) error {
	ri.afterReq.afterListTournamentRecordsAroundOwnerFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.TournamentRecordList, in *api.ListTournamentRecordsAroundOwnerRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeUnlinkApple(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountApple) (*api.AccountApple, error)) error {
	ri.beforeReq.beforeUnlinkAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountApple) (*api.AccountApple, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterUnlinkApple(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountApple) error) error {
	ri.afterReq.afterUnlinkAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountApple) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeUnlinkCustom(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountCustom) (*api.AccountCustom, error)) error {
	ri.beforeReq.beforeUnlinkCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountCustom) (*api.AccountCustom, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterUnlinkCustom(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountCustom) error) error {
	ri.afterReq.afterUnlinkCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountCustom) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeUnlinkDevice(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountDevice) (*api.AccountDevice, error)) error {
	ri.beforeReq.beforeUnlinkDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountDevice) (*api.AccountDevice, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterUnlinkDevice(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountDevice) error) error {
	ri.afterReq.afterUnlinkDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountDevice) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeUnlinkEmail(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountEmail) (*api.AccountEmail, error)) error {
	ri.beforeReq.beforeUnlinkEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountEmail) (*api.AccountEmail, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterUnlinkEmail(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountEmail) error) error {
	ri.afterReq.afterUnlinkEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountEmail) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeUnlinkFacebook(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountFacebook) (*api.AccountFacebook, error)) error {
	ri.beforeReq.beforeUnlinkFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebook) (*api.AccountFacebook, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterUnlinkFacebook(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountFacebook) error) error {
	ri.afterReq.afterUnlinkFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebook) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeUnlinkFacebookInstantGame(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountFacebookInstantGame) (*api.AccountFacebookInstantGame, error)) error {
	ri.beforeReq.beforeUnlinkFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebookInstantGame) (*api.AccountFacebookInstantGame, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterUnlinkFacebookInstantGame(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountFacebookInstantGame) error) error {
	ri.afterReq.afterUnlinkFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebookInstantGame) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeUnlinkGameCenter(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGameCenter) (*api.AccountGameCenter, error)) error {
	ri.beforeReq.beforeUnlinkGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGameCenter) (*api.AccountGameCenter, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterUnlinkGameCenter(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGameCenter) error) error {
	ri.afterReq.afterUnlinkGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGameCenter) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeUnlinkGoogle(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGoogle) (*api.AccountGoogle, error)) error {
	ri.beforeReq.beforeUnlinkGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGoogle) (*api.AccountGoogle, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterUnlinkGoogle(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGoogle) error) error {
	ri.afterReq.afterUnlinkGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGoogle) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeUnlinkSteam(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountSteam) (*api.AccountSteam, error)) error {
	ri.beforeReq.beforeUnlinkSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountSteam) (*api.AccountSteam, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterUnlinkSteam(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountSteam) error) error {
	ri.afterReq.afterUnlinkSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountSteam) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeGetUsers(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.GetUsersRequest) (*api.GetUsersRequest, error)) error {
	ri.beforeReq.beforeGetUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.GetUsersRequest) (*api.GetUsersRequest, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterGetUsers(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Users, in *api.GetUsersRequest) error) error {
	ri.afterReq.afterGetUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Users, in *api.GetUsersRequest) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, out, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterBeforeEvent(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.Event) (*api.Event, error)) error {
	ri.beforeReq.beforeEventFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.Event) (*api.Event, error, codes.Code) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeBefore, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		result, fnErr := fn(ctx, ri.logger, ri.db, ri.nk, in)
		if fnErr != nil {
			if runtimeErr, ok := fnErr.(*runtime.Error); ok {
				if runtimeErr.Code <= 0 || runtimeErr.Code >= 17 {
					// If error is present but code is invalid then default to 13 (Internal) as the error code.
					return result, runtimeErr, codes.Internal
				}
				return result, runtimeErr, codes.Code(runtimeErr.Code)
			}
			// Not a runtime error that contains a code.
			return result, fnErr, codes.Internal
		}
		return result, nil, codes.OK
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterAfterEvent(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.Event) error) error {
	ri.afterReq.afterEventFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.Event) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeAfter, nil, expiry, userID, username, vars, "", clientIP, clientPort)
		return fn(ctx, ri.logger, ri.db, ri.nk, in)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterMatchmakerMatched(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, entries []runtime.MatchmakerEntry) (string, error)) error {
	ri.matchmakerMatched = func(ctx context.Context, entries []*MatchmakerEntry) (string, bool, error) {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeMatchmaker, nil, 0, "", "", nil, "", "", "")
		runtimeEntries := make([]runtime.MatchmakerEntry, len(entries))
		for i, entry := range entries {
			runtimeEntries[i] = runtime.MatchmakerEntry(entry)
		}
		matchID, err := fn(ctx, ri.logger, ri.db, ri.nk, runtimeEntries)
		if err != nil {
			return "", false, err
		}
		return matchID, matchID != "", nil
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterTournamentEnd(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, tournament *api.Tournament, end, reset int64) error) error {
	ri.tournamentEnd = func(ctx context.Context, tournament *api.Tournament, end, reset int64) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeTournamentEnd, nil, 0, "", "", nil, "", "", "")
		return fn(ctx, ri.logger, ri.db, ri.nk, tournament, end, reset)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterTournamentReset(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, tournament *api.Tournament, end, reset int64) error) error {
	ri.tournamentReset = func(ctx context.Context, tournament *api.Tournament, end, reset int64) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeTournamentReset, nil, 0, "", "", nil, "", "", "")
		return fn(ctx, ri.logger, ri.db, ri.nk, tournament, end, reset)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterLeaderboardReset(fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, leaderboard runtime.Leaderboard, reset int64) error) error {
	ri.leaderboardReset = func(ctx context.Context, leaderboard runtime.Leaderboard, reset int64) error {
		ctx = NewRuntimeGoContext(ctx, ri.node, ri.env, RuntimeExecutionModeLeaderboardReset, nil, 0, "", "", nil, "", "", "")
		return fn(ctx, ri.logger, ri.db, ri.nk, leaderboard, reset)
	}
	return nil
}

func (ri *RuntimeCInitializer) RegisterMatch(name string, fn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error)) error {
	ri.matchLock.Lock()
	ri.match[name] = fn
	ri.matchLock.Unlock()
	return nil
}

type RuntimeProviderC struct {
	logger               *zap.Logger
	db                   *sql.DB
	jsonpbMarshaler      *jsonpb.Marshaler
	jsonpbUnmarshaler    *jsonpb.Unmarshaler
	config               Config
	socialClient         *social.Client
	leaderboardCache     LeaderboardCache
	leaderboardRankCache LeaderboardRankCache
	sessionRegistry      SessionRegistry
	matchRegistry        MatchRegistry
	tracker              Tracker
	streamManager        StreamManager
	router               MessageRouter
	eventFn              RuntimeEventCustomFunction
	matchCreateFn        RuntimeMatchCreateFunction
	poolCh               chan *RuntimeJS
	maxCount             uint32
	currentCount         *atomic.Uint32
	newFn                func() *RuntimeJS
	metrics              *Metrics
}


type Symbol interface{}

func CheckRuntimeProviderC(logger *zap.Logger, config Config) error {
	// modCache, err := cacheJavascriptModules(logger, config.GetRuntime().Path, config.GetRuntime().JsEntrypoint)
	// if err != nil {
	// 	return err
	// }
	// rp := &RuntimeProviderC{
	// 	logger: logger,
	// 	config: config,
	// }
	// _, _, err = evalRuntimeModules(rp, modCache, nil, nil, func(RuntimeExecutionMode, string) {}, true)
	// if err != nil {
	// 	logger.Error("Failed to load C module.", zap.Error(err))
	// }
	// return err

	return errors.New("TODO")
}

func NewRuntimeProviderC(logger, startupLogger *zap.Logger, db *sql.DB, jsonpbMarshaler *jsonpb.Marshaler, jsonpbUnmarshaler *jsonpb.Unmarshaler, config Config, socialClient *social.Client, leaderboardCache LeaderboardCache, leaderboardRankCache LeaderboardRankCache, leaderboardScheduler LeaderboardScheduler, sessionRegistry SessionRegistry, matchRegistry MatchRegistry, tracker Tracker, metrics *Metrics, streamManager StreamManager, router MessageRouter, eventFn RuntimeEventCustomFunction, rootPath string, paths []string, matchProvider *MatchProvider) ([]string, map[string]RuntimeRpcFunction, map[string]RuntimeBeforeRtFunction, map[string]RuntimeAfterRtFunction, *RuntimeBeforeReqFunctions, *RuntimeAfterReqFunctions, RuntimeMatchmakerMatchedFunction, RuntimeTournamentEndFunction, RuntimeTournamentResetFunction, RuntimeLeaderboardResetFunction, error) {
	runtimeLogger := NewRuntimeGoLogger(logger)
	node := config.GetName()
	env := config.GetRuntime().Environment
	nk := NewRuntimeCNakamaModule(logger, db, jsonpbMarshaler, config, socialClient, leaderboardCache, leaderboardRankCache, leaderboardScheduler, sessionRegistry, matchRegistry, tracker, streamManager, router)
	match := make(map[string]func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error), 0)
	matchLock := &sync.RWMutex{}
	initializer := &RuntimeCInitializer{
		logger: runtimeLogger,
		db:     db,
		node:   node,
		env:    env,
		nk:     nk,

		rpc: make(map[string]RuntimeRpcFunction, 0),

		beforeRt: make(map[string]RuntimeBeforeRtFunction, 0),
		afterRt:  make(map[string]RuntimeAfterRtFunction, 0),

		beforeReq: &RuntimeBeforeReqFunctions{},
		afterReq:  &RuntimeAfterReqFunctions{},

		eventFunctions:        make([]RuntimeEventFunction, 0),
		sessionStartFunctions: make([]RuntimeEventFunction, 0),
		sessionEndFunctions:   make([]RuntimeEventFunction, 0),

		match:     match,
		matchLock: matchLock,
	}

	// The baseline context that will be passed to all InitModule calls.
	ctx := NewRuntimeCContext(context.Background(), node, env, RuntimeExecutionModeRunOnce, nil, 0, "", "", nil, "", "", "")

	startupLogger.Info("Initialising C runtime provider", zap.String("path", rootPath))

	modulePaths := make([]string, 0)
	for _, path := range paths {
		// Skip everything except shared object files.
		if strings.ToLower(filepath.Ext(path)) != ".so" {
			continue
		}

		// Open the library, and look up the required initialisation function.
		relPath, name, fn, err := openCModule(startupLogger, rootPath, path)
		if err != nil {
			// Errors are already logged in the function above.
			return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
		}

		// Run the initialisation.
		if err = fn(ctx, runtimeLogger, db, nk, initializer); err != nil {
			startupLogger.Fatal("Error returned by InitModule function in Go module", zap.String("name", name), zap.Error(err))
			return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, errors.New("error returned by InitModule function in C module")
		}
		modulePaths = append(modulePaths, relPath)
	}

	startupLogger.Info("C runtime modules loaded")

	// modCache, err := cacheJavascriptModules(startupLogger, path, entrypoint)
	// if err != nil {
	// 	startupLogger.Fatal("Failed to load JavaScript files", zap.Error(err))
	// }

	// jsJsonpbMarshaler := &jsonpb.Marshaler{
	// 	OrigName:     false,
	// 	EnumsAsInts:  jsonpbMarshaler.EnumsAsInts,
	// 	EmitDefaults: jsonpbMarshaler.EmitDefaults,
	// 	Indent:       jsonpbMarshaler.Indent,
	// }

	// runtimeProviderC := &RuntimeProviderC{
	// 	config:               config,
	// 	logger:               logger,
	// 	db:                   db,
	// 	eventFn:              eventFn,
	// 	matchCreateFn:        matchProvider.CreateMatch,
	// 	matchRegistry:        matchRegistry,
	// 	socialClient:         socialClient,
	// 	leaderboardCache:     leaderboardCache,
	// 	leaderboardRankCache: leaderboardRankCache,
	// 	sessionRegistry:      sessionRegistry,
	// 	tracker:              tracker,
	// 	streamManager:        streamManager,
	// 	router:               router,
	// 	metrics:              metrics,
	// 	poolCh:               make(chan *RuntimeJS, config.GetRuntime().JsMaxCount),
	// 	maxCount:             uint32(config.GetRuntime().JsMaxCount),
	// 	currentCount:         atomic.NewUint32(uint32(config.GetRuntime().JsMinCount)),
	// }

	rpcFunctions := make(map[string]RuntimeRpcFunction, 0)
	beforeRtFunctions := make(map[string]RuntimeBeforeRtFunction, 0)
	afterRtFunctions := make(map[string]RuntimeAfterRtFunction, 0)
	beforeReqFunctions := &RuntimeBeforeReqFunctions{}
	afterReqFunctions := &RuntimeAfterReqFunctions{}
	var matchmakerMatchedFunction RuntimeMatchmakerMatchedFunction
	var tournamentEndFunction RuntimeTournamentEndFunction
	var tournamentResetFunction RuntimeTournamentResetFunction
	var leaderboardResetFunction RuntimeLeaderboardResetFunction

	// callbacks, matchCallbacks, err := evalRuntimeModules(runtimeProviderJS, modCache, matchProvider, leaderboardScheduler, func(mode RuntimeExecutionMode, id string) {
	// 	switch mode {
	// 	case RuntimeExecutionModeRPC:
	// 		rpcFunctions[id] = func(ctx context.Context, queryParams map[string][]string, userID, username string, vars map[string]string, expiry int64, sessionID, clientIP, clientPort, payload string) (string, error, codes.Code) {
	// 			return runtimeProviderJS.Rpc(ctx, id, queryParams, userID, username, vars, expiry, sessionID, clientIP, clientPort, payload)
	// 		}
	// 	case RuntimeExecutionModeBefore:
	// 		if strings.HasPrefix(id, strings.ToLower(RTAPI_PREFIX)) {
	// 			beforeRtFunctions[id] = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, sessionID, clientIP, clientPort string, envelope *rtapi.Envelope) (*rtapi.Envelope, error) {
	// 				return runtimeProviderJS.BeforeRt(ctx, id, logger, userID, username, vars, expiry, sessionID, clientIP, clientPort, envelope)
	// 			}
	// 		} else if strings.HasPrefix(id, strings.ToLower(API_PREFIX)) {
	// 			shortID := strings.TrimPrefix(id, strings.ToLower(API_PREFIX))
	// 			switch shortID {
	// 			case "getaccount":
	// 				beforeReqFunctions.beforeGetAccountFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string) (error, codes.Code) {
	// 					_, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil)
	// 					if err != nil {
	// 						return err, code
	// 					}
	// 					return nil, 0
	// 				}
	// 			case "updateaccount":
	// 				beforeReqFunctions.beforeUpdateAccountFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.UpdateAccountRequest) (*api.UpdateAccountRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.UpdateAccountRequest), nil, 0
	// 				}
	// 			case "sessionrefresh":
	// 				beforeReqFunctions.beforeSessionRefreshFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.SessionRefreshRequest) (*api.SessionRefreshRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.SessionRefreshRequest), nil, 0
	// 				}
	// 			case "authenticateapple":
	// 				beforeReqFunctions.beforeAuthenticateAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateAppleRequest) (*api.AuthenticateAppleRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AuthenticateAppleRequest), nil, 0
	// 				}
	// 			case "authenticatecustom":
	// 				beforeReqFunctions.beforeAuthenticateCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateCustomRequest) (*api.AuthenticateCustomRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AuthenticateCustomRequest), nil, 0
	// 				}
	// 			case "authenticatedevice":
	// 				beforeReqFunctions.beforeAuthenticateDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateDeviceRequest) (*api.AuthenticateDeviceRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AuthenticateDeviceRequest), nil, 0
	// 				}
	// 			case "authenticateemail":
	// 				beforeReqFunctions.beforeAuthenticateEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateEmailRequest) (*api.AuthenticateEmailRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AuthenticateEmailRequest), nil, 0
	// 				}
	// 			case "authenticatefacebook":
	// 				beforeReqFunctions.beforeAuthenticateFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateFacebookRequest) (*api.AuthenticateFacebookRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AuthenticateFacebookRequest), nil, 0
	// 				}
	// 			case "authenticatefacebookinstantgame":
	// 				beforeReqFunctions.beforeAuthenticateFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateFacebookInstantGameRequest) (*api.AuthenticateFacebookInstantGameRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AuthenticateFacebookInstantGameRequest), nil, 0
	// 				}
	// 			case "authenticategamecenter":
	// 				beforeReqFunctions.beforeAuthenticateGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateGameCenterRequest) (*api.AuthenticateGameCenterRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AuthenticateGameCenterRequest), nil, 0
	// 				}
	// 			case "authenticategoogle":
	// 				beforeReqFunctions.beforeAuthenticateGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateGoogleRequest) (*api.AuthenticateGoogleRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AuthenticateGoogleRequest), nil, 0
	// 				}
	// 			case "authenticatesteam":
	// 				beforeReqFunctions.beforeAuthenticateSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AuthenticateSteamRequest) (*api.AuthenticateSteamRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AuthenticateSteamRequest), nil, 0
	// 				}
	// 			case "listchannelmessages":
	// 				beforeReqFunctions.beforeListChannelMessagesFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListChannelMessagesRequest) (*api.ListChannelMessagesRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.ListChannelMessagesRequest), nil, 0
	// 				}
	// 			case "listfriends":
	// 				beforeReqFunctions.beforeListFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListFriendsRequest) (*api.ListFriendsRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.ListFriendsRequest), nil, 0
	// 				}
	// 			case "addfriends":
	// 				beforeReqFunctions.beforeAddFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AddFriendsRequest) (*api.AddFriendsRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AddFriendsRequest), nil, 0
	// 				}
	// 			case "deletefriends":
	// 				beforeReqFunctions.beforeDeleteFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteFriendsRequest) (*api.DeleteFriendsRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.DeleteFriendsRequest), nil, 0
	// 				}
	// 			case "blockfriends":
	// 				beforeReqFunctions.beforeBlockFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.BlockFriendsRequest) (*api.BlockFriendsRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.BlockFriendsRequest), nil, 0
	// 				}
	// 			case "importfacebookfriends":
	// 				beforeReqFunctions.beforeImportFacebookFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ImportFacebookFriendsRequest) (*api.ImportFacebookFriendsRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.ImportFacebookFriendsRequest), nil, 0
	// 				}
	// 			case "creategroup":
	// 				beforeReqFunctions.beforeCreateGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.CreateGroupRequest) (*api.CreateGroupRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.CreateGroupRequest), nil, 0
	// 				}
	// 			case "updategroup":
	// 				beforeReqFunctions.beforeUpdateGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.UpdateGroupRequest) (*api.UpdateGroupRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.UpdateGroupRequest), nil, 0
	// 				}
	// 			case "deletegroup":
	// 				beforeReqFunctions.beforeDeleteGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteGroupRequest) (*api.DeleteGroupRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.DeleteGroupRequest), nil, 0
	// 				}
	// 			case "joingroup":
	// 				beforeReqFunctions.beforeJoinGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.JoinGroupRequest) (*api.JoinGroupRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.JoinGroupRequest), nil, 0
	// 				}
	// 			case "leavegroup":
	// 				beforeReqFunctions.beforeLeaveGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.LeaveGroupRequest) (*api.LeaveGroupRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.LeaveGroupRequest), nil, 0
	// 				}
	// 			case "addgroupusers":
	// 				beforeReqFunctions.beforeAddGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AddGroupUsersRequest) (*api.AddGroupUsersRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AddGroupUsersRequest), nil, 0
	// 				}
	// 			case "bangroupusers":
	// 				beforeReqFunctions.beforeBanGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.BanGroupUsersRequest) (*api.BanGroupUsersRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.BanGroupUsersRequest), nil, 0
	// 				}
	// 			case "kickgroupusers":
	// 				beforeReqFunctions.beforeKickGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.KickGroupUsersRequest) (*api.KickGroupUsersRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.KickGroupUsersRequest), nil, 0
	// 				}
	// 			case "promotegroupusers":
	// 				beforeReqFunctions.beforePromoteGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.PromoteGroupUsersRequest) (*api.PromoteGroupUsersRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.PromoteGroupUsersRequest), nil, 0
	// 				}
	// 			case "listgroupusers":
	// 				beforeReqFunctions.beforeListGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListGroupUsersRequest) (*api.ListGroupUsersRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.ListGroupUsersRequest), nil, 0
	// 				}
	// 			case "listusergroups":
	// 				beforeReqFunctions.beforeListUserGroupsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListUserGroupsRequest) (*api.ListUserGroupsRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.ListUserGroupsRequest), nil, 0
	// 				}
	// 			case "listgroups":
	// 				beforeReqFunctions.beforeListGroupsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListGroupsRequest) (*api.ListGroupsRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.ListGroupsRequest), nil, 0
	// 				}
	// 			case "deleteleaderboardrecord":
	// 				beforeReqFunctions.beforeDeleteLeaderboardRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteLeaderboardRecordRequest) (*api.DeleteLeaderboardRecordRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.DeleteLeaderboardRecordRequest), nil, 0
	// 				}
	// 			case "listleaderboardrecords":
	// 				beforeReqFunctions.beforeListLeaderboardRecordsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListLeaderboardRecordsRequest) (*api.ListLeaderboardRecordsRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.ListLeaderboardRecordsRequest), nil, 0
	// 				}
	// 			case "writeleaderboardrecord":
	// 				beforeReqFunctions.beforeWriteLeaderboardRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.WriteLeaderboardRecordRequest) (*api.WriteLeaderboardRecordRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.WriteLeaderboardRecordRequest), nil, 0
	// 				}
	// 			case "listleaderboardrecordsaroundowner":
	// 				beforeReqFunctions.beforeListLeaderboardRecordsAroundOwnerFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListLeaderboardRecordsAroundOwnerRequest) (*api.ListLeaderboardRecordsAroundOwnerRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.ListLeaderboardRecordsAroundOwnerRequest), nil, 0
	// 				}
	// 			case "linkapple":
	// 				beforeReqFunctions.beforeLinkAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountApple) (*api.AccountApple, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountApple), nil, 0
	// 				}
	// 			case "linkcustom":
	// 				beforeReqFunctions.beforeLinkCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountCustom) (*api.AccountCustom, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountCustom), nil, 0
	// 				}
	// 			case "linkdevice":
	// 				beforeReqFunctions.beforeLinkDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountDevice) (*api.AccountDevice, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountDevice), nil, 0
	// 				}
	// 			case "linkemail":
	// 				beforeReqFunctions.beforeLinkEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountEmail) (*api.AccountEmail, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountEmail), nil, 0
	// 				}
	// 			case "linkfacebook":
	// 				beforeReqFunctions.beforeLinkFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.LinkFacebookRequest) (*api.LinkFacebookRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.LinkFacebookRequest), nil, 0
	// 				}
	// 			case "linkfacebookinstantgame":
	// 				beforeReqFunctions.beforeLinkFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebookInstantGame) (*api.AccountFacebookInstantGame, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountFacebookInstantGame), nil, 0
	// 				}
	// 			case "linkgamecenter":
	// 				beforeReqFunctions.beforeLinkGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGameCenter) (*api.AccountGameCenter, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountGameCenter), nil, 0
	// 				}
	// 			case "linkgoogle":
	// 				beforeReqFunctions.beforeLinkGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGoogle) (*api.AccountGoogle, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountGoogle), nil, 0
	// 				}
	// 			case "linksteam":
	// 				beforeReqFunctions.beforeLinkSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountSteam) (*api.AccountSteam, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountSteam), nil, 0
	// 				}
	// 			case "listmatches":
	// 				beforeReqFunctions.beforeListMatchesFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListMatchesRequest) (*api.ListMatchesRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.ListMatchesRequest), nil, 0
	// 				}
	// 			case "listnotifications":
	// 				beforeReqFunctions.beforeListNotificationsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListNotificationsRequest) (*api.ListNotificationsRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.ListNotificationsRequest), nil, 0
	// 				}
	// 			case "deletenotification":
	// 				beforeReqFunctions.beforeDeleteNotificationFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteNotificationsRequest) (*api.DeleteNotificationsRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.DeleteNotificationsRequest), nil, 0
	// 				}
	// 			case "liststorageobjects":
	// 				beforeReqFunctions.beforeListStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListStorageObjectsRequest) (*api.ListStorageObjectsRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.ListStorageObjectsRequest), nil, 0
	// 				}
	// 			case "readstorageobjects":
	// 				beforeReqFunctions.beforeReadStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ReadStorageObjectsRequest) (*api.ReadStorageObjectsRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.ReadStorageObjectsRequest), nil, 0
	// 				}
	// 			case "writestorageobjects":
	// 				beforeReqFunctions.beforeWriteStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.WriteStorageObjectsRequest) (*api.WriteStorageObjectsRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.WriteStorageObjectsRequest), nil, 0
	// 				}
	// 			case "deletestorageobjects":
	// 				beforeReqFunctions.beforeDeleteStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteStorageObjectsRequest) (*api.DeleteStorageObjectsRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.DeleteStorageObjectsRequest), nil, 0
	// 				}
	// 			case "jointournament":
	// 				beforeReqFunctions.beforeJoinTournamentFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.JoinTournamentRequest) (*api.JoinTournamentRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.JoinTournamentRequest), nil, 0
	// 				}
	// 			case "listtournamentrecords":
	// 				beforeReqFunctions.beforeListTournamentRecordsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListTournamentRecordsRequest) (*api.ListTournamentRecordsRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.ListTournamentRecordsRequest), nil, 0
	// 				}
	// 			case "listtournaments":
	// 				beforeReqFunctions.beforeListTournamentsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListTournamentsRequest) (*api.ListTournamentsRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.ListTournamentsRequest), nil, 0
	// 				}
	// 			case "writetournamentrecord":
	// 				beforeReqFunctions.beforeWriteTournamentRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.WriteTournamentRecordRequest) (*api.WriteTournamentRecordRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.WriteTournamentRecordRequest), nil, 0
	// 				}
	// 			case "listtournamentrecordsaroundowner":
	// 				beforeReqFunctions.beforeListTournamentRecordsAroundOwnerFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ListTournamentRecordsAroundOwnerRequest) (*api.ListTournamentRecordsAroundOwnerRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.ListTournamentRecordsAroundOwnerRequest), nil, 0
	// 				}
	// 			case "unlinkapple":
	// 				beforeReqFunctions.beforeUnlinkAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountApple) (*api.AccountApple, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountApple), nil, 0
	// 				}
	// 			case "unlinkcustom":
	// 				beforeReqFunctions.beforeUnlinkCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountCustom) (*api.AccountCustom, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountCustom), nil, 0
	// 				}
	// 			case "unlinkdevice":
	// 				beforeReqFunctions.beforeUnlinkDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountDevice) (*api.AccountDevice, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountDevice), nil, 0
	// 				}
	// 			case "unlinkemail":
	// 				beforeReqFunctions.beforeUnlinkEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountEmail) (*api.AccountEmail, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountEmail), nil, 0
	// 				}
	// 			case "unlinkfacebook":
	// 				beforeReqFunctions.beforeUnlinkFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebook) (*api.AccountFacebook, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountFacebook), nil, 0
	// 				}
	// 			case "unlinkfacebookinstantgame":
	// 				beforeReqFunctions.beforeUnlinkFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebookInstantGame) (*api.AccountFacebookInstantGame, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountFacebookInstantGame), nil, 0
	// 				}
	// 			case "unlinkgamecenter":
	// 				beforeReqFunctions.beforeUnlinkGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGameCenter) (*api.AccountGameCenter, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountGameCenter), nil, 0
	// 				}
	// 			case "unlinkgoogle":
	// 				beforeReqFunctions.beforeUnlinkGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGoogle) (*api.AccountGoogle, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountGoogle), nil, 0
	// 				}
	// 			case "unlinksteam":
	// 				beforeReqFunctions.beforeUnlinkSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountSteam) (*api.AccountSteam, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.AccountSteam), nil, 0
	// 				}
	// 			case "getusers":
	// 				beforeReqFunctions.beforeGetUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.GetUsersRequest) (*api.GetUsersRequest, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.GetUsersRequest), nil, 0
	// 				}
	// 			case "event":
	// 				beforeReqFunctions.beforeEventFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.Event) (*api.Event, error, codes.Code) {
	// 					result, err, code := runtimeProviderJS.BeforeReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, in)
	// 					if result == nil || err != nil {
	// 						return nil, err, code
	// 					}
	// 					return result.(*api.Event), nil, 0
	// 				}
	// 			}
	// 		}
	// 	case RuntimeExecutionModeAfter:
	// 		if strings.HasPrefix(id, strings.ToLower(RTAPI_PREFIX)) {
	// 			afterRtFunctions[id] = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, sessionID, clientIP, clientPort string, envelope *rtapi.Envelope) error {
	// 				return runtimeProviderJS.AfterRt(ctx, id, logger, userID, username, vars, expiry, sessionID, clientIP, clientPort, envelope)
	// 			}
	// 		} else if strings.HasPrefix(id, strings.ToLower(API_PREFIX)) {
	// 			shortID := strings.TrimPrefix(id, strings.ToLower(API_PREFIX))
	// 			switch shortID {
	// 			case "getaccount":
	// 				afterReqFunctions.afterGetAccountFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Account) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, nil)
	// 				}
	// 			case "updateaccount":
	// 				afterReqFunctions.afterUpdateAccountFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.UpdateAccountRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "sessionrefresh":
	// 				afterReqFunctions.afterSessionRefreshFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.SessionRefreshRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "authenticateapple":
	// 				afterReqFunctions.afterAuthenticateAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateAppleRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "authenticatecustom":
	// 				afterReqFunctions.afterAuthenticateCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateCustomRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "authenticatedevice":
	// 				afterReqFunctions.afterAuthenticateDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateDeviceRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "authenticateemail":
	// 				afterReqFunctions.afterAuthenticateEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateEmailRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "authenticatefacebook":
	// 				afterReqFunctions.afterAuthenticateFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateFacebookRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "authenticatefacebookinstantgame":
	// 				afterReqFunctions.afterAuthenticateFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateFacebookInstantGameRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "authenticategamecenter":
	// 				afterReqFunctions.afterAuthenticateGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateGameCenterRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "authenticategoogle":
	// 				afterReqFunctions.afterAuthenticateGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateGoogleRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "authenticatesteam":
	// 				afterReqFunctions.afterAuthenticateSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Session, in *api.AuthenticateSteamRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "listchannelmessages":
	// 				afterReqFunctions.afterListChannelMessagesFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.ChannelMessageList, in *api.ListChannelMessagesRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "listfriends":
	// 				afterReqFunctions.afterListFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.FriendList) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, nil)
	// 				}
	// 			case "addfriends":
	// 				afterReqFunctions.afterAddFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AddFriendsRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "deletefriends":
	// 				afterReqFunctions.afterDeleteFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteFriendsRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "blockfriends":
	// 				afterReqFunctions.afterBlockFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.BlockFriendsRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "importfacebookfriends":
	// 				afterReqFunctions.afterImportFacebookFriendsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.ImportFacebookFriendsRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "creategroup":
	// 				afterReqFunctions.afterCreateGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Group, in *api.CreateGroupRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "updategroup":
	// 				afterReqFunctions.afterUpdateGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.UpdateGroupRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "deletegroup":
	// 				afterReqFunctions.afterDeleteGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteGroupRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "joingroup":
	// 				afterReqFunctions.afterJoinGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.JoinGroupRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "leavegroup":
	// 				afterReqFunctions.afterLeaveGroupFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.LeaveGroupRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "addgroupusers":
	// 				afterReqFunctions.afterAddGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AddGroupUsersRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "bangroupusers":
	// 				afterReqFunctions.afterBanGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.BanGroupUsersRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "kickgroupusers":
	// 				afterReqFunctions.afterKickGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.KickGroupUsersRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "promotegroupusers":
	// 				afterReqFunctions.afterPromoteGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.PromoteGroupUsersRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "listgroupusers":
	// 				afterReqFunctions.afterListGroupUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.GroupUserList, in *api.ListGroupUsersRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "listusergroups":
	// 				afterReqFunctions.afterListUserGroupsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.UserGroupList, in *api.ListUserGroupsRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "listgroups":
	// 				afterReqFunctions.afterListGroupsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.GroupList, in *api.ListGroupsRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "deleteleaderboardrecord":
	// 				afterReqFunctions.afterDeleteLeaderboardRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteLeaderboardRecordRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "listleaderboardrecords":
	// 				afterReqFunctions.afterListLeaderboardRecordsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.LeaderboardRecordList, in *api.ListLeaderboardRecordsRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "writeleaderboardrecord":
	// 				afterReqFunctions.afterWriteLeaderboardRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.LeaderboardRecord, in *api.WriteLeaderboardRecordRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "listleaderboardrecordsaroundowner":
	// 				afterReqFunctions.afterListLeaderboardRecordsAroundOwnerFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.LeaderboardRecordList, in *api.ListLeaderboardRecordsAroundOwnerRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "linkapple":
	// 				afterReqFunctions.afterLinkAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountApple) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "linkcustom":
	// 				afterReqFunctions.afterLinkCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountCustom) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "linkdevice":
	// 				afterReqFunctions.afterLinkDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountDevice) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "linkemail":
	// 				afterReqFunctions.afterLinkEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountEmail) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "linkfacebook":
	// 				afterReqFunctions.afterLinkFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.LinkFacebookRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "linkfacebookinstantgame":
	// 				afterReqFunctions.afterLinkFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebookInstantGame) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "linkgamecenter":
	// 				afterReqFunctions.afterLinkGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGameCenter) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "linkgoogle":
	// 				afterReqFunctions.afterLinkGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGoogle) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "linksteam":
	// 				afterReqFunctions.afterLinkSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountSteam) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "listmatches":
	// 				afterReqFunctions.afterListMatchesFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.MatchList, in *api.ListMatchesRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "listnotifications":
	// 				afterReqFunctions.afterListNotificationsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.NotificationList, in *api.ListNotificationsRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "deletenotification":
	// 				afterReqFunctions.afterDeleteNotificationFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteNotificationsRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "liststorageobjects":
	// 				afterReqFunctions.afterListStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.StorageObjectList, in *api.ListStorageObjectsRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "readstorageobjects":
	// 				afterReqFunctions.afterReadStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.StorageObjects, in *api.ReadStorageObjectsRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "writestorageobjects":
	// 				afterReqFunctions.afterWriteStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.StorageObjectAcks, in *api.WriteStorageObjectsRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "deletestorageobjects":
	// 				afterReqFunctions.afterDeleteStorageObjectsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.DeleteStorageObjectsRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "jointournament":
	// 				afterReqFunctions.afterJoinTournamentFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.JoinTournamentRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "listtournamentrecords":
	// 				afterReqFunctions.afterListTournamentRecordsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.TournamentRecordList, in *api.ListTournamentRecordsRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "listtournaments":
	// 				afterReqFunctions.afterListTournamentsFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.TournamentList, in *api.ListTournamentsRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "writetournamentrecord":
	// 				afterReqFunctions.afterWriteTournamentRecordFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.LeaderboardRecord, in *api.WriteTournamentRecordRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "listtournamentrecordsaroundowner":
	// 				afterReqFunctions.afterListTournamentRecordsAroundOwnerFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.TournamentRecordList, in *api.ListTournamentRecordsAroundOwnerRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "unlinkapple":
	// 				afterReqFunctions.afterUnlinkAppleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountApple) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "unlinkcustom":
	// 				afterReqFunctions.afterUnlinkCustomFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountCustom) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "unlinkdevice":
	// 				afterReqFunctions.afterUnlinkDeviceFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountDevice) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "unlinkemail":
	// 				afterReqFunctions.afterUnlinkEmailFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountEmail) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "unlinkfacebook":
	// 				afterReqFunctions.afterUnlinkFacebookFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebook) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "unlinkfacebookinstantgame":
	// 				afterReqFunctions.afterUnlinkFacebookInstantGameFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountFacebookInstantGame) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "unlinkgamecenter":
	// 				afterReqFunctions.afterUnlinkGameCenterFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGameCenter) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "unlinkgoogle":
	// 				afterReqFunctions.afterUnlinkGoogleFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountGoogle) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "unlinksteam":
	// 				afterReqFunctions.afterUnlinkSteamFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.AccountSteam) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			case "getusers":
	// 				afterReqFunctions.afterGetUsersFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, out *api.Users, in *api.GetUsersRequest) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, out, in)
	// 				}
	// 			case "event":
	// 				afterReqFunctions.afterEventFunction = func(ctx context.Context, logger *zap.Logger, userID, username string, vars map[string]string, expiry int64, clientIP, clientPort string, in *api.Event) error {
	// 					return runtimeProviderJS.AfterReq(ctx, id, logger, userID, username, vars, expiry, clientIP, clientPort, nil, in)
	// 				}
	// 			}
	// 		}
	// 	case RuntimeExecutionModeMatchmaker:
	// 		matchmakerMatchedFunction = func(ctx context.Context, entries []*MatchmakerEntry) (string, bool, error) {
	// 			return runtimeProviderJS.MatchmakerMatched(ctx, entries)
	// 		}
	// 	case RuntimeExecutionModeTournamentEnd:
	// 		tournamentEndFunction = func(ctx context.Context, tournament *api.Tournament, end, reset int64) error {
	// 			return runtimeProviderJS.TournamentEnd(ctx, tournament, end, reset)
	// 		}
	// 	case RuntimeExecutionModeTournamentReset:
	// 		tournamentResetFunction = func(ctx context.Context, tournament *api.Tournament, end, reset int64) error {
	// 			return runtimeProviderJS.TournamentReset(ctx, tournament, end, reset)
	// 		}
	// 	case RuntimeExecutionModeLeaderboardReset:
	// 		leaderboardResetFunction = func(ctx context.Context, leaderboard runtime.Leaderboard, reset int64) error {
	// 			return runtimeProviderJS.LeaderboardReset(ctx, leaderboard, reset)
	// 		}
	// 	}
	// }, false)
	// if err != nil {
	// 	logger.Error("Failed to eval JavaScript modules.", zap.Error(err))
	// 	return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	// }

	// matchProvider.RegisterCreateFn("javascript",
	// 	func(ctx context.Context, logger *zap.Logger, id uuid.UUID, node string, stopped *atomic.Bool, name string) (RuntimeMatchCore, error) {
	// 		mc := matchCallbacks.Get(name)
	// 		if mc == nil {
	// 			return nil, nil
	// 		}

	// 		return NewRuntimeJavascriptMatchCore(logger, name, db, jsonpbMarshaler, jsonpbUnmarshaler, config, socialClient, leaderboardCache, leaderboardRankCache, leaderboardScheduler, sessionRegistry, matchRegistry, tracker, streamManager, router, matchProvider.CreateMatch, eventFn, id, node, stopped, mc)
	// 	})

	// runtimeProviderJS.newFn = func() *RuntimeJS {
	// 	runtime := goja.New()

	// 	jsLogger := NewJsLogger(logger)
	// 	jsLoggerValue := runtime.ToValue(jsLogger.Constructor(runtime))
	// 	jsLoggerInst, err := runtime.New(jsLoggerValue)
	// 	if err != nil {
	// 		logger.Fatal("Failed to initialize JavaScript runtime", zap.Error(err))
	// 	}

	// 	nakamaModule := NewRuntimeJavascriptNakamaModule(logger, db, jsonpbMarshaler, jsonpbUnmarshaler, config, socialClient, leaderboardCache, leaderboardRankCache, leaderboardScheduler, sessionRegistry, matchRegistry, tracker, streamManager, router, eventFn, matchProvider.CreateMatch)
	// 	nk := runtime.ToValue(nakamaModule.Constructor(runtime))
	// 	nkInst, err := runtime.New(nk)
	// 	if err != nil {
	// 		logger.Fatal("Failed to initialize JavaScript runtime", zap.Error(err))
	// 	}

	// 	return &RuntimeJS{
	// 		logger:       logger,
	// 		jsLoggerInst: jsLoggerInst,
	// 		nkInst:       nkInst,
	// 		node:         config.GetName(),
	// 		vm:           runtime,
	// 		env:          runtime.ToValue(config.GetRuntime().Environment),
	// 		callbacks:    callbacks,
	// 	}
	// }

	// startupLogger.Info("JavaScript runtime modules loaded")

	// // Warm up the pool.
	// startupLogger.Info("Allocating minimum JavaScript runtime pool", zap.Int("count", config.GetRuntime().JsMinCount))
	// if len(modCache.Names) > 0 {
	// 	// Only if there are runtime modules to load.
	// 	for i := 0; i < config.GetRuntime().JsMinCount; i++ {
	// 		runtimeProviderJS.poolCh <- runtimeProviderJS.newFn()
	// 	}
	// 	runtimeProviderJS.metrics.GaugeJsRuntimes(float64(config.GetRuntime().JsMinCount))
	// }
	// startupLogger.Info("Allocated minimum JavaScript runtime pool")

	return modulePaths, rpcFunctions, beforeRtFunctions, afterRtFunctions, beforeReqFunctions, afterReqFunctions, matchmakerMatchedFunction, tournamentEndFunction, tournamentResetFunction, leaderboardResetFunction, nil
}

func openCModule(logger *zap.Logger, rootPath, path string) (string, string, func(context.Context, runtime.Logger, *sql.DB, runtime.NakamaModule, runtime.Initializer) error, error) {
	relPath, _ := filepath.Rel(rootPath, path)
	name := strings.TrimSuffix(relPath, filepath.Ext(relPath))

	// Open the shared library.
	l, err := OpenSharedLib(path)
	if err != nil {
		logger.Error("Could not open C module", zap.String("path", path), zap.Error(err))
		return "", "", nil, err
	}

	// Look up the required initialisation function.
	f, err := l.lookup("init_module")
	if err != nil {
		logger.Fatal("Error looking up init_module function in C module", zap.String("name", name))
		return "", "", nil, err
	}

	// Ensure the function has the correct signature.
	fn, ok := f.(func(context.Context, runtime.Logger, *sql.DB, runtime.NakamaModule, runtime.Initializer) error)
	if !ok {
		logger.Fatal("Error reading init_module function in C module", zap.String("name", name))
		return "", "", nil, errors.New("error reading init_module function in C module")
	}

	return relPath, name, fn, nil
}