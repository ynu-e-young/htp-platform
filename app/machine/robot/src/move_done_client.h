//
// Created by Homin Su on 2022/2/19.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_SRC_MOVE_DONE_CLIENT_H_
#define HTP_PLATFORM_MACHINE_ROBOT_SRC_MOVE_DONE_CLIENT_H_

#include "apis/htpp/machine/service/v1/machine.grpc.pb.h"
#include <grpc++/grpc++.h>

#include "move_done_req.h"
#include "robot/utils/get_cred.h"

class MoveDoneClient {
 public:
  explicit MoveDoneClient(const std::string &_server_address)
      : stub_(::htpp::machine::service::v1::Machine::NewStub(
      grpc::CreateChannel(_server_address, grpc::InsecureChannelCredentials()))) {}

 public:
  bool MoveDone(const MoveDoneRequestBody &_body);

 private:
  std::unique_ptr<::htpp::machine::service::v1::Machine::Stub> stub_{};
};

#endif //HTP_PLATFORM_MACHINE_ROBOT_SRC_MOVE_DONE_CLIENT_H_
