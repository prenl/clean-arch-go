package main

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal/delivery"
	"architecture_go/services/contact/internal/repository"
	"architecture_go/services/contact/internal/usecase"
	"context"
	"log"
	"net/http"
)

func main() {
	db, err := postgres.Connect("localhost", 5432, "postgres", "1112", "clean-arch-go")

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	defer db.Close(context.Background())
	log.Println("Connected to the database")

	repo := repository.NewContactRepository()
	usecase := usecase.NewContactUseCase(repo)
	delivery := delivery.NewContactDelivery(usecase)

	_ = delivery

	log.Println("Server is starting on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

}
