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

// NOTE: In order to implement a c-module, you must provide the following functions:
//
// Module entrypoint:
// void nk_init_module(void *logger);
//
// Match initializer:
// void *nk_init_match(void *logger);

#include <stddef.h>

typedef unsigned char NkU8;
typedef unsigned short NkU16;
typedef unsigned int NkU32;
typedef unsigned long long NkU64;

typedef signed char NkI8;
typedef short NkI16;
typedef int NkI32;
typedef long long NkI64;

typedef float NkF32;
typedef double NkF64;

// typedef __SIZE_TYPE__ Ptr;
typedef struct
{
	const char *p;
	ptrdiff_t n;
} NkString;

#ifdef __cplusplus
extern "C"
{
#endif
	typedef void (*NkLoggerFieldsFn)(void *ptr);
	typedef void (*NkLoggerLevelFn)(void *ptr, NkString s);
	typedef void *(*NkLoggerWithFieldFn)(void *ptr, NkString key, NkString value);
	typedef void (*NkLoggerWithFieldsFn)(void *ptr);

	typedef struct
	{
		void *ptr;
		NkLoggerLevelFn debug;
		NkLoggerLevelFn error;
		NkLoggerFieldsFn fields;
		NkLoggerLevelFn info;
		NkLoggerLevelFn warn;
		NkLoggerWithFieldFn with_field;
		NkLoggerWithFieldsFn with_fields;
	} NkLogger;


#ifdef __cplusplus
}
#endif
