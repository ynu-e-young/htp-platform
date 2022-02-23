package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
	machineV1 "htp-platform/api/machine/service/v1"
	"htp-platform/app/machine/service/internal/biz"
	"htp-platform/app/machine/service/internal/data/ent"
	"htp-platform/app/machine/service/internal/data/ent/cronjob"
	"strconv"
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
	machineIdStr := strconv.FormatInt(c.MachineId, 10)
	checkName := c.CheckName

	if _, ok := r.data.cr[machineIdStr+":"+checkName]; ok {
		return machineV1.ErrorCronConflict("cron job already exist, machine id: %s, check name: %s", machineIdStr, checkName)
	} else {
		cr := cron.New()

		// 设置定时器
		if _, err := cr.AddFunc(c.CronString, jobFunc); err != nil {
			return machineV1.ErrorCronSetupFailed("set up cron job failed, cron string: %s", c.CronString)
		}

		r.data.cr[machineIdStr+":"+checkName] = cr
		cr.Start()
	}

	return nil
}

func (r *cronRepo) DelCronJob(ctx context.Context, machineId int64, checkName string) error {
	machineIdStr := strconv.FormatInt(machineId, 10)

	if _, ok := r.data.cr[machineIdStr+":"+checkName]; ok {
		// 停止定时任务然后删除
		r.data.cr[machineIdStr+":"+checkName].Stop()
		delete(r.data.cr, machineIdStr+":"+checkName)
	} else {
		return machineV1.ErrorNotFoundError("cron job not found, machine id: %s, check name: %s", machineIdStr, checkName)
	}

	return nil
}

func (r *cronRepo) FindByMachineId(ctx context.Context, machineId int64) ([]*biz.Cron, error) {
	targets, err := r.data.db.CronJob.
		Query().
		Where(cronjob.MachineIDEQ(machineId)).
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
			MachineId:        target.MachineID,
			CronString:       target.CronString,
			ID:               target.ID,
		})
	}

	return bcs, nil
}

func (r *cronRepo) Create(ctx context.Context, cr *biz.Cron) (*biz.Cron, error) {
	if po, err := r.data.db.CronJob.
		Query().
		Where(cronjob.And(cronjob.MachineIDEQ(cr.MachineId), cronjob.CheckNameEQ(cr.CheckName))).
		All(ctx);
		len(po) == 0 || (err != nil && ent.IsNotFound(err)) {
		po, err := r.data.db.CronJob.
			Create().
			SetMachineID(cr.MachineId).
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
			MachineId:        po.MachineID,
			CronString:       po.CronString,
			ID:               po.ID,
		}, nil
	} else {
		return nil, machineV1.ErrorCreateFailed("check name: %s already exist", cr.CheckName)
	}
}

func (r *cronRepo) Update(ctx context.Context, cr *biz.Cron) (*biz.Cron, error) {
	po, err := r.data.db.CronJob.
		UpdateOneID(cr.ID).
		SetMachineID(cr.MachineId).
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
		MachineId:        po.MachineID,
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
		MachineId:        po.MachineID,
		CronString:       po.CronString,
		ID:               po.ID,
	}, nil
}
