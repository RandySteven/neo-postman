package firestores

import (
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type Firestores struct {
	bucket *storage.BucketHandle
}

func NewFirestores(ctx context.Context) (*Firestores, error) {
	config := &firebase.Config{
		StorageBucket: "<BUCKET_NAME>.appspot.com",
	}
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		return nil, err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return nil, err
	}

	return &Firestores{
		bucket: bucket,
	}, nil
}
