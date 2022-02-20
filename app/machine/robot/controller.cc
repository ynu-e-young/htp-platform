//
// Created by HominSu on 2021/5/8.
//

#include "controller.h"

#include "internal/conf/config.h"
#include "internal/data/grpc/move_done_client.h"

#include <iostream>
#include <string>
#include <thread>
#include <chrono>

Controller::Controller() = default;

Controller::~Controller() {
  if (system_status_) {
    ShutDownSystem();
  }
}

/**
 * @brief 动作队列是否为空
 * @return bool
 */
bool Controller::motion_coord_empty() {
  // 读取 check_coord，共享锁
  std::shared_lock<std::shared_mutex> lock(mutex_);
  return motion_coord_.empty();
}

/**
 * @brief 弹出动作队列队头的坐标
 * @return 坐标
 */
std::array<double, 6> Controller::motion_coord_pop() {
  // 弹出 check_coord，独占锁
  std::unique_lock<std::shared_mutex> lock(mutex_);
  auto coord = motion_coord_.front();
  motion_coord_.pop();

  return coord;
}

/**
 * @brief 加入坐标到动作队列
 * @param _coord 六位 double array
 */
void Controller::motion_coord_push(const std::array<double, 6> &_coord) {
  // 写入 check_coord，独占锁
  std::unique_lock<std::shared_mutex> lock(mutex_);
  motion_coord_.push(_coord);
}

/**
 * @brief 检查队列是否为空
 * @return bool
 */
bool Controller::check_coord_empty() {
  // 读取 check_coord，共享锁
  std::shared_lock<std::shared_mutex> lock(mutex_);
  return check_coord_.empty();
}

/**
 * @brief 弹出检查队列队头的坐标
 * @return 包含检查动作名称和坐标的结构体
 */
Coord Controller::check_coord_pop() {
  // 弹出 check_coord，独占锁
  std::unique_lock<std::shared_mutex> lock(mutex_);
  auto coord = check_coord_.front();
  check_coord_.pop();

  return coord;
}

/**
 * @brief 加入坐标到检查队列
 * @param _coord 包含检查动作名称和坐标的结构体
 */
void Controller::check_coord_push(const Coord &_coord) {
  // 写入 check_coord，独占锁
  std::unique_lock<std::shared_mutex> lock(mutex_);
  check_coord_.push(_coord);
}

/**
 * @brief 读取电机状态
 * @return 以 6 位 bitset 数组的形式返回
 */
std::array<std::bitset<7>, 6> Controller::motor_status() {
  // 读取 motor_status_ 寄存器，共享锁
  std::shared_lock<std::shared_mutex> lock(mutex_);
  return motor_status_;
}

/**
 * @brief 读取当前电机的实际位
 * @return 6 位 int 数组
 */
std::array<int, 6> Controller::motor_real_pos() {
  // 读取 motor_real_pos_ 寄存器，共享锁
  std::shared_lock<std::shared_mutex> lock(mutex_);
  return motor_real_pos_;
}

/**
 * @brief 读取当前电机的指令位置
 * @return 6 位 int 数组
 */
std::array<int, 6> Controller::motor_expected_pos() {
  // 读取 motor_expected_pos_ 寄存器，共享锁
  std::shared_lock<std::shared_mutex> lock(mutex_);
  return motor_expected_pos_;
}

/**
 * @brief 读取回零状态
 * @return bool
 */
bool Controller::zrn() {
  // 读取 zrn_ ，共享锁
  std::shared_lock<std::shared_mutex> lock(mutex_);
  return zrn_;
}

/**
 * @brief 设置回零状态
 * @param _zrn bool 回零状态
 */
void Controller::set_zrn(bool _zrn) {
  // 写入 motor_real_pos_ 寄存器，独占锁
  std::unique_lock<std::shared_mutex> lock(mutex_);
  zrn_ = _zrn;
}

/**
 * @brief 初始化机器
 * @details 初始化串口然后初始化机器的每一轴
 */
void Controller::InitialSystem() {
  // 初始化串口
  do {
    try {
      motors_.MotorSerialInit();
      break;
    } catch (std::runtime_error &runtime_error) {
      std::cout << runtime_error.what() << std::endl;
    } catch (std::exception &e) {
      std::cout << e.what() << std::endl;
    }

    // 每隔一秒尝试重新连接串口
    std::this_thread::sleep_for(std::chrono::seconds(1));
  } while (is_running());

  // 初始化6个控制器
  motors_.MotorSystemInitialization(0x01);
  motors_.MotorSystemInitialization(0x02);
  motors_.MotorSystemInitialization(0x03);
  motors_.MotorSystemInitialization(0x04);
  motors_.MotorSystemInitialization(0x05);
  motors_.MotorSystemInitialization(0x06);

  // 初始化坐标结算
  cal_len_.OnInit();

  // 设置零位参考
  SetZero();

  system_status_ = true;
}

/**
 * @brief 关闭系统
 */
void Controller::ShutDownSystem() {
  // 关闭全部轴
  motors_.MotorClose();

  system_status_ = false;
}

/**
 * @brief 手动设置零位
 */
void Controller::SetZero() {
  long size;
  double m_xyz[] = {0, 0, 0};
  double m_ges[] = {0, 0, 0};
  double len[6]{0};

  if (0 == cal_len_.OnPos(m_xyz, m_ges, len, size)) {
    int index = 0;
    for (const auto &i: len) {
      zero_[index] = i;
      ++index;
    }
  }
}

/**
 * @brief 计算脉冲数
 * @details 该函数将输入的绳子的长度和零位参考长度相减，并计算出步进电机控制器实际需要接受的脉冲数
 * @param _len 位置
 * @param _zero 零位位置
 * @return 脉冲数
 */
int Controller::CalMotorPulse(double _len, double _zero) {
  int alex = static_cast<int>(_len - _zero) * kPulsePerRevolution;
  return (0 - alex) + kOffset;
}

/**
 * @brief 执行坐标和姿态
 * @details 该函数将具体的移动位置和移动姿态写入到步进电机控制器的 PR0 中，并以120rpm的速度移动到目标位置
 * @param _m_xyz 运动坐标
 * @param _m_ges 运动姿态
 */
void Controller::Move(double _m_xyz[], double _m_ges[]) {
  long size;
  double len[6]{0};
  long ret = cal_len_.OnPos(_m_xyz, _m_ges, len, size);

#if DEBUG
  std::cout << "Move: [" << _m_xyz[0] << ", " << _m_xyz[1] << ", " << _m_xyz[2] << ", " << _m_ges[0] << ", "
            << _m_ges[1] << "], delay: " << _m_ges[2] << std::endl;
#endif

  if (0 != ret) {
    std::cout << "ret: " << ret << std::endl;
    return;
  }

#if DEBUG
  std::cout << "pluse: [";
#endif
  int index = 0;
  for (const auto &i: len) {
    auto pulse = CalMotorPulse(i, zero_[index]);
#if DEBUG
    std::cout << pulse << ", ";
#endif
    motor_expected_pos_[index] = pulse;

    // 写入PR0
    motors_.MotorWriteProcedureReg(0x01 + index,
                                   0x0001,
                                   pulse,
                                   120,
                                   0);
    // 触发PR0
    motors_.MotorProcedureTrigger(0x01 + index, 0, true);

    ++index;
  }
#if DEBUG
  std::cout << "]" << std::endl;
#endif
}

/**
 * @brief 获取单个控制器的信息
 * @param _addr 控制器地址
 */
void Controller::GetMotorAllStatus(unsigned char _addr) {
  motors_.MotorGetAllReg(_addr);

  for (int index = 0; index < 6; ++index) {
    motor_reg_status_[_addr - 1][index] = motors_.motor_status_[_addr - 1][index];
  }
}

void Controller::GetMotorStatusReg(unsigned char _addr) {
  motors_.MotorGetStatusReg(_addr);

  motor_reg_status_[_addr - 1][0] = motors_.motor_status_[_addr - 1][0];
}

void Controller::GetMotorPosReg(unsigned char _addr) {
  motors_.MotorGetPosReg(_addr);

  for (int index = 1; index < 5; ++index) {
    motor_reg_status_[_addr - 1][index] = motors_.motor_status_[_addr - 1][index];
  }
}

/**
 * @brief 获取全部控制器信息
 * @details 轮询 6 个控制器，获取每个控制器的 0x1003 寄存器
 * @param _option 读取的寄存器标签
 */
void Controller::GetAllMotorStatus(const int &_option) {
  if (kMotorStatus == _option) {
    for (unsigned char addr = 0x01; addr <= 0x06; ++addr) {
      GetMotorStatusReg(addr);
    }
  } else if (kMotorPos == _option) {
    for (unsigned char addr = 0x01; addr <= 0x06; ++addr) {
      GetMotorPosReg(addr);
    }
  } else if (kMotorAll == _option) {
    for (unsigned char addr = 0x01; addr <= 0x06; ++addr) {
      GetMotorAllStatus(addr);
    }
  }
}

/**
 * @brief 检查回零是否完成
 * @details 该函数通过对控制器的 0x1003 寄存器进行读取，获取到详细当前控制器的回零状态以及路径执行状态，用来判断当前动作知否执行完成
 * @return bool 返回值为 true，说明回零动作已完成，返回值为 false，说明回零动作未完成
 */
bool Controller::ZRNDone() {
  int flag = 0; // 存储完成条件

  // 更新控制器信息数组
  GetAllMotorStatus(kMotorStatus);

  {
    // 写入motor_status_寄存器，独占锁
    std::unique_lock<std::shared_mutex> lock(mutex_);
    // 遍历控制器信息数组
    for (int index = 0; index < 6; ++index) {
      motor_status_[index] = motor_reg_status_[index][0];  // 转换为bitset，便于读取信息
//    std::cout << "addr: " << index << " " << motion_status << std::endl;
    }
  }

  for (int index = 0; index < 6; ++index) {
    if (1 == motor_status_[index][6]) {
      ++flag;
    }
  }

  if (6 == flag) {
    return true;
  }
  return false;
}

/**
 * @brief 检查动作是否执行完成
 * @details 该函数通过对控制器的 0x1003 寄存器进行读取，获取到详细当前控制器的回零状态以及路径执行状态，用来判断当前动作知否执行完成
 * @return bool 返回值为 true，说明回零动作已经执行完成，返回值为 false，说明动作未完成
 */
bool Controller::MoveDone() {
  int flag = 0; // 存储完成条件

  // 更新控制器信息数组
  GetAllMotorStatus(kMotorStatus);

  {
    // 写入 motor_status_ 寄存器，独占锁
    std::unique_lock<std::shared_mutex> lock(mutex_);
    // 遍历控制器信息数组
    for (int index = 0; index < 6; ++index) {
      motor_status_[index] = motor_reg_status_[index][0];  // 转换为bitset，便于读取信息
//    std::cout << "addr: " << index << " " << motion_status << std::endl;
    }
  }

  for (int index = 0; index < 6; ++index) {
    if (1 == motor_status_[index][5]) {
      ++flag;
    }
  }

  if (6 == flag) {
    return true;
  }
  return false;
}

/**
 * @brief 电机回零
 */
void Controller::MotorZRN() {
  // 初始化回零路径
//  for (unsigned char index = 0x01; index <= 0x06; ++index) {
//    motors_.MotorZRNRegInit(index);
//  }

  // 触发回零
  for (unsigned char index = 0x01; index <= 0x06; ++index) {
    motors_.MotorZRN(index, true);
  }

  // 每100ms 轮询回零是否完成
  while (!ZRNDone()) {
    UpdateMotorRealPos();
    std::this_thread::sleep_for(std::chrono::milliseconds(10));
  }

  // 每100ms 轮询回零后移动到-510000 的动作是否完成
  while (!MoveDone()) {
    UpdateMotorRealPos();
    std::this_thread::sleep_for(std::chrono::milliseconds(10));
  }

  set_zrn(false);
}

/**
 * @brief 更新每个电机的实际位置
 */
void Controller::UpdateMotorRealPos() {
  GetAllMotorStatus(kMotorPos); // 更新电机实际位置

  // 写入 motor_real_pos_ 寄存器，独占锁
  std::unique_lock<std::shared_mutex> lock(mutex_);
  // 写入电机的实际位置
  for (int index = 0; index < 6; ++index) {
    motor_real_pos_[index] = (motor_reg_status_[index][3] << 16) + (motor_reg_status_[index][4]);
  }
}

/**
 * @brief 调用 grpc client，发送动作完成的信息
 * @param _req MoveDoneRequestBody 包含执行完成的动作坐标
 * @return bool 是否发送成功
 */
bool Controller::Send(const MoveDoneRequestBody &_req) {
  auto server_address =
      (Config::Get()->BasicSetting()->server_rpc_address() + ":" + Config::Get()->BasicSetting()->server_rpc_port());
  MoveDoneClient client(server_address);

  auto status = client.MoveDone(_req);
#if DEBUG
  std::cout << "Send Status: " << status << std::endl;
#endif
  return status;
}

/*

void Controller::Camera() {
  motion_coord_.push(std::array<double, 6>{0, 0, 0, 30, 30, 0});
  motion_coord_.push(std::array<double, 6>{0, 0, 0, 30, -30, 0});
  motion_coord_.push(std::array<double, 6>{0, 0, 0, 0, -30, 0});
  motion_coord_.push(std::array<double, 6>{0, 0, 0, -30, -30, 0});
  motion_coord_.push(std::array<double, 6>{0, 0, 0, -30, 30, 0});
  motion_coord_.push(std::array<double, 6>{0, 0, 0, 0, 30, 0});
}
*/

/**
 * @brief 线程函数
 * @details 该函数实现了该线程的工作逻辑，先执行回零函数，在进入线程循环工作，检测网络通讯模块是否向队列中传入移动数据，若队列中存在运动数据，就依次执行
 */
void Controller::Main() {
  // 初始化系统
  InitialSystem();

  while (is_running()) {
    // 检测回零信号
    if (zrn()) {
      MotorZRN();
    } else if (!motion_coord_empty()) { // 动作队列
      auto motion = motion_coord_pop();

      // 执行动作
      Move(motion.begin(), motion.begin() + 3);

      // 等待运动
      while (!MoveDone()) {
        UpdateMotorRealPos();
        std::this_thread::sleep_for(std::chrono::milliseconds(10));
      }

      // 等待
      if (motion.at(5) > 0) {
        std::this_thread::sleep_for(std::chrono::milliseconds(static_cast<long>(motion.at(5))));
      }

      // 发送完成信息
      Send(MoveDoneRequestBody{
          motion.at(0),
          motion.at(1),
          motion.at(2),
          motion.at(3),
          motion.at(4),
          false,
          motion.at(5),
          Config::Get()->BasicSetting()->uuid()
      });
    } else if (!check_coord_empty()) {  // 检查队列
      auto check = check_coord_pop();
      auto motion = check.coord_;

      // 执行动作
      Move(motion.begin(), motion.begin() + 3);

      // 等待运动
      while (!MoveDone()) {
        UpdateMotorRealPos();
        std::this_thread::sleep_for(std::chrono::milliseconds(10));
      }

      // 等待
      if (motion.at(5) > 0) {
        std::this_thread::sleep_for(std::chrono::milliseconds(static_cast<long>(motion.at(5))));
      }

      // 发送完成信息
      Send(MoveDoneRequestBody{
          motion.at(0),
          motion.at(1),
          motion.at(2),
          motion.at(3),
          motion.at(4),
          true,
          motion.at(5),
          Config::Get()->BasicSetting()->uuid(),
          check.check_name_
      });
    } else {
      // 线程休眠10ms
      std::this_thread::sleep_for(std::chrono::milliseconds(10));
    }
  }

  ShutDownSystem(); // 关闭系统
}
