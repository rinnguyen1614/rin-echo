package repository

import (
	"rin-echo/common/uow"
	"rin-echo/system/domain"
	"time"

	"gorm.io/gorm"
)

type (
	userRepository struct {
		*uow.RepositoryOfEntity
	}
)

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		RepositoryOfEntity: uow.NewRepositoryOfEntity(db, &domain.User{}),
	}
}

func (repo userRepository) UpdatePassword(user *domain.User, pwd string) error {
	user.SetPassword(pwd)
	return repo.UpdateWithoutHooksWithPrimaryKey(user.ID, map[string]interface{}{"password": user.Password})
}

func (repo userRepository) UpdateAvatar(id uint, path string) error {
	return repo.UpdateWithoutHooksWithPrimaryKey(id, map[string]interface{}{"avatar_path": path})
}

func (repo userRepository) UpdateProfile(id uint, fullName string, dateOfBirth *time.Time, gender uint) error {
	return repo.UpdateWithPrimaryKey(
		id,
		map[string]interface{}{
			"full_name":     fullName,
			"date_of_birth": dateOfBirth,
			"gender":        gender,
		})
}

func (repo userRepository) ChangePhone(id uint, phone string, phoneVerificationCodeHashed string) error {
	return repo.UpdateWithPrimaryKey(
		id,
		map[string]interface{}{
			"phone":                          phone,
			"phone_verified":                 false,
			"phone_verification_code_hashed": phoneVerificationCodeHashed,
		})
}

func (repo userRepository) VerifyPhone(id uint) error {
	return repo.UpdateWithPrimaryKey(
		id,
		map[string]interface{}{
			"phone_verified": true,
		})
}

func (repo userRepository) ChangeEmail(id uint, email string, emailVerificationCodeHashed string) error {
	return repo.UpdateWithPrimaryKey(
		id,
		map[string]interface{}{
			"email":                          email,
			"email_verified":                 false,
			"email_verification_code_hashed": emailVerificationCodeHashed,
		})
}

func (repo userRepository) VerifyEmail(id uint) error {
	return repo.UpdateWithPrimaryKey(
		id,
		map[string]interface{}{
			"email_verified": true,
		})
}

func (repo userRepository) FirstByUsernameOrEmail(usernameOrEmail string, preloads map[string][]interface{}) (*domain.User, error) {
	var user domain.User

	if err := repo.First(&user, map[string][]interface{}{"username": {usernameOrEmail}}, preloads); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) WithTransaction(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		RepositoryOfEntity: repo.RepositoryOfEntity.WithTransaction(db),
	}
}
