package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gocv.io/x/gocv"
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
