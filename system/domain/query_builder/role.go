package query_builder

import (
	"fmt"
	"rin-echo/system/domain"

	uow "rin-echo/common/uow"

	iuow "rin-echo/common/uow/interfaces"
)

type RoleQueryBuilder struct {
	iuow.QueryBuilder
}

func NewRoleQueryBuilder() RoleQueryBuilder {
	qbuilder, err := uow.NewQueryBuilder(&domain.Role{})
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

func (q *RoleQueryBuilder) WhereSlug(slug string) {
	q.SetCondition("slug", slug)
}

func (q *RoleQueryBuilder) WhereSlugsIn(slugs []string) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "slug", "IN"), slugs)
}

func (q *RoleQueryBuilder) WhereSlugsNotIn(slugs []string) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "slug", "NOT IN"), slugs)
}
