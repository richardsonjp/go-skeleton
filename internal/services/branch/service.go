package branch

import (
	"context"
	"github.com/google/uuid"
	"go/skeleton/internal/model"
	"go/skeleton/internal/model/enum"
	"go/skeleton/internal/services/constant"
	"go/skeleton/pkg/utils/errors"
)

func (s *branchService) CreateBranch(ctx context.Context, params BranchCreatePayload, roleName string) error {
	// this prevents user to create 'Branch' freely
	if roleName != constant.MASTER_ROLE {
		return errors.NewGenericError(errors.INVALID_ROLE_ACTION)
	}
	data := s.setData(params)
	err := s.branchRepo.Create(ctx, *data)
	if err != nil {
		return err
	}

	return nil
}

func (s *branchService) GetByID(ctx context.Context, ID uint) (*BranchResponse, error) {
	branch, err := s.branchRepo.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	response := &BranchResponse{}
	response.Transformer(*branch)

	return response, nil
}

func (s *branchService) setData(params BranchCreatePayload) *model.Branch {
	return &model.Branch{
		Name:     params.Name,
		Code:     params.Code,
		UniqueID: uuid.New().String(),
		Status:   enum.ACTIVE,
	}
}
