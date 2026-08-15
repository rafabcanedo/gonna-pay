package domains

import "time"

type CostDomainInterface interface {
	GetID() string
	SetID(id string)

	GetUserID() string
	SetUserID(id string)

	GetGroupID() string
	SetGroupID(id string)

	GetGroupName() string
	SetGroupName(name string)

	GetSplitCount() int
	SetSplitCount(count int)

	GetCostName() string
	SetCostName(name string)

	GetTotalValue() float64
	SetTotalValue(value float64)

	GetOwnerPercentage() float64
	SetOwnerPercentage(percentage float64)

	GetCategory() string
	SetCategory(category string)

	GetCreatedAt() time.Time
	SetCreatedAt(t time.Time)

	GetUpdatedAt() time.Time
	SetUpdatedAt(t time.Time)

	GetSplits() []SplitDomain
	SetSplits(splits []SplitDomain)

	OwnerValue() float64
}
