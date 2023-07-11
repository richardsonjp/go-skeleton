package authentication

import (
	"context"
	"go/skeleton/internal/model/enum"
	"go/skeleton/internal/services/credential"
	"go/skeleton/pkg/utils/errors"
	timeutil "go/skeleton/pkg/utils/time"
	"golang.org/x/crypto/bcrypt"
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
		credential, err := s.credentialService.CreateCredential(ctx, user.ID)
		if err != nil {
			return err
		}

		err = s.userService.UpdateLastLogin(ctx, user.ID)
		if err != nil {
			return err
		}

		branch, err := s.branchService.GetByID(ctx, user.BranchID)
		if err != nil {
			return err
		}
		response.Transformer(*user, *credential, branch.Name)

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

	branch, err := s.branchService.GetByID(ctx, user.BranchID)
	if err != nil {
		return nil, returnResponseErrorUnathorized()
	}

	session := &SessionData{}
	session.Transformer(*user, *credentialD, *branch)

	return session, nil
}

func (s *authenticationService) AuthenticateLogout(ctx context.Context, session SessionData) error {
	err := s.credentialService.InactiveCredential(ctx, session.Credential.Secret)
	return err
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
	if expiredAt.Before(timeutil.Now()) {
		return nil, errors.NewGenericError(errors.TOKEN_UNAUTHORIZED)
	}

	return credential, nil
}
