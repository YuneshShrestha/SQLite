package sqlite

import (
	"io"
	"os"

	"rsqlite/internal/domain"
)

type FileReader struct{}

func NewFileReader() *FileReader {
	return &FileReader{}
}

func (r *FileReader) ReadHeader(
	path string,
) (domain.DatabaseHeader, error) {
	file, err := os.Open(path)
	if err != nil {
		return domain.DatabaseHeader{}, err
	}
	defer file.Close()

	data := make([]byte, 100)

	_, err = io.ReadFull(file, data)
	if err != nil {
		return domain.DatabaseHeader{}, err
	}

	return ParseHeader(data)
}
