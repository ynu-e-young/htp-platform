package service

import (
	"context"
	v1 "htp-platform/api/machine/service/v1"
	"htp-platform/app/machine/service/internal/biz"
)

func (s *MachineService) FindByUserId(ctx context.Context, in *v1.FindByUserIdRequest) (*v1.MachinesReply, error) {
	ms, err := s.mu.FindByUserId(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}

	var ret []*v1.MachineStruct
	for _, m := range ms {
		ret = append(ret, &v1.MachineStruct{
			MachineId: m.MachineId,
			UserId:    m.UserId,
			Address:   m.Address,
		})
	}

	return &v1.MachinesReply{
		Machines: ret,
	}, nil
}

func (s *MachineService) Create(ctx context.Context, in *v1.CreateRequest) (*v1.MachineReply, error) {
	machine := in.GetMachine()
	m, err := s.mu.Create(ctx, &biz.Machine{
		UserId:  machine.UserId,
		Address: machine.Address,
	})
	if err != nil {
		return nil, err
	}

	return &v1.MachineReply{Machine: &v1.MachineStruct{
		MachineId: m.MachineId,
		UserId:    m.UserId,
		Address:   m.Address,
	}}, nil
}

func (s *MachineService) Update(ctx context.Context, in *v1.UpdateRequest) (*v1.MachineReply, error) {
	machine := in.GetMachine()
	m, err := s.mu.Update(ctx, &biz.Machine{
		MachineId: machine.MachineId,
		UserId:    machine.UserId,
		Address:   machine.Address,
	})
	if err != nil {
		return nil, err
	}

	return &v1.MachineReply{Machine: &v1.MachineStruct{
		MachineId: m.MachineId,
		UserId:    m.UserId,
		Address:   m.Address,
	}}, nil
}

func (s *MachineService) Get(ctx context.Context, in *v1.GetRequest) (*v1.MachineReply, error) {
	m, err := s.mu.Get(ctx, in.GetMachineId())
	if err != nil {
		return nil, err
	}

	return &v1.MachineReply{Machine: &v1.MachineStruct{
		MachineId: m.MachineId,
		UserId:    m.UserId,
		Address:   m.Address,
	}}, nil
}
