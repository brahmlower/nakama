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

extern int initmodule(const void *, NkContext, NkLogger, NkDb, NkModule, NkInitializer);

extern NkString contextvalue(const void *, NkString key);

extern void loggerdebug(const void *, NkString);
extern void loggererror(const void *, NkString);
extern void loggerinfo(const void *, NkString);
extern void loggerwarn(const void *, NkString);

extern int moduleauthenticateapple(const void *, NkContext, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int moduleauthenticatecustom(const void *, NkContext, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int moduleauthenticatedevice(const void *, NkContext, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int moduleauthenticateemail(const void *, NkContext, NkString, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int moduleauthenticatefacebook(const void *, NkContext, NkString, bool, NkString, bool, NkString *, NkString *, NkString *, bool *l);
extern int moduleauthenticatefacebookinstantgame(const void *, NkContext, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int moduleauthenticategamecenter(const void *, NkContext, NkString, NkString, NkI64, NkString, NkString, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int moduleauthenticategoogle(const void *, NkContext, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int moduleauthenticatesteam(const void *, NkContext, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);

*/
import "C"

import (
	"context"
	"database/sql"
	"errors"

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
	ret.registerrpc = C.NkInitializerRpcFn(C.initializerregisterrpc);
	ret.registerbeforert = C.NkInitializerBeforeRtFn(C.initializerregisterbeforert);
	ret.registerafterrt = C.NkInitializerAfterRtFn(C.initializerregisterafterrt);
	ret.registermatchmakermatched = C.NkInitializerMatchmakerMatchedFn(C.initializerregistermatchmakermatched);
	ret.registermatch = C.NkInitializerMatchFn(C.initializerregistermatch);
	ret.registertournamentend = C.NkInitializerTournamentFn(C.initializerregistertournamentend);
	ret.registertournamentreset = C.NkInitializerTournamentFn(C.initializerregistertournamentreset);
	ret.registerleaderboardreset = C.NkInitializerLeaderBoardFn(C.initializerregisterleaderboardreset);
	ret.registerbeforegetaccount = C.NkInitializerBeforeGetAccountFn(C.initializerregisterbeforegetaccount);
	ret.registeraftergetaccount = C.NkInitializerAfterGetAccountFn(C.initializerregisteraftergetaccount);
	ret.registerbeforeupdateaccount = C.NkInitializerBeforeUpdateAccountFn(C.initializerregisterbeforeupdateaccount);
	ret.registerafterupdateaccount = C.NkInitializerAfterUpdateAccountFn(C.initializerregisterafterupdateaccount);
	ret.registerbeforesessionrefresh = C.NkInitializerBeforeSessionRefreshFn(C.initializerregisterbeforesessionrefresh);
	ret.registeraftersessionrefresh = C.NkInitializerAfterSessionRefreshFn(C.initializerregisteraftersessionrefresh);
	ret.registerbeforeauthenticateapple = C.NkInitializerBeforeAuthenticateAppleFn(C.initializerregisterbeforeauthenticateapple);
	ret.registerafterauthenticateapple = C.NkInitializerAfterAuthenticateAppleFn(C.initializerregisterafterauthenticateapple);
	ret.registerbeforeauthenticatecustom = C.NkInitializerBeforeAuthenticateCustomFn(C.initializerregisterbeforeauthenticatecustom);
	ret.registerafterauthenticatecustom = C.NkInitializerAfterAuthenticateCustomFn(C.initializerregisterafterauthenticatecustom);
	ret.registerbeforeauthenticatedevice = C.NkInitializerBeforeAuthenticateDeviceFn(C.initializerregisterbeforeauthenticatedevice);
	ret.registerafterauthenticatedevice = C.NkInitializerAfterAuthenticateDeviceFn(C.initializerregisterafterauthenticatedevice);
	ret.registerbeforeauthenticateemail = C.NkInitializerBeforeAuthenticateEmailFn(C.initializerregisterbeforeauthenticateemail);
	ret.registerafterauthenticateemail = C.NkInitializerAfterAuthenticateEmailFn(C.initializerregisterafterauthenticateemail);
	ret.registerbeforeauthenticatefacebook = C.NkInitializerBeforeAuthenticateFacebookFn(C.initializerregisterbeforeauthenticatefacebook);
	ret.registerafterauthenticatefacebook = C.NkInitializerAfterAuthenticateFacebookFn(C.initializerregisterafterauthenticatefacebook);
	ret.registerbeforeauthenticatefacebookinstantgame = C.NkInitializerBeforeAuthenticateFacebookInstantGameFn(C.initializerregisterbeforeauthenticatefacebookinstantgame);
	ret.registerafterauthenticatefacebookinstantgame = C.NkInitializerAfterAuthenticateFacebookInstantGameFn(C.initializerregisterafterauthenticatefacebookinstantgame);
	ret.registerbeforeauthenticategamecenter = C.NkInitializerBeforeAuthenticateGameCenterFn(C.initializerregisterbeforeauthenticategamecenter);
	ret.registerafterauthenticategamecenter = C.NkInitializerAfterAuthenticateGameCenterFn(C.initializerregisterafterauthenticategamecenter);
	ret.registerbeforeauthenticategoogle = C.NkInitializerBeforeAuthenticateGoogleFn(C.initializerregisterbeforeauthenticategoogle);
	ret.registerafterauthenticategoogle = C.NkInitializerAfterAuthenticateGoogleFn(C.initializerregisterafterauthenticategoogle);
	ret.registerbeforeauthenticatesteam = C.NkInitializerBeforeAuthenticateSteamFn(C.initializerregisterbeforeauthenticatesteam);
	ret.registerafterauthenticatesteam = C.NkInitializerAfterAuthenticateSteamFn(C.initializerregisterafterauthenticatesteam);
	ret.registerbeforelistchannelmessages = C.NkInitializerBeforeListChannelMessagesFn(C.initializerregisterbeforelistchannelmessages);
	ret.registerafterlistchannelmessages = C.NkInitializerAfterListChannelMessagesFn(C.initializerregisterafterlistchannelmessages);
	ret.registerbeforelistfriends = C.NkInitializerBeforeListFriendsFn(C.initializerregisterbeforelistfriends);
	ret.registerafterlistfriends = C.NkInitializerAfterListFriendsFn(C.initializerregisterafterlistfriends);
	ret.registerbeforeaddfriends = C.NkInitializerBeforeAddFriendsFn(C.initializerregisterbeforeaddfriends);
	ret.registerafteraddfriends = C.NkInitializerAfterAddFriendsFn(C.initializerregisterafteraddfriends);
	ret.registerbeforedeletefriends = C.NkInitializerBeforeDeleteFriendsFn(C.initializerregisterbeforedeletefriends);
	ret.registerafterdeletefriends = C.NkInitializerAfterDeleteFriendsFn(C.initializerregisterafterdeletefriends);
	ret.registerbeforeblockfriends = C.NkInitializerBeforeBlockFriendsFn(C.initializerregisterbeforeblockfriends);
	ret.registerafterblockfriends = C.NkInitializerAfterBlockFriendsFn(C.initializerregisterafterblockfriends);
	ret.registerbeforeimportfacebookfriends = C.NkInitializerBeforeImportFacebookFriendsFn(C.initializerregisterbeforeimportfacebookfriends);
	ret.registerafterimportfacebookfriends = C.NkInitializerAfterImportFacebookFriendsFn(C.initializerregisterafterimportfacebookfriends);
	ret.registerbeforecreategroup = C.NkInitializerBeforeCreateGroupFn(C.initializerregisterbeforecreategroup);
	ret.registeraftercreategroup = C.NkInitializerAfterCreateGroupFn(C.initializerregisteraftercreategroup);
	ret.registerbeforeupdategroup = C.NkInitializerBeforeUpdateGroupFn(C.initializerregisterbeforeupdategroup);
	ret.registerafterupdategroup = C.NkInitializerAfterUpdateGroupFn(C.initializerregisterafterupdategroup);
	ret.registerbeforedeletegroup = C.NkInitializerBeforeDeleteGroupFn(C.initializerregisterbeforedeletegroup);
	ret.registerafterdeletegroup = C.NkInitializerAfterDeleteGroupFn(C.initializerregisterafterdeletegroup);
	ret.registerbeforejoingroup = C.NkInitializerBeforeJoinGroupFn(C.initializerregisterbeforejoingroup);
	ret.registerafterjoingroup = C.NkInitializerAfterJoinGroupFn(C.initializerregisterafterjoingroup);
	ret.registerbeforeleavegroup = C.NkInitializerBeforeLeaveGroupFn(C.initializerregisterbeforeleavegroup);
	ret.registerafterleavegroup = C.NkInitializerAfterLeaveGroupFn(C.initializerregisterafterleavegroup);
	ret.registerbeforeaddgroupusers = C.NkInitializerBeforeAddGroupUsersFn(C.initializerregisterbeforeaddgroupusers);
	ret.registerafteraddgroupusers = C.NkInitializerAfterAddGroupUsersFn(C.initializerregisterafteraddgroupusers);
	ret.registerbeforebangroupusers = C.NkInitializerBeforeBanGroupUsersFn(C.initializerregisterbeforebangroupusers);
	ret.registerafterbangroupusers = C.NkInitializerAfterBanGroupUsersFn(C.initializerregisterafterbangroupusers);
	ret.registerbeforekickgroupusers = C.NkInitializerBeforeKickGroupUsersFn(C.initializerregisterbeforekickgroupusers);
	ret.registerafterkickgroupusers = C.NkInitializerAfterKickGroupUsersFn(C.initializerregisterafterkickgroupusers);
	ret.registerbeforepromotegroupusers = C.NkInitializerBeforePromoteGroupUsersFn(C.initializerregisterbeforepromotegroupusers);
	ret.registerafterpromotegroupusers = C.NkInitializerAfterPromoteGroupUsersFn(C.initializerregisterafterpromotegroupusers);
	ret.registerbeforedemotegroupusers = C.NkInitializerBeforeDemoteGroupUsersFn(C.initializerregisterbeforedemotegroupusers);
	ret.registerafterdemotegroupusers = C.NkInitializerAfterDemoteGroupUsersFn(C.initializerregisterafterdemotegroupusers);
	ret.registerbeforelistgroupusers = C.NkInitializerBeforeListGroupUsersFn(C.initializerregisterbeforelistgroupusers);
	ret.registerafterlistgroupusers = C.NkInitializerAfterListGroupUsersFn(C.initializerregisterafterlistgroupusers);
	ret.registerbeforelistusergroups = C.NkInitializerBeforeListUserGroupsFn(C.initializerregisterbeforelistusergroups);
	ret.registerafterlistusergroups = C.NkInitializerAfterListUserGroupsFn(C.initializerregisterafterlistusergroups);
	ret.registerbeforelistgroups = C.NkInitializerBeforeListGroupsFn(C.initializerregisterbeforelistgroups);
	ret.registerafterlistgroups = C.NkInitializerAfterListGroupsFn(C.initializerregisterafterlistgroups);
	ret.registerbeforedeleteleaderboardrecord = C.NkInitializerBeforeDeleteLeaderboardRecordFn(C.initializerregisterbeforedeleteleaderboardrecord);
	ret.registerafterdeleteleaderboardrecord = C.NkInitializerAfterDeleteLeaderboardRecordFn(C.initializerregisterafterdeleteleaderboardrecord);
	ret.registerbeforelistleaderboardrecords = C.NkInitializerBeforeListLeaderboardRecordsFn(C.initializerregisterbeforelistleaderboardrecords);
	ret.registerafterlistleaderboardrecords = C.NkInitializerAfterListLeaderboardRecordsFn(C.initializerregisterafterlistleaderboardrecords);
	ret.registerbeforewriteleaderboardrecord = C.NkInitializerBeforeWriteLeaderboardRecordFn(C.initializerregisterbeforewriteleaderboardrecord);
	ret.registerafterwriteleaderboardrecord = C.NkInitializerAfterWriteLeaderboardRecordFn(C.initializerregisterafterwriteleaderboardrecord);
	ret.registerbeforelistleaderboardrecordsaroundowner = C.NkInitializerBeforeListLeaderboardRecordsAroundOwnerFn(C.initializerregisterbeforelistleaderboardrecordsaroundowner);
	ret.registerafterlistleaderboardrecordsaroundowner = C.NkInitializerAfterListLeaderboardRecordsAroundOwnerFn(C.initializerregisterafterlistleaderboardrecordsaroundowner);
	ret.registerbeforelinkapple = C.NkInitializerBeforeLinkAppleFn(C.initializerregisterbeforelinkapple);
	ret.registerafterlinkapple = C.NkInitializerAfterLinkAppleFn(C.initializerregisterafterlinkapple);
	ret.registerbeforelinkcustom = C.NkInitializerBeforeLinkCustomFn(C.initializerregisterbeforelinkcustom);
	ret.registerafterlinkcustom = C.NkInitializerAfterLinkCustomFn(C.initializerregisterafterlinkcustom);
	ret.registerbeforelinkdevice = C.NkInitializerBeforeLinkDeviceFn(C.initializerregisterbeforelinkdevice);
	ret.registerafterlinkdevice = C.NkInitializerAfterLinkDeviceFn(C.initializerregisterafterlinkdevice);
	ret.registerbeforelinkemail = C.NkInitializerBeforeLinkEmailFn(C.initializerregisterbeforelinkemail);
	ret.registerafterlinkemail = C.NkInitializerAfterLinkEmailFn(C.initializerregisterafterlinkemail);
	ret.registerbeforelinkfacebook = C.NkInitializerBeforeLinkFacebookFn(C.initializerregisterbeforelinkfacebook);
	ret.registerafterlinkfacebook = C.NkInitializerAfterLinkFacebookFn(C.initializerregisterafterlinkfacebook);
	ret.registerbeforelinkfacebookinstantgame = C.NkInitializerBeforeLinkFacebookInstantGameFn(C.initializerregisterbeforelinkfacebookinstantgame);
	ret.registerafterlinkfacebookinstantgame = C.NkInitializerAfterLinkFacebookInstantGameFn(C.initializerregisterafterlinkfacebookinstantgame);
	ret.registerbeforelinkgamecenter = C.NkInitializerBeforeLinkGameCenterFn(C.initializerregisterbeforelinkgamecenter);
	ret.registerafterlinkgamecenter = C.NkInitializerAfterLinkGameCenterFn(C.initializerregisterafterlinkgamecenter);
	ret.registerbeforelinkgoogle = C.NkInitializerBeforeLinkGoogleFn(C.initializerregisterbeforelinkgoogle);
	ret.registerafterlinkgoogle = C.NkInitializerAfterLinkGoogleFn(C.initializerregisterafterlinkgoogle);
	ret.registerbeforelinksteam = C.NkInitializerBeforeLinkSteamFn(C.initializerregisterbeforelinksteam);
	ret.registerafterlinksteam = C.NkInitializerAfterLinkSteamFn(C.initializerregisterafterlinksteam);
	ret.registerbeforelistmatches = C.NkInitializerBeforeListMatchesFn(C.initializerregisterbeforelistmatches);
	ret.registerafterlistmatches = C.NkInitializerAfterListMatchesFn(C.initializerregisterafterlistmatches);
	ret.registerbeforelistnotifications = C.NkInitializerBeforeListNotificationsFn(C.initializerregisterbeforelistnotifications);
	ret.registerafterlistnotifications = C.NkInitializerAfterListNotificationsFn(C.initializerregisterafterlistnotifications);
	ret.registerbeforedeletenotifications = C.NkInitializerBeforeDeleteNotificationsFn(C.initializerregisterbeforedeletenotifications);
	ret.registerafterdeletenotifications = C.NkInitializerAfterDeleteNotificationsFn(C.initializerregisterafterdeletenotifications);
	ret.registerbeforeliststorageobjects = C.NkInitializerBeforeListStorageObjectsFn(C.initializerregisterbeforeliststorageobjects);
	ret.registerafterliststorageobjects = C.NkInitializerAfterListStorageObjectsFn(C.initializerregisterafterliststorageobjects);
	ret.registerbeforereadstorageobjects = C.NkInitializerBeforeReadStorageObjectsFn(C.initializerregisterbeforereadstorageobjects);
	ret.registerafterreadstorageobjects = C.NkInitializerAfterReadStorageObjectsFn(C.initializerregisterafterreadstorageobjects);
	ret.registerbeforewritestorageobjects = C.NkInitializerBeforeWriteStorageObjectsFn(C.initializerregisterbeforewritestorageobjects);
	ret.registerafterwritestorageobjects = C.NkInitializerAfterWriteStorageObjectsFn(C.initializerregisterafterwritestorageobjects);
	ret.registerbeforedeletestorageobjects = C.NkInitializerBeforeDeleteStorageObjectsFn(C.initializerregisterbeforedeletestorageobjects);
	ret.registerafterdeletestorageobjects = C.NkInitializerAfterDeleteStorageObjectsFn(C.initializerregisterafterdeletestorageobjects);
	ret.registerbeforejointournament = C.NkInitializerBeforeJoinTournamentFn(C.initializerregisterbeforejointournament);
	ret.registerafterjointournament = C.NkInitializerAfterJoinTournamentFn(C.initializerregisterafterjointournament);
	ret.registerbeforelisttournamentrecords = C.NkInitializerBeforeListTournamentRecordsFn(C.initializerregisterbeforelisttournamentrecords);
	ret.registerafterlisttournamentrecords = C.NkInitializerAfterListTournamentRecordsFn(C.initializerregisterafterlisttournamentrecords);
	ret.registerbeforelisttournaments = C.NkInitializerBeforeListTournamentsFn(C.initializerregisterbeforelisttournaments);
	ret.registerafterlisttournaments = C.NkInitializerAfterListTournamentsFn(C.initializerregisterafterlisttournaments);
	ret.registerbeforewritetournamentrecord = C.NkInitializerBeforeWriteTournamentRecordFn(C.initializerregisterbeforewritetournamentrecord);
	ret.registerafterwritetournamentrecord = C.NkInitializerAfterWriteTournamentRecordFn(C.initializerregisterafterwritetournamentrecord);
	ret.registerbeforelisttournamentrecordsaroundowner = C.NkInitializerBeforeListTournamentRecordsAroundOwnerFn(C.initializerregisterbeforelisttournamentrecordsaroundowner);
	ret.registerafterlisttournamentrecordsaroundowner = C.NkInitializerAfterListTournamentRecordsAroundOwnerFn(C.initializerregisterafterlisttournamentrecordsaroundowner);
	ret.registerbeforeunlinkapple = C.NkInitializerBeforeUnlinkAppleFn(C.initializerregisterbeforeunlinkapple);
	ret.registerafterunlinkapple = C.NkInitializerAfterUnlinkAppleFn(C.initializerregisterafterunlinkapple);
	ret.registerbeforeunlinkcustom = C.NkInitializerBeforeUnlinkCustomFn(C.initializerregisterbeforeunlinkcustom);
	ret.registerafterunlinkcustom = C.NkInitializerAfterUnlinkCustomFn(C.initializerregisterafterunlinkcustom);
	ret.registerbeforeunlinkdevice = C.NkInitializerBeforeUnlinkDeviceFn(C.initializerregisterbeforeunlinkdevice);
	ret.registerafterunlinkdevice = C.NkInitializerAfterUnlinkDeviceFn(C.initializerregisterafterunlinkdevice);
	ret.registerbeforeunlinkemail = C.NkInitializerBeforeUnlinkEmailFn(C.initializerregisterbeforeunlinkemail);
	ret.registerafterunlinkemail = C.NkInitializerAfterUnlinkEmailFn(C.initializerregisterafterunlinkemail);
	ret.registerbeforeunlinkfacebook = C.NkInitializerBeforeUnlinkFacebookFn(C.initializerregisterbeforeunlinkfacebook);
	ret.registerafterunlinkfacebook = C.NkInitializerAfterUnlinkFacebookFn(C.initializerregisterafterunlinkfacebook);
	ret.registerbeforeunlinkfacebookinstantgame = C.NkInitializerBeforeUnlinkFacebookInstantGameFn(C.initializerregisterbeforeunlinkfacebookinstantgame);
	ret.registerafterunlinkfacebookinstantgame = C.NkInitializerAfterUnlinkFacebookInstantGameFn(C.initializerregisterafterunlinkfacebookinstantgame);
	ret.registerbeforeunlinkgamecenter = C.NkInitializerBeforeUnlinkGameCenterFn(C.initializerregisterbeforeunlinkgamecenter);
	ret.registerafterunlinkgamecenter = C.NkInitializerAfterUnlinkGameCenterFn(C.initializerregisterafterunlinkgamecenter);
	ret.registerbeforeunlinkgoogle = C.NkInitializerBeforeUnlinkGoogleFn(C.initializerregisterbeforeunlinkgoogle);
	ret.registerafterunlinkgoogle = C.NkInitializerAfterUnlinkGoogleFn(C.initializerregisterafterunlinkgoogle);
	ret.registerbeforeunlinksteam = C.NkInitializerBeforeUnlinkSteamFn(C.initializerregisterbeforeunlinksteam);
	ret.registerafterunlinksteam = C.NkInitializerAfterUnlinkSteamFn(C.initializerregisterafterunlinksteam);
	ret.registerbeforegetusers = C.NkInitializerBeforeGetUsersFn(C.initializerregisterbeforegetusers);
	ret.registeraftergetusers = C.NkInitializerAfterGetUsersFn(C.initializerregisteraftergetusers);
	ret.registerevent = C.NkInitializerEventFn(C.initializerregisterevent);
	ret.registereventsessionstart = C.NkInitializerEventSessionStartFn(C.initializerregistereventsessionstart);
	ret.registereventsessionend = C.NkInitializerEventSessionEndFn(C.initializerregistereventsessionend);

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
	call := &RuntimeCNakamaModuleCall{}
	call.nk = nk

	ret := C.NkModule{}
	ret.ptr = pointer.Save(call)
	ret.authenticateapple = C.NkModuleAuthenticateFn(C.moduleauthenticateapple)
	ret.authenticatecustom = C.NkModuleAuthenticateFn(C.moduleauthenticatecustom)
	ret.authenticatedevice = C.NkModuleAuthenticateFn(C.moduleauthenticatedevice)
	ret.authenticateemail = C.NkModuleAuthenticateEmailFn(C.moduleauthenticateemail)
	ret.authenticatefacebook = C.NkModuleAuthenticateFacebookFn(C.moduleauthenticatefacebook)
	ret.authenticatefacebookinstantgame = C.NkModuleAuthenticateFn(C.moduleauthenticatefacebookinstantgame)
	ret.authenticategamecenter = C.NkModuleAuthenticateGameCenterFn(C.moduleauthenticategamecenter)
	ret.authenticategoogle = C.NkModuleAuthenticateFn(C.moduleauthenticategoogle)
	ret.authenticatesteam = C.NkModuleAuthenticateFn(C.moduleauthenticatesteam)
	res.authenticatetokengenerate = C.NkModuleAuthenticateTokenGenerateFn(C.moduleauthenticatetokengenerate);
	res.accountgetid = C.NkModuleAccountGetIdFn(C.moduleaccountgetid);
	res.accountsgetid = C.NkModuleAccountsGetIdFn(C.moduleaccountsgetid);
	res.accountupdateid = C.NkModuleAccountUpdateIdFn(C.moduleaccountupdateid);
	res.accountdeleteid = C.NkModuleAccountDeleteIdFn(C.moduleaccountdeleteid);
	res.accountexportid = C.NkModuleAccountExportIdFn(C.moduleaccountexportid);
	res.usersgetid = C.NkModuleUsersGetIdFn(C.moduleusersgetid);
	res.usersgetusername = C.NkModuleUsersGetUsernameFn(C.moduleusersgetusername);
	res.usersbanid = C.NkModuleUsersBanIdFn(C.moduleusersbanid);
	res.usersunbanid = C.NkModuleUsersUnbanIdFn(C.moduleusersunbanid);
	res.linkapple = C.NkModuleLinkAppleFn(C.modulelinkapple);
	res.linkcustom = C.NkModuleLinkCustomFn(C.modulelinkcustom);
	res.linkdevice = C.NkModuleLinkDeviceFn(C.modulelinkdevice);
	res.linkemail = C.NkModuleLinkEmailFn(C.modulelinkemail);
	res.linkfacebook = C.NkModuleLinkFacebookFn(C.modulelinkfacebook);
	res.linkfacebookinstantgame = C.NkModuleLinkFacebookInstantGameFn(C.modulelinkfacebookinstantgame);
	res.linkgamecenter = C.NkModuleLinkGameCenterFn(C.modulelinkgamecenter);
	res.linkgoogle = C.NkModuleLinkGoogleFn(C.modulelinkgoogle);
	res.linksteam = C.NkModuleLinkSteamFn(C.modulelinksteam);
	res.unlinkapple = C.NkModuleUnlinkAppleFn(C.moduleunlinkapple);
	res.unlinkcustom = C.NkModuleUnlinkCustomFn(C.moduleunlinkcustom);
	res.unlinkdevice = C.NkModuleUnlinkDeviceFn(C.moduleunlinkdevice);
	res.unlinkemail = C.NkModuleUnlinkEmailFn(C.moduleunlinkemail);
	res.unlinkfacebook = C.NkModuleUnlinkFacebookFn(C.moduleunlinkfacebook);
	res.unlinkfacebookinstantgame = C.NkModuleUnlinkFacebookInstantGameFn(C.moduleunlinkfacebookinstantgame);
	res.unlinkgamecenter = C.NkModuleUnlinkGameCenterFn(C.moduleunlinkgamecenter);
	res.unlinkgoogle = C.NkModuleUnlinkGoogleFn(C.moduleunlinkgoogle);
	res.unlinksteam = C.NkModuleunlinkSteamFn(C.moduleunlinksteam);
	res.streamuserlist = C.NkModuleStreamUserListFn(C.modulestreamuserlist);
	res.streamuserget = C.NkModuleStreamUserGetFn(C.modulestreamuserget);
	res.streamuserjoin = C.NkModuleStreamUserJoinFn(C.modulestreamuserjoin);
	res.streamuserupdate = C.NkModuleStreamuserUpdateFn(C.modulestreamuserupdate);
	res.streamuserleave = C.NkModuleStreamUserLeaveFn(C.modulestreamuserleave);
	res.streamuserkick = C.NkModuleStreamUserKickFn(C.modulestreamuserkick);
	res.streamcount = C.NkModuleStreamCountFn(C.modulestreamcount);
	res.streamclose = C.NkModuleStreamCloseFn(C.modulestreamclose);
	res.streamsend = C.NkModuleStreamSendFn(C.modulestreamsend);
	res.streamsendraw = C.NkModuleStreamSendRawFn(C.modulestreamsendraw);
	res.sessiondisconnect = C.NkModuleSessionDisconnectFn(C.modulesessiondisconnect);
	res.matchcreate = C.NkModuleMatchCreateFn(C.modulematchcreate);
	res.matchget = C.NkModuleMatchGetFn(C.modulematchget);
	res.matchlist = C.NkModuleMatchListFn(C.modulematchlist);
	res.notificationsend = C.NkModuleNotificationSendFn(C.modulenotificationsend);
	res.notificationssend = C.NkModuleNotificationsSendFn(C.modulenotificationssend);
	res.walletupdate = C.NkModuleWalletUpdateFn(C.modulewalletupdate);
	res.walletsupdate = C.NkModuleWalletsUpdateFn(C.modulewalletsupdate);
	res.walletledgerupdate = C.NkModuleWalletLedgerUpdateFn(C.modulewalletledgerupdate);
	res.walletledgerlist = C.NkModuleWalletLedgerListFn(C.modulewalletledgerlist);
	res.storagelist = C.NkModuleStorageListFn(C.modulestoragelist);
	res.storageread = C.NkModuleStorageReadFn(C.modulestorageread);
	res.storagewrite = C.NkModuleStorageWriteFn(C.modulestoragewrite);
	res.storagedelete = C.NkModuleStorageDeleteFn(C.modulestoragedelete);
	res.multiupdate = C.NkModuleMultiUpdateFn(C.modulemultiupdate);
	res.leaderboardcreate = C.NkModuleLeaderboardCreateFn(C.moduleleaderboardcreate);
	res.leaderboarddelete = C.NkModuleLeaderboardDeleteFn(C.moduleleaderboarddelete);
	res.leaderboardrecordslist = C.NkModuleLeaderboardRecordsListFn(C.moduleleaderboardrecordslist);
	res.leaderboardrecordwrite = C.NkModuleLeaderboardRecordWriteFn(C.moduleleaderboardrecordwrite);
	res.leaderboardrecorddelete = C.NkModuleLeaderboardRecordDeleteFn(C.moduleleaderboardrecorddelete);
	res.tournamentcreate = C.NkModuleTournamentCreateFn(C.moduletournamentcreate);
	res.tournamentdelete = C.NkModuleTournamentDeleteFn(C.moduletournamentdelete);
	res.tournamentaddattempt = C.NkModuleTournamentAddAttemptFn(C.moduletournamentaddattempt);
	res.tournamentjoin = C.NkModuleTournamentJoinFn(C.moduletournamentjoin);
	res.tournamentsgetid = C.NkModuleTournamentsGetIdFn(C.moduletournamentsgetid);
	res.tournamentlist = C.NkModuleTournamentListFn(C.moduletournamentlist);
	res.tournamentrecordslist = C.NkModuleTournamentRecordsListFn(C.moduletournamentrecordslist);
	res.tournamentrecordwrite = C.NkModuleTournamentRecordWriteFn(C.moduletournamentrecordwrite);
	res.tournamentrecordshaystack = C.NkModuleTournamentRecordsHaystackFn(C.moduletournamentrecordshaystack);
	res.groupgetid = C.NkModuleGroupGetIdFn(C.modulegroupgetid);
	res.groupcreate = C.NkModuleGroupCreateFn(C.modulegroupcreate);
	res.groupupdate = C.NkModuleGroupUpdateFn(C.modulegroupupdate);
	res.groupdelete = C.NkModuleGroupDeleteFn(C.modulegroupdelete);
	res.groupuserjoin = C.NkModuleGroupUserJoinFn(C.modulegroupuserjoin);
	res.groupuserleave = C.NkModuleGroupUserLeaveFn(C.modulegroupuserleave);
	res.groupusersadd = C.NkModuleGroupUsersAddFn(C.modulegroupusersadd);
	res.groupuserskick = C.NkModuleGroupUsersKickFn(C.modulegroupuserskick);
	res.groupuserspromote = C.NkModuleGroupUsersPromoteFn(C.modulegroupuserspromote);
	res.groupusersdemote = C.NkModuleGroupUsersDemoteFn(C.modulegroupusersdemote);
	res.groupuserslist = C.NkModuleGroupUsersListFn(C.modulegroupuserslist);
	res.usergroupslist = C.NkModuleUserGroupsListFn(C.moduleusergroupslist);
	res.friendslist = C.NkModuleFriendsListFn(C.modulefriendslist);
	res.event = C.NkModuleEventFn(C.moduleevent);

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

	C.initmodule(c.syms.initModule, cCtx, cLogger, cDb, cNk, cInitializer)

	for _, alloc := range pointer.Restore(cCtx.ptr).(*RuntimeCContextCall).allocs {
		C.free(alloc)
	}

	for _, alloc := range pointer.Restore(cLogger.ptr).(*RuntimeCNakamaModuleCall).allocs {
		C.free(alloc)
	}

	pointer.Unref(cLogger.ptr)
	pointer.Unref(cNk.ptr)

	return nil
}
