package user

// UserCreatePayload for register User
type UserCreatePayload struct {
	Name     string `json:"name" binding:"required" validate:"min=3,max=100"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required" validate:"email=true,min=8,max=50"`
	Phone    string `json:"phone" binding:"required" validate:"min=8,max=20"`
	RoleID   uint   `json:"role" binding:"required"`
}

type UserGetFilterPayload struct {
	Pagination struct {
		Page  int    `form:"page" query:"page" validate:"omitempty,numeric"`
		Limit int    `form:"limit" query:"limit" validate:"omitempty,numeric"`
		Sort  string `form:"sort" query:"sort" validate:"omitempty"`
	}
	Search string `form:"search" query:"search"`
	Status string `form:"status" query:"status"`
}
