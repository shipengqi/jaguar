package main

import (
	"context"
	"log"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	pb "{{ .App.ModuleName }}/pkg/api/proto/v1"
)

func main() {
	// this is a retry example
	// by default, retriable codes are ResourceExhausted,Unavailable
	opts := []retry.CallOption{
		retry.WithMax(5),
		retry.WithPerRetryTimeout(5 * time.Second),
		retry.WithCodes(codes.Unauthenticated),
	}

	conn, err := grpc.Dial("127.0.0.1:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(opts...)),
		grpc.WithStreamInterceptor(retry.StreamClientInterceptor(opts...)),
	)
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer func() { _ = conn.Close() }()

	client := pb.NewUserClient(conn)
	// there are two ways to create metadata
	// 1. metadata.New
	md := metadata.New(map[string]string{"authorization": "bearer example-only"})
	// 2. metadata.Pairs
	// metadata.Pairs("authorization", "bearer example-only")

	// setup metadata
	// OutgoingContext is used to send requests
	// InComingContext is used to receive requests
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, err := client.Create(ctx, &pb.CreateRequest{Request: "gRPC"})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	log.Printf("resp: %s", resp.GetResponse())
}