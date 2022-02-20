package pkg

import (
	"google.golang.org/grpc"
	robotV1 "htp-platform/api/machine/robot/v1"
	"log"
)

// 获取当前设备 (uuid) 的 rpc client
func machineClient(address string) (robotV1.RobotClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return NewControlServiceClient(conn), nil
}
