package uow

import (
	"rin-echo/common"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/utils"

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

func (r RepositoryOfEntity) FindID(ctx common.Context, dest interface{}, ids []uint, preloads map[string][]interface{}) error {
	return r.Find(ctx, dest, map[string][]interface{}{"id": {ids}}, preloads)
}

func (r RepositoryOfEntity) FirstID(ctx common.Context, dest interface{}, id uint, preloads map[string][]interface{}) error {
	return r.First(ctx, dest, map[string][]interface{}{"id": {id}}, preloads)
}

func (r RepositoryOfEntity) FindUUID(ctx common.Context, dest interface{}, uuids []utils.UUID, preloads map[string][]interface{}) error {
	return r.Find(ctx, dest, map[string][]interface{}{"uuid": {uuids}}, preloads)
}

func (r RepositoryOfEntity) FirstUUID(ctx common.Context, dest interface{}, uuid utils.UUID, preloads map[string][]interface{}) error {
	return r.First(ctx, dest, map[string][]interface{}{"uuid": {uuid}}, preloads)
}
