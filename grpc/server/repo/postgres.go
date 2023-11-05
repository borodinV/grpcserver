package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Port     string
	Username string
	DBName   string
	SSLMode  string
	Password string
}

const (
	LibraryTable = "library"
)

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres",
		fmt.Sprintf("port=%s user=%s dbname=%s sslmode=%s password=%s",
			cfg.Port, cfg.Username, cfg.DBName, cfg.SSLMode, cfg.Password))

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
