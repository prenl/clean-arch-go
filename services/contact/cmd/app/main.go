package main

import (
	"architecture_go/pkg/store/postgres"
	"context"
	"log"
)

func main() {
	db, err := postgres.Connect("localhost", 5432, "postgres", "1112", "clean-arch-go")

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	defer db.Close(context.Background())

	log.Println("Connected to the database")
}
