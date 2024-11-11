package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/godverv/hello_world/internal/config"
	"github.com/godverv/hello_world/internal/transport/grpc/api/docs"
	"github.com/godverv/hello_world/pkg/api"
)

type Impl struct {
	api.UnimplementedHelloWorldAPIServer

	version string

	db *sql.DB
}

func New(db *sql.DB, cfg config.Config) *Impl {
	a := &Impl{
		version: cfg.AppInfo.Version,
		db:      db,
	}

	return a
}

func (a *Impl) Register(server *grpc.Server) {
	api.RegisterHelloWorldAPIServer(server, a)
}

func (a *Impl) Gateway(ctx context.Context, endpoint string, opts ...grpc.DialOption) (route string, handler http.Handler) {
	gwHttpMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard, &runtime.JSONPb{},
		),
	)

	err := api.RegisterHelloWorldAPIHandlerFromEndpoint(ctx,
		gwHttpMux,
		endpoint,
		opts)
	if err != nil {
		logrus.Errorf("error registering grpc2http handler: %s", err)
	}

	return "/api/", gwHttpMux
}

func (a *Impl) SwaggerUI() (path string, handler http.Handler) {
	return docs.Swagger()
}
