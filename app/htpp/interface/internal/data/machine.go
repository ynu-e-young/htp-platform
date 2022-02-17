package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	interfaceV1 "htp-platform/api/htpp/interface/v1"
	machineV1 "htp-platform/api/machine/service/v1"
	"htp-platform/app/htpp/interface/internal/biz"
)

var _ biz.MachineRepo = (*machineRepo)(nil)

type machineRepo struct {
	data *Data
	log  *log.Helper
}

func NewMachineRepo(data *Data, logger log.Logger) biz.MachineRepo {
	return &machineRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/machine")),
	}
}

func (r *machineRepo) FindByUserId(ctx context.Context, userId int64) ([]*biz.Machine, error) {
	reply, err := r.data.mc.FindByUserId(ctx, &machineV1.FindByUserIdRequest{
		UserId: userId,
	})
	if err != nil {
		return nil, machineV1.ErrorNotFoundError("find userId: %d not found, err: %v", userId, err)
	}

	var ret []*biz.Machine
	for _, m := range reply.Machines {
		ret = append(ret, &biz.Machine{
			MachineId: m.MachineId,
			UserId:    m.UserId,
			Address:   m.Address,
		})
	}

	return ret, nil
}

func (r *machineRepo) Create(ctx context.Context, machine *biz.Machine) (*biz.Machine, error) {
	reply, err := r.data.mc.Create(ctx, &machineV1.CreateRequest{
		Machine: &machineV1.MachineStruct{
			MachineId: machine.MachineId,
			UserId:    machine.UserId,
			Address:   machine.Address,
		}})
	if err != nil {
		return nil, interfaceV1.ErrorAddressConflict("create machine conflict, err: %v", err)
	}

	m := reply.GetMachine()

	return &biz.Machine{
		MachineId: m.MachineId,
		UserId:    m.UserId,
		Address:   m.Address,
	}, nil
}

func (r *machineRepo) Update(ctx context.Context, machine *biz.Machine) (*biz.Machine, error) {
	reply, err := r.data.mc.Update(ctx, &machineV1.UpdateRequest{
		Machine: &machineV1.MachineStruct{
			MachineId: machine.MachineId,
			UserId:    machine.UserId,
			Address:   machine.Address,
		}})
	if err != nil {
		return nil, interfaceV1.ErrorAddressConflict("update machine conflict, err: %v", err)
	}

	m := reply.GetMachine()

	return &biz.Machine{
		MachineId: m.MachineId,
		UserId:    m.UserId,
		Address:   m.Address,
	}, nil
}

func (r *machineRepo) Get(ctx context.Context, machineId int64) (*biz.Machine, error) {
	reply, err := r.data.mc.Get(ctx, &machineV1.GetRequest{
		MachineId: machineId,
	})
	if err != nil {
		return nil, interfaceV1.ErrorNotFoundError("find id: %s not found, err: %v", machineId, err)
	}

	m := reply.GetMachine()

	return &biz.Machine{
		MachineId: m.MachineId,
		UserId:    m.UserId,
		Address:   m.Address,
	}, nil
}
