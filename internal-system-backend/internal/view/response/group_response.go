package response

import (
	"time"

	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
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

func NewGroupResponse(d domains.GroupDomainInterface) GroupResponse {
	return GroupResponse{
		ID:        d.GetID(),
		Name:      d.GetName(),
		Category:  d.GetCategory(),
		CreatedAt: d.GetCreatedAt(),
		UpdatedAt: d.GetUpdatedAt(),
	}
}

func NewGroupDetailResponse(d domains.GroupDomainInterface) GroupDetailResponse {
	members := make([]MemberResponse, len(d.GetMembers()))
	for i, m := range d.GetMembers() {
		members[i] = MemberResponse{ID: m.ID, Name: m.Name, Email: m.Email}
	}
	return GroupDetailResponse{
		ID:        d.GetID(),
		Name:      d.GetName(),
		Category:  d.GetCategory(),
		CreatedAt: d.GetCreatedAt(),
		UpdatedAt: d.GetUpdatedAt(),
		Members:   members,
	}
}

func NewGroupResponseList(groups []domains.GroupDomainInterface) []GroupResponse {
	out := make([]GroupResponse, len(groups))
	for i, g := range groups {
		out[i] = NewGroupResponse(g)
	}
	return out
}
