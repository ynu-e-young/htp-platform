//
// Created by HominSu on 2021/4/27.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_INTERNAL_DATA_MOTOR_GET_CONFIG_GET_SETTING_H
#define HTP_PLATFORM_MACHINE_ROBOT_INTERNAL_DATA_MOTOR_GET_CONFIG_GET_SETTING_H

#include <string>

/**
 * @brief 该类实现从 JSON 文件获取串口相关信息
 */
class GetSerialInfo {
 public:
  std::string serial_port_{}; ///< 串口号
  int baud_rate_{}; ///< 串口波特率
  int data_bits_{}; ///< 数据位
  int stop_bits_{}; ///< 停止位
  int parity_{};  ///< 校验位
  bool flow_control_{}; ///< 流控制
  bool clocal_{}; ///< 本地连接

 public:
  GetSerialInfo();
  ~GetSerialInfo();

 public:
  void LoadJsonInfo();
};

/**
 * @brief 该类实现从 JSON 文件获取机器相关信息
 */
class GetPlatInfo {
 public:
  long ltmp_{}; ///< 轴的数量
  double m_ancher_[6][3]{}; ///< 长度的相关信息
  double m_plate_[6][3]{};  ///< 末端执行器的相关信息

 public:
  GetPlatInfo();
  ~GetPlatInfo();

 public:
  void LoadJsonInfo();
};

class GetUrlInfo {
 public:
  std::string get_url_;
  std::string post_url_;

 public:
  GetUrlInfo();
  ~GetUrlInfo();

 public:
  void LoadJsonInfo();
};

#endif //HTP_PLATFORM_MACHINE_ROBOT_INTERNAL_DATA_MOTOR_GET_CONFIG_GET_SETTING_H
