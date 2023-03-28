//
// Created by Homin Su on 2022/8/1.
//

#include "robot/utils/logger.h"

#include <unistd.h>

#include "robot/utils/timestamp.h"

namespace htp_platform::logger {

#ifndef NDEBUG
Level log_level = DEBUG;
#else
Level log_level = INFO;
#endif

int log_fd = STDOUT_FILENO;

void set_log_level(Level _level) {
  HTP_PLATFORM_ASSERT((_level >= TRACE && _level <= FATAL) && "log level out of range");
  log_level = _level;
}

void set_log_fd(int _fd) {
  HTP_PLATFORM_ASSERT(_fd > 0 && "file descriptor should greater than 0");
  log_fd = _fd;
}

void timestamp(char *_buf, std::size_t _size) {
  htp_platform::clock::to_string(_buf, _size);
}

} // namespace htp_platform::logger