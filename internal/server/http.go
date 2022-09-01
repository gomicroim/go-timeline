package server

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	nethttp "net/http"
	v1 "timeline-service/api/timeline/v1"
	"timeline-service/internal/conf"
	"timeline-service/internal/service"
)

//MiddlewareCors 设置跨域请求头
func MiddlewareCors() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if ts, ok := transport.FromServerContext(ctx); ok {
				if ht, ok := ts.(http.Transporter); ok {
					ht.ReplyHeader().Set("Access-Control-Allow-Origin", "*")
					ht.ReplyHeader().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,PATCH,DELETE")
					ht.ReplyHeader().Set("Access-Control-Allow-Credentials", "true")
					ht.ReplyHeader().Set("Access-Control-Allow-Headers", "Content-Type,"+
						"X-Requested-With,Access-Control-Allow-Credentials,User-Agent,Content-Length,Authorization")
				}
			}
			return handler(ctx, req)
		}
	}
}

func NotFoundHandler(res nethttp.ResponseWriter, req *nethttp.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,PATCH,DELETE")
	res.Header().Set("Access-Control-Allow-Credentials", "true")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type,"+
		"X-Requested-With,Access-Control-Allow-Credentials,User-Agent,Content-Length,Authorization")

	log.Info("NotFoundHandler")

	errCode := errors.NotFound("page not found", "page not found")
	buffer, _ := json.Marshal(errCode)
	//res.WriteHeader(400)
	_, _ = res.Write(buffer)
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, timeline *service.TimelineService, logger log.Logger) *http.Server {
	nethttp.HandleFunc("/", NotFoundHandler)

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			MiddlewareCors(),
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
	v1.RegisterTimelineHTTPServer(srv, timeline)
	return srv
}
