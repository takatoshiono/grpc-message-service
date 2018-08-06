package cloud_datastore

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"

	"github.com/takatoshiono/grpc-message-service/entity"
)

type ConversationRespository struct {
	client *datastore.Client
}

func NewConversationRepository() (*ConversationRespository, error) {
	client, err := NewCloudDatastoreClient()
	if err != nil {
		return nil, nil
	}
	return &ConversationRespository{client: client}, nil
}

func (r *ConversationRespository) Save(c *entity.Conversation) (*entity.Conversation, error) {
	ctx := context.Background()
	k := datastore.NameKey("Conversation", c.Id, nil)
	_, err := r.client.Put(ctx, k, c)
	if err != nil {
		log.Printf("Failed to save conversation: %v", err)
		return c, err
	}
	return c, nil
}
