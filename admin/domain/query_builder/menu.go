package query_builder

import (
	"fmt"
	"rin-echo/admin/domain"

	uow "rin-echo/common/uow"

	iuow "rin-echo/common/uow/interfaces"
)

type MenuQueryBuilder struct {
	iuow.QueryBuilder
}

func NewMenuQueryBuilder() MenuQueryBuilder {
	qbuilder, err := uow.NewQueryBuilder(&domain.Menu{})
	if err != nil {
		panic(err)
	}

	return MenuQueryBuilder{
		QueryBuilder: qbuilder,
	}
}

func (q *MenuQueryBuilder) WhereID(id uint) {
	q.SetCondition("id", id)
}

func (q *MenuQueryBuilder) WhereIDsIn(ids []uint) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "id", "IN"), ids)
}

func (q *MenuQueryBuilder) WhereIDsNotIn(ids []uint) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "id", "NOT IN"), ids)
}

func (q *MenuQueryBuilder) WhereSlug(slug string) {
	q.SetCondition("slug", slug)
}

func (q *MenuQueryBuilder) WhereSlugsIn(slugs []string) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "slug", "IN"), slugs)
}

func (q *MenuQueryBuilder) WhereSlugsNotIn(slugs []string) {
	q.SetCondition(fmt.Sprintf("%v %v ?", "slug", "NOT IN"), slugs)
}
