//
// Created by Homin Su on 2021/9/10.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_ROBOT_INTERNAL_CONF_CONFIG_H_
#define HTP_PLATFORM_MACHINE_ROBOT_ROBOT_INTERNAL_CONF_CONFIG_H_

#include <memory>

#include "basic_setting.h"

class Config {
 private:
  std::unique_ptr<Basic::Setting> basic_setting_;

 public:
  ~Config();

  /**
   * @brief 单件模式
   * @return Config *
   */
  static Config *Get() {
    static Config c;
    return &c;
  }

 private:
  Config();
  void Init();

 public:
  [[nodiscard]] const std::unique_ptr<Basic::Setting> &BasicSetting() const;
};

#endif //HTP_PLATFORM_MACHINE_ROBOT_ROBOT_INTERNAL_CONF_CONFIG_H_
