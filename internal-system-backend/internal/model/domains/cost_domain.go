package domains

import "time"

type SplitDomain struct {
	ID          string
	ContactID   string
	ContactName string
	Value       float64
	Percentage  float64
}

type costDomain struct {
	id              string
	userID          string
	groupID         string
	groupName       string
	costName        string
	totalValue      float64
	ownerPercentage float64
	category        string
	createdAt       time.Time
	updatedAt       time.Time
	splitCount      int
	splits          []SplitDomain
}

func NewCostDomain(userID, groupID, costName, category string, totalValue, ownerPercentage float64) CostDomainInterface {
	return &costDomain{
		userID:          userID,
		groupID:         groupID,
		costName:        costName,
		category:        category,
		totalValue:      totalValue,
		ownerPercentage: ownerPercentage,
	}
}

func NewCostDomainWithID(id, userID, groupID, costName, category string, totalValue, ownerPercentage float64, createdAt, updatedAt time.Time, splits []SplitDomain) CostDomainInterface {
	return &costDomain{
		id:              id,
		userID:          userID,
		groupID:         groupID,
		costName:        costName,
		category:        category,
		totalValue:      totalValue,
		ownerPercentage: ownerPercentage,
		createdAt:       createdAt,
		updatedAt:       updatedAt,
		splits:          splits,
	}
}

func (c *costDomain) GetID() string {
	return c.id
}

func (c *costDomain) SetID(id string) {
	c.id = id
}

func (c *costDomain) GetUserID() string {
	return c.userID
}

func (c *costDomain) SetUserID(id string) {
	c.userID = id
}

func (c *costDomain) GetGroupID() string {
	return c.groupID
}

func (c *costDomain) SetGroupID(id string) {
	c.groupID = id
}

func (c *costDomain) GetGroupName() string {
	return c.groupName
}

func (c *costDomain) SetGroupName(name string) {
	c.groupName = name
}

func (c *costDomain) GetSplitCount() int {
	return c.splitCount
}

func (c *costDomain) SetSplitCount(count int) {
	c.splitCount = count
}

func (c *costDomain) GetCostName() string {
	return c.costName
}

func (c *costDomain) SetCostName(name string) {
	c.costName = name
}

func (c *costDomain) GetTotalValue() float64 {
	return c.totalValue
}

func (c *costDomain) SetTotalValue(value float64) {
	c.totalValue = value
}

func (c *costDomain) GetOwnerPercentage() float64 {
	return c.ownerPercentage
}

func (c *costDomain) SetOwnerPercentage(percentage float64) {
	c.ownerPercentage = percentage
}

func (c *costDomain) GetCategory() string {
	return c.category
}

func (c *costDomain) SetCategory(category string) {
	c.category = category
}

func (c *costDomain) GetCreatedAt() time.Time {
	return c.createdAt
}

func (c *costDomain) SetCreatedAt(t time.Time) {
	c.createdAt = t
}

func (c *costDomain) GetUpdatedAt() time.Time {
	return c.updatedAt
}

func (c *costDomain) SetUpdatedAt(t time.Time) {
	c.updatedAt = t
}

func (c *costDomain) GetSplits() []SplitDomain {
	return c.splits
}

func (c *costDomain) SetSplits(splits []SplitDomain) {
	c.splits = splits
}

func (c *costDomain) OwnerValue() float64 {
	return c.totalValue * c.ownerPercentage / 100
}
