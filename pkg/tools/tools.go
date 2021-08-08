package tools

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/VxVxN/log"
)

type closer interface {
	Close() error
}

func Close(file closer, message string) {
	if err := file.Close(); err != nil {
		log.Error.Printf("%s: %v", message, err)
	}
}

func CloseMongoCursor(cur *mongo.Cursor, ctx context.Context) {
	if err := cur.Close(ctx); err != nil {
		log.Error.Printf("Failed to close mongo cursor: %v", err)
	}
}

func ContainString(slice []string, elem string) bool {
	for _, elemSlice := range slice {
		if elemSlice == elem {
			return true
		}
	}
	return false
}
