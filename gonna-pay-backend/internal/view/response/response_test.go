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

func TestNewUserResponse(t *testing.T) {
	t.Run("all fields are mapped", func(t *testing.T) {
		user := &domains.User{
			ID:    "user-1",
			Name:  "Rafael",
			Email: "rafael@email.com",
			Phone: "11999999999",
		}

		r := response.NewUserResponse(user)

		assert.Equal(t, "user-1", r.ID)
		assert.Equal(t, "Rafael", r.Name)
		assert.Equal(t, "rafael@email.com", r.Email)
		assert.Equal(t, "11999999999", r.Phone)
	})
}

func TestNewUserResponseList(t *testing.T) {
	t.Run("empty list returns empty slice", func(t *testing.T) {
		r := response.NewUserResponseList([]*domains.User{})

		assert.NotNil(t, r)
		assert.Empty(t, r)
	})

	t.Run("all items are mapped", func(t *testing.T) {
		users := []*domains.User{
			{ID: "user-1", Name: "Rafael", Email: "rafael@email.com", Phone: "11999999999"},
			{ID: "user-2", Name: "Ana", Email: "ana@email.com", Phone: "11888888888"},
		}

		r := response.NewUserResponseList(users)

		assert.Len(t, r, 2)
		assert.Equal(t, "user-1", r[0].ID)
		assert.Equal(t, "user-2", r[1].ID)
	})
}

func TestNewContactResponse(t *testing.T) {
	t.Run("all fields are mapped", func(t *testing.T) {
		contact := &domains.Contact{
			ID:       "contact-1",
			Name:     "Ana",
			Email:    "ana@email.com",
			Phone:    "11999999999",
			Category: "Friend",
		}

		r := response.NewContactResponse(contact)

		assert.Equal(t, "contact-1", r.ID)
		assert.Equal(t, "Ana", r.Name)
		assert.Equal(t, "ana@email.com", r.Email)
		assert.Equal(t, "11999999999", r.Phone)
		assert.Equal(t, "Friend", r.Category)
	})
}

func TestNewContactResponseList(t *testing.T) {
	t.Run("empty list returns empty slice", func(t *testing.T) {
		r := response.NewContactResponseList([]*domains.Contact{})

		assert.NotNil(t, r)
		assert.Empty(t, r)
	})

	t.Run("all items are mapped", func(t *testing.T) {
		contacts := []*domains.Contact{
			{ID: "contact-1", Name: "Ana", Email: "ana@email.com", Phone: "11999999999", Category: "Friend"},
			{ID: "contact-2", Name: "Pedro", Email: "pedro@email.com", Phone: "11888888888", Category: "Work"},
		}

		r := response.NewContactResponseList(contacts)

		assert.Len(t, r, 2)
		assert.Equal(t, "contact-1", r[0].ID)
		assert.Equal(t, "contact-2", r[1].ID)
	})
}

func TestNewContactFrequencyResponseList(t *testing.T) {
	t.Run("all fields are mapped", func(t *testing.T) {
		contacts := []domains.ContactFrequency{
			{ContactID: "contact-1", ContactName: "Ana", SharedCosts: 3},
		}

		r := response.NewContactFrequencyResponseList(contacts)

		assert.Len(t, r, 1)
		assert.Equal(t, "contact-1", r[0].ContactID)
		assert.Equal(t, "Ana", r[0].ContactName)
		assert.Equal(t, 3, r[0].SharedCosts)
	})

	t.Run("empty list returns empty slice", func(t *testing.T) {
		r := response.NewContactFrequencyResponseList([]domains.ContactFrequency{})

		assert.NotNil(t, r)
		assert.Empty(t, r)
	})

	t.Run("multiple items preserve order", func(t *testing.T) {
		contacts := []domains.ContactFrequency{
			{ContactID: "contact-1", ContactName: "Ana", SharedCosts: 5},
			{ContactID: "contact-2", ContactName: "Pedro", SharedCosts: 2},
		}

		r := response.NewContactFrequencyResponseList(contacts)

		assert.Len(t, r, 2)
		assert.Equal(t, "contact-1", r[0].ContactID)
		assert.Equal(t, 5, r[0].SharedCosts)
		assert.Equal(t, "contact-2", r[1].ContactID)
		assert.Equal(t, 2, r[1].SharedCosts)
	})
}
