package tools

import (
	"fmt"
	"io"
	"os"
)

// CopyFile - copy one file
func CopyFile(src, dst string) error {
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

// GetFileNamesInDir - return the file names from the directory, if a directory is found, then skip it
func GetFileNamesInDir(pathToDir string) ([]string, error) {
	dirEntry, err := os.ReadDir(pathToDir)
	if err != nil {

	}
	result := make([]string, 0, len(dirEntry))
	for _, entry := range dirEntry {
		if entry.IsDir() {
			continue
		}
		result = append(result, entry.Name())
	}
	return result, nil
}
