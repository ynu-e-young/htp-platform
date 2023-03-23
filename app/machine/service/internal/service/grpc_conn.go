package service

import (
	robotV1 "github.com/ynu-e-young/apis-go/htpp/machine/robot/v1"
	"google.golang.org/grpc"
)

// 获取当前设备 (uuid) 的 rpc client
func (s *MachineService) robotClient(address string) (robotV1.RobotClient, func(), error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		s.log.Error(err)
		return nil, nil, err
	}

	return robotV1.NewRobotClient(conn), func() {
		if err := conn.Close(); err != nil {
			s.log.Error(err)
		}
	}, nil
}
