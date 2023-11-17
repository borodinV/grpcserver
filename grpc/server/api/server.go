package api

import (
	"context"
	"grpc/server/app"
	"grpc/server/proto"
)

type Repo interface {
	AddBook(ctx context.Context, book *app.Book) (int32, error)
	GetBook(ctx context.Context, book *app.Book) (*app.Book, error)
	UpdateBook(ctx context.Context, book *app.Book) error
	DeleteBook(ctx context.Context, book *app.Book) error
	SearchBookByName(ctx context.Context, book *app.Book) ([]app.Book, error)
	GetAll(ctx context.Context) ([]app.Book, error)
}

type Server struct {
	proto.UnimplementedLibraryServer
	repo Repo
}

func NewServer(repo Repo) *Server {
	return &Server{repo: repo}
}
func (s *Server) AddBook(ctx context.Context, book *proto.Book) (*proto.BookID, error) {

	var input = &app.Book{
		Id:     book.Id,
		Name:   book.Name,
		Author: book.Author,
		Year:   book.Year,
	}

	id, err := s.repo.AddBook(ctx, input)
	if err != nil {
		return nil, err
	}

	return &proto.BookID{Id: id}, nil
}
func (s *Server) GetBook(ctx context.Context, bookId *proto.BookID) (*proto.Book, error) {

	var input = &app.Book{Id: bookId.Id}

	book, err := s.repo.GetBook(ctx, input)
	if err != nil {
		return nil, err
	}

	return &proto.Book{
		Id:     book.Id,
		Name:   book.Name,
		Author: book.Author,
		Year:   book.Year,
	}, nil
}
func (s *Server) UpdateBook(ctx context.Context, book *proto.Book) (*proto.Empty, error) {

	var input = &app.Book{
		Id:     book.Id,
		Name:   book.Name,
		Author: book.Author,
		Year:   book.Year,
	}

	err := s.repo.UpdateBook(ctx, input)
	if err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}
func (s *Server) DeleteBook(ctx context.Context, bookId *proto.BookID) (*proto.Empty, error) {

	var input = &app.Book{Id: bookId.Id}

	err := s.repo.DeleteBook(ctx, input)
	if err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}
func (s *Server) SearchBookByName(ctx context.Context, book *proto.BookName) (*proto.BookList, error) {

	var input = &app.Book{Name: book.Name}

	books, err := s.repo.SearchBookByName(ctx, input)
	if err != nil {
		return nil, err
	}

	resultSlice := make([]*proto.Book, 0, 10)

	for _, value := range books {

		var resultBook = &proto.Book{
			Id:     value.Id,
			Name:   value.Name,
			Author: value.Author,
			Year:   value.Year,
		}

		resultSlice = append(resultSlice, resultBook)
	}

	return &proto.BookList{Books: resultSlice}, nil
}
func (s *Server) GetAll(ctx context.Context, in *proto.Empty) (*proto.BookList, error) {

	books, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resultSlice := make([]*proto.Book, 0, 10)

	for _, value := range books {

		var resultBook = &proto.Book{
			Id:     value.Id,
			Name:   value.Name,
			Author: value.Author,
			Year:   value.Year,
		}

		resultSlice = append(resultSlice, resultBook)
	}

	return &proto.BookList{Books: resultSlice}, nil
}
