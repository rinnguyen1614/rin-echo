package uow

import (
	"gorm.io/gorm"
)

type RepositoryOfEntity struct {
	*Repository
}

func NewRepositoryOfEntity(store *gorm.DB, model interface{}) *RepositoryOfEntity {
	return &RepositoryOfEntity{
		Repository: NewRepository(store, model),
	}
}

func (r RepositoryOfEntity) FindID(dest interface{}, ids []uint, preloads map[string][]interface{}) error {
	return r.Find(dest, map[string][]interface{}{"id": {ids}}, preloads)
}

func (r RepositoryOfEntity) GetID(dest interface{}, ids uint, preloads map[string][]interface{}) error {
	return r.Get(dest, map[string][]interface{}{"id": {ids}}, preloads)
}

func (r RepositoryOfEntity) FirstID(dest interface{}, id uint, preloads map[string][]interface{}) error {
	return r.First(dest, map[string][]interface{}{"id": {id}}, preloads)
}

func (r RepositoryOfEntity) CountID(ids []uint) int64 {
	return r.Count(map[string][]interface{}{"id": {ids}})
}

func (r RepositoryOfEntity) ContainsID(ids []uint) bool {
	return r.Contains(map[string][]interface{}{"id": {ids}})
}

func (r *RepositoryOfEntity) WithTransaction(db *gorm.DB) *RepositoryOfEntity {
	return &RepositoryOfEntity{Repository: r.Repository.WithTransaction(db)}
}
