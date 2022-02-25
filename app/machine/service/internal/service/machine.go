package service

import (
	"context"
	robotV1 "htp-platform/api/machine/robot/v1"
	machineV1 "htp-platform/api/machine/service/v1"
	"htp-platform/app/machine/service/internal/biz"
	"os"
	"sort"
	"strconv"
	"time"
)

func (s *MachineService) FindByUserId(ctx context.Context, in *machineV1.FindByUserIdRequest) (*machineV1.MachinesReply, error) {
	ms, err := s.mu.FindByUserId(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}

	var ret []*machineV1.MachineStruct
	for _, m := range ms {
		ret = append(ret, &machineV1.MachineStruct{
			MachineId: m.MachineId,
			UserId:    m.UserId,
			Address:   m.Address,
		})
	}

	return &machineV1.MachinesReply{
		Machines: ret,
	}, nil
}

func (s *MachineService) Create(ctx context.Context, in *machineV1.CreateRequest) (*machineV1.MachineReply, error) {
	machine := in.GetMachine()
	m, err := s.mu.Create(ctx, &biz.Machine{
		UserId:  machine.UserId,
		Address: machine.Address,
	})
	if err != nil {
		return nil, err
	}

	return &machineV1.MachineReply{Machine: &machineV1.MachineStruct{
		MachineId: m.MachineId,
		UserId:    m.UserId,
		Address:   m.Address,
	}}, nil
}

func (s *MachineService) Update(ctx context.Context, in *machineV1.UpdateRequest) (*machineV1.MachineReply, error) {
	machine := in.GetMachine()
	m, err := s.mu.Update(ctx, &biz.Machine{
		MachineId: machine.MachineId,
		UserId:    machine.UserId,
		Address:   machine.Address,
	})
	if err != nil {
		return nil, err
	}

	return &machineV1.MachineReply{Machine: &machineV1.MachineStruct{
		MachineId: m.MachineId,
		UserId:    m.UserId,
		Address:   m.Address,
	}}, nil
}

func (s *MachineService) Get(ctx context.Context, in *machineV1.GetRequest) (*machineV1.MachineReply, error) {
	m, err := s.mu.Get(ctx, in.GetMachineId())
	if err != nil {
		return nil, err
	}

	return &machineV1.MachineReply{Machine: &machineV1.MachineStruct{
		MachineId: m.MachineId,
		UserId:    m.UserId,
		Address:   m.Address,
	}}, nil
}

func (s *MachineService) saveImage(fileName string, data []byte) error {
	if err := os.WriteFile(fileName, data, 0644); err != nil {
		return err
	}
	return nil
}

func (s *MachineService) ReadAllWithBinaryAndCalArea(ctx context.Context, machineIdStr string) {
	captures, err := s.cu.ReadAllWithBinaryAndCalArea(ctx)
	if err != nil {
		s.log.Error(err)
		return
	}

	for _, capture := range captures {
		fileName := s.dcf.Images.Dir + "/" + time.Now().String() + ".jpg"
		machineId, err := strconv.ParseInt(machineIdStr, 10, 64)
		if err != nil {
			s.log.Error(err)
		}

		if _, err := s.cu.CreateLog(ctx, &biz.CaptureLog{
			MachineId: machineId,
			Pixels:    capture.Pixels,
			Area:      capture.Area,
			ImageName: fileName,
			OssUrl:    "",
		}); err != nil {
			s.log.Error(err)
		}

		if err = s.saveImage(fileName, capture.Data); err != nil {
			s.log.Error(err)
		}

		//go s.OssUpload(fileName, capture.Data)
	}
}

func (s *MachineService) Move(ctx context.Context, in *machineV1.MoveRequest) (*machineV1.MoveReply, error) {
	m, err := s.mu.Get(ctx, in.GetMachineId())
	if err != nil {
		return nil, err
	}

	rCli, cleanup, err := s.robotClient(m.Address)
	if err != nil {
		return nil, err
	}
	defer cleanup()

	reply, err := rCli.AppendCoordinate(ctx, &robotV1.CoordinateRequest{
		X:         in.GetX(),
		Y:         in.GetY(),
		Z:         in.GetZ(),
		Rx:        in.GetRx(),
		Ry:        in.GetRy(),
		Check:     in.GetCheck(),
		Delay:     in.GetDelay(),
		Uuid:      strconv.FormatInt(in.GetMachineId(), 10),
		CheckName: in.GetCheckName(),
	})
	if err != nil {
		return nil, err
	}

	return &machineV1.MoveReply{Status: reply.GetStatus()}, nil
}

func (s *MachineService) Zero(ctx context.Context, in *machineV1.ZeroRequest) (*machineV1.ZeroReply, error) {
	m, err := s.mu.Get(ctx, in.GetMachineId())
	if err != nil {
		return nil, err
	}

	rCli, cleanup, err := s.robotClient(m.Address)
	if err != nil {
		return nil, err
	}
	defer cleanup()

	reply, err := rCli.Zero(ctx, &robotV1.ZeroRequest{Zero: true})
	if err != nil {
		return nil, err
	}

	return &machineV1.ZeroReply{Status: reply.GetStatus()}, nil
}

func (s *MachineService) GetMotorStatus(ctx context.Context, in *machineV1.GetMotorStatusRequest) (*machineV1.GetMotorStatusReply, error) {
	m, err := s.mu.Get(ctx, in.GetMachineId())
	if err != nil {
		return nil, err
	}

	rCli, cleanup, err := s.robotClient(m.Address)
	if err != nil {
		return nil, err
	}
	defer cleanup()

	reply, err := rCli.GetMotorStatus(ctx, &robotV1.MotorInfoRequest{Status: true})
	if err != nil {
		return nil, err
	}

	var infos []*machineV1.MotorInfo
	for _, info := range reply.GetMotorInfo() {
		mStatus := info.GetMotorStatus()

		infos = append(infos, &machineV1.MotorInfo{
			MotorStatus: &machineV1.MotorStatus{
				Fault:                 mStatus.GetFault(),
				Enabling:              mStatus.GetEnabling(),
				Running:               mStatus.GetRunning(),
				InstructionCompletion: mStatus.GetInstructionCompletion(),
				PathCompletion:        mStatus.GetPathCompletion(),
				ZeroCompletion:        mStatus.GetZeroCompletion(),
			},
			InstrPos:   info.GetInstrPos(),
			CurrentPos: info.GetCurrentPos(),
		})
	}

	return &machineV1.GetMotorStatusReply{MotorInfo: infos}, nil
}

func (s *MachineService) MoveDone(ctx context.Context, in *machineV1.MoveDoneRequest) (*machineV1.MoveDoneReply, error) {
	ret := &machineV1.MoveDoneReply{Status: false}

	//// 先检查设备是否存在
	//exists, err := rdOps.IsDeviceExists(moveDoneInfo.Uuid)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// 不存在就返回
	//if !exists {
	//	return ret, nil
	//}
	//
	//// 存在的话就将记录序列化
	//bytes, err := json.Marshal(moveDoneInfo)
	//if err != nil {
	//	return nil, err
	//}
	//
	////将序列化的记录写入 Redis
	//_, err = rdOps.AddMoveRecord(moveDoneInfo.Uuid, string(bytes))
	//if err != nil {
	//	return ret, err
	//}

	if in.GetCheck() {
		s.ReadAllWithBinaryAndCalArea(ctx, in.GetUuid())
	}

	ret.Status = true
	return ret, nil
}

func (s *MachineService) jobFunc(machineId int64, checkName string, coordinates []*biz.CheckCoordinate) func() {
	return func() {
		// 将坐标 Slice 以结构体中的 Seq 进行降序排序
		sort.Slice(coordinates, func(i, j int) bool {
			return coordinates[i].Seq < coordinates[j].Seq
		})

		// 通过 grpc 服务将坐标发送到客户端
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		m, err := s.mu.Get(ctx, machineId)
		if err != nil {
			s.log.Error(err)
			return
		}

		rCli, cleanup, err := s.robotClient(m.Address)
		if err != nil {
			s.log.Error(err)
			return
		}
		defer cleanup()

		// 遍历坐标，发送到机器
		for _, coordinate := range coordinates {
			coordinate.Crd.CheckName = checkName

			_, err := rCli.AppendCoordinate(ctx, &robotV1.CoordinateRequest{
				X:         coordinate.Crd.X,
				Y:         coordinate.Crd.Y,
				Z:         coordinate.Crd.Z,
				Rx:        coordinate.Crd.Rx,
				Ry:        coordinate.Crd.Ry,
				Check:     coordinate.Crd.Check,
				Delay:     coordinate.Crd.Delay,
				Uuid:      strconv.FormatInt(coordinate.Crd.MachineId, 10),
				CheckName: coordinate.Crd.CheckName,
			})
			if err != nil {
				s.log.Error(err)
				return
			}
		}
	}
}

func (s *MachineService) CreateCronJob(ctx context.Context, in *machineV1.CreateCronJobRequest) (*machineV1.CronJobReply, error) {
	c := in.GetCronJob()

	var bCcs []*biz.CheckCoordinate
	for _, coordinate := range c.GetCheckCoordinates() {
		crd := coordinate.GetCrd()
		bCrd := &biz.Coordinate{
			X:         crd.GetX(),
			Y:         crd.GetY(),
			Z:         crd.GetZ(),
			Rx:        crd.GetRx(),
			Ry:        crd.GetRy(),
			Check:     crd.GetCheck(),
			Delay:     crd.GetDelay(),
			MachineId: crd.GetMachineId(),
			CheckName: crd.GetCheckName(),
		}

		bCcs = append(bCcs, &biz.CheckCoordinate{
			Seq: coordinate.GetSeq(),
			Crd: bCrd,
		})
	}

	po, err := s.cr.Create(ctx, &biz.Cron{
		CheckCoordinates: bCcs,
		CheckName:        c.GetCheckName(),
		MachineId:        c.GetMachineId(),
		CronString:       c.GetCronString(),
	})
	if err != nil {
		return nil, err
	}

	if err := s.cr.AddCronJob(ctx, &biz.Cron{
		CheckCoordinates: bCcs,
		CheckName:        c.GetCheckName(),
		MachineId:        c.GetMachineId(),
		CronString:       c.GetCronString(),
	}, s.jobFunc(c.GetMachineId(), c.GetCheckName(), bCcs)); err != nil {
		return nil, err
	}

	cj := in.GetCronJob()
	cj.Id = po.ID
	for _, cCrd := range cj.CheckCoordinates {
		cCrd.GetCrd().CheckName = c.GetCheckName()
	}

	return &machineV1.CronJobReply{CronJob: cj}, nil
}

func (s *MachineService) DeleteCronJob(ctx context.Context, in *machineV1.DeleteCronJobRequest) (*machineV1.DeleteCronJobReply, error) {
	po, err := s.cr.Get(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	if err := s.cr.Delete(ctx, in.GetId()); err != nil {
		return nil, err
	}

	if err := s.cr.DelCronJob(ctx, po.MachineId, po.CheckName); err != nil {
		return nil, err
	}

	return &machineV1.DeleteCronJobReply{Num: 1}, nil
}

func (s *MachineService) ListCronJob(ctx context.Context, in *machineV1.ListCronJobRequest) (*machineV1.CronJobsReply, error) {
	targets, err := s.cr.FindByMachineId(ctx, in.GetMachineId())
	if err != nil {
		return nil, err
	}

	var cjs []*machineV1.CronJob
	for _, target := range targets {
		var cc []*machineV1.CheckCoordinate
		for _, coordinate := range target.CheckCoordinates {
			crd := coordinate.Crd
			cc = append(cc, &machineV1.CheckCoordinate{
				Seq: coordinate.Seq,
				Crd: &machineV1.Coordinate{
					X:         crd.X,
					Y:         crd.Y,
					Z:         crd.Z,
					Rx:        crd.Rx,
					Ry:        crd.Ry,
					Check:     crd.Check,
					Delay:     crd.Delay,
					MachineId: crd.MachineId,
					CheckName: crd.CheckName,
				},
			})
		}
		cjs = append(cjs, &machineV1.CronJob{
			Id:               target.ID,
			MachineId:        target.MachineId,
			CheckName:        target.CheckName,
			CronString:       target.CronString,
			CheckCoordinates: cc,
		})
	}

	return &machineV1.CronJobsReply{CronJobs: cjs}, nil
}
