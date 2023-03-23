package service

import (
	"context"

	interfaceV1 "github.com/ynu-e-young/apis-go/htpp/htpp/interface/v1"

	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/biz"
)

func (s *InterfaceService) CreateCronJob(ctx context.Context, in *interfaceV1.CreateCronJobRequest) (*interfaceV1.CronJobReply, error) {
	iCj := in.GetCronJob()

	var bCcs []*biz.CheckCoordinate
	for _, rCcd := range iCj.CheckCoordinates {
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

	po, err := s.cju.Create(ctx, &biz.CronJob{
		CheckCoordinates: bCcs,
		CheckName:        iCj.GetCheckName(),
		MachineId:        iCj.GetMachineId(),
		CronString:       iCj.GetCronString(),
	})
	if err != nil {
		return nil, err
	}

	var ccs []*interfaceV1.CheckCoordinate
	for _, ccd := range po.CheckCoordinates {
		mCrd := ccd.Crd
		crd := &interfaceV1.Coordinate{
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
		ccs = append(ccs, &interfaceV1.CheckCoordinate{
			Seq: ccd.Seq,
			Crd: crd,
		})
	}

	return &interfaceV1.CronJobReply{CronJob: &interfaceV1.CronJob{
		Id:               po.ID,
		MachineId:        po.MachineId,
		CheckName:        po.CheckName,
		CronString:       po.CronString,
		CheckCoordinates: ccs,
	}}, nil
}

func (s *InterfaceService) DeleteCronJob(ctx context.Context, in *interfaceV1.DeleteCronJobRequest) (*interfaceV1.DeleteCronJobReply, error) {
	num, err := s.cju.Delete(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	return &interfaceV1.DeleteCronJobReply{Num: num}, nil
}

func (s *InterfaceService) ListCronJob(ctx context.Context, in *interfaceV1.ListCronJobRequest) (*interfaceV1.CronJobsReply, error) {
	targets, err := s.cju.List(ctx, in.GetMachineId())
	if err != nil {
		return nil, err
	}

	var cjs []*interfaceV1.CronJob
	for _, target := range targets {
		var cc []*interfaceV1.CheckCoordinate
		for _, checkCoordinate := range target.CheckCoordinates {
			tCrd := checkCoordinate.Crd

			cc = append(cc, &interfaceV1.CheckCoordinate{
				Seq: checkCoordinate.Seq,
				Crd: &interfaceV1.Coordinate{
					X:         tCrd.X,
					Y:         tCrd.Y,
					Z:         tCrd.Z,
					Rx:        tCrd.Rx,
					Ry:        tCrd.Ry,
					Check:     tCrd.Check,
					Delay:     tCrd.Delay,
					MachineId: tCrd.MachineId,
					CheckName: tCrd.CheckName,
				},
			})
		}

		cjs = append(cjs, &interfaceV1.CronJob{
			Id:               target.ID,
			MachineId:        target.MachineId,
			CheckName:        target.CheckName,
			CronString:       target.CronString,
			CheckCoordinates: cc,
		})
	}

	return &interfaceV1.CronJobsReply{CronJobs: cjs}, nil
}
