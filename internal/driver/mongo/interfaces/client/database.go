package client

import "go.mongodb.org/mongo-driver/mongo/options"

type Database interface {
	Collection(name string, opts ...*options.CollectionOptions) Collection
}
