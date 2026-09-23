package service

import (
	"errors"
	"math"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/configuration/logger"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository"
)

type CostService interface {
	Create(cost *domains.Cost, ownerPercentage *float64) (*domains.Cost, error)
	Update(id, userID string, cost *domains.Cost, ownerPercentage *float64) (*domains.Cost, error)
	FindAll(userID string, page, limit int, filters domains.CostFilters) ([]*domains.Cost, int64, error)
	FindByID(id, userID string) (*domains.Cost, error)
	Delete(id, userID string) error
	FindStats(userID string, filters domains.CostFilters) (*domains.CostStats, error)
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
		group, err := s.repo.GetGroupByID(cost.GroupID)
		if err != nil {
			if !errors.Is(err, domains.ErrNotFound) {
				logger.Error("error fetching group for cost creation", err)
			}
			return nil, err
		}
		if group.OwnerID != cost.UserID {
			return nil, domains.NewForbiddenError("access denied")
		}

		groupMembers, err := s.repo.GetGroupMembers(cost.GroupID)
		if err != nil {
			logger.Error("error fetching group members for cost creation", err)
			return nil, err
		}
		members = groupMembers

		memberCount := len(members)
		if ownerPercentage != nil {
			if *ownerPercentage <= 0 || *ownerPercentage >= 100 {
				return nil, domains.NewInvalidInputError("ownerPercentage must be between 0 and 100 (exclusive)")
			}
			cost.OwnerPercentage = *ownerPercentage
		} else {
			cost.OwnerPercentage = math.Round((100.0/float64(memberCount+1))*100) / 100
		}

		memberPercentage := math.Round(((100-cost.OwnerPercentage)/float64(len(members)))*100) / 100
		memberValue := math.Round((cost.TotalValue*memberPercentage/100)*100) / 100

		cost.Splits = make([]domains.Split, len(members))
		for i, m := range members {
			cost.Splits[i] = domains.Split{
				ContactID:   m.ID,
				ContactName: m.Name,
				Value:       memberValue,
				Percentage:  memberPercentage,
			}
		}
	} else {
		cost.OwnerPercentage = 100
	}

	created, err := s.repo.Create(cost)
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

		if len(existing.Splits) > 0 {
			memberPercentage := math.Round(((100-existing.OwnerPercentage)/float64(len(existing.Splits)))*100) / 100
			memberValue := math.Round((existing.TotalValue*memberPercentage/100)*100) / 100
			for i := range existing.Splits {
				existing.Splits[i].Percentage = memberPercentage
				existing.Splits[i].Value = memberValue
			}
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

func (s *costService) FindAll(userID string, page, limit int, filters domains.CostFilters) ([]*domains.Cost, int64, error) {
	offset := (page - 1) * limit

	costs, total, err := s.repo.FindAll(userID, limit, offset, filters)
	if err != nil {
		logger.Error("error finding all costs", err)
		return nil, 0, err
	}

	return costs, total, nil
}

func (s *costService) FindByID(id, userID string) (*domains.Cost, error) {
	return s.findAndAuthorize(id, userID)
}

func (s *costService) FindStats(userID string, filters domains.CostFilters) (*domains.CostStats, error) {
	stats, err := s.repo.FindStats(userID, filters)
	if err != nil {
		logger.Error("error finding cost stats", err)
		return nil, err
	}

	return stats, nil
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
