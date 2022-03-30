package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CronJob struct {
	CheckCoordinates []*CheckCoordinate
	CheckName        string
	MachineId        string
	CronString       string
	ID               int64
}

type CheckCoordinate struct {
	Seq int64       `json:"seq,omitempty"`
	Crd *Coordinate `json:"crd,omitempty"`
}

type CronJobRepo interface {
	Create(ctx context.Context, cj *CronJob) (*CronJob, error)
	Delete(ctx context.Context, id int64) (int64, error)
	List(ctx context.Context, machineId string) ([]*CronJob, error)
}

type CronJobUsecase struct {
	repo CronJobRepo

	log *log.Helper
}

func NewCronJobUsecase(repo CronJobRepo, logger log.Logger) *CronJobUsecase {
	return &CronJobUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (cju *CronJobUsecase) Create(ctx context.Context, cj *CronJob) (*CronJob, error) {
	return cju.repo.Create(ctx, cj)
}

func (cju *CronJobUsecase) Delete(ctx context.Context, id int64) (int64, error) {
	return cju.repo.Delete(ctx, id)
}

func (cju *CronJobUsecase) List(ctx context.Context, machineId string) ([]*CronJob, error) {
	return cju.repo.List(ctx, machineId)
}
