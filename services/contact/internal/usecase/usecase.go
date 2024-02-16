package usecase

import (
	"architecture_go/services/contact/internal/domain"
	"architecture_go/services/contact/internal/repository"
)

type ContactUseCaseImpl struct {
	repository domain.ContactRepository
}

func NewContactUseCase(repository repository.ContactRepository) ContactUseCase {
	return &ContactUseCaseImpl{
		repository: repository,
	}
}

type ContactUseCase interface {
	// Contact model
	CreateContact(contact domain.Contact) (int, error)
	GetContact(id int) (*domain.Contact, error)
	UpdateContact(contact domain.Contact) error
	DeleteContact(id int) error

	// Group model
	CreateGroup(group domain.Group) (int, error)
	GetGroup(id int) (*domain.Group, error)

	// ContactGroup model
	AddContactToGroup(contactID, groupID int) error
}
