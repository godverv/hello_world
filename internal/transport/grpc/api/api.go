package api

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/godverv/hello_world/pkg/api"
)

type Api struct {
	api.UnimplementedHelloWorldAPIServer

	version string
}

func New() *Api {
	a := &Api{}

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
