//
// Created by HominSu on 2021/5/2.
//

#include "serial.h"
#include "robot/internal/utils/logger.h"

#include <fcntl.h>    // 包含文件控件，例如 O_RDWR
#include <termios.h>  // 包含 POSIX 终端控件定义
#include <unistd.h>   // write(), read(), close()
#include <sys/file.h> // 排他锁

#include <cassert>
#include <cerrno>
#include <cstring>

/**
 * @brief 显示打开的串口信息
 */
void Serial::ShowTermios() const {
  INFO(R"(
Termios info:
	input mode flags: %ld
	output mode flags: %ld
	control mode flags: %ld
	local mode flags: %ld
	input speed: %s
	output speed: %s
)",
       this->Opt.c_iflag,
       this->Opt.c_oflag,
       this->Opt.c_cflag,
       this->Opt.c_lflag,
       this->Opt.c_cc,
       this->Opt.c_cc);
}

/**
 * @brief 设置校验位
 */
void Serial::SetParity() {
  switch (this->parity_) {
    case 0: // 无校验
      this->Opt.c_cflag &= ~PARENB;
      this->Opt.c_iflag &= ~INPCK;
      break;
    case 2: // 偶校验
      this->Opt.c_cflag |= PARENB;
      this->Opt.c_cflag &= ~PARODD;
      this->Opt.c_iflag |= INPCK;
      break;
    case 3: // 奇校验
      this->Opt.c_cflag |= (PARODD | PARENB);
      this->Opt.c_iflag |= INPCK;
      break;
    default:  // 默认不设校验位
      this->Opt.c_cflag &= ~PARENB;
      this->Opt.c_iflag &= ~INPCK;
      break;
  }
}

/**
 * @brief 设置停止巍峨
 */
void Serial::SetStopBits() {
  // 设置停止位
  if (this->stop_bits_ == 2) {
    this->Opt.c_cflag |= CSTOPB;  // 设置两位停止位
  } else {
    this->Opt.c_cflag &= ~CSTOPB; // 设置一位停止位
  }
}

/**
 * @brief 设置数据位
 */
void Serial::SetDataBits() {
  // 设置数据位
  this->Opt.c_cflag &= ~CSIZE; // 清除所有大小位
  switch (this->data_bits_) {
    case 8: this->Opt.c_cflag |= CS8;
      break;
    case 5: this->Opt.c_cflag |= CS5;
      break;
    case 6: this->Opt.c_cflag |= CS6;
      break;
    case 7: this->Opt.c_cflag |= CS7;
      break;
    default: this->Opt.c_cflag |= CS8;
      break;  // 默认8位数据
  }
}

/**
 * @brief 设置流控制
 */
void Serial::SetFlowControl() {
  if (this->flow_control_) {
    this->Opt.c_cflag |= CRTSCTS;  // 启用RTS/CTS硬件流控制
  } else {
    this->Opt.c_cflag &= ~CRTSCTS; // 禁用RTS/CTS硬件流控制（最常见）
  }
}

/**
 * @brief 设置 CLOCAL 标志位, 本地连接
 */
void Serial::SetClocal() {
  if (this->clocal_) {
    this->Opt.c_cflag |= CREAD | CLOCAL;
  }
}

/**
 * @brief 设置 RAW 模式
 */
void Serial::SetLocalModes() {
  // 禁用规范模式，不禁用可能会导致字节丢失
  this->Opt.c_lflag &= ~ICANON;

  // Echo，如果设1，发送的字符将会被回显，以防万一便禁用
  this->Opt.c_lflag &= ~ECHO; // 禁用 echo
  this->Opt.c_lflag &= ~ECHOE; // 禁用 erasure
  this->Opt.c_lflag &= ~ECHONL; // 禁用换行 echo

  // 禁用信号字符，不使用串行端口，清除以下位
  this->Opt.c_lflag &= ~ISIG; // 禁止解释INTR，QUIT和SUSP
}

/**
 * @brief 设置输入模式
 */
void Serial::SetInputModes() {
  // 禁用软件流控制
  this->Opt.c_iflag &= ~(IXON | IXOFF | IXANY); // Turn off s/w flow ctrl

  // 禁用接收时字节的特殊处理
  this->Opt.c_iflag &=
      ~(IGNBRK | BRKINT | PARMRK | ISTRIP | INLCR | IGNCR | ICRNL); // Disable any special handling of received bytes
}

/**
 * @brief 设置输出模式
 */
void Serial::SetOutputModes() {
  // 在配置串行端口时，我们要禁用对输出字符/字节的任何特殊处理
  this->Opt.c_oflag &= ~OPOST; // Prevent special interpretation of output bytes (e.g. newline chars)
  this->Opt.c_oflag &= ~ONLCR; // Prevent conversion of newline to carriage return/line feed
}

/**
 * @brief 设置串口接收的最小时间和最小字节数
 */
void Serial::SetVminAndVTime() {
  // Wait for up to 0.1*MAX_ARDUINO_WAIT_TIME s (MAX_ARDUINO_WAIT_TIME deciseconds), returning as soon as any data is received.
  this->Opt.c_cc[VTIME] = MAX_ARDUINO_WAIT_TIME;
  this->Opt.c_cc[VMIN] = 0;
}

/**
 * @brief 设置传输波特率
 */
void Serial::SetBaudRate() {
  // 设置读写速率
  switch (this->baud_rate_) {
    case 9600: {
      cfsetspeed(&this->Opt, B9600);
      break;
    }
    case 19200: {
      cfsetspeed(&this->Opt, B19200);
      break;
    }
    case 38400: {
      cfsetspeed(&this->Opt, B38400);
      break;
    }
    case 57600: {
      cfsetspeed(&this->Opt, B57600);
      break;
    }
    case 115200: {
      cfsetspeed(&this->Opt, B115200);
      break;
    }
    default: {
      cfsetspeed(&this->Opt, B9600);
      break;
    }
  }
}

/**
 * @brief 初始化串口
 */
void Serial::SerialInitial() {
  fcntl(this->serial_port_, F_SETFL, 0);
  INFO("Test Port %s  has been successfully opened and %d is the file description",
       this->port_number_.c_str(),
       this->serial_port_);

  // 获取当前的串口设置
  if (tcgetattr(this->serial_port_, &this->Opt)) {
    SYSERR("Error  from tcgetattr");
  }

  SetParity();
  SetStopBits();
  SetDataBits();
  SetDataBits();
  SetFlowControl();
  SetClocal();
  SetLocalModes();
  SetInputModes();
  SetOutputModes();
  SetBaudRate();

  this->ShowTermios();

  // Save tty settings, also checking for error
  if (tcsetattr(this->serial_port_, TCSANOW, &this->Opt)) {
    SYSERR("Error  from tcgetattr");
  }

  this->connected_ = true;

  // 获取非阻塞排他锁
  // EACCES：访问出错
  // EAGAIN：文件已被锁定，或者太多的内存已被锁定
  // EBADF：fd不是有效的文件描述词
  // EINVAL：一个或者多个参数无效
  // ENFILE：已达到系统对打开文件的限制
  // ENODEV：指定文件所在的文件系统不支持内存映射
  // ENOMEM：内存不足，或者进程已超出最大内存映射数量
  // EPERM：权能不足，操作不允许
  // ETXTBSY：已写的方式打开文件，同时指定MAP_DENYWRITE标志
  // SIGSEGV：试着向只读区写入
  // SIGBUS：试着访问不属于进程的内存区
  if (flock(this->serial_port_, LOCK_EX | LOCK_NB)) {
    throw std::runtime_error("Serial port with file descriptor " +
        std::to_string(this->serial_port_) + " is already locked by another process.");
  }
}

/**
 * @brief 打开串口
 * @details 如果无法打开，抛出异常
 */
void Serial::SerialOpen() {
  this->connected_ = false;

  // 尝试打开串口
  this->serial_port_ = open(this->port_number_.c_str(), O_RDWR | O_NOCTTY | O_NDELAY);

  try {
    if (-1 == this->serial_port_) {
      throw std::runtime_error("Open_port: Unable to open " + std::string(this->port_number_));
    } else {
      SerialInitial();
    }
  } catch (std::runtime_error &runtime_error) {
    throw runtime_error;
  } catch (std::exception &e) {
    throw e;
  } catch (...) {
    FATAL("catch unknown wrong!");
  }
}

/**
 * @brief 串口类构造函数
 * @details 该函数用于初始化后打开的串口
 * @param _port_number 串口号
 * @param _baud_rate 传输波特率
 * @param _data_bits 数据位
 * @param _stop_bits 停止位
 * @param _parity 校验位
 * @param _flow_control 流控制
 * @param _clocal 本机连接
 */
Serial::Serial(std::string _port_number,
               int32_t _baud_rate,
               int32_t _data_bits,
               int32_t _stop_bits,
               int32_t _parity,
               bool _flow_control,
               bool _clocal) {
  // 初始化串口参数
  this->port_number_ = std::move(_port_number);
  this->baud_rate_ = _baud_rate;
  this->data_bits_ = _data_bits;
  this->stop_bits_ = _stop_bits;
  this->parity_ = _parity;
  this->flow_control_ = _flow_control;
  this->clocal_ = _clocal;
}

/**
 * @brief 析构。关闭串口
 */
Serial::~Serial() {
  if (this->connected_) {
    this->connected_ = false;
    close(this->serial_port_);
  }
}

/**
 * @brief 从串口读取数据
 * @param recv_array 接收的数组的引用 std::array<unsigned char, 256> 类型
 * @return 返回接收的实际字节数 size_t 类型
 */
size_t Serial::Read(std::array<unsigned char, 256> &recv_array) {
  // todo: 更改__buffer值
  // Allocate memory for read buffer, set size according to your needs
  unsigned char buffer[256]{0};
  unsigned char *index = buffer;

  // 通常不会执行这个memset()调用，但不加会有乱码
  memset(&buffer, 0, sizeof(buffer));

  // 读取字节。read()的行为取决于上面的配置设置，特别是VMIN和VTIME
  size_t num_bytes = read(this->serial_port_, &buffer, sizeof(buffer));

  if (num_bytes < 0) {
    SYSERR("Error reading");
  }

  for (auto &i: recv_array) {
    i = *index;
    ++index;
  }

  return num_bytes;
}

/**
 * @brief 向串口发送数据
 * @param _send_array 发送的 vector 的引用 std::vector<unsigned char>
 * @return 实际发送字节数 size_t 类型
 */
size_t Serial::Write(const std::vector<unsigned char> &_send_array) {
  // 将vector转换成对应大小的数组
  unsigned char buffer[_send_array.size()];
  unsigned char *index = buffer;
  for (auto i: _send_array) {
    *index++ = i;
  }

  // 发送数据
  size_t len = write(this->serial_port_, buffer, _send_array.size());

  return len;
}

/**
 * @brief 是否已经建立连接
 * @return bool 类型，返回值为 true，说明已经建立连接，反之亦然
 */
bool Serial::IsConnected() {
  return this->connected_;
}

/**
 * @brief 关闭串口
 */
void Serial::CloseSerialPort() {
  if (this->connected_) {
    this->connected_ = false;
    close(this->serial_port_);
  }
}

void Serial::SerialFlush() {
  tcflush(this->serial_port_, TCIFLUSH);
}