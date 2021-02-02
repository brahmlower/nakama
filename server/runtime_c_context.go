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

// #include "../include/nakama.h"
import "C"

import (
	"context"
	"fmt"
	"unsafe"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/mattn/go-pointer"
)

type RuntimeCContextCall struct {
	ctx    context.Context
	allocs []unsafe.Pointer
}

func (c *RuntimeCContextCall) goStringNk(s string) *C.NkString {
	var ret *C.NkString
	ret.p = C.CString(s)
	ret.n = C.long(len(s))

	c.allocs = append(c.allocs, unsafe.Pointer(ret.p))

	return ret
}

//export cContextValue
func cContextValue(pCtx unsafe.Pointer, key C.NkString, outValue *C.NkString) {
	call := pointer.Restore(pCtx).(*RuntimeCContextCall)
	retValue := call.ctx.Value(nkStringGo(key))
	outValue = call.goStringNk(fmt.Sprintf("%v", retValue))
}

func NewRuntimeCContext(ctx context.Context, node string, env map[string]string, mode RuntimeExecutionMode, queryParams map[string][]string, sessionExpiry int64, userID, username string, vars map[string]string, sessionID, clientIP, clientPort string) context.Context {
	ctx = context.WithValue(ctx, runtime.RUNTIME_CTX_ENV, env)
	ctx = context.WithValue(ctx, runtime.RUNTIME_CTX_MODE, mode.String())
	ctx = context.WithValue(ctx, runtime.RUNTIME_CTX_NODE, node)

	if queryParams != nil {
		ctx = context.WithValue(ctx, runtime.RUNTIME_CTX_QUERY_PARAMS, queryParams)
	}

	if userID != "" {
		ctx = context.WithValue(ctx, runtime.RUNTIME_CTX_USER_ID, userID)
		ctx = context.WithValue(ctx, runtime.RUNTIME_CTX_USERNAME, username)
		if vars != nil {
			ctx = context.WithValue(ctx, runtime.RUNTIME_CTX_VARS, vars)
		}
		ctx = context.WithValue(ctx, runtime.RUNTIME_CTX_USER_SESSION_EXP, sessionExpiry)
		if sessionID != "" {
			ctx = context.WithValue(ctx, runtime.RUNTIME_CTX_SESSION_ID, sessionID)
		}
	}

	if clientIP != "" {
		ctx = context.WithValue(ctx, runtime.RUNTIME_CTX_CLIENT_IP, clientIP)
	}
	if clientPort != "" {
		ctx = context.WithValue(ctx, runtime.RUNTIME_CTX_CLIENT_PORT, clientPort)
	}

	return ctx
}
