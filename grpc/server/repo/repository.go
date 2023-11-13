package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"grpc/server/app"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}
func (r *Repository) AddBook(ctx context.Context, book *app.Book) (int32, error) {

	var result int32

	row := r.db.QueryRow("insert into library (name, author, year) values ($1, $2, $3) returning id",
		book.Name, book.Author, book.Year)

	if err := row.Scan(&result); err != nil {
		return 0, err
	}

	return result, nil

}
func (r *Repository) GetBook(ctx context.Context, book *app.Book) (*app.Book, error) {

	var result app.Book

	err := r.db.Get(&result, "select * from library where id = $1", book.Id)
	if err != nil {
		return nil, err
	}

	return &result, nil

}
func (r *Repository) UpdateBook(ctx context.Context, book *app.Book) (string, error) {

	response := "OK"

	_, err := r.db.Exec("update library set name = $1, author = $2, year = $3 where id = $4",
		book.Name, book.Author, book.Year, book.Id)
	if err != nil {
		return "", err
	}

	return response, err

}
func (r *Repository) DeleteBook(ctx context.Context, book *app.Book) (string, error) {

	response := "Book Deleted!"

	_, err := r.db.Exec("delete from library where id = $1", book.Id)
	if err != nil {
		return "", err
	}

	return response, nil
}
func (r *Repository) SearchBookByName(ctx context.Context, book *app.Book) ([]*app.Book, error) {

	resultSlice := make([]*app.Book, 0, 10)

	rows, err := r.db.Query("select * from library where name=$1", book.Name)
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

		resultSlice = append(resultSlice, &result)
	}

	return resultSlice, nil
}
func (r *Repository) GetAll(ctx context.Context, in string) ([]*app.Book, error) {

	resultSlice := make([]*app.Book, 0, 10)

	rows, err := r.db.Query("select * from library")
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

		resultSlice = append(resultSlice, &book)
	}

	return resultSlice, nil
}
