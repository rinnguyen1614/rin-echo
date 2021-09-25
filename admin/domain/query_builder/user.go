package query_builder

import (
	"fmt"
	"rin-echo/admin/domain"
	"rin-echo/common/query"
)

type UserQueryBuilder struct {
	query.QueryBuilder
}

func NewUserQueryBuilder() UserQueryBuilder {
	qbuilder, err := query.NewQueryBuilder(&domain.User{})
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

func (q *UserQueryBuilder) WhereIDIn(id uint) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "id", "IN"), id)
}

func (q *UserQueryBuilder) WhereIDNotIn(id uint) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "id", "NOT IN"), id)
}
