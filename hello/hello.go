package hello

import (
	"context"
	"log"
)

type Server struct{}

func (s *Server) GetMessage(ctx context.Context, in *RequestMessage) (*ResponseMessage, error) {
	log.Printf("Receive message body from client: %s", in.Req)
	return &ResponseMessage{Res: "Hello " + in.Req}, nil
}
