package application

import (
	"rsqlite/internal/domain"
)

type DatabaseReader interface {
	ReadHeader(path string) (domain.DatabaseHeader, error)
}

type ReadHeaderUseCase struct {
	reader DatabaseReader
}

func NewReadHeaderUseCase(
	reader DatabaseReader,
) *ReadHeaderUseCase {
	return &ReadHeaderUseCase{
		reader: reader,
	}
}

func (uc *ReadHeaderUseCase) Execute(
	path string,
) (domain.DatabaseHeader, error) {
	return uc.reader.ReadHeader(path)
}
