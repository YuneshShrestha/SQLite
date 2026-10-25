package sqlite

import (
	"encoding/binary"
	"errors"
	"fmt"

	"rsqlite/internal/domain"
)

var sqliteMagic = []byte("SQLite format 3\x00")

func ParseHeader(data []byte) (domain.DatabaseHeader, error) {
	if len(data) < 100 {
		return domain.DatabaseHeader{}, errors.New(
			"database header must contain at least 100 bytes",
		)
	}

	if string(data[:16]) != string(sqliteMagic) {
		return domain.DatabaseHeader{}, errors.New(
			"invalid SQLite database file",
		)
	}

	// Bytes 16 and 17 contain the page size.
	rawPageSize := binary.BigEndian.Uint16(data[16:18])

	pageSize := uint32(rawPageSize)

	// SQLite uses 1 to represent 65536.
	/*
		SQLite stores the page size using only 2 bytes.
		- 2 bytes can represent values from 0 to 65,535.
		- But SQLite allows a page size of 65,536 bytes.
		- So 65,536 cannot fit in those 2 bytes.
		- That's why SQLite uses 1 to represent 65536.
	*/
	if rawPageSize == 1 {
		pageSize = 65536
	}

	if pageSize < 512 || pageSize > 65536 {
		return domain.DatabaseHeader{}, fmt.Errorf(
			"invalid page size: %d", pageSize,
		)
	}

	// Check that the page size is a power of two.
	if pageSize&(pageSize-1) != 0 {
		return domain.DatabaseHeader{}, fmt.Errorf(
			"page size %d is not a power of two", pageSize,
		)
	}

	return domain.DatabaseHeader{
		PageSize: pageSize,
	}, nil
}
