package lib

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

// GetShortenedURL - creates shortened url
func GetShortenedURL(key string, projectID string) (string, error) {
	return getShortenedURL(key, projectID)
}

func getShortenedURL(key string, projectID string) (string, error) {
	ctx := context.Background()

	// Create a datastore client. In a typical application, you would create
	// a single client which is reused for every datastore operation.
	dsClient, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a new client: %v", err)
		return "", err
	}

	k := datastore.NameKey("URLShortener", key, nil)
	e := new(URLShortener)
	if err := dsClient.Get(ctx, k, e); err != nil {
		log.Fatalf("Failed to get entity: %v", err)
		return "", err
	}

	return e.Value, nil
}
