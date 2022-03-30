package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	machineV1 "htp-platform/api/machine/service/v1"
	"htp-platform/app/machine/service/internal/biz"
	"htp-platform/app/machine/service/internal/data/ent"
	"htp-platform/app/machine/service/internal/data/ent/machine"
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
	targets, err := r.data.db.Machine.
		Query().
		Where(machine.UserIDEQ(userId)).
		All(ctx)

	if err != nil && ent.IsNotFound(err) {
		return nil, machineV1.ErrorNotFoundError("find userId: %d not found, err: %v", userId, err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, machineV1.ErrorUnknownError("unknown err: %v", err)
	}

	var machines []*biz.Machine
	for _, target := range targets {
		machines = append(machines, &biz.Machine{
			MachineId: target.ID.String(),
			UserId:    target.UserID,
			Address:   target.Address,
		})
	}

	return machines, nil
}

func (r *machineRepo) Create(ctx context.Context, machine *biz.Machine) (*biz.Machine, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, machineV1.ErrorUuidGenerateFailed("create machine uuid failed, err: %v", err)
	}

	po, err := r.data.db.Machine.
		Create().
		SetID(u).
		SetUserID(machine.UserId).
		SetAddress(machine.Address).
		Save(ctx)
	if err != nil && ent.IsConstraintError(err) {
		return nil, machineV1.ErrorAddressConflict("create machine conflict, err: %v", err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, machineV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.Machine{
		MachineId: po.ID.String(),
		UserId:    po.UserID,
		Address:   po.Address,
	}, nil
}

func (r *machineRepo) Update(ctx context.Context, machine *biz.Machine) (*biz.Machine, error) {
	u, err := uuid.Parse(machine.MachineId)
	if err != nil {
		return nil, machineV1.ErrorUuidParseFailed("update machine conflict, err: %v", err)
	}

	po, err := r.data.db.Machine.
		UpdateOneID(u).
		SetUserID(machine.UserId).
		SetAddress(machine.Address).
		Save(ctx)
	if err != nil && ent.IsConstraintError(err) {
		return nil, machineV1.ErrorAddressConflict("update machine conflict, err: %v", err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, machineV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.Machine{
		MachineId: po.ID.String(),
		UserId:    po.UserID,
		Address:   po.Address,
	}, nil
}

func (r *machineRepo) Get(ctx context.Context, machineId string) (*biz.Machine, error) {
	u, err := uuid.Parse(machineId)
	if err != nil {
		return nil, machineV1.ErrorUuidParseFailed("update machine conflict, err: %v", err)
	}

	po, err := r.data.db.Machine.Get(ctx, u)
	if err != nil && ent.IsNotFound(err) {
		return nil, machineV1.ErrorNotFoundError("find id: %s not found, err: %v", machineId, err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, machineV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.Machine{
		MachineId: po.ID.String(),
		UserId:    po.UserID,
		Address:   po.Address,
	}, nil
}
