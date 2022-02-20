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

type Coordinate struct {
	X         float64
	Y         float64
	Z         float64
	Rx        float64
	Ry        float64
	Check     bool
	Delay     float64
	MachineId int64
	CheckName string
}

type MotorInfo struct {
	MotorStatus *MotorStatus
	InstrPos    int64
	CurrentPos  int64
}

type MotorStatus struct {
	Fault                 bool
	Enabling              bool
	Running               bool
	InstructionCompletion bool
	PathCompletion        bool
	ZeroCompletion        bool
}

type MachineRepo interface {
	FindByUserId(ctx context.Context, userId int64) ([]*Machine, error)
	Create(ctx context.Context, machine *Machine) (*Machine, error)
	Update(ctx context.Context, machine *Machine) (*Machine, error)
	Get(ctx context.Context, machineId int64) (*Machine, error)

	Move(ctx context.Context, coordinate *Coordinate) (bool, error)
	Zero(ctx context.Context, machineId int64) (bool, error)
	GetMotorStatus(ctx context.Context, machineId int64) ([]*MotorInfo, error)
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

func (uc *MachineUsecase) Move(ctx context.Context, coordinate *Coordinate) (bool, error) {
	return uc.repo.Move(ctx, coordinate)
}

func (uc *MachineUsecase) Zero(ctx context.Context, machineId int64) (bool, error) {
	return uc.repo.Zero(ctx, machineId)
}

func (uc *MachineUsecase) GetMotorStatus(ctx context.Context, machineId int64) ([]*MotorInfo, error) {
	return uc.repo.GetMotorStatus(ctx, machineId)
}
