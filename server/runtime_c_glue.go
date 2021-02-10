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

NkI64 *mallocnki64(NkU32 cnt)
{
	return malloc(sizeof(NkI64) * cnt);
}

NkUser *mallocnkuser(NkU32 cnt)
{
	return malloc(sizeof(NkUser) * cnt);
}

void nkaccountset(NkAccount *ptr, NkU32 idx, NkAccount acct)
{
	ptr[idx] = acct;
}

void nkuserset(NkUser *ptr, NkU32 idx, NkUser acct)
{
	ptr[idx] = acct;
}

// Start of internal implementations

void nkmapanyi(NkMapAny map, NkU32 idx, NkString **key, NkAny **val)
{
	*key = &map.keys[idx];
	*val = &map.vals[idx];
}

void nkmapstringi(NkMapString map, NkU32 idx, NkString **key, NkString **val)
{
	*key = &map.keys[idx];
	*val = &map.vals[idx];
}

void nkstringi(NkString *p, NkU32 idx, NkString **s)
{
	*s = &p[idx];
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
*/
import "C"

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

func goString(s C.NkString) string {
	return C.GoStringN(s.p, C.int(s.n))
}

func goStringArray(ptr *C.NkString, cnt C.NkU32) []string {
	cCnt := int(cnt)
	ret := make([]string, cCnt)
	for i := 0; i < cCnt; i++ {
		var s *C.NkString
		C.nkstringi(ptr, C.NkU32(i), &s)
		ret[i] = goString(*s)
	}

	return ret
}
