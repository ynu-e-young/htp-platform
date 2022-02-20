//
// Created by Homin Su on 2022/2/19.
//

#include "move_done_client.h"

bool MoveDoneClient::MoveDone(const MoveDoneRequestBody &_body) {
  machine::service::v1::MoveDoneRequest req;  // 请求
  machine::service::v1::MoveDoneReply reply;  // 响应

  // 设置请求体
  req.set_x(_body.x);
  req.set_y(_body.y);
  req.set_z(_body.z);
  req.set_rx(_body.rx);
  req.set_ry(_body.ry);
  req.set_check(_body.check);
  req.set_delay(_body.delay);
  req.set_uuid(_body.uuid);
  if(_body.check){
    req.set_check_name(_body.check_name);
  }

  // context
  grpc::ClientContext context;

  // 调用
  grpc::Status status = stub_->MoveDone(&context, req, &reply);

  // 返回响应或者错误
  if (status.ok()) {
    return reply.status();
  } else {
    std::cerr << "RPC failed, error_code: [" << status.error_code() << "], error_message: ["
              << status.error_message() << "]" << std::endl;
    return false;
  }
}
