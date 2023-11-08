package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/wrapperspb"
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
	return s.repo.AddBook(ctx, book)
}
func (s *Server) GetBook(ctx context.Context, bookId *proto.BookID) (*proto.Book, error) {
	return s.repo.GetBook(ctx, bookId)
}
func (s *Server) UpdateBook(ctx context.Context, book *proto.Book) (*wrappers.StringValue, error) {
	return s.repo.UpdateBook(ctx, book)
}
func (s *Server) DeleteBook(ctx context.Context, bookId *proto.BookID) (*wrappers.StringValue, error) {
	return s.repo.DeleteBook(ctx, bookId)
}
func (s *Server) SearchBookByName(ctx context.Context, book *proto.BookName) (*proto.BookList, error) {
	return s.repo.SearchBookByName(ctx, book)
}
func (s *Server) GetAll(ctx context.Context, in *wrapperspb.StringValue) (*proto.BookList, error) {
	return s.repo.GetAll(ctx, in)
}
