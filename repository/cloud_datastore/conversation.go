package cloud_datastore

import (
	"github.com/takatoshiono/grpc-message-service/entity"
)

type ConversationRespository struct{}

func NewConversationRepository() *ConversationRespository {
	return &ConversationRespository{}
}

func (r *ConversationRespository) Save(c *entity.Conversation) (*entity.Conversation, error) {
	return c, nil
}
