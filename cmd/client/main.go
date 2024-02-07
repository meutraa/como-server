package main

import (
	"context"
	"log"
	"time"

	"github.com/meutraa/como/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewComoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Login(ctx, &proto.AuthRequest{Username: "Paulieoncsitheoncstiheoncsitehoncisetohnicseotnhicesotnhicesonth", Password: "12345678"})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}

	log.Printf("Response: %s", res.GetToken())
}
