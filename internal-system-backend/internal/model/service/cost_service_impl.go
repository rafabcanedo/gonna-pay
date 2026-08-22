package service

import (
	"errors"
	"math"

	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/configuration/logger"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/repository"
)

type CostService interface {
	Create(cost *domains.Cost, ownerPercentage *float64) (*domains.Cost, error)
	Update(id, userID string, cost *domains.Cost, ownerPercentage *float64) (*domains.Cost, error)
	FindAll(userID string) ([]*domains.Cost, error)
	FindByID(id, userID string) (*domains.Cost, error)
	Delete(id, userID string) error
}

type costService struct {
	repo repository.CostRepository
}

func NewCostService(repo repository.CostRepository) CostService {
	return &costService{repo: repo}
}

func (s *costService) Create(cost *domains.Cost, ownerPercentage *float64) (*domains.Cost, error) {
	var members []domains.Member

	if cost.GroupID != "" {
		group, err := s.repo.GetGroupMembers(cost.GroupID)
		if err != nil {
			logger.Error("error fetching group members for cost creation", err)
			return nil, err
		}
		members = group

		memberCount := len(members)
		if ownerPercentage != nil {
			if *ownerPercentage <= 0 || *ownerPercentage >= 100 {
				return nil, domains.NewInvalidInputError("ownerPercentage must be between 0 and 100 (exclusive)")
			}
			cost.OwnerPercentage = *ownerPercentage
		} else {
			cost.OwnerPercentage = math.Round((100.0/float64(memberCount+1))*100) / 100
		}
	} else {
		cost.OwnerPercentage = 100
	}

	created, err := s.repo.Create(cost, members)
	if err != nil {
		logger.Error("error creating cost", err)
		return nil, err
	}

	return created, nil
}

func (s *costService) findAndAuthorize(id, userID string) (*domains.Cost, error) {
	cost, err := s.repo.FindByID(id)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding cost", err)
		}
		return nil, err
	}

	if cost.UserID != userID {
		return nil, domains.NewForbiddenError("access denied")
	}

	return cost, nil
}

func (s *costService) Update(id, userID string, cost *domains.Cost, ownerPercentage *float64) (*domains.Cost, error) {
	existing, err := s.findAndAuthorize(id, userID)
	if err != nil {
		return nil, err
	}

	if cost.CostName != "" {
		existing.CostName = cost.CostName
	}
	if cost.TotalValue != 0 {
		existing.TotalValue = cost.TotalValue
	}
	if cost.Category != "" {
		existing.Category = cost.Category
	}

	if existing.GroupID != "" {
		if ownerPercentage != nil {
			if *ownerPercentage <= 0 || *ownerPercentage >= 100 {
				return nil, domains.NewInvalidInputError("ownerPercentage must be between 0 and 100 (exclusive)")
			}
			existing.OwnerPercentage = *ownerPercentage
		}
	} else {
		existing.OwnerPercentage = 100
	}

	updated, err := s.repo.Update(id, existing)
	if err != nil {
		logger.Error("error updating cost", err)
		return nil, err
	}

	return updated, nil
}

func (s *costService) FindAll(userID string) ([]*domains.Cost, error) {
	costs, err := s.repo.FindAll(userID)
	if err != nil {
		logger.Error("error finding all costs", err)
		return nil, err
	}

	return costs, nil
}

func (s *costService) FindByID(id, userID string) (*domains.Cost, error) {
	return s.findAndAuthorize(id, userID)
}

func (s *costService) Delete(id, userID string) error {
	if _, err := s.findAndAuthorize(id, userID); err != nil {
		return err
	}

	if err := s.repo.Delete(id); err != nil {
		logger.Error("error deleting cost", err)
		return err
	}

	return nil
}
