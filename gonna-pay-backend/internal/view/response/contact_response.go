package response

import "github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"

type ContactResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Category string `json:"category"`
}

func NewContactResponse(c *domains.Contact) ContactResponse {
	return ContactResponse{
		ID:       c.ID,
		Name:     c.Name,
		Email:    c.Email,
		Phone:    c.Phone,
		Category: c.Category,
	}
}

func NewContactResponseList(contacts []*domains.Contact) []ContactResponse {
	out := make([]ContactResponse, len(contacts))
	for i, c := range contacts {
		out[i] = NewContactResponse(c)
	}
	return out
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
