package service

import (
	"context"
	"errors"
	"math"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/configuration/logger"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository"
)

type CostService interface {
	Create(ctx context.Context, cost *domains.Cost, ownerPercentage *float64) (*domains.Cost, error)
	Update(ctx context.Context, id, userID string, cost *domains.Cost, ownerPercentage *float64) (*domains.Cost, error)
	FindAll(ctx context.Context, userID string, page, limit int, filters domains.CostFilters) ([]*domains.Cost, int64, error)
	FindByID(ctx context.Context, id, userID string) (*domains.Cost, error)
	Delete(ctx context.Context, id, userID string) error
	FindStats(ctx context.Context, userID string, filters domains.CostFilters) (*domains.CostStats, error)
}

type costService struct {
	repo      repository.CostRepository
	groupRepo repository.GroupRepository
}

func NewCostService(repo repository.CostRepository, groupRepo repository.GroupRepository) CostService {
	return &costService{repo: repo, groupRepo: groupRepo}
}

func (s *costService) Create(ctx context.Context, cost *domains.Cost, ownerPercentage *float64) (*domains.Cost, error) {
	var members []domains.Member

	if cost.GroupID != "" {
		group, err := s.groupRepo.FindByID(ctx, cost.GroupID)
		if err != nil {
			if !errors.Is(err, domains.ErrNotFound) {
				logger.Error("error fetching group for cost creation", err)
			}
			return nil, err
		}
		if group.OwnerID != cost.UserID {
			return nil, domains.NewForbiddenError("access denied")
		}

		groupMembers, err := s.groupRepo.GetMembers(ctx, cost.GroupID)
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

	created, err := s.repo.Create(ctx, cost)
	if err != nil {
		logger.Error("error creating cost", err)
		return nil, err
	}

	return created, nil
}

func (s *costService) findAndAuthorize(ctx context.Context, id, userID string) (*domains.Cost, error) {
	cost, err := s.repo.FindByID(ctx, id)
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

func (s *costService) Update(ctx context.Context, id, userID string, cost *domains.Cost, ownerPercentage *float64) (*domains.Cost, error) {
	existing, err := s.findAndAuthorize(ctx, id, userID)
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

	updated, err := s.repo.Update(ctx, id, existing)
	if err != nil {
		logger.Error("error updating cost", err)
		return nil, err
	}

	return updated, nil
}

func (s *costService) FindAll(ctx context.Context, userID string, page, limit int, filters domains.CostFilters) ([]*domains.Cost, int64, error) {
	offset := (page - 1) * limit

	costs, total, err := s.repo.FindAll(ctx, userID, limit, offset, filters)
	if err != nil {
		logger.Error("error finding all costs", err)
		return nil, 0, err
	}

	return costs, total, nil
}

func (s *costService) FindByID(ctx context.Context, id, userID string) (*domains.Cost, error) {
	return s.findAndAuthorize(ctx, id, userID)
}

func (s *costService) FindStats(ctx context.Context, userID string, filters domains.CostFilters) (*domains.CostStats, error) {
	stats, err := s.repo.FindStats(ctx, userID, filters)
	if err != nil {
		logger.Error("error finding cost stats", err)
		return nil, err
	}

	return stats, nil
}

func (s *costService) Delete(ctx context.Context, id, userID string) error {
	if _, err := s.findAndAuthorize(ctx, id, userID); err != nil {
		return err
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		logger.Error("error deleting cost", err)
		return err
	}

	return nil
}
