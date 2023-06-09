package credential

import (
	"context"
	"github.com/google/uuid"
	"mgw/mgw-resi/internal/model"
	"mgw/mgw-resi/internal/model/enum"
	timeutil "mgw/mgw-resi/pkg/utils/time"
)

func (s *credentialService) CreateCredential(ctx context.Context, userID uint) (*CredentialResponse, error) {
	data := s.setData(userID)
	err := s.credentialRepo.Create(ctx, *data)
	if err != nil {
		return nil, err
	}
	response := &CredentialResponse{}
	response.Transformer(*data)

	return response, nil
}

func (s *credentialService) GetSecret(ctx context.Context, secret string) (*CredentialResponse, error) {
	credential, err := s.credentialRepo.GetSecret(ctx, secret)
	if err != nil {
		return nil, err
	}
	response := &CredentialResponse{}
	response.Transformer(*credential)

	return response, nil
}

func (s *credentialService) InactiveCredential(ctx context.Context, secret string) error {
	credential, err := s.credentialRepo.GetSecret(ctx, secret)
	if err != nil {
		return err
	}

	credential.Status = enum.INACTIVE
	_, err = s.credentialRepo.Update(ctx, *credential, "status")
	if err != nil {
		return nil
	}

	return nil
}

func (s *credentialService) setData(userID uint) *model.Credential {
	return &model.Credential{
		UserID:    userID,
		Secret:    uuid.New().String(),
		Status:    enum.ACTIVE,
		ExpiredAt: timeutil.HoursAdd(timeutil.Now(), 1),
	}
}