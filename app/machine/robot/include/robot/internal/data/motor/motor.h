//
// Created by HominSu on 2021/5/7.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_ROBOT_INTERNAL_DATA_MOTOR_MOTOR_H
#define HTP_PLATFORM_MACHINE_ROBOT_ROBOT_INTERNAL_DATA_MOTOR_MOTOR_H

#include "motor_serial_comm.h"
#include "motor_config/motor_path.h"
#include "motor_config/motor_status_reg.h"

#include <vector>

const int kPulsePerRevolution = 2000;    ///< 2000 脉冲一转
const int kMaxMotor = 6;    ///< 控制器数量
const int kMotorStatusLength = 6;    ///< 每一路控制器读取的寄存器个数
const int kMotorProcedureLength = 8;    ///< 设置单个 PR 路径所需的所有寄存器的数目

/**
 * @brief 封装了对控制器的基础操作
 */
class Motor {
 public:
  unsigned short motor_status_[kMaxMotor][kMotorStatusLength]{};  ///< 记录所有电机的所有状态信息

 private:
  MotorSerialComm motor_port_;  ///< 串口号
  std::vector<unsigned short> motor_Dio_; ///< 保存控制器 Dio 设置的具体参数
  std::vector<unsigned short> motor_basic_; ///< 保存控制器的运动方向以及电机的电流
  std::vector<unsigned short> motor_procedure_; ///< 保存预设的 PR 路径
  std::vector<unsigned short> motor_zrn_; ///< 保存回零的具体参数
  std::vector<unsigned short> motor_status_reg_;  ///< 保存从控制器寄存器读取的信息

 public:
  Motor();
  ~Motor();

 public:
  void MotorSerialInit(const ::std::shared_ptr<config::Bootstrap>& _bootstrap);
  void MotorClose();

  // 一系列初始化函数
  void MotorSystemInitialization(unsigned char _addr);
  void MotorDioInit(unsigned char _addr);
  void MotorBasicParmInit(unsigned char _addr);
  void MotorProcedureRegInit(unsigned char _addr);
  void MotorZRNRegInit(unsigned char _addr);

  // 回零和设置零位
  void MotorZRN(unsigned char _addr, bool _trigger);
  void MotorZeroSet(unsigned char _addr, bool _trigger);

  // 写入PR路径和触发PR路径
  void MotorWriteProcedureReg(unsigned char _addr,
                              unsigned short _motion_mode,
                              int _pulse,
                              unsigned short _speed,
                              unsigned char _procedure_path);
  void MotorProcedureTrigger(unsigned char _addr, unsigned char _procedure_path, bool _trigger);

  // 读取寄存器
  unsigned short MotorReadSingleReg(unsigned char _addr,
                                    unsigned short _reg_addr);
  unsigned short *MotorReadMultipleReg(unsigned char _addr, unsigned char _reg_addr, unsigned short _reg_num);
  void MotorGetAllReg(unsigned char _addr);
  void MotorGetStatusReg(unsigned char _addr);
  void MotorGetPosReg(unsigned char _addr);
};

#endif //HTP_PLATFORM_MACHINE_ROBOT_ROBOT_INTERNAL_DATA_MOTOR_MOTOR_H
