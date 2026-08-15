package response

import "github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"

type ContactResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Category string `json:"category"`
}

func NewContactResponse(d domains.ContactDomainInterface) ContactResponse {
	return ContactResponse{
		ID:       d.GetID(),
		Name:     d.GetName(),
		Email:    d.GetEmail(),
		Phone:    d.GetPhone(),
		Category: d.GetCategory(),
	}
}

func NewContactResponseList(contacts []domains.ContactDomainInterface) []ContactResponse {
	out := make([]ContactResponse, len(contacts))
	for i, c := range contacts {
		out[i] = NewContactResponse(c)
	}
	return out
}
