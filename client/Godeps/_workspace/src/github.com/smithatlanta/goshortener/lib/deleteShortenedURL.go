package lib

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

func deleteShortenedURL(key string, projectID string) error {
	ctx := context.Background()

	// Create a datastore client. In a typical application, you would create
	// a single client which is reused for every datastore operation.
	dsClient, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a new client: %v", err)
		return err
	}

	k := datastore.NameKey("Entity", key, nil)
	if err := dsClient.Delete(ctx, k); err != nil {
		log.Fatalf("Failed to delete entity: %v", err)
		return err
	}

	return nil
}
