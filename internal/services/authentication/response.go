package authentication

import (
	"go-skeleton/internal/services/context_path"
	"go-skeleton/internal/services/credential"
	"go-skeleton/internal/services/role"
	"go-skeleton/internal/services/user"
)

type AuthenticateSessionResponse struct {
	SID          string      `json:"sid"`
	ExpiredAt    string      `json:"expired_at"`
	Profile      UserProfile `json:"profile"`
	FrontendPath []string    `json:"frontend_path"`
}

type SessionData struct {
	User        user.UserResponse
	Role        role.RoleResponse
	Credential  credential.CredentialResponse
	BackendPath []context_path.BackendPath
}

type UserProfile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type FrontendPathResponse struct {
	FrontendPath []string `json:"frontend_path"`
}

type HomeProfile struct {
	User user.UserResponse `json:"user"`
	Role role.RoleResponse `json:"role"`
}
