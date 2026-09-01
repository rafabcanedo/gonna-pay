package domains

import "time"

type Member struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Group struct {
	ID        string    `json:"id"`
	OwnerID   string    `json:"ownerId"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Members   []Member  `json:"members"`
}

func NewGroup(ownerID, name, category string) *Group {
	return &Group{
		OwnerID:  ownerID,
		Name:     name,
		Category: category,
	}
}

func NewGroupForUpdate(id, ownerID, name, category string) *Group {
	return &Group{
		ID:       id,
		OwnerID:  ownerID,
		Name:     name,
		Category: category,
	}
}

func NewGroupWithID(id, ownerID, name, category string, createdAt, updatedAt time.Time, members []Member) *Group {
	return &Group{
		ID:        id,
		OwnerID:   ownerID,
		Name:      name,
		Category:  category,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Members:   members,
	}
}
