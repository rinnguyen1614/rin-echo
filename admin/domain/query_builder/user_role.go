package query_builder

import (
	"fmt"
	"rin-echo/admin/domain"
	"rin-echo/common/query"

	"gorm.io/gorm"
)

type UserRoleQueryBuilder struct {
	query.QueryBuilder
}

func NewUserRoleQueryBuilder(db *gorm.DB) UserRoleQueryBuilder {
	qbuilder, err := query.NewQueryBuilder(db, &domain.UserRole{})
	if err != nil {
		panic(err)
	}

	return UserRoleQueryBuilder{
		QueryBuilder: qbuilder,
	}
}

func (q *UserRoleQueryBuilder) WhereID(id uint) {
	q.SetCondition("id", id)
}

func (q *UserRoleQueryBuilder) WhereIDsIn(ids []uint) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "id", "IN"), ids)
}

func (q *UserRoleQueryBuilder) WhereIDsNotIn(ids []uint) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "id", "NOT IN"), ids)
}
