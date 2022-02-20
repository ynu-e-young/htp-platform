//
// Created by HominSu on 2021/4/29.
//

#include <cmath>

#include "cal_len.h"

CalLen::CalLen() {
  m_dem_size_ = 0;
  m_max_pull_ang_ = 60 * ARMD2R;
}

/**
 * @brief 将齐次变换矩阵用作算子，实现绕 x 轴的转动
 * @details 齐次变换矩阵中左上角的 3*3 矩阵为旋转矩阵，右上角 3*1 列向量为位置向量，由于只涉及转动，故位置向量设为 0，最后一行的行向量 (0,0,0,1) 是为了便于计算。初始化为单位矩阵，然后向旋转矩阵写入绕 x 轴旋转后的姿态。旋转矩阵中的每一列代表旋转后的坐标系的一轴在初始坐标系上的投影，由于只考虑方向，故都使用单位向量。绕 x 轴转动ang，x 轴保持不变，故旋转后的 x 轴相对于初始坐标系的投影为 (1,0,0), 类似可算出 y 的投影为(0, cos(ang), sin(ang)), z 的投影为(0, -sin(ang), cos(ang))，据此写入旋转矩阵，便可得到绕 x 轴转动的变换矩阵算子
 * @param ang 绕 x 轴转动的角度
 * @param m 变换算子，实现绕  x轴转动角度 ang
 */
// 相对x轴转动的变换矩阵，ang指定绕x轴转了多少,m为转动后的姿态
void CalLen::RotX(double ang, double m[4][4]) {
  double cosv = cos(ang);
  double sinv = sin(ang);
  // 生成变换矩阵
  for (long i = 0; i < 4; i++) {
    for (long j = 0; j < 4; j++) {
      if (i == j) {
        m[i][j] = 1;
      } else {
        m[i][j] = 0;
      }
    }
  }
  // 绕x轴转动ang角度后得到的姿态
  m[1][1] = cosv;
  m[1][2] = -sinv;
  m[2][1] = sinv;
  m[2][2] = cosv;
}

/**
 * @brief 将齐次变换矩阵用作算子，实现绕 y 轴的转动
 * @details 齐次变换矩阵中左上角的 3*3 矩阵为旋转矩阵，右上角 3*1 列向量为位置向量，由于只涉及转动，故位置向量设为 0，最后一行的行向量 (0,0,0,1) 是为了便于计算。初始化为单位矩阵，然后向旋转矩阵写入绕 y 轴旋转后的姿态。旋转矩阵中的每一列代表旋转后的坐标系的一轴在初始坐标系上的投影，由于只考虑方向，故都使用单位向量。绕 y 轴转动 ang，y 轴保持不变，故旋转后的 y 轴相对于初始坐标系的投影为 (0,1,0)，类似可算出 x 的投影为 (cos(ang),0,-sin(ang))，z 的投影为 (sin(ang),0,cos(ang))，据此写入旋转矩阵，便可得到绕 x 轴转动的变换矩阵算子
 * @param ang 绕 y 轴转动的角度
 * @param m 变换算子，实现绕 y 轴转动角度 ang
 */
// 相对y轴转动的变换矩阵，ang指定绕y轴转了多少,m为转动后的姿态
void CalLen::RotY(double ang, double m[4][4]) {
  double cosv = cos(ang);
  double sinv = sin(ang);
  for (long i = 0; i < 4; i++) {
    for (long j = 0; j < 4; j++) {
      if (i == j) {
        m[i][j] = 1;
      } else {
        m[i][j] = 0;
      }
    }
  }
  // 绕y轴转动ang角度后得到的姿态
  m[0][0] = cosv;
  m[0][2] = sinv;
  m[2][0] = -sinv;
  m[2][2] = cosv;
}

/**
 * @brief 将齐次变换矩阵用作算子，实现绕 z 轴的转动
 * @details 齐次变换矩阵中左上角的 3*3 矩阵为旋转矩阵，右上角 3*1 列向量为位置向量，由于只涉及转动，故位置向量设为 0, 最后一行的行向量 (0,0,0,1) 是为了便于计算。初始化为单位矩阵，然后向旋转矩阵写入绕 z 轴旋转后的姿态。旋转矩阵中的每一列代表旋转后的坐标系的一轴在初始坐标系上的投影，由于只考虑方向，故都使用单位向量。绕 z 轴转动ang, z 轴保持不变，故旋转后的 z 轴相对于初始坐标系的投影为 (0,0,1), 类似可算出 x 的投影为 (cos(ang), sin(ang), 0), y 的投影为(-sin(ang), cos(ang), 0),据此写入旋转矩阵，便可得到绕 x 轴转动的变换矩阵算子
 * @param ang 绕 z 轴转动的角度
 * @param m 变换算子，实现绕 z 轴转动角度 ang
 */
// 相对z轴转动的变换矩阵，ang指定绕z轴转了多少,m为转动后的姿态
void CalLen::RotZ(double ang, double m[4][4]) {
  double cosv = cos(ang);
  double sinv = sin(ang);
  for (long i = 0; i < 4; i++) {
    for (long j = 0; j < 4; j++) {
      if (i == j) {
        m[i][j] = 1;
      } else {
        m[i][j] = 0;
      }
    }
  }
  // 绕z轴转动ang角度后得到的姿态
  m[0][0] = cosv;
  m[0][1] = -sinv;
  m[1][0] = sinv;
  m[1][1] = cosv;
}

/**
 * @brief 两个 4 阶矩阵相乘，实现连续转动
 * @details l 和 r 分别为两个绕轴转动的变换矩阵算子, 将其相乘，可得连续绕轴转动两次的变换矩阵算子
 * @param l 绕某轴旋转的变换矩阵算子
 * @param r 绕某轴旋转的变换矩阵算子
 * @param ret 根据 l 和 r 指定的轴和角度连续旋转两次的变换矩阵算子
 * @return 返回值为 true
 */

bool CalLen::MatrixMult(double l[4][4], double r[4][4], double ret[4][4]) {
  int i, j, k, order;
  order = 4;
  for (i = 0; i < order; i++) {
    for (j = 0; j < order; j++) {
      ret[i][j] = l[i][0] * r[0][j];
      for (k = 1; k < order; k++) {
        ret[i][j] += l[i][k] * r[k][j];
      }
    }
  }
  return true;
}

/**
 * @brief 实现坐标系变换，描述末端一点的运动情况
 * @details 使用齐次变换矩阵，实现将物体坐标系上的点变换到空间坐标系下进行描述
 * @param l 变换矩阵，表示当前物体坐标系相对于空间坐标系的姿态
 * @param r 物体坐标系上一点
 * @param result 物体坐标系下 r 对应的点变换到空间坐标系下的表达法
 */
void CalLen::MatrixVectMult(double l[4][4], double r[3], double result[3]) {
  int i, j, order;
  order = 4;
  for (i = 0; i < 3; i++) {
    result[i] = l[i][3];    // 先将l中最后一列前三行写入，剩下的3*3就与r[3]维度对应，才能进行相乘
    for (j = 0; j < 3; j++) {
      result[i] += l[i][j] * r[j];
    }
  }
}

// 该函数用于获取绳索及平台在各个自由度上的初始信息
// path指定一个文件路径
/**
 * @brief 末端初始信息
 * @details 获取末端与绳的各个连接点在物体坐标系上的坐标
 * @return 返回值为 true
 */
bool CalLen::OnInit() {
  GetPlatInfo plat_info;
  plat_info.LoadJsonInfo();

  m_dem_size_ = 0;

  long ltmp = plat_info.ltmp_;

  for (long i = 0; i < ltmp; i++) {

    m_ancher_[i][0] = plat_info.m_ancher_[i][0];
    m_ancher_[i][1] = plat_info.m_ancher_[i][1];
    m_ancher_[i][2] = plat_info.m_ancher_[i][2];

    m_plate_[i][0] = plat_info.m_plate_[i][0];
    m_plate_[i][1] = plat_info.m_plate_[i][1];
    m_plate_[i][2] = plat_info.m_plate_[i][2];

    m_dem_size_ = ltmp;  //
  }
  return true;
}

/**
 * @brief 实现末端的移动与转动
 * @details 给定朝各轴移动的长度与方向和绕各轴旋转的角度，计算所需的绳索改变长度
 * @param xyz 位置向量，表示移动方式
 * @param ges 绕x、y、z轴转动的角度
 * @param retv 实现运动所需的绳索变化长度
 * @param retnum 绳索个数
 * @return
 */
long CalLen::OnPos(const double xyz[3], const double ges[3], double *retv, long &retnum) {

  double rx[4][4]; // 绕x轴旋转的变换矩阵
  double ry[4][4]; // 绕y轴旋转的变换矩阵
  double gps[4][4]; // 绕x轴、y轴连续旋转的变换矩阵
  // 初始化两个旋转算子，rx表示绕x轴转ges[0] * ARMD2R角度
  RotX(ges[0] * ARMD2R, rx);
  RotY(ges[1] * ARMD2R, ry);
  MatrixMult(rx, ry, gps); // 绕x轴、y轴转动后的姿态
  // xyz代表移动方向，写入gps变换矩阵的对应位置
  gps[0][3] = xyz[0];
  gps[1][3] = xyz[1];
  gps[2][3] = xyz[2];

  for (long i = 0; i < m_dem_size_; i++) {
    double v[3]; // 物体坐标系上的点
    // v是末端坐标系上的一个点
    v[0] = m_plate_[i][0];
    v[1] = m_plate_[i][1];
    v[2] = m_plate_[i][2];
    double r[3]; // 将点变换到空间坐标系下的表达法
    MatrixVectMult(gps, v, r);  // 把v从末端坐标系变换到笛卡尔坐标系下表示，结果为r
    double dx = r[0] - m_ancher_[i][0]; // 获取x轴上的变化量
    double dy = r[1] - m_ancher_[i][1]; // 获取y轴上的变化量
    double dz = r[2] - m_ancher_[i][2]; // 获取z轴上的变化量
    retv[i] = sqrt(dx * dx + dy * dy + dz * dz);
    if (fabs(retv[i]) * cos(m_max_pull_ang_) > fabs(dz)) {
      return i + 1;
    }
    double rr1 = m_ancher_[i][0] * m_ancher_[i][0] + m_ancher_[i][1] * m_ancher_[i][1];
    double rr2 = r[0] * r[0] + r[1] * r[1];
    if (rr1 < rr2) {
      return i + 11;
    }
  }
  retnum = m_dem_size_;
  return 0;
}