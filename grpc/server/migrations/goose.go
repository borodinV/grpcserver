package migrations

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/spf13/viper"
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
func InitMigrations(dir string, command string) error {

	cfg := repo.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	}

	err := Run(dir, command, cfg)
	if err != nil {
		return err
	}

	return nil
}
