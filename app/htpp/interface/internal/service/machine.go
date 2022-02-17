package service

import (
	"context"
	interfaceV1 "htp-platform/api/htpp/interface/v1"
	"htp-platform/app/htpp/interface/internal/biz"
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
