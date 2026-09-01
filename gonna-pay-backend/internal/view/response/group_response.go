package response

import (
	"time"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
)

type MemberResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GroupResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GroupDetailResponse struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Category  string           `json:"category"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	Members   []MemberResponse `json:"members"`
}

func NewGroupResponse(g *domains.Group) GroupResponse {
	return GroupResponse{
		ID:        g.ID,
		Name:      g.Name,
		Category:  g.Category,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}

func NewGroupDetailResponse(g *domains.Group) GroupDetailResponse {
	members := make([]MemberResponse, len(g.Members))
	for i, m := range g.Members {
		members[i] = MemberResponse{ID: m.ID, Name: m.Name, Email: m.Email}
	}
	return GroupDetailResponse{
		ID:        g.ID,
		Name:      g.Name,
		Category:  g.Category,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
		Members:   members,
	}
}

func NewGroupResponseList(groups []*domains.Group) []GroupResponse {
	out := make([]GroupResponse, len(groups))
	for i, g := range groups {
		out[i] = NewGroupResponse(g)
	}
	return out
}
