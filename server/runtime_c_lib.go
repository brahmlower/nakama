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
#include "../include/nakama.h"

extern void initmodule(void *, NkContext, NkLogger, NkDb, NkModule, NkInitializer);

extern NkString contextvalue(void *, NkString key);

extern void loggerdebug(void *, NkString);
extern void loggererror(void *, NkString);
extern void loggerinfo(void *, NkString);
extern void loggerwarn(void *, NkString);

extern int moduleauthenticateapple(void *, NkContext, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int moduleauthenticatecustom(void *, NkContext, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int moduleauthenticatedevice(void *, NkContext, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int moduleauthenticateemail(void *, NkContext, NkString, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int moduleauthenticatefacebook(void *, NkContext, NkString, bool, NkString, bool, NkString *, NkString *, NkString *, bool *l);
extern int moduleauthenticatefacebookinstantgame(void *, NkContext, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int moduleauthenticategamecenter(void *, NkContext, NkString, NkString, NkI64, NkString, NkString, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int moduleauthenticategoogle(void *, NkContext, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);
extern int moduleauthenticatesteam(void *, NkContext, NkString, NkString, bool, NkString *, NkString *, NkString *, bool *);

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
