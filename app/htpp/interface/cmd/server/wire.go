//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/biz"
	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/conf"
	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/data"
	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/server"
	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Jwt, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
