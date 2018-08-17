package cloud_datastore

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/datastore"

	"github.com/takatoshiono/grpc-message-service/entity"
)

type ConversationRespository struct {
	client *datastore.Client
}

type conversation struct {
	ID        string
	CreatedAt time.Time
}

func NewConversationRepository() (*ConversationRespository, error) {
	client, err := NewCloudDatastoreClient()
	if err != nil {
		return nil, nil
	}
	return &ConversationRespository{client: client}, nil
}

func (r *ConversationRespository) Save(e *entity.Conversation) (*entity.Conversation, error) {
	ctx := context.Background()
	k := datastore.NameKey("Conversation", e.Name(), nil)
	c := &conversation{ID: e.ID, CreatedAt: e.CreatedAt}
	_, err := r.client.Put(ctx, k, c)
	if err != nil {
		log.Printf("Failed to save conversation: %v", err)
		return e, err
	}
	return e, nil
}

func (r *ConversationRespository) Get(name string) (*entity.Conversation, error) {
	ctx := context.Background()
	k := datastore.NameKey("Conversation", name, nil)
	c := new(conversation)
	if err := r.client.Get(ctx, k, c); err != nil {
		log.Printf("Failed to get conversation: %v", err)
		return nil, err
	}
	return &entity.Conversation{ID: c.ID, CreatedAt: c.CreatedAt}, nil
}
