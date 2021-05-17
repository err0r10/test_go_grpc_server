package main

import (
	"log"
	"net"

	"github.com/err0r10/test_go_grpc_server/hello"
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

	grpcServer := grpc.NewServer()

	hello.RegisterHelloServer(grpcServer, &hello.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
