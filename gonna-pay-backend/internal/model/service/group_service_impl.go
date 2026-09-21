package service

import (
	"errors"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/configuration/logger"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository"
)

type GroupService interface {
	Create(group *domains.Group, memberIDs []string) (*domains.Group, error)
	FindAll(ownerID string, page, limit int, filters domains.GroupFilters) ([]*domains.Group, int64, error)
	FindByID(id, ownerID string) (*domains.Group, error)
	Update(group *domains.Group, ownerID string) (*domains.Group, error)
	Delete(id, ownerID string) error
	AddMember(groupID, contactID, ownerID string) error
	RemoveMember(groupID, contactID, ownerID string) error
}

type groupService struct {
	repo repository.GroupRepository
}

func NewGroupService(repo repository.GroupRepository) GroupService {
	return &groupService{repo: repo}
}

func (s *groupService) Create(group *domains.Group, memberIDs []string) (*domains.Group, error) {
	for _, contactID := range memberIDs {
		owned, err := s.repo.IsContactOwnedBy(contactID, group.OwnerID)
		if err != nil {
			logger.Error("error checking contact ownership on group create", err)
			return nil, err
		}
		if !owned {
			return nil, domains.NewForbiddenError("one or more contacts do not belong to the authenticated user")
		}
	}

	created, err := s.repo.Create(group, memberIDs)
	if err != nil {
		logger.Error("error creating group", err)
		return nil, err
	}

	return created, nil
}

func (s *groupService) FindAll(ownerID string, page, limit int, filters domains.GroupFilters) ([]*domains.Group, int64, error) {
	offset := (page - 1) * limit

	groups, total, err := s.repo.FindAll(ownerID, limit, offset, filters)
	if err != nil {
		logger.Error("error finding all groups", err)
		return nil, 0, err
	}

	return groups, total, nil
}

func (s *groupService) findAndAuthorize(id, ownerID string) (*domains.Group, error) {
	group, err := s.repo.FindByID(id)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding group", err)
		}
		return nil, err
	}

	if group.OwnerID != ownerID {
		return nil, domains.NewForbiddenError("access denied")
	}

	return group, nil
}

func (s *groupService) FindByID(id, ownerID string) (*domains.Group, error) {
	return s.findAndAuthorize(id, ownerID)
}

func (s *groupService) Update(group *domains.Group, ownerID string) (*domains.Group, error) {
	existing, err := s.findAndAuthorize(group.ID, ownerID)
	if err != nil {
		return nil, err
	}

	if group.Name != "" {
		existing.Name = group.Name
	}
	if group.Category != "" {
		existing.Category = group.Category
	}

	updated, err := s.repo.Update(existing)
	if err != nil {
		logger.Error("error updating group", err)
		return nil, err
	}

	return updated, nil
}

func (s *groupService) Delete(id, ownerID string) error {
	if _, err := s.findAndAuthorize(id, ownerID); err != nil {
		return err
	}

	if err := s.repo.Delete(id); err != nil {
		logger.Error("error deleting group", err)
		return err
	}

	return nil
}

func (s *groupService) AddMember(groupID, contactID, ownerID string) error {
	if _, err := s.findAndAuthorize(groupID, ownerID); err != nil {
		return err
	}

	owned, err := s.repo.IsContactOwnedBy(contactID, ownerID)
	if err != nil {
		logger.Error("error checking contact ownership", err)
		return err
	}
	if !owned {
		return domains.NewForbiddenError("contact does not belong to the authenticated user")
	}

	exists, err := s.repo.MemberExists(groupID, contactID)
	if err != nil {
		logger.Error("error checking member existence", err)
		return err
	}
	if exists {
		return domains.NewConflictError("contact is already a member of this group")
	}

	if err := s.repo.AddMember(groupID, contactID); err != nil {
		logger.Error("error adding member to group", err)
		return err
	}

	return nil
}

func (s *groupService) RemoveMember(groupID, contactID, ownerID string) error {
	if _, err := s.findAndAuthorize(groupID, ownerID); err != nil {
		return err
	}

	exists, err := s.repo.MemberExists(groupID, contactID)
	if err != nil {
		logger.Error("error checking member existence", err)
		return err
	}
	if !exists {
		return domains.NewNotFoundError("member not found in this group")
	}

	if err := s.repo.RemoveMember(groupID, contactID); err != nil {
		logger.Error("error removing member from group", err)
		return err
	}

	return nil
}
