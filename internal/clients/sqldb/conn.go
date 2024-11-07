package sqldb

import (
	"database/sql"

	"github.com/Red-Sock/toolbox/closer"
	"github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka/resources"

	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

func New(cfg resources.SqlResource) (*sql.DB, error) {
	dialect := cfg.SqlDialect()
	connStr := cfg.ConnectionString()

	conn, err := sql.Open(dialect, connStr)
	if err != nil {
		return nil, errors.Wrap(err, "error checking connection to postgres")
	}

	closer.Add(func() error {
		return conn.Close()
	})

	goose.SetLogger(logrus.StandardLogger())
	err = goose.SetDialect(dialect)
	if err != nil {
		return nil, errors.Wrap(err, "error setting dialect")
	}

	mig := cfg.MigrationFolder()
	if mig == "" {
		mig = "./migrations"
	}

	err = goose.Up(conn, mig)
	if err != nil {
		return nil, errors.Wrap(err, "error performing up")
	}

	return conn, nil
}
