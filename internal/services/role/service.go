package role

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/internal/model/enum"
	"go-skeleton/pkg/utils/transformer"
)

func (s *roleService) CreateRole(ctx context.Context, payload RoleCreatePayload) error {
	return s.txRepo.Run(ctx, func(ctx context.Context) error {
		data := s.setData(payload)
		result, err := s.roleRepo.Create(ctx, *data)
		if err != nil {
			return err
		}

		return s.contextPathService.CreateRBAC(ctx, result.ID, payload.RBAC)
	})
}

func (s *roleService) UpdateRole(ctx context.Context, payload RoleUpdatePayload) error {
	return s.txRepo.Run(ctx, func(ctx context.Context) error {
		data, err := s.roleRepo.GetByID(ctx, payload.ID)
		if err != nil {
			return err
		}

		var updatedFields []string
		if data.Name != payload.Name { // do this, because name is unique for this column
			data.Name = payload.Name
			updatedFields = append(updatedFields, "name")
		}

		if payload.Status != "" {
			data.Status = enum.NewGenericStatus(payload.Status)
			updatedFields = append(updatedFields, "status")
		}

		_, err = s.roleRepo.Update(ctx, *data, updatedFields...)
		if err != nil {
			return err
		}

		return s.contextPathService.CreateRBAC(ctx, data.ID, payload.RBAC)
	})
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

func (s *roleService) GetListRole(ctx context.Context, filter RoleGetFilterByPayload) (*transformer.Pagination, error) {
	data, list, err := s.roleRepo.GetListRole(ctx, model.Pagination{
		Limit: filter.Pagination.Limit,
		Page:  filter.Pagination.Page,
		Sort:  filter.Pagination.Sort,
	})
	if err != nil {
		return nil, err
	}

	response := []RoleResponse{}
	for _, v := range data {
		var result RoleResponse
		result.Transformer(*v)
		response = append(response, result)
	}

	return &transformer.Pagination{
		Data: response,
		Meta: transformer.Meta{
			Pagination: transformer.PaginationMeta{
				Limit:      list.Limit,
				Page:       list.Page,
				Sort:       list.Sort,
				TotalRows:  list.TotalRows,
				TotalPages: list.TotalPages,
			},
		},
	}, nil
}

func (s *roleService) setData(payload RoleCreatePayload) *model.Role {
	return &model.Role{
		Name:   payload.Name,
		Status: enum.NewGenericStatus(payload.Status),
	}
}
