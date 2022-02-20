package service

import (
	"context"
	robotV1 "htp-platform/api/machine/robot/v1"
	machineV1 "htp-platform/api/machine/service/v1"
	"htp-platform/app/machine/service/internal/biz"
	"os"
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

func (s *MachineService) saveImage(data []byte) error {
	if err := os.WriteFile(s.dcf.Images.Dir+"/"+time.Now().String()+".jpg", data, 0644); err != nil {
		return err
	}
	return nil
}

func (s *MachineService) ReadAllWithBinaryAndCalArea(ctx context.Context) {
	captures, err := s.cu.ReadAll(ctx)
	if err != nil {
		s.log.Error(err)
		return
	}

	for _, capture := range captures {
		if err = s.saveImage(capture.Data); err != nil {
			s.log.Error(err)
		}
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
		go s.ReadAllWithBinaryAndCalArea(ctx)
	}

	ret.Status = true
	return ret, nil
}
