package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	captureV1 "htp-platform/api/capture/service/v1"
	"htp-platform/app/machine/service/internal/conf"
	"htp-platform/app/machine/service/internal/data/ent"
	"htp-platform/app/machine/service/internal/data/ent/migrate"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	consulAPI "github.com/hashicorp/consul/api"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewEntClient,
	NewMachineRepo,
	NewCaptureRepo,
	NewDiscovery,
	NewCaptureServiceClient,
)

// Data .
type Data struct {
	cc captureV1.CaptureClient

	db *ent.Client
}

// NewData .
func NewData(entClient *ent.Client, cc captureV1.CaptureClient, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(log.With(logger, "module", "machine-service/data"))

	d := &Data{
		cc: cc,
		db: entClient,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			helper.Error(err)
		}
	}, nil
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	helper := log.NewHelper(log.With(logger, "module", "user-service/data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		helper.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		helper.Fatalf("failed creating schema resources: %v", err)
	}
	return client
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
