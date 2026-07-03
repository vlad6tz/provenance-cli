package utils

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
)

var (
	ErrFileNotFound = errors.New("the specified file could not be found")
	ErrEmptyFile    = errors.New("the file is empty")
)

type FileMetadata struct {
	Name     string
	Size     int64
	MimeType string
}

func GetFileMetadata(path string) (*FileMetadata, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileNotFound
		}
		return nil, err
	}

	if info.Size() == 0 {
		return nil, ErrEmptyFile
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Get MIME type using standard 512 buffer
	buffer := make([]byte, 512)
	n, _ := file.Read(buffer)
	mimeType := http.DetectContentType(buffer[:n])

	return &FileMetadata{
		Name:     filepath.Base(path),
		Size:     info.Size(),
		MimeType: mimeType,
	}, nil
}
