//
// Created by HominSu on 2021/5/7.
//

#include "motor.h"

#include <stdexcept>

Motor::Motor() = default;

Motor::~Motor() = default;

/**
 * @brief 初始化数据传输串口
 * @details 该函数初始化一个传输串口，用于与控制器的交互，若无法初始化，会抛出异常
 */
void Motor::MotorSerialInit(const ::std::shared_ptr<config::Bootstrap>& _bootstrap) {
  try {
    this->motor_port_.InitialCommander(_bootstrap);
  } catch (std::runtime_error &runtime_error) {
    throw runtime_error;
  } catch (std::exception &e) {
    throw e;
  }
}

/**
 * @brief 关闭控制器
 * @details 急停所有电机确保安全，并关闭串口
 */
void Motor::MotorClose() {
  // 急停全部电机并关闭串口
  MotorProcedureTrigger(0x01, 0, false);
  MotorProcedureTrigger(0x02, 0, false);
  MotorProcedureTrigger(0x03, 0, false);
  MotorProcedureTrigger(0x04, 0, false);
  MotorProcedureTrigger(0x05, 0, false);
  MotorProcedureTrigger(0x06, 0, false);

  this->motor_port_.CloseCommander();
}

/**
 * @brief 初始化控制器
 * @param _addr 16 进制控制器地址
 */
void Motor::MotorSystemInitialization(unsigned char _addr) {
  for (auto &i : kMotorDio) {
    this->motor_Dio_.push_back(i);
  }

  for (auto &i : kMotorBasic) {
    this->motor_basic_.push_back(i);
  }

  for (auto &i : kMotorProcedure) {
    this->motor_procedure_.push_back(i);
  }

  for (auto &i : kMotorZRN) {
    this->motor_zrn_.push_back(i);
  }

  for (auto &i : kMotorStatusReg) {
    this->motor_status_reg_.push_back(i);
  }

  MotorDioInit(_addr);  // 初始化DIO
  MotorBasicParmInit(_addr);  // 设置电机运动方向以及峰值电流

  // TODO:取消PR路径的初始化，降低初始化时间
//  MotorProcedureRegInit(_addr); // 设置PR路径
  MotorZRNRegInit(_addr); // 设置回零信息
}

/**
 * @brief 初始化 DIO
 * @param _addr 16 进制控制器地址
 */
void Motor::MotorDioInit(unsigned char _addr) {
  /**
   * DIO相关功能码:
   * 0x20: 触发命令
   * 0x21: 回零触发;
   * 0x22: 强制急停;
   * 0x23: 正向JOG;
   * 0x24: 方向JOG;
   * 0x25: 正向限位;
   * 0x26: 反向限位;
   * 0x27: 原点信号;
   * 0x28: 路径地址0;
   * 0x29: 路径地址1;
   * 0x2A: 路径地址2;
   * 0x2B: 路径地址3;
  */

  /**
   * 向控制器写入的数据：
   * 0x0088, 0x0027, 0x0025, 0x0026, 0x0000,
   * 0x0000, 0x0000, 0x0000, 0x0000, 0x0000
   */

  this->motor_port_.WriteInstr(_addr, 0x0145, this->motor_Dio_[0]); // 设置DI1
  this->motor_port_.WriteInstr(_addr, 0x0147, this->motor_Dio_[1]); // 设置DI2为0x0027 原点信号
  this->motor_port_.WriteInstr(_addr, 0x0149, this->motor_Dio_[2]); // 设置DI3为0x0025 正向限位
  this->motor_port_.WriteInstr(_addr, 0x014b, this->motor_Dio_[3]); // 设置DI4为0x0026 反向限位
  this->motor_port_.WriteInstr(_addr, 0x014d, this->motor_Dio_[4]); // 设置DI5 无效
  this->motor_port_.WriteInstr(_addr, 0x014f, this->motor_Dio_[5]); // 设置DI6 无效
  this->motor_port_.WriteInstr(_addr, 0x0151, this->motor_Dio_[6]); // 设置DI7 无效

  this->motor_port_.WriteInstr(_addr, 0x0157, this->motor_Dio_[7]); // 设置DO1 无效
  this->motor_port_.WriteInstr(_addr, 0x0159, this->motor_Dio_[8]); // 设置DO2 无效
  this->motor_port_.WriteInstr(_addr, 0x015b, this->motor_Dio_[9]); // 设置DO3 无效
}

/**
 * @brief 设置电机运动方向以及峰值电流
 * @param _addr 16 进制控制器地址
 */
void Motor::MotorBasicParmInit(unsigned char _addr) {
  /**
   * 0x0007 寄存器: 控制电机运动方向, 0 为正, 1 为负
   * 0x0191 寄存器: 电机峰值电流, 范围 0-65, 默认值为 10, 即 1 安
   *
   * 写入的数据
   * 0x0000, 0x000f
   */
  this->motor_port_.WriteInstr(_addr, 0x0007, this->motor_basic_[0]); // 0x0000 设置电机运动方向为正
  this->motor_port_.WriteInstr(_addr, 0x0191, this->motor_basic_[1]); // 0x000f 设置电机峰值电流为1.5安
}

/**
 * @brief 设置 PR 路径
 * @param _addr 16 进制控制器地址
 */
void Motor::MotorProcedureRegInit(unsigned char _addr) {
  /**
   * 该函数中向控制器写入的数据：
   */
  /**
   * PR路径 0:
   * 运动模式路径 0设置为: 位置定位, 不插断, 不重叠, 绝对位置, 不跳转
   * 位置 H: 0x0000
   * 位置 L: 0x16e0 (-59680)DEC
   * 运行速度: 300rpm
   * 加速时间: 100ms/1000rpm
   * 减速时间: 100ms/1000rpm
   * 停顿时间: 0
   * 特殊参数: 0
   */

  /**
   * PR路径 1:
   * 运动模式路径 1设置为: 位置定位, 不插断, 不重叠, 相对位置, 不跳转
   * 位置 H: 0x0000
   * 位置 L: 0x2710 (10000)DEC
   * 运行速度: 300rpm
   * 加速时间: 100ms/1000rpm
   * 减速时间: 100ms/1000rpm
   * 停顿时间: 0
   * 特殊参数: 0
   */

  /**
   * PR路径 2:
   * 运动模式路径 2设置为: 位置定位, 不插断, 不重叠, 相对位置, 不跳转
   * 位置 H: 0xFFFF
   * 位置 L: 0xD8F0 (-10000)DEC
   * 运行速度: 300rpm
   * 加速时间: 100ms/1000rpm
   * 减速时间: 100ms/1000rpm
   * 停顿时间: 0
   * 特殊参数: 0
   *
   * 后面 PR 路径均为 0
   */

  // 向前8个PR路径写入预设
  unsigned short index = 0;
  for (auto &i : this->motor_procedure_) {
    this->motor_port_.WriteInstr(_addr, 0x6200 + index, i);
    ++index;
  }
}

/**
 * @brief 设置回零的信息
 * @param _addr 控制器地址
 */
void Motor::MotorZRNRegInit(unsigned char _addr) {
  /**
   * 限位、JOG和急停功能：
   * 0x6000 寄存器: PR 控制参数 Bit1: 软件限位是否有效; 0x6006 寄存器: 正限位 H 软件限位正向位置高位; 0x6007 寄存器: 正限位 L 软件限位正向位置高位; 0x6008 寄存器: 负限位 H 软件限位负向位置高位; 0x6009 寄存器: 负限位 L 软件限位负向位置高位
   *
   * 0x600a 寄存器: 回零模式; Bit0: 回零方向 =0: 反向, =1:正向; Bit1: 回零后是否移动到指定位置 =0: 否, =1:是; Bit2：回零模式 =0: 限位回零, =1:原点回零
   *
   * 原点信号在坐标轴上的位置: 0x600b 寄存器: 零位位置 H 高16位; 0x600c 寄存器: 零位位置 L 低16位
   *
   * 回零后，电机移动到指定位置停止。若回零模式 bit1 使能，则回零后移动到该绝对位置; 0x600d 寄存器: 回零停止位置 H 高 16 位; 0x600e 寄存器: 回零停止位置 H 低 16 位
   *
   * 0x600f 寄存器: 回零高速 回零的第一段速度，单位 rpm
   * 0x6010 寄存器: 回零低速 回零的第二段速度，单位 rpm
   * 0x6011 寄存器: 回零加速时间 回零的加速度，单位 ms/1000rpm
   * 0x6012 寄存器: 回零减速时间 回零的减速度，单位 ms/1000rpm
   *
   * 0x6016 寄存器: 限位急停时间 限位后的减速时间, 单位: ms
   * 0x6017 寄存器: STOP 急停时间 急停后的减速时间, 单位: ms
   */

  /**
   * 写入的数据: 0x0000, 0x001e, 0x8480, 0xffe1, 0x7b80, 0x0007, 0x0000, 0x0000, 0xFFF8, 0x37D0, 0x0064, 0x000a, 0x0064, 0x0064, 0x0064, 0x0032
   */

  // 向指定寄存器写入数据
  this->motor_port_.WriteInstr(_addr, 0x6000, this->motor_zrn_[0]);
  this->motor_port_.WriteInstr(_addr, 0x6006, this->motor_zrn_[1]); // 设置正限位H(软件限位正向位置高位) 0x001e
  this->motor_port_.WriteInstr(_addr, 0x6007, this->motor_zrn_[2]); // 设置正限位L(软件限位正向位置低位) 0x8480 (2000000)DEC
  this->motor_port_.WriteInstr(_addr, 0x6008, this->motor_zrn_[3]); // 设置负限位H(软件限位反向位置高位) 0xffe1
  this->motor_port_.WriteInstr(_addr, 0x6009, this->motor_zrn_[4]); // 设置负限位L(软件限位反向位置低位) 0x7b80 (-2000000)DEC
  this->motor_port_.WriteInstr(_addr, 0x600a, this->motor_zrn_[5]); // 设置回零模式 回零方向: 正向, 回零后移动到指定位置, 回零模式: 原点回零
  this->motor_port_.WriteInstr(_addr, 0x600b, this->motor_zrn_[6]); // 零位位置H 0x0000
  this->motor_port_.WriteInstr(_addr, 0x600c, this->motor_zrn_[7]); // 零位位置L 0x0000
  this->motor_port_.WriteInstr(_addr, 0x600d, this->motor_zrn_[8]); // 回零停止位置H 0xFFF8
  this->motor_port_.WriteInstr(_addr, 0x600e, this->motor_zrn_[9]); // 回零停止位置L 0x37D0, 即回零后移动到-510000(Dec)
  this->motor_port_.WriteInstr(_addr, 0x600f, this->motor_zrn_[10]);  // 回零的第一段速度 100rpm
  this->motor_port_.WriteInstr(_addr, 0x6010, this->motor_zrn_[11]);  // 回零的第二段速度 10rpm
  this->motor_port_.WriteInstr(_addr, 0x6011, this->motor_zrn_[12]);  // 回零的加速度 100ms/1000rpm
  this->motor_port_.WriteInstr(_addr, 0x6012, this->motor_zrn_[13]);  // 回零的减速度 100ms/1000rpm
  this->motor_port_.WriteInstr(_addr, 0x6016, this->motor_zrn_[14]);  // 限位后的减速时间 100ms
  this->motor_port_.WriteInstr(_addr, 0x6017, this->motor_zrn_[15]);  // 急停后的减速时间 50ms
}

/**
 * 触发回零
 * @param _addr 控制器地址
 * @param _trigger 是否触发
 */
void Motor::MotorZRN(unsigned char _addr, bool _trigger) {
  // 向0x6002寄存器写入0x020: 回零;
  if (_trigger) {
    this->motor_port_.WriteInstr(_addr, 0x6002, 0x20);
  }
}

/**
 * @brief 当前位置手动设零
 * @param _addr 控制器地址
 * @param _trigger 是否触发
 */
void Motor::MotorZeroSet(unsigned char _addr, bool _trigger) {
  // 向0x6002寄存器写入0x021: 当前位置手动设零;
  if (_trigger) {
    this->motor_port_.WriteInstr(_addr, 0x6002, 0x21);
  }
}

/**
 * @brief 向控制器写入 PR 路径信息
 * @param _addr 控制器地址
 * @param _motion_mode 运动模式
 * @param _pulse 运动脉冲数
 * @param _speed 运动速度
 * @param _procedure_path PR 路径
 */
void Motor::MotorWriteProcedureReg(unsigned char _addr,
                                   unsigned short _motion_mode,
                                   int _pulse,
                                   unsigned short _speed,
                                   unsigned char _procedure_path) {
  // 写入运动模式
  this->motor_port_.WriteInstr(_addr, 0x6200 + 0 + _procedure_path * kMotorProcedureLength, _motion_mode);

  // 写入Pulse的高16位
  this->motor_port_.WriteInstr(_addr,
                               0x6200 + 1 + _procedure_path * kMotorProcedureLength,
                               (_pulse & 0xffff0000) >> 16);
  // 写入Pulse的低16位
  this->motor_port_.WriteInstr(_addr,
                               0x6200 + 2 + _procedure_path * kMotorProcedureLength,
                               _pulse & 0x0000ffff);
  // 写入速度
  this->motor_port_.WriteInstr(_addr, 0x6200 + 3 + _procedure_path * kMotorProcedureLength, _speed);
}

/**
 * @brief 触发 PR 路径
 * @param _addr 控制器地址
 * @param _procedure_path PR 路径
 * @param _trigger true 触发, false 急停
 */
void Motor::MotorProcedureTrigger(unsigned char _addr, unsigned char _procedure_path, bool _trigger) {
  if (_trigger) {
    this->motor_port_.WriteInstr(_addr, 0x6002, 0x10 + _procedure_path); // P段定位, 触发指定路径
  } else {
    this->motor_port_.WriteInstr(_addr, 0x6002, 0x40); // 急停
  }
}

/**
 * @brief 读取单个寄存器
 * @param _addr 控制器地址
 * @param _reg_addr 寄存器地址
 * @return 读取到的值
 */
unsigned short Motor::MotorReadSingleReg(unsigned char _addr,
                                         unsigned short _reg_addr) {
  bool read_flag = false;

  this->motor_port_.ReadSingleReg(_addr, _reg_addr, &read_flag);
  // 如果读到了数据
  if (read_flag) {
    return this->motor_port_.read_reg_[0];
  }
  return 0;
}

/**
 * @brief 读取多个连续寄存器
 * @param _addr 控制器地址
 * @param _reg_addr 寄存器起始地址
 * @param _reg_num 读取寄存器数量
 * @return 读取到的值
 */
unsigned short *Motor::MotorReadMultipleReg(unsigned char _addr, unsigned char _reg_addr, unsigned short _reg_num) {
  bool read_flag = false;
  unsigned short *recv_buffer;

  recv_buffer = this->motor_port_.ReadMultipleReg(_addr, _reg_addr, _reg_num, &read_flag);
  // 如果读到了数据
  if (read_flag) {
    return recv_buffer;
  }
  return nullptr;
}

/**
 * @brief 获取控制器寄存器的参数
 * @details 读取单个控制器的电机的运行状态、命令位置高位、命令位置低位、电机位置高位、电机位置低位、触发寄存器
 * @param _addr 读取的控制器地址
 */
void Motor::MotorGetAllReg(unsigned char _addr) {
  // 读取0x1003寄存器
  this->motor_status_[_addr - 1][0] = MotorReadSingleReg(_addr, this->motor_status_reg_.at(0));

  // 读取0x602a, 0x602b, 0x602c, 0x602d寄存器
  this->motor_status_[_addr - 1][1] = MotorReadSingleReg(_addr, this->motor_status_reg_.at(1));
  this->motor_status_[_addr - 1][2] = MotorReadSingleReg(_addr, this->motor_status_reg_.at(2));
  this->motor_status_[_addr - 1][3] = MotorReadSingleReg(_addr, this->motor_status_reg_.at(3));
  this->motor_status_[_addr - 1][4] = MotorReadSingleReg(_addr, this->motor_status_reg_.at(4));

//  MotorReadMultipleReg(_addr, this->motor_status_reg_.at(1), 4);

  // 读取0x6002寄存器
  //this->motor_reg_status_[_addr][5] = MotorReadSingleReg(_addr, motor_status_reg_.at(5));
}

/**
 * @brief 获取控制器 0x1003 寄存器的参数
 * @details 读取单个控制器的电机的运行状态
 * @param _addr 读取的控制器地址
 */
void Motor::MotorGetStatusReg(unsigned char _addr) {
  // 读取0x1003寄存器
  this->motor_status_[_addr - 1][0] = MotorReadSingleReg(_addr, this->motor_status_reg_.at(0));
}

/**
 * @brief 获取控制器 0x602a, 0x602b, 0x602c, 0x602d 寄存器的参数
 * @details 读取单个控制器的电机的命令位置高位、命令位置低位、电机位置高位、电机位置低位、触发寄存器
 * @param _addr 读取的控制器地址
 */
void Motor::MotorGetPosReg(unsigned char _addr) {
  // 读取0x602a, 0x602b, 0x602c, 0x602d寄存器
  this->motor_status_[_addr - 1][1] = MotorReadSingleReg(_addr, this->motor_status_reg_.at(1));
  this->motor_status_[_addr - 1][2] = MotorReadSingleReg(_addr, this->motor_status_reg_.at(2));
  this->motor_status_[_addr - 1][3] = MotorReadSingleReg(_addr, this->motor_status_reg_.at(3));
  this->motor_status_[_addr - 1][4] = MotorReadSingleReg(_addr, this->motor_status_reg_.at(4));

  //  MotorReadMultipleReg(_addr, this->motor_status_reg_.at(1), 4);
}



