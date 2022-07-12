package repository

import (
	"rin-echo/system/domain"

	"github.com/rinnguyen1614/rin-echo-core/uow"

	"gorm.io/gorm"
)

type (
	userRoleRepository struct {
		*uow.RepositoryOfEntity
	}
)

func NewUserRoleRepository(db *gorm.DB) domain.UserRoleRepository {
	return &userRoleRepository{
		RepositoryOfEntity: uow.NewRepositoryOfEntity(db, &domain.UserRole{}),
	}
}

func (repo userRoleRepository) FindByUsers(userIDs []uint, preloads map[string][]interface{}) (domain.UserRoles, error) {
	var des domain.UserRoles

	if err := repo.Find(&des, map[string][]interface{}{"user_id": {userIDs}}, preloads); err != nil {
		return nil, err
	}

	return des, nil
}

func (repo userRoleRepository) FindByRoles(roleIDs []uint, preloads map[string][]interface{}) (domain.UserRoles, error) {
	var des domain.UserRoles

	if err := repo.Find(&des, map[string][]interface{}{"role_id": {roleIDs}}, preloads); err != nil {
		return nil, err
	}

	return des, nil
}

func (repo userRoleRepository) FindByUsersAndRoles(userIDs []uint, roleIDs []uint, preloads map[string][]interface{}) (domain.UserRoles, error) {
	var des domain.UserRoles

	if err := repo.Find(&des, map[string][]interface{}{"role_id": {roleIDs}, "user_id": {userIDs}}, preloads); err != nil {
		return nil, err
	}

	return des, nil
}

func (repo *userRoleRepository) WithTransaction(db *gorm.DB) domain.UserRoleRepository {
	return &userRoleRepository{
		RepositoryOfEntity: repo.RepositoryOfEntity.WithTransaction(db),
	}
}
