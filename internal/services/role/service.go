package role

import (
	"context"
	"mgw/mgw-resi/internal/model"
	"mgw/mgw-resi/internal/model/enum"
)

func (s *roleService) CreateRole(ctx context.Context, name string) error {
	data := s.setData(name)
	err := s.roleRepo.Create(ctx, *data)
	if err != nil {
		return err
	}

	return nil
}

func (s *roleService) GetByID(ctx context.Context, ID uint) (*RoleResponse, error) {
	role, err := s.roleRepo.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	response := &RoleResponse{}
	response.Transformer(*role)

	return response, nil
}

func (s *roleService) setData(name string) *model.Role {
	return &model.Role{
		Name:   name,
		Status: enum.ACTIVE,
	}
}
