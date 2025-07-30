package authentication

import (
	"go-skeleton/internal/services/context_path"
	"go-skeleton/internal/services/credential"
	"go-skeleton/internal/services/role"
	"go-skeleton/internal/services/user"
)

func (c *AuthenticateSessionResponse) Transformer(user user.UserResponse, credential credential.CredentialResponse, frontendList []string) {
	c.SID = credential.Secret
	c.ExpiredAt = credential.ExpiredAt
	c.Profile.Name = user.Name
	c.Profile.Phone = user.Phone
	c.Profile.Email = user.Email
	c.FrontendPath = frontendList
}

func (c *SessionData) Transformer(user user.UserResponse, credential credential.CredentialResponse, backendList []context_path.BackendPath, role role.RoleResponse) {
	c.User = user
	c.Role = role
	c.Credential = credential
	c.BackendPath = backendList
}
