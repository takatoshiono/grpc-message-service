package entity

import "time"

type Conversation struct {
	Id        string
	CreatedAt int64
}

func NewConversation() *Conversation {
	return &Conversation{Id: "abcd1234", CreatedAt: time.Now().Unix()}
}
