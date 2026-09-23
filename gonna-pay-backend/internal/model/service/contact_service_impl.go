package service

import (
	"context"
	"errors"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/configuration/logger"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository"
)

type ContactService interface {
	Create(ctx context.Context, contact *domains.Contact) (*domains.Contact, error)
	FindAll(ctx context.Context, ownerID string, page, limit int, filters domains.ContactFilters) ([]*domains.Contact, int64, error)
	FindByID(ctx context.Context, id, ownerID string) (*domains.Contact, error)
	Update(ctx context.Context, contact *domains.Contact) (*domains.Contact, error)
	Delete(ctx context.Context, id, ownerID string) error
	FindContactsByFrequency(ctx context.Context, userID string, limit int) ([]domains.ContactFrequency, error)
	FindStats(ctx context.Context, ownerID string) (*domains.ContactStats, error)
}

type contactService struct {
	repo repository.ContactRepository
}

func NewContactService(repo repository.ContactRepository) ContactService {
	return &contactService{repo: repo}
}

func (s *contactService) Create(ctx context.Context, contact *domains.Contact) (*domains.Contact, error) {
	exists, err := s.repo.ExistsByEmailAndOwner(ctx, contact.Email, contact.OwnerID)
	if err != nil {
		logger.Error("error checking contact uniqueness", err)
		return nil, err
	}
	if exists {
		return nil, domains.NewConflictError("contact with this email already exists")
	}

	created, err := s.repo.Create(ctx, contact)
	if err != nil {
		logger.Error("error creating contact", err)
		return nil, err
	}

	return created, nil
}

func (s *contactService) FindAll(ctx context.Context, ownerID string, page, limit int, filters domains.ContactFilters) ([]*domains.Contact, int64, error) {
	offset := (page - 1) * limit

	contacts, total, err := s.repo.FindAll(ctx, ownerID, limit, offset, filters)
	if err != nil {
		logger.Error("error finding all contacts", err)
		return nil, 0, err
	}

	return contacts, total, nil
}

func (s *contactService) findAndAuthorize(ctx context.Context, id, ownerID string) (*domains.Contact, error) {
	contact, err := s.repo.FindByID(ctx, id)
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

func (s *contactService) FindByID(ctx context.Context, id, ownerID string) (*domains.Contact, error) {
	return s.findAndAuthorize(ctx, id, ownerID)
}

func (s *contactService) Update(ctx context.Context, contact *domains.Contact) (*domains.Contact, error) {
	existing, err := s.findAndAuthorize(ctx, contact.ID, contact.OwnerID)
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

	updated, err := s.repo.Update(ctx, existing)
	if err != nil {
		logger.Error("error updating contact", err)
		return nil, err
	}

	return updated, nil
}

func (s *contactService) Delete(ctx context.Context, id, ownerID string) error {
	if _, err := s.findAndAuthorize(ctx, id, ownerID); err != nil {
		return err
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		logger.Error("error deleting contact", err)
		return err
	}

	return nil
}

func (s *contactService) FindStats(ctx context.Context, ownerID string) (*domains.ContactStats, error) {
	stats, err := s.repo.FindStats(ctx, ownerID)
	if err != nil {
		logger.Error("error finding contact stats", err)
		return nil, err
	}

	return stats, nil
}

func (s *contactService) FindContactsByFrequency(ctx context.Context, userID string, limit int) ([]domains.ContactFrequency, error) {
	if limit <= 0 {
		limit = 5
	}

	contacts, err := s.repo.FindContactsByFrequency(ctx, userID, limit)
	if err != nil {
		logger.Error("error finding top contacts", err)
		return nil, err
	}

	return contacts, nil
}
