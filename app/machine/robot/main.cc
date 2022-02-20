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

#include "controller.h"
#include "internal/conf/config.h"
#include "internal/server/robot_server.h"

#include <iostream>
#include <csignal>

bool is_running = true;

void OnSignal(int) {
  is_running = false;
}

int main(int _argc, char *_argv[]) {
  if (_argc == 2 && (std::string(_argv[1]) == "-v" || std::string(_argv[1]) == "--version")) {
    printf("%s homepage url: %s\n", PROJECT_NAME, HOMEPAGE_URL);
    exit(EXIT_SUCCESS);
  }

  // 触发下面的信号就退出
  signal(SIGINT, OnSignal);
  signal(SIGQUIT, OnSignal);
  signal(SIGTERM, OnSignal);

  // 启动控制器
  Controller::Get()->Start();

  // 获取 rpc 地址
  auto local_address =
      Config::Get()->BasicSetting()->local_rpc_address() + ":" + Config::Get()->BasicSetting()->local_rpc_port();

  // 实例化 rpc 服务
  RobotServer server(local_address, Controller::Get());

  // 启动 rpc 服务
  server.Start();

  Controller::Get()->ThreadSleep(std::chrono::milliseconds(100));


  while (true) {
    // 退出
    if (!is_running) {
      break;
    }

    // 读取寄存器状态和电机实际位置信息，写入curl类中并发送
    auto status = Controller::Get()->motor_status();
    auto pos = Controller::Get()->motor_real_pos();
    // 显示寄存器状态以及电机实际位置
    std::cout << "========================" << std::endl;
    for (int i = 0; i < 6; ++i) {
      std::cout << "addr: " << i + 1 << ": ";
      std::cout << status[i] << ", " << pos[i] << ", " << std::endl;
    }
    std::cout << "========================" << std::endl << std::endl;

    std::this_thread::sleep_for(std::chrono::seconds (1));
  }

  server.Stop();
  Controller::Get()->Stop();

  return 0;
}
