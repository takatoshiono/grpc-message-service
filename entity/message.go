package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID             string
	ConversationID string
	Sender         string
	Body           string
	CreatedAt      time.Time
}

func NewMessage(c Conversation, sender string, body string) *Message {
	return &Message{ID: uuid.New().String(), ConversationID: c.ID, Sender: sender, Body: body, CreatedAt: time.Now()}
}

func (m *Message) Name() string {
	return fmt.Sprintf("conversations/%s/messages/%s", m.ConversationID, m.ID)
}
