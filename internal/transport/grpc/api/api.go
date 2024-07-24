package api

import (
	"context"
	"database/sql"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/godverv/hello_world/pkg/api"
)

type Api struct {
	api.UnimplementedHelloWorldAPIServer

	version string

	db *sql.DB
}

func New(db *sql.DB) *Api {
	a := &Api{
		version: "v0.0.1",
		db:      db,
	}

	return a
}

func (a *Api) Register(server grpc.ServiceRegistrar) {
	api.RegisterHelloWorldAPIServer(server, a)
}

func (a *Api) RegisterGw(ctx context.Context, mux *runtime.ServeMux, addr string) error {
	return api.RegisterHelloWorldAPIHandlerFromEndpoint(
		ctx,
		mux,
		addr,
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		})
}
