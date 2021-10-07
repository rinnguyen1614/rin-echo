package query_builder

import (
	"fmt"
	"rin-echo/admin/domain"
	"rin-echo/common/query"

	"gorm.io/gorm"
)

type RoleQueryBuilder struct {
	query.QueryBuilder
}

func NewRoleQueryBuilder(db *gorm.DB) RoleQueryBuilder {
	qbuilder, err := query.NewQueryBuilder(db, &domain.Role{})
	if err != nil {
		panic(err)
	}

	return RoleQueryBuilder{
		QueryBuilder: qbuilder,
	}
}

func (q *RoleQueryBuilder) WhereID(id uint) {
	q.SetCondition("id", id)
}

func (q *RoleQueryBuilder) WhereIDsIn(ids []uint) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "id", "IN"), ids)
}

func (q *RoleQueryBuilder) WhereIDsNotIn(ids []uint) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "id", "NOT IN"), ids)
}
