package firestores

import (
	"cloud.google.com/go/firestore"
	"context"
)

type Firestores struct {
	client *firestore.Client
}

func NewFirestores(ctx context.Context) (*Firestores, error) {
	client, err := firestore.NewClient(ctx, "")
	if err != nil {
		return nil, err
	}
	return &Firestores{
		client: client,
	}, nil
}
