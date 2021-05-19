package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/brianvoe/gofakeit"
	"github.com/err0r10/test_go_grpc_server/hello"
)

const (
	port = ":5003"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := hello.NewHelloClient(conn)

	response, err := c.GetMessage(context.Background(), &hello.RequestMessage{Req: gofakeit.BeerName()})
	if err != nil {
		log.Fatalf("Error wher calling GetMessage: %s", err)
	}
	log.Printf("Response from server: %s", response.Res)
}
