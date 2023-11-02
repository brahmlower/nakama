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
	"bytes"
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/rtapi"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/heroiclabs/nakama/v3/internal/cronexpr"
	"github.com/heroiclabs/nakama/v3/internal/satori"
	"github.com/heroiclabs/nakama/v3/social"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
)

type RuntimeCNakamaModule struct {
	sync.RWMutex
	logger               *zap.Logger
	db                   *sql.DB
	protojsonMarshaler   *protojson.MarshalOptions
	config               Config
	socialClient         *social.Client
	leaderboardCache     LeaderboardCache
	leaderboardRankCache LeaderboardRankCache
	leaderboardScheduler LeaderboardScheduler
	sessionRegistry      SessionRegistry
	sessionCache         SessionCache
	statusRegistry       *StatusRegistry
	matchRegistry        MatchRegistry
	tracker              Tracker
	metrics              Metrics
	streamManager        StreamManager
	router               MessageRouter
	satori               runtime.Satori
	storageIndex         StorageIndex

	eventFn RuntimeEventCustomFunction

	node string

	matchCreateFn RuntimeMatchCreateFunction
}

var _ runtime.NakamaModule = (*RuntimeCNakamaModule)(nil)

func NewRuntimeCNakamaModule(
	logger *zap.Logger,
	db *sql.DB,
	protojsonMarshaler *protojson.MarshalOptions,
	config Config,
	socialClient *social.Client,
	leaderboardCache LeaderboardCache,
	leaderboardRankCache LeaderboardRankCache,
	leaderboardScheduler LeaderboardScheduler,
	storageIndex StorageIndex,
	sessionRegistry SessionRegistry,
	sessionCache SessionCache,
	statusRegistry *StatusRegistry,
	matchRegistry MatchRegistry,
	tracker Tracker,
	metrics Metrics,
	streamManager StreamManager,
	router MessageRouter,
) *RuntimeCNakamaModule {
	return &RuntimeCNakamaModule{
		logger:               logger,
		db:                   db,
		protojsonMarshaler:   protojsonMarshaler,
		config:               config,
		socialClient:         socialClient,
		leaderboardCache:     leaderboardCache,
		leaderboardRankCache: leaderboardRankCache,
		leaderboardScheduler: leaderboardScheduler,
		sessionRegistry:      sessionRegistry,
		sessionCache:         sessionCache,
		statusRegistry:       statusRegistry,
		matchRegistry:        matchRegistry,
		tracker:              tracker,
		metrics:              metrics,
		streamManager:        streamManager,
		router:               router,
		storageIndex:         storageIndex,

		node: config.GetName(),

		satori: satori.NewSatoriClient(logger, config.GetSatori().Url, config.GetSatori().ApiKeyName, config.GetSatori().ApiKey, config.GetSatori().SigningKey),
	}
}

func (n *RuntimeCNakamaModule) AuthenticateApple(ctx context.Context, token, username string, create bool) (string, string, bool, error) {
	if n.config.GetSocial().Apple.BundleId == "" {
		return "", "", false, errors.New("Apple authentication is not configured")
	}

	if token == "" {
		return "", "", false, errors.New("expects token string")
	}

	if username == "" {
		username = generateUsername()
	} else if invalidCharsRegex.MatchString(username) {
		return "", "", false, errors.New("expects username to be valid, no spaces or control characters allowed")
	} else if len(username) > 128 {
		return "", "", false, errors.New("expects id to be valid, must be 1-128 bytes")
	}

	return AuthenticateApple(ctx, n.logger, n.db, n.socialClient, n.config.GetSocial().Apple.BundleId, token, username, create)
}

func (n *RuntimeCNakamaModule) AuthenticateCustom(ctx context.Context, id, username string, create bool) (string, string, bool, error) {
	if id == "" {
		return "", "", false, errors.New("expects id string")
	} else if invalidCharsRegex.MatchString(id) {
		return "", "", false, errors.New("expects id to be valid, no spaces or control characters allowed")
	} else if len(id) < 6 || len(id) > 128 {
		return "", "", false, errors.New("expects id to be valid, must be 6-128 bytes")
	}

	if username == "" {
		username = generateUsername()
	} else if invalidCharsRegex.MatchString(username) {
		return "", "", false, errors.New("expects username to be valid, no spaces or control characters allowed")
	} else if len(username) > 128 {
		return "", "", false, errors.New("expects id to be valid, must be 1-128 bytes")
	}

	return AuthenticateCustom(ctx, n.logger, n.db, id, username, create)
}

func (n *RuntimeCNakamaModule) AuthenticateDevice(ctx context.Context, id, username string, create bool) (string, string, bool, error) {
	if id == "" {
		return "", "", false, errors.New("expects id string")
	} else if invalidCharsRegex.MatchString(id) {
		return "", "", false, errors.New("expects id to be valid, no spaces or control characters allowed")
	} else if len(id) < 10 || len(id) > 128 {
		return "", "", false, errors.New("expects id to be valid, must be 10-128 bytes")
	}

	if username == "" {
		username = generateUsername()
	} else if invalidCharsRegex.MatchString(username) {
		return "", "", false, errors.New("expects username to be valid, no spaces or control characters allowed")
	} else if len(username) > 128 {
		return "", "", false, errors.New("expects id to be valid, must be 1-128 bytes")
	}

	return AuthenticateDevice(ctx, n.logger, n.db, id, username, create)
}

func (n *RuntimeCNakamaModule) AuthenticateEmail(ctx context.Context, email, password, username string, create bool) (string, string, bool, error) {
	var attemptUsernameLogin bool
	if email == "" {
		attemptUsernameLogin = true
	} else if invalidCharsRegex.MatchString(email) {
		return "", "", false, errors.New("expects email to be valid, no spaces or control characters allowed")
	} else if !emailRegex.MatchString(email) {
		return "", "", false, errors.New("expects email to be valid, invalid email address format")
	} else if len(email) < 10 || len(email) > 255 {
		return "", "", false, errors.New("expects email to be valid, must be 10-255 bytes")
	}

	if password == "" {
		return "", "", false, errors.New("expects password string")
	} else if len(password) < 8 {
		return "", "", false, errors.New("expects password to be valid, must be longer than 8 characters")
	}

	if username == "" {
		if attemptUsernameLogin {
			return "", "", false, errors.New("expects username string when email is not supplied")
		}

		username = generateUsername()
	} else if invalidCharsRegex.MatchString(username) {
		return "", "", false, errors.New("expects username to be valid, no spaces or control characters allowed")
	} else if len(username) > 128 {
		return "", "", false, errors.New("expects id to be valid, must be 1-128 bytes")
	}

	if attemptUsernameLogin {
		dbUserID, err := AuthenticateUsername(ctx, n.logger, n.db, username, password)
		return dbUserID, username, false, err
	}

	cleanEmail := strings.ToLower(email)

	return AuthenticateEmail(ctx, n.logger, n.db, cleanEmail, password, username, create)
}

func (n *RuntimeCNakamaModule) AuthenticateFacebook(ctx context.Context, token string, importFriends bool, username string, create bool) (string, string, bool, error) {
	if token == "" {
		return "", "", false, errors.New("expects access token string")
	}

	if username == "" {
		username = generateUsername()
	} else if invalidCharsRegex.MatchString(username) {
		return "", "", false, errors.New("expects username to be valid, no spaces or control characters allowed")
	} else if len(username) > 128 {
		return "", "", false, errors.New("expects id to be valid, must be 1-128 bytes")
	}

	dbUserID, dbUsername, created, importFriendsPossible, err := AuthenticateFacebook(ctx, n.logger, n.db, n.socialClient, n.config.GetSocial().FacebookLimitedLogin.AppId, token, username, create)
	if err == nil && importFriends && importFriendsPossible {
		// Errors are logged before this point and failure here does not invalidate the whole operation.
		_ = importFacebookFriends(ctx, n.logger, n.db, n.router, n.socialClient, uuid.FromStringOrNil(dbUserID), dbUsername, token, false)
	}

	return dbUserID, dbUsername, created, err
}

func (n *RuntimeCNakamaModule) AuthenticateFacebookInstantGame(ctx context.Context, signedPlayerInfo string, username string, create bool) (string, string, bool, error) {
	if signedPlayerInfo == "" {
		return "", "", false, errors.New("expects signed player info")
	}

	if username == "" {
		username = generateUsername()
	} else if invalidCharsRegex.MatchString(username) {
		return "", "", false, errors.New("expects username to be valid, no spaces or control characters allowed")
	} else if len(username) > 128 {
		return "", "", false, errors.New("expects id to be valid, must be 1-128 bytes")
	}

	return AuthenticateFacebookInstantGame(ctx, n.logger, n.db, n.socialClient, n.config.GetSocial().FacebookInstantGame.AppSecret, signedPlayerInfo, username, create)
}

func (n *RuntimeCNakamaModule) AuthenticateGameCenter(ctx context.Context, playerID, bundleID string, timestamp int64, salt, signature, publicKeyUrl, username string, create bool) (string, string, bool, error) {
	if playerID == "" {
		return "", "", false, errors.New("expects player ID string")
	}
	if bundleID == "" {
		return "", "", false, errors.New("expects bundle ID string")
	}
	if timestamp == 0 {
		return "", "", false, errors.New("expects timestamp value")
	}
	if salt == "" {
		return "", "", false, errors.New("expects salt string")
	}
	if signature == "" {
		return "", "", false, errors.New("expects signature string")
	}
	if publicKeyUrl == "" {
		return "", "", false, errors.New("expects public key URL string")
	}

	if username == "" {
		username = generateUsername()
	} else if invalidCharsRegex.MatchString(username) {
		return "", "", false, errors.New("expects username to be valid, no spaces or control characters allowed")
	} else if len(username) > 128 {
		return "", "", false, errors.New("expects id to be valid, must be 1-128 bytes")
	}

	return AuthenticateGameCenter(ctx, n.logger, n.db, n.socialClient, playerID, bundleID, timestamp, salt, signature, publicKeyUrl, username, create)
}

func (n *RuntimeCNakamaModule) AuthenticateGoogle(ctx context.Context, token, username string, create bool) (string, string, bool, error) {
	if token == "" {
		return "", "", false, errors.New("expects ID token string")
	}

	if username == "" {
		username = generateUsername()
	} else if invalidCharsRegex.MatchString(username) {
		return "", "", false, errors.New("expects username to be valid, no spaces or control characters allowed")
	} else if len(username) > 128 {
		return "", "", false, errors.New("expects id to be valid, must be 1-128 bytes")
	}

	return AuthenticateGoogle(ctx, n.logger, n.db, n.socialClient, token, username, create)
}

func (n *RuntimeCNakamaModule) AuthenticateSteam(ctx context.Context, token, username string, create bool) (string, string, bool, error) {
	if n.config.GetSocial().Steam.PublisherKey == "" || n.config.GetSocial().Steam.AppID == 0 {
		return "", "", false, errors.New("Steam authentication is not configured")
	}

	if token == "" {
		return "", "", false, errors.New("expects token string")
	}

	if username == "" {
		username = generateUsername()
	} else if invalidCharsRegex.MatchString(username) {
		return "", "", false, errors.New("expects username to be valid, no spaces or control characters allowed")
	} else if len(username) > 128 {
		return "", "", false, errors.New("expects id to be valid, must be 1-128 bytes")
	}

	userID, username, _, created, err := AuthenticateSteam(ctx, n.logger, n.db, n.socialClient, n.config.GetSocial().Steam.AppID, n.config.GetSocial().Steam.PublisherKey, token, username, create)

	return userID, username, created, err
}

func (n *RuntimeCNakamaModule) AuthenticateTokenGenerate(userID, username string, exp int64, vars map[string]string) (string, int64, error) {
	if userID == "" {
		return "", 0, errors.New("expects user id")
	}
	_, err := uuid.FromString(userID)
	if err != nil {
		return "", 0, errors.New("expects valid user id")
	}

	if username == "" {
		return "", 0, errors.New("expects username")
	}

	if exp == 0 {
		// If expiry is 0 or not set, use standard configured expiry.
		exp = time.Now().UTC().Add(time.Duration(n.config.GetSession().TokenExpirySec) * time.Second).Unix()
	}

	tokenId := uuid.Must(uuid.NewV4()).String()
	token, exp := generateTokenWithExpiry(n.config.GetSession().EncryptionKey, tokenId, userID, username, vars, exp)
	return token, exp, nil
}

func (n *RuntimeCNakamaModule) AccountGetId(ctx context.Context, userID string) (*api.Account, error) {
	u, err := uuid.FromString(userID)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	account, err := GetAccount(ctx, n.logger, n.db, n.statusRegistry, u)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (n *RuntimeCNakamaModule) AccountsGetId(ctx context.Context, userIDs []string) ([]*api.Account, error) {
	if len(userIDs) == 0 {
		return make([]*api.Account, 0), nil
	}

	for _, id := range userIDs {
		if _, err := uuid.FromString(id); err != nil {
			return nil, errors.New("each user id must be a valid id string")
		}
	}

	return GetAccounts(ctx, n.logger, n.db, n.statusRegistry, userIDs)
}

func (n *RuntimeCNakamaModule) AccountUpdateId(ctx context.Context, userID, username string, metadata map[string]interface{}, displayName, timezone, location, langTag, avatarUrl string) error {
	u, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("expects user ID to be a valid identifier")
	}

	var metadataWrapper *wrappers.StringValue
	if metadata != nil {
		metadataBytes, err := json.Marshal(metadata)
		if err != nil {
			return errors.Errorf("error encoding metadata: %v", err.Error())
		}
		metadataWrapper = &wrappers.StringValue{Value: string(metadataBytes)}
	}

	var displayNameWrapper *wrappers.StringValue
	if displayName != "" {
		displayNameWrapper = &wrappers.StringValue{Value: displayName}
	}
	var timezoneWrapper *wrappers.StringValue
	if timezone != "" {
		timezoneWrapper = &wrappers.StringValue{Value: timezone}
	}
	var locationWrapper *wrappers.StringValue
	if location != "" {
		locationWrapper = &wrappers.StringValue{Value: location}
	}
	var langWrapper *wrappers.StringValue
	if langTag != "" {
		langWrapper = &wrappers.StringValue{Value: langTag}
	}
	var avatarWrapper *wrappers.StringValue
	if avatarUrl != "" {
		avatarWrapper = &wrappers.StringValue{Value: avatarUrl}
	}

	return UpdateAccounts(ctx, n.logger, n.db, []*accountUpdate{{
		userID:      u,
		username:    username,
		displayName: displayNameWrapper,
		timezone:    timezoneWrapper,
		location:    locationWrapper,
		langTag:     langWrapper,
		avatarURL:   avatarWrapper,
		metadata:    metadataWrapper,
	}})
}

func (n *RuntimeCNakamaModule) AccountDeleteId(ctx context.Context, userID string, recorded bool) error {
	u, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("expects user ID to be a valid identifier")
	}

	return DeleteAccount(ctx, n.logger, n.db, n.config, n.leaderboardCache, n.leaderboardRankCache, n.sessionRegistry, n.sessionCache, n.tracker, u, recorded)
}

func (n *RuntimeCNakamaModule) AccountExportId(ctx context.Context, userID string) (string, error) {
	u, err := uuid.FromString(userID)
	if err != nil {
		return "", errors.New("expects user ID to be a valid identifier")
	}

	export, err := ExportAccount(ctx, n.logger, n.db, u)
	if err != nil {
		return "", errors.Errorf("error exporting account: %v", err.Error())
	}

	exportBytes, err := n.protojsonMarshaler.Marshal(export)
	if err != nil {
		return "", errors.Errorf("error encoding account export: %v", err.Error())
	}

	return string(exportBytes), nil
}

func (n *RuntimeCNakamaModule) UsersGetId(ctx context.Context, userIDs []string, facebookIDs []string) ([]*api.User, error) {
	if len(userIDs) == 0 {
		return make([]*api.User, 0), nil
	}

	for _, id := range userIDs {
		if _, err := uuid.FromString(id); err != nil {
			return nil, errors.New("each user id must be a valid id string")
		}
	}

	users, err := GetUsers(ctx, n.logger, n.db, n.statusRegistry, userIDs, nil, nil)
	if err != nil {
		return nil, err
	}

	return users.Users, nil
}

func (n *RuntimeCNakamaModule) UsersGetUsername(ctx context.Context, usernames []string) ([]*api.User, error) {
	if len(usernames) == 0 {
		return make([]*api.User, 0), nil
	}

	for _, username := range usernames {
		if username == "" {
			return nil, errors.New("each username must be a string")
		}
	}

	users, err := GetUsers(ctx, n.logger, n.db, n.statusRegistry, nil, usernames, nil)
	if err != nil {
		return nil, err
	}

	return users.Users, nil
}

func (n *RuntimeCNakamaModule) UsersGetRandom(ctx context.Context, count int) ([]*api.User, error) {
	if count == 0 {
		return make([]*api.User, 0), nil
	}

	if count < 0 || count > 1000 {
		return nil, errors.New("count must be 0-1000")
	}

	return GetRandomUsers(ctx, n.logger, n.db, n.statusRegistry, count)
}

func (n *RuntimeCNakamaModule) UsersBanId(ctx context.Context, userIDs []string) error {
	if len(userIDs) == 0 {
		return nil
	}

	userUIDs := make([]uuid.UUID, len(userIDs))
	for i, id := range userIDs {
		uid, err := uuid.FromString(id)
		if err != nil {
			return errors.New("each user id must be a valid id string")
		}
		userUIDs[i] = uid
	}

	return BanUsers(ctx, n.logger, n.db, n.config, n.sessionCache, n.sessionRegistry, n.tracker, userUIDs)
}

func (n *RuntimeCNakamaModule) UsersUnbanId(ctx context.Context, userIDs []string) error {
	if len(userIDs) == 0 {
		return nil
	}

	userUIDs := make([]uuid.UUID, len(userIDs))
	for i, id := range userIDs {
		uid, err := uuid.FromString(id)
		if err != nil {
			return errors.New("each user id must be a valid id string")
		}
		userUIDs[i] = uid
	}

	return UnbanUsers(ctx, n.logger, n.db, n.sessionCache, userUIDs)
}

func (n *RuntimeCNakamaModule) LinkApple(ctx context.Context, userID, token string) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return LinkApple(ctx, n.logger, n.db, n.config, n.socialClient, id, token)
}

func (n *RuntimeCNakamaModule) LinkCustom(ctx context.Context, userID, customID string) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return LinkCustom(ctx, n.logger, n.db, id, customID)
}

func (n *RuntimeCNakamaModule) LinkDevice(ctx context.Context, userID, deviceID string) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return LinkDevice(ctx, n.logger, n.db, id, deviceID)
}

func (n *RuntimeCNakamaModule) LinkEmail(ctx context.Context, userID, email, password string) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return LinkEmail(ctx, n.logger, n.db, id, email, password)
}

func (n *RuntimeCNakamaModule) LinkFacebook(ctx context.Context, userID, username, token string, importFriends bool) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return LinkFacebook(ctx, n.logger, n.db, n.socialClient, n.router, id, username, n.config.GetSocial().FacebookLimitedLogin.AppId, token, importFriends)
}

func (n *RuntimeCNakamaModule) LinkFacebookInstantGame(ctx context.Context, userID, signedPlayerInfo string) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return LinkFacebookInstantGame(ctx, n.logger, n.db, n.config, n.socialClient, id, signedPlayerInfo)
}

func (n *RuntimeCNakamaModule) LinkGameCenter(ctx context.Context, userID, playerID, bundleID string, timestamp int64, salt, signature, publicKeyUrl string) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return LinkGameCenter(ctx, n.logger, n.db, n.socialClient, id, playerID, bundleID, timestamp, salt, signature, publicKeyUrl)
}

func (n *RuntimeCNakamaModule) LinkGoogle(ctx context.Context, userID, token string) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return LinkGoogle(ctx, n.logger, n.db, n.socialClient, id, token)
}

func (n *RuntimeCNakamaModule) LinkSteam(ctx context.Context, userID, username string, token string, importFriends bool) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return LinkSteam(ctx, n.logger, n.db, n.config, n.socialClient, n.router, id, username, token, importFriends)
}

func (n *RuntimeCNakamaModule) ReadFile(relPath string) (*os.File, error) {
	return FileRead(n.config.GetRuntime().Path, relPath)
}

func (n *RuntimeCNakamaModule) UnlinkApple(ctx context.Context, userID, token string) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return UnlinkApple(ctx, n.logger, n.db, n.config, n.socialClient, id, token)
}

func (n *RuntimeCNakamaModule) UnlinkCustom(ctx context.Context, userID, customID string) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return UnlinkCustom(ctx, n.logger, n.db, id, customID)
}

func (n *RuntimeCNakamaModule) UnlinkDevice(ctx context.Context, userID, deviceID string) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return UnlinkDevice(ctx, n.logger, n.db, id, deviceID)
}

func (n *RuntimeCNakamaModule) UnlinkEmail(ctx context.Context, userID, email string) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return UnlinkEmail(ctx, n.logger, n.db, id, email)
}

func (n *RuntimeCNakamaModule) UnlinkFacebook(ctx context.Context, userID, token string) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return UnlinkFacebook(ctx, n.logger, n.db, n.socialClient, n.config.GetSocial().FacebookLimitedLogin.AppId, id, token)
}

func (n *RuntimeCNakamaModule) UnlinkFacebookInstantGame(ctx context.Context, userID, signedPlayerInfo string) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return UnlinkFacebookInstantGame(ctx, n.logger, n.db, n.config, n.socialClient, id, signedPlayerInfo)
}

func (n *RuntimeCNakamaModule) UnlinkGameCenter(ctx context.Context, userID, playerID, bundleID string, timestamp int64, salt, signature, publicKeyUrl string) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return UnlinkGameCenter(ctx, n.logger, n.db, n.socialClient, id, playerID, bundleID, timestamp, salt, signature, publicKeyUrl)
}

func (n *RuntimeCNakamaModule) UnlinkGoogle(ctx context.Context, userID, token string) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return UnlinkGoogle(ctx, n.logger, n.db, n.socialClient, id, token)
}

func (n *RuntimeCNakamaModule) UnlinkSteam(ctx context.Context, userID, token string) error {
	id, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("user ID must be a valid identifier")
	}

	return UnlinkSteam(ctx, n.logger, n.db, n.config, n.socialClient, id, token)
}

func (n *RuntimeCNakamaModule) StreamUserList(mode uint8, subject, subcontext, label string, includeHidden, includeNotHidden bool) ([]runtime.Presence, error) {
	stream := PresenceStream{
		Mode:  mode,
		Label: label,
	}
	var err error
	if subject != "" {
		stream.Subject, err = uuid.FromString(subject)
		if err != nil {
			return nil, errors.New("stream subject must be a valid identifier")
		}
	}
	if subcontext != "" {
		stream.Subcontext, err = uuid.FromString(subcontext)
		if err != nil {
			return nil, errors.New("stream subcontext must be a valid identifier")
		}
	}

	presences := n.tracker.ListByStream(stream, includeHidden, includeNotHidden)
	runtimePresences := make([]runtime.Presence, len(presences))
	for i, p := range presences {
		runtimePresences[i] = runtime.Presence(p)
	}
	return runtimePresences, nil
}

func (n *RuntimeCNakamaModule) StreamUserGet(mode uint8, subject, subcontext, label, userID, sessionID string) (runtime.PresenceMeta, error) {
	uid, err := uuid.FromString(userID)
	if err != nil {
		return nil, errors.New("expects valid user id")
	}

	sid, err := uuid.FromString(sessionID)
	if err != nil {
		return nil, errors.New("expects valid session id")
	}

	stream := PresenceStream{
		Mode:  mode,
		Label: label,
	}
	if subject != "" {
		stream.Subject, err = uuid.FromString(subject)
		if err != nil {
			return nil, errors.New("stream subject must be a valid identifier")
		}
	}
	if subcontext != "" {
		stream.Subcontext, err = uuid.FromString(subcontext)
		if err != nil {
			return nil, errors.New("stream subcontext must be a valid identifier")
		}
	}

	if meta := n.tracker.GetLocalBySessionIDStreamUserID(sid, stream, uid); meta != nil {
		return meta, nil
	}
	return nil, nil
}

func (n *RuntimeCNakamaModule) StreamUserJoin(mode uint8, subject, subcontext, label, userID, sessionID string, hidden, persistence bool, status string) (bool, error) {
	uid, err := uuid.FromString(userID)
	if err != nil {
		return false, errors.New("expects valid user id")
	}

	sid, err := uuid.FromString(sessionID)
	if err != nil {
		return false, errors.New("expects valid session id")
	}

	stream := PresenceStream{
		Mode:  mode,
		Label: label,
	}
	if subject != "" {
		stream.Subject, err = uuid.FromString(subject)
		if err != nil {
			return false, errors.New("stream subject must be a valid identifier")
		}
	}
	if subcontext != "" {
		stream.Subcontext, err = uuid.FromString(subcontext)
		if err != nil {
			return false, errors.New("stream subcontext must be a valid identifier")
		}
	}

	success, newlyTracked, err := n.streamManager.UserJoin(stream, uid, sid, hidden, persistence, status)
	if err != nil {
		return false, err
	}
	if !success {
		return false, errors.New("tracker rejected new presence, session is closing")
	}

	return newlyTracked, nil
}

func (n *RuntimeCNakamaModule) StreamUserUpdate(mode uint8, subject, subcontext, label, userID, sessionID string, hidden, persistence bool, status string) error {
	uid, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("expects valid user id")
	}

	sid, err := uuid.FromString(sessionID)
	if err != nil {
		return errors.New("expects valid session id")
	}

	stream := PresenceStream{
		Mode:  mode,
		Label: label,
	}
	if subject != "" {
		stream.Subject, err = uuid.FromString(subject)
		if err != nil {
			return errors.New("stream subject must be a valid identifier")
		}
	}
	if subcontext != "" {
		stream.Subcontext, err = uuid.FromString(subcontext)
		if err != nil {
			return errors.New("stream subcontext must be a valid identifier")
		}
	}

	success, err := n.streamManager.UserUpdate(stream, uid, sid, hidden, persistence, status)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("tracker rejected updated presence, session is closing")
	}

	return nil
}

func (n *RuntimeCNakamaModule) StreamUserLeave(mode uint8, subject, subcontext, label, userID, sessionID string) error {
	uid, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("expects valid user id")
	}

	sid, err := uuid.FromString(sessionID)
	if err != nil {
		return errors.New("expects valid session id")
	}

	stream := PresenceStream{
		Mode:  mode,
		Label: label,
	}
	if subject != "" {
		stream.Subject, err = uuid.FromString(subject)
		if err != nil {
			return errors.New("stream subject must be a valid identifier")
		}
	}
	if subcontext != "" {
		stream.Subcontext, err = uuid.FromString(subcontext)
		if err != nil {
			return errors.New("stream subcontext must be a valid identifier")
		}
	}

	return n.streamManager.UserLeave(stream, uid, sid)
}

func (n *RuntimeCNakamaModule) StreamUserKick(mode uint8, subject, subcontext, label string, presence runtime.Presence) error {
	uid, err := uuid.FromString(presence.GetUserId())
	if err != nil {
		return errors.New("expects valid user id")
	}

	sid, err := uuid.FromString(presence.GetSessionId())
	if err != nil {
		return errors.New("expects valid session id")
	}

	stream := PresenceStream{
		Mode:  mode,
		Label: label,
	}
	if subject != "" {
		stream.Subject, err = uuid.FromString(subject)
		if err != nil {
			return errors.New("stream subject must be a valid identifier")
		}
	}
	if subcontext != "" {
		stream.Subcontext, err = uuid.FromString(subcontext)
		if err != nil {
			return errors.New("stream subcontext must be a valid identifier")
		}
	}

	return n.streamManager.UserLeave(stream, uid, sid)
}

func (n *RuntimeCNakamaModule) StreamCount(mode uint8, subject, subcontext, label string) (int, error) {
	stream := PresenceStream{
		Mode:  mode,
		Label: label,
	}
	var err error
	if subject != "" {
		stream.Subject, err = uuid.FromString(subject)
		if err != nil {
			return 0, errors.New("stream subject must be a valid identifier")
		}
	}
	if subcontext != "" {
		stream.Subcontext, err = uuid.FromString(subcontext)
		if err != nil {
			return 0, errors.New("stream subcontext must be a valid identifier")
		}
	}

	return n.tracker.CountByStream(stream), nil
}

func (n *RuntimeCNakamaModule) StreamClose(mode uint8, subject, subcontext, label string) error {
	stream := PresenceStream{
		Mode:  mode,
		Label: label,
	}
	var err error
	if subject != "" {
		stream.Subject, err = uuid.FromString(subject)
		if err != nil {
			return errors.New("stream subject must be a valid identifier")
		}
	}
	if subcontext != "" {
		stream.Subcontext, err = uuid.FromString(subcontext)
		if err != nil {
			return errors.New("stream subcontext must be a valid identifier")
		}
	}

	n.tracker.UntrackByStream(stream)

	return nil
}

func (n *RuntimeCNakamaModule) StreamSend(mode uint8, subject, subcontext, label, data string, presences []runtime.Presence, reliable bool) error {
	stream := PresenceStream{
		Mode:  mode,
		Label: label,
	}
	var err error
	if subject != "" {
		stream.Subject, err = uuid.FromString(subject)
		if err != nil {
			return errors.New("stream subject must be a valid identifier")
		}
	}
	if subcontext != "" {
		stream.Subcontext, err = uuid.FromString(subcontext)
		if err != nil {
			return errors.New("stream subcontext must be a valid identifier")
		}
	}

	var presenceIDs []*PresenceID
	if l := len(presences); l != 0 {
		presenceIDs = make([]*PresenceID, 0, l)
		for _, presence := range presences {
			sessionID, err := uuid.FromString(presence.GetSessionId())
			if err != nil {
				return errors.New("expects each presence session id to be a valid identifier")
			}
			node := presence.GetNodeId()
			if node == "" {
				node = n.node
			}

			presenceIDs = append(presenceIDs, &PresenceID{
				SessionID: sessionID,
				Node:      node,
			})
		}
	}

	streamWire := &rtapi.Stream{
		Mode:  int32(stream.Mode),
		Label: stream.Label,
	}
	if stream.Subject != uuid.Nil {
		streamWire.Subject = stream.Subject.String()
	}
	if stream.Subcontext != uuid.Nil {
		streamWire.Subcontext = stream.Subcontext.String()
	}
	msg := &rtapi.Envelope{Message: &rtapi.Envelope_StreamData{StreamData: &rtapi.StreamData{
		Stream: streamWire,
		// No sender.
		Data:     data,
		Reliable: reliable,
	}}}

	if len(presenceIDs) == 0 {
		// Sending to whole stream.
		n.router.SendToStream(n.logger, stream, msg, reliable)
	} else {
		// Sending to a subset of stream users.
		n.router.SendToPresenceIDs(n.logger, presenceIDs, msg, reliable)
	}

	return nil
}

func (n *RuntimeCNakamaModule) StreamSendRaw(mode uint8, subject, subcontext, label string, msg *rtapi.Envelope, presences []runtime.Presence, reliable bool) error {
	stream := PresenceStream{
		Mode:  mode,
		Label: label,
	}
	var err error
	if subject != "" {
		stream.Subject, err = uuid.FromString(subject)
		if err != nil {
			return errors.New("stream subject must be a valid identifier")
		}
	}
	if subcontext != "" {
		stream.Subcontext, err = uuid.FromString(subcontext)
		if err != nil {
			return errors.New("stream subcontext must be a valid identifier")
		}
	}
	if msg == nil {
		return errors.New("expects a valid message")
	}

	var presenceIDs []*PresenceID
	if l := len(presences); l != 0 {
		presenceIDs = make([]*PresenceID, 0, l)
		for _, presence := range presences {
			sessionID, err := uuid.FromString(presence.GetSessionId())
			if err != nil {
				return errors.New("expects each presence session id to be a valid identifier")
			}
			node := presence.GetNodeId()
			if node == "" {
				node = n.node
			}

			presenceIDs = append(presenceIDs, &PresenceID{
				SessionID: sessionID,
				Node:      node,
			})
		}
	}

	if len(presenceIDs) == 0 {
		// Sending to whole stream.
		n.router.SendToStream(n.logger, stream, msg, reliable)
	} else {
		// Sending to a subset of stream users.
		n.router.SendToPresenceIDs(n.logger, presenceIDs, msg, reliable)
	}

	return nil
}

func (n *RuntimeCNakamaModule) SessionDisconnect(ctx context.Context, sessionID string, reason ...runtime.PresenceReason) error {
	sid, err := uuid.FromString(sessionID)
	if err != nil {
		return errors.New("expects valid session id")
	}

	return n.sessionRegistry.Disconnect(ctx, sid, false, reason...)
}

func (n *RuntimeCNakamaModule) SessionLogout(userID, token, refreshToken string) error {
	uid, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("expects valid user id")
	}

	return SessionLogout(n.config, n.sessionCache, uid, token, refreshToken)
}

func (n *RuntimeCNakamaModule) MatchCreate(ctx context.Context, module string, params map[string]interface{}) (string, error) {
	if module == "" {
		return "", errors.New("expects module name")
	}

	n.RLock()
	fn := n.matchCreateFn
	n.RUnlock()

	return n.matchRegistry.CreateMatch(ctx, fn, module, params)
}

func (n *RuntimeCNakamaModule) MatchGet(ctx context.Context, id string) (*api.Match, error) {
	result, _, err := n.matchRegistry.GetMatch(ctx, id)
	return result, err
}

// func (n *RuntimeCNakamaModule) MatchList(ctx context.Context, limit int, authoritative bool, label string, minSize, maxSize *int, query string) ([]*api.Match, []string, error) {
func (n *RuntimeCNakamaModule) MatchList(ctx context.Context, limit int, authoritative bool, label string, minSize, maxSize *int, query string) ([]*api.Match, error) {
	authoritativeWrapper := &wrappers.BoolValue{Value: authoritative}
	var labelWrapper *wrappers.StringValue
	if label != "" {
		labelWrapper = &wrappers.StringValue{Value: label}
	}
	var queryWrapper *wrappers.StringValue
	if query != "" {
		queryWrapper = &wrappers.StringValue{Value: query}
	}
	var minSizeWrapper *wrappers.Int32Value
	if minSize != nil {
		minSizeWrapper = &wrappers.Int32Value{Value: int32(*minSize)}
	}
	var maxSizeWrapper *wrappers.Int32Value
	if maxSize != nil {
		maxSizeWrapper = &wrappers.Int32Value{Value: int32(*maxSize)}
	}

	matches, _, err := n.matchRegistry.ListMatches(ctx, limit, authoritativeWrapper, labelWrapper, minSizeWrapper, maxSizeWrapper, queryWrapper, nil)
	return matches, err
}

func (n *RuntimeCNakamaModule) MatchSignal(ctx context.Context, id string, data string) (string, error) {
	return n.matchRegistry.Signal(ctx, id, data)
}

func (n *RuntimeCNakamaModule) NotificationSend(ctx context.Context, userID, subject string, content map[string]interface{}, code int, sender string, persistent bool) error {
	uid, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("expects userID to be a valid UUID")
	}

	if subject == "" {
		return errors.New("expects subject to be a non-empty string")
	}

	contentBytes, err := json.Marshal(content)
	if err != nil {
		return errors.Errorf("failed to convert content: %s", err.Error())
	}
	contentString := string(contentBytes)

	if code <= 0 {
		return errors.New("expects code to number above 0")
	}

	senderID := uuid.Nil.String()
	if sender != "" {
		suid, err := uuid.FromString(sender)
		if err != nil {
			return errors.New("expects sender to either be an empty string or a valid UUID")
		}
		senderID = suid.String()
	}

	nots := []*api.Notification{{
		Id:         uuid.Must(uuid.NewV4()).String(),
		Subject:    subject,
		Content:    contentString,
		Code:       int32(code),
		SenderId:   senderID,
		Persistent: persistent,
		CreateTime: &timestamp.Timestamp{Seconds: time.Now().UTC().Unix()},
	}}
	notifications := map[uuid.UUID][]*api.Notification{
		uid: nots,
	}

	return NotificationSend(ctx, n.logger, n.db, n.router, notifications)
}

func (n *RuntimeCNakamaModule) NotificationsSend(ctx context.Context, notifications []*runtime.NotificationSend) error {
	ns := make(map[uuid.UUID][]*api.Notification)

	for _, notification := range notifications {
		uid, err := uuid.FromString(notification.UserID)
		if err != nil {
			return errors.New("expects userID to be a valid UUID")
		}

		if notification.Subject == "" {
			return errors.New("expects subject to be a non-empty string")
		}

		contentBytes, err := json.Marshal(notification.Content)
		if err != nil {
			return errors.Errorf("failed to convert content: %s", err.Error())
		}
		contentString := string(contentBytes)

		if notification.Code <= 0 {
			return errors.New("expects code to number above 0")
		}

		senderID := uuid.Nil.String()
		if notification.Sender != "" {
			suid, err := uuid.FromString(notification.Sender)
			if err != nil {
				return errors.New("expects sender to either be an empty string or a valid UUID")
			}
			senderID = suid.String()
		}

		no := ns[uid]
		if no == nil {
			no = make([]*api.Notification, 0)
		}
		no = append(no, &api.Notification{
			Id:         uuid.Must(uuid.NewV4()).String(),
			Subject:    notification.Subject,
			Content:    contentString,
			Code:       int32(notification.Code),
			SenderId:   senderID,
			Persistent: notification.Persistent,
			CreateTime: &timestamp.Timestamp{Seconds: time.Now().UTC().Unix()},
		})
		ns[uid] = no
	}

	return NotificationSend(ctx, n.logger, n.db, n.router, ns)
}

func (n *RuntimeCNakamaModule) NotificationSendAll(ctx context.Context, subject string, content map[string]interface{}, code int, persistent bool) error {
	// TODO: Implement this
	return nil
}

func (n *RuntimeCNakamaModule) NotificationsDelete(ctx context.Context, notifications []*runtime.NotificationDelete) error {
	// TODO: Implement this
	return nil
}

func (n *RuntimeCNakamaModule) WalletUpdate(ctx context.Context, userID string, changeset map[string]int64, metadata map[string]interface{}, updateLedger bool) (map[string]int64, map[string]int64, error) {
	uid, err := uuid.FromString(userID)
	if err != nil {
		return nil, nil, errors.New("expects a valid user id")
	}

	metadataBytes := []byte("{}")
	if metadata != nil {
		metadataBytes, err = json.Marshal(metadata)
		if err != nil {
			return nil, nil, errors.Errorf("failed to convert metadata: %s", err.Error())
		}
	}

	results, err := UpdateWallets(ctx, n.logger, n.db, []*walletUpdate{{
		UserID:    uid,
		Changeset: changeset,
		Metadata:  string(metadataBytes),
	}}, updateLedger)
	if err != nil {
		if len(results) == 0 {
			return nil, nil, err
		}
		return results[0].Updated, results[0].Previous, err
	}

	return results[0].Updated, results[0].Previous, nil
}

func (n *RuntimeCNakamaModule) WalletsUpdate(ctx context.Context, updates []*runtime.WalletUpdate, updateLedger bool) ([]*runtime.WalletUpdateResult, error) {
	size := len(updates)
	if size == 0 {
		return nil, nil
	}

	walletUpdates := make([]*walletUpdate, size)

	for i, update := range updates {
		uid, err := uuid.FromString(update.UserID)
		if err != nil {
			return nil, errors.New("expects a valid user id")
		}

		metadataBytes := []byte("{}")
		if update.Metadata != nil {
			metadataBytes, err = json.Marshal(update.Metadata)
			if err != nil {
				return nil, errors.Errorf("failed to convert metadata: %s", err.Error())
			}
		}

		walletUpdates[i] = &walletUpdate{
			UserID:    uid,
			Changeset: update.Changeset,
			Metadata:  string(metadataBytes),
		}
	}

	return UpdateWallets(ctx, n.logger, n.db, walletUpdates, updateLedger)
}

func (n *RuntimeCNakamaModule) WalletLedgerUpdate(ctx context.Context, itemID string, metadata map[string]interface{}) (runtime.WalletLedgerItem, error) {
	id, err := uuid.FromString(itemID)
	if err != nil {
		return nil, errors.New("expects a valid item id")
	}

	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		return nil, errors.Errorf("failed to convert metadata: %s", err.Error())
	}

	return UpdateWalletLedger(ctx, n.logger, n.db, id, string(metadataBytes))
}

func (n *RuntimeCNakamaModule) WalletLedgerList(ctx context.Context, userID string, limit int, cursor string) ([]runtime.WalletLedgerItem, string, error) {
	uid, err := uuid.FromString(userID)
	if err != nil {
		return nil, "", errors.New("expects a valid user id")
	}

	if limit < 0 || limit > 100 {
		return nil, "", errors.New("expects limit to be 0-100")
	}

	items, newCursor, _, err := ListWalletLedger(ctx, n.logger, n.db, uid, &limit, cursor)
	if err != nil {
		return nil, "", err
	}

	runtimeItems := make([]runtime.WalletLedgerItem, len(items))
	for i, item := range items {
		runtimeItems[i] = runtime.WalletLedgerItem(item)
	}
	return runtimeItems, newCursor, nil
}

func (n *RuntimeCNakamaModule) StorageList(ctx context.Context, callerID, userID, collection string, limit int, cursor string) ([]*api.StorageObject, string, error) {
	cid := uuid.Nil
	if callerID != "" {
		u, err := uuid.FromString(callerID)
		if err != nil {
			return nil, "", errors.New("expects an empty or valid caller id")
		}
		cid = u
	}

	var uid *uuid.UUID
	if userID != "" {
		u, err := uuid.FromString(userID)
		if err != nil {
			return nil, "", errors.New("expects an empty or valid user id")
		}
		uid = &u
	}

	objectList, _, err := StorageListObjects(ctx, n.logger, n.db, cid, uid, collection, limit, cursor)
	if err != nil {
		return nil, "", err
	}

	return objectList.Objects, objectList.Cursor, nil
}

func (n *RuntimeCNakamaModule) StorageRead(ctx context.Context, reads []*runtime.StorageRead) ([]*api.StorageObject, error) {
	size := len(reads)
	if size == 0 {
		return make([]*api.StorageObject, 0), nil
	}
	objectIDs := make([]*api.ReadStorageObjectId, size)

	for i, read := range reads {
		if read.Collection == "" {
			return nil, errors.New("expects collection to be a non-empty string")
		}
		if read.Key == "" {
			return nil, errors.New("expects key to be a non-empty string")
		}
		uid := uuid.Nil
		var err error
		if read.UserID != "" {
			uid, err = uuid.FromString(read.UserID)
			if err != nil {
				return nil, errors.New("expects an empty or valid user id")
			}
		}

		objectIDs[i] = &api.ReadStorageObjectId{
			Collection: read.Collection,
			Key:        read.Key,
			UserId:     uid.String(),
		}
	}

	objects, err := StorageReadObjects(ctx, n.logger, n.db, uuid.Nil, objectIDs)
	if err != nil {
		return nil, err
	}

	return objects.Objects, nil
}

func (n *RuntimeCNakamaModule) StorageWrite(ctx context.Context, writes []*runtime.StorageWrite) ([]*api.StorageObjectAck, error) {
	size := len(writes)
	if size == 0 {
		return make([]*api.StorageObjectAck, 0), nil
	}

	ops := make(StorageOpWrites, 0, size)

	for _, write := range writes {
		if write.Collection == "" {
			return nil, errors.New("expects collection to be a non-empty string")
		}
		if write.Key == "" {
			return nil, errors.New("expects key to be a non-empty string")
		}
		if write.UserID != "" {
			if _, err := uuid.FromString(write.UserID); err != nil {
				return nil, errors.New("expects an empty or valid user id")
			}
		}
		if maybeJSON := []byte(write.Value); !json.Valid(maybeJSON) || bytes.TrimSpace(maybeJSON)[0] != byteBracket {
			return nil, errors.New("value must be a JSON-encoded object")
		}

		op := &StorageOpWrite{
			Object: &api.WriteStorageObject{
				Collection:      write.Collection,
				Key:             write.Key,
				Value:           write.Value,
				Version:         write.Version,
				PermissionRead:  &wrappers.Int32Value{Value: int32(write.PermissionRead)},
				PermissionWrite: &wrappers.Int32Value{Value: int32(write.PermissionWrite)},
			},
		}
		if write.UserID == "" {
			op.OwnerID = uuid.Nil.String()
		} else {
			op.OwnerID = write.UserID
		}

		ops = append(ops, op)
	}

	acks, _, err := StorageWriteObjects(ctx, n.logger, n.db, n.metrics, n.storageIndex, true, ops)
	if err != nil {
		return nil, err
	}

	return acks.Acks, nil
}

func (n *RuntimeCNakamaModule) StorageDelete(ctx context.Context, deletes []*runtime.StorageDelete) error {
	size := len(deletes)
	if size == 0 {
		return nil
	}

	ops := make(StorageOpDeletes, 0, size)

	for _, del := range deletes {
		if del.Collection == "" {
			return errors.New("expects collection to be a non-empty string")
		}
		if del.Key == "" {
			return errors.New("expects key to be a non-empty string")
		}
		if del.UserID != "" {
			if _, err := uuid.FromString(del.UserID); err != nil {
				return errors.New("expects an empty or valid user id")
			}
		}

		op := &StorageOpDelete{
			ObjectID: &api.DeleteStorageObjectId{
				Collection: del.Collection,
				Key:        del.Key,
				Version:    del.Version,
			},
		}
		if del.UserID == "" {
			op.OwnerID = uuid.Nil.String()
		} else {
			op.OwnerID = del.UserID
		}

		ops = append(ops, op)
	}

	_, err := StorageDeleteObjects(ctx, n.logger, n.db, n.storageIndex, true, ops)

	return err
}

func (n *RuntimeCNakamaModule) StorageIndexList(ctx context.Context, callerID, indexName, query string, limit int) (*api.StorageObjects, error) {
	cid := uuid.Nil
	if callerID != "" {
		id, err := uuid.FromString(callerID)
		if err != nil {
			return nil, errors.New("expects caller id to be empty or a valid user id")
		}
		cid = id
	}

	if indexName == "" {
		return nil, errors.New("expects a non-empty indexName")
	}
	if limit < 1 || limit > 100 {
		return nil, errors.New("limit must be 1-100")
	}

	return n.storageIndex.List(ctx, cid, indexName, query, limit)
}

func (n *RuntimeCNakamaModule) MultiUpdate(ctx context.Context, accountUpdates []*runtime.AccountUpdate, storageWrites []*runtime.StorageWrite, walletUpdates []*runtime.WalletUpdate, updateLedger bool) ([]*api.StorageObjectAck, []*runtime.WalletUpdateResult, error) {
	// Process account update inputs.
	accountUpdateOps := make([]*accountUpdate, 0, len(accountUpdates))
	for _, update := range accountUpdates {
		u, err := uuid.FromString(update.UserID)
		if err != nil {
			return nil, nil, errors.New("expects user ID to be a valid identifier")
		}

		var metadataWrapper *wrappers.StringValue
		if update.Metadata != nil {
			metadataBytes, err := json.Marshal(update.Metadata)
			if err != nil {
				return nil, nil, errors.Errorf("error encoding metadata: %v", err.Error())
			}
			metadataWrapper = &wrappers.StringValue{Value: string(metadataBytes)}
		}

		var displayNameWrapper *wrappers.StringValue
		if update.DisplayName != "" {
			displayNameWrapper = &wrappers.StringValue{Value: update.DisplayName}
		}
		var timezoneWrapper *wrappers.StringValue
		if update.Timezone != "" {
			timezoneWrapper = &wrappers.StringValue{Value: update.Timezone}
		}
		var locationWrapper *wrappers.StringValue
		if update.Location != "" {
			locationWrapper = &wrappers.StringValue{Value: update.Location}
		}
		var langWrapper *wrappers.StringValue
		if update.LangTag != "" {
			langWrapper = &wrappers.StringValue{Value: update.LangTag}
		}
		var avatarWrapper *wrappers.StringValue
		if update.AvatarUrl != "" {
			avatarWrapper = &wrappers.StringValue{Value: update.AvatarUrl}
		}

		accountUpdateOps = append(accountUpdateOps, &accountUpdate{
			userID:      u,
			username:    update.Username,
			displayName: displayNameWrapper,
			timezone:    timezoneWrapper,
			location:    locationWrapper,
			langTag:     langWrapper,
			avatarURL:   avatarWrapper,
			metadata:    metadataWrapper,
		})
	}

	// Process storage write inputs.
	storageWriteOps := make(StorageOpWrites, 0, len(storageWrites))
	for _, write := range storageWrites {
		if write.Collection == "" {
			return nil, nil, errors.New("expects collection to be a non-empty string")
		}
		if write.Key == "" {
			return nil, nil, errors.New("expects key to be a non-empty string")
		}
		if write.UserID != "" {
			if _, err := uuid.FromString(write.UserID); err != nil {
				return nil, nil, errors.New("expects an empty or valid user id")
			}
		}
		if maybeJSON := []byte(write.Value); !json.Valid(maybeJSON) || bytes.TrimSpace(maybeJSON)[0] != byteBracket {
			return nil, nil, errors.New("value must be a JSON-encoded object")
		}

		op := &StorageOpWrite{
			Object: &api.WriteStorageObject{
				Collection:      write.Collection,
				Key:             write.Key,
				Value:           write.Value,
				Version:         write.Version,
				PermissionRead:  &wrappers.Int32Value{Value: int32(write.PermissionRead)},
				PermissionWrite: &wrappers.Int32Value{Value: int32(write.PermissionWrite)},
			},
		}
		if write.UserID == "" {
			op.OwnerID = uuid.Nil.String()
		} else {
			op.OwnerID = write.UserID
		}

		storageWriteOps = append(storageWriteOps, op)
	}

	// Process wallet update inputs.
	walletUpdateOps := make([]*walletUpdate, len(walletUpdates))
	for i, update := range walletUpdates {
		uid, err := uuid.FromString(update.UserID)
		if err != nil {
			return nil, nil, errors.New("expects a valid user id")
		}

		metadataBytes := []byte("{}")
		if update.Metadata != nil {
			metadataBytes, err = json.Marshal(update.Metadata)
			if err != nil {
				return nil, nil, errors.Errorf("failed to convert metadata: %s", err.Error())
			}
		}

		walletUpdateOps[i] = &walletUpdate{
			UserID:    uid,
			Changeset: update.Changeset,
			Metadata:  string(metadataBytes),
		}
	}

	return MultiUpdate(ctx, n.logger, n.db, n.metrics, accountUpdateOps, storageWriteOps, walletUpdateOps, updateLedger)
}

func (n *RuntimeCNakamaModule) LeaderboardCreate(ctx context.Context, id string, authoritative bool, sortOrder, operator, resetSchedule string, metadata map[string]interface{}) error {
	if id == "" {
		return errors.New("expects a leaderboard ID string")
	}

	sort := LeaderboardSortOrderDescending
	switch sortOrder {
	case "desc":
		sort = LeaderboardSortOrderDescending
	case "asc":
		sort = LeaderboardSortOrderAscending
	default:
		return errors.New("expects sort order to be 'asc' or 'desc'")
	}

	oper := LeaderboardOperatorBest
	switch operator {
	case "best":
		oper = LeaderboardOperatorBest
	case "set":
		oper = LeaderboardOperatorSet
	case "incr":
		oper = LeaderboardOperatorIncrement
	default:
		return errors.New("expects sort order to be 'best', 'set', or 'incr'")
	}

	if resetSchedule != "" {
		if _, err := cronexpr.Parse(resetSchedule); err != nil {
			return errors.New("expects reset schedule to be a valid CRON expression")
		}
	}

	metadataStr := "{}"
	if metadata != nil {
		metadataBytes, err := json.Marshal(metadata)
		if err != nil {
			return errors.Errorf("error encoding metadata: %v", err.Error())
		}
		metadataStr = string(metadataBytes)
	}

	_, created, err := n.leaderboardCache.Create(ctx, id, authoritative, sort, oper, resetSchedule, metadataStr)
	if err != nil {
		return err
	}

	if created {
		// Only need to update the scheduler for newly created leaderboards.
		n.leaderboardScheduler.Update()
	}

	return nil
}

func (n *RuntimeCNakamaModule) LeaderboardDelete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("expects a leaderboard ID string")
	}

	_, err := n.leaderboardCache.Delete(ctx, n.leaderboardRankCache, n.leaderboardScheduler, id)
	if err != nil {
		return err
	}

	return nil
}

func (n *RuntimeCNakamaModule) LeaderboardList(limit int, cursor string) (*api.LeaderboardList, error) {
	// TODO: Implement this
	return nil, nil
}

func (n *RuntimeCNakamaModule) LeaderboardRecordsList(ctx context.Context, id string, ownerIDs []string, limit int, cursor string, expiry int64) ([]*api.LeaderboardRecord, []*api.LeaderboardRecord, string, string, error) {
	if id == "" {
		return nil, nil, "", "", errors.New("expects a leaderboard ID string")
	}

	for _, o := range ownerIDs {
		if _, err := uuid.FromString(o); err != nil {
			return nil, nil, "", "", errors.New("expects each owner ID to be a valid identifier")
		}
	}

	var limitWrapper *wrappers.Int32Value
	if limit < 0 || limit > 10000 {
		return nil, nil, "", "", errors.New("expects limit to be 0-10000")
	}
	limitWrapper = &wrappers.Int32Value{Value: int32(limit)}

	if expiry < 0 {
		return nil, nil, "", "", errors.New("expects expiry to equal or greater than 0")
	}

	list, err := LeaderboardRecordsList(ctx, n.logger, n.db, n.leaderboardCache, n.leaderboardRankCache, id, limitWrapper, cursor, ownerIDs, expiry)
	if err != nil {
		return nil, nil, "", "", err
	}

	return list.Records, list.OwnerRecords, list.NextCursor, list.PrevCursor, nil
}

func (n *RuntimeCNakamaModule) LeaderboardRecordsListCursorFromRank(id string, rank, overrideExpiry int64) (string, error) {
	// TODO: Implement this
	return "", nil
}

func (n *RuntimeCNakamaModule) LeaderboardRecordWrite(ctx context.Context, id, ownerID, username string, score, subscore int64, metadata map[string]interface{}, overrideOperator *int) (*api.LeaderboardRecord, error) {
	if id == "" {
		return nil, errors.New("expects a leaderboard ID string")
	}

	if _, err := uuid.FromString(ownerID); err != nil {
		return nil, errors.New("expects owner ID to be a valid identifier")
	}

	// Username is optional.

	if score < 0 {
		return nil, errors.New("expects score to be >= 0")
	}
	if subscore < 0 {
		return nil, errors.New("expects subscore to be >= 0")
	}

	metadataStr := ""
	if metadata != nil {
		metadataBytes, err := json.Marshal(metadata)
		if err != nil {
			return nil, errors.Errorf("error encoding metadata: %v", err.Error())
		}
		metadataStr = string(metadataBytes)
	}

	operator := api.Operator_NO_OVERRIDE
	if overrideOperator != nil {
		if _, ok := api.Operator_name[int32(*overrideOperator)]; !ok {
			return nil, ErrInvalidOperator
		}
		operator = api.Operator(*overrideOperator)
	}

	return LeaderboardRecordWrite(ctx, n.logger, n.db, n.leaderboardCache, n.leaderboardRankCache, uuid.Nil, id, ownerID, username, score, subscore, metadataStr, operator)
}

func (n *RuntimeCNakamaModule) LeaderboardRecordDelete(ctx context.Context, id, ownerID string) error {
	if id == "" {
		return errors.New("expects a leaderboard ID string")
	}

	if _, err := uuid.FromString(ownerID); err != nil {
		return errors.New("expects owner ID to be a valid identifier")
	}

	return LeaderboardRecordDelete(ctx, n.logger, n.db, n.leaderboardCache, n.leaderboardRankCache, uuid.Nil, id, ownerID)
}

func (n *RuntimeCNakamaModule) LeaderboardRecordsHaystack(ctx context.Context, id, ownerID string, limit int, cursor string, expiry int64) (*api.LeaderboardRecordList, error) {
	// TODO: Implement this
	return nil, nil
}

func (n *RuntimeCNakamaModule) LeaderboardsGetId(ctx context.Context, IDs []string) ([]*api.Leaderboard, error) {
	return LeaderboardsGet(n.leaderboardCache, IDs), nil
}

func (n *RuntimeCNakamaModule) TournamentCreate(ctx context.Context, id string, authoritative bool, sortOrder, operator, resetSchedule string, metadata map[string]interface{}, title, description string, category, startTime, endTime, duration, maxSize, maxNumScore int, joinRequired bool) error {
	if id == "" {
		return errors.New("expects a tournament ID string")
	}

	sort := LeaderboardSortOrderDescending
	switch sortOrder {
	case "desc":
		sort = LeaderboardSortOrderDescending
	case "asc":
		sort = LeaderboardSortOrderAscending
	default:
		return errors.New("expects sort order to be 'asc' or 'desc'")
	}

	oper := LeaderboardOperatorBest
	switch operator {
	case "best":
		oper = LeaderboardOperatorBest
	case "set":
		oper = LeaderboardOperatorSet
	case "incr":
		oper = LeaderboardOperatorIncrement
	default:
		return errors.New("expects sort order to be 'best', 'set', or 'incr'")
	}

	if resetSchedule != "" {
		if _, err := cronexpr.Parse(resetSchedule); err != nil {
			return errors.New("expects reset schedule to be a valid CRON expression")
		}
	}

	metadataStr := "{}"
	if metadata != nil {
		metadataBytes, err := json.Marshal(metadata)
		if err != nil {
			return errors.Errorf("error encoding metadata: %v", err.Error())
		}
		metadataStr = string(metadataBytes)
	}

	if category < 0 || category >= 128 {
		return errors.New("category must be 0-127")
	}
	if startTime < 0 {
		return errors.New("startTime must be >= 0")
	}
	if endTime < 0 {
		return errors.New("endTime must be >= 0")
	}
	if endTime != 0 && endTime < startTime {
		return errors.New("endTime must be >= startTime")
	}
	if duration < 0 {
		return errors.New("duration must be >= 0")
	}
	if maxSize < 0 {
		return errors.New("maxSize must be >= 0")
	}
	if maxNumScore < 0 {
		return errors.New("maxNumScore must be >= 0")
	}

	return TournamentCreate(ctx, n.logger, n.leaderboardCache, n.leaderboardScheduler, id, authoritative, sort, oper, resetSchedule, metadataStr, title, description, category, startTime, endTime, duration, maxSize, maxNumScore, joinRequired)
}

func (n *RuntimeCNakamaModule) TournamentDelete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("expects a tournament ID string")
	}

	return TournamentDelete(ctx, n.leaderboardCache, n.leaderboardRankCache, n.leaderboardScheduler, id)
}

func (n *RuntimeCNakamaModule) TournamentAddAttempt(ctx context.Context, id, ownerID string, count int) error {
	if id == "" {
		return errors.New("expects a tournament ID string")
	}

	if ownerID == "" {
		return errors.New("expects a owner ID string")
	} else if _, err := uuid.FromString(ownerID); err != nil {
		return errors.New("expects owner ID to be a valid identifier")
	}

	if count == 0 {
		return errors.New("expects an attempt count number != 0")
	}

	return TournamentAddAttempt(ctx, n.logger, n.db, n.leaderboardCache, id, ownerID, count)
}

func (n *RuntimeCNakamaModule) TournamentJoin(ctx context.Context, id, ownerID, username string) error {
	if id == "" {
		return errors.New("expects a tournament ID string")
	}

	if ownerID == "" {
		return errors.New("expects a owner ID string")
	}
	ownerUID, err := uuid.FromString(ownerID)
	if err != nil {
		return errors.New("expects owner ID to be a valid identifier")
	}

	if username == "" {
		return errors.New("expects a username string")
	}

	return TournamentJoin(ctx, n.logger, n.db, n.leaderboardCache, n.leaderboardRankCache, ownerUID, username, id)
}

func (n *RuntimeCNakamaModule) TournamentsGetId(ctx context.Context, tournamentIDs []string) ([]*api.Tournament, error) {
	if len(tournamentIDs) == 0 {
		return []*api.Tournament{}, nil
	}

	return TournamentsGet(ctx, n.logger, n.db, n.leaderboardCache, tournamentIDs)
}

func (n *RuntimeCNakamaModule) TournamentList(ctx context.Context, categoryStart, categoryEnd, startTime, endTime, limit int, cursor string) (*api.TournamentList, error) {
	if categoryStart < 0 || categoryStart >= 128 {
		return nil, errors.New("categoryStart must be 0-127")
	}
	if categoryEnd < 0 || categoryEnd >= 128 {
		return nil, errors.New("categoryEnd must be 0-127")
	}
	if startTime < 0 {
		return nil, errors.New("startTime must be >= 0")
	}
	if endTime < 0 {
		return nil, errors.New("endTime must be >= 0")
	}
	if endTime < startTime {
		return nil, errors.New("endTime must be >= startTime")
	}

	if limit < 1 || limit > 100 {
		return nil, errors.New("limit must be 1-100")
	}

	var cursorPtr *TournamentListCursor
	if cursor != "" {
		cb, err := base64.StdEncoding.DecodeString(cursor)
		if err != nil {
			return nil, errors.New("expects cursor to be valid when provided")
		}
		cursorPtr = &TournamentListCursor{}
		if err := gob.NewDecoder(bytes.NewReader(cb)).Decode(cursorPtr); err != nil {
			return nil, errors.New("expects cursor to be valid when provided")
		}
	}

	return TournamentList(ctx, n.logger, n.db, n.leaderboardCache, categoryStart, categoryEnd, startTime, endTime, limit, cursorPtr)
}

func (n *RuntimeCNakamaModule) TournamentRecordsList(ctx context.Context, tournamentId string, ownerIDs []string, limit int, cursor string, overrideExpiry int64) ([]*api.LeaderboardRecord, []*api.LeaderboardRecord, string, string, error) {
	if tournamentId == "" {
		return nil, nil, "", "", errors.New("expects a tournament ID strings")
	}
	for _, ownerID := range ownerIDs {
		if _, err := uuid.FromString(ownerID); err != nil {
			return nil, nil, "", "", errors.New("One or more ownerIDs are invalid.")
		}
	}
	var limitWrapper *wrappers.Int32Value
	if limit < 0 || limit > 10000 {
		return nil, nil, "", "", errors.New("expects limit to be 0-10000")
	}
	limitWrapper = &wrappers.Int32Value{Value: int32(limit)}

	if overrideExpiry < 0 {
		return nil, nil, "", "", errors.New("expects expiry to equal or greater than 0")
	}

	records, err := TournamentRecordsList(ctx, n.logger, n.db, n.leaderboardCache, n.leaderboardRankCache, tournamentId, ownerIDs, limitWrapper, cursor, overrideExpiry)
	if err != nil {
		return nil, nil, "", "", err
	}

	return records.Records, records.OwnerRecords, records.PrevCursor, records.NextCursor, nil
}

func (n *RuntimeCNakamaModule) TournamentRecordWrite(ctx context.Context, id, ownerID, username string, score, subscore int64, metadata map[string]interface{}, overrideOperator *int) (*api.LeaderboardRecord, error) {
	if id == "" {
		return nil, errors.New("expects a tournament ID string")
	}

	owner, err := uuid.FromString(ownerID)
	if err != nil {
		return nil, errors.New("expects owner ID to be a valid identifier")
	}

	// Username is optional.

	metadataStr := "{}"
	if metadata != nil {
		metadataBytes, err := json.Marshal(metadata)
		if err != nil {
			return nil, errors.Errorf("error encoding metadata: %v", err.Error())
		}
		metadataStr = string(metadataBytes)
	}

	operator := api.Operator_NO_OVERRIDE
	if overrideOperator != nil {
		if _, ok := api.Operator_name[int32(*overrideOperator)]; !ok {
			return nil, ErrInvalidOperator
		}
		operator = api.Operator(*overrideOperator)
	}

	return TournamentRecordWrite(ctx, n.logger, n.db, n.leaderboardCache, n.leaderboardRankCache, uuid.Nil, id, owner, username, score, subscore, metadataStr, operator)
}

func (n *RuntimeCNakamaModule) TournamentRecordDelete(ctx context.Context, id, ownerID string) error {
	if id == "" {
		return errors.New("expects a tournament ID string")
	}

	if _, err := uuid.FromString(ownerID); err != nil {
		return errors.New("expects owner ID to be a valid identifier")
	}

	return TournamentRecordDelete(ctx, n.logger, n.db, n.leaderboardCache, n.leaderboardRankCache, uuid.Nil, id, ownerID)
}

func (n *RuntimeCNakamaModule) TournamentRecordsHaystack(ctx context.Context, id, ownerID string, limit int, cursor string, expiry int64) (*api.TournamentRecordList, error) {
	if id == "" {
		return nil, errors.New("expects a tournament ID string")
	}

	owner, err := uuid.FromString(ownerID)
	if err != nil {
		return nil, errors.New("expects owner ID to be a valid identifier")
	}

	if limit < 1 || limit > 100 {
		return nil, errors.New("limit must be 1-100")
	}

	if cursor == "" {
		return nil, errors.New("expects a cursor string")
	}

	if expiry < 0 {
		return nil, errors.New("expiry should be time since epoch in seconds and has to be a positive integer")
	}

	return TournamentRecordsHaystack(ctx, n.logger, n.db, n.leaderboardCache, n.leaderboardRankCache, id, cursor, owner, limit, expiry)
}

func (n *RuntimeCNakamaModule) PurchaseValidateApple(ctx context.Context, userID, receipt string, persist bool, passwordOverride ...string) (*api.ValidatePurchaseResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (n *RuntimeCNakamaModule) PurchaseValidateGoogle(ctx context.Context, userID, receipt string, persist bool, overrides ...struct {
	ClientEmail string
	PrivateKey  string
}) (*api.ValidatePurchaseResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (n *RuntimeCNakamaModule) PurchaseValidateHuawei(ctx context.Context, userID, signature, inAppPurchaseData string, persist bool) (*api.ValidatePurchaseResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (n *RuntimeCNakamaModule) PurchasesList(ctx context.Context, userID string, limit int, cursor string) (*api.PurchaseList, error) {
	// TODO: Implement this
	return nil, nil
}

func (n *RuntimeCNakamaModule) PurchaseGetByTransactionId(ctx context.Context, transactionID string) (*api.ValidatedPurchase, error) {
	// TODO: Implement this
	return nil, nil
}

func (n *RuntimeCNakamaModule) SubscriptionValidateApple(ctx context.Context, userID, receipt string, persist bool, passwordOverride ...string) (*api.ValidateSubscriptionResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (n *RuntimeCNakamaModule) SubscriptionValidateGoogle(ctx context.Context, userID, receipt string, persist bool, overrides ...struct {
	ClientEmail string
	PrivateKey  string
}) (*api.ValidateSubscriptionResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (n *RuntimeCNakamaModule) SubscriptionsList(ctx context.Context, userID string, limit int, cursor string) (*api.SubscriptionList, error) {
	// TODO: Implement this
	return nil, nil
}

func (n *RuntimeCNakamaModule) SubscriptionGetByProductId(ctx context.Context, userID, productID string) (*api.ValidatedSubscription, error) {
	// TODO: Implement this
	return nil, nil
}

func (n *RuntimeCNakamaModule) GroupsGetId(ctx context.Context, groupIDs []string) ([]*api.Group, error) {
	if len(groupIDs) == 0 {
		return make([]*api.Group, 0), nil
	}

	for _, id := range groupIDs {
		if _, err := uuid.FromString(id); err != nil {
			return nil, errors.New("each group id must be a valid id string")
		}
	}

	groups, err := GetGroups(ctx, n.logger, n.db, groupIDs)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (n *RuntimeCNakamaModule) GroupCreate(ctx context.Context, userID, name, creatorID, langTag, description, avatarUrl string, open bool, metadata map[string]interface{}, maxCount int) (*api.Group, error) {
	uid, err := uuid.FromString(userID)
	if err != nil {
		return nil, errors.New("expects user ID to be a valid identifier")
	}

	if name == "" {
		return nil, errors.New("expects group name not be empty")
	}

	cid := uuid.Nil
	if creatorID != "" {
		cid, err = uuid.FromString(creatorID)
		if err != nil {
			return nil, errors.New("expects creator ID to be a valid identifier")
		}
	}

	metadataStr := ""
	if metadata != nil {
		metadataBytes, err := json.Marshal(metadata)
		if err != nil {
			return nil, errors.Errorf("error encoding metadata: %v", err.Error())
		}
		metadataStr = string(metadataBytes)
	}

	if maxCount < 1 {
		return nil, errors.New("expects max_count to be >= 1")
	}

	return CreateGroup(ctx, n.logger, n.db, uid, cid, name, langTag, description, avatarUrl, metadataStr, open, maxCount)
}

func (n *RuntimeCNakamaModule) GroupUpdate(ctx context.Context, id, userID, name, creatorID, langTag, description, avatarUrl string, open bool, metadata map[string]interface{}, maxCount int) error {
	// func (n *RuntimeGoNakamaModule) GroupUpdate(ctx context.Context, id, name, creatorID, langTag, description, avatarUrl string, open bool, metadata map[string]interface{}, maxCount int) error {
	groupID, err := uuid.FromString(id)
	if err != nil {
		return errors.New("expects group ID to be a valid identifier")
	}

	var nameWrapper *wrappers.StringValue
	if name != "" {
		nameWrapper = &wrappers.StringValue{Value: name}
	}

	creator := uuid.Nil
	if creatorID != "" {
		var err error
		creator, err = uuid.FromString(creatorID)
		if err != nil {
			return errors.New("expects creator ID to be a valid identifier")
		}
	}

	var langTagWrapper *wrappers.StringValue
	if langTag != "" {
		langTagWrapper = &wrappers.StringValue{Value: langTag}
	}

	var descriptionWrapper *wrappers.StringValue
	if description != "" {
		descriptionWrapper = &wrappers.StringValue{Value: description}
	}

	var avatarURLWrapper *wrappers.StringValue
	if avatarUrl != "" {
		avatarURLWrapper = &wrappers.StringValue{Value: avatarUrl}
	}

	openWrapper := &wrappers.BoolValue{Value: open}

	var metadataWrapper *wrappers.StringValue
	if metadata != nil {
		metadataBytes, err := json.Marshal(metadata)
		if err != nil {
			return errors.Errorf("error encoding metadata: %v", err.Error())
		}
		metadataWrapper = &wrappers.StringValue{Value: string(metadataBytes)}
	}

	maxCountValue := 0
	if maxCount > 0 && maxCount <= 100 {
		maxCountValue = maxCount
	}

	return UpdateGroup(ctx, n.logger, n.db, groupID, uuid.Nil, creator, nameWrapper, langTagWrapper, descriptionWrapper, avatarURLWrapper, metadataWrapper, openWrapper, maxCountValue)
}

func (n *RuntimeCNakamaModule) GroupDelete(ctx context.Context, id string) error {
	groupID, err := uuid.FromString(id)
	if err != nil {
		return errors.New("expects group ID to be a valid identifier")
	}

	return DeleteGroup(ctx, n.logger, n.db, groupID, uuid.Nil)
}

func (n *RuntimeCNakamaModule) GroupUserJoin(ctx context.Context, groupID, userID, username string) error {
	group, err := uuid.FromString(groupID)
	if err != nil {
		return errors.New("expects group ID to be a valid identifier")
	}

	user, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("expects user ID to be a valid identifier")
	}

	if username == "" {
		return errors.New("expects a username string")
	}

	return JoinGroup(ctx, n.logger, n.db, n.router, group, user, username)
}

func (n *RuntimeCNakamaModule) GroupUserLeave(ctx context.Context, groupID, userID, username string) error {
	group, err := uuid.FromString(groupID)
	if err != nil {
		return errors.New("expects group ID to be a valid identifier")
	}

	user, err := uuid.FromString(userID)
	if err != nil {
		return errors.New("expects user ID to be a valid identifier")
	}

	if username == "" {
		return errors.New("expects a username string")
	}

	return LeaveGroup(ctx, n.logger, n.db, n.tracker, n.router, n.streamManager, group, user, username)
}

// TODO: plumb through callerID if needed
func (n *RuntimeCNakamaModule) GroupUsersAdd(ctx context.Context, callerID, groupID string, userIDs []string) error {
	group, err := uuid.FromString(groupID)
	if err != nil {
		return errors.New("expects group ID to be a valid identifier")
	}

	if len(userIDs) == 0 {
		return nil
	}

	users := make([]uuid.UUID, 0, len(userIDs))
	for _, userID := range userIDs {
		uid, err := uuid.FromString(userID)
		if err != nil {
			return errors.New("expects each user ID to be a valid identifier")
		}
		if uid == uuid.Nil {
			return errors.New("cannot add the root user")
		}
		users = append(users, uid)
	}

	return AddGroupUsers(ctx, n.logger, n.db, n.router, uuid.Nil, group, users)
}

func (n *RuntimeCNakamaModule) GroupUsersBan(ctx context.Context, callerID, groupID string, userIDs []string) error {
	// TODO: Implement this
	return nil
}

// TODO: plumb through callerID if needed
func (n *RuntimeCNakamaModule) GroupUsersKick(ctx context.Context, callerID, groupID string, userIDs []string) error {
	group, err := uuid.FromString(groupID)
	if err != nil {
		return errors.New("expects group ID to be a valid identifier")
	}

	if len(userIDs) == 0 {
		return nil
	}

	users := make([]uuid.UUID, 0, len(userIDs))
	for _, userID := range userIDs {
		uid, err := uuid.FromString(userID)
		if err != nil {
			return errors.New("expects each user ID to be a valid identifier")
		}
		if uid == uuid.Nil {
			return errors.New("cannot kick the root user")
		}
		users = append(users, uid)
	}

	return KickGroupUsers(ctx, n.logger, n.db, n.tracker, n.router, n.streamManager, uuid.Nil, group, users, false)
}

// TODO: plumb through callerID if needed
func (n *RuntimeCNakamaModule) GroupUsersPromote(ctx context.Context, callerID, groupID string, userIDs []string) error {
	group, err := uuid.FromString(groupID)
	if err != nil {
		return errors.New("expects group ID to be a valid identifier")
	}

	if len(userIDs) == 0 {
		return nil
	}

	users := make([]uuid.UUID, 0, len(userIDs))
	for _, userID := range userIDs {
		uid, err := uuid.FromString(userID)
		if err != nil {
			return errors.New("expects each user ID to be a valid identifier")
		}
		if uid == uuid.Nil {
			return errors.New("cannot promote the root user")
		}
		users = append(users, uid)
	}

	return PromoteGroupUsers(ctx, n.logger, n.db, n.router, uuid.Nil, group, users)
}

// TODO: plumb through callerID if needed
func (n *RuntimeCNakamaModule) GroupUsersDemote(ctx context.Context, callerID, groupID string, userIDs []string) error {
	group, err := uuid.FromString(groupID)
	if err != nil {
		return errors.New("expects group ID to be a valid identifier")
	}

	if len(userIDs) == 0 {
		return nil
	}

	users := make([]uuid.UUID, 0, len(userIDs))
	for _, userID := range userIDs {
		uid, err := uuid.FromString(userID)
		if err != nil {
			return errors.New("expects each user ID to be a valid identifier")
		}
		if uid == uuid.Nil {
			return errors.New("cannot demote the root user")
		}
		users = append(users, uid)
	}

	return DemoteGroupUsers(ctx, n.logger, n.db, n.router, uuid.Nil, group, users)
}

func (n *RuntimeCNakamaModule) GroupUsersList(ctx context.Context, id string, limit int, state *int, cursor string) ([]*api.GroupUserList_GroupUser, string, error) {
	groupID, err := uuid.FromString(id)
	if err != nil {
		return nil, "", errors.New("expects group ID to be a valid identifier")
	}

	if limit < 1 || limit > 100 {
		return nil, "", errors.New("expects limit to be 1-100")
	}

	var stateWrapper *wrappers.Int32Value
	if state != nil {
		stateValue := *state
		if stateValue < 0 || stateValue > 4 {
			return nil, "", errors.New("expects state to be 0-4")
		}
		stateWrapper = &wrappers.Int32Value{Value: int32(stateValue)}
	}

	users, err := ListGroupUsers(ctx, n.logger, n.db, n.statusRegistry, groupID, limit, stateWrapper, cursor)
	if err != nil {
		return nil, "", err
	}

	return users.GroupUsers, users.Cursor, nil
}

func (n *RuntimeCNakamaModule) GroupsList(ctx context.Context, name, langTag string, members *int, open *bool, limit int, cursor string) ([]*api.Group, string, error) {
	// TODO: Implement this
	return nil, "", nil
}

func (n *RuntimeCNakamaModule) GroupsGetRandom(ctx context.Context, count int) ([]*api.Group, error) {
	// TODO: Implement this
	return nil, nil
}

func (n *RuntimeCNakamaModule) UserGroupsList(ctx context.Context, userID string, limit int, state *int, cursor string) ([]*api.UserGroupList_UserGroup, string, error) {
	uid, err := uuid.FromString(userID)
	if err != nil {
		return nil, "", errors.New("expects user ID to be a valid identifier")
	}

	if limit < 1 || limit > 100 {
		return nil, "", errors.New("expects limit to be 1-100")
	}

	var stateWrapper *wrappers.Int32Value
	if state != nil {
		stateValue := *state
		if stateValue < 0 || stateValue > 4 {
			return nil, "", errors.New("expects state to be 0-4")
		}
		stateWrapper = &wrappers.Int32Value{Value: int32(stateValue)}
	}

	groups, err := ListUserGroups(ctx, n.logger, n.db, uid, limit, stateWrapper, cursor)
	if err != nil {
		return nil, "", err
	}

	return groups.UserGroups, groups.Cursor, nil
}

func (n *RuntimeCNakamaModule) Event(ctx context.Context, evt *api.Event) error {
	if ctx == nil {
		return errors.New("expects a non-nil context")
	}
	if evt == nil {
		return errors.New("expects a non-nil event")
	}

	n.RLock()
	fn := n.eventFn
	n.RUnlock()
	if fn != nil {
		fn(ctx, evt)
	}

	return nil
}

func (n *RuntimeCNakamaModule) MetricsCounterAdd(name string, tags map[string]string, delta int64) {
	n.metrics.CustomCounter(name, tags, delta)
}

func (n *RuntimeCNakamaModule) MetricsGaugeSet(name string, tags map[string]string, value float64) {
	n.metrics.CustomGauge(name, tags, value)
}

func (n *RuntimeCNakamaModule) MetricsTimerRecord(name string, tags map[string]string, value time.Duration) {
	n.metrics.CustomTimer(name, tags, value)
}

func (n *RuntimeCNakamaModule) FriendsList(ctx context.Context, userID string, limit int, state *int, cursor string) ([]*api.Friend, string, error) {
	uid, err := uuid.FromString(userID)
	if err != nil {
		return nil, "", errors.New("expects user ID to be a valid identifier")
	}

	if limit < 1 || limit > 100 {
		return nil, "", errors.New("expects limit to be 1-100")
	}

	var stateWrapper *wrappers.Int32Value
	if state != nil {
		stateValue := *state
		if stateValue < 0 || stateValue > 3 {
			return nil, "", errors.New("expects state to be 0-3")
		}
		stateWrapper = &wrappers.Int32Value{Value: int32(stateValue)}
	}

	friends, err := ListFriends(ctx, n.logger, n.db, n.statusRegistry, uid, limit, stateWrapper, cursor)
	if err != nil {
		return nil, "", err
	}

	return friends.Friends, friends.Cursor, nil
}

func (n *RuntimeCNakamaModule) FriendsAdd(ctx context.Context, userID string, username string, ids []string, usernames []string) error {
	// TODO: Implement this
	return nil
}

func (n *RuntimeCNakamaModule) FriendsDelete(ctx context.Context, userID string, username string, ids []string, usernames []string) error {
	// TODO: Implement this
	return nil
}

func (n *RuntimeCNakamaModule) FriendsBlock(ctx context.Context, userID string, username string, ids []string, usernames []string) error {
	// TODO: Implement this
	return nil
}

func (n *RuntimeCNakamaModule) SetEventFn(fn RuntimeEventCustomFunction) {
	n.Lock()
	n.eventFn = fn
	n.Unlock()
}

func (n *RuntimeCNakamaModule) ChannelMessageSend(ctx context.Context, channelId string, content map[string]interface{}, senderId, senderUsername string, persist bool) (*rtapi.ChannelMessageAck, error) {
	// TODO: Implement this
	return nil, nil
}

func (n *RuntimeCNakamaModule) ChannelMessageUpdate(ctx context.Context, channelId, messageId string, content map[string]interface{}, senderId, senderUsername string, persist bool) (*rtapi.ChannelMessageAck, error) {
	// TODO: Implement this
	return nil, nil
}

func (n *RuntimeCNakamaModule) ChannelMessageRemove(ctx context.Context, channelId, messageId string, senderId, senderUsername string, persist bool) (*rtapi.ChannelMessageAck, error) {
	// TODO: Implement this
	return nil, nil
}

func (n *RuntimeCNakamaModule) ChannelMessagesList(ctx context.Context, channelId string, limit int, forward bool, cursor string) ([]*api.ChannelMessage, string, string, error) {
	// TODO: Implement this
	return nil, "", "", nil
}

func (n *RuntimeCNakamaModule) ChannelIdBuild(ctx context.Context, senderId, target string, chanType runtime.ChannelType) (string, error) {
	// TODO: Implement this
	return "", nil
}

// @group satori
// @summary Get the Satori client.
// @return satori(runtime.Satori) The Satori client.
func (n *RuntimeCNakamaModule) GetSatori() runtime.Satori {
	return n.satori
}
