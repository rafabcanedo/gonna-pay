package response

import (
	"time"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
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

type CostCategoryBreakdownResponse struct {
	Category   string  `json:"category"`
	Total      float64 `json:"total"`
	Percentage float64 `json:"percentage"`
}

type CostStatsResponse struct {
	ThisMonth  float64                          `json:"thisMonth"`
	InSplits   float64                          `json:"inSplits"`
	Solo       float64                          `json:"solo"`
	ByCategory []CostCategoryBreakdownResponse  `json:"byCategory"`
}

type CostsListResponse struct {
	Data       []CostResponse    `json:"data"`
	Page       int               `json:"page"`
	Limit      int               `json:"limit"`
	Total      int64             `json:"total"`
	TotalPages int               `json:"totalPages"`
	Stats      CostStatsResponse `json:"stats"`
}

func NewCostsListResponse(costs []CostResponse, page, limit int, total int64, stats *domains.CostStats) CostsListResponse {
	totalPages := int(total) / limit
	if int(total)%limit != 0 {
		totalPages++
	}

	breakdown := make([]CostCategoryBreakdownResponse, len(stats.ByCategory))
	for i, b := range stats.ByCategory {
		breakdown[i] = CostCategoryBreakdownResponse{
			Category:   b.Category,
			Total:      b.Total,
			Percentage: b.Percentage,
		}
	}

	return CostsListResponse{
		Data:       costs,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
		Stats: CostStatsResponse{
			ThisMonth:  stats.ThisMonth,
			InSplits:   stats.InSplits,
			Solo:       stats.Solo,
			ByCategory: breakdown,
		},
	}
}

func NewCostResponseList(costs []*domains.Cost) []CostResponse {
	out := make([]CostResponse, len(costs))
	for i, c := range costs {
		out[i] = NewCostResponse(c)
	}
	return out
}
