package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "htp-platform/api/htpp/interface/v1"
	"htp-platform/app/htpp/interface/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewInterfaceService)

// InterfaceService is a realworld service.
type InterfaceService struct {
	v1.UnimplementedInterfaceServer

	uu  *biz.UserUsecase
	log *log.Helper
}

// NewInterfaceService new a realworld service.
func NewInterfaceService(uu *biz.UserUsecase, logger log.Logger) *InterfaceService {
	return &InterfaceService{
		uu:  uu,
		log: log.NewHelper(log.With(logger, "module", "service/server-service")),
	}
}
