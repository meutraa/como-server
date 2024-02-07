package main

import (
	"log"
	"net"

	"github.com/meutraa/como/internal/proto"
	"github.com/meutraa/como/internal/server"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterComoServer(s, &server.Server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
