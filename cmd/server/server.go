package main

import (
	"context"
	"time"

	pb "github.com/takatoshiono/grpc-message-service/proto"
)

type server struct{}

func NewMessageServer() *server {
	return &server{}
}

func (s *server) CreateConversation(context.Context, *pb.CreateConversationRequest) (*pb.Conversation, error) {
	conversation := &pb.Conversation{Id: 1, CreatedAt: time.Now().Unix()}
	return conversation, nil
}

func (s *server) CreateMessage(context.Context, *pb.CreateMessageRequest) (*pb.Message, error) {
	message := &pb.Message{ConversationId: 1, Body: "Hello", CreatedAt: time.Now().Unix()}
	return message, nil
}

func (s *server) GetMessages(context.Context, *pb.GetMessagesRequest) (*pb.Messages, error) {
	message := &pb.Message{ConversationId: 1, Body: "Hello", CreatedAt: time.Now().Unix()}
	messages := &pb.Messages{Messages: []*pb.Message{message}}
	return messages, nil
}
