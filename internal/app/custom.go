// Code generated by RedSock CLI.
// DO EDIT, DON'T DELETE THIS FILE.

package app

import (
	"context"

	"github.com/godverv/hello_world/internal/transport"
	"github.com/godverv/hello_world/internal/transport/docs"
	"github.com/godverv/hello_world/internal/transport/grpc/api"
)

type Custom struct {
	ServerManager *transport.ServersManager
}

func (c *Custom) Init(a *App) error {
	grpcImpl := api.New(a.Sqlite, a.Cfg)

	a.ServerMaster.AddImplementation(grpcImpl)

	a.ServerMaster.AddHttpHandler(docs.Swagger())

	return nil
}

func (c *Custom) Start(ctx context.Context) error {
	return nil
}

func (c *Custom) Stop() error {
	return nil
}
