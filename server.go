package main

import (
	"log"
	"net"

	"github.com/err0r10/test_go_grpc_server/hello"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":5003")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	hello.RegisterHelloServer(grpcServer, &hello.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
