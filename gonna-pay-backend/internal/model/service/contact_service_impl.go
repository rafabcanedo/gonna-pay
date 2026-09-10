package service

import (
	"errors"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/configuration/logger"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository"
)

type ContactService interface {
	Create(contact *domains.Contact) (*domains.Contact, error)
	FindAll(ownerID string) ([]*domains.Contact, error)
	FindByID(id, ownerID string) (*domains.Contact, error)
	Update(contact *domains.Contact) (*domains.Contact, error)
	Delete(id, ownerID string) error
	FindContactsByFrequency(userID string, limit int) ([]domains.ContactFrequency, error)
}

type contactService struct {
	repo repository.ContactRepository
}

func NewContactService(repo repository.ContactRepository) ContactService {
	return &contactService{repo: repo}
}

func (s *contactService) Create(contact *domains.Contact) (*domains.Contact, error) {
	exists, err := s.repo.ExistsByEmailAndOwner(contact.Email, contact.OwnerID)
	if err != nil {
		logger.Error("error checking contact uniqueness", err)
		return nil, err
	}
	if exists {
		return nil, domains.NewConflictError("contact with this email already exists")
	}

	created, err := s.repo.Create(contact)
	if err != nil {
		logger.Error("error creating contact", err)
		return nil, err
	}

	return created, nil
}

func (s *contactService) FindAll(ownerID string) ([]*domains.Contact, error) {
	contacts, err := s.repo.FindAll(ownerID)
	if err != nil {
		logger.Error("error finding all contacts", err)
		return nil, err
	}

	return contacts, nil
}

func (s *contactService) findAndAuthorize(id, ownerID string) (*domains.Contact, error) {
	contact, err := s.repo.FindByID(id)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding contact", err)
		}
		return nil, err
	}

	if contact.OwnerID != ownerID {
		return nil, domains.NewForbiddenError("access denied")
	}

	return contact, nil
}

func (s *contactService) FindByID(id, ownerID string) (*domains.Contact, error) {
	return s.findAndAuthorize(id, ownerID)
}

func (s *contactService) Update(contact *domains.Contact) (*domains.Contact, error) {
	existing, err := s.findAndAuthorize(contact.ID, contact.OwnerID)
	if err != nil {
		return nil, err
	}

	if contact.Name != "" {
		existing.Name = contact.Name
	}
	if contact.Email != "" {
		existing.Email = contact.Email
	}
	if contact.Phone != "" {
		existing.Phone = contact.Phone
	}
	if contact.Category != "" {
		existing.Category = contact.Category
	}

	updated, err := s.repo.Update(existing)
	if err != nil {
		logger.Error("error updating contact", err)
		return nil, err
	}

	return updated, nil
}

func (s *contactService) Delete(id, ownerID string) error {
	if _, err := s.findAndAuthorize(id, ownerID); err != nil {
		return err
	}

	if err := s.repo.Delete(id); err != nil {
		logger.Error("error deleting contact", err)
		return err
	}

	return nil
}

func (s *contactService) FindContactsByFrequency(userID string, limit int) ([]domains.ContactFrequency, error) {
	if limit <= 0 {
		limit = 5
	}

	contacts, err := s.repo.FindContactsByFrequency(userID, limit)
	if err != nil {
		logger.Error("error finding top contacts", err)
		return nil, err
	}

	return contacts, nil
}
