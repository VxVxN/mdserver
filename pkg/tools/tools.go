package tools

import (
	"os"

	"github.com/VxVxN/log"
)

func CloseFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Error.Printf("Failed to close file: %v", err)
	}
}
