//
// Created by HominSu on 2021/4/27.
//

#include "get_setting.h"

#include <json/json.h>

#include <iostream>
#include <sstream>
#include <fstream>

GetSerialInfo::GetSerialInfo() = default;

GetSerialInfo::~GetSerialInfo() = default;

/**
 * @brief 从 JSON 文件读取串口相关设置
 */
void GetSerialInfo::LoadJsonInfo() {
  Json::Value root, serial_info;
  JSONCPP_STRING errs;
  Json::CharReaderBuilder builder;
  builder["collectComments"] = true;

  std::ifstream ifs;
  ifs.open("/data/conf/serial_port_config.json");

  if (!ifs.is_open()) {
    std::cout << "Error opening file\n";
    exit(1);
  }

  if (!parseFromStream(builder, ifs, &root, &errs)) {
    std::cout << errs << std::endl;
    exit(2);
  }

  serial_info = root["serial_info"];

  this->serial_port_ = serial_info["serial_port"].asString();
  this->baud_rate_ = serial_info["baud_rate"].asInt();
  this->data_bits_ = serial_info["data_bits"].asInt();
  this->stop_bits_ = serial_info["stop_bits"].asInt();
  this->parity_ = serial_info["parity"].asInt();
  this->flow_control_ = serial_info["flow_control"].asBool();
  this->clocal_ = serial_info["clocal"].asBool();
}

GetPlatInfo::GetPlatInfo() = default;

GetPlatInfo::~GetPlatInfo() = default;

/**
 * @brief 从 JSON 文件读取机器相关设置
 */
void GetPlatInfo::LoadJsonInfo() {
  Json::Value root, plat_info;
  JSONCPP_STRING errs;
  Json::CharReaderBuilder builder;
  builder["collectComments"] = true;

  std::ifstream ifs;
  ifs.open("/data/conf/plat_info.json");

  if (!ifs.is_open()) {
    std::cout << "Error opening file\n";
    exit(1);
  }

  if (!parseFromStream(builder, ifs, &root, &errs)) {
    std::cout << errs << std::endl;
    exit(2);
  }

  plat_info = root["plat_info"];

  this->ltmp_ = plat_info["ltmp"].asInt();

  const Json::Value m_ancher_array = plat_info["m_ancher"];
  double *tmp = &this->m_ancher_[0][0];

  for (const auto &i : m_ancher_array) {
    for (const auto &j : i) {
      *tmp++ = j.asDouble();
    }
  }

  const Json::Value m_plate_array = plat_info["m_plate"];
  tmp = &this->m_plate_[0][0];
  for (const auto &i : m_plate_array) {
    for (const auto &j : i) {
      *tmp++ = j.asDouble();
    }
  }
}

GetUrlInfo::GetUrlInfo() = default;

GetUrlInfo::~GetUrlInfo() = default;

void GetUrlInfo::LoadJsonInfo() {
  Json::Value root, GetUrl, PostUrl;
  JSONCPP_STRING errs;
  Json::CharReaderBuilder builder;
  builder["collectComments"] = true;

  std::ifstream ifs;
  ifs.open("/data/conf/url_config.json");

  if (!ifs.is_open()) {
    std::cout << "Error opening file\n";
    exit(1);
  }

  if (!parseFromStream(builder, ifs, &root, &errs)) {
    std::cout << errs << std::endl;
    exit(2);
  }

  std::stringstream ss;

  PostUrl = root["post_url"];
  ss << PostUrl["ip"].asString();
  ss << ':';
  ss << PostUrl["port"].asString();
  ss << '/';
  ss << PostUrl["api"].asString();
  this->post_url_ = ss.str();
  std::cout << "url: " << ss.str() << std::endl;

  ss.str("");

  GetUrl = root["get_url"];
  ss << GetUrl["ip"].asString();
  ss << ':';
  ss << GetUrl["port"].asString();
  ss << '/';
  ss << GetUrl["api"].asString();
  this->get_url_ = ss.str();
}


