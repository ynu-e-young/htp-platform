package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	interfaceV1 "github.com/ynu-e-young/apis-go/htpp/htpp/interface/v1"

	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/biz"
	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/conf"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewInterfaceService)

// InterfaceService is an interface service.
type InterfaceService struct {
	interfaceV1.UnimplementedInterfaceServer

	uu  *biz.UserUsecase
	cu  *biz.CaptureUsecase
	mu  *biz.MachineUsecase
	cju *biz.CronJobUsecase
	dcf *conf.Data
	log *log.Helper
}

// NewInterfaceService new an interface service.
func NewInterfaceService(
	uu *biz.UserUsecase,
	cu *biz.CaptureUsecase,
	mu *biz.MachineUsecase,
	cju *biz.CronJobUsecase,
	conf *conf.Data,
	logger log.Logger) *InterfaceService {
	return &InterfaceService{
		uu:  uu,
		cu:  cu,
		mu:  mu,
		cju: cju,
		dcf: conf,
		log: log.NewHelper(log.With(logger, "module", "service/server-service")),
	}
}
