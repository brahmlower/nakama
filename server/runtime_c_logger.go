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
	"unsafe"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/mattn/go-pointer"
)

//export cLoggerDebug
func cLoggerDebug(p unsafe.Pointer, s C.NkString) {
	pointer.Restore(p).(runtime.Logger).Debug(GoStringN(s))
}

//export cLoggerError
func cLoggerError(p unsafe.Pointer, s C.NkString) {
	pointer.Restore(p).(runtime.Logger).Error(GoStringN(s))
}

//export cLoggerInfo
func cLoggerInfo(p unsafe.Pointer, s C.NkString) {
	pointer.Restore(p).(runtime.Logger).Info(GoStringN(s))
}

//export cLoggerWarn
func cLoggerWarn(p unsafe.Pointer, s C.NkString) {
	pointer.Restore(p).(runtime.Logger).Warn(GoStringN(s))
}