package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Machine struct {
	MachineId int64
	UserId    int64
	Address   string
}

type MachineRepo interface {
	FindByUserId(ctx context.Context, userId int64) ([]*Machine, error)
	Create(ctx context.Context, machine *Machine) (*Machine, error)
	Update(ctx context.Context, machine *Machine) (*Machine, error)
	Get(ctx context.Context, machineId int64) (*Machine, error)
}

type MachineUsecase struct {
	repo MachineRepo

	log *log.Helper
}

func NewMachineUsecase(repo MachineRepo, logger log.Logger) *MachineUsecase {
	return &MachineUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (mu *MachineUsecase) FindByUserId(ctx context.Context, userId int64) ([]*Machine, error) {
	return mu.repo.FindByUserId(ctx, userId)
}

func (mu *MachineUsecase) Create(ctx context.Context, machine *Machine) (*Machine, error) {
	return mu.repo.Create(ctx, machine)
}

func (mu *MachineUsecase) Update(ctx context.Context, machine *Machine) (*Machine, error) {
	return mu.repo.Update(ctx, machine)
}

func (mu *MachineUsecase) Get(ctx context.Context, machineId int64) (*Machine, error) {
	return mu.repo.Get(ctx, machineId)
}
