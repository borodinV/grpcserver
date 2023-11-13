package migrations

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"grpc/server/repo"
)

func Run(dir string, command string, cfg repo.Config) error {

	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.SSLMode, cfg.Password))

	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	if err = goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err = goose.Run(command, db, dir); err != nil {
		return err
	}

	return db.Close()
}
