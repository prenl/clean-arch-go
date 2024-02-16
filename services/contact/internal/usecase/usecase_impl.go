package usecase

import (
	"architecture_go/services/contact/internal/domain"
	"architecture_go/services/contact/internal/repository"
)

type ContactUseCaseImpl struct {
	repository repository.ContactRepository
}

func NewContactUseCase(repository repository.ContactRepository) ContactUseCase {
	return &ContactUseCaseImpl{
		repository: repository,
	}
}

func (c *ContactUseCaseImpl) CreateContact(contact domain.Contact) (int, error) {
	return c.repository.CreateContact(contact)
}

func (c *ContactUseCaseImpl) GetContact(id int) (*domain.Contact, error) {
	return c.repository.GetContact(id)
}

func (c *ContactUseCaseImpl) UpdateContact(contact domain.Contact) error {
	return c.repository.UpdateContact(contact)
}

func (c *ContactUseCaseImpl) DeleteContact(id int) error {
	return c.repository.DeleteContact(id)
}

func (c *ContactUseCaseImpl) CreateGroup(group domain.Group) (int, error) {
	return c.repository.CreateGroup(group)
}

func (c *ContactUseCaseImpl) GetGroup(id int) (*domain.Group, error) {
	return c.repository.GetGroup(id)
}

func (c *ContactUseCaseImpl) AddContactToGroup(contactID, groupID int) error {
	return c.repository.AddContactToGroup(contactID, groupID)
}
