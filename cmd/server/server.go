package main

import (
	"context"
	"log"

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
	cRepo, err := cloud_datastore.NewConversationRepository()
	if err != nil {
		log.Fatalf("failed to create ConversationRepository: %v", err)
	}

	c, err := cRepo.Get(in.Parent)
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
	r, err := cloud_datastore.NewMessageRepository()
	if err != nil {
		log.Fatalf("Failed to create MessageRepository: %v", err)
	}

	entities, err := r.List(in.Parent)
	if err != nil {
		log.Fatalf("Failed to get messages: %v", err)
	}

	var messages = make([]*pb.Message, len(entities))
	for i, e := range entities {
		createTime, err := ptypes.TimestampProto(e.CreatedAt)
		if err != nil {
			log.Fatalf("Failed to create Timestamp proto: %v", err)
		}
		messages[i] = &pb.Message{Name: e.Name(), Sender: e.Sender, Body: e.Body, CreateTime: createTime}
	}
	return &pb.Messages{Messages: messages}, nil
}
