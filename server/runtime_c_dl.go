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
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

// DLOpen loads the c-module symbols from the given module filename.
func DLOpen(name string) (CSymbols, error) {
	syms := CSymbols{}

	cName := C.CString(name)
	lib := C.dlopen(cName, C.RTLD_LAZY)
	C.free(unsafe.Pointer(cName))
	if lib == nil {
		return syms, errors.New("Unable to open c-module")
	}

	syms.Load(lib)

	if syms.initModule == nil {
		return syms, errors.New("Unable to open c-module: module initialization function not found")
	}

	return syms, nil
}

// DLSym gets the unsafe.Pointer to a function with the given name from the given library.
func DLSym(lib unsafe.Pointer, name string) unsafe.Pointer {
	cName := C.CString(name)
	sym := C.dlsym(lib, cName)
	C.free(unsafe.Pointer(cName))

	return sym
}
