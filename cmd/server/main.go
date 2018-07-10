package main

import (
	"log"
	"net"

	pb "github.com/takatoshiono/grpc-message-service/proto"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":50101")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	msgSrv := NewMessageServer()
	pb.RegisterMessageServiceServer(s, msgSrv)
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
