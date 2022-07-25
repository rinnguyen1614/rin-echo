package domain

import (
	"time"
)

type (
	Entity struct {
		ID uint
		// UUID utils.UUID `gorm:"unique,type:uuid;default:uuid_generate_v4();autoincrement"`
	}

	Entities []*Entity
)

func (e Entities) IDs() []uint {
	var ids []uint
	for _, v := range e {
		ids = append(ids, v.ID)
	}
	return ids
}

type (
	CreationEntity struct {
		Entity

		CreatedAt     time.Time
		CreatorUserID *uint
	}

	CreationAuditedEntity struct {
		CreationEntity

		ModifiedAt     time.Time
		ModifierUserID *uint
	}

	FullAuditedEntity struct {
		CreationAuditedEntity

		DeletedAt     *time.Time
		DeleterUserID *uint
	}
)

func (e *CreationEntity) BeforeCreate() {

}
