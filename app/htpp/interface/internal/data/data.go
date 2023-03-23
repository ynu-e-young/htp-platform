package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	captureV1 "github.com/ynu-e-young/apis-go/htpp/capture/service/v1"
	machineV1 "github.com/ynu-e-young/apis-go/htpp/machine/service/v1"
	userV1 "github.com/ynu-e-young/apis-go/htpp/user/service/v1"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	consulAPI "github.com/hashicorp/consul/api"

	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	NewCaptureRepo,
	NewMachineRepo,
	NewCronJobRepo,
	NewDiscovery,
	NewUserServiceClient,
	NewCaptureServiceClient,
	NewMachineServiceClient,
)

// Data .
type Data struct {
	uc userV1.UserClient
	cc captureV1.CaptureClient
	mc machineV1.MachineClient

	helper *log.Helper
}

// NewData .
func NewData(uc userV1.UserClient, cc captureV1.CaptureClient, mc machineV1.MachineClient, logger log.Logger) (*Data, error) {
	helper := log.NewHelper(log.With(logger, "module", "htpp-interface/data"))

	return &Data{
		uc:     uc,
		cc:     cc,
		mc:     mc,
		helper: helper,
	}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewUserServiceClient(r registry.Discovery) userV1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///htp-platform.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := userV1.NewUserClient(conn)
	return c
}

func NewCaptureServiceClient(r registry.Discovery) captureV1.CaptureClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///htp-platform.capture.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := captureV1.NewCaptureClient(conn)
	return c
}

func NewMachineServiceClient(r registry.Discovery) machineV1.MachineClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///htp-platform.machine.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := machineV1.NewMachineClient(conn)
	return c
}
