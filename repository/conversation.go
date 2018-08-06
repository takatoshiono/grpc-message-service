package repository

import "github.com/takatoshiono/grpc-message-service/entity"

type ConversationRespository interface {
	Save(*entity.Conversation) (*entity.Conversation, error)
}
