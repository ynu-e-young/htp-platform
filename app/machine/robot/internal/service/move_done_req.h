//
// Created by Homin Su on 2022/2/19.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_INTERNAL_SERVICE_MOVE_DONE_REQ_H_
#define HTP_PLATFORM_MACHINE_ROBOT_INTERNAL_SERVICE_MOVE_DONE_REQ_H_

struct MoveDoneRequestBody {
  double x;         // x
  double y;         // y
  double z;         // z
  double rx;        // 绕 x 角度
  double ry;        // 绕 y 角度
  bool check;       // 是否为检查动作
  double delay;     // 每个动作的延时
  std::string uuid; // 设备的唯一识别码
  std::string check_name; // 检查名称
};

struct Coord {
  std::array<double, 6> coord_;
  std::string check_name_;
};

#endif //HTP_PLATFORM_MACHINE_ROBOT_INTERNAL_SERVICE_MOVE_DONE_REQ_H_
