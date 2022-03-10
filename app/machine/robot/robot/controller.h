//
// Created by HominSu on 2021/5/8.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_ROBOT_CONTROLLER_H_
#define HTP_PLATFORM_MACHINE_ROBOT_ROBOT_CONTROLLER_H_

#include "robot/cal_len.h"
#include "robot/internal/data/motor/motor.h"
#include "robot/internal/utils/x_thread.h"
#include "robot/internal/service/robot_impl.h"
#include "robot/internal/service/move_done_req.h"

#include <bitset>
#include <array>
#include <queue>
#include <shared_mutex>

const int kOffset = -510000;
const int kMotorStatus = 1;
const int kMotorPos = 2;
const int kMotorAll = 3;

/**
 * @brief Controller 类封装了简单的运动函数
 * @details 该类实现了对于执行运动以及回零动作的简单封装，继承XThread基类，简单的实现了多线程功能
 */
class Controller : public XThread {
 private:
  friend RobotImpl;  ///< 声明 CommandServiceImpl 为友元类，确保能向 Controller 中的队列中写入数据

  bool system_status_ = false; ///< 系统运行状态
  Motor motors_;    ///< 电机控制对象
  CalLen cal_len_;  ///< 坐标解算对象

  std::queue<std::array<double, 6>> motion_coord_;  ///< 运动的坐标以及姿态
  std::queue<Coord> check_coord_{};
  std::array<std::bitset<7>, 6> motor_status_{};
  std::array<int, 6> motor_real_pos_{};
  std::array<int, 6> motor_expected_pos_{};

  bool zrn_ = true;  ///< 回零标识
  std::array<double, 6> zero_{0}; ///< 零位参考
  std::array<std::array<unsigned short, kMotorStatusLength>, 6> motor_reg_status_{}; ///< 保存控制器寄存器信息

  mutable std::shared_mutex mutex_;

 public:
  ~Controller();

  /**
   * @brief 单件模式，只初始化一次
   * @return Controller *
   */
  static Controller *Get() {
    static Controller c;
    return &c;
  }

 private:
  Controller();
  inline static bool Send(const MoveDoneRequestBody &_req);

 public:
  void InitialSystem();

 private:
  void Main() override;

  void ShutDownSystem();
  void SetZero();
  void MotorZRN();

  inline static int CalMotorPulse(double _len, double _zero);
  void Move(double _m_xyz[], double _m_ges[]);
  void GetMotorAllStatus(unsigned char _addr);
  void GetMotorStatusReg(unsigned char _addr);
  void GetMotorPosReg(unsigned char _addr);
  void GetAllMotorStatus(const int &_option);
  void UpdateMotorRealPos();
  bool ZRNDone();
  bool MoveDone();

  template<typename _num>
  inline _num Abs(_num x) { return x < 0 ? -x : x; }

  template<typename _data>
  inline _data GetShiftData(_data _high, _data _low) {
    return (_high << 16) + _low;
  }

  // 写入、读取器
 public:
  std::array<std::bitset<7>, 6> motor_status();
  std::array<int, 6> motor_real_pos();
  std::array<int, 6> motor_expected_pos();

 private:
  bool zrn();
  void set_zrn(bool _zrn);

  bool motion_coord_empty();
  std::array<double, 6> motion_coord_pop();
  void motion_coord_push(const std::array<double, 6> &_coord);

  bool check_coord_empty();
  Coord check_coord_pop();
  void check_coord_push(const Coord &_coord);

};

#endif //HTP_PLATFORM_MACHINE_ROBOT_ROBOT_CONTROLLER_H_
