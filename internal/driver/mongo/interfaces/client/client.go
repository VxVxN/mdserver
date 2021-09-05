package client

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client interface {
	Database(name string, opts ...*options.DatabaseOptions) Database
	Disconnect(ctx context.Context) error
}
