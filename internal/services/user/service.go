package user

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/internal/model/enum"
	timeutil "go-skeleton/pkg/utils/time"
	"go-skeleton/pkg/utils/transformer"
	"go-skeleton/pkg/utils/wording"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) CreateUser(ctx context.Context, params UserCreatePayload) error {
	// create user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	params.Password = string(hashedPassword)

	data := s.setData(params)
	err = s.userRepo.Create(ctx, *data)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*UserResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, wording.NormalizeEmail(email))
	if err != nil {
		return nil, err
	}
	response := &UserResponse{}
	response.Transformer(*user)

	return response, nil
}

func (s *userService) UpdateLastLogin(ctx context.Context, userID uint) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}
	now := timeutil.Now()
	user.LastLoginAt = &now
	_, err = s.userRepo.Update(ctx, *user, "last_login_at")

	return nil
}

func (s *userService) GetByID(ctx context.Context, userID uint) (*UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	response := &UserResponse{}
	response.Transformer(*user)

	return response, nil
}

func (s *userService) GetListUser(ctx context.Context, filter UserGetFilterPayload) (*transformer.Pagination, error) {
	filterValue := make(map[string]string)

	if filter.Search != "" { //search
		filterValue["search"] = filter.Search
	}

	if filter.Status != "" {
		filterValue["status"] = filter.Status
	}

	data, list, err := s.userRepo.GetListUser(ctx, model.Pagination{
		Limit: filter.Pagination.Limit,
		Page:  filter.Pagination.Page,
		Sort:  filter.Pagination.Sort,
	}, filterValue)
	if err != nil {
		return nil, err
	}

	response := []UserListResponse{}
	for _, v := range data {
		var result UserListResponse
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

func (s *userService) setData(params UserCreatePayload) *model.User {
	newUser := &model.User{}
	newUser.Name = params.Name
	newUser.Password = params.Password
	newUser.Email = wording.NormalizeEmail(params.Email)
	newUser.Phone = params.Phone
	newUser.Status = enum.ACTIVE
	newUser.LastLoginAt = nil
	newUser.RoleID = params.RoleID
	return newUser
}
