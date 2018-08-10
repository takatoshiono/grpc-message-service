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
	k := datastore.NameKey("Conversation", e.ID, nil)
	c := &conversation{CreatedAt: e.CreatedAt}
	_, err := r.client.Put(ctx, k, c)
	if err != nil {
		log.Printf("Failed to save conversation: %v", err)
		return e, err
	}
	return e, nil
}

func (r *ConversationRespository) Get(ID string) (*entity.Conversation, error) {
	ctx := context.Background()
	k := datastore.NameKey("Conversation", ID, nil)
	c := new(conversation)
	if err := r.client.Get(ctx, k, c); err != nil {
		log.Printf("Failed to get conversation: %v", err)
		return nil, err
	}
	return conversationEntity(ID, c), nil
}

func conversationEntity(ID string, c *conversation) *entity.Conversation {
	return &entity.Conversation{ID: ID, CreatedAt: c.CreatedAt}
}
