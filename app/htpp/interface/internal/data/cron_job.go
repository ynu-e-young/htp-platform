package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	machineV1 "htp-platform/api/machine/service/v1"
	"htp-platform/app/htpp/interface/internal/biz"
)

var _ biz.CronJobRepo = (*cronJobRepo)(nil)

type cronJobRepo struct {
	data *Data
	log  *log.Helper
}

func NewCronJobRepo(data *Data, logger log.Logger) biz.CronJobRepo {
	return &cronJobRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/cronJob")),
	}
}

func (r *cronJobRepo) Create(ctx context.Context, cj *biz.CronJob) (*biz.CronJob, error) {
	var ccs []*machineV1.CheckCoordinate
	for _, ccd := range cj.CheckCoordinates {
		mCrd := ccd.Crd
		crd := &machineV1.Coordinate{
			X:         mCrd.X,
			Y:         mCrd.Y,
			Z:         mCrd.Z,
			Rx:        mCrd.Rx,
			Ry:        mCrd.Ry,
			Check:     mCrd.Check,
			Delay:     mCrd.Delay,
			MachineId: mCrd.MachineId,
			CheckName: mCrd.CheckName,
		}
		ccs = append(ccs, &machineV1.CheckCoordinate{
			Seq: ccd.Seq,
			Crd: crd,
		})
	}

	reply, err := r.data.mc.CreateCronJob(ctx, &machineV1.CreateCronJobRequest{CronJob: &machineV1.CronJob{
		MachineId:        cj.MachineId,
		CheckName:        cj.CheckName,
		CronString:       cj.CronString,
		CheckCoordinates: ccs,
	}})
	if err != nil {
		return nil, err
	}

	rcj := reply.GetCronJob()
	var bCcs []*biz.CheckCoordinate
	for _, rCcd := range rcj.CheckCoordinates {
		bCrd := rCcd.GetCrd()
		crd := &biz.Coordinate{
			X:         bCrd.GetX(),
			Y:         bCrd.GetY(),
			Z:         bCrd.GetZ(),
			Rx:        bCrd.GetRx(),
			Ry:        bCrd.GetRy(),
			Check:     bCrd.GetCheck(),
			Delay:     bCrd.GetDelay(),
			MachineId: bCrd.GetMachineId(),
			CheckName: bCrd.GetCheckName(),
		}
		bCcs = append(bCcs, &biz.CheckCoordinate{
			Seq: rCcd.GetSeq(),
			Crd: crd,
		})
	}

	return &biz.CronJob{
		CheckCoordinates: bCcs,
		CheckName:        rcj.GetCheckName(),
		MachineId:        rcj.GetMachineId(),
		CronString:       rcj.GetCronString(),
		ID:               rcj.GetId(),
	}, nil
}

func (r *cronJobRepo) Delete(ctx context.Context, id int64) (int64, error) {
	reply, err := r.data.mc.DeleteCronJob(ctx, &machineV1.DeleteCronJobRequest{Id: id})
	if err != nil {
		return 0, err
	}

	return reply.GetNum(), nil
}

func (r *cronJobRepo) List(ctx context.Context, machineId string) ([]*biz.CronJob, error) {
	reply, err := r.data.mc.ListCronJob(ctx, &machineV1.ListCronJobRequest{MachineId: machineId})
	if err != nil {
		return nil, err
	}

	rCjs := reply.GetCronJobs()
	var bCjs []*biz.CronJob
	for _, rCj := range rCjs {
		rCods := rCj.GetCheckCoordinates()
		var bCods []*biz.CheckCoordinate
		for _, rCod := range rCods {
			rCrd := rCod.GetCrd()
			bCods = append(bCods, &biz.CheckCoordinate{
				Seq: rCod.GetSeq(),
				Crd: &biz.Coordinate{
					X:         rCrd.GetX(),
					Y:         rCrd.GetY(),
					Z:         rCrd.GetZ(),
					Rx:        rCrd.GetRx(),
					Ry:        rCrd.GetRy(),
					Check:     rCrd.GetCheck(),
					Delay:     rCrd.GetDelay(),
					MachineId: rCrd.GetMachineId(),
					CheckName: rCrd.GetCheckName(),
				},
			})
		}

		bCjs = append(bCjs, &biz.CronJob{
			CheckCoordinates: bCods,
			CheckName:        rCj.GetCheckName(),
			MachineId:        rCj.GetMachineId(),
			CronString:       rCj.GetCronString(),
			ID:               rCj.GetId(),
		})
	}

	return bCjs, err
}
