package server

import (
	"context"
	"fmt"
	"log"

	"github.com/bufbuild/protovalidate-go"
	"github.com/meutraa/como/internal/proto"
)

func (s *Server) Login(ctx context.Context, in *proto.AuthRequest) (*proto.AuthResponse, error) {
	log.Printf("Received: %v", in.GetUsername())
	v, err := protovalidate.New()
	if nil != err {
		fmt.Println("failed to initialize validator:", err)
		return nil, err
	}
	if err = v.Validate(in); err != nil {
		fmt.Println("validation failed:", err)
		return nil, err
	}

	return &proto.AuthResponse{Token: "Hello " + in.GetUsername()}, nil
}
