package branch

import "go/skeleton/internal/model"

func (c *BranchResponse) Transformer(data model.Branch) {
	c.ID = data.ID
	c.Name = data.Name
	c.Code = data.UniqueID
	c.UniqueID = data.UniqueID
	c.Status = data.Status.String()
}
