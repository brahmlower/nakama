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

// Start of extern implementations

bool *mallocbool(NkU32 cnt)
{
	return malloc(sizeof(bool) * cnt);
}

NkAccount *mallocnkaccount(NkU32 cnt)
{
	return malloc(sizeof(NkAccount) * cnt);
}

NkFriend *mallocnkfriend(NkU32 cnt)
{
	return malloc(sizeof(NkFriend) * cnt);
}

NkGroup *mallocnkgroup(NkU32 cnt)
{
	return malloc(sizeof(NkGroup) * cnt);
}

NkGroupUserListGroupUser *mallocnkgroupuserlistgroupuser(NkU32 cnt)
{
	return malloc(sizeof(NkGroupUserListGroupUser) * cnt);
}

NkString *mallocnkstring(NkU32 cnt)
{
	return malloc(sizeof(NkAccount) * cnt);
}

NkI64 *mallocnki64(NkU32 cnt)
{
	return malloc(sizeof(NkI64) * cnt);
}

NkMapI64 *mallocnkmapi64(NkU32 cnt)
{
	NkMapI64 *ret = malloc(sizeof(NkMapI64));
	ret->cnt = cnt;
	ret->keys = mallocnkstring(cnt);
	ret->vals = mallocnki64(cnt);
}

NkLeaderboardRecord *mallocnkleaderboardrecord(NkU32 cnt)
{
	return malloc(sizeof(NkLeaderboardRecord) * cnt);
}

NkMatch *mallocnkmatch(NkU32 cnt)
{
	return malloc(sizeof(NkMatch) * cnt);
}

NkPresence *mallocnkpresence(NkU32 cnt)
{
	return malloc(sizeof(NkPresence) * cnt);
}

NkPresenceMeta *mallocnkpresencemeta(NkU32 cnt)
{
	return malloc(sizeof(NkPresenceMeta) * cnt);
}

NkU32 *mallocnku32(NkU32 cnt)
{
	return malloc(sizeof(NkU32) * cnt);
}

NkU64 *mallocnku64(NkU32 cnt)
{
	return malloc(sizeof(NkU64) * cnt);
}

NkStorageObject *mallocnkstorageobject(NkU32 cnt)
{
	return malloc(sizeof(NkStorageObject) * cnt);
}

NkStorageObjectAck *mallocnkstorageobjectack(NkU32 cnt)
{
	return malloc(sizeof(NkStorageObjectAck) * cnt);
}

NkTournament *mallocnktournament(NkU32 cnt)
{
	return malloc(sizeof(NkTournament) * cnt);
}

NkTournamentList *mallocnktournamentlist(NkU32 cnt)
{
	return malloc(sizeof(NkTournamentList) * cnt);
}

NkUser *mallocnkuser(NkU32 cnt)
{
	return malloc(sizeof(NkUser) * cnt);
}

NkUserGroupListUserGroup *mallocnkusergrouplistusergroup(NkU32 cnt)
{
	return malloc(sizeof(NkUserGroupListUserGroup) * cnt);
}

NkWalletLedgerItem *mallocnkwalletledgeritem(NkU32 cnt)
{
	return malloc(sizeof(NkWalletLedgerItem) * cnt);
}

NkWalletUpdateResult *mallocnkwalletupdateresult(NkU32 cnt)
{
	return malloc(sizeof(NkWalletUpdateResult) * cnt);
}

void nkaccountset(NkAccount *accts, NkU32 idx, NkAccount acct)
{
	accts[idx] = acct;
}

void nkfriendset(NkFriend *friends, NkU32 idx, NkFriend friend)
{
	friends[idx] = friend;
}

void nkgroupset(NkGroup *groups, NkU32 idx, NkGroup group)
{
	groups[idx] = group;
}

void nkleaderboardrecordset(NkLeaderboardRecord *records, NkU32 idx, NkLeaderboardRecord record)
{
	records[idx] = record;
}

void nkgroupuserlistgroupuserset(NkGroupUserListGroupUser *users, NkU32 idx, NkGroupUserListGroupUser user)
{
	users[idx] = user;
}

void nkmapi64set(NkMapI64 *ptr, NkU32 idx, NkString key, NkI64 val)
{
	ptr->keys[idx] = key;
	ptr->vals[idx] = val;
}

void nkmatchset(NkMatch *matches, NkU32 idx, NkMatch match)
{
	matches[idx] = match;
}

void nkpresenceset(NkPresence *presences, NkU32 idx, NkPresence presence)
{
	presences[idx] = presence;
}

void nkstorageobjectset(NkStorageObject *objs, NkU32 idx, NkStorageObject obj)
{
	objs[idx] = obj;
}

void nkstorageobjectackset(NkStorageObjectAck *acks, NkU32 idx, NkStorageObjectAck ack)
{
	acks[idx] = ack;
}

void nktournamentset(NkTournament *tournaments, NkU32 idx, NkTournament tournament)
{
	tournaments[idx] = tournament;
}

void nkusergrouplistusergroupset(NkUserGroupListUserGroup *groups, NkU32 idx, NkUserGroupListUserGroup group)
{
	groups[idx] = group;
}

void nkuserset(NkUser *accts, NkU32 idx, NkUser acct)
{
	accts[idx] = acct;
}

void nkwalletledgeritemset(NkWalletLedgerItem *items, NkU32 idx, NkWalletLedgerItem item)
{
	items[idx] = item;
}

void nkwalletupdateresultset(NkWalletUpdateResult *results, NkU32 idx, NkWalletUpdateResult result)
{
	results[idx] = result;
}

// Start of internal implementations

void nkaccountupdateget(NkAccountUpdate *updates, NkU32 idx, NkAccountUpdate **outupdate)
{
	*outupdate = &updates[idx];
}

void nkmapanyget(NkMapAny map, NkU32 idx, NkString **key, NkAny **val)
{
	*key = &map.keys[idx];
	*val = &map.vals[idx];
}

void nkmapi64get(NkMapI64 map, NkU32 idx, NkString **key, NkI64 **val)
{
	*key = &map.keys[idx];
	*val = &map.vals[idx];
}

void nkmapstringget(NkMapString map, NkU32 idx, NkString **key, NkString **val)
{
	*key = &map.keys[idx];
	*val = &map.vals[idx];
}

void nknotificationsendget(NkNotificationSend *sends, NkU32 idx, NkNotificationSend **outsend)
{
	*outsend = &sends[idx];
}

void nkpresenceget(NkPresence *presences, NkU32 idx, NkPresence **outpresence)
{
	*outpresence = &presences[idx];
}

void nkstoragedeleteget(NkStorageDelete *deletes, NkU32 idx, NkStorageDelete **outdelete)
{
	*outdelete = &deletes[idx];
}

void nkstoragereadget(NkStorageRead *reads, NkU32 idx, NkStorageRead **outread)
{
	*outread = &reads[idx];
}

void nkstoragewriteget(NkStorageWrite *writes, NkU32 idx, NkStorageWrite **outwrite)
{
	*outwrite = &writes[idx];
}

void nkstringget(NkString *strings, NkU32 idx, NkString **outstring)
{
	*outstring = &strings[idx];
}

NkU8 nkvalueu8(NkValue val)
{
	return val.u8;
}

NkU16 nkvalueu16(NkValue val)
{
	return val.u16;
}

NkU32 nkvalueu32(NkValue val)
{
	return val.u32;
}

NkU64 nkvalueu64(NkValue val)
{
	return val.u64;
}

NkI8 nkvaluei8(NkValue val)
{
	return val.i8;
}

NkI16 nkvaluei16(NkValue val)
{
	return val.i16;
}

NkI32 nkvaluei32(NkValue val)
{
	return val.i32;
}

NkI64 nkvaluei64(NkValue val)
{
	return val.i64;
}

NkF32 nkvaluef32(NkValue val)
{
	return val.f32;
}

NkF64 nkvaluef64(NkValue val)
{
	return val.f64;
}

NkString nkvaluestring(NkValue val)
{
	return val.s;
}

void nkwalletupdatei(NkWalletUpdate *arr, NkU32 idx, NkWalletUpdate **u)
{
	*u = &arr[idx];
}
*/
import "C"

import (
	"github.com/heroiclabs/nakama-common/rtapi"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/mattn/go-pointer"
)

func cNkMapI64(m map[string]int64) *C.NkMapI64 {
	cnt := C.NkU32(len(m))
	ret := C.mallocnkmapi64(cnt)

	idx := 0
	for key, val := range m {
		C.nkmapi64set(ret, C.NkU32(idx), cNkString(key), C.NkI64(val))
		idx++
	}

	return ret
}

func cNkString(s string) C.NkString {
	ret := C.NkString{}
	ret.p = C.CString(s)
	ret.n = C.long(len(s))

	return ret
}

func goAccountUpdate(update C.NkAccountUpdate) *runtime.AccountUpdate {
	ret := &runtime.AccountUpdate{}
	ret.AvatarUrl = goString(update.avatarurl)
	ret.DisplayName = goString(update.displayname)
	ret.LangTag = goString(update.langtag)
	ret.Location = goString(update.location)
	ret.Metadata = goMapAny(update.metadata)
	ret.Timezone = goString(update.timezone)
	ret.UserID = goString(update.userid)
	ret.Username = goString(update.username)

	return ret
}

func goAccountUpdates(ptr *C.NkAccountUpdate, cnt C.NkU32) []*runtime.AccountUpdate {
	goCnt := int(cnt)
	ret := make([]*runtime.AccountUpdate, goCnt)
	for i := 0; i < goCnt; i++ {
		var u *C.NkAccountUpdate
		C.nkaccountupdateget(ptr, C.NkU32(i), &u)
		ret[i] = goAccountUpdate(*u)
	}

	return ret
}

func goEnvelope(e C.NkEnvelope) *rtapi.Envelope {
	return pointer.Restore(e.ptr).(*rtapi.Envelope)
}

func goMapAny(m C.NkMapAny) map[string]interface{} {
	cnt := int(m.cnt)
	ret := make(map[string]interface{}, cnt)
	for i := 0; i < cnt; i++ {
		var key *C.NkString
		var val *C.NkAny
		C.nkmapanyget(m, C.NkU32(i), &key, &val)
		ret[goString(*key)] = goValueAny(*val)
	}

	return ret
}

func goMapI64(m C.NkMapI64) map[string]int64 {
	cnt := int(m.cnt)
	ret := make(map[string]int64, cnt)
	for i := 0; i < cnt; i++ {
		var key *C.NkString
		var val *C.NkI64
		C.nkmapi64get(m, C.NkU32(i), &key, &val)
		ret[goString(*key)] = int64(*val)
	}

	return ret
}

func goMapString(m C.NkMapString) map[string]string {
	cnt := int(m.cnt)
	ret := make(map[string]string, cnt)
	for i := 0; i < cnt; i++ {
		var key, val *C.NkString
		C.nkmapstringget(m, C.NkU32(i), &key, &val)
		ret[goString(*key)] = goString(*val)
	}

	return ret
}

func goNotificationSends(ptr *C.NkNotificationSend, cnt C.NkU32) []*runtime.NotificationSend {
	goCnt := int(cnt)
	ret := make([]*runtime.NotificationSend, goCnt)
	for i := 0; i < goCnt; i++ {
		var n *C.NkNotificationSend
		C.nknotificationsendget(ptr, C.NkU32(i), &n)
		ret[i] = pointer.Restore((*n).ptr).(*runtime.NotificationSend)
	}

	return ret
}

func goPresences(ptr *C.NkPresence, cnt C.NkU32) []runtime.Presence {
	goCnt := int(cnt)
	ret := make([]runtime.Presence, goCnt)
	for i := 0; i < goCnt; i++ {
		var p *C.NkPresence
		C.nkpresenceget(ptr, C.NkU32(i), &p)
		ret[i] = pointer.Restore((*p).ptr).(runtime.Presence)
	}

	return ret
}

func goStorageDelete(del C.NkStorageDelete) *runtime.StorageDelete {
	ret := &runtime.StorageDelete{}
	ret.Collection = goString(del.collection)
	ret.Key = goString(del.key)
	ret.UserID = goString(del.userid)
	ret.Version = goString(del.version)

	return ret
}

func goStorageDeletes(ptr *C.NkStorageDelete, cnt C.NkU32) []*runtime.StorageDelete {
	goCnt := int(cnt)
	ret := make([]*runtime.StorageDelete, goCnt)
	for i := 0; i < goCnt; i++ {
		var d *C.NkStorageDelete
		C.nkstoragedeleteget(ptr, C.NkU32(i), &d)
		ret[i] = goStorageDelete(*d)
	}

	return ret
}

func goStorageRead(read C.NkStorageRead) *runtime.StorageRead {
	ret := &runtime.StorageRead{}
	ret.Collection = goString(read.collection)
	ret.Key = goString(read.key)
	ret.UserID = goString(read.userid)

	return ret
}

func goStorageReads(ptr *C.NkStorageRead, cnt C.NkU32) []*runtime.StorageRead {
	goCnt := int(cnt)
	ret := make([]*runtime.StorageRead, goCnt)
	for i := 0; i < goCnt; i++ {
		var r *C.NkStorageRead
		C.nkstoragereadget(ptr, C.NkU32(i), &r)
		ret[i] = goStorageRead(*r)
	}

	return ret
}

func goStorageWrite(read C.NkStorageWrite) *runtime.StorageWrite {
	ret := &runtime.StorageWrite{}
	ret.Collection = goString(read.collection)
	ret.Key = goString(read.key)
	ret.PermissionRead = int(read.permissionread)
	ret.PermissionWrite = int(read.permissionwrite)
	ret.UserID = goString(read.userid)
	ret.Value = goString(read.value)
	ret.Version = goString(read.version)

	return ret
}

func goStorageWrites(ptr *C.NkStorageWrite, cnt C.NkU32) []*runtime.StorageWrite {
	goCnt := int(cnt)
	ret := make([]*runtime.StorageWrite, goCnt)
	for i := 0; i < goCnt; i++ {
		var w *C.NkStorageWrite
		C.nkstoragewriteget(ptr, C.NkU32(i), &w)
		ret[i] = goStorageWrite(*w)
	}

	return ret
}

func goString(s C.NkString) string {
	return C.GoStringN(s.p, C.int(s.n))
}

func goStrings(ptr *C.NkString, cnt C.NkU32) []string {
	goCnt := int(cnt)
	ret := make([]string, goCnt)
	for i := 0; i < goCnt; i++ {
		var s *C.NkString
		C.nkstringget(ptr, C.NkU32(i), &s)
		ret[i] = goString(*s)
	}

	return ret
}

func goValueAny(a C.NkAny) interface{} {
	switch a.ty {
	case C.NK_U8_T:
		return uint8(C.nkvalueu8(a.val))
	case C.NK_U16_T:
		return uint16(C.nkvalueu16(a.val))
	case C.NK_U32_T:
		return uint32(C.nkvalueu32(a.val))
	case C.NK_U64_T:
		return uint64(C.nkvalueu64(a.val))
	case C.NK_I8_T:
		return int8(C.nkvaluei8(a.val))
	case C.NK_I16_T:
		return int16(C.nkvaluei16(a.val))
	case C.NK_I32_T:
		return int32(C.nkvaluei32(a.val))
	case C.NK_I64_T:
		return int64(C.nkvaluei64(a.val))
	case C.NK_F32_T:
		return float32(C.nkvaluef32(a.val))
	case C.NK_F64_T:
		return float64(C.nkvaluef64(a.val))
	case C.NK_STRING_T:
		return goString(C.nkvaluestring(a.val))
	}

	panic("Not implemented")
}

func goWalletUpdates(ptr *C.NkWalletUpdate, cnt C.NkU32) []*runtime.WalletUpdate {
	goCnt := int(cnt)
	ret := make([]*runtime.WalletUpdate, goCnt)
	for i := 0; i < goCnt; i++ {
		var u *C.NkWalletUpdate
		C.nkwalletupdatei(ptr, C.NkU32(i), &u)
		ret[i] = &runtime.WalletUpdate{}
		ret[i].Changeset = goMapI64(u.changeset)
		ret[i].Metadata = goMapAny(u.metadata)
		ret[i].UserID = goString(u.userid)
	}

	return ret
}
