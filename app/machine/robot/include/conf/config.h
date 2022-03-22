//
// Created by Homin Su on 2022/3/21.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_CONF_CONFIG_H_
#define HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_CONF_CONFIG_H_

#include "google/protobuf/message.h"
#include "google/protobuf/util/json_util.h"

#include <cstdio>

#include <exception>

class Config {
 public:
  void Load(const char *_path);
  void Scan(::google::protobuf::Message *_message);

 private:
  void read(FILE *_in);

  ::std::string get_str() {
    return {buffer_.begin(), buffer_.end()};
  }

 private:
  std::vector<char> buffer_;
};

inline void Config::Load(const char *_path) {
  FILE *input;
#if defined(_MSC_VER)
  fopen_s(&input, _path, "r");
#else
  input = fopen(_path, "r");
#endif
  if (input == nullptr) { exit(EXIT_FAILURE); }
  read(input);
  fclose(input);
}

inline void Config::read(FILE *_in) {
  char buf[65536];
  if (!buffer_.empty()) { buffer_.clear(); }
  while (true) {
    size_t n = fread(buf, 1, sizeof(buf), _in);
    if (n == 0) { break; }
    buffer_.insert(buffer_.end(), buf, buf + n);
  }
}

void Config::Scan(::google::protobuf::Message *_message) {
  if (!google::protobuf::util::JsonStringToMessage(get_str(), _message).ok()) {
    throw ::std::runtime_error("Parse json data to message failed");
  }
}

#endif //HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_CONF_CONFIG_H_
