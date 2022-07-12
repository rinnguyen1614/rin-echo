package response

import (
	"rin-echo/system/domain"
	"time"

	"github.com/rinnguyen1614/rin-echo-core/model"
	"github.com/rinnguyen1614/rin-echo-core/utils"
)

type Profile struct {
	model.Model

	UUID          utils.UUID
	Email         string     `json:"email"`
	EmailVerified bool       `json:"email_verified"`
	Phone         string     `json:"phone"`
	PhoneVerified bool       `json:"phone_verified"`
	Username      string     `json:"username"`
	FullName      string     `json:"full_name"`
	AvatarPath    string     `json:"avatar_path"`
	DateOfBirth   *time.Time `json:"date_of_birth"`
	Gender        uint       `json:"gender"`
}

func NewProfile(user domain.User) Profile {
	return Profile{
		Model:         model.NewModel(user.ID),
		UUID:          user.UUID,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		Username:      user.Username,
		FullName:      user.FullName,
		AvatarPath:    user.AvatarPath,
		DateOfBirth:   user.DateOfBirth,
		Phone:         user.Phone,
		PhoneVerified: user.PhoneVerified,
		Gender:        user.Gender,
	}
}

type UserMenu struct {
	model.Model

	Name      string    `json:"name" `
	Slug      string    `json:"slug" `
	ParentID  uint      `json:"parent_id" `
	Path      string    `json:"path" `
	Component string    `json:"component"`
	Sort      int       `json:"sort"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Icon      string    `json:"icon"`
	Children  UserMenus `json:"children"`
}

func NewUserMenu(e domain.Menu) UserMenu {
	return UserMenu{
		Model:     model.NewModel(e.ID),
		Name:      e.Name,
		Slug:      e.Slug,
		Path:      e.Path,
		ParentID:  *e.ParentID,
		Component: e.Component,
		Sort:      e.Sort,
		Type:      e.Type,
		Title:     e.Title,
		Icon:      e.Icon,
	}
}

type UserMenus []UserMenu

type UserPermission struct {
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}

type UserPermissions []*UserPermission

func NewUserPermissions(fields []map[string]interface{}) UserPermissions {
	var (
		mapByParentSlug = make(map[string]*UserPermission)
		result          UserPermissions
	)

	for _, field := range fields {
		parentSlug := field["parent_slug"].(string)
		m, ok := mapByParentSlug[parentSlug]
		if !ok {
			m = &UserPermission{
				Name:    parentSlug,
				Actions: make([]string, 0),
			}
			result = append(result, m)
			mapByParentSlug[parentSlug] = m
		}
		m.Actions = append(m.Actions, field["slug"].(string))
	}

	return result
}
