package uow

import (
	iuow "rin-echo/common/uow/interfaces"

	"gorm.io/gorm"
)

type RepositoryOfEntity struct {
	iuow.Repository
}

func NewRepositoryOfEntity(store *gorm.DB, model interface{}) iuow.RepositoryOfEntity {
	return &RepositoryOfEntity{
		Repository: NewRepository(store, model),
	}
}

func (r RepositoryOfEntity) FindID(dest interface{}, ids []uint, preloads map[string][]interface{}) error {
	return r.Find(dest, map[string][]interface{}{"id": {ids}}, preloads)
}

func (r RepositoryOfEntity) GetID(dest interface{}, ids []uint, preloads map[string][]interface{}) error {
	return r.Get(dest, map[string][]interface{}{"id": {ids}}, preloads)
}

func (r RepositoryOfEntity) FirstID(dest interface{}, id uint, preloads map[string][]interface{}) error {
	return r.First(dest, map[string][]interface{}{"id": {id}}, preloads)
}
