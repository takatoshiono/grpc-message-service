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
