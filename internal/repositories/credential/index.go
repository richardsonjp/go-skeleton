package credential

import (
	"context"
	"mgw/mgw-resi/internal/model"
	"mgw/mgw-resi/pkg/clients/db"
)

type CredentialRepo interface {
	Create(ctx context.Context, m model.Credential) error
	Update(ctx context.Context, m model.Credential, updatedFields ...string) (int64, error)
	GetSecret(ctx context.Context, secret string) (*model.Credential, error)
}

type credentialRepo struct {
	dbdget db.DBGormDelegate
}

func NewCredentialRepo(dbdget db.DBGormDelegate) CredentialRepo {
	return &credentialRepo{
		dbdget: dbdget,
	}
}