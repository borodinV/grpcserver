package main

import "grpc/server/repo"

type Server struct {
	repo *repo.Repository
}

func NewServer(repo *repo.Repository) *Server {
	return &Server{repo: repo}
}
