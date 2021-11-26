package model

import (
	"rin-echo/common/domain"
	"time"
)

type Model struct {
	ID uint `json:"id"`
}

func NewModel(id uint) Model {
	return Model{
		ID: id,
	}
}

func NewModelWithEntity(e domain.Entity) Model {
	return Model{
		ID: e.ID,
	}
}

type CreationModel struct {
	Model

	CreatedAt     time.Time `json:"created_at"`
	CreatorUserID *uint     `json:"creator_user_id"`
}

func NewCreationModelWithEntity(e domain.CreationEntity) CreationModel {
	return CreationModel{
		Model:         NewModelWithEntity(e.Entity),
		CreatedAt:     e.CreatedAt,
		CreatorUserID: e.CreatorUserID,
	}
}

type CreationAuditedModel struct {
	CreationModel

	ModifiedAt     time.Time `json:"modified_at"`
	ModifierUserID *uint     `json:"modifier_user_id"`
}

func NewCreationAuditedModelWithEntity(e domain.CreationAuditedEntity) CreationAuditedModel {
	return CreationAuditedModel{
		CreationModel:  NewCreationModelWithEntity(e.CreationEntity),
		ModifiedAt:     e.ModifiedAt,
		ModifierUserID: e.ModifierUserID,
	}
}

type FullAuditedEntityModel struct {
	CreationAuditedModel

	DeletedAt     *time.Time `json:"deleted_at"`
	DeleterUserID *uint      `json:"deleter_user_id"`
}

func NewFullAuditedModelWithEntity(e domain.FullAuditedEntity) FullAuditedEntityModel {
	return FullAuditedEntityModel{
		CreationAuditedModel: NewCreationAuditedModelWithEntity(e.CreationAuditedEntity),
		DeletedAt:            e.DeletedAt,
		DeleterUserID:        e.DeleterUserID,
	}
}

type QueryResult struct {
	Records  interface{} `json:"records"`
	Total    int64       `json:"total"`
	Page     int         `json:"page,omitempty"`
	PageSize int         `json:"page_size,omitempty"`
}

func NewQueryResult(records interface{}, total int64, limit int, offset int) *QueryResult {
	return &QueryResult{
		Records:  records,
		Total:    total,
		Page:     offset/limit + 1,
		PageSize: limit,
	}
}
