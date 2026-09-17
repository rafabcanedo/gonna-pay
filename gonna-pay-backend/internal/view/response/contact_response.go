package response

import (
	"time"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
)

type ContactResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewContactResponse(c *domains.Contact) ContactResponse {
	return ContactResponse{
		ID:        c.ID,
		Name:      c.Name,
		Email:     c.Email,
		Phone:     c.Phone,
		Category:  c.Category,
		CreatedAt: c.CreatedAt,
	}
}

func NewContactResponseList(contacts []*domains.Contact) []ContactResponse {
	out := make([]ContactResponse, len(contacts))
	for i, c := range contacts {
		out[i] = NewContactResponse(c)
	}
	return out
}

type ContactStatsResponse struct {
	ByCategory map[string]int64 `json:"byCategory"`
}

type ContactsListResponse struct {
	Data       []ContactResponse    `json:"data"`
	Page       int                  `json:"page"`
	Limit      int                  `json:"limit"`
	Total      int64                `json:"total"`
	TotalPages int                  `json:"totalPages"`
	Stats      ContactStatsResponse `json:"stats"`
}

func NewContactsListResponse(contacts []ContactResponse, page, limit int, total int64, stats *domains.ContactStats) ContactsListResponse {
	totalPages := int(total) / limit
	if int(total)%limit != 0 {
		totalPages++
	}

	return ContactsListResponse{
		Data:       contacts,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
		Stats: ContactStatsResponse{
			ByCategory: stats.ByCategory,
		},
	}
}

type ContactFrequencyResponse struct {
	ContactID   string `json:"contactId"`
	ContactName string `json:"contactName"`
	SharedCosts int    `json:"sharedCosts"`
}

func NewContactFrequencyResponseList(contacts []domains.ContactFrequency) []ContactFrequencyResponse {
	out := make([]ContactFrequencyResponse, len(contacts))
	for i, c := range contacts {
		out[i] = ContactFrequencyResponse{
			ContactID:   c.ContactID,
			ContactName: c.ContactName,
			SharedCosts: c.SharedCosts,
		}
	}
	return out
}
