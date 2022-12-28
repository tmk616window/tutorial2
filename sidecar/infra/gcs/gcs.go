package gcs

import (
	"context"

	"cloud.google.com/go/storage"
)

type client struct {
    *storage.Client
}

type ClientInterface interface {

}

func NewClient () (client, error) {
	c, err := storage.NewClient(context.Background())
	if err != nil {
		return client{nil}, err
	}

	return client{c}, nil
}