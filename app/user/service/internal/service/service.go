package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "htp-platform/api/user/service/v1"
	"htp-platform/app/user/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

// UserService is a user service.
type UserService struct {
	v1.UnimplementedUserServer

	uu  *biz.UserUsecase
	log *log.Helper
}

// NewUserService new a user service.
func NewUserService(uu *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		uu:  uu,
		log: log.NewHelper(log.With(logger, "module", "service/server-service")),
	}
}
