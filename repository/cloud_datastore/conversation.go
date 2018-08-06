package cloud_datastore

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"

	"github.com/takatoshiono/grpc-message-service/entity"
)

const projectID = "grpc-message-service"

type ConversationRespository struct {
	client *datastore.Client
}

func NewConversationRepository() *ConversationRespository {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &ConversationRespository{client: client}
}

func (r *ConversationRespository) Save(c *entity.Conversation) (*entity.Conversation, error) {
	ctx := context.Background()
	k := datastore.NameKey("Conversation", c.Id, nil)
	_, err := r.client.Put(ctx, k, c)
	if err != nil {
		log.Fatalf("Failed to save conversation: %v", err)
	}
	return c, nil
}
