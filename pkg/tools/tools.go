package tools

import (
	"context"
	"fmt"
	"io"
	"os"

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

func Copy(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer Close(source, "Failed to close source file, when copy file")

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer Close(destination, "Failed to close destination file, when copy file")
	_, err = io.Copy(destination, source)
	return err
}
