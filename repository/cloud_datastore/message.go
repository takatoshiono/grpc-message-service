package cloud_datastore

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/takatoshiono/grpc-message-service/entity"
)

type MessageRepository struct {
	client *datastore.Client
}

type message struct {
	ID        string
	Sender    string
	Body      string
	CreatedAt time.Time
}

func NewMessageRepository() (*MessageRepository, error) {
	client, err := NewCloudDatastoreClient()
	if err != nil {
		return nil, nil
	}
	return &MessageRepository{client: client}, nil
}

func (r *MessageRepository) Save(e *entity.Message) (*entity.Message, error) {
	ctx := context.Background()

	c := entity.Conversation{ID: e.ConversationID}
	parentKey := datastore.NameKey("Conversation", c.Name(), nil)

	k := datastore.NameKey("Message", e.Name(), parentKey)
	m := &message{ID: e.ID, Sender: e.Sender, Body: e.Body, CreatedAt: e.CreatedAt}
	_, err := r.client.Put(ctx, k, m)
	if err != nil {
		log.Printf("Failed to create message: %v", err)
		return e, err
	}
	return e, nil
}

func (r *MessageRepository) List(conversationName string) ([]*entity.Message, error) {
	k := datastore.NameKey("Conversation", conversationName, nil)
	q := datastore.NewQuery("Message").Ancestor(k)

	ctx := context.Background()
	var messages []*message
	_, err := r.client.GetAll(ctx, q, &messages)
	if err != nil {
		log.Printf("Failed to get messages: %v", err)
		return nil, err
	}

	entities := make([]*entity.Message, len(messages))
	for i, m := range messages {
		entities[i] = &entity.Message{ID: m.ID, Sender: m.Sender, Body: m.Body, CreatedAt: m.CreatedAt}
	}
	return entities, err
}
