package mocks

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_interface "github.com/VxVxN/mdserver/internal/driver/mongo/interfaces/client"
)

type MongoClient struct{}

func (c *MongoClient) Database(name string, opts ...*options.DatabaseOptions) *Database {
	return new(Database)
}

func (c *MongoClient) Disconnect(ctx context.Context) error {
	return nil
}

type Database struct{}

type Collection struct{}

func (db *Database) Collection(name string, opts ...*options.CollectionOptions) _interface.Collection {
	return new(Collection)
}

func (coll *Collection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return new(mongo.Cursor), nil
}

func (coll *Collection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return new(mongo.SingleResult)
}

func (coll *Collection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return new(mongo.InsertOneResult), nil
}

func (coll *Collection) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return new(mongo.InsertManyResult), nil
}

func (coll *Collection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return new(mongo.UpdateResult), nil
}

func (coll *Collection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return new(mongo.UpdateResult), nil
}

func (coll *Collection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return new(mongo.DeleteResult), nil
}

func (coll *Collection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return new(mongo.DeleteResult), nil
}
