package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go_api/load_config"
	"go_api/proto_gen"
	"go_api/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func runningGRPCGateWay(config load_config.Config) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	err := proto_gen.RegisterDemoHandlerServer(context.Background(), mux, &service.Server{})
	if err != nil {
		panic(err)
	}
	log.Printf("GRPC-Gateway: Listening to %s ...\n", config.GatewayAddr)
	log.Fatalln(http.ListenAndServe(config.GatewayAddr, mux))
}

func runningGRPCServer(config load_config.Config) {
	lis, err := net.Listen("tcp", config.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto_gen.RegisterDemoServer(s, &service.Server{})
	log.Printf("server listening at %v\n", lis.Addr())
	go func() {
		log.Fatal(s.Serve(lis))
	}()
}

func main() {
	config := load_config.Conf
	runningGRPCServer(config)
	runningGRPCGateWay(config)
}
