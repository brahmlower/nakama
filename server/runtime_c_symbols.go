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

/*

This script was used to create the include/nakama.h file - it is not normally built by CGO because
these symbols are not public exports: they are given directly to the C libraries as *void.

go build -trimpath -mod=vendor -buildmode=c-shared server/runtime_c_symbols.go
rm include/nakama.h
touch include/nakama.h
echo -e "// Copyright 2021 The Nakama Authors" | tee -a include/nakama.h
echo -e "//" | tee -a include/nakama.h
echo -e "// Licensed under the Apache License, Version 2.0 (the "License");" | tee -a include/nakama.h
echo -e "// you may not use this file except in compliance with the License." | tee -a include/nakama.h
echo -e "// You may obtain a copy of the License at" | tee -a include/nakama.h
echo -e "//" | tee -a include/nakama.h
echo -e "// http://www.apache.org/licenses/LICENSE-2.0" | tee -a include/nakama.h
echo -e "//" | tee -a include/nakama.h
echo -e "// Unless required by applicable law or agreed to in writing, software" | tee -a include/nakama.h
echo -e "// distributed under the License is distributed on an "AS IS" BASIS," | tee -a include/nakama.h
echo -e "// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied." | tee -a include/nakama.h
echo -e "// See the License for the specific language governing permissions and" | tee -a include/nakama.h
echo -e "// limitations under the License." | tee -a include/nakama.h
echo -e "" | tee -a include/nakama.h
cat runtime_c_symbols.h >> include/nakama.h

package main

import "C"

func main() {}

*/

package server

//export HelloWorld
func HelloWorld(arg1, arg2 int, arg3 string) int64 {
	return 0
}
