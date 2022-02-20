//
// Created by HominSu on 2021/5/6.
//

#include <cstring>

#include <iostream>
#include <thread>
#include <vector>
#include <array>

#include "motor_serial_comm.h"

MotorSerialComm::MotorSerialComm() = default;

MotorSerialComm::~MotorSerialComm() {
  if (this->serial_ != nullptr) {
    this->serial_->CloseSerialPort();
    delete this->serial_;
  }
  this->serial_ = nullptr;
}

/**
 * @brief 初始化一个传输串口
 * @details 通过实例化一个 GetSerialInfo 对象，从 JSON 文件中获取相关的串口通信预设，并初始化串口
 */
void MotorSerialComm::InitialCommander() {
  // 从Json文件获取串口设置
  GetSerialInfo get_serial_info;
  get_serial_info.LoadJsonInfo();

  // 通过读取到的信息来初始化串口
  this->serial_ = new Serial(get_serial_info.serial_port_,
                             get_serial_info.baud_rate_,
                             get_serial_info.data_bits_,
                             get_serial_info.stop_bits_,
                             get_serial_info.parity_,
                             get_serial_info.flow_control_,
                             get_serial_info.clocal_);

  // 打开并初始化串口
  try {
    this->serial_->SerialOpen();
  } catch (std::runtime_error &runtime_error) {
    throw runtime_error;
  } catch (std::exception &e) {
    throw e;
  }
}

/**
 * @brief 关闭串口
 */
void MotorSerialComm::CloseCommander() {
  if (this->serial_ != nullptr) {
    this->serial_->CloseSerialPort();
    delete this->serial_;
  }
  this->serial_ = nullptr;
}

/**
 * @brief 生成 CRC 校验码
 * @param _ptr 生成校验码的目标数组
 * @param _len 目标数组的长度
 * @return 返回一个 unsigned short 类型的 CRC 校验码
 */
unsigned short MotorSerialComm::Crc16Convert(unsigned char *_ptr, unsigned short _len) {
  unsigned short crc = 0xffff;
  unsigned short i;
  unsigned char ch;

  for (i = 0; i < _len; i++) {
    ch = *_ptr++;
    crc = KCrcTables[(ch ^ crc) & 15] ^ (crc >> 4);
    crc = KCrcTables[((ch >> 4) ^ crc) & 15] ^ (crc >> 4);
  }

  return crc;
}

/**
 * @brief 向控制器写入指令
 * @details 使用 Modbus 协议规定的指令格式，实现具体的写指令
 * @param _equ_addr 8 位下位机地址
 * @param _reg_addr 16 位寄存器地址
 * @param _reg_data 写入的数据
 */
void MotorSerialComm::WriteInstrFirst(unsigned char _equ_addr, unsigned short _reg_addr, unsigned short _reg_data) {
  std::vector<unsigned char> write_data;
  unsigned char temp_data[6];
  unsigned short Crc16;

  temp_data[0] = _equ_addr;                 // 下位机地址，8位
  temp_data[1] = 0x06;                    // 功能码，表明要向下位机寄存器写入数据
  temp_data[2] = (_reg_addr >> 8) & 0xFF;       // 寄存器地址高八位
  temp_data[3] = (_reg_addr & 0xFF);          // 寄存器地址低八位
  temp_data[4] = (_reg_data >> 8) & 0xFF;        // 写入数据的高八位
  temp_data[5] = (_reg_data & 0xFF);           // 写入数据的低八位

  // 生成这条指令的CRC码
  Crc16 = Crc16Convert(temp_data, sizeof(temp_data));
  for (unsigned char &i : temp_data) {
    write_data.push_back(i);
  }

  // 将指令与CRC码组合起来，成为一条完整的读取指令
  write_data.push_back(Crc16 & 0xFF);
  write_data.push_back((Crc16 >> 8) & 0xFF);

  this->serial_->Write(write_data);

  // 线程休眠8ms
  std::this_thread::sleep_for(std::chrono::milliseconds(10));
}

/**
 * @brief 向控制器的写入读取信号
 * @details 使用 Modbus 协议规定的指令格式，向控制器写入具体的读指令
 * @param _equ_addr 8 位下位机地址
 * @param _reg_addr 16 位寄存器地址
 * @param _reg_num 寄存器数量
 */
void MotorSerialComm::WriteReadSignal(unsigned char _equ_addr, unsigned short _reg_addr, unsigned short _reg_num) {
  std::vector<unsigned char> write_data;  // 读取寄存器的指令+CRC校验码
  unsigned char temp_data[6]; // 存放指令，指令由6个8位数据组成
  unsigned short Crc16;

  temp_data[0] = _equ_addr; // 下位机地址，8位
  temp_data[1] = 0x03;      // 0x03是功能码，表示上位机想要读取下位机的寄存器
  temp_data[2] = _reg_addr >> 8; // 寄存器起始地址高八位
  temp_data[3] = _reg_addr & 0xFF;      // 寄存器起始地址低八位
  temp_data[4] = _reg_num >> 8;  // 寄存器数量高八位
  temp_data[5] = (_reg_num & 0xFF);       // 寄存器数量低八位

  // 生成这条指令的CRC码
  Crc16 = Crc16Convert(temp_data, sizeof(temp_data));
  for (unsigned char &i : temp_data) {
    write_data.push_back(i);
  }

  // 将指令与CRC码组合起来，成为一条完整的读取指令
  write_data.push_back(Crc16 & 0xFF);
  write_data.push_back((Crc16 >> 8) & 0xFF);

  this->serial_->Write(write_data);

  // 线程休眠8ms
  std::this_thread::sleep_for(std::chrono::milliseconds(10));
}

/**
 * @brief 向下位机发出读取寄存器的指令，同时接收下位机发回的反馈信息
 * @param _equ_addr 8 位下位机地址
 * @param _reg_addr 16 位寄存器地址
 * @param _reg_num 读取的寄存器数量
 */
void MotorSerialComm::ReadInstr(unsigned char _equ_addr, unsigned short _reg_addr, unsigned short _reg_num) {
// 定义两个寄存器，分别存储返回数据的8位
  unsigned short temp1 = 0;
  unsigned short temp2 = 0;

  WriteReadSignal(_equ_addr, _reg_addr, _reg_num);  // 向下位机发出读取数据的指令

  std::array<unsigned char, 256> read_buf;
  this->serial_->Read(read_buf); // 将下位机反馈回的信息存入缓冲区

  unsigned char feedback_addr = read_buf[0];  // 获取反馈回的地址
  unsigned char feedback_com = read_buf[1]; // 获取反馈回的功能码

  this->read_flag_ = false;

  if (_equ_addr == feedback_addr) {  // 判断反馈地址是不是请求的下位机的地址
    if (0x03 == feedback_com) { // 判断反馈的功能码是不是0x03
      unsigned char temp = 3; // 指向_read_buf中存放返回数据的起始地址
      for (size_t i = 0; i < _reg_num; i++) { // 读出所有返回的数据，_reg_com是上位机指定的要读取的寄存器个数
        temp1 = (read_buf[temp] & 0xFF) << 8; // 读出数据的高8位，然后将数据左移8位，使低8位全为0，便于后续合并
        temp2 = read_buf[temp + 1] & 0xFF;  // 读出数据的低8位
        this->read_reg_[i] = temp1 + temp2; // 将两段数据直接相加，合并为完整数据，放入接收区
        temp = temp + 2;  // 指向下一个寄存器返回数据的起始地址
      }
      this->read_flag_ = true;  // 设置标志位为真，表明读取数据了
    }
  }
  std::this_thread::sleep_for(std::chrono::milliseconds(10));
}

/**
 * @brief 向下位机的一个寄存器写入数据
 * @param _equ_addr 8 位下位机地址
 * @param _reg_addr 16 位寄存器地址
 * @param _reg_data 写入的数据
 */
void MotorSerialComm::WriteInstr(unsigned char _equ_addr, unsigned short _reg_addr, unsigned short _reg_data) {
  WriteInstrFirst(_equ_addr, _reg_addr, _reg_data);
//  this->serial_->SerialFlush(); // 清空缓冲区
}

/**
 * 读取下位机的一个寄存器，封装成单次读取的信号
 * @param _addr 8 位下位机地址
 * @param _reg_addr 16 位寄存器地址
 * @param _read_reg 返回读取到的数据所在的起始位置
 * @param _read_flag 返回读取状态
 */
void MotorSerialComm::ReadSingleReg(unsigned char _addr,
                                    unsigned short _reg_addr,
                                    bool *_read_flag) {
  memset(this->read_reg_, 0, sizeof(this->read_reg_));  // 将接受区全部初始化为0
  ReadInstr(_addr, _reg_addr, 1);  // 发送读取指令，寄存器数量为1

  if (true == this->read_flag_) { // 标志位为真，即成功进行了读取
    *_read_flag = true; // 设置读取位为真
  } else {
    *_read_flag = false;
  }
}

/**
 * @brief 读取下位机的多个寄存器，封装成多次读取的信号
 * @param _addr 8 位下位机地址
 * @param _reg_addr 16 位寄存器地址
 * @param _reg_num 读取的寄存器个数
 * @param _read_flag 返回读取状态
 * @return 返回一个存放读取到的寄存器的信息的 unsigned short 数组
 */
unsigned short *MotorSerialComm::ReadMultipleReg(unsigned char _addr,
                                                 unsigned short _reg_addr,
                                                 unsigned short _reg_num,
                                                 bool *_read_flag) {
  memset(this->read_reg_, 0, sizeof(this->read_reg_));
  ReadInstr(_addr, _reg_addr, _reg_num); // 发送读取指令，可指定寄存器数量

  if (true == read_flag_) {
    *_read_flag = true;
    return this->read_reg_;  // 返回读取到的数据所在的位置
  } else {
    *_read_flag = false;
  }
  // 添加条件外返回
  return nullptr;
}

