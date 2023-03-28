//
// Created by Homin Su on 2022/2/19.
//

#ifndef HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_CAL_LEN_H_
#define HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_CAL_LEN_H_

#define ARMD2R 0.01745329251994329576923690768489
#define MAX_DEM_SIZE 8

#include "conf/conf.pb.h"

/**
 * @brief 坐标解算
 * @details 通过坐标解算，将输入的坐标转换成对应没跟绳子的长度
 */
class CalLen {
 public:
  explicit CalLen();
  // path指定文件路径，读取文件中包含的绳索和末端相关信息
  bool OnInit(const ::std::shared_ptr<config::PlatInfo> &_plat_info);
  long OnPos(const double xyz[3], const double ges[3], double *retv, long &retnum);

 public:
  double m_max_pull_ang_; ///< 最大拉动角度
  double m_ancher_[MAX_DEM_SIZE][3]{};  ///< 长度相关信息
  double m_plate_[MAX_DEM_SIZE][3]{}; ///< 末端执行器相关信息
  int32_t m_dem_size_; ///< 最大轴数

 protected:
  // 记录与x、y、z轴的夹角，可用于描述转动自由度
  void RotX(double ang, double m[4][4]);
  void RotY(double ang, double m[4][4]);
  void RotZ(double ang, double m[4][4]);
  // 矩阵与矩阵相乘
  bool MatrixMult(double l[4][4], double r[4][4], double ret[4][4]);
  // 矩阵与向量相乘
  void MatrixVectMult(double l[4][4], double r[3], double result[3]);

};

#endif //HTP_PLATFORM_MACHINE_ROBOT_INCLUDE_ROBOT_CAL_LEN_H_
