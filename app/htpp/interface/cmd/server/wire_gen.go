// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/biz"
	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/conf"
	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/data"
	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/server"
	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, jwt *conf.Jwt, logger log.Logger) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	userClient := data.NewUserServiceClient(discovery)
	captureClient := data.NewCaptureServiceClient(discovery)
	machineClient := data.NewMachineServiceClient(discovery)
	dataData, err := data.NewData(userClient, captureClient, machineClient, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, jwt, logger)
	captureRepo := data.NewCaptureRepo(dataData, logger)
	captureUsecase := biz.NewCaptureUsecase(captureRepo, logger)
	machineRepo := data.NewMachineRepo(dataData, logger)
	machineUsecase := biz.NewMachineUsecase(machineRepo, logger)
	cronJobRepo := data.NewCronJobRepo(dataData, logger)
	cronJobUsecase := biz.NewCronJobUsecase(cronJobRepo, logger)
	interfaceService := service.NewInterfaceService(userUsecase, captureUsecase, machineUsecase, cronJobUsecase, confData, logger)
	httpServer := server.NewHTTPServer(confServer, jwt, interfaceService, logger)
	grpcServer := server.NewGRPCServer(confServer, interfaceService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
	}, nil
}
