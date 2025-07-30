package role

import "go-skeleton/internal/model"

func (c *RoleResponse) Transformer(data model.Role) {
	c.ID = data.ID
	c.Name = data.Name
	c.Status = data.Status.String()
}
