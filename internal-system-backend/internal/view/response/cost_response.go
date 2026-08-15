package response

import (
	"time"

	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
)

type SplitResponse struct {
	ID          string  `json:"id"`
	ContactID   string  `json:"contactId"`
	ContactName string  `json:"contactName"`
	Value       float64 `json:"value"`
	Percentage  float64 `json:"percentage"`
}

type CostResponse struct {
	ID              string    `json:"id"`
	CostName        string    `json:"costName"`
	TotalValue      float64   `json:"totalValue"`
	OwnerPercentage float64   `json:"ownerPercentage"`
	OwnerValue      float64   `json:"ownerValue"`
	Category        string    `json:"category"`
	GroupID         string    `json:"groupId"`
	GroupName       string    `json:"groupName"`
	SplitCount      int       `json:"splitCount"`
	CreatedAt       time.Time `json:"createdAt"`
}

type CostDetailResponse struct {
	ID              string          `json:"id"`
	CostName        string          `json:"costName"`
	TotalValue      float64         `json:"totalValue"`
	OwnerPercentage float64         `json:"ownerPercentage"`
	OwnerValue      float64         `json:"ownerValue"`
	Category        string          `json:"category"`
	GroupID         string          `json:"groupId"`
	GroupName       string          `json:"groupName"`
	SplitCount      int             `json:"splitCount"`
	CreatedAt       time.Time       `json:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt"`
	Splits          []SplitResponse `json:"splits"`
}

func NewCostResponse(c *domains.Cost) CostResponse {
	return CostResponse{
		ID:              c.ID,
		CostName:        c.CostName,
		TotalValue:      c.TotalValue,
		OwnerPercentage: c.OwnerPercentage,
		OwnerValue:      c.OwnerValue(),
		Category:        c.Category,
		GroupID:         c.GroupID,
		GroupName:       c.GroupName,
		SplitCount:      c.SplitCount,
		CreatedAt:       c.CreatedAt,
	}
}

func NewCostDetailResponse(c *domains.Cost) CostDetailResponse {
	splits := make([]SplitResponse, len(c.Splits))
	for i, s := range c.Splits {
		splits[i] = SplitResponse{
			ID:          s.ID,
			ContactID:   s.ContactID,
			ContactName: s.ContactName,
			Value:       s.Value,
			Percentage:  s.Percentage,
		}
	}
	return CostDetailResponse{
		ID:              c.ID,
		CostName:        c.CostName,
		TotalValue:      c.TotalValue,
		OwnerPercentage: c.OwnerPercentage,
		OwnerValue:      c.OwnerValue(),
		Category:        c.Category,
		GroupID:         c.GroupID,
		GroupName:       c.GroupName,
		SplitCount:      c.SplitCount,
		CreatedAt:       c.CreatedAt,
		UpdatedAt:       c.UpdatedAt,
		Splits:          splits,
	}
}

func NewCostResponseList(costs []*domains.Cost) []CostResponse {
	out := make([]CostResponse, len(costs))
	for i, c := range costs {
		out[i] = NewCostResponse(c)
	}
	return out
}
