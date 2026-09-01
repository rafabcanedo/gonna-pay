package response

import "github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func NewUserResponse(u *domains.User) UserResponse {
	return UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Phone: u.Phone,
	}
}

func NewUserResponseList(users []*domains.User) []UserResponse {
	out := make([]UserResponse, len(users))
	for i, u := range users {
		out[i] = NewUserResponse(u)
	}
	return out
}
