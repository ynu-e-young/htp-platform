//
// Created by Homin Su on 2022/2/19.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_APP_MACHINE_ROBOT_INTERNAL_SERVICE_ROBOT_IMPL_H_
#define HTP_PLATFORM_MACHINE_ROBOT_APP_MACHINE_ROBOT_INTERNAL_SERVICE_ROBOT_IMPL_H_

#include "../../../../../api/machine/robot/v1/cpp/robot.grpc.pb.h"

#include <grpc++/grpc++.h>
#include <queue>

class Controller;

class RobotImpl final : public machine::robot::v1::Robot::Service {
 private:
  Controller *controller_ptr_{};  ///< controller 对象指针，用于写入坐标数据

 public:
  explicit RobotImpl(Controller *_controller_ptr) :
      controller_ptr_(_controller_ptr) {};

  ::grpc::Status AppendCoordinate(::grpc::ServerContext *_context,
                                  const ::machine::robot::v1::CoordinateRequest *_req,
                                  ::machine::robot::v1::CoordinateReply *_reply) override;

  ::grpc::Status Zero(::grpc::ServerContext *_context,
                      const ::machine::robot::v1::ZeroRequest *_req,
                      ::machine::robot::v1::ZeroReply *_reply) override;

  ::grpc::Status GetMotorStatus(::grpc::ServerContext *_context,
                                const ::machine::robot::v1::MotorInfoRequest *_req,
                                ::machine::robot::v1::MotorInfoReply *_reply) override;
};

#endif //HTP_PLATFORM_MACHINE_ROBOT_APP_MACHINE_ROBOT_INTERNAL_SERVICE_ROBOT_IMPL_H_
