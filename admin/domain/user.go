package domain

import (
	"errors"
	"rin-echo/common/domain"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	domain.FullAuditedEntity
	Username string `gorm:"unique;<-:create"`
	Password string
	FullName string
	Email    *string `gorm:"unique"`
	// PhoneNumber string
	// PhotoURL    string
	// ProviderId  string

	// Disabled      bool
	// EmailVerified bool

	UserRoles []*UserRole
}

func NewUser(username string, password string, fullName string, email string) (User, error) {
	u := User{
		Username: username,
		FullName: fullName,
		Email:    &email,
	}

	if err := u.SetPassword(password); err != nil {
		return User{}, err
	}

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

func HashPassword(pwd string) (string, error) {
	if len(pwd) == 0 {
		return "", errors.New("password_not_empty")
	}

	h, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(h), err
}
