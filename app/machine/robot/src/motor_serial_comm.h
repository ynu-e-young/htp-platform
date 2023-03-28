//
// Created by HominSu on 2021/5/6.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_SRC_MOTOR_SERIAL_COMM_H_
#define HTP_PLATFORM_MACHINE_ROBOT_SRC_MOTOR_SERIAL_COMM_H_

#include "conf/conf.pb.h"
#include "serial.h"

///< 这是一个常量数组，是用在 CRC 计算里的，是 CRC 计算算法的一部分
const unsigned short KCrcTables[] = {
    0x0000, 0xCC01, 0xD801, 0x1400, 0xF001, 0x3C00, 0x2800, 0xE401,
    0xA001, 0x6C00, 0x7800, 0xB401, 0x5000, 0x9C01, 0x8801, 0x4400
};

/**
 * @brief 控制器通信类
 * @details 封装了与步进电机控制器通信的函数，以及控制器初始化的函数
 */
class MotorSerialComm {
 public:
  unsigned short read_reg_[256]{};  ///< 接收数据缓存区
  bool read_flag_ = false;  ///< 读取标识，表明可以从缓冲区读数据

 private:
  Serial *serial_{};  ///< 串口对象

 public:
  MotorSerialComm();
  ~MotorSerialComm();

 public:
  void InitialCommander(const ::std::shared_ptr<config::Bootstrap> &_bootstrap);
  void CloseCommander();

  // 计算CRC校验码
  unsigned short Crc16Convert(unsigned char *_ptr, unsigned short _len);

  void WriteInstrFirst(unsigned char _equ_addr, unsigned short _reg_addr, unsigned short _reg_data);
  void WriteReadSignal(unsigned char _equ_addr, unsigned short _reg_addr, unsigned short _reg_num);
  void ReadInstr(unsigned char _equ_addr, unsigned short _reg_addr, unsigned short _reg_num);

  // 这三个函数是封装好的控制器读写函数
  void WriteInstr(unsigned char _equ_addr, unsigned short _reg_addr, unsigned short _reg_data);
  void ReadSingleReg(unsigned char _addr, unsigned short _reg_addr, bool *_read_flag);
  unsigned short *ReadMultipleReg(unsigned char _addr,
                                  unsigned short _reg_addr,
                                  unsigned short _reg_num,
                                  bool *_read_flag);
};

#endif //HTP_PLATFORM_MACHINE_ROBOT_SRC_MOTOR_SERIAL_COMM_H_
