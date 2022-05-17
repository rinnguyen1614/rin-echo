package repository

import (
	"rin-echo/common/uow"
	"rin-echo/system/domain"

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
