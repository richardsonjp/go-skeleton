package user

type UserResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Password    string `json:"-"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	RoleID      uint   `json:"role_id"`
	BranchID    uint   `json:"branch_id"`
	Status      string `json:"status"`
}
