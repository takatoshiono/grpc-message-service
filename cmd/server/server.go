package main

import (
	"context"
	"log"
	"strings"

	"github.com/golang/protobuf/ptypes"
	"github.com/takatoshiono/grpc-message-service/entity"
	pb "github.com/takatoshiono/grpc-message-service/proto"
	"github.com/takatoshiono/grpc-message-service/repository/cloud_datastore"
)

type server struct{}

func NewMessageServer() *server {
	return &server{}
}

func (s *server) CreateConversation(ctx context.Context, in *pb.CreateConversationRequest) (*pb.Conversation, error) {
	c := entity.NewConversation()
	r, err := cloud_datastore.NewConversationRepository()
	if err != nil {
		// FIXME: What should we do here?
		log.Fatalf("Failed to create ConversationRepository: %v", err)
	}
	r.Save(c)
	createTime, err := ptypes.TimestampProto(c.CreatedAt)
	if err != nil {
		log.Fatalf("Failed to create Timestamp proto: %v", err)
	}
	return &pb.Conversation{Name: c.Name(), CreateTime: createTime}, nil
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
