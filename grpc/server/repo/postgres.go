package repo

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"grpc/server/app"
)

type Config struct {
	Host     string
	Port     string
	Username string
	DBName   string
	SSLMode  string
	Password string
}

type PostgresDB struct {
	db *sqlx.DB
}

func NewPostgresDB(db *sqlx.DB) *PostgresDB {
	return &PostgresDB{db: db}
}
func InitPostgresDB() (*sqlx.DB, error) {

	cfg := &Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	}

	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.SSLMode, cfg.Password))

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
func (r *PostgresDB) AddBook(ctx context.Context, book *app.Book) (int32, error) {

	var result int32

	row := r.db.QueryRow("insert into library (name, author, year) values ($1, $2, $3) returning id",
		book.Name, book.Author, book.Year)

	if err := row.Scan(&result); err != nil {
		return 0, err
	}

	return result, nil

}
func (r *PostgresDB) GetBook(ctx context.Context, book *app.Book) (*app.Book, error) {

	var result app.Book

	err := r.db.Get(&result, "select id, name, author, year from library where id = $1", book.Id)
	if err != nil {
		return nil, err
	}

	return &result, nil

}
func (r *PostgresDB) UpdateBook(ctx context.Context, book *app.Book) error {

	_, err := r.db.Exec("update library set name = $1, author = $2, year = $3 where id = $4",
		book.Name, book.Author, book.Year, book.Id)
	if err != nil {
		return err
	}

	return err

}
func (r *PostgresDB) DeleteBook(ctx context.Context, book *app.Book) error {

	_, err := r.db.Exec("delete from library where id = $1", book.Id)
	if err != nil {
		return err
	}

	return nil
}
func (r *PostgresDB) SearchBookByName(ctx context.Context, book *app.Book) ([]app.Book, error) {

	resultSlice := make([]app.Book, 0, 10)

	rows, err := r.db.Query("select id, name, author, year from library where name=$1", book.Name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var result app.Book

		err := rows.Scan(&result.Id, &result.Name, &result.Author, &result.Year)
		if err != nil {
			return nil, err
		}

		resultSlice = append(resultSlice, result)
	}

	return resultSlice, nil
}
func (r *PostgresDB) GetAll(ctx context.Context) ([]app.Book, error) {

	resultSlice := make([]app.Book, 0, 10)

	rows, err := r.db.Query("select id, name, author, year from library")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var book app.Book

		err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}

		resultSlice = append(resultSlice, book)
	}

	return resultSlice, nil
}
