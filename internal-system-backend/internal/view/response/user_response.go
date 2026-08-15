package response

import "github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func NewUserResponse(d domains.UserDomainInterface) UserResponse {
	return UserResponse{
		ID:    d.GetID(),
		Name:  d.GetName(),
		Email: d.GetEmail(),
		Phone: d.GetPhone(),
	}
}

func NewUserResponseList(users []domains.UserDomainInterface) []UserResponse {
	out := make([]UserResponse, len(users))
	for i, u := range users {
		out[i] = NewUserResponse(u)
	}
	return out
}
