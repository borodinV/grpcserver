package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/jmoiron/sqlx"
	"grpc/server/proto"
	"grpc/server/repo"
	"strconv"
)

type Server struct {
	proto.UnimplementedLibraryServer
	db *sqlx.DB
}

func NewServer(db *sqlx.DB) *Server {
	return &Server{db: db}
}

func (s *Server) AddBook(ctx context.Context, book *proto.Book) (*proto.BookID, error) {

	var bookIdInt int

	queryString := fmt.Sprintf("insert into %s (name, author, year) values ($1, $2, $3) returning id", repo.LibraryTable)
	row := s.db.QueryRow(queryString, book.Name, book.Author, book.Year)

	if err := row.Scan(&bookIdInt); err != nil {
		return nil, err
	}

	var bookId = proto.BookID{Id: strconv.Itoa(bookIdInt)}

	return &bookId, nil
}
func (s *Server) GetBook(ctx context.Context, bookId *proto.BookID) (*proto.Book, error) {

	var book proto.Book

	queryString := fmt.Sprintf("select * from %s where id = $1", repo.LibraryTable)
	err := s.db.Get(&book, queryString, bookId.Id)
	if err != nil {
		return nil, err
	}

	return &book, nil

}
func (s *Server) UpdateBook(ctx context.Context, book *proto.Book) (*wrappers.StringValue, error) {

	var response = wrappers.StringValue{Value: "OK"}

	queryString := fmt.Sprintf("update %s set name = $1, author = $2, year = $3 where id = $4", repo.LibraryTable)
	_, err := s.db.Exec(queryString, book.Name, book.Author, book.Year, book.Id)
	if err != nil {
		return nil, err
	}
	return &response, err
}
func (s *Server) DeleteBook(ctx context.Context, bookId *proto.BookID) (*wrappers.StringValue, error) {

	var response = wrappers.StringValue{Value: "Book Deleted!"}

	queryString := fmt.Sprintf("delete from %s where id = $1", repo.LibraryTable)
	_, err := s.db.Exec(queryString, bookId.Id)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (s *Server) SearchBookByName(book *proto.BookName, stream proto.Library_SearchBookByNameServer) error {

	queryString := fmt.Sprintf("select * from %s where name=$1", repo.LibraryTable)
	rows, err := s.db.Query(queryString, book.Name)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {

		var books proto.Book

		err := rows.Scan(&books.Id, &books.Name, &books.Author, &books.Year)
		if err != nil {
			return err
		}
		err = stream.Send(&books)
		if err != nil {
			return err
		}
	}
	return nil
}
