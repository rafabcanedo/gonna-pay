package service

import (
	"errors"

	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/configuration/logger"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/repository"
)

type GroupService interface {
	Create(domain domains.GroupDomainInterface, memberIDs []string) (domains.GroupDomainInterface, error)
	FindAll(ownerID string) ([]domains.GroupDomainInterface, error)
	FindByID(id, ownerID string) (domains.GroupDomainInterface, error)
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

func (s *groupService) Create(domain domains.GroupDomainInterface, memberIDs []string) (domains.GroupDomainInterface, error) {
	for _, contactID := range memberIDs {
		owned, err := s.repo.IsContactOwnedBy(contactID, domain.GetOwnerID())
		if err != nil {
			logger.Error("error checking contact ownership on group create", err)
			return nil, err
		}
		if !owned {
			return nil, domains.NewForbiddenError("one or more contacts do not belong to the authenticated user")
		}
	}

	created, err := s.repo.Create(domain, memberIDs)
	if err != nil {
		logger.Error("error creating group", err)
		return nil, err
	}

	return created, nil
}

func (s *groupService) FindAll(ownerID string) ([]domains.GroupDomainInterface, error) {
	groups, err := s.repo.FindAll(ownerID)
	if err != nil {
		logger.Error("error finding all groups", err)
		return nil, err
	}

	return groups, nil
}

func (s *groupService) FindByID(id, ownerID string) (domains.GroupDomainInterface, error) {
	group, err := s.repo.FindByID(id)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding group by id", err)
		}
		return nil, err
	}

	if group.GetOwnerID() != ownerID {
		return nil, domains.NewForbiddenError("access denied")
	}

	return group, nil
}

func (s *groupService) Delete(id, ownerID string) error {
	group, err := s.repo.FindByID(id)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding group for delete", err)
		}
		return err
	}

	if group.GetOwnerID() != ownerID {
		return domains.NewForbiddenError("access denied")
	}

	if err := s.repo.Delete(id); err != nil {
		logger.Error("error deleting group", err)
		return err
	}

	return nil
}

func (s *groupService) AddMember(groupID, contactID, ownerID string) error {
	group, err := s.repo.FindByID(groupID)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding group for add member", err)
		}
		return err
	}

	if group.GetOwnerID() != ownerID {
		return domains.NewForbiddenError("access denied")
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
	group, err := s.repo.FindByID(groupID)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding group for remove member", err)
		}
		return err
	}

	if group.GetOwnerID() != ownerID {
		return domains.NewForbiddenError("access denied")
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
