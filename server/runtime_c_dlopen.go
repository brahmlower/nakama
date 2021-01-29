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

// +build linux,cgo darwin,cgo freebsd,cgo

package server

/*
#cgo linux LDFLAGS: -ldl
#include <dlfcn.h>
#include <limits.h>
#include <stdlib.h>
#include <stdint.h>

#include <stdio.h>

static uintptr_t sharedLibOpen(const char* path, char** err) {
	void* h = dlopen(path, RTLD_NOW|RTLD_GLOBAL);
	if (h == NULL) {
		*err = (char*)dlerror();
	}
	return (uintptr_t)h;
}

static void* sharedLibLookup(uintptr_t h, const char* name, char** err) {
	void* r = dlsym((void*)h, name);
	if (r == NULL) {
		*err = (char*)dlerror();
	}
	return r;
}
*/
import "C"

import (
	"errors"
	"sync"
	"unsafe"
)

func OpenSharedLib(name string) (*SharedLib, error) {
	cPath := make([]byte, C.PATH_MAX+1)
	cRelName := make([]byte, len(name)+1)
	copy(cRelName, name)
	if C.realpath(
		(*C.char)(unsafe.Pointer(&cRelName[0])),
		(*C.char)(unsafe.Pointer(&cPath[0]))) == nil {
		return nil, errors.New(`OpenSharedLib("` + name + `"): realpath failed`)
	}

	filepath := C.GoString((*C.char)(unsafe.Pointer(&cPath[0])))

	sharedLibMu.Lock()
	if l := sharedLibs[filepath]; l != nil {
		sharedLibMu.Unlock()
		if l.err != nil {
			return nil, errors.New(`OpenSharedLib("` + name + `"): ` + *l.err + ` (previous failure)`)
		}
		<-l.loaded
		return l, nil
	}
	var cErr *C.char
	h := C.sharedLibOpen((*C.char)(unsafe.Pointer(&cPath[0])), &cErr)
	if h == 0 {
		sharedLibMu.Unlock()
		return nil, errors.New(`OpenSharedLib("` + name + `"): ` + C.GoString(cErr))
	}
	if len(name) > 3 && name[len(name)-3:] == ".so" {
		name = name[:len(name)-3]
	}
	if sharedLibs == nil {
		sharedLibs = make(map[string]*SharedLib)
	}
	syms, err := loadCSymbols(h)
	if err != nil {
		sharedLibs[filepath] = &SharedLib{
			err:  err,
			path: filepath,
		}
		sharedLibMu.Unlock()
		return nil, errors.New(`OpenSharedLib("` + name + `"): ` + *err)
	}
	// This function can be called from the init function of a plugin.
	// Drop a placeholder in the map so subsequent opens can wait on it.
	l := &SharedLib{
		loaded: make(chan struct{}),
		path:   filepath,
		syms:   syms,
	}
	sharedLibs[filepath] = l
	sharedLibMu.Unlock()

	// p, _ := l.lookup("c_nk_init_module")
	// C.cNkInitModule(p, 2468)

	close(l.loaded)
	return l, nil
}

func loadCSymbol(h C.ulong, syms map[string]CSymbol, name string) *string {
	var cErr *C.char
	var cName *C.char
	cName = C.CString(name)
	syms[name] = C.sharedLibLookup(h, cName, &cErr)
	C.free(unsafe.Pointer(cName))
	if cErr != nil {
		err := C.GoString(cErr)
		return &err
	}

	return nil
}

func loadCSymbols(h C.ulong) (map[string]CSymbol, *string) {
	syms := make(map[string]CSymbol)
	if err := loadCSymbol(h, syms, "c_nk_init_module"); err != nil {
		return nil, err
	}

	return syms, nil
}

func (l *SharedLib) lookup(symName string) (CSymbol, error) {
	if s := l.syms[symName]; s != nil {
		return s, nil
	}
	return nil, errors.New("shared library: symbol " + symName + " not found in " + l.path)
}

var (
	sharedLibMu sync.Mutex
	sharedLibs  map[string]*SharedLib
)
