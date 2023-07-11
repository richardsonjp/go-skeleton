package authentication

import (
	"go/skeleton/internal/services/branch"
	"go/skeleton/internal/services/credential"
	"go/skeleton/internal/services/user"
)

func (c *AuthenticateSessionResponse) Transformer(user user.UserResponse, credential credential.CredentialResponse, branchName string) {
	c.SID = credential.Secret
	c.ExpiredAt = credential.ExpiredAt
	c.Profile.Name = user.Name
	c.Profile.PhoneNumber = user.PhoneNumber
	c.Profile.Email = user.Email
	c.Profile.BranchName = branchName
}

func (c *SessionData) Transformer(user user.UserResponse, credential credential.CredentialResponse, branch branch.BranchResponse) {
	c.User = user
	c.Credential = credential
	c.Branch = branch
}
