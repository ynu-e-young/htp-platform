package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "htp-platform/api/machine/service/v1"
	"htp-platform/app/machine/service/internal/biz"
	"htp-platform/app/machine/service/internal/conf"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewMachineService)

type MachineService struct {
	v1.UnimplementedMachineServer

	cu  *biz.CaptureUsecase
	cr  *biz.CronUsecase
	mu  *biz.MachineUsecase
	dcf *conf.Data
	log *log.Helper
}

func NewMachineService(cu *biz.CaptureUsecase, cr *biz.CronUsecase, mu *biz.MachineUsecase, dcf *conf.Data, logger log.Logger) *MachineService {
	return &MachineService{
		cu:  cu,
		cr:  cr,
		mu:  mu,
		dcf: dcf,
		log: log.NewHelper(log.With(logger, "module", "service/server-service")),
	}
}
