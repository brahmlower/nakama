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

extern void initmodule(void *, NkLogger);
extern void loggerdebug(void *, NkString);
extern void loggererror(void *, NkString);
extern void loggerinfo(void *, NkString);
extern void loggerwarn(void *, NkString);
*/
import "C"

import (
	"context"
	"database/sql"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/mattn/go-pointer"
)

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

	cLogger := C.NkLogger{}
	cLogger.ptr = pointer.Save(logger)
	cLogger.debug = C.NkLoggerLevelFn(C.loggerdebug)
	cLogger.error = C.NkLoggerLevelFn(C.loggererror)
	cLogger.info = C.NkLoggerLevelFn(C.loggerinfo)
	cLogger.warn = C.NkLoggerLevelFn(C.loggerwarn)

	C.initmodule(c.syms.initModule, cLogger)

	pointer.Unref(cLogger.ptr)

	return nil
}
