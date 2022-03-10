//
// Created by Homin Su on 2022/2/19.
//

#include "robot_server.h"

#include "robot/internal/service/robot_impl.h"
#include "robot/internal/utils/get_cred.h"

#include <fstream>

/**
 * @brief 停止线程任务
 * @details 重载 XThread 中的 Stop()，调用 XThread 中的 StopWith，传入关闭 grpc server 的 lambda
 */
void RobotServer::Stop() {
  std::function<void()> func = [this]() {
    this->server_->Shutdown();
  };
  XThread::StopWith(func);
}

/**
 * @brief 线程主函数，启动 rpc 服务，然后阻塞等待退出
 */
void RobotServer::Main() {
  RobotImpl service(controller_ptr_);
  grpc::ServerBuilder builder;

  builder.AddListeningPort(local_address_, grpc::InsecureServerCredentials());
  builder.RegisterService(&service);

  std::unique_ptr<grpc::Server> server(builder.BuildAndStart());
  server_ = std::move(server);

#if DEBUG
  std::cout << "Server listening on " << local_address_ << std::endl;
#endif

  server_->Wait();

#if DEBUG
  std::cout << "Rpc Service Shut Down" << std::endl;
#endif
}
