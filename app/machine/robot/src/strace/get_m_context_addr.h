//
// Created by Homin Su on 2023/3/28.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_SRC_STRACE_GET_M_CONTEXT_ADDR_H_
#define HTP_PLATFORM_MACHINE_ROBOT_SRC_STRACE_GET_M_CONTEXT_ADDR_H_

#include "strace/strace.h"

#if STRACE_APPLE
#include <sys/ucontext.h>
#else
#include <ucontext.h>
#endif

namespace strace {
inline namespace v1 {

void *getRaisedAddress(ucontext_t *uc);

} // namespace v1
} // namespace strace


#endif //HTP_PLATFORM_MACHINE_ROBOT_SRC_STRACE_GET_M_CONTEXT_ADDR_H_
