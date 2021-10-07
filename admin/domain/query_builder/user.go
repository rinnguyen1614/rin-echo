package query_builder

import (
	"fmt"
	"rin-echo/admin/domain"
	"rin-echo/common/query"

	"gorm.io/gorm"
)

type UserQueryBuilder struct {
	query.QueryBuilder
}

func NewUserQueryBuilder(db *gorm.DB) UserQueryBuilder {
	qbuilder, err := query.NewQueryBuilder(db, &domain.User{})
	if err != nil {
		panic(err)
	}

	return UserQueryBuilder{
		QueryBuilder: qbuilder,
	}
}

func (q *UserQueryBuilder) WhereID(id uint) {
	q.SetCondition("id", id)
}

func (q *UserQueryBuilder) WhereIDsIn(ids []uint) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "id", "IN"), ids)
}

func (q *UserQueryBuilder) WhereIDsNotIn(ids []uint) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "id", "NOT IN"), ids)
}
