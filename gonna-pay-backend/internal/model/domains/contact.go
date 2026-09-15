package domains

import "time"

type Contact struct {
	ID        string    `json:"id"`
	OwnerID   string    `json:"ownerId"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdAt"`
}

type ContactFrequency struct {
	ContactID   string
	ContactName string
	SharedCosts int
}

func NewContact(ownerID, name, email, phone, category string) *Contact {
	return &Contact{
		OwnerID:  ownerID,
		Name:     name,
		Email:    email,
		Phone:    phone,
		Category: category,
	}
}

func NewContactWithID(id, ownerID, name, email, phone, category string, createdAt time.Time) *Contact {
	return &Contact{
		ID:        id,
		OwnerID:   ownerID,
		Name:      name,
		Email:     email,
		Phone:     phone,
		Category:  category,
		CreatedAt: createdAt,
	}
}
