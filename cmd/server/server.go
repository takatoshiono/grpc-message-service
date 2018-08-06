package main

import (
	"context"
	"strings"

	"github.com/golang/protobuf/ptypes"
	pb "github.com/takatoshiono/grpc-message-service/proto"
)

type server struct{}

func NewMessageServer() *server {
	return &server{}
}

func (s *server) CreateConversation(ctx context.Context, in *pb.CreateConversationRequest) (*pb.Conversation, error) {
	conversation := &pb.Conversation{Name: "conversations/abc123", CreateTime: ptypes.TimestampNow()}
	return conversation, nil

	//	c := entity.NewConversation()
	//	r, err := cloud_datastore.NewConversationRepository()
	//	if err != nil {
	//		// FIXME: What should we do here?
	//		log.Fatalf("Failed to create ConversationRepository: %v", err)
	//	}
	//	r.Save(c)
	//
	//	return &pb.Conversation{Id: c.Id, CreatedAt: c.CreatedAt}, nil
}

func (s *server) CreateMessage(ctx context.Context, in *pb.CreateMessageRequest) (*pb.Message, error) {
	name := strings.Join([]string{in.Parent, "messages/msg111"}, "/")
	message := &pb.Message{Name: name, Sender: in.Message.Sender, Body: in.Message.Body, CreateTime: ptypes.TimestampNow()}
	return message, nil
}

func (s *server) GetMessages(ctx context.Context, in *pb.GetMessagesRequest) (*pb.Messages, error) {
	name := strings.Join([]string{in.Parent, "messages/msg111"}, "/")
	message := &pb.Message{Name: name, Sender: "alice", Body: "Hello", CreateTime: ptypes.TimestampNow()}
	messages := &pb.Messages{Messages: []*pb.Message{message}}
	return messages, nil
}
