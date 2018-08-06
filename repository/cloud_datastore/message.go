package cloud_datastore

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/takatoshiono/grpc-message-service/entity"
)

type MessageRepository struct {
	client *datastore.Client
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{client: NewCloudDatastoreClient()}
}

func (r *MessageRepository) Save(m *entity.Message) (*entity.Message, error) {
	ctx := context.Background()
	parentKey := datastore.NameKey("Conversation", m.ConversationId, nil)
	k := datastore.NameKey("Message", m.Id, parentKey)
	_, err := r.client.Put(ctx, k, m)
	if err != nil {
		log.Fatalf("Failed to create message: %v", err)
		return m, err
	}
	return m, nil
}
