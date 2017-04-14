package lib

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

// CreateShortenedURL - creates shortened url
func CreateShortenedURL(key string, value string, projectID string) error {
	return createShortenedURL(key, value, projectID)
}

func createShortenedURL(key string, value string, projectID string) error {
	ctx := context.Background()

	// Create a datastore client. In a typical application, you would create
	// a single client which is reused for every datastore operation.
	dsClient, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a new client: %v", err)
		return err
	}

	k := datastore.NameKey("URLShortener", key, nil)
	e := new(URLShortener)
	if err := dsClient.Get(ctx, k, e); err != nil {
		// log.Fatalf("Failed to get entity: %v", err)
		// return err
	}

	e.Value = value

	if _, err := dsClient.Put(ctx, k, e); err != nil {
		log.Fatalf("Failed to put entity: %v", err)
		return err
	}

	return nil
}
