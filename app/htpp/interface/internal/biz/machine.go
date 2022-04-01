package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Machine struct {
	MachineId string
	UserId    int64
	Address   string
}

type Coordinate struct {
	X         float64 `json:"x,omitempty"`
	Y         float64 `json:"y,omitempty"`
	Z         float64 `json:"z,omitempty"`
	Rx        float64 `json:"rx,omitempty"`
	Ry        float64 `json:"ry,omitempty"`
	Check     bool    `json:"check,omitempty"`
	Delay     float64 `json:"delay,omitempty"`
	MachineId string  `json:"machine_id,omitempty"`
	CheckName string  `json:"check_name,omitempty"`
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
	Get(ctx context.Context, machineId string) (*Machine, error)

	Move(ctx context.Context, coordinate *Coordinate) (bool, error)
	Zero(ctx context.Context, machineId string) (bool, error)
	GetMotorStatus(ctx context.Context, machineId string) ([]*MotorInfo, error)
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

func (uc *MachineUsecase) GetMachine(ctx context.Context, machineId string) (*Machine, error) {
	m, err := uc.repo.Get(ctx, machineId)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (uc *MachineUsecase) Move(ctx context.Context, coordinate *Coordinate) (bool, error) {
	return uc.repo.Move(ctx, coordinate)
}

func (uc *MachineUsecase) Zero(ctx context.Context, machineId string) (bool, error) {
	return uc.repo.Zero(ctx, machineId)
}

func (uc *MachineUsecase) GetMotorStatus(ctx context.Context, machineId string) ([]*MotorInfo, error) {
	return uc.repo.GetMotorStatus(ctx, machineId)
}
