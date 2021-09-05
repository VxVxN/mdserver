package mongo_client

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_interface "github.com/VxVxN/mdserver/internal/driver/mongo/interfaces/client"
)

type Database struct {
	*mongo.Database
}

func (db *Database) Collection(name string, opts ...*options.CollectionOptions) _interface.Collection {
	return &Collection{db.Database.Collection(name, opts...)}
}
