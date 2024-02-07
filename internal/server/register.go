package server

import (
	"context"

	"github.com/meutraa/como/internal/proto"
)

func (s *Server) Register(ctx context.Context, in *proto.AuthRequest) (*proto.AuthResponse, error) {
	// validate the request, hash the password, save the user to the database

	return &proto.AuthResponse{Token: "Hello " + in.GetUsername()}, nil
}
