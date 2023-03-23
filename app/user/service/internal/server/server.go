package server

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	consulAPI "github.com/hashicorp/consul/api"

	"github.com/ynu-e-young/htp-platform/app/user/service/internal/conf"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
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
