package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gocv.io/x/gocv"
	"math"
)

type Capture struct {
	Mat *gocv.Mat
}

type CaptureRepo interface {
	ReadOne(ctx context.Context, device int) (*Capture, error)
	ReadAll(ctx context.Context) ([]*Capture, error)
}

type CaptureUsecase struct {
	repo CaptureRepo

	log *log.Helper
}

func NewCaptureUsecase(repo CaptureRepo, logger log.Logger) *CaptureUsecase {
	return &CaptureUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *CaptureUsecase) ReadOne(ctx context.Context, device int) (*Capture, error) {
	return uc.repo.ReadOne(ctx, device)
}

func (uc *CaptureUsecase) ReadAll(ctx context.Context) ([]*Capture, error) {
	return uc.repo.ReadAll(ctx)
}

func (uc *CaptureUsecase) Binary(ctx context.Context, capture *Capture) (*Capture, error) {
	img := capture.Mat
	tmp := gocv.NewMatWithSizes(img.Size(), img.Type())

	// get s channel from HSV channel
	gocv.CvtColor(*img, &tmp, gocv.ColorRGBToHSV)
	HsvS := gocv.Split(tmp)[1]

	sThresh := gocv.NewMat()
	gocv.Threshold(HsvS, &sThresh, 85, 255, gocv.ThresholdBinary)

	// 中值滤波
	sMBlur := gocv.NewMat()
	gocv.MedianBlur(sThresh, &sMBlur, 5)

	// get b channel from LAB channel
	gocv.CvtColor(*img, &tmp, gocv.ColorRGBToLab)
	LabB := gocv.Split(tmp)[2]

	bThresh := gocv.NewMat()
	gocv.Threshold(LabB, &bThresh, 160, 255, gocv.ThresholdBinary)

	bs := gocv.NewMat()
	gocv.BitwiseOr(sMBlur, bThresh, &bs)

	uc.applyMask(*img, bs, &tmp)

	// get a channel from LAB channel
	gocv.CvtColor(tmp, &tmp, gocv.ColorRGBToLab)
	LabA := gocv.Split(tmp)[1]

	gocv.Threshold(LabA, &tmp, 115, 255, gocv.ThresholdBinaryInv)

	return &Capture{
		Mat: &tmp,
	}, nil
}

func (uc *CaptureUsecase) BaweraOpen(grayImg gocv.Mat, thresh int) gocv.Mat {
	ret := grayImg.Clone()

	labels := gocv.NewMat()
	stats := gocv.NewMat()
	centroids := gocv.NewMat()

	// labels: 和大小相同的标记图,
	// stats: components * 5 的矩阵, 表示每个连通区域的外接矩形和面积（pixel）
	// centroids: components * 2 的矩阵, 表示每个连通区域的质心
	components := gocv.ConnectedComponentsWithStats(grayImg, &labels, &stats, &centroids)

	for i := 1; i < components-1; i++ {
		if int(stats.GetUCharAt(i, 4)) < thresh {
			x0 := int(stats.GetUCharAt(i, 0))
			y0 := int(stats.GetUCharAt(i, 1))
			x1 := x0 + int(stats.GetUCharAt(i, 2))
			y1 := y0 + int(stats.GetUCharAt(i, 3))

			for row := y0; row < y1; row++ {
				for col := x0; col < x1; col++ {
					if int(labels.GetUCharAt(row, col)) == i {
						ret.SetUCharAt(row, col, 0)
					}
				}
			}
		}
	}
	return ret
}

func (uc *CaptureUsecase) Area(pixels, distanceTop float64) float64 {
	return math.Pow(0.000122*(distanceTop-0.304)/0.304, 2) * pixels
}

func (uc *CaptureUsecase) Pixels(grayImg gocv.Mat) int {
	return gocv.CountNonZero(grayImg)
}

func (uc *CaptureUsecase) applyMask(src, mask gocv.Mat, dst *gocv.Mat) {
	maskInv := gocv.NewMatWithSizes(mask.Size(), mask.Type())
	gocv.BitwiseNot(mask, &maskInv)

	mask3 := gocv.NewMatWithSizes(src.Size(), src.Type())
	gocv.Merge([]gocv.Mat{mask, mask, mask}, &mask3)

	mask3Inv := gocv.NewMatWithSizes(src.Size(), src.Type())
	gocv.Merge([]gocv.Mat{maskInv, maskInv, maskInv}, &mask3Inv)

	gocv.BitwiseAnd(src, mask3, dst)
	gocv.BitwiseXor(*dst, mask3Inv, dst)
}
