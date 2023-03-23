package service

import (
	"context"

	interfaceV1 "github.com/ynu-e-young/apis-go/htpp/htpp/interface/v1"

	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/biz"
)

func (s *InterfaceService) RegisterMachine(ctx context.Context, in *interfaceV1.RegisterMachineRequest) (*interfaceV1.MachineReply, error) {
	m := in.GetMachine()
	rv, err := s.mu.RegisterMachine(ctx, &biz.Machine{
		MachineId: m.MachineId,
		UserId:    m.UserId,
		Address:   m.Address,
	})
	if err != nil {
		return nil, err
	}

	return &interfaceV1.MachineReply{Machine: &interfaceV1.MachineStruct{
		MachineId: rv.MachineId,
		UserId:    rv.UserId,
		Address:   rv.Address,
	}}, err
}

func (s *InterfaceService) UpdateMachine(ctx context.Context, in *interfaceV1.UpdateMachineRequest) (*interfaceV1.MachineReply, error) {
	m := in.GetMachine()
	rv, err := s.mu.UpdateMachine(ctx, &biz.Machine{
		MachineId: m.MachineId,
		UserId:    m.UserId,
		Address:   m.Address,
	})
	if err != nil {
		return nil, err
	}

	return &interfaceV1.MachineReply{Machine: &interfaceV1.MachineStruct{
		MachineId: rv.MachineId,
		UserId:    rv.UserId,
		Address:   rv.Address,
	}}, err
}

func (s *InterfaceService) GetMachine(ctx context.Context, in *interfaceV1.GetMachineRequest) (*interfaceV1.MachineReply, error) {
	rv, err := s.mu.GetMachine(ctx, in.GetMachineId())
	if err != nil {
		return nil, err
	}

	return &interfaceV1.MachineReply{Machine: &interfaceV1.MachineStruct{
		MachineId: rv.MachineId,
		UserId:    rv.UserId,
		Address:   rv.Address,
	}}, err
}

func (s *InterfaceService) GetCurrentUserMachines(ctx context.Context, in *interfaceV1.GetCurrentUserMachinesRequest) (*interfaceV1.MachinesReply, error) {
	rv, err := s.mu.GetCurrentUserMachines(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}

	var ms []*interfaceV1.MachineStruct
	for _, m := range rv {
		ms = append(ms, &interfaceV1.MachineStruct{
			MachineId: m.MachineId,
			UserId:    m.UserId,
			Address:   m.Address,
		})
	}

	return &interfaceV1.MachinesReply{Machines: ms}, err
}

func (s *InterfaceService) Move(ctx context.Context, in *interfaceV1.MoveRequest) (*interfaceV1.MoveReply, error) {
	ok, err := s.mu.Move(ctx, &biz.Coordinate{
		X:         in.GetX(),
		Y:         in.GetY(),
		Z:         in.GetZ(),
		Rx:        in.GetRx(),
		Ry:        in.GetRy(),
		Check:     in.GetCheck(),
		Delay:     in.GetDelay(),
		MachineId: in.GetMachineId(),
		CheckName: in.GetCheckName(),
	})
	if err != nil {
		return nil, err
	}

	return &interfaceV1.MoveReply{Status: ok}, nil
}

func (s *InterfaceService) Zero(ctx context.Context, in *interfaceV1.ZeroRequest) (*interfaceV1.ZeroReply, error) {
	ok, err := s.mu.Zero(ctx, in.GetMachineId())
	if err != nil {
		return nil, err
	}

	return &interfaceV1.ZeroReply{Status: ok}, nil
}

func (s *InterfaceService) GetMotorStatus(ctx context.Context, in *interfaceV1.GetMotorStatusRequest) (*interfaceV1.GetMotorStatusReply, error) {
	rv, err := s.mu.GetMotorStatus(ctx, in.GetMachineId())
	if err != nil {
		return nil, err
	}

	var mis []*interfaceV1.MotorInfo
	for _, info := range rv {
		mStatus := info.MotorStatus
		mis = append(mis, &interfaceV1.MotorInfo{
			MotorStatus: &interfaceV1.MotorStatus{
				Fault:                 mStatus.Fault,
				Enabling:              mStatus.Enabling,
				Running:               mStatus.Running,
				InstructionCompletion: mStatus.InstructionCompletion,
				PathCompletion:        mStatus.PathCompletion,
				ZeroCompletion:        mStatus.ZeroCompletion,
			},
			InstrPos:   info.InstrPos,
			CurrentPos: info.CurrentPos,
		})
	}

	return &interfaceV1.GetMotorStatusReply{MotorInfo: mis}, nil
}
