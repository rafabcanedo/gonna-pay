package domains

import "time"

type Split struct {
	ID          string  `json:"id"`
	ContactID   string  `json:"contactId"`
	ContactName string  `json:"contactName"`
	Value       float64 `json:"value"`
	Percentage  float64 `json:"percentage"`
}

type Cost struct {
	ID              string    `json:"id"`
	UserID          string    `json:"userId"`
	GroupID         string    `json:"groupId"`
	GroupName       string    `json:"groupName"`
	CostName        string    `json:"costName"`
	TotalValue      float64   `json:"totalValue"`
	OwnerPercentage float64   `json:"ownerPercentage"`
	Category        string    `json:"category"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	SplitCount      int       `json:"splitCount"`
	Splits          []Split   `json:"splits"`
}

func NewCost(userID, groupID, costName, category string, totalValue, ownerPercentage float64) *Cost {
	return &Cost{
		UserID:          userID,
		GroupID:         groupID,
		CostName:        costName,
		Category:        category,
		TotalValue:      totalValue,
		OwnerPercentage: ownerPercentage,
	}
}

func NewCostWithID(id, userID, groupID, groupName, costName, category string, totalValue, ownerPercentage float64, createdAt, updatedAt time.Time, splits []Split) *Cost {
	return &Cost{
		ID:              id,
		UserID:          userID,
		GroupID:         groupID,
		GroupName:       groupName,
		CostName:        costName,
		Category:        category,
		TotalValue:      totalValue,
		OwnerPercentage: ownerPercentage,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
		SplitCount:      len(splits),
		Splits:          splits,
	}
}

func (c *Cost) OwnerValue() float64 {
	return c.TotalValue * c.OwnerPercentage / 100
}

type CostCategoryBreakdown struct {
	Category   string
	Total      float64
	Percentage float64
}

type CostStats struct {
	ThisMonth  float64
	InSplits   float64
	Solo       float64
	ByCategory []CostCategoryBreakdown
}
