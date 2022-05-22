package domain

import (
	"rin-echo/common/domain"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/utils"
	"rin-echo/system/errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	domain.FullAuditedEntity

	UUID        utils.UUID
	Username    string `gorm:"unique;<-:create"`
	Password    string
	FullName    string
	AvatarPath  string
	Email       string `gorm:"unique"`
	DateOfBirth *time.Time
	//PhoneNumber string
	// PhotoURL    string
	// ProviderId  string

	// Disabled      bool
	// EmailVerified bool

	UserRoles UserRoles
}

func NewUser(username string, password string, fullName string, email string, roleIDs []uint) (*User, error) {
	u := &User{
		Username: username,
		FullName: fullName,
		Email:    email,
		UUID:     utils.MustUUID(),
	}

	if err := u.SetPassword(password); err != nil {
		return nil, err
	}

	u.AssignToRoles(roleIDs)

	return u, nil
}

func (u *User) SetPassword(pwd string) error {
	pwdh, err := HashPassword(pwd)
	if err != nil {
		return err
	}
	u.Password = pwdh
	return nil
}

func (u *User) CheckPassword(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	return err == nil
}

func (u *User) AssignToRole(roleID uint) {
	uR, _ := NewUserRole(u.ID, roleID)
	u.UserRoles = append(u.UserRoles, uR)
}

func (u *User) AssignToRoles(roleIDs []uint) {
	for _, roleID := range roleIDs {
		u.AssignToRole(roleID)
	}
}

func (u User) CompareUserRoles(newUserRoles UserRoles) (userRoleNews, userRoleDels UserRoles) {
	var (
		oldByRoleID = u.UserRoles.ToMapByRoleID()
		newByRoleID = newUserRoles.ToMapByRoleID()
	)

	if len(newUserRoles) != 0 {
		for rID, ur := range newByRoleID {
			_, ok := oldByRoleID[rID]
			if ok {
				delete(oldByRoleID, rID)
			} else {
				userRoleNews = append(userRoleNews, ur)
			}
		}

		for _, ur := range oldByRoleID {
			userRoleDels = append(userRoleDels, ur)
		}
	} else {
		userRoleDels = u.UserRoles
	}

	return
}

func HashPassword(pwd string) (string, error) {
	if len(pwd) == 0 {
		return "", errors.ErrPasswordRequired
	}

	h, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(h), err
}

type UserRepository interface {
	iuow.RepositoryOfEntity

	UpdatePassword(user *User, pwd string) error

	UpdateAvatar(id uint, path string) error

	UpdateProfile(id uint, fullName, email string, dateOfBirth *time.Time) error

	FirstByUsernameOrEmail(usernameOrEmail string, preloads map[string][]interface{}) (*User, error)

	WithTransaction(db *gorm.DB) UserRepository
}
