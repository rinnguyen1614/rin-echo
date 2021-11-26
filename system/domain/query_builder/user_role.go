package query_builder

import (
	"fmt"
	uow "rin-echo/common/uow"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/system/domain"
)

type UserRoleQueryBuilder struct {
	iuow.QueryBuilder
}

func NewUserRoleQueryBuilder() UserRoleQueryBuilder {
	qbuilder, err := uow.NewQueryBuilder(&domain.UserRole{})
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
