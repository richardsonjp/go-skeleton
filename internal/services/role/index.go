package role

import (
	"context"
	repos "go-skeleton/internal/repositories"
	"go-skeleton/internal/repositories/tx"
	"go-skeleton/internal/services/context_path"
	"go-skeleton/pkg/utils/transformer"
)

type RoleService interface {
	CreateRole(ctx context.Context, payload RoleCreatePayload) error
	UpdateRole(ctx context.Context, payload RoleUpdatePayload) error
	GetByID(ctx context.Context, ID uint) (*RoleResponse, error)
	GetListRole(ctx context.Context, filter RoleGetFilterByPayload) (*transformer.Pagination, error)
}

type roleService struct {
	txRepo             tx.TxRepo
	roleRepo           repos.RoleRepo
	contextPathService context_path.ContextPathService
}

func NewRoleService(txRepo tx.TxRepo,
	roleRepo repos.RoleRepo,
	contextPathService context_path.ContextPathService) RoleService {
	return &roleService{
		txRepo:             txRepo,
		roleRepo:           roleRepo,
		contextPathService: contextPathService,
	}
}
