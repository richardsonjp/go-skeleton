package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"mgw/mgw-resi/internal/model"
	"mgw/mgw-resi/internal/model/enum"
	"mgw/mgw-resi/internal/services/constant"
	"mgw/mgw-resi/pkg/utils/errors"
	timeutil "mgw/mgw-resi/pkg/utils/time"
)

func (s *userService) CreateUser(ctx context.Context, accessorRoleID uint, params UserCreatePayload, branchID uint) error {
	// validate creator role
	role, err := s.roleService.GetByID(ctx, accessorRoleID)
	if err != nil {
		return err
	}
	if role.Name != constant.ADMIN_ROLE {
		return errors.NewGenericError(errors.INVALID_ROLE_ACTION)
	}

	// create user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	params.Password = string(hashedPassword)

	data := s.setData(params, branchID)
	err = s.userRepo.Create(ctx, *data)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*UserResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
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
	user.LastLogin = &now
	_, err = s.userRepo.Update(ctx, *user, "last_login")

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

func (s *userService) setData(params UserCreatePayload, branchID uint) *model.User {
	newUser := &model.User{}
	newUser.Name = params.Name
	newUser.Password = params.Password
	newUser.Email = params.Email
	newUser.PhoneNumber = params.PhoneNumber
	newUser.BranchID = branchID
	newUser.Status = enum.ACTIVE
	newUser.LastLogin = nil
	if params.RoleID != 0 {
		newUser.RoleID = params.RoleID
	}
	return newUser
}
