package main

import (
	"context"
	"go_api/proto_gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func GRPCReqHandling(ctx context.Context, s proto_gen.DemoClient) {
	res, err := s.HelloWorld(ctx, &proto_gen.EmptyMessage{})
	if err != nil {
		panic(err)
	}
	println(res.Msg)
}

func main() {
	// Create gRPC connection from client to server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto_gen.NewDemoClient(conn)

	// Test GRPC API
	GRPCReqHandling(context.Background(), c)
}
