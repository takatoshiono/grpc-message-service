package repository

import "github.com/takatoshiono/grpc-message-service/entity"

type MessageRepository interface {
	Save(m *entity.Message) (*entity.Message, error)
}
