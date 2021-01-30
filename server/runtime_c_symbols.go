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

void Free(void *p)
{
	free(p);
}

void *NkLogLevelFnVoid(NkLogLevelFn f)
{
	return (void *)f;
}

NkLogLevelFn VoidNkLogLevelFn(void *f)
{
	return (NkLogLevelFn)f;
}

NkLogger *NewNkLogger(NkLogLevelFn debug, NkLogLevelFn warn)
{
	NkLogger *res = malloc(sizeof(NkLogger));
	if (res == NULL)
		return NULL;

	res->debug = debug;
	res->warn = warn;

	return res;
}
*/
import "C"

import (
	"context"
	"database/sql"
	"unsafe"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/mattn/go-pointer"
	"github.com/rainycape/dl"
)

type CSymbol interface{}

type CSymbols struct {
	initModule func(*C.NkLogger)
}

func NewCSymbols(lib *dl.DL) (CSymbols, error) {
	syms := CSymbols{}

	if err := lib.Sym("nk_init_module", &syms.initModule); err != nil {
		return syms, err
	}

	return syms, nil
}

func (c *CSymbols) InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	cLoggerDebug := C.VoidNkLogLevelFn(pointer.Save(func(s *C.char) {
		logger.Debug(C.GoString(s))
	}))
	cLoggerWarn := C.VoidNkLogLevelFn(pointer.Save(func(s *C.char) {
		logger.Warn(C.GoString(s))
	}))
	cLogger := C.NewNkLogger(cLoggerDebug, cLoggerWarn)
	c.initModule(cLogger)
	pointer.Unref(C.NkLogLevelFnVoid(cLoggerDebug))
	pointer.Unref(C.NkLogLevelFnVoid(cLoggerWarn))
	C.Free(unsafe.Pointer(cLogger))

	return nil
}
