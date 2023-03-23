//
// Created by Homin Su on 2022/2/19.
//

#include "robot_impl.h"

#include "conf/conf.pb.h"
#include "robot/controller.h"

::grpc::Status RobotImpl::AppendCoordinate(::grpc::ServerContext *_context,
                                           const ::htpp::machine::robot::v1::CoordinateRequest *_req,
                                           ::htpp::machine::robot::v1::CoordinateReply *_reply) {
  (void) _context;
  // 确认 uuid 相同
  if (bootstrap_->uuid() == _req->uuid()) {
    // 将坐标转存
    auto coord = std::array<double, 6>{
        _req->x(),
        _req->y(),
        _req->z(),
        _req->rx(),
        _req->ry(),
        _req->delay()
    };

    // 如果是检查，将坐标和检查名字写入到 check_coord_ 中，否则将坐标写入 motion_coord_ 中
    if (_req->check()) {
      controller_ptr_->check_coord_push(Coord{coord, _req->check_name()});
    } else {
      controller_ptr_->motion_coord_push(coord);
    }

    _reply->set_status(true);
  } else {
    _reply->set_status(false);
  }

  return ::grpc::Status::OK;
}

::grpc::Status RobotImpl::Zero(::grpc::ServerContext *_context,
                               const ::htpp::machine::robot::v1::ZeroRequest *_req,
                               ::htpp::machine::robot::v1::ZeroReply *_reply) {
  (void) _context;
  // 设置回零
  if (bootstrap_->uuid() == _req->uuid()) {
    controller_ptr_->set_zrn(true);
    _reply->set_status(true);
  } else {
    _reply->set_status(false);
  }

  return ::grpc::Status::OK;
}

::grpc::Status RobotImpl::GetMotorStatus(::grpc::ServerContext *_context,
                                         const ::htpp::machine::robot::v1::MotorInfoRequest *_req,
                                         ::htpp::machine::robot::v1::MotorInfoReply *_reply) {
  (void) _context;
  if (bootstrap_->uuid() == _req->uuid()) {
    // 获取控制器的一系列状态信息
    auto status = controller_ptr_->motor_status();
    auto real_pos = controller_ptr_->motor_real_pos();
    auto expect_pos = controller_ptr_->motor_expected_pos();

    // 写入到响应中
    for (int index = 0; index < kMaxMotor; index++) {
      auto motor_info = _reply->add_motor_info();

      motor_info->set_instr_pos(expect_pos[index]);
      motor_info->set_current_pos(real_pos[index]);

      auto motor_status = motor_info->mutable_motor_status();
      motor_status->set_fault(status[index][0]);
      motor_status->set_enabling(status[index][1]);
      motor_status->set_running(status[index][2]);
      motor_status->set_instruction_completion(status[index][4]);
      motor_status->set_path_completion(status[index][5]);
      motor_status->set_zero_completion(status[index][6]);
    }
  }

  return ::grpc::Status::OK;
}
