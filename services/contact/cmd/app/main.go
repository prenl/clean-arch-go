package main

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal/domain"
	"architecture_go/services/contact/internal/repository"
	"context"
	"log"
)

func main() {
	log.Println("Starting the application...")
    conn, err := postgres.Connect("localhost", 5432, "postgres", "1112", "clean-arch-go")
	log.Println("Connected to the database...")
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer conn.Close(context.Background())

    contactRepo := repository.NewContactRepository(conn)
	contactRepo.CreateContact(domain.Contact{ FirstName: "Yelnur", MiddleName: "Maratovich", LastName: "Abdrakhmanov", PhoneNumber: "+77015130153"})

	contact, err := contactRepo.GetContact(1)
	log.Println("Retrieving contact...")

	if err != nil {
		log.Fatalf("Failed to retrieve contact: %v", err)
	}
	log.Printf("Retrieved contact: %v", contact)
}
