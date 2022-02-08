//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"htp-platform/app/user/service/internal/biz"
	"htp-platform/app/user/service/internal/conf"
	"htp-platform/app/user/service/internal/data"
	"htp-platform/app/user/service/internal/server"
	"htp-platform/app/user/service/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
