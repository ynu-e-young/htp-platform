package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Capture struct {
	Data []byte
}

type CaptureRepo interface {
	ReadOne(ctx context.Context, device int64) (*Capture, error)
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

func (uc *CaptureUsecase) ReadOne(ctx context.Context, device int64) (*Capture, error) {
	capture, err := uc.repo.ReadOne(ctx, device)
	if err != nil {
		return nil, err
	}

	return capture, nil
}

func (uc *CaptureUsecase) ReadAll(ctx context.Context) ([]*Capture, error) {
	captures, err := uc.repo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}

	return captures, nil
}
