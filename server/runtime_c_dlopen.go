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
		if l.err != "" {
			return nil, errors.New(`OpenSharedLib("` + name + `"): ` + l.err + ` (previous failure)`)
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
	libpath := ""
	//libpath, syms, errstr := lastmoduleinit()
	// if errstr != "" {
	// 	sharedLibs[filepath] = &SharedLib{
	// 		libpath: libpath,
	// 		err:        errstr,
	// 	}
	// 	sharedLibMu.Unlock()
	// 	return nil, errors.New(`OpenSharedLib("` + name + `"): ` + errstr)
	// }
	// This function can be called from the init function of a plugin.
	// Drop a placeholder in the map so subsequent opens can wait on it.
	l := &SharedLib{
		libpath: libpath,
		loaded:     make(chan struct{}),
	}
	sharedLibs[filepath] = l
	sharedLibMu.Unlock()

	// initStr := make([]byte, len(pluginpath)+len("..inittask")+1) // +1 for terminating NUL
	// copy(initStr, pluginpath)
	// copy(initStr[len(pluginpath):], "..inittask")

	// initTask := C.pluginLookup(h, (*C.char)(unsafe.Pointer(&initStr[0])), &cErr)
	// if initTask != nil {
	// 	doInit(initTask)
	// }

	/*/ Fill out the value of each plugin symbol.
	updatedSyms := map[string]interface{}{}
	for symName, sym := range syms {
		isFunc := symName[0] == '.'
		if isFunc {
			delete(syms, symName)
			symName = symName[1:]
		}

		fullName := pluginpath + "." + symName
		cname := make([]byte, len(fullName)+1)
		copy(cname, fullName)

		p := C.pluginLookup(h, (*C.char)(unsafe.Pointer(&cname[0])), &cErr)
		if p == nil {
			return nil, errors.New(`plugin.Open("` + name + `"): could not find symbol ` + symName + `: ` + C.GoString(cErr))
		}
		valp := (*[2]unsafe.Pointer)(unsafe.Pointer(&sym))
		if isFunc {
			(*valp)[1] = unsafe.Pointer(&p)
		} else {
			(*valp)[1] = p
		}
		// we can't add to syms during iteration as we'll end up processing
		// some symbols twice with the inability to tell if the symbol is a function
		updatedSyms[symName] = sym
	}
	p.syms = updatedSyms*/

	close(l.loaded)
	return l, nil
}

func (l *SharedLib) lookup(symName string) (Symbol, error) {
	if s := l.syms[symName]; s != nil {
		return s, nil
	}
	return nil, errors.New("shared library: symbol " + symName + " not found in " + l.libpath)
}

var (
	sharedLibMu sync.Mutex
	sharedLibs   map[string]*SharedLib
)
