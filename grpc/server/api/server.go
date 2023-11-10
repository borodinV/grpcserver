package api

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"grpc/server/app"
	"grpc/server/proto"
	"grpc/server/repo"
)

type Server struct {
	proto.UnimplementedLibraryServer
	repo *repo.Repository
}

func NewServer(repo *repo.Repository) *Server {
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
func (s *Server) UpdateBook(ctx context.Context, book *proto.Book) (*wrappers.StringValue, error) {

	var input = &app.Book{
		Id:     book.Id,
		Name:   book.Name,
		Author: book.Author,
		Year:   book.Year,
	}

	response, err := s.repo.UpdateBook(ctx, input)
	if err != nil {
		return nil, err
	}

	return &wrappers.StringValue{Value: response}, nil
}
func (s *Server) DeleteBook(ctx context.Context, bookId *proto.BookID) (*wrappers.StringValue, error) {

	var input = &app.Book{Id: bookId.Id}

	response, err := s.repo.DeleteBook(ctx, input)
	if err != nil {
		return nil, err
	}

	return &wrappers.StringValue{Value: response}, nil
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
func (s *Server) GetAll(ctx context.Context, in *wrapperspb.StringValue) (*proto.BookList, error) {

	books, err := s.repo.GetAll(ctx, "")
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
