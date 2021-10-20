package repository

import (
	"rin-echo/admin/domain"
	"rin-echo/common/uow"
	iuow "rin-echo/common/uow/interfaces"

	"gorm.io/gorm"
)

type UserRepository struct {
	iuow.RepositoryOfEntity
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		RepositoryOfEntity: uow.NewRepositoryOfEntity(db, &domain.User{}),
	}
}

func (r UserRepository) UpdatePassword(user *domain.User, pwd string) error {
	user.SetPassword(pwd)
	return r.UpdateWithoutHooksWithPrimaryKey(user.ID, map[string]interface{}{"password": user.Password})
}

func (r UserRepository) FindByUsernameOrEmail(usernameOrEmail string, preloads map[string][]interface{}) (*domain.User, error) {
	var user domain.User

	if err := r.First(&user, map[string][]interface{}{"username": {usernameOrEmail}}, preloads); err != nil {
		return nil, err
	}

	return &user, nil
}
