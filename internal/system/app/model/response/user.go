package response

import (
	"time"

	"github.com/rinnguyen1614/rin-echo/internal/core/model"
	"github.com/rinnguyen1614/rin-echo/internal/system/domain"
)

type User struct {
	model.FullAuditedEntityModel

	Email         string     `json:"email"`
	EmailVerified bool       `json:"email_verified"`
	Phone         string     `json:"phone"`
	PhoneVerified bool       `json:"phone_verified"`
	Username      string     `json:"username"`
	FullName      string     `json:"full_name"`
	AvatarPath    string     `json:"avatar_path"`
	DateOfBirth   *time.Time `json:"date_of_birth"`
	Gender        uint       `json:"gender"`

	UserRoles []UserRole `json:"user_roles"`
}

type Users []*User

type UserRole struct {
	model.FullAuditedEntityModel
	Role struct {
		model.Model
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"role"`
}

func NewUser(e domain.User) User {
	return User{
		FullAuditedEntityModel: model.NewFullAuditedModelWithEntity(e.FullAuditedEntity),
		Username:               e.Username,
		FullName:               e.FullName,
		Email:                  e.Email,
		AvatarPath:             e.AvatarPath,
	}
}
