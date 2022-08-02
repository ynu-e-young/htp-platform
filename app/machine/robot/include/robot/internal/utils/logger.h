//
// Created by Homin Su on 2022/8/1.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_INTERNAL_UTILS_LOGGER_H_
#define HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_INTERNAL_UTILS_LOGGER_H_

#include "robot/robot.h"

#include <cerrno>
#include <cstdio>
#include <cstring>

#include <sstream>

#define LOG_LEVEL(HTP_PLATFORM) \
  HTP_PLATFORM(TRACE)           \
  HTP_PLATFORM(DEBUG)           \
  HTP_PLATFORM(INFO)            \
  HTP_PLATFORM(WARN)            \
  HTP_PLATFORM(ERROR)           \
  HTP_PLATFORM(FATAL)           \
  //

namespace htp_platform::logger {

enum Level {
#define LOG_NAME(_name) _name,
  LOG_LEVEL(LOG_NAME)
#undef LOG_NAME
};

extern Level log_level;
extern int log_fd;

void set_log_level(Level _level);
void set_log_fd(int _fd);

void timestamp(char *_buf, std::size_t size);

inline const char *level_str(Level _level) {
  const static char *level_str_table[] = {
#define LOG_STR(_name) #_name,
      LOG_LEVEL(LOG_STR)
#undef LOG_STR
  };

  HTP_PLATFORM_ASSERT(_level >= 0 && _level < HTP_PLATFORM_LENGTH(level_str_table));
  return level_str_table[_level];
}

} // namespace htp_platform::logger

#undef LOG_LEVEL

#define LOG_BASE(_log_fd_, _level_, _file_, _line_, _abort_, _formatter_, ...) \
  do {                                                                         \
    char buf[32]{0};                                                           \
    htp_platform::logger::timestamp(buf, sizeof(buf));                         \
    int err = dprintf((_log_fd_), "[%s] [%s] "#_formatter_" - %s:%d\n", htp_platform::logger::level_str((_level_)), buf, ##__VA_ARGS__, strrchr((_file_), '/') + 1, (_line_)); \
    if (err == -1) {                                                           \
      fprintf(stderr, "log failed");                                           \
    }                                                                          \
    if ((_abort_)) {                                                           \
      abort();                                                                 \
    }                                                                          \
  } while (0)                                                                  \
  //

#define LOG_SYS(_log_fd_, _file_, _line_, _abort_, _formatter_, ...) \
  do {                                                               \
    char buf[32]{0};                                                 \
    htp_platform::logger::timestamp(buf, sizeof(buf));               \
    dprintf((_log_fd_), "[%s] [%s] "#_formatter_": %s - %s:%d\n", (_abort_) ? "SYSFA" : "SYSER", buf, ##__VA_ARGS__, strerror(errno), strrchr((_file_), '/') + 1, (_line_)); \
    if ((_abort_)) {                                                 \
      abort();                                                       \
    }                                                                \
  } while (0)                                                        \
  //

#define TRACE(_formatter_, ...) \
  do {                          \
    if (htp_platform::logger::log_level <= htp_platform::logger::Level::TRACE) { \
      LOG_BASE(htp_platform::logger::log_fd, htp_platform::logger::Level::TRACE, __FILE__, __LINE__, false, _formatter_, ##__VA_ARGS__); \
    }                           \
  } while (0)                   \
  //

#define DEBUG(_formatter_, ...) \
  do {                          \
    if (htp_platform::logger::log_level <= htp_platform::logger::Level::DEBUG) { \
      LOG_BASE(htp_platform::logger::log_fd, htp_platform::logger::Level::DEBUG, __FILE__, __LINE__, false, _formatter_, ##__VA_ARGS__); \
    }                           \
  } while (0)                   \
  //

#define INFO(_formatter_, ...) \
  do {                         \
    if (htp_platform::logger::log_level <= htp_platform::logger::Level::INFO) { \
      LOG_BASE(htp_platform::logger::log_fd, htp_platform::logger::Level::INFO, __FILE__, __LINE__, false, _formatter_, ##__VA_ARGS__); \
    }                          \
  } while (0)                  \
  //

#define WARN(_formatter_, ...) \
  do {                         \
    if (htp_platform::logger::log_level <= htp_platform::logger::Level::WARN) { \
      LOG_BASE(htp_platform::logger::log_fd, htp_platform::logger::Level::WARN, __FILE__, __LINE__, false, _formatter_, ##__VA_ARGS__); \
    }                          \
  } while (0)                  \
  //

#define ERROR(_formatter_, ...) LOG_BASE(htp_platform::logger::log_fd, htp_platform::logger::Level::ERROR, __FILE__, __LINE__, false, _formatter_, ##__VA_ARGS__)

#define FATAL(_formatter_, ...) LOG_BASE(htp_platform::logger::log_fd, htp_platform::logger::Level::FATAL, __FILE__, __LINE__, true, _formatter_, ##__VA_ARGS__)

#define SYSERR(_formatter_, ...) LOG_SYS(htp_platform::logger::log_fd, __FILE__, __LINE__, false, _formatter_, ##__VA_ARGS__)

#define SYSFATAL(_formatter_, ...) LOG_SYS(htp_platform::logger::log_fd, __FILE__, __LINE__, true, _formatter_, ##__VA_ARGS__)


#endif //HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_INTERNAL_UTILS_LOGGER_H_
