package pool

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

type StorageObject interface {
	GetStorageClient() storage.Client
}

type storageObject struct {
	client storage.Client
}

var instance *storageObject

func GetInstance() StorageObject {
	if instance == nil {
		instance = new(storageObject)
	}
	return instance
}

func (s *storageObject) GetStorageClient() storage.Client {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "histalker-b9672-firebase-adminsdk-vt2os-7293908177.json")
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create a storage client: %v", err)
	}
	return *client
}
