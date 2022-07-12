package query_builder

import (
	"fmt"
	"rin-echo/system/domain"

	uow "github.com/rinnguyen1614/rin-echo-core/uow"
	iuow "github.com/rinnguyen1614/rin-echo-core/uow/interfaces"
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
