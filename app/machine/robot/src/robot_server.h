//
// Created by Homin Su on 2022/2/19.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_SRC_ROBOT_SERVER_H_
#define HTP_PLATFORM_MACHINE_ROBOT_SRC_ROBOT_SERVER_H_

#include <queue>

#include <grpc++/grpc++.h>

#include "conf/conf.pb.h"
#include "robot/utils/x_thread.h"

class Controller;

class RobotServer : public XThread {
 private:
  Controller *controller_ptr_{};
  ::std::shared_ptr<config::Bootstrap> bootstrap_;
  std::unique_ptr<grpc::Server> server_{};  ///< rpc 服务句柄，用智能指针管理

 public:
  explicit RobotServer(Controller *_controller_ptr, ::std::shared_ptr<config::Bootstrap> _bootstrap)
      : controller_ptr_(_controller_ptr), bootstrap_(std::move(_bootstrap)) {};

 public:
  void Main() override;
  void Stop() override;
};

#endif //HTP_PLATFORM_MACHINE_ROBOT_SRC_ROBOT_SERVER_H_
