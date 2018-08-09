package entity

import (
	"time"

	"github.com/google/uuid"
)

type Conversation struct {
	ID        string
	CreatedAt int64
}

func NewConversation() *Conversation {
	return &Conversation{ID: uuid.New().String(), CreatedAt: time.Now().Unix()}
}
