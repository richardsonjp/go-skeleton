package authentication

type Signin struct {
	Email    string `json:"email" binding:"required" validate:"email=true,min=8,max=50"`
	Password string `json:"password" binding:"required"`
}

type ListRBAC struct {
	RoleID uint `json:"role_id"`
}
