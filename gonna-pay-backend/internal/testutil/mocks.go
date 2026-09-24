package testutil

import (
	"time"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
)

type UserMock struct {
	User  *domains.User
	ID    string
	Name  string
	Email string
	Phone string
}

type ContactMock struct {
	Contact  *domains.Contact
	ID       string
	OwnerID  string
	Name     string
	Email    string
	Phone    string
	Category string
}

type GroupMock struct {
	Group    *domains.Group
	ID       string
	OwnerID  string
	Name     string
	Category string
}

type CostMock struct {
	Cost            *domains.Cost
	ID              string
	UserID          string
	GroupID         string
	Name            string
	Category        string
	TotalValue      float64
	OwnerPercentage float64
	SplitValue      float64
	SplitPercentage float64
}

func NewUserMock() UserMock {
	u := domains.NewUserWithID(UserID, UserName, UserEmail, "", UserPhone, false)
	return UserMock{
		User:  u,
		ID:    UserID,
		Name:  UserName,
		Email: UserEmail,
		Phone: UserPhone,
	}
}

func NewContactMock() ContactMock {
	c := domains.NewContactWithID(ContactID, UserID, ContactName, ContactEmail, ContactPhone, ContactCategory, time.Now())
	return ContactMock{
		Contact:  c,
		ID:       ContactID,
		OwnerID:  UserID,
		Name:     ContactName,
		Email:    ContactEmail,
		Phone:    ContactPhone,
		Category: ContactCategory,
	}
}

func NewGroupMock() GroupMock {
	now := time.Now()
	g := domains.NewGroupWithID(GroupID, UserID, GroupName, GroupCategory, now, now, nil)
	return GroupMock{
		Group:    g,
		ID:       GroupID,
		OwnerID:  UserID,
		Name:     GroupName,
		Category: GroupCategory,
	}
}

func NewCostMock() CostMock {
	const (
		totalValue = 100.0
		ownerPct   = 50.0
		splitValue = 50.0
		splitPct   = 50.0
	)
	now := time.Now()
	splits := []domains.Split{
		{ID: SplitID, ContactID: ContactID, ContactName: ContactName, Value: splitValue, Percentage: splitPct},
	}
	c := domains.NewCostWithID(CostID, UserID, GroupID, "Grupo A", CostName, CostCategory, totalValue, ownerPct, now, now, splits)
	return CostMock{
		Cost:            c,
		ID:              CostID,
		UserID:          UserID,
		GroupID:         GroupID,
		Name:            CostName,
		Category:        CostCategory,
		TotalValue:      totalValue,
		OwnerPercentage: ownerPct,
		SplitValue:      splitValue,
		SplitPercentage: splitPct,
	}
}
