package main

import (
	"fmt"
	"os"
	"rsqlite/internal/application"
	"rsqlite/internal/infrastructure/sqlite"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: rsqlite <database-file>")
		os.Exit(1)
	}

	reader := sqlite.NewFileReader()

	useCase := application.NewReadHeaderUseCase(reader)

	header, err := useCase.Execute(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("SQLite database opened successfully!")
	fmt.Printf("Page size: %d bytes\n", header.PageSize)
}
