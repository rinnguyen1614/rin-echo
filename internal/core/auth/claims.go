package auth

import (
	"github.com/rinnguyen1614/rin-echo/internal/core/utils"
)

type (
	ClaimsSession struct {
		ID       uint       `json:"id"`
		UUID     utils.UUID `json:"uuid"`
		Username string     `json:"username"`
		FullName string     `json:"fullname"`
	}
)

func (c ClaimsSession) UserID() uint {
	return c.ID
}

func (c ClaimsSession) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"ID":       c.ID,
		"UUID":     c.UUID,
		"Username": c.Username,
		"FullName": c.FullName,
	}
}
