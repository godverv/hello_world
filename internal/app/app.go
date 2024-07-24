package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/sirupsen/logrus"

	"github.com/godverv/hello_world/internal/clients/sqlite"
	"github.com/godverv/hello_world/internal/config"
	"github.com/godverv/hello_world/internal/transport"
	"github.com/godverv/hello_world/internal/transport/grpc"
	"github.com/godverv/hello_world/internal/transport/grpc/api"
	"github.com/godverv/hello_world/internal/utils/closer"
)

type App struct {
	manager *transport.ServersManager
}

func New() (App, error) {
	logrus.Println("starting app")

	cfg, err := config.Load()
	if err != nil {
		return App{}, errors.Wrap(err, "error reading config")
	}

	if cfg.GetAppInfo().StartupDuration == 0 {
		return App{}, errors.New("no startup duration in config")
	}
	app := App{}

	grpcServerCfg, err := cfg.GetServers().GRPC(config.ServerGrpc)
	if err != nil {
		return App{}, errors.Wrap(err, "error getting grpc config")
	}

	dbConf, err := cfg.GetDataSources().Sqlite(config.ResourceSqlite)
	if err != nil {
		return App{}, errors.Wrap(err, "error getting sqlite config")
	}

	storage, err := sqlite.NewStorage(dbConf)
	if err != nil {
		return App{}, errors.Wrap(err, "error creating storage")
	}

	grpcServer, err := grpc.NewServer(grpcServerCfg, api.New(storage))
	if err != nil {
		return App{}, errors.Wrap(err, "error creating grpc server")
	}

	app.manager = transport.NewManager()
	app.manager.AddServer(grpcServer)

	return app, nil
}

func (a *App) Start() error {
	ctx := context.Background()
	err := a.manager.Start(ctx)
	if err != nil {
		return errors.Wrap(err, "error starting manager")
	}

	waitingForTheEnd()

	logrus.Println("shutting down the app")
	err = closer.Close()
	if err != nil {
		logrus.Fatalf("errors while shutting down application %s", err.Error())
	}

	return nil
}

// rscli comment: an obligatory function for tool to work properly.
// must be called in the main function above
// also this is the LP's song name reference, so no linting rules can be applied to the function name
func waitingForTheEnd() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
