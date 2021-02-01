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

typedef void (*InitModuleFn)(NkLogger logger);

extern void cLoggerDebug(void *, NkString);
extern void cLoggerError(void *, NkString);
extern void cLoggerInfo(void *, NkString);
extern void cLoggerWarn(void *, NkString);

void initmodule(void *ptr, NkLogger logger)
{
	InitModuleFn fn = (InitModuleFn)ptr;
	fn(logger);
}

void loggerdebug(void *ptr, NkString s)
{
	cLoggerDebug(ptr, s);
}

void loggererror(void *ptr, NkString s)
{
	cLoggerError(ptr, s);
}

void loggerinfo(void *ptr, NkString s)
{
	cLoggerInfo(ptr, s);
}

void loggerwarn(void *ptr, NkString s)
{
	cLoggerWarn(ptr, s);
}
*/
import "C"
