package user

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	RoleID   uint   `json:"role_id"`
	Status   string `json:"status"`
}

type UserListResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	RoleName    string `json:"role_name"`
	Status      string `json:"status"`
	LastLoginAt string `json:"last_login_at"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
