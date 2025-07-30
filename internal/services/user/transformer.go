package user

import (
	"go-skeleton/internal/model"
	"go-skeleton/internal/repositories/user"
)

func (c *UserResponse) Transformer(data model.User) {
	c.ID = data.ID
	c.Name = data.Name
	c.Password = data.Password
	c.Email = data.Email
	c.Phone = data.Phone
	c.RoleID = data.RoleID
	c.Status = data.Status.String()
}

func (c *UserListResponse) Transformer(data user.UserList) {
	c.ID = data.ID
	c.Name = data.Name
	c.Email = data.Email
	c.Phone = data.Phone
	c.RoleName = data.RoleName
	c.Status = data.Status
	c.LastLoginAt = data.LastLoginAt
	c.CreatedAt = data.CreatedAt
	c.UpdatedAt = data.UpdatedAt
}
