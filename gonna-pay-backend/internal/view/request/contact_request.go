package request

type CreateContactRequest struct {
	Name     string `json:"name"     binding:"required"`
	Email    string `json:"email"    binding:"required,email"`
	Phone    string `json:"phone"    binding:"required"`
	Category string `json:"category" binding:"required,oneof=Family Friend Work"`
}

type UpdateContactRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"    binding:"omitempty,email"`
	Phone    string `json:"phone"`
	Category string `json:"category" binding:"omitempty,oneof=Family Friend Work"`
}
