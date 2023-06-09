package user

// UserCreatePayload for register User
type UserCreatePayload struct {
	Name        string `json:"name" binding:"required" validate:"min=3,max=100"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required" validate:"email=true,min=8,max=50"`
	PhoneNumber string `json:"phone_number" binding:"required" validate:"min=8,max=20"`
	RoleID      uint   `json:"role" binding:"required"`
}
