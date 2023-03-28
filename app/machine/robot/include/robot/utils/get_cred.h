//
// Created by Homin Su on 2021/9/10.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_UTILS_GET_CRED_H_
#define HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_UTILS_GET_CRED_H_

#include <memory>
#include <string>

#include <grpc++/grpc++.h>

class GetCred {
 public:
  static std::string GetFileContents(const std::string &_path);
  static std::shared_ptr<grpc::ServerCredentials> GetServerCred();
  static std::shared_ptr<grpc::ChannelCredentials> GetClientCred();
};

#endif //HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_UTILS_GET_CRED_H_
