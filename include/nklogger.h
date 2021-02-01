// Copyright 2021 The Nakama Authors
//
// Licensed under the Apache License, Version 2.0 (the License);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an AS IS BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#ifndef NKLOGGER_H_INCLUDED
#define NKLOGGER_H_INCLUDED

#include "hashmap.h"
#include "nktypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

	typedef struct hashmap_s (*NkLoggerFieldsFn)(void *ptr);

	typedef void (*NkLoggerLevelFn)(void *ptr, NkString s);

	typedef struct NkLogger (*NkLoggerWithFieldFn)(void *ptr, NkString key, NkString value);

	typedef struct NkLogger (*NkLoggerWithFieldsFn)(void *ptr, struct hashmap_s fields);

	typedef struct
	{
		void *ptr;
		NkLoggerLevelFn debug;
		NkLoggerLevelFn error;
		NkLoggerFieldsFn fields;
		NkLoggerLevelFn info;
		NkLoggerLevelFn warn;
		NkLoggerWithFieldFn withfield;
		NkLoggerWithFieldsFn withfields;
	} NkLogger;

#ifdef __cplusplus
}
#endif

#endif