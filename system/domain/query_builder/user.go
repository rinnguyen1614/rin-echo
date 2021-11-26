package query_builder

import (
	"fmt"
	uow "rin-echo/common/uow"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/system/domain"
)

type UserQueryBuilder struct {
	iuow.QueryBuilder
}

func NewUserQueryBuilder() UserQueryBuilder {
	qbuilder, err := uow.NewQueryBuilder(&domain.User{})
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