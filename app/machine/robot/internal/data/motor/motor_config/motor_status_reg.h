//
// Created by HominSu on 2021/5/2.
//

#ifndef SYSTEM_MOTOR_MOTOR_CONFIG_MOTOR_STATUS_REG_H_
#define SYSTEM_MOTOR_MOTOR_CONFIG_MOTOR_STATUS_REG_H_

///< 保存了需要读取的控制器的寄存器的地址
const unsigned short kMotorStatusReg[] = {
    0x1003, 0x602a, 0x602b, 0x602c, 0x602d, 0x6002
};

#endif //SYSTEM_MOTOR_MOTOR_CONFIG_MOTOR_STATUS_REG_H_
