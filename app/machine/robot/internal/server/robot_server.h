//
// Created by Homin Su on 2022/2/19.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_INTERNAL_SERVER_ROBOT_SERVER_H_
#define HTP_PLATFORM_MACHINE_ROBOT_INTERNAL_SERVER_ROBOT_SERVER_H_

#include "../utils/x_thread.h"

#include <grpc++/grpc++.h>
#include <queue>


class Controller;

class RobotServer : public XThread {
 private:
  Controller *controller_ptr_{};
  std::string local_address_{};  ///< rpc 服务地址：ip+端口
  std::unique_ptr<grpc::Server> server_{};  ///< rpc 服务句柄，用智能指针管理

 public:
  explicit RobotServer(std::string _local_address, Controller *_controller_ptr)
      : local_address_(std::move(_local_address)), controller_ptr_(_controller_ptr) {};

 public:
  void Main() override;
  void Stop() override;
};

#endif //HTP_PLATFORM_MACHINE_ROBOT_INTERNAL_SERVER_ROBOT_SERVER_H_
