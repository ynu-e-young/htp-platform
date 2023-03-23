package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	machineV1 "github.com/ynu-e-young/apis-go/htpp/machine/service/v1"

	"github.com/ynu-e-young/htp-platform/app/machine/service/internal/biz"
	"github.com/ynu-e-young/htp-platform/app/machine/service/internal/data/ent"
	"github.com/ynu-e-young/htp-platform/app/machine/service/internal/data/ent/cronjob"
)

var _ biz.CronRepo = (*cronRepo)(nil)

type cronRepo struct {
	data *Data
	log  *log.Helper
}

func NewCronRepo(data *Data, logger log.Logger) biz.CronRepo {
	return &cronRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/cron")),
	}
}

func (r *cronRepo) AddCronJob(ctx context.Context, c *biz.Cron, jobFunc func()) error {
	checkName := c.CheckName

	if _, ok := r.data.cr[c.MachineId+":"+checkName]; ok {
		return machineV1.ErrorCronConflict("cron job already exist, machine id: %s, check name: %s", c.MachineId, checkName)
	} else {
		cr := cron.New()

		// 设置定时器
		if _, err := cr.AddFunc(c.CronString, jobFunc); err != nil {
			return machineV1.ErrorCronSetupFailed("set up cron job failed, cron string: %s", c.CronString)
		}

		r.data.cr[c.MachineId+":"+checkName] = cr
		cr.Start()
	}

	return nil
}

func (r *cronRepo) DelCronJob(ctx context.Context, machineId, checkName string) error {
	if _, ok := r.data.cr[machineId+":"+checkName]; ok {
		// 停止定时任务然后删除
		r.data.cr[machineId+":"+checkName].Stop()
		delete(r.data.cr, machineId+":"+checkName)
	} else {
		return machineV1.ErrorNotFoundError("cron job not found, machine id: %s, check name: %s", machineId, checkName)
	}

	return nil
}

func (r *cronRepo) FindByMachineId(ctx context.Context, machineId string) ([]*biz.Cron, error) {
	u, err := uuid.Parse(machineId)
	if err != nil {
		return nil, machineV1.ErrorUuidParseFailed("update machine conflict, err: %v", err)
	}

	targets, err := r.data.db.CronJob.
		Query().
		Where(cronjob.MachineIDEQ(u)).
		All(ctx)

	if err != nil && ent.IsNotFound(err) {
		return nil, machineV1.ErrorNotFoundError("find machineId: %d not found, err: %v", machineId, err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, machineV1.ErrorUnknownError("unknown err: %v", err)
	}

	var bcs []*biz.Cron
	for _, target := range targets {
		bcs = append(bcs, &biz.Cron{
			CheckCoordinates: target.Coordinates,
			CheckName:        target.CheckName,
			MachineId:        target.MachineID.String(),
			CronString:       target.CronString,
			ID:               target.ID,
		})
	}

	return bcs, nil
}

func (r *cronRepo) Create(ctx context.Context, cr *biz.Cron) (*biz.Cron, error) {
	u, err := uuid.Parse(cr.MachineId)
	if err != nil {
		return nil, machineV1.ErrorUuidParseFailed("update machine conflict, err: %v", err)
	}

	if po, err := r.data.db.CronJob.
		Query().
		Where(cronjob.And(cronjob.MachineIDEQ(u), cronjob.CheckNameEQ(cr.CheckName))).
		All(ctx); len(po) == 0 || (err != nil && ent.IsNotFound(err)) {
		poi, err := r.data.db.CronJob.
			Create().
			SetMachineID(u).
			SetCheckName(cr.CheckName).
			SetCronString(cr.CronString).
			SetCoordinates(cr.CheckCoordinates).
			Save(ctx)
		if err != nil && ent.IsConstraintError(err) {
			return nil, machineV1.ErrorAddressConflict("create cronjob conflict, err: %v", err)
		}
		if err != nil {
			r.log.Errorf("unknown err: %v", err)
			return nil, machineV1.ErrorUnknownError("unknown err: %v", err)
		}

		return &biz.Cron{
			CheckCoordinates: poi.Coordinates,
			CheckName:        poi.CheckName,
			MachineId:        poi.MachineID.String(),
			CronString:       poi.CronString,
			ID:               poi.ID,
		}, nil
	} else {
		return nil, machineV1.ErrorCreateFailed("check name: %s already exist", cr.CheckName)
	}
}

func (r *cronRepo) Update(ctx context.Context, cr *biz.Cron) (*biz.Cron, error) {
	u, err := uuid.Parse(cr.MachineId)
	if err != nil {
		return nil, machineV1.ErrorUuidParseFailed("update machine conflict, err: %v", err)
	}

	po, err := r.data.db.CronJob.
		UpdateOneID(cr.ID).
		SetMachineID(u).
		SetCheckName(cr.CheckName).
		SetCronString(cr.CronString).
		SetCoordinates(cr.CheckCoordinates).
		Save(ctx)
	if err != nil && ent.IsConstraintError(err) {
		return nil, machineV1.ErrorAddressConflict("create cronjob conflict, err: %v", err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, machineV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.Cron{
		CheckCoordinates: po.Coordinates,
		CheckName:        po.CheckName,
		MachineId:        po.MachineID.String(),
		CronString:       po.CronString,
		ID:               po.ID,
	}, nil
}

func (r *cronRepo) Delete(ctx context.Context, id int64) error {
	err := r.data.db.CronJob.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return machineV1.ErrorUnknownError("unknown err: %v", err)
	}

	return nil
}

func (r *cronRepo) Get(ctx context.Context, id int64) (*biz.Cron, error) {
	po, err := r.data.db.CronJob.Get(ctx, id)
	if err != nil && ent.IsNotFound(err) {
		return nil, machineV1.ErrorNotFoundError("find Id: %d not found, err: %v", id, err)
	}
	if err != nil {
		r.log.Errorf("unknown err: %v", err)
		return nil, machineV1.ErrorUnknownError("unknown err: %v", err)
	}

	return &biz.Cron{
		CheckCoordinates: po.Coordinates,
		CheckName:        po.CheckName,
		MachineId:        po.MachineID.String(),
		CronString:       po.CronString,
		ID:               po.ID,
	}, nil
}
