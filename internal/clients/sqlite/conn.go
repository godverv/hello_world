package sqlite

import (
	"database/sql"
	"os"
	"path"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka/resources"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

const (
	dialect = "sqlite"

	inMemory = "file::memory:?mode=memory&cache=shared"
)

func NewStorage(cfg *resources.Sqlite) (*sql.DB, error) {
	databaseLocation := cfg.Path
	if databaseLocation == "" {
		logrus.Warning("no path for file, running in memory mode")
		databaseLocation = inMemory
	} else {
		err := os.MkdirAll(path.Dir(cfg.Path), 0777)
		if err != nil {
			return nil, errors.Wrap(err, "error creating dir for sqlite db")
		}

	}

	conn, err := sql.Open(dialect, databaseLocation)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database connection")
	}

	err = checkMigrate(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error performing migrations")
	}

	return conn, nil
}

func checkMigrate(conn *sql.DB) error {
	goose.SetLogger(logrus.StandardLogger())
	err := goose.SetDialect(dialect)
	if err != nil {
		return errors.Wrap(err, "error setting dialect")
	}

	err = goose.Up(conn, "./migrations")
	if err != nil {
		return errors.Wrap(err, "error performing up")
	}

	return nil
}
