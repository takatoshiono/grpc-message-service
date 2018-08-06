package cloud_datastore

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

const projectID = "grpc-message-service"

func NewCloudDatastoreClient() *datastore.Client {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}
