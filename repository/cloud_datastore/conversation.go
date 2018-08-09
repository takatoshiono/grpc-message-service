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
