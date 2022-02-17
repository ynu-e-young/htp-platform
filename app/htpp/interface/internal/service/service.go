package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "htp-platform/api/htpp/interface/v1"
	"htp-platform/app/htpp/interface/internal/biz"
	"htp-platform/app/htpp/interface/internal/conf"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewInterfaceService)

// InterfaceService is an interface service.
type InterfaceService struct {
	v1.UnimplementedInterfaceServer

	uu  *biz.UserUsecase
	cu  *biz.CaptureUsecase
	mu  *biz.MachineUsecase
	dcf *conf.Data
	log *log.Helper
}

// NewInterfaceService new an interface service.
func NewInterfaceService(uu *biz.UserUsecase, cu *biz.CaptureUsecase, mu *biz.MachineUsecase, conf *conf.Data, logger log.Logger) *InterfaceService {
	return &InterfaceService{
		uu:  uu,
		cu:  cu,
		mu:  mu,
		dcf: conf,
		log: log.NewHelper(log.With(logger, "module", "service/server-service")),
	}
}
