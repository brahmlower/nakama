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

import (
	"unsafe"

	"github.com/mattn/go-pointer"
)

type RuntimeCCall struct {
	ptr    unsafe.Pointer
	refs []unsafe.Pointer
}

func (r *RuntimeCCall) unref() {
	for _, ptr := range r.refs {
		pointer.Unref(ptr)
	}

	pointer.Unref(r.ptr)
}