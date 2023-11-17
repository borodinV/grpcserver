package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"grpc/server/app"
)

type Repo interface {
	AddBook(ctx context.Context, book *app.Book) (int32, error)
	GetBook(ctx context.Context, book *app.Book) (*app.Book, error)
	UpdateBook(ctx context.Context, book *app.Book) error
	DeleteBook(ctx context.Context, book *app.Book) error
	SearchBookByName(ctx context.Context, book *app.Book) ([]app.Book, error)
	GetAll(ctx context.Context) ([]app.Book, error)
}
type Repository struct {
	Repo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewPostgresDB(db)}
}
