package context_path

import (
	"context"
	repos "go-skeleton/internal/repositories"
)

type ContextPathService interface {
	GetListPath(ctx context.Context, RoleID uint) (*AuthorizedPathResponse, error)
	GetListRBAC(ctx context.Context, roleID uint) ([]map[string]ModulePermission, error)
	CreateRBAC(ctx context.Context, roleID uint, rbacs []map[string]ModulePermission) error
}

type contextPathService struct {
	contextPathRepo repos.ContextPathRepo
}

func NewContextPathService(contextPathRepo repos.ContextPathRepo) ContextPathService {
	return &contextPathService{
		contextPathRepo: contextPathRepo,
	}
}
