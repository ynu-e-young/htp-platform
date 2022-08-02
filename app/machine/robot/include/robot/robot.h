//
// Created by Homin Su on 2022/8/1.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_ROBOT_H_
#define HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_ROBOT_H_

#ifndef HTP_PLATFORM_ASSERT
#include <cassert>
#define HTP_PLATFORM_ASSERT(x) assert(x)
#endif // HTP_PLATFORM_ASSERT

#ifndef HTP_PLATFORM_LENGTH
#define HTP_PLATFORM_LENGTH(CONST_ARRAY) sizeof(CONST_ARRAY) / sizeof(CONST_ARRAY[0])
#endif // HTP_PLATFORM_LENGTH

#ifndef HTP_PLATFORM_STR_LENGTH
#if defined(_MSC_VER)
#define HTP_PLATFORM_STR_LENGTH(CONST_STR) _countof(CONST_STR)
#else
#define HTP_PLATFORM_STR_LENGTH(CONST_STR) sizeof(CONST_STR) / sizeof(CONST_STR[0])
#endif
#endif // HTP_PLATFORM_STR_LENGTH

#if defined(_WIN64) || defined(WIN64) || defined(_WIN32) || defined(WIN32)
#if defined(_WIN64) || defined(WIN64)
#define HTP_PLATFORM_ARCH_64 1
#else
#define HTP_PLATFORM_ARCH_32 1
#endif
#define HTP_PLATFORM_PLATFORM_STRING "windows"
#define HTP_PLATFORM_WINDOWS 1
#elif defined(__linux__)
#define HTP_PLATFORM_PLATFORM_STRING "linux"
#define HTP_PLATFORM_LINUX 1
#ifdef _LP64
#define HTP_PLATFORM_ARCH_64 1
#else /* _LP64 */
#define HTP_PLATFORM_ARCH_32 1
#endif /* _LP64 */
#elif defined(__APPLE__)
#define HTP_PLATFORM_PLATFORM_STRING "osx"
#define HTP_PLATFORM_APPLE 1
#ifdef _LP64
#define HTP_PLATFORM_ARCH_64 1
#else /* _LP64 */
#define HTP_PLATFORM_ARCH_32 1
#endif /* _LP64 */
#endif

#ifndef HTP_PLATFORM_WINDOWS
#define HTP_PLATFORM_WINDOWS 0
#endif
#ifndef HTP_PLATFORM_LINUX
#define HTP_PLATFORM_LINUX 0
#endif
#ifndef HTP_PLATFORM_APPLE
#define HTP_PLATFORM_APPLE 0
#endif

#endif //HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_ROBOT_H_
