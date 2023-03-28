//
// Created by Homin Su on 2022/2/19.
//

#include "robot_server.h"

#include "robot_impl.h"
#include "robot/utils/get_cred.h"
#include "robot/utils/logger.h"


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
  RobotImpl service(controller_ptr_, bootstrap_);
  grpc::ServerBuilder builder;

  builder.AddListeningPort(bootstrap_->local().grpc().addr(), grpc::InsecureServerCredentials());
  builder.RegisterService(&service);

  std::unique_ptr<grpc::Server> server(builder.BuildAndStart());
  server_ = std::move(server);

  DEBUG("Server listening on %s", bootstrap_->local().grpc().addr().c_str());

  server_->Wait();

  DEBUG("Rpc Service Shut Down");
}
