package query_builder

import (
	"fmt"

	uow "github.com/rinnguyen1614/rin-echo/internal/core/uow"

	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"

	"github.com/rinnguyen1614/rin-echo/internal/system/domain"
)

type MenuQueryBuilder struct {
	iuow.QueryBuilder
}

func NewMenuQueryBuilder() MenuQueryBuilder {
	qbuilder, err := uow.NewQueryBuilder(&domain.Menu{})
	if err != nil {
		panic(err)
	}

	qbuilder.SetOrder("parent_id", "asc NULLS FIRST")
	qbuilder.SetOrder("id", "asc")

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
