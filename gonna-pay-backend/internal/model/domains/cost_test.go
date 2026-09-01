package domains_test

import (
	"testing"
	"time"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/stretchr/testify/assert"
)

func TestCost_OwnerValue(t *testing.T) {
	t.Run("100% owner gets full value", func(t *testing.T) {
		cost := domains.NewCost("user-1", "", "Jantar", "Dinner", 200.0, 100)
		assert.Equal(t, 200.0, cost.OwnerValue())
	})

	t.Run("50% owner gets half value", func(t *testing.T) {
		cost := domains.NewCost("user-1", "group-1", "Jantar", "Dinner", 200.0, 50)
		assert.Equal(t, 100.0, cost.OwnerValue())
	})

	t.Run("0% owner gets zero", func(t *testing.T) {
		cost := domains.NewCost("user-1", "group-1", "Jantar", "Dinner", 200.0, 0)
		assert.Equal(t, 0.0, cost.OwnerValue())
	})
}

func TestNewCostWithID(t *testing.T) {
	now := time.Now()

	t.Run("SplitCount matches len(splits)", func(t *testing.T) {
		splits := []domains.Split{
			{ID: "split-1", ContactID: "contact-1", ContactName: "Ana", Value: 50.0, Percentage: 50.0},
			{ID: "split-2", ContactID: "contact-2", ContactName: "Pedro", Value: 50.0, Percentage: 50.0},
		}
		cost := domains.NewCostWithID("cost-1", "user-1", "group-1", "Grupo A", "Jantar", "Dinner", 100.0, 50.0, now, now, splits)

		assert.Equal(t, 2, cost.SplitCount)
		assert.Len(t, cost.Splits, 2)
	})

	t.Run("SplitCount is zero when no splits", func(t *testing.T) {
		cost := domains.NewCostWithID("cost-1", "user-1", "", "", "Jantar", "Dinner", 100.0, 100.0, now, now, nil)

		assert.Equal(t, 0, cost.SplitCount)
		assert.Empty(t, cost.Splits)
	})
}
