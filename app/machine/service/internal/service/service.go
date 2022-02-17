package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "htp-platform/api/machine/service/v1"
	"htp-platform/app/machine/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewMachineService)

type MachineService struct {
	v1.UnimplementedMachineServer

	mu  *biz.MachineUsecase
	log *log.Helper
}

func NewMachineService(mu *biz.MachineUsecase, logger log.Logger) *MachineService {
	return &MachineService{
		mu:  mu,
		log: log.NewHelper(log.With(logger, "module", "service/server-service")),
	}
}
