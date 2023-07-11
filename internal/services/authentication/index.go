package authentication

import (
	"context"
	"go/skeleton/internal/repositories/tx"
	"go/skeleton/internal/services/branch"
	"go/skeleton/internal/services/credential"
	"go/skeleton/internal/services/user"
)

type AuthenticationService interface {
	AuthenticateSignin(ctx context.Context, params Signin) (*AuthenticateSessionResponse, error)
	AuthenticateLogout(ctx context.Context, session SessionData) error
	GetSession(ctx context.Context, secret string) (*SessionData, error)
}

type authenticationService struct {
	txRepo            tx.TxRepo
	userService       user.UserService
	credentialService credential.CredentialService
	branchService     branch.BranchService
}

func NewAuthenticationService(txRepo tx.TxRepo, userService user.UserService, credentialService credential.CredentialService, branchService branch.BranchService) AuthenticationService {
	return &authenticationService{
		txRepo:            txRepo,
		userService:       userService,
		credentialService: credentialService,
		branchService:     branchService,
	}
}
