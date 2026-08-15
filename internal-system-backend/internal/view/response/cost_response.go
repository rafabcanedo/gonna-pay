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

func NewCostResponse(d domains.CostDomainInterface) CostResponse {
	return CostResponse{
		ID:              d.GetID(),
		CostName:        d.GetCostName(),
		TotalValue:      d.GetTotalValue(),
		OwnerPercentage: d.GetOwnerPercentage(),
		OwnerValue:      d.OwnerValue(),
		Category:        d.GetCategory(),
		GroupID:         d.GetGroupID(),
		GroupName:       d.GetGroupName(),
		SplitCount:      d.GetSplitCount(),
		CreatedAt:       d.GetCreatedAt(),
	}
}

func NewCostDetailResponse(d domains.CostDomainInterface) CostDetailResponse {
	splits := make([]SplitResponse, len(d.GetSplits()))
	for i, s := range d.GetSplits() {
		splits[i] = SplitResponse{
			ID:          s.ID,
			ContactID:   s.ContactID,
			ContactName: s.ContactName,
			Value:       s.Value,
			Percentage:  s.Percentage,
		}
	}
	return CostDetailResponse{
		ID:              d.GetID(),
		CostName:        d.GetCostName(),
		TotalValue:      d.GetTotalValue(),
		OwnerPercentage: d.GetOwnerPercentage(),
		OwnerValue:      d.OwnerValue(),
		Category:        d.GetCategory(),
		GroupID:         d.GetGroupID(),
		GroupName:       d.GetGroupName(),
		SplitCount:      d.GetSplitCount(),
		CreatedAt:       d.GetCreatedAt(),
		UpdatedAt:       d.GetUpdatedAt(),
		Splits:          splits,
	}
}

func NewCostResponseList(costs []domains.CostDomainInterface) []CostResponse {
	out := make([]CostResponse, len(costs))
	for i, c := range costs {
		out[i] = NewCostResponse(c)
	}
	return out
}
