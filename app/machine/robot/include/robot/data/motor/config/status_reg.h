//
// Created by Homin Su on 2021/5/12.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_DATA_MOTOR_CONFIG_STATUS_REG_H_
#define HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_DATA_MOTOR_CONFIG_STATUS_REG_H_

///< 保存了需要读取的控制器的寄存器的地址
const unsigned short kMotorStatusReg[] = {
    0x1003, 0x602a, 0x602b, 0x602c, 0x602d, 0x6002
};

#endif //HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_DATA_MOTOR_CONFIG_STATUS_REG_H_