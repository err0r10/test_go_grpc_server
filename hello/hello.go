package hello

import (
	"context"
	"log"
)

type Server struct{}

func (s *Server) GetMessage(ctx context.Context, in *RequestMessage) (*ResponseMessage, error) {
	log.Printf("Receive message body from client: %s", in.req)
	return &ResponseMessage{res: "Hello " + in.req}
}
