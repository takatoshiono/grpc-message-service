package repository

import "github.com/takatoshiono/grpc-message-service/entity"

type MessageRepository interface {
	Save(e *entity.Message) (*entity.Message, error)
	List(conversationName string) ([]*entity.Message, error)
}
