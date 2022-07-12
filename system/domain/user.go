package domain

import (
	"rin-echo/system/errors"
	"time"

	"github.com/rinnguyen1614/rin-echo-core/domain"
	iuow "github.com/rinnguyen1614/rin-echo-core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo-core/utils"

	"github.com/thoas/go-funk"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	Genders       = []uint{1, 2, 3}
	GenderDefault = Genders[0]
)

type User struct {
	domain.FullAuditedEntity

	UUID                        utils.UUID
	Username                    string `gorm:"unique;<-:create"`
	Password                    string
	FullName                    string
	AvatarPath                  string
	Email                       string `gorm:"unique"`
	EmailVerified               bool
	EmailVerificationCodeHashed string
	DateOfBirth                 *time.Time
	Phone                       string
	PhoneVerified               bool
	PhoneVerificationCodeHashed string
	Gender                      uint `gorm:"column:gender;size:1;default:1"`
	// ProviderId  string
	// Disabled      bool

	UserRoles UserRoles
	Addresses Addresses
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
	u.SetGenderDefault()
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

func (u *User) SetGenderDefault() {
	u.Gender = GenderDefault
}

func (u *User) SetGender(gender uint) error {
	if !funk.Contains(Genders, gender) {
		return errors.ErrGenderNotFound
	}
	u.Gender = gender
	return nil
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

	UpdateProfile(id uint, fullName string, dateOfBirth *time.Time, gender uint) error

	ChangePhone(id uint, phone string, phoneVerificationCodeHashed string) error

	VerifyPhone(id uint) error

	ChangeEmail(id uint, email string, emailVerificationCodeHashed string) error

	VerifyEmail(id uint) error

	FirstByUsernameOrEmail(usernameOrEmail string, preloads map[string][]interface{}) (*User, error)

	WithTransaction(db *gorm.DB) UserRepository
}
