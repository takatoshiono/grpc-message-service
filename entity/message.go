package entity

import "time"

type Message struct {
	Id             string
	ConversationId string
	Sender         string
	Body           string
	CreatedAt      int64
}

func NewMessage(c Conversation, sender string, body string) *Message {
	return &Message{Id: "m123", ConversationId: c.Id, Sender: sender, Body: body, CreatedAt: time.Now().Unix()}
}
