package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Cron struct {
	CheckCoordinates []*CheckCoordinate
	CheckName        string
	MachineId        int64
	CronString       string
	ID               int64
}

type CheckCoordinate struct {
	Seq int64       `json:"seq,omitempty"`
	Crd *Coordinate `json:"crd,omitempty"`
}

type Coordinate struct {
	X         float64 `json:"x,omitempty"`
	Y         float64 `json:"y,omitempty"`
	Z         float64 `json:"z,omitempty"`
	Rx        float64 `json:"rx,omitempty"`
	Ry        float64 `json:"ry,omitempty"`
	Check     bool    `json:"check,omitempty"`
	Delay     float64 `json:"delay,omitempty"`
	MachineId int64   `json:"machine_id,omitempty"`
	CheckName string  `json:"check_name,omitempty"`
}

type CronRepo interface {
	AddCronJob(ctx context.Context, c *Cron, jobFunc func()) error
	DelCronJob(ctx context.Context, machineId int64, checkName string) error

	FindByMachineId(ctx context.Context, machineId int64) ([]*Cron, error)
	Create(ctx context.Context, cr *Cron) (*Cron, error)
	Update(ctx context.Context, cr *Cron) (*Cron, error)
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*Cron, error)
}

type CronUsecase struct {
	repo CronRepo

	log *log.Helper
}

func NewCronUsecase(repo CronRepo, logger log.Logger) *CronUsecase {
	return &CronUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *CronUsecase) AddCronJob(ctx context.Context, c *Cron, jobFunc func()) error {
	return uc.repo.AddCronJob(ctx, c, jobFunc)
}

func (uc *CronUsecase) DelCronJob(ctx context.Context, machineId int64, checkName string) error {
	return uc.repo.DelCronJob(ctx, machineId, checkName)
}

func (uc *CronUsecase) FindByMachineId(ctx context.Context, machineId int64) ([]*Cron, error) {
	return uc.repo.FindByMachineId(ctx, machineId)
}

func (uc *CronUsecase) Create(ctx context.Context, cr *Cron) (*Cron, error) {
	return uc.repo.Create(ctx, cr)
}

func (uc *CronUsecase) Update(ctx context.Context, cr *Cron) (*Cron, error) {
	return uc.repo.Update(ctx, cr)
}

func (uc *CronUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *CronUsecase) Get(ctx context.Context, id int64) (*Cron, error) {
	return uc.repo.Get(ctx, id)
}
