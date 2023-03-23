package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	interfaceV1 "github.com/ynu-e-young/apis-go/htpp/htpp/interface/v1"

	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/conf"
	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/pkg/middleware/auth"
	"github.com/ynu-e-young/htp-platform/app/htpp/interface/internal/service"
)

func NewSkipRoutersMatcher() selector.MatchFunc {
	skipList := make(map[string]struct{})
	skipList["/htpp.interface.v1.Interface/Login"] = struct{}{}
	skipList["/htpp.interface.v1.Interface/Register"] = struct{}{}

	return func(ctx context.Context, operation string) bool {
		if _, ok := skipList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, jc *conf.Jwt, is *service.InterfaceService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			selector.Server(
				auth.JwtAuth(jc.GetSecret()),
			).Match(NewSkipRoutersMatcher()).Build(),
		),
		http.Filter(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	interfaceV1.RegisterInterfaceHTTPServer(srv, is)
	return srv
}
