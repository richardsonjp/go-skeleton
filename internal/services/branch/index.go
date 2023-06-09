package branch

import (
	"context"
	repos "mgw/mgw-resi/internal/repositories"
)

type BranchService interface {
	CreateBranch(ctx context.Context, params BranchCreatePayload, roleName string) error
	GetByID(ctx context.Context, ID uint) (*BranchResponse, error)
}

type branchService struct {
	branchRepo repos.BranchRepo
}

func NewBranchService(branchRepo repos.BranchRepo) BranchService {
	return &branchService{
		branchRepo: branchRepo,
	}
}
