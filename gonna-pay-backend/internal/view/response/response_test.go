package response_test

import (
	"testing"
	"time"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/view/response"
	"github.com/stretchr/testify/assert"
)

func TestNewCostResponse(t *testing.T) {
	now := time.Now()

	t.Run("OwnerValue is calculated correctly", func(t *testing.T) {
		cost := &domains.Cost{
			ID: "cost-1", CostName: "Jantar", TotalValue: 200.0,
			OwnerPercentage: 50.0, Category: "Dinner", CreatedAt: now,
		}

		r := response.NewCostResponse(cost)

		assert.Equal(t, 100.0, r.OwnerValue)
	})

	t.Run("SplitCount is included", func(t *testing.T) {
		cost := &domains.Cost{
			ID: "cost-1", TotalValue: 100.0, OwnerPercentage: 100.0,
			SplitCount: 3, CreatedAt: now,
		}

		r := response.NewCostResponse(cost)

		assert.Equal(t, 3, r.SplitCount)
	})

	t.Run("Splits field is not present in list response", func(t *testing.T) {
		cost := &domains.Cost{
			ID: "cost-1", TotalValue: 100.0, OwnerPercentage: 100.0,
			Splits: []domains.Split{{ID: "split-1"}}, CreatedAt: now,
		}

		r := response.NewCostResponse(cost)

		assert.Equal(t, "cost-1", r.ID)
	})
}

func TestNewCostDetailResponse(t *testing.T) {
	now := time.Now()

	t.Run("Splits are mapped correctly", func(t *testing.T) {
		cost := &domains.Cost{
			ID: "cost-1", TotalValue: 200.0, OwnerPercentage: 50.0,
			CreatedAt: now, UpdatedAt: now,
			Splits: []domains.Split{
				{ID: "split-1", ContactID: "contact-1", ContactName: "Ana", Value: 100.0, Percentage: 50.0},
			},
		}

		r := response.NewCostDetailResponse(cost)

		assert.Len(t, r.Splits, 1)
		assert.Equal(t, "split-1", r.Splits[0].ID)
		assert.Equal(t, "Ana", r.Splits[0].ContactName)
		assert.Equal(t, 100.0, r.Splits[0].Value)
		assert.Equal(t, 50.0, r.Splits[0].Percentage)
	})

	t.Run("OwnerValue is calculated correctly", func(t *testing.T) {
		cost := &domains.Cost{
			ID: "cost-1", TotalValue: 300.0, OwnerPercentage: 33.33,
			CreatedAt: now, UpdatedAt: now,
		}

		r := response.NewCostDetailResponse(cost)

		assert.InDelta(t, 99.99, r.OwnerValue, 0.01)
	})

	t.Run("UpdatedAt is included", func(t *testing.T) {
		cost := &domains.Cost{
			ID: "cost-1", TotalValue: 100.0, OwnerPercentage: 100.0,
			CreatedAt: now, UpdatedAt: now,
		}

		r := response.NewCostDetailResponse(cost)

		assert.Equal(t, now, r.UpdatedAt)
	})

	t.Run("empty splits returns empty slice", func(t *testing.T) {
		cost := &domains.Cost{
			ID: "cost-1", TotalValue: 100.0, OwnerPercentage: 100.0,
			CreatedAt: now, UpdatedAt: now, Splits: nil,
		}

		r := response.NewCostDetailResponse(cost)

		assert.Empty(t, r.Splits)
	})
}

func TestNewCostResponseList(t *testing.T) {
	now := time.Now()

	t.Run("empty list returns empty slice", func(t *testing.T) {
		r := response.NewCostResponseList([]*domains.Cost{})

		assert.NotNil(t, r)
		assert.Empty(t, r)
	})

	t.Run("all items are mapped", func(t *testing.T) {
		costs := []*domains.Cost{
			{ID: "cost-1", TotalValue: 100.0, OwnerPercentage: 100.0, CreatedAt: now},
			{ID: "cost-2", TotalValue: 200.0, OwnerPercentage: 50.0, CreatedAt: now},
		}

		r := response.NewCostResponseList(costs)

		assert.Len(t, r, 2)
		assert.Equal(t, "cost-1", r[0].ID)
		assert.Equal(t, "cost-2", r[1].ID)
	})
}

func TestNewGroupDetailResponse(t *testing.T) {
	now := time.Now()

	t.Run("Members are mapped correctly", func(t *testing.T) {
		group := domains.NewGroupWithID("group-1", "user-1", "Viagem", "Travel", now, now, []domains.Member{
			{ID: "contact-1", Name: "Ana", Email: "ana@email.com"},
			{ID: "contact-2", Name: "Pedro", Email: "pedro@email.com"},
		})

		r := response.NewGroupDetailResponse(group)

		assert.Len(t, r.Members, 2)
		assert.Equal(t, "contact-1", r.Members[0].ID)
		assert.Equal(t, "Ana", r.Members[0].Name)
		assert.Equal(t, "ana@email.com", r.Members[0].Email)
	})

	t.Run("empty members returns empty slice", func(t *testing.T) {
		group := domains.NewGroupWithID("group-1", "user-1", "Viagem", "Travel", now, now, nil)

		r := response.NewGroupDetailResponse(group)

		assert.NotNil(t, r.Members)
		assert.Empty(t, r.Members)
	})
}
