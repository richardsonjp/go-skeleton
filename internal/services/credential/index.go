package credential

import (
	"context"
	repos "go/skeleton/internal/repositories"
)

type CredentialService interface {
	CreateCredential(ctx context.Context, userID uint) (*CredentialResponse, error)
	InactiveCredential(ctx context.Context, secret string) error
	GetSecret(ctx context.Context, secret string) (*CredentialResponse, error)
}

type credentialService struct {
	credentialRepo repos.CredentialRepo
}

func NewCredentialService(credentialRepo repos.CredentialRepo) CredentialService {
	return &credentialService{
		credentialRepo: credentialRepo,
	}
}
