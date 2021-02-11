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

NkTournament *mallocnktournament(NkU32 cnt)
{
	return malloc(sizeof(NkTournament) * cnt);
}

NkUser *mallocnkuser(NkU32 cnt)
{
	return malloc(sizeof(NkUser) * cnt);
}

NkWalletLedgerItem *mallocnkwalletledgeritem(NkU32 cnt)
{
	return malloc(sizeof(NkWalletLedgerItem) * cnt);
}

NkWalletUpdateResult *mallocnkwalletupdateresult(NkU32 cnt)
{
	return malloc(sizeof(NkWalletUpdateResult) * cnt);
}

void nkaccountset(NkAccount *ptr, NkU32 idx, NkAccount acct)
{
	ptr[idx] = acct;
}

void nkmapi64set(NkMapI64 *ptr, NkU32 idx, NkString key, NkI64 val)
{
	ptr->keys[idx] = key;
	ptr->vals[idx] = val;
}

void nkmatchset(NkMatch *ptr, NkU32 idx, NkMatch match)
{
	ptr[idx] = match;
}

void nkpresenceset(NkPresence *ptr, NkU32 idx, NkPresence presence)
{
	ptr[idx] = presence;
}

void nkstorageobjectset(NkStorageObject *ptr, NkU32 idx, NkStorageObject obj)
{
	ptr[idx] = obj;
}

void nktournamentset(NkTournament *ptr, NkU32 idx, NkTournament tournament)
{
	ptr[idx] = tournament;
}

void nkuserset(NkUser *ptr, NkU32 idx, NkUser acct)
{
	ptr[idx] = acct;
}

void nkwalletledgeritemset(NkWalletLedgerItem *ptr, NkU32 idx, NkWalletLedgerItem item)
{
	ptr[idx] = item;
}

void nkwalletupdateresultset(NkWalletUpdateResult *ptr, NkU32 idx, NkWalletUpdateResult result)
{
	ptr[idx] = result;
}

// Start of internal implementations

void nkmapanyi(NkMapAny map, NkU32 idx, NkString **key, NkAny **val)
{
	*key = &map.keys[idx];
	*val = &map.vals[idx];
}

void nkmapi64i(NkMapI64 map, NkU32 idx, NkString **key, NkI64 **val)
{
	*key = &map.keys[idx];
	*val = &map.vals[idx];
}

void nkmapstringi(NkMapString map, NkU32 idx, NkString **key, NkString **val)
{
	*key = &map.keys[idx];
	*val = &map.vals[idx];
}

void nknotificationsendi(NkNotificationSend *arr, NkU32 idx, NkNotificationSend **n)
{
	*n = &arr[idx];
}

void nkpresencei(NkPresence *arr, NkU32 idx, NkPresence **p)
{
	*p = &arr[idx];
}

void nkstringi(NkString *arr, NkU32 idx, NkString **s)
{
	*s = &arr[idx];
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

func goEnvelope(e C.NkEnvelope) *rtapi.Envelope {
	return pointer.Restore(e.ptr).(*rtapi.Envelope)
}

func goMapAny(m C.NkMapAny) map[string]interface{} {
	cnt := int(m.cnt)
	ret := make(map[string]interface{}, cnt)
	for i := 0; i < cnt; i++ {
		var key *C.NkString
		var val *C.NkAny
		C.nkmapanyi(m, C.NkU32(i), &key, &val)
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
		C.nkmapi64i(m, C.NkU32(i), &key, &val)
		ret[goString(*key)] = int64(*val)
	}

	return ret
}

func goMapString(m C.NkMapString) map[string]string {
	cnt := int(m.cnt)
	ret := make(map[string]string, cnt)
	for i := 0; i < cnt; i++ {
		var key, val *C.NkString
		C.nkmapstringi(m, C.NkU32(i), &key, &val)
		ret[goString(*key)] = goString(*val)
	}

	return ret
}

func goNotificationSends(ptr *C.NkNotificationSend, cnt C.NkU32) []*runtime.NotificationSend {
	goCnt := int(cnt)
	ret := make([]*runtime.NotificationSend, goCnt)
	for i := 0; i < goCnt; i++ {
		var n *C.NkNotificationSend
		C.nknotificationsendi(ptr, C.NkU32(i), &n)
		ret[i] = pointer.Restore((*n).ptr).(*runtime.NotificationSend)
	}

	return ret
}

func goPresences(ptr *C.NkPresence, cnt C.NkU32) []runtime.Presence {
	goCnt := int(cnt)
	ret := make([]runtime.Presence, goCnt)
	for i := 0; i < goCnt; i++ {
		var p *C.NkPresence
		C.nkpresencei(ptr, C.NkU32(i), &p)
		ret[i] = pointer.Restore((*p).ptr).(runtime.Presence)
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
		C.nkstringi(ptr, C.NkU32(i), &s)
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
