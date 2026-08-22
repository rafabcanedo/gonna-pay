package testutil

import (
	"time"

	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
)

func NewUserFixture() *domains.User {
	return domains.NewUserWithID("user-1", "Rafael", "rafael@email.com", "", "11999999999")
}

func NewContactFixture() *domains.Contact {
	return domains.NewContactWithID("contact-1", "user-1", "Ana", "ana@email.com", "11999999999", "Friend")
}

func NewGroupFixture() *domains.Group {
	now := time.Now()
	return domains.NewGroupWithID("group-1", "user-1", "Viagem", "Travel", now, now, nil)
}

func NewCostFixture() *domains.Cost {
	now := time.Now()
	splits := []domains.Split{
		{ID: "split-1", ContactID: "contact-1", ContactName: "Ana", Value: 50.0, Percentage: 50.0},
	}
	return domains.NewCostWithID("cost-1", "user-1", "group-1", "Grupo A", "Jantar", "Dinner", 100.0, 50.0, now, now, splits)
}
