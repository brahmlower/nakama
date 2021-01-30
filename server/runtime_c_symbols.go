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

void *NkLogLevelFnVoid(NkLogLevelFn f)
{
  return (void *)f;
}

NkLogLevelFn VoidNkLogLevelFn(void *f)
{
  return (NkLogLevelFn)f;
}

*/
import "C"

import (
	"context"
	"database/sql"

	//"unsafe"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/mattn/go-pointer"
	"github.com/rainycape/dl"
)

// //export HelloWorld
// func HelloWorld(arg1, arg2 int, arg3 string) int64 {
// 	return 0
// }
type CLogger struct {
}

type CSymbol interface{}

type CSymbols struct {
	initModule func(C.NkLogger)
}



func NewCSymbols(lib *dl.DL) (CSymbols, error) {
	syms := CSymbols{}
	if err := lib.Sym("c_nk_init_module", syms.InitModule); err != nil {
		return syms, err
	}

	return syms, nil
}

func (c *CSymbols) InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	cLogger := C.NkLogger{
		debug: C.VoidNkLogLevelFn(pointer.Save(func(s *C.char) {
			logger.Debug(C.GoString(s))
		})),
		warn: C.VoidNkLogLevelFn(pointer.Save(func(s *C.char) {
			logger.Warn(C.GoString(s))
		})),
	}
	c.initModule(cLogger)
	pointer.Unref(C.NkLogLevelFnVoid(cLogger.debug))
	pointer.Unref(C.NkLogLevelFnVoid(cLogger.warn))

	
	return nil
}
