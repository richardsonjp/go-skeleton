package credential

import (
	"go-skeleton/internal/model"
	timeutil "go-skeleton/pkg/utils/time"
)

func (c *CredentialResponse) Transformer(data model.Credential) {
	c.ID = data.ID
	c.UserID = data.UserID
	c.Secret = data.Secret
	c.Status = data.Status.String()
	c.ExpiredAt = timeutil.StrFormat(data.ExpiredAt, timeutil.ISO8601TimeWithoutZone)
}
