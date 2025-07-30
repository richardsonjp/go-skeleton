package authentication

import (
	"context"
	"go-skeleton/internal/repositories/tx"
	"go-skeleton/internal/services/context_path"
	"go-skeleton/internal/services/credential"
	"go-skeleton/internal/services/role"
	"go-skeleton/internal/services/user"
)

type AuthenticationService interface {
	AuthenticateSignin(ctx context.Context, params Signin) (*AuthenticateSessionResponse, error)
	AuthenticateLogout(ctx context.Context, session SessionData) error
	FrontendPath(ctx context.Context, session SessionData) (*FrontendPathResponse, error)
	GetSession(ctx context.Context, secret string) (*SessionData, error)
	ExtendSession(ctx context.Context, secret string) error
	GetHomeProfile(session SessionData) (*HomeProfile, error)
	GetListRBAC(ctx context.Context, roleID uint) ([]map[string]context_path.ModulePermission, error)
}

type authenticationService struct {
	txRepo             tx.TxRepo
	userService        user.UserService
	credentialService  credential.CredentialService
	contextPathService context_path.ContextPathService
	roleService        role.RoleService
}

func NewAuthenticationService(txRepo tx.TxRepo, userService user.UserService,
	credentialService credential.CredentialService,
	contextPathService context_path.ContextPathService,
	roleService role.RoleService) AuthenticationService {
	return &authenticationService{
		txRepo:             txRepo,
		userService:        userService,
		credentialService:  credentialService,
		contextPathService: contextPathService,
		roleService:        roleService,
	}
}
