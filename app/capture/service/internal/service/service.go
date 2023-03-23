package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	captureV1 "github.com/ynu-e-young/apis-go/htpp/capture/service/v1"

	"github.com/ynu-e-young/htp-platform/app/capture/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCaptureService)

// CaptureService is a user service.
type CaptureService struct {
	captureV1.UnimplementedCaptureServer

	uu  *biz.CaptureUsecase
	log *log.Helper
}

// NewCaptureService new a user service.
func NewCaptureService(uu *biz.CaptureUsecase, logger log.Logger) *CaptureService {
	return &CaptureService{
		uu:  uu,
		log: log.NewHelper(log.With(logger, "module", "service/server-service")),
	}
}
