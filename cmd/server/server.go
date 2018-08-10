package main

import (
	"context"
	"fmt"
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
	// parent must have a relative resource name of the Conversation.
	parentNameParts := strings.Split(in.Parent, "/")
	if !(len(parentNameParts) == 2 && parentNameParts[0] == "conversations") {
		return nil, fmt.Errorf("parent not found: %s", in.Parent)
	}

	cRepo, err := cloud_datastore.NewConversationRepository()
	if err != nil {
		log.Fatalf("failed to create ConversationRepository: %v", err)
	}

	c, err := cRepo.Get(parentNameParts[1])
	if err != nil {
		log.Fatalf("Failed to get conversation: %v", err)
	}

	m := entity.NewMessage(*c, in.Message.Sender, in.Message.Body)
	r, err := cloud_datastore.NewMessageRepository()
	if err != nil {
		log.Fatalf("Failed to create MessageRepository: %v", err)
	}
	r.Save(m)
	createTime, err := ptypes.TimestampProto(m.CreatedAt)
	if err != nil {
		log.Fatalf("Failed to create Timestamp proto: %v", err)
	}
	return &pb.Message{Name: m.Name(), Sender: m.Sender, Body: m.Body, CreateTime: createTime}, nil
}

func (s *server) GetMessages(ctx context.Context, in *pb.GetMessagesRequest) (*pb.Messages, error) {
	name := strings.Join([]string{in.Parent, "messages/msg111"}, "/")
	message := &pb.Message{Name: name, Sender: "alice", Body: "Hello", CreateTime: ptypes.TimestampNow()}
	messages := &pb.Messages{Messages: []*pb.Message{message}}
	return messages, nil
}
