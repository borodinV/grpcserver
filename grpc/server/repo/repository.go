package repo

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"grpc/server/proto"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}
func (r *Repository) AddBook(ctx context.Context, book *proto.Book) (*proto.BookID, error) {

	var bookIdInt int32

	row := r.db.QueryRow("insert into library (name, author, year) values ($1, $2, $3) returning id",
		book.Name, book.Author, book.Year)

	if err := row.Scan(&bookIdInt); err != nil {
		return nil, err
	}

	var bookId = proto.BookID{Id: bookIdInt}

	return &bookId, nil

}
func (r *Repository) GetBook(ctx context.Context, bookId *proto.BookID) (*proto.Book, error) {

	var resultBook proto.Book

	err := r.db.Get(&resultBook, "select * from library where id = $1", bookId.Id)
	if err != nil {
		return nil, err
	}

	return &resultBook, nil

}
func (r *Repository) UpdateBook(ctx context.Context, book *proto.Book) (*wrappers.StringValue, error) {

	var response = wrappers.StringValue{Value: "OK"}

	_, err := r.db.Exec("update library set name = $1, author = $2, year = $3 where id = $4",
		book.Name, book.Author, book.Year, book.Id)
	if err != nil {
		return nil, err
	}

	return &response, err

}
func (r *Repository) DeleteBook(ctx context.Context, bookId *proto.BookID) (*wrappers.StringValue, error) {

	var response = wrappers.StringValue{Value: "Book Deleted!"}

	_, err := r.db.Exec("delete from library where id = $1", bookId.Id)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
func (r *Repository) SearchBookByName(ctx context.Context, book *proto.BookName) (*proto.BookList, error) {

	var resultBookList proto.BookList

	rows, err := r.db.Query("select * from library where name=$1", book.Name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var book proto.Book

		err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}

		resultBookList.Books = append(resultBookList.Books, &book)
	}

	return &resultBookList, nil
}
func (r *Repository) GetAll(ctx context.Context, in *wrapperspb.StringValue) (*proto.BookList, error) {

	var resultBookList proto.BookList

	rows, err := r.db.Query("select * from library")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var book proto.Book

		err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}

		resultBookList.Books = append(resultBookList.Books, &book)
	}

	return &resultBookList, nil
}
