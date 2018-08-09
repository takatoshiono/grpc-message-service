package entity

import (
	"time"

	"github.com/google/uuid"
)

type Conversation struct {
	ID        string
	CreatedAt time.Time
}

func NewConversation() *Conversation {
	return &Conversation{ID: uuid.New().String(), CreatedAt: time.Now()}
}
