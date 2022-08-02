//
// Created by Homin Su on 2022/8/1.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_INTERNAL_UTILS_TIMESTAMP_H_
#define HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_INTERNAL_UTILS_TIMESTAMP_H_

#include <cstdio>

#include <chrono>

namespace htp_platform {

using nanoseconds = std::chrono::nanoseconds;
using microseconds = std::chrono::microseconds;
using milliseconds = std::chrono::milliseconds;
using seconds = std::chrono::seconds;
using minutes = std::chrono::minutes;
using hours = std::chrono::hours;

using system_clock = std::chrono::system_clock;
using time_point = std::chrono::time_point<system_clock, nanoseconds>;

namespace clock {

inline time_point now() { return system_clock::now(); }

inline time_point after_now(nanoseconds _interval) { return now() + _interval; }

inline time_point before_now(nanoseconds _interval) { return now() - _interval; }

inline int to_string(char *_data, std::size_t _size) {
  auto n = system_clock::now();
  auto t_s = std::chrono::duration_cast<seconds>(n.time_since_epoch());
  auto t_ms = std::chrono::duration_cast<milliseconds>(n.time_since_epoch());
  auto ms = static_cast<int>((t_ms - t_s).count());
  auto t = system_clock::to_time_t(n);

  tm tm_time{};
  gmtime_r(&t, &tm_time);

  return snprintf(_data,
                  _size,
                  "%4d-%02d-%02dT%02d:%02d:%02d.%03dZ",
                  tm_time.tm_year + 1900,
                  tm_time.tm_mon + 1,
                  tm_time.tm_mday,
                  tm_time.tm_hour,
                  tm_time.tm_min,
                  tm_time.tm_sec,
                  ms);
}

} // namespace clock

} // namespace htp_platform

#endif //HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_INTERNAL_UTILS_TIMESTAMP_H_
