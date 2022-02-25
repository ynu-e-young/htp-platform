package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Capture struct {
	Data   []byte
	Pixels int64
	Area   float64
}

type CaptureLog struct {
	Id        int64
	MachineId int64
	Pixels    int64
	Area      float64
	ImageName string
	OssUrl    string
}

type CaptureRepo interface {
	ReadOne(ctx context.Context, device int64) (*Capture, error)
	ReadAll(ctx context.Context) ([]*Capture, error)
	ReadOneWithBinary(ctx context.Context, device int64) (*Capture, error)
	ReadAllWithBinary(ctx context.Context) ([]*Capture, error)
	ReadOneWithBinaryAndCalArea(ctx context.Context, device int64) (*Capture, error)
	ReadAllWithBinaryAndCalArea(ctx context.Context) ([]*Capture, error)

	FindLogsByMachineId(ctx context.Context, machineId int64) ([]*CaptureLog, error)
	CreateLog(ctx context.Context, captureLog *CaptureLog) (*CaptureLog, error)
	GetLog(ctx context.Context, id int64) (*CaptureLog, error)
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

func (uc *CaptureUsecase) ReadOneWithBinary(ctx context.Context, device int64) (*Capture, error) {
	capture, err := uc.repo.ReadOneWithBinary(ctx, device)
	if err != nil {
		return nil, err
	}

	return capture, nil
}

func (uc *CaptureUsecase) ReadAllWithBinary(ctx context.Context) ([]*Capture, error) {
	captures, err := uc.repo.ReadAllWithBinary(ctx)
	if err != nil {
		return nil, err
	}

	return captures, nil
}

func (uc *CaptureUsecase) ReadOneWithBinaryAndCalArea(ctx context.Context, device int64) (*Capture, error) {
	capture, err := uc.repo.ReadOneWithBinaryAndCalArea(ctx, device)
	if err != nil {
		return nil, err
	}

	return capture, nil
}

func (uc *CaptureUsecase) ReadAllWithBinaryAndCalArea(ctx context.Context) ([]*Capture, error) {
	captures, err := uc.repo.ReadAllWithBinaryAndCalArea(ctx)
	if err != nil {
		return nil, err
	}

	return captures, nil
}

func (uc *CaptureUsecase) FindLogsByMachineId(ctx context.Context, machineId int64) ([]*CaptureLog, error) {
	return uc.repo.FindLogsByMachineId(ctx, machineId)
}

func (uc *CaptureUsecase) CreateLog(ctx context.Context, captureLog *CaptureLog) (*CaptureLog, error) {
	return uc.repo.CreateLog(ctx, captureLog)
}

func (uc *CaptureUsecase) GetLog(ctx context.Context, id int64) (*CaptureLog, error) {
	return uc.repo.GetLog(ctx, id)
}
