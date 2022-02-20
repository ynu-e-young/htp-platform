//
// Created by Homin Su on 2021/8/27.
//

#include <iostream>
#include <fstream>

#include "config.h"

Config::Config() {
  Init();
}

Config::~Config() = default;

/**
 * @brief 初始化基础设置
 */
void Config::Init() {
  nlohmann::json j;
  std::ifstream fp("../conf/basic_setting.json");
  if (!fp.is_open()) {
    std::cerr << "open \"../conf/basic_setting.json\" failed" << std::endl;
    exit(-1);
  }
  fp >> j;
  basic_setting_ = std::move(std::make_unique<Basic::Setting>(j));
}

/**
 * @brief 获取设置
 * @return
 */
const std::unique_ptr<Basic::Setting> &Config::BasicSetting() const {
  return basic_setting_;
}