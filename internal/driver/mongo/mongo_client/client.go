package mongo_client

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_interface "github.com/VxVxN/mdserver/internal/driver/mongo/interfaces/client"
)

type MongoClient struct {
	*mongo.Client
}

func NewClient() *MongoClient {
	return &MongoClient{}
}

func (c *MongoClient) Connect(ctx context.Context, uri string) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	c.Client = client
	return err
}

func (c *MongoClient) Database(name string, opts ...*options.DatabaseOptions) _interface.Database {
	return &Database{c.Client.Database(name, opts...)}
}

func (c *MongoClient) Disconnect(ctx context.Context) error {
	return c.Client.Disconnect(ctx)
}
