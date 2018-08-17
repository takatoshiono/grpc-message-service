package repository

import "github.com/takatoshiono/grpc-message-service/entity"

type ConversationRespository interface {
	Save(e *entity.Conversation) (*entity.Conversation, error)
	Get(name string) (*entity.Conversation, error)
}
