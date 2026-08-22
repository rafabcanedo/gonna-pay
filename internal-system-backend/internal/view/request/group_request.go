package request

type CreateGroupRequest struct {
	Name      string   `json:"name"      binding:"required"`
	Category  string   `json:"category"  binding:"required,oneof=Dinner Lunch Entertainment Travel Others"`
	MemberIDs []string `json:"memberIds"`
}

type UpdateGroupRequest struct {
	Name     string `json:"name"`
	Category string `json:"category" binding:"omitempty,oneof=Dinner Lunch Entertainment Travel Others"`
}
