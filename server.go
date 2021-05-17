package main

import (
	"log"
	"net"

	"github.com/err0r10/test_go_grpc_server"
	"google.golang.org/grpc"
)

const (
	port = ":5003"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := test_go_grpc_server.Server{}

	grpcServer := grpc.NewServer()

	test_go_grpc_server.RegisterHelloServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	grpc_server
}
