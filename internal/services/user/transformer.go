package user

import "mgw/mgw-resi/internal/model"

func (c *UserResponse) Transformer(data model.User) {
	c.ID = data.ID
	c.Name = data.Name
	c.Password = data.Password
	c.Email = data.Email
	c.PhoneNumber = data.PhoneNumber
	c.RoleID = data.RoleID
	c.BranchID = data.BranchID
	c.Status = data.Status.String()
}
