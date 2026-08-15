package service

import (
	"errors"

	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/configuration/logger"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/repository"
)

type ContactService interface {
	Create(domain domains.ContactDomainInterface) (domains.ContactDomainInterface, error)
	FindAll(ownerID string) ([]domains.ContactDomainInterface, error)
	FindByID(id, ownerID string) (domains.ContactDomainInterface, error)
	Update(domain domains.ContactDomainInterface) (domains.ContactDomainInterface, error)
	Delete(id, ownerID string) error
}

type contactService struct {
	repo repository.ContactRepository
}

func NewContactService(repo repository.ContactRepository) ContactService {
	return &contactService{repo: repo}
}

func (s *contactService) Create(domain domains.ContactDomainInterface) (domains.ContactDomainInterface, error) {
	created, err := s.repo.Create(domain)
	if err != nil {
		logger.Error("error creating contact", err)
		return nil, err
	}

	return created, nil
}

func (s *contactService) FindAll(ownerID string) ([]domains.ContactDomainInterface, error) {
	contacts, err := s.repo.FindAll(ownerID)
	if err != nil {
		logger.Error("error finding all contacts", err)
		return nil, err
	}

	return contacts, nil
}

func (s *contactService) FindByID(id, ownerID string) (domains.ContactDomainInterface, error) {
	contact, err := s.repo.FindByID(id)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding contact by id", err)
		}
		return nil, err
	}

	if contact.GetOwnerID() != ownerID {
		return nil, domains.NewForbiddenError("access denied")
	}

	return contact, nil
}

func (s *contactService) Update(domain domains.ContactDomainInterface) (domains.ContactDomainInterface, error) {
	existing, err := s.repo.FindByID(domain.GetID())
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding contact for update", err)
		}
		return nil, err
	}

	if existing.GetOwnerID() != domain.GetOwnerID() {
		return nil, domains.NewForbiddenError("access denied")
	}

	updated, err := s.repo.Update(domain)
	if err != nil {
		logger.Error("error updating contact", err)
		return nil, err
	}

	return updated, nil
}

func (s *contactService) Delete(id, ownerID string) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding contact for delete", err)
		}
		return err
	}

	if existing.GetOwnerID() != ownerID {
		return domains.NewForbiddenError("access denied")
	}

	if err := s.repo.Delete(id); err != nil {
		logger.Error("error deleting contact", err)
		return err
	}

	return nil
}
