package service

import (
	"errors"
	"math"

	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/configuration/logger"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/repository"
)

type CostService interface {
	Create(domain domains.CostDomainInterface, ownerPercentage *float64) (domains.CostDomainInterface, error)
	Update(id, userID string, domain domains.CostDomainInterface) (domains.CostDomainInterface, error)
	FindAll(userID string) ([]domains.CostDomainInterface, error)
	FindByID(id, userID string) (domains.CostDomainInterface, error)
	Delete(id, userID string) error
}

type costService struct {
	repo repository.CostRepository
}

func NewCostService(repo repository.CostRepository) CostService {
	return &costService{repo: repo}
}

func (s *costService) Create(domain domains.CostDomainInterface, ownerPercentage *float64) (domains.CostDomainInterface, error) {
	var memberIDs []string

	if domain.GetGroupID() != "" {
		group, err := s.repo.GetGroupMemberIDs(domain.GetGroupID())
		if err != nil {
			logger.Error("error fetching group members for cost creation", err)
			return nil, err
		}
		memberIDs = group

		memberCount := len(memberIDs)
		if ownerPercentage != nil {
			if *ownerPercentage <= 0 || *ownerPercentage >= 100 {
				return nil, domains.NewInvalidInputError("ownerPercentage must be between 0 and 100 (exclusive)")
			}
			domain.SetOwnerPercentage(*ownerPercentage)
		} else {
			domain.SetOwnerPercentage(math.Round((100.0/float64(memberCount+1))*100) / 100)
		}
	} else {
		domain.SetOwnerPercentage(100)
	}

	created, err := s.repo.Create(domain, memberIDs)
	if err != nil {
		logger.Error("error creating cost", err)
		return nil, err
	}

	return created, nil
}

func (s *costService) Update(id, userID string, domain domains.CostDomainInterface) (domains.CostDomainInterface, error) {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding cost for update", err)
		}
		return nil, err
	}

	if existing.GetUserID() != userID {
		return nil, domains.NewForbiddenError("access denied")
	}

	updated, err := s.repo.Update(id, domain)
	if err != nil {
		logger.Error("error updating cost", err)
		return nil, err
	}

	return updated, nil
}

func (s *costService) FindAll(userID string) ([]domains.CostDomainInterface, error) {
	costs, err := s.repo.FindAll(userID)
	if err != nil {
		logger.Error("error finding all costs", err)
		return nil, err
	}

	return costs, nil
}

func (s *costService) FindByID(id, userID string) (domains.CostDomainInterface, error) {
	cost, err := s.repo.FindByID(id)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding cost by id", err)
		}
		return nil, err
	}

	if cost.GetUserID() != userID {
		return nil, domains.NewForbiddenError("access denied")
	}

	return cost, nil
}

func (s *costService) Delete(id, userID string) error {
	cost, err := s.repo.FindByID(id)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding cost for delete", err)
		}
		return err
	}

	if cost.GetUserID() != userID {
		return domains.NewForbiddenError("access denied")
	}

	if err := s.repo.Delete(id); err != nil {
		logger.Error("error deleting cost", err)
		return err
	}

	return nil
}
