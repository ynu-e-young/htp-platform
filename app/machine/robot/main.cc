//
// Created by HominSu on 2021/5/8.
//

/**
 * @mainpage 绳驱并联机器人控制系统
 * @authors 苏浩明 刘金卓 杨启元 范宸佑
 * @version v1.0
 * @date 2021-5-19
 */

/**
 * @file main.cc
 * @brief main.cc是主程序的入口
 * @details main函数完成的工作有：初始化 Controller、Curl 两个线程，将 Curl中获取的坐标传到 Controller 中
 */

#include "conf/config.h"
#include "conf/conf.pb.h"
#include "robot/controller.h"
#include "robot/internal/server/robot_server.h"
#include "robot/internal/utils/logger.h"

#include "unistd.h"

#include <csignal>
#include <sstream>

bool is_running = true;

void OnSignal(int) {
  is_running = false;
}

int main(int _argc, char *_argv[]) {
#ifndef NDEBUG
  htp_platform::logger::set_log_level(htp_platform::logger::Level::DEBUG);
#else
  htp_platform::logger::set_log_level(htp_platform::logger::Level::INFO);
#endif
  htp_platform::logger::set_log_fd(STDOUT_FILENO);

  if (_argc == 2 && (std::string(_argv[1]) == "-v" || std::string(_argv[1]) == "--version")) {
    printf("%s homepage url: %s\n", PROJECT_NAME, HOMEPAGE_URL);
    exit(EXIT_SUCCESS);
  }

  // 触发下面的信号就退出
  signal(SIGINT, OnSignal);
  signal(SIGQUIT, OnSignal);
  signal(SIGTERM, OnSignal);

  auto bootstrap = ::std::make_shared<config::Bootstrap>();
  auto plat_info = ::std::make_shared<config::PlatInfo>();

  Config conf;

  conf.Load("/data/conf/config.json");
  conf.Scan(bootstrap.get());
  conf.Load("/data/conf/robot.json");
  conf.Scan(plat_info.get());

  Controller::Get()->SetConfig(bootstrap, plat_info);
  // 启动控制器
  Controller::Get()->Start();

  // 实例化 rpc 服务
  RobotServer server(Controller::Get(), bootstrap);

  // 启动 rpc 服务
  server.Start();

  Controller::Get()->ThreadSleep(std::chrono::milliseconds(100));

  std::stringstream ss;
  while (true) {
    // 退出
    if (!is_running) {
      break;
    }

    // 读取寄存器状态和电机实际位置信息，写入curl类中并发送
    auto status = Controller::Get()->motor_status();
    auto pos = Controller::Get()->motor_real_pos();
    // 显示寄存器状态以及电机实际位置
    ss << std::endl << "========================" << std::endl;
    for (int i = 0; i < 6; ++i) {
      ss << "addr: " << i + 1 << ": ";
      ss << "" << status[i] << ", " << pos[i] << ", " << std::endl;
    }
    ss << "========================";
    INFO("%s", ss.str().c_str());
    ss.clear();
    ss.str("");

    std::this_thread::sleep_for(std::chrono::seconds(1));
  }

  server.Stop();
  Controller::Get()->Stop();

  return 0;
}
