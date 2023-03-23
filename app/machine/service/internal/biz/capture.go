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

type CaptureSrc struct {
	Proc   []byte
	Src    []byte
	Pixels int64
	Area   float64
}

type CaptureLog struct {
	Id         int64
	MachineId  string
	Pixels     int64
	Area       float64
	SrcName    string
	ProcName   string
	SrcOssUrl  string
	ProcOssUrl string
}

type CaptureRepo interface {
	ReadOne(ctx context.Context, device int64) (*Capture, error)
	ReadAll(ctx context.Context) ([]*Capture, error)

	ReadOneWithBinary(ctx context.Context, device int64) (*Capture, error)
	ReadOneWithBinaryAndSrc(ctx context.Context, device int64) (*CaptureSrc, error)

	ReadAllWithBinary(ctx context.Context) ([]*Capture, error)
	ReadAllWithBinaryAndSrc(ctx context.Context) ([]*CaptureSrc, error)

	ReadOneWithBinaryAndCalArea(ctx context.Context, device int64) (*Capture, error)
	ReadOneWithBinaryAndCalAreaAndSrc(ctx context.Context, device int64) (*CaptureSrc, error)

	ReadAllWithBinaryAndCalArea(ctx context.Context) ([]*Capture, error)
	ReadAllWithBinaryAndCalAreaAndSrc(ctx context.Context) ([]*CaptureSrc, error)

	FindLogsByMachineId(ctx context.Context, machineId string) ([]*CaptureLog, error)
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
	return uc.repo.ReadOne(ctx, device)
}

func (uc *CaptureUsecase) ReadAll(ctx context.Context) ([]*Capture, error) {
	return uc.repo.ReadAll(ctx)
}

func (uc *CaptureUsecase) ReadOneWithBinary(ctx context.Context, device int64) (*Capture, error) {
	return uc.repo.ReadOneWithBinary(ctx, device)
}

func (uc *CaptureUsecase) ReadOneWithBinaryAndSrc(ctx context.Context, device int64) (*CaptureSrc, error) {
	return uc.repo.ReadOneWithBinaryAndSrc(ctx, device)
}

func (uc *CaptureUsecase) ReadAllWithBinary(ctx context.Context) ([]*Capture, error) {
	return uc.repo.ReadAllWithBinary(ctx)
}

func (uc *CaptureUsecase) ReadAllWithBinaryAndSrc(ctx context.Context) ([]*CaptureSrc, error) {
	return uc.repo.ReadAllWithBinaryAndSrc(ctx)
}

func (uc *CaptureUsecase) ReadOneWithBinaryAndCalArea(ctx context.Context, device int64) (*Capture, error) {
	return uc.repo.ReadOneWithBinaryAndCalArea(ctx, device)
}

func (uc *CaptureUsecase) ReadOneWithBinaryAndCalAreaAndSrc(ctx context.Context, device int64) (*CaptureSrc, error) {
	return uc.repo.ReadOneWithBinaryAndCalAreaAndSrc(ctx, device)
}

func (uc *CaptureUsecase) ReadAllWithBinaryAndCalArea(ctx context.Context) ([]*Capture, error) {
	return uc.repo.ReadAllWithBinaryAndCalArea(ctx)
}

func (uc *CaptureUsecase) ReadAllWithBinaryAndCalAreaAndSrc(ctx context.Context) ([]*CaptureSrc, error) {
	return uc.repo.ReadAllWithBinaryAndCalAreaAndSrc(ctx)
}

func (uc *CaptureUsecase) FindLogsByMachineId(ctx context.Context, machineId string) ([]*CaptureLog, error) {
	return uc.repo.FindLogsByMachineId(ctx, machineId)
}

func (uc *CaptureUsecase) CreateLog(ctx context.Context, captureLog *CaptureLog) (*CaptureLog, error) {
	return uc.repo.CreateLog(ctx, captureLog)
}

func (uc *CaptureUsecase) GetLog(ctx context.Context, id int64) (*CaptureLog, error) {
	return uc.repo.GetLog(ctx, id)
}
