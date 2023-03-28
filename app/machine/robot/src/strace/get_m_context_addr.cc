// MIT License
//
// Copyright (c) 2022 HominSu
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

//
// Created by Homin Su on 2022/10/15.
//

#include "get_m_context_addr.h"

#if STRACE_APPLE && STRACE_aarch64
#include <mach/mach.h>
#endif

namespace strace {
inline namespace v1 {

void *getRaisedAddress(ucontext_t *uc) {
  // @formatter:off
#if STRACE_APPLE && !defined(MAC_OS_X_VERSION_10_6)
  #if defined(STRACE_ARCH_64)
  return (void *) uc->uc_mcontext->__ss.__rip;
#elif defined(STRACE_ARCH_32)
  return (void *) uc->uc_mcontext->__ss.__eip;
  #else
  return (void *) uc->uc_mcontext->__ss.__srr0;
#endif
#elif STRACE_APPLE && defined(MAC_OS_X_VERSION_10_6)
  #if defined(_STRUCT_X86_THREAD_STATE64) && !STRACE_i386
  return (void *) uc->uc_mcontext->__ss.__rip;
#elif STRACE_i386
  return (void *) uc->uc_mcontext->__ss.__eip;
#else
  return (void *) arm_thread_state64_get_pc(uc->uc_mcontext->__ss);
#endif
#elif STRACE_LINUX
#if STRACE_i386 || ((STRACE_X86_64 || STRACE_x86_64) && STRACE_ARCH_32)
  return (void *) uc->uc_mcontext.gregs[14];
#elif STRACE_X86_64 || STRACE_x86_64
  return (void *) uc->uc_mcontext.gregs[16];
#elif STRACE_ia64
  return (void *) uc->uc_mcontext.sc_ip;
  #elif STRACE_arm
  return (void *) uc->uc_mcontext.arm_pc;
  #elif !STRACE_aarch64
  return (void *) uc->uc_mcontext.pc;
  #else
  STRACE_UNUSED(uc);
  return nullptr;
#endif
#else
  STRACE_UNUSED(uc);
  return nullptr;
#endif
  // @formatter:on
}

} // namespace v1
} // namespace strace