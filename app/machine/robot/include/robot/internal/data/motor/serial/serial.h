//
// Created by HominSu on 2021/5/2.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_ROBOT_INTERNAL_DATA_MOTOR_SERIAL_SERIAL_H_
#define HTP_PLATFORM_MACHINE_ROBOT_ROBOT_INTERNAL_DATA_MOTOR_SERIAL_SERIAL_H_

#include <termios.h>  // 包含 POSIX 终端控件定义

#include <vector>
#include <array>
#include <string>

#define MAX_ARDUINO_WAIT_TIME 255

/**
 * @brief 该类实现串口的读写功能
 * @details 通过 linux 串口编程，该类封装好了对应的串口设置以及串口读写函数
 */
class Serial {
 private:
  bool connected_{};  ///< 表示当前串口的连接状态
  struct termios Opt{}; ///< linux 下的串口结构体
  int serial_port_{}; ///< 串口编号

  std::string port_number_ = "/dev/ttyS4"; ///< 串口名称
  int32_t baud_rate_ = 115200; ///< 串口波特率
  int32_t data_bits_ = 8; ///< 数据位
  int32_t stop_bits_ = 1; ///< 停止位
  int32_t parity_ = 0;    ///< 校验位
  bool flow_control_ = false; ///< 流控制
  bool clocal_ = true;  ///< 本地连接

 public:
  explicit Serial(std::string _port_number,
                  int32_t _baud_rate,
                  int32_t _data_bits,
                  int32_t _stop_bits,
                  int32_t _parity,
                  bool _flow_control,
                  bool _clocal);
  ~Serial();

  void ShowTermios() const;
  void SerialOpen();
  void SerialInitial();
  void CloseSerialPort();
  void SerialFlush();

  size_t Read(std::array<unsigned char, 256> &recv_array);
  size_t Write(const std::vector<unsigned char> &_send_array);
  bool IsConnected();

 private:
  void SetParity();
  void SetStopBits();
  void SetDataBits();
  void SetFlowControl();
  void SetClocal();
  void SetLocalModes();
  void SetInputModes();
  void SetOutputModes();
  void SetVminAndVTime();
  void SetBaudRate();
};

#endif //HTP_PLATFORM_MACHINE_ROBOT_ROBOT_INTERNAL_DATA_MOTOR_SERIAL_SERIAL_H_
