package authentication

import (
	"mgw/mgw-resi/internal/services/branch"
	"mgw/mgw-resi/internal/services/credential"
	"mgw/mgw-resi/internal/services/user"
)

type AuthenticateSessionResponse struct {
	SID       string      `json:"sid"`
	ExpiredAt string      `json:"expired_at"`
	Profile   UserProfile `json:"profile"`
}

type SessionData struct {
	User       user.UserResponse
	Credential credential.CredentialResponse
	Branch     branch.BranchResponse
}

type UserProfile struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	BranchName  string `json:"branch_name"`
}
