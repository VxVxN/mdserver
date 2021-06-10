package tools

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/VxVxN/log"
)

func CloseFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Error.Printf("Failed to close file: %v", err)
	}
}

func CloseMongoCursor(cur *mongo.Cursor, ctx context.Context) {
	if err := cur.Close(ctx); err != nil {
		log.Error.Printf("Failed to close mongo cursor: %v", err)
	}
}
