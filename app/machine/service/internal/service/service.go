package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/ynu-e-young/apis-go/htpp/machine/service/v1"

	"github.com/ynu-e-young/htp-platform/app/machine/service/internal/biz"
	"github.com/ynu-e-young/htp-platform/app/machine/service/internal/conf"
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
