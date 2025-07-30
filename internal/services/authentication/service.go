package authentication

import (
	"context"
	"go-skeleton/internal/model/enum"
	"go-skeleton/internal/services/context_path"
	"go-skeleton/internal/services/credential"
	"go-skeleton/pkg/utils/errors"
	timeutil "go-skeleton/pkg/utils/time"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (s *authenticationService) AuthenticateSignin(ctx context.Context, params Signin) (*AuthenticateSessionResponse, error) {
	user, err := s.userService.GetByEmail(ctx, params.Email)
	if err != nil {
		return nil, errors.NewGenericError(errors.INVALID_EMAIL_OR_PASSWORD)
	} else if user.Status != enum.ACTIVE.String() {
		return nil, errors.NewGenericError(errors.USER_REMOVED)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return nil, errors.NewGenericError(errors.INVALID_EMAIL_OR_PASSWORD)
	}

	response := &AuthenticateSessionResponse{}
	errTx := s.txRepo.Run(ctx, func(ctx context.Context) error {
		err = s.userService.UpdateLastLogin(ctx, user.ID)
		if err != nil {
			return err
		}

		credential, err := s.credentialService.CreateCredential(ctx, user.ID)
		if err != nil {
			return err
		}

		pathList, err := s.contextPathService.GetListPath(ctx, user.RoleID)
		if err != nil {
			return err
		}
		response.Transformer(*user, *credential, pathList.FrontendPath)

		return nil
	})
	if errTx != nil {
		return nil, errTx
	}

	return response, nil
}

func (s *authenticationService) GetSession(ctx context.Context, secret string) (*SessionData, error) {
	returnResponseErrorUnathorized := func() error {
		return errors.NewGenericError(errors.TOKEN_UNAUTHORIZED)
	}
	credentialD, err := s.validateCredential(ctx, secret)
	if err != nil {
		return nil, returnResponseErrorUnathorized()
	}

	user, err := s.userService.GetByID(ctx, credentialD.UserID)
	if err != nil {
		return nil, returnResponseErrorUnathorized()
	}

	role, err := s.roleService.GetByID(ctx, user.RoleID)
	if err != nil {
		return nil, returnResponseErrorUnathorized()
	}

	pathList, err := s.contextPathService.GetListPath(ctx, user.RoleID)
	if err != nil {
		return nil, err
	}

	session := &SessionData{}
	session.Transformer(*user, *credentialD, pathList.BackendPath, *role)

	return session, nil
}

func (s *authenticationService) ExtendSession(ctx context.Context, secret string) error {
	return s.credentialService.ExtendCredential(ctx, secret, 1*time.Hour)
}

func (s *authenticationService) AuthenticateLogout(ctx context.Context, session SessionData) error {
	err := s.credentialService.InactiveCredential(ctx, session.Credential.Secret)
	return err
}

func (s *authenticationService) FrontendPath(ctx context.Context, session SessionData) (*FrontendPathResponse, error) {
	path, err := s.contextPathService.GetListPath(ctx, session.User.RoleID)
	if err != nil {
		return nil, err
	}
	return &FrontendPathResponse{FrontendPath: path.FrontendPath}, nil
}

func (s *authenticationService) GetListRBAC(ctx context.Context, roleID uint) ([]map[string]context_path.ModulePermission, error) {
	path, err := s.contextPathService.GetListRBAC(ctx, roleID)
	if err != nil {
		return nil, err
	}

	return path, nil
}

func (s *authenticationService) validateCredential(ctx context.Context, secret string) (*credential.CredentialResponse, error) {
	credential, err := s.credentialService.GetSecret(ctx, secret)
	if err != nil {
		return nil, err
	}
	if credential.Status != enum.ACTIVE.String() {
		return nil, errors.NewGenericError(errors.TOKEN_UNAUTHORIZED)
	}
	expiredAt, err := timeutil.Parse(credential.ExpiredAt, timeutil.ISO8601TimeWithoutZone)
	if err != nil {
		return nil, err
	}
	now, err := timeutil.Parse(timeutil.NowStr(), timeutil.ISO8601TimeWithoutZone)
	if expiredAt.Before(now) {
		return nil, errors.NewGenericError(errors.TOKEN_UNAUTHORIZED)
	}

	return credential, nil
}

func (s *authenticationService) GetHomeProfile(session SessionData) (*HomeProfile, error) {
	return &HomeProfile{
		User: session.User,
		Role: session.Role,
	}, nil
}
