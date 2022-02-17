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

func (uc *MachineUsecase) RegisterMachine(ctx context.Context, machine *Machine) (*Machine, error) {
	m, err := uc.repo.Create(ctx, machine)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (uc *MachineUsecase) UpdateMachine(ctx context.Context, machine *Machine) (*Machine, error) {
	m, err := uc.repo.Update(ctx, machine)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (uc *MachineUsecase) GetCurrentUserMachines(ctx context.Context, userId int64) ([]*Machine, error) {
	ms, err := uc.repo.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return ms, nil
}

func (uc *MachineUsecase) GetMachine(ctx context.Context, machineId int64) (*Machine, error) {
	m, err := uc.repo.Get(ctx, machineId)
	if err != nil {
		return nil, err
	}

	return m, nil
}
