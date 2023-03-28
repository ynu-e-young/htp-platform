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
// Created by Homin Su on 2022/10/14.
//

#include "strace/stack_trace.h"

#include <csignal>
#include <cstdio>
#include <cstring>
#include <execinfo.h>

#include "get_m_context_addr.h"
#include "robot/utils/logger.h"
#include "strace/strace.h"

namespace strace {
inline namespace v1 {

#ifdef __cplusplus
extern "C" {
#endif

void print_stack(ucontext_t *uc) {
  const int kMaxFrames = 100;
  void *frames[kMaxFrames + 1];

  auto num_frames = backtrace(frames + 1, kMaxFrames);
  auto addr = ::strace::v1::getRaisedAddress(uc);
  if (addr != nullptr) {
    frames[0] = addr;
    auto data = backtrace_symbols(frames, 1);
    if (data != nullptr) ERROR("%s", data[0]);
  }
  auto data = backtrace_symbols(frames + 1, num_frames);
  if (data != nullptr) { for (int i = 0; i < num_frames; ++i) { ERROR("%s", data[i]); }}
}

void sigHandler(int sig, siginfo_t *info, void *ucontext) {
  STRACE_UNUSED(info);

  auto *uc = (ucontext_t *) ucontext;
  auto raised_address = ::strace::v1::getRaisedAddress(uc);

  fprintf(stderr, "Crashed by signal %d (%s)\n", sig, strsignal(sig));
  if (raised_address != nullptr) {
    ERROR("Crashed running the instruction at: %p\n", raised_address);
  }
  if (sig == SIGSEGV || sig == SIGBUS) {
    ERROR("Accessing address: %p\n", info->si_addr);
  }
  print_stack(uc);

  std::raise(sig);
}

void InstallSignalHandlers(int signals[], int length) {
  struct sigaction act{};

  sigemptyset(&act.sa_mask);
  act.sa_flags = SA_NODEFER | SA_RESETHAND | SA_SIGINFO;
  act.sa_sigaction = sigHandler;

  for (int i = 0; i < length; ++i) {
    sigaction(signals[i], &act, nullptr);
  }
}

#ifdef __cplusplus
}
#endif

} // namespace v1
} // namespace strace